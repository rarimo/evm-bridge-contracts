import { expect } from "chai";
import { Bridge, ERC1155MintableBurnable } from "@ethers-v5";
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
import { Bundle } from "@/test/utils/types";

describe("ERC1155Handler", () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let SECOND: SignerWithAddress;

  let bridge: Bridge;
  let erc1155: ERC1155MintableBurnable;

  beforeEach(async () => {
    [OWNER, SECOND] = await ethers.getSigners();

    const Bridge = await ethers.getContractFactory("Bridge");
    const BundleExecutorImplementation = await ethers.getContractFactory("BundleExecutorImplementation");
    const ERC1155MintableBurnable = await ethers.getContractFactory("ERC1155MintableBurnable");

    bridge = await Bridge.deploy();
    const bundleImplementation = await BundleExecutorImplementation.deploy(ZERO_ADDR);

    erc1155 = await ERC1155MintableBurnable.deploy(OWNER.address, "");
    await erc1155.mintTo(OWNER.address, 10, wei("10"), "URI");

    await bridge.__Bridge_init(ZERO_ADDR, bundleImplementation.address, CHAIN_NAME, OWNER.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("erc1155", () => {
    describe("#mintTo", () => {
      it("should not mint if caller is not the owner", async () => {
        const tx = erc1155.connect(SECOND).mintTo(OWNER.address, 11, wei("10"), "URI2");

        await expect(tx).to.be.revertedWith("Ownable: caller is not the owner");
      });

      it("should mint if all conditions are met", async () => {
        await erc1155.mintTo(OWNER.address, 11, wei("11"), "URI2");

        expect(await erc1155.balanceOf(OWNER.address, 11)).to.be.eq(wei("11"));
        expect(await erc1155.uri(11)).to.be.eq("URI2");
      });
    });

    describe("#burnFrom", () => {
      it("should not burn by a third party if not approved", async () => {
        const tx = erc1155.connect(SECOND).burnFrom(OWNER.address, 10, wei("10"));

        await expect(tx).to.be.revertedWith("ERC1155MintableBurnable: not approved");
      });

      it("should burn by a third party if approved for all", async () => {
        await erc1155.setApprovalForAll(SECOND.address, true);
        await erc1155.connect(SECOND).burnFrom(OWNER.address, 10, wei("10"));

        expect(await erc1155.balanceOf(OWNER.address, 10)).to.be.eq(0);
      });

      it("should burn own tokens", async () => {
        await erc1155.burnFrom(OWNER.address, 10, wei("10"));

        expect(await erc1155.balanceOf(OWNER.address, 10)).to.be.eq(0);
      });
    });
  });

  describe("handler", () => {
    beforeEach(async () => {
      await erc1155.transferOwnership(bridge.address);
    });

    describe("#depositERC1155", () => {
      it("should revert if the caller is not a facade", async () => {
        const tx = bridge.connect(SECOND).depositERC1155({
          token: erc1155.address,
          tokenId: 10,
          amount: wei("10"),
          bundle: RANDOM_BUNDLE,
          network: CHAIN_NAME,
          receiver: RECEIVER,
          isWrapped: false,
        });

        await expect(tx).to.be.revertedWith("Bundler: not a facade");
      });

      it("should emit the deposit event with proper params", async () => {
        const tx = bridge.depositERC1155({
          token: erc1155.address,
          tokenId: 10,
          amount: wei("10"),
          bundle: RANDOM_BUNDLE,
          network: CHAIN_NAME,
          receiver: RECEIVER,
          isWrapped: false,
        });

        await expect(tx)
          .to.emit(bridge, "DepositedERC1155")
          .withArgs(
            erc1155.address,
            10,
            wei("10"),
            encodeSalt(RANDOM_BUNDLE.salt, OWNER),
            RANDOM_BUNDLE.bundle,
            CHAIN_NAME,
            RECEIVER,
            false
          );
      });
    });

    describe("#withdrawERC1155", () => {
      context("if deposited", () => {
        beforeEach(async () => {
          await erc1155.safeTransferFrom(OWNER.address, bridge.address, 10, wei("10"), "0x");
        });

        it("should revert if the caller is not a facade", async () => {
          const tx = bridge.connect(SECOND).withdrawERC1155({
            token: erc1155.address,
            tokenId: 10,
            tokenURI: "URI",
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
          const tx = bridge.withdrawERC1155({
            token: ZERO_ADDR,
            tokenId: 10,
            tokenURI: "URI",
            amount: wei("10"),
            bundle: EMPTY_BUNDLE,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
            proof: "0x",
            isWrapped: false,
          });

          await expect(tx).to.be.revertedWith("ERC1155Handler: zero token");
        });

        it("should not withdraw if zero receiver", async () => {
          const tx = bridge.withdrawERC1155({
            token: erc1155.address,
            tokenId: 10,
            tokenURI: "URI",
            amount: wei("10"),
            bundle: EMPTY_BUNDLE,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: ZERO_ADDR,
            proof: "0x",
            isWrapped: false,
          });

          await expect(tx).to.be.revertedWith("ERC1155Handler: zero receiver");
        });

        it("should withdraw properly", async () => {
          const tx = await bridge.withdrawERC1155({
            token: erc1155.address,
            tokenId: 10,
            tokenURI: "URI",
            amount: wei("10"),
            bundle: EMPTY_BUNDLE,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
            proof: "0x",
            isWrapped: false,
          });

          expect(await erc1155.balanceOf(bridge.address, 10)).to.eq(0);
          expect(await erc1155.balanceOf(OWNER.address, 10)).to.eq(wei("10"));

          await expect(tx)
            .to.emit(bridge, "WithdrawnERC1155")
            .withArgs(
              erc1155.address,
              10,
              "URI",
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
          await erc1155.burnFrom(OWNER.address, 10, wei("10"));
        });

        it("should withdraw properly", async () => {
          const tx = await bridge.withdrawERC1155({
            token: erc1155.address,
            tokenId: 10,
            tokenURI: "URI",
            amount: wei("10"),
            bundle: EMPTY_BUNDLE,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
            proof: "0x",
            isWrapped: true,
          });

          expect(await erc1155.balanceOf(bridge.address, 10)).to.eq(0);
          expect(await erc1155.balanceOf(OWNER.address, 10)).to.eq(wei("10"));

          await expect(tx)
            .to.emit(bridge, "WithdrawnERC1155")
            .withArgs(
              erc1155.address,
              10,
              "URI",
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

    describe("#withdrawERC1155Bundle", () => {
      let bundleApprove: Bundle;
      let proxyAddress: string;

      beforeEach(async () => {
        bundleApprove = {
          bundle: ethers.utils.defaultAbiCoder.encode(
            ["address[]", "uint256[]", "bytes[]"],
            [[erc1155.address], [0], [erc1155.interface.encodeFunctionData("setApprovalForAll", [OWNER.address, true])]]
          ),
          salt: RANDOM_SALT,
        };
        proxyAddress = await bridge.determineProxyAddress(bundleApprove.salt);
      });

      context("if deposited", () => {
        beforeEach(async () => {
          await erc1155.safeTransferFrom(OWNER.address, bridge.address, 10, wei("10"), "0x");
        });

        it("should revert if the caller is not a facade", async () => {
          const tx = bridge.connect(SECOND).withdrawERC1155Bundle({
            token: erc1155.address,
            tokenId: 10,
            tokenURI: "URI",
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
          const tx = await bridge.withdrawERC1155Bundle({
            token: erc1155.address,
            tokenId: 10,
            tokenURI: "URI",
            amount: wei("10"),
            bundle: bundleApprove,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
            proof: "0x",
            isWrapped: false,
          });

          expect(await erc1155.balanceOf(bridge.address, 10)).to.eq(0);
          expect(await erc1155.balanceOf(proxyAddress, 10)).to.eq(wei("10"));
          expect(await erc1155.isApprovedForAll(proxyAddress, OWNER.address)).to.be.true;

          await expect(tx)
            .to.emit(bridge, "WithdrawnERC1155")
            .withArgs(
              erc1155.address,
              10,
              "URI",
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
          await erc1155.burnFrom(OWNER.address, 10, wei("10"));
        });

        it("should withdraw bundle properly", async () => {
          const tx = await bridge.withdrawERC1155Bundle({
            token: erc1155.address,
            tokenId: 10,
            tokenURI: "URI",
            amount: wei("10"),
            bundle: bundleApprove,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
            proof: "0x",
            isWrapped: true,
          });

          expect(await erc1155.balanceOf(bridge.address, 10)).to.eq(0);
          expect(await erc1155.balanceOf(proxyAddress, 10)).to.eq(wei("10"));
          expect(await erc1155.isApprovedForAll(proxyAddress, OWNER.address)).to.be.true;
          expect(await erc1155.uri(10)).to.be.eq("URI");

          await expect(tx)
            .to.emit(bridge, "WithdrawnERC1155")
            .withArgs(
              erc1155.address,
              10,
              "URI",
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
