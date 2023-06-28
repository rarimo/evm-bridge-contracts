import { expect } from "chai";
import { Bridge, SBTMintableBurnable } from "@ethers-v5";
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

describe("SBTHandler", () => {
  const reverter = new Reverter();

  let OWNER: SignerWithAddress;
  let SECOND: SignerWithAddress;

  let bridge: Bridge;
  let sbt: SBTMintableBurnable;

  beforeEach(async () => {
    [OWNER, SECOND] = await ethers.getSigners();

    const Bridge = await ethers.getContractFactory("Bridge");
    const BundleExecutorImplementation = await ethers.getContractFactory("BundleExecutorImplementation");
    const SBTMintableBurnable = await ethers.getContractFactory("SBTMintableBurnable");

    bridge = await Bridge.deploy();
    const bundleImplementation = await BundleExecutorImplementation.deploy(ZERO_ADDR);

    sbt = await SBTMintableBurnable.deploy("SBTMock", "SBTMock", OWNER.address, "");
    await sbt.attestTo(OWNER.address, 10, "URI");

    await bridge.__Bridge_init(ZERO_ADDR, bundleImplementation.address, CHAIN_NAME, OWNER.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("sbt", () => {
    describe("#attestTo", () => {
      it("should not attest if caller is not the owner", async () => {
        const tx = sbt.connect(SECOND).attestTo(OWNER.address, 1, "URI2");

        await expect(tx).to.be.revertedWith("Ownable: caller is not the owner");
      });

      it("should not attest if receiver already has a token", async () => {
        const tx = sbt.attestTo(OWNER.address, 1, "URI2");

        await expect(tx).to.be.revertedWith("SBTMintableBurnable: receiver already has a token");
      });

      it("should attest if all conditions are met", async () => {
        const tx = sbt.attestTo(SECOND.address, 1, "URI2");

        await expect(tx).to.changeTokenBalance(sbt, SECOND, 1);
        await expect(tx).to.emit(sbt, "Locked").withArgs(1);

        expect(await sbt.tokenURI(1)).to.eq("URI2");
      });
    });

    describe("#burn", () => {
      it("should not burn own token if user doesn't have a token", async () => {
        const tx = sbt.connect(SECOND).burn();

        await expect(tx).to.be.revertedWith("SBTMintableBurnable: user doesn't have a token");
      });

      it("should burn own tokens if all conditions are met", async () => {
        const tx = sbt.burn();

        await expect(tx).to.changeTokenBalance(sbt, OWNER, -1);
      });
    });

    describe("#locked", () => {
      it("should revert if token is assigned to zero address", async () => {
        const tx = sbt.locked(1);

        await expect(tx).to.be.revertedWith("SBTMintableBurnable: token is assigned to zero address");
      });

      it("should be locked if all conditions are met", async () => {
        expect(await sbt.locked(10)).to.be.true;
      });
    });

    describe("#supportsInterface", () => {
      it("should support correct interfaces", async () => {
        expect(await sbt.supportsInterface("0x780e9d63")).to.be.false; // IERC721Enumerable
        expect(await sbt.supportsInterface("0x80ac58cd")).to.be.true; // IERC721
        expect(await sbt.supportsInterface("0x01ffc9a7")).to.be.true; // IERC165
        expect(await sbt.supportsInterface("0xb45a3c0e")).to.be.true; // IERC5192
        expect(await sbt.supportsInterface("0xf6809be0")).to.be.true; // ISBT
        expect(await sbt.supportsInterface("0x00000000")).to.be.false;
      });
    });

    describe("#transfer", async () => {
      it("should not be transferable", async () => {
        const tx = sbt.transferFrom(OWNER.address, SECOND.address, 10);

        expect(tx).to.be.revertedWith("SBTMintableBurnable: not transferable");
      });
    });
  });

  describe("handler", () => {
    beforeEach(async () => {
      await sbt.transferOwnership(bridge.address);
    });

    describe("#depositSBT", () => {
      it("should revert if the caller is not a facade", async () => {
        const tx = bridge.connect(SECOND).depositSBT({
          token: sbt.address,
          tokenId: 10,
          bundle: RANDOM_BUNDLE,
          network: CHAIN_NAME,
          receiver: RECEIVER,
        });

        await expect(tx).to.be.revertedWith("Bundler: not a facade");
      });

      it("should emit the deposit event with proper params", async () => {
        const tx = bridge.depositSBT({
          token: sbt.address,
          tokenId: 10,
          bundle: RANDOM_BUNDLE,
          network: CHAIN_NAME,
          receiver: RECEIVER,
        });

        await expect(tx)
          .to.emit(bridge, "DepositedSBT")
          .withArgs(sbt.address, 10, encodeSalt(RANDOM_BUNDLE.salt, OWNER), RANDOM_BUNDLE.bundle, CHAIN_NAME, RECEIVER);
      });
    });

    describe("#withdrawSBT", () => {
      beforeEach(async () => {
        await sbt.burn();
      });

      it("should revert if the caller is not a facade", async () => {
        const tx = bridge.connect(SECOND).withdrawSBT({
          token: sbt.address,
          tokenId: 10,
          tokenURI: "URI",
          bundle: EMPTY_BUNDLE,
          originHash: RANDOM_ORIGIN_HASH,
          receiver: OWNER.address,
          proof: "0x",
        });

        await expect(tx).to.be.revertedWith("Bundler: not a facade");
      });

      it("should not withdraw if zero token", async () => {
        const tx = bridge.withdrawSBT({
          token: ZERO_ADDR,
          tokenId: 10,
          tokenURI: "URI",
          bundle: EMPTY_BUNDLE,
          originHash: RANDOM_ORIGIN_HASH,
          receiver: OWNER.address,
          proof: "0x",
        });

        await expect(tx).to.be.revertedWith("SBTHandler: zero token");
      });

      it("should not withdraw if zero receiver", async () => {
        const tx = bridge.withdrawSBT({
          token: sbt.address,
          tokenId: 10,
          tokenURI: "URI",
          bundle: EMPTY_BUNDLE,
          originHash: RANDOM_ORIGIN_HASH,
          receiver: ZERO_ADDR,
          proof: "0x",
        });

        await expect(tx).to.be.revertedWith("SBTHandler: zero receiver");
      });

      it("should withdraw properly if all conditions are met", async () => {
        const tx = bridge.withdrawSBT({
          token: sbt.address,
          tokenId: 10,
          tokenURI: "URI",
          bundle: EMPTY_BUNDLE,
          originHash: RANDOM_ORIGIN_HASH,
          receiver: OWNER.address,
          proof: "0x",
        });

        await expect(tx).to.changeTokenBalances(sbt, [bridge, OWNER], [0, 1]);
        expect(await sbt.tokenURI(10)).to.eq("URI");
      });
    });

    describe("#withdrawSBTBundle", () => {
      let bundleApprove: Bundle;
      let proxyAddress: string;

      beforeEach(async () => {
        await sbt.burn();

        bundleApprove = {
          bundle: ethers.utils.defaultAbiCoder.encode(
            ["address[]", "uint256[]", "bytes[]"],
            [[sbt.address], [0], [sbt.interface.encodeFunctionData("approve", [OWNER.address, 10])]]
          ),
          salt: RANDOM_SALT,
        };
        proxyAddress = await bridge.determineProxyAddress(bundleApprove.salt);
      });

      it("should revert if the caller is not a facade", async () => {
        const tx = bridge.connect(SECOND).withdrawSBTBundle({
          token: sbt.address,
          tokenId: 10,
          tokenURI: "URI",
          bundle: bundleApprove,
          originHash: RANDOM_ORIGIN_HASH,
          receiver: OWNER.address,
          proof: "0x",
        });

        await expect(tx).to.be.revertedWith("Bundler: not a facade");
      });

      it("should withdraw bundle properly if all conditions are met", async () => {
        const tx = bridge.withdrawSBTBundle({
          token: sbt.address,
          tokenId: 10,
          tokenURI: "URI",
          bundle: bundleApprove,
          originHash: RANDOM_ORIGIN_HASH,
          receiver: OWNER.address,
          proof: "0x",
        });

        await expect(tx).to.changeTokenBalances(sbt, [bridge, proxyAddress, OWNER], [0, 1, 0]);
        expect(await sbt.getApproved(10)).to.eq(OWNER.address);
        expect(await sbt.tokenURI(10)).to.eq("URI");
      });
    });
  });
});
