const { accounts, wei } = require("../../scripts/helpers/utils");
const { rawSign } = require("../helpers/signer");
const { OWNER_PRIVATE_KEY } = require("../helpers/keys");
const truffleAssert = require("truffle-assertions");

const ERC1967Proxy = artifacts.require("ERC1967ProxyMock");
const Bridge = artifacts.require("Bridge");
const BundleImpl = artifacts.require("BundleExecutorImplementation");

Bridge.numberFormat = "BigNumber";
ERC1967Proxy.numberFormat = "BigNumber";
BundleImpl.numberFormat = "BigNumber";

describe("Upgradeable", () => {
  const chainName = "ethereum";

  let OWNER;
  let SECOND;

  let bridge;
  let newBridge;

  let bundleImpl;

  let proxy;
  let proxyBridge;

  before("setup", async () => {
    OWNER = await accounts(0);
    SECOND = await accounts(1);
  });

  beforeEach("setup", async () => {
    bundleImpl = await BundleImpl.new();
    bridge = await Bridge.new();
    newBridge = await Bridge.new();
    proxy = await ERC1967Proxy.new(bridge.address, []);
    proxyBridge = await Bridge.at(proxy.address);

    await proxyBridge.__Bridge_init(OWNER, bundleImpl.address, chainName);
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
    const hashToSign = web3.utils.soliditySha3(
      { value: newBridge.address, type: "address" },
      { value: chainName, type: "string" },
      { value: "0", type: "uint256" },
      { value: proxyBridge.address, type: "address" }
    );

    const signature = rawSign(hashToSign, OWNER_PRIVATE_KEY);

    await truffleAssert.passes(proxyBridge.upgradeToWithSig(newBridge.address, signature));
  });

  it("should fail upgrade if nonce is wrong", async () => {
    const hashToSign = web3.utils.soliditySha3(
      { value: newBridge.address, type: "address" },
      { value: chainName, type: "string" },
      { value: "1", type: "uint256" },
      { value: proxyBridge.address, type: "address" }
    );

    const signature = rawSign(hashToSign, OWNER_PRIVATE_KEY);

    await truffleAssert.reverts(
      proxyBridge.upgradeToWithSig(newBridge.address, signature),
      "Signers: invalid signature"
    );
  });

  it("should not upgrade from nonproxy", async () => {
    const hashToSign = web3.utils.soliditySha3(
      { value: newBridge.address, type: "address" },
      { value: chainName, type: "string" },
      { value: "0", type: "uint256" },
      { value: proxyBridge.address, type: "address" }
    );

    const signature = rawSign(hashToSign, OWNER_PRIVATE_KEY);

    await truffleAssert.reverts(
      bridge.upgradeToWithSig(newBridge.address, signature),
      "Function must be called through delegatecall"
    );
  });

  it("should receive ether through proxy", async () => {
    await truffleAssert.passes(
      await web3.eth.sendTransaction({ from: OWNER, to: proxyBridge.address, value: wei("1") }),
      "pass"
    );
  });
});
