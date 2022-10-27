const { assert } = require("chai");
const { accounts, wei, toBN } = require("../../scripts/helpers/utils");
const { getBytesDoSomething } = require("./bundler-utils");
const truffleAssert = require("truffle-assertions");

const BundleProxy = artifacts.require("BundleExecutorProxy");
const BundleImpl = artifacts.require("BundleExecutorImplementation");
const Bundler = artifacts.require("BundlerMock");
const BundleReceiver = artifacts.require("BundleReceiverMock");

BundleProxy.numberFormat = "BigNumber";
BundleImpl.numberFormat = "BigNumber";
Bundler.numberFormat = "BigNumber";
BundleReceiver.numberFormat = "BigNumber";

describe("Bundler", () => {
  let OWNER;
  let SECOND;

  let bundler;
  let bundleReceiver;
  let bundleImpl;
  let bundleProxy;

  before("setup", async () => {
    OWNER = await accounts(0);
    SECOND = await accounts(1);
  });

  beforeEach("setup", async () => {
    bundleImpl = await BundleImpl.new();
    bundleProxy = await BundleProxy.new(OWNER, OWNER);
    bundleReceiver = await BundleReceiver.new();
    bundler = await Bundler.new();

    await bundler.__BundlerMock_init(bundleImpl.address);
  });

  describe("access", () => {
    it("should not initialize twice", async () => {
      await truffleAssert.reverts(
        bundler.__Bundler_init(bundleImpl.address),
        "Initializable: contract is not initializing"
      );
    });
  });

  describe("proxy address", () => {
    it("proxy address should be different", async () => {
      const salt1 = web3.utils.soliditySha3({ value: 1, type: "uint256" });
      const salt2 = web3.utils.soliditySha3({ value: 2, type: "uint256" });

      let bundleProxy1 = await bundler.determineProxyAddress(salt1);
      let bundleProxy2 = await bundler.determineProxyAddress(salt2);

      assert.notEqual(bundleProxy1, bundleProxy2);
    });

    it("should not be possible to destroy proxy not from bridge", async () => {
      await truffleAssert.reverts(bundleProxy.destroy({ from: SECOND }));
      await truffleAssert.passes(bundleProxy.destroy(), "pass");
    });
  });

  describe("bundling", () => {
    it("should send ether to the receiver", async () => {
      const salt = web3.utils.soliditySha3({ value: 1, type: "uint256" });
      const bundle = web3.eth.abi.encodeParameters(
        ["address[]", "uint256[]", "bytes[]"],
        [[bundleReceiver.address], [wei("1")], ["0x"]]
      );

      const bundleProxy = await bundler.determineProxyAddress(salt);

      await web3.eth.sendTransaction({ from: OWNER, to: bundleProxy, value: wei("100") });

      assert.equal(await web3.eth.getBalance(bundleProxy), wei("100"));
      assert.equal(await web3.eth.getBalance(bundleReceiver.address), "0");
      assert.equal(await web3.eth.getBalance(bundler.address), "0");

      await bundler.bundleUp({ salt: salt, bundle: bundle });

      assert.equal(await web3.eth.getBalance(bundleProxy), "0");
      assert.equal(await web3.eth.getBalance(bundleReceiver.address), wei("1"));
      assert.equal(await web3.eth.getBalance(bundler.address), wei("99"));
    });

    it("should revert because of failure during execution", async () => {
      const salt = web3.utils.soliditySha3({ value: 123123, type: "uint256" });
      const bundle = web3.eth.abi.encodeParameters(
        ["address[]", "uint256[]", "bytes[]"],
        [
          [bundleReceiver.address, bundleReceiver.address],
          [0, 0],
          [getBytesDoSomething(false), getBytesDoSomething(true)],
        ]
      );

      await truffleAssert.reverts(
        bundler.bundleUp({ salt: salt, bundle: bundle }),
        "BundleExecutorImplementation: call reverted"
      );
    });

    it("should revert if nonexisting function is called", async () => {
      const salt = web3.utils.soliditySha3({ value: 321321, type: "uint256" });
      const bundle = web3.eth.abi.encodeParameters(
        ["address[]", "uint256[]", "bytes[]"],
        [[bundler.address], [0], [getBytesDoSomething(false)]]
      );

      await truffleAssert.reverts(
        bundler.bundleUp({ salt: salt, bundle: bundle }),
        "BundleExecutorImplementation: call reverted"
      );
    });

    it("should revert if not enough ether", async () => {
      const salt = web3.utils.soliditySha3({ value: 1, type: "uint256" });
      const bundle = web3.eth.abi.encodeParameters(
        ["address[]", "uint256[]", "bytes[]"],
        [[bundleReceiver.address], [wei("1")], ["0x"]]
      );

      await truffleAssert.reverts(
        bundler.bundleUp({ salt: salt, bundle: bundle }),
        "BundleExecutorImplementation: call reverted"
      );
    });
  });
});
