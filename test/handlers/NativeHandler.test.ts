import { expect } from "chai";
import { Bridge, WrappedNativeMock } from "@ethers-v5";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "hardhat";
import { ZERO_ADDR } from "@/scripts/utils/constants";
import { Reverter } from "@/test/helpers/reverter";
import {
  CHAIN_NAME,
  EMPTY_BUNDLE,
  RANDOM_BUNDLE,
  RANDOM_ORIGIN_HASH,
  RANDOM_SALT,
  RECEIVER,
} from "@/test/utils/constants";
import { encodeSalt } from "@/test/utils/utils";
import { wei } from "@/scripts/utils/utils";
import { setBalance } from "@nomicfoundation/hardhat-network-helpers";
import { Bundle } from "@/test/utils/types";

describe("NativeHandler", () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let SECOND: SignerWithAddress;

  let bridge: Bridge;
  let weth9: WrappedNativeMock;

  beforeEach(async () => {
    [OWNER, SECOND] = await ethers.getSigners();

    const Bridge = await ethers.getContractFactory("Bridge");
    const BundleExecutorImplementation = await ethers.getContractFactory("BundleExecutorImplementation");
    const WrappedNativeMock = await ethers.getContractFactory("WrappedNativeMock");

    bridge = await Bridge.deploy();
    weth9 = await WrappedNativeMock.deploy();
    const bundleImplementation = await BundleExecutorImplementation.deploy(weth9.address);

    await bridge.__Bridge_init(ZERO_ADDR, bundleImplementation.address, CHAIN_NAME, OWNER.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("#depositNative", () => {
    it("should revert if the caller is not a facade", async () => {
      const tx = bridge.connect(SECOND).depositNative(
        {
          amount: wei("10"),
          bundle: RANDOM_BUNDLE,
          network: CHAIN_NAME,
          receiver: RECEIVER,
        },
        { value: wei("10") }
      );

      await expect(tx).to.be.revertedWith("Bundler: not a facade");
    });

    it("should emit the deposit event with proper params", async () => {
      const tx = bridge.depositNative(
        {
          amount: wei("10"),
          bundle: RANDOM_BUNDLE,
          network: CHAIN_NAME,
          receiver: RECEIVER,
        },
        { value: wei("10") }
      );

      await expect(tx)
        .to.emit(bridge, "DepositedNative")
        .withArgs(wei("10"), encodeSalt(RANDOM_BUNDLE.salt, OWNER), RANDOM_BUNDLE.bundle, CHAIN_NAME, RECEIVER);
    });
  });

  context("if deposited", () => {
    beforeEach(async () => {
      await setBalance(bridge.address, wei("10"));
    });

    describe("#withdrawNative", () => {
      it("should revert if the caller is not a facade", async () => {
        const tx = bridge.connect(SECOND).withdrawNative({
          amount: wei("10"),
          bundle: EMPTY_BUNDLE,
          originHash: RANDOM_ORIGIN_HASH,
          receiver: OWNER.address,
          proof: "0x",
        });

        await expect(tx).to.be.revertedWith("Bundler: not a facade");
      });

      it("should not withdraw if zero receiver", async () => {
        const tx = bridge.withdrawNative({
          amount: wei("10"),
          bundle: EMPTY_BUNDLE,
          originHash: RANDOM_ORIGIN_HASH,
          receiver: ZERO_ADDR,
          proof: "0x",
        });

        await expect(tx).to.be.revertedWith("NativeHandler: receiver is zero");
      });

      it("should not withdraw if failed to send eth", async () => {
        const tx = bridge.withdrawNative({
          amount: wei("10"),
          bundle: EMPTY_BUNDLE,
          originHash: RANDOM_ORIGIN_HASH,
          receiver: weth9.address,
          proof: "0x",
        });

        await expect(tx).to.be.revertedWith("NativeHandler: failed to send eth");
      });

      it("should withdraw properly if all conditions are met", async () => {
        const tx = bridge.withdrawNative({
          amount: wei("10"),
          bundle: EMPTY_BUNDLE,
          originHash: RANDOM_ORIGIN_HASH,
          receiver: OWNER.address,
          proof: "0x",
        });

        await expect(tx).to.changeEtherBalances([bridge, OWNER], [-wei("10"), wei("10")]);
      });
    });

    describe("#withdrawNativeBundle", () => {
      let bundleTransfer: Bundle;
      let proxyAddress: string;

      beforeEach(async () => {
        bundleTransfer = {
          bundle: ethers.utils.defaultAbiCoder.encode(
            ["address[]", "uint256[]", "bytes[]"],
            [[OWNER.address], [wei("5")], ["0x"]]
          ),
          salt: RANDOM_SALT,
        };
        proxyAddress = await bridge.determineProxyAddress(bundleTransfer.salt);
      });

      it("should revert if the caller is not a facade", async () => {
        const tx = bridge.connect(SECOND).withdrawNativeBundle({
          amount: wei("10"),
          bundle: bundleTransfer,
          originHash: RANDOM_ORIGIN_HASH,
          receiver: OWNER.address,
          proof: "0x",
        });

        await expect(tx).to.be.revertedWith("Bundler: not a facade");
      });

      it("should withdraw bundle properly if all conditions are met", async () => {
        const tx = bridge.withdrawNativeBundle({
          amount: wei("10"),
          bundle: bundleTransfer,
          originHash: RANDOM_ORIGIN_HASH,
          receiver: OWNER.address,
          proof: "0x",
        });

        await expect(tx).to.changeEtherBalances([bridge, proxyAddress, OWNER], [-wei("10"), 0, wei("5")]);
        await expect(tx).to.changeTokenBalances(weth9, [bridge, proxyAddress, OWNER], [0, wei("5"), 0]);
      });
    });
  });
});
