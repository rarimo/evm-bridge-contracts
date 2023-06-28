import { expect } from "chai";
import { Bridge, ERC721MintableBurnable } from "@ethers-v5";
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
import { Bundle } from "@/test/utils/types";

describe("ERC721Handler", () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let SECOND: SignerWithAddress;

  let bridge: Bridge;
  let erc721: ERC721MintableBurnable;

  beforeEach(async () => {
    [OWNER, SECOND] = await ethers.getSigners();

    const Bridge = await ethers.getContractFactory("Bridge");
    const BundleExecutorImplementation = await ethers.getContractFactory("BundleExecutorImplementation");
    const ERC721MintableBurnable = await ethers.getContractFactory("ERC721MintableBurnable");

    bridge = await Bridge.deploy();
    const bundleImplementation = await BundleExecutorImplementation.deploy(ZERO_ADDR);

    erc721 = await ERC721MintableBurnable.deploy("ERC721Mock", "ERC721Mock", OWNER.address, "");
    await erc721.mintTo(OWNER.address, 10, "URI");

    await bridge.__Bridge_init(ZERO_ADDR, bundleImplementation.address, CHAIN_NAME, OWNER.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("erc721", () => {
    describe("#mintTo", () => {
      it("should not mint if caller is not the owner", async () => {
        const tx = erc721.connect(SECOND).mintTo(OWNER.address, 1, "URI2");

        await expect(tx).to.be.revertedWith("Ownable: caller is not the owner");
      });

      it("should mint if all conditions are met", async () => {
        const tx = erc721.mintTo(OWNER.address, 1, "URI2");

        await expect(tx).to.changeTokenBalance(erc721, OWNER, 1);

        expect(await erc721.tokenURI(1)).to.eq("URI2");
      });
    });

    describe("#burnFrom", () => {
      it("should not burn by a third party if insufficient allowance", async () => {
        const tx = erc721.connect(SECOND).burnFrom(OWNER.address, 10);

        await expect(tx).to.be.revertedWith("ERC721MintableBurnable: not approved");
      });

      it("should not burn if the payer is not the token owner", async () => {
        const tx = erc721.burnFrom(SECOND.address, 10);

        await expect(tx).to.be.revertedWith("ERC721MintableBurnable: not approved");
      });

      it("should burn by a third party if the token is approved", async () => {
        await erc721.approve(SECOND.address, 10);

        const tx = erc721.connect(SECOND).burnFrom(OWNER.address, 10);

        await expect(tx).to.changeTokenBalances(erc721, [OWNER, SECOND], [-1, 0]);
      });

      it("should burn by a third party if approved for all", async () => {
        await erc721.setApprovalForAll(SECOND.address, true);

        const tx = erc721.connect(SECOND).burnFrom(OWNER.address, 10);

        await expect(tx).to.changeTokenBalances(erc721, [OWNER, SECOND], [-1, 0]);
      });

      it("should burn own tokens", async () => {
        const tx = erc721.burnFrom(OWNER.address, 10);

        await expect(tx).to.changeTokenBalance(erc721, OWNER, -1);
      });
    });

    describe("#supportsInterface", () => {
      it("should support correct interfaces", async () => {
        expect(await erc721.supportsInterface("0x780e9d63")).to.be.true; // IERC721Enumerable
        expect(await erc721.supportsInterface("0x80ac58cd")).to.be.true; // IERC721
        expect(await erc721.supportsInterface("0x01ffc9a7")).to.be.true; // IERC165
        expect(await erc721.supportsInterface("0xb45a3c0e")).to.be.false; // IERC5192
        expect(await erc721.supportsInterface("0xf6809be0")).to.be.false; // ISBT
        expect(await erc721.supportsInterface("0x00000000")).to.be.false;
      });
    });
  });

  describe("handler", () => {
    beforeEach(async () => {
      await erc721.transferOwnership(bridge.address);
    });

    describe("#depositERC721", () => {
      it("should revert if the caller is not a facade", async () => {
        const tx = bridge.connect(SECOND).depositERC721({
          token: erc721.address,
          tokenId: 10,
          bundle: RANDOM_BUNDLE,
          network: CHAIN_NAME,
          receiver: RECEIVER,
          isWrapped: false,
        });

        await expect(tx).to.be.revertedWith("Bundler: not a facade");
      });

      it("should emit the deposit event with proper params", async () => {
        const tx = bridge.depositERC721({
          token: erc721.address,
          tokenId: 10,
          bundle: RANDOM_BUNDLE,
          network: CHAIN_NAME,
          receiver: RECEIVER,
          isWrapped: false,
        });

        await expect(tx)
          .to.emit(bridge, "DepositedERC721")
          .withArgs(
            erc721.address,
            10,
            encodeSalt(RANDOM_BUNDLE.salt, OWNER),
            RANDOM_BUNDLE.bundle,
            CHAIN_NAME,
            RECEIVER,
            false
          );
      });
    });

    describe("#withdrawERC721", () => {
      context("if deposited", () => {
        beforeEach(async () => {
          await erc721["safeTransferFrom(address,address,uint256)"](OWNER.address, bridge.address, 10);
        });

        it("should revert if the caller is not a facade", async () => {
          const tx = bridge.connect(SECOND).withdrawERC721({
            token: erc721.address,
            tokenId: 10,
            tokenURI: "URI",
            bundle: EMPTY_BUNDLE,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
            proof: "0x",
            isWrapped: false,
          });

          await expect(tx).to.be.revertedWith("Bundler: not a facade");
        });

        it("should not withdraw if zero token", async () => {
          const tx = bridge.withdrawERC721({
            token: ZERO_ADDR,
            tokenId: 10,
            tokenURI: "URI",
            bundle: EMPTY_BUNDLE,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
            proof: "0x",
            isWrapped: false,
          });

          await expect(tx).to.be.revertedWith("ERC721Handler: zero token");
        });

        it("should not withdraw if zero receiver", async () => {
          const tx = bridge.withdrawERC721({
            token: erc721.address,
            tokenId: 10,
            tokenURI: "URI",
            bundle: EMPTY_BUNDLE,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: ZERO_ADDR,
            proof: "0x",
            isWrapped: false,
          });

          await expect(tx).to.be.revertedWith("ERC721Handler: zero receiver");
        });

        it("should withdraw properly", async () => {
          const tx = bridge.withdrawERC721({
            token: erc721.address,
            tokenId: 10,
            tokenURI: "URI",
            bundle: EMPTY_BUNDLE,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
            proof: "0x",
            isWrapped: false,
          });

          await expect(tx).to.changeTokenBalances(erc721, [bridge, OWNER], [-1, 1]);
        });
      });

      context("if burned", () => {
        beforeEach(async () => {
          await erc721.burnFrom(OWNER.address, 10);
        });

        it("should withdraw properly", async () => {
          const tx = bridge.withdrawERC721({
            token: erc721.address,
            tokenId: 10,
            tokenURI: "URI",
            bundle: EMPTY_BUNDLE,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
            proof: "0x",
            isWrapped: true,
          });

          await expect(tx).to.changeTokenBalances(erc721, [bridge, OWNER], [0, 1]);
          expect(await erc721.tokenURI(10)).to.eq("URI");
        });
      });
    });

    describe("#withdrawERC721Bundle", () => {
      let bundleApprove: Bundle;
      let proxyAddress: string;

      beforeEach(async () => {
        bundleApprove = {
          bundle: ethers.utils.defaultAbiCoder.encode(
            ["address[]", "uint256[]", "bytes[]"],
            [[erc721.address], [0], [erc721.interface.encodeFunctionData("approve", [OWNER.address, 10])]]
          ),
          salt: RANDOM_SALT,
        };
        proxyAddress = await bridge.determineProxyAddress(bundleApprove.salt);
      });

      context("if deposited", () => {
        beforeEach(async () => {
          await erc721["safeTransferFrom(address,address,uint256)"](OWNER.address, bridge.address, 10);
        });

        it("should revert if the caller is not a facade", async () => {
          const tx = bridge.connect(SECOND).withdrawERC721Bundle({
            token: erc721.address,
            tokenId: 10,
            tokenURI: "URI",
            bundle: bundleApprove,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
            proof: "0x",
            isWrapped: false,
          });

          await expect(tx).to.be.revertedWith("Bundler: not a facade");
        });

        it("should withdraw bundle properly", async () => {
          const tx = bridge.withdrawERC721Bundle({
            token: erc721.address,
            tokenId: 10,
            tokenURI: "URI",
            bundle: bundleApprove,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
            proof: "0x",
            isWrapped: false,
          });

          await expect(tx).to.changeTokenBalances(erc721, [bridge, proxyAddress, OWNER], [-1, 1, 0]);
          expect(await erc721.getApproved(10)).to.eq(OWNER.address);
        });
      });

      context("if burned", () => {
        beforeEach(async () => {
          await erc721.burnFrom(OWNER.address, 10);
        });

        it("should withdraw bundle properly", async () => {
          const tx = bridge.withdrawERC721Bundle({
            token: erc721.address,
            tokenId: 10,
            tokenURI: "URI",
            bundle: bundleApprove,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
            proof: "0x",
            isWrapped: true,
          });

          await expect(tx).to.changeTokenBalances(erc721, [bridge, proxyAddress, OWNER], [0, 1, 0]);
          expect(await erc721.getApproved(10)).to.eq(OWNER.address);
          expect(await erc721.tokenURI(10)).to.eq("URI");
        });
      });
    });
  });
});
