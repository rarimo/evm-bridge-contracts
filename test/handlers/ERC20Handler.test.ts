import { expect } from "chai";
import { Bridge, ERC20MintableBurnable } from "@ethers-v5";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "hardhat";
import { wei } from "@/scripts/utils/utils";
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
import { Bundle } from "@/test/utils/types";

describe("ERC20Handler", () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let SECOND: SignerWithAddress;

  let bridge: Bridge;
  let erc20: ERC20MintableBurnable;

  beforeEach(async () => {
    [OWNER, SECOND] = await ethers.getSigners();

    const Bridge = await ethers.getContractFactory("Bridge");
    const BundleExecutorImplementation = await ethers.getContractFactory("BundleExecutorImplementation");
    const ERC20MintableBurnable = await ethers.getContractFactory("ERC20MintableBurnable");

    bridge = await Bridge.deploy();
    const bundleImplementation = await BundleExecutorImplementation.deploy(ZERO_ADDR);

    erc20 = await ERC20MintableBurnable.deploy("ERC20Mock", "ERC20Mock", 18, OWNER.address);
    await erc20.mintTo(OWNER.address, wei("10"));

    await bridge.__Bridge_init(ZERO_ADDR, bundleImplementation.address, CHAIN_NAME, OWNER.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("erc20", () => {
    describe("#mintTo", () => {
      it("should not mint if caller is not the owner", async () => {
        const tx = erc20.connect(SECOND).mintTo(OWNER.address, wei("10"));

        await expect(tx).to.be.revertedWith("Ownable: caller is not the owner");
      });

      it("should mint if all conditions are met", async () => {
        const tx = erc20.mintTo(OWNER.address, wei("10"));

        await expect(tx).to.changeTokenBalance(erc20, OWNER, wei("10"));
      });
    });

    describe("#burnFrom", () => {
      it("should not burn by a third party if insufficient allowance", async () => {
        const tx = erc20.connect(SECOND).burnFrom(OWNER.address, wei("10"));

        await expect(tx).to.be.revertedWith("ERC20: insufficient allowance");
      });

      it("should burn by a third party if all conditions are met", async () => {
        await erc20.approve(SECOND.address, wei("10"));

        const tx = erc20.connect(SECOND).burnFrom(OWNER.address, wei("10"));

        await expect(tx).to.changeTokenBalances(erc20, [OWNER, SECOND], [-wei("10"), 0]);
      });

      it("should burn own tokens", async () => {
        const tx = erc20.burnFrom(OWNER.address, wei("10"));

        await expect(tx).to.changeTokenBalance(erc20, OWNER, -wei("10"));
      });
    });

    describe("#decimals", () => {
      it("should set decimals properly", async () => {
        expect(await erc20.decimals()).to.be.eq(18);
      });
    });
  });

  describe("handler", () => {
    beforeEach(async () => {
      await erc20.transferOwnership(bridge.address);
    });

    describe("#depositERC20", () => {
      it("should revert if the caller is not a facade", async () => {
        const tx = bridge.connect(SECOND).depositERC20({
          token: erc20.address,
          amount: wei("10"),
          bundle: RANDOM_BUNDLE,
          network: CHAIN_NAME,
          receiver: RECEIVER,
          isWrapped: false,
        });

        await expect(tx).to.be.revertedWith("Bundler: not a facade");
      });

      it("should emit the deposit event with proper params", async () => {
        const tx = bridge.depositERC20({
          token: erc20.address,
          amount: wei("10"),
          bundle: RANDOM_BUNDLE,
          network: CHAIN_NAME,
          receiver: RECEIVER,
          isWrapped: false,
        });

        await expect(tx)
          .to.emit(bridge, "DepositedERC20")
          .withArgs(
            erc20.address,
            wei("10"),
            encodeSalt(RANDOM_BUNDLE.salt, OWNER),
            RANDOM_BUNDLE.bundle,
            CHAIN_NAME,
            RECEIVER,
            false
          );
      });
    });

    describe("#withdrawERC20", () => {
      context("if deposited", () => {
        beforeEach(async () => {
          await erc20.transfer(bridge.address, wei("10"));
        });

        it("should revert if the caller is not a facade", async () => {
          const tx = bridge.connect(SECOND).withdrawERC20({
            token: erc20.address,
            amount: wei("10"),
            bundle: EMPTY_BUNDLE,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
            proof: "0x",
            isWrapped: false,
          });

          await expect(tx).to.be.revertedWith("Bundler: not a facade");
        });

        it("should not withdraw if zero token", async () => {
          const tx = bridge.withdrawERC20({
            token: ZERO_ADDR,
            amount: wei("10"),
            bundle: EMPTY_BUNDLE,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
            proof: "0x",
            isWrapped: false,
          });

          await expect(tx).to.be.revertedWith("ERC20Handler: zero token");
        });

        it("should not withdraw if zero receiver", async () => {
          const tx = bridge.withdrawERC20({
            token: erc20.address,
            amount: wei("10"),
            bundle: EMPTY_BUNDLE,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: ZERO_ADDR,
            proof: "0x",
            isWrapped: false,
          });

          await expect(tx).to.be.revertedWith("ERC20Handler: zero receiver");
        });

        it("should withdraw properly", async () => {
          const tx = bridge.withdrawERC20({
            token: erc20.address,
            amount: wei("10"),
            bundle: EMPTY_BUNDLE,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
            proof: "0x",
            isWrapped: false,
          });

          await expect(tx).to.changeTokenBalances(erc20, [bridge, OWNER], [-wei("10"), wei("10")]);
          await expect(tx)
            .to.emit(bridge, "WithdrawnERC20")
            .withArgs(
              erc20.address,
              wei("10"),
              EMPTY_BUNDLE.salt,
              EMPTY_BUNDLE.bundle,
              RANDOM_ORIGIN_HASH,
              OWNER.address,
              "0x",
              false
            );
        });
      });

      context("if burned", () => {
        beforeEach(async () => {
          await erc20.burnFrom(OWNER.address, wei("10"));
        });

        it("should withdraw properly", async () => {
          const tx = bridge.withdrawERC20({
            token: erc20.address,
            amount: wei("10"),
            bundle: EMPTY_BUNDLE,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
            proof: "0x",
            isWrapped: true,
          });

          await expect(tx).to.changeTokenBalances(erc20, [bridge, OWNER], [0, wei("10")]);
          await expect(tx)
            .to.emit(bridge, "WithdrawnERC20")
            .withArgs(
              erc20.address,
              wei("10"),
              EMPTY_BUNDLE.salt,
              EMPTY_BUNDLE.bundle,
              RANDOM_ORIGIN_HASH,
              OWNER.address,
              "0x",
              true
            );
        });
      });
    });

    describe("#withdrawERC20Bundle", () => {
      let bundleApprove: Bundle;
      let proxyAddress: string;

      beforeEach(async () => {
        bundleApprove = {
          bundle: ethers.utils.defaultAbiCoder.encode(
            ["address[]", "uint256[]", "bytes[]"],
            [[erc20.address], [0], [erc20.interface.encodeFunctionData("approve", [OWNER.address, wei("5")])]]
          ),
          salt: RANDOM_SALT,
        };
        proxyAddress = await bridge.determineProxyAddress(bundleApprove.salt);
      });

      context("if deposited", () => {
        beforeEach(async () => {
          await erc20.transfer(bridge.address, wei("10"));
        });

        it("should revert if the caller is not a facade", async () => {
          const tx = bridge.connect(SECOND).withdrawERC20Bundle({
            token: erc20.address,
            amount: wei("10"),
            bundle: bundleApprove,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
            proof: "0x",
            isWrapped: false,
          });

          await expect(tx).to.be.revertedWith("Bundler: not a facade");
        });

        it("should withdraw bundle properly", async () => {
          const tx = bridge.withdrawERC20Bundle({
            token: erc20.address,
            amount: wei("10"),
            bundle: bundleApprove,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
            proof: "0x",
            isWrapped: false,
          });

          await expect(tx).to.changeTokenBalances(erc20, [bridge, proxyAddress, OWNER], [-wei("10"), wei("10"), 0]);
          expect(await erc20.allowance(proxyAddress, OWNER.address)).to.eq(wei("5"));

          await expect(tx)
            .to.emit(bridge, "WithdrawnERC20")
            .withArgs(
              erc20.address,
              wei("10"),
              bundleApprove.salt,
              bundleApprove.bundle,
              RANDOM_ORIGIN_HASH,
              proxyAddress,
              "0x",
              false
            );
        });
      });

      context("if burned", () => {
        beforeEach(async () => {
          await erc20.burnFrom(OWNER.address, wei("10"));
        });

        it("should withdraw bundle properly", async () => {
          const tx = bridge.withdrawERC20Bundle({
            token: erc20.address,
            amount: wei("10"),
            bundle: bundleApprove,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
            proof: "0x",
            isWrapped: true,
          });

          await expect(tx).to.changeTokenBalances(erc20, [bridge, proxyAddress, OWNER], [0, wei("10"), 0]);
          expect(await erc20.allowance(proxyAddress, OWNER.address)).to.eq(wei("5"));

          await expect(tx)
            .to.emit(bridge, "WithdrawnERC20")
            .withArgs(
              erc20.address,
              wei("10"),
              bundleApprove.salt,
              bundleApprove.bundle,
              RANDOM_ORIGIN_HASH,
              proxyAddress,
              "0x",
              true
            );
        });
      });
    });
  });
});
