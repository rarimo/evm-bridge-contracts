const { accounts } = require("../../scripts/helpers/utils");
const truffleAssert = require("truffle-assertions");
const ethSigUtil = require("@metamask/eth-sig-util");

const ERC1967Proxy = artifacts.require("ERC1967ProxyMock");
const Bridge = artifacts.require("Bridge");

Bridge.numberFormat = "BigNumber";
ERC1967Proxy.numberFormat = "BigNumber";

const OWNER_PRIVATE_KEY = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80";

describe("Upgradeable", () => {
  let OWNER;
  let SECOND;

  let bridge;
  let newBridge;
  let proxy;
  let proxyBridge;

  before("setup", async () => {
    OWNER = await accounts(0);
    SECOND = await accounts(1);
  });

  beforeEach("setup", async () => {
    bridge = await Bridge.new();
    newBridge = await Bridge.new();
    proxy = await ERC1967Proxy.new(bridge.address, []);
    proxyBridge = await Bridge.at(proxy.address);

    await proxyBridge.__Bridge_init(OWNER, "1");
  });

  it("should fail classical upgrade", async () => {
    await truffleAssert.reverts(proxyBridge.upgradeTo(newBridge.address), "Bridge: this upgrade method is off");
  });

  it("should fail classical upgrade with data", async () => {
    await truffleAssert.reverts(
      proxyBridge.upgradeToAndCall(newBridge.address, []),
      "Bridge: this upgrade method is off"
    );
  });

  it("should upgrade if signature is ok", async () => {
    const ownerKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");

    const hashToSign = web3.utils.soliditySha3(
      { value: newBridge.address, type: "address" },
      { value: "31337", type: "uint256" },
      { value: "0", type: "uint256" },
      { value: proxyBridge.address, type: "address" }
    );

    const signature = ethSigUtil.personalSign({ privateKey: ownerKey, data: hashToSign });

    await truffleAssert.passes(proxyBridge.upgradeToWithSig(newBridge.address, signature));
  });

  it("should fail upgrade if nonce is wrong", async () => {
    const ownerKey = Buffer.from(OWNER_PRIVATE_KEY, "hex");

    const hashToSign = web3.utils.soliditySha3(
      { value: newBridge.address, type: "address" },
      { value: "31337", type: "uint256" },
      { value: "1", type: "uint256" },
      { value: proxyBridge.address, type: "address" }
    );

    const signature = ethSigUtil.personalSign({ privateKey: ownerKey, data: hashToSign });

    await truffleAssert.reverts(
      proxyBridge.upgradeToWithSig(newBridge.address, signature),
      "Signers: invalid signature"
    );
  });
});
