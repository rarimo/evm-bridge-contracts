import { expect } from "chai";
import { Bridge, ERC20MintableBurnable, FeeManagerMock } from "@ethers-v5";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "hardhat";
import { wei } from "@/scripts/utils/utils";
import { CHAIN_NAME, COMMISSION_ADDRESS, ETHEREUM_ADDRESS } from "@/test/utils/constants";
import { Reverter } from "@/test/helpers/reverter";
import { SignHelper } from "@/test/utils/signature";
import { Wallet } from "ethers";
import { ZERO_ADDR } from "@/scripts/utils/constants";
import { setBalance } from "@nomicfoundation/hardhat-network-helpers";

describe("FeeManager", () => {
  const reverter = new Reverter();

  let signHelper: SignHelper;

  let OWNER: SignerWithAddress;
  let SIGNER: Wallet;

  let feeManager: FeeManagerMock;
  let bridge: Bridge;
  let feeToken: ERC20MintableBurnable;
  let feeTokens: Array<string>;
  let feeAmounts: Array<bigint>;

  before("setup", async () => {
    [OWNER] = await ethers.getSigners();
    SIGNER = ethers.Wallet.createRandom();

    const ERC1967ProxyMock = await ethers.getContractFactory("ERC1967ProxyMock");
    const FeeManagerMock = await ethers.getContractFactory("FeeManagerMock");
    const Bridge = await ethers.getContractFactory("Bridge");
    const ERC20MintableBurnable = await ethers.getContractFactory("ERC20MintableBurnable");

    feeManager = await FeeManagerMock.deploy();
    bridge = await Bridge.deploy();

    feeToken = await ERC20MintableBurnable.deploy("ERC20FeeMock", "ERC20FeeMock", 18, OWNER.address);
    await feeToken.mintTo(OWNER.address, wei("2"));

    feeTokens = [ETHEREUM_ADDRESS, feeToken.address];
    feeAmounts = [wei("1"), wei("2")];

    const erc1967 = await ERC1967ProxyMock.deploy(feeManager.address, "0x");

    feeManager = FeeManagerMock.attach(erc1967.address);
    await feeManager.__FeeManagerMock_init(bridge.address);

    await bridge.__Bridge_init(SIGNER.address, ZERO_ADDR, CHAIN_NAME, feeManager.address);

    signHelper = new SignHelper(SIGNER, CHAIN_NAME, bridge.address, feeManager.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("proxy", () => {
    let feeManagerV2: FeeManagerMock;

    beforeEach(async () => {
      const FeeManagerMock = await ethers.getContractFactory("FeeManagerMock");

      feeManagerV2 = await FeeManagerMock.deploy();
    });

    it("should not upgrade if wrong method", async () => {
      const tx = feeManager.upgradeTo(feeManagerV2.address);

      await expect(tx).to.be.revertedWith("FeeManager: this upgrade method is off");
    });

    it("should not upgrade if zero address", async () => {
      const signature = signHelper.signFacadeAuthorizeUpgrade(ZERO_ADDR, 0);

      const tx = feeManager.upgradeToWithSig(ZERO_ADDR, signature);

      await expect(tx).to.be.revertedWith("FeeManager: zero address");
    });

    it("should not upgrade if invalid signature", async () => {
      const signature = signHelper.signFacadeAuthorizeUpgrade(feeManagerV2.address, 1);

      const tx = feeManager.upgradeToWithSig(feeManagerV2.address, signature);

      await expect(tx).to.be.revertedWith("Signers: invalid signature");
    });

    it("should not call upgrade method on implementation", async () => {
      const signature = signHelper.signFacadeAuthorizeUpgrade(feeManager.address, 0);

      const tx = feeManagerV2.upgradeToWithSig(feeManager.address, signature);

      await expect(tx).to.be.revertedWith("Function must be called through delegatecall");
    });

    it("should not initialize parent contract", async () => {
      const tx = feeManager.__FeeManager_init(bridge.address);

      await expect(tx).to.be.revertedWith("Initializable: contract is not initializing");
    });

    it("should upgrade if all conditions are met", async () => {
      const signature = signHelper.signFacadeAuthorizeUpgrade(feeManagerV2.address, 0);

      await feeManager.upgradeToWithSig(feeManagerV2.address, signature);

      const ERC1967Proxy = await ethers.getContractFactory("ERC1967ProxyMock");

      expect(await (await ERC1967Proxy.attach(feeManager.address)).implementation()).to.eq(feeManagerV2.address);
    });
  });

  describe("#addFeeToken", () => {
    it("should not add fee token if params lengths mismatch", async () => {
      const signature = signHelper.signAddFeeToken([], feeAmounts, 0);

      const tx = feeManager.addFeeToken({ feeTokens: [], feeAmounts, signature });

      await expect(tx).to.be.revertedWith("FeeManager: params lengths mismatch");
    });

    it("should not add fee token if token already exists", async () => {
      const signature = signHelper.signAddFeeToken([ETHEREUM_ADDRESS, ETHEREUM_ADDRESS], feeAmounts, 0);

      const tx = feeManager.addFeeToken({ feeTokens: [ETHEREUM_ADDRESS, ETHEREUM_ADDRESS], feeAmounts, signature });

      await expect(tx).to.be.revertedWith("FeeManager: token already exists");
    });

    it("should not add fee token if invalid signature", async () => {
      const signature = signHelper.signAddFeeToken(feeTokens, feeAmounts, 1);

      const tx = feeManager.addFeeToken({ feeTokens, feeAmounts, signature });

      await expect(tx).to.be.revertedWith("Signers: invalid signature");
    });

    it("should add fee token if all conditions are met", async () => {
      const signature = signHelper.signAddFeeToken(feeTokens, feeAmounts, 0);

      const tx = feeManager.addFeeToken({ feeTokens, feeAmounts, signature });

      await expect(tx)
        .to.emit(feeManager, "AddedFeeToken")
        .withArgs(feeTokens[0], feeAmounts[0])
        .to.emit(feeManager, "AddedFeeToken")
        .withArgs(feeTokens[1], feeAmounts[1]);

      expect(await feeManager.getCommission(feeTokens[0])).to.eq(feeAmounts[0]);
      expect(await feeManager.getCommission(feeTokens[1])).to.eq(feeAmounts[1]);
    });
  });

  context("if fee tokens are added", () => {
    beforeEach(async () => {
      const signature = signHelper.signAddFeeToken(feeTokens, feeAmounts, 0);

      await feeManager.addFeeToken({ feeTokens, feeAmounts, signature });
    });

    describe("#removeFeeToken", () => {
      it("should not remove fee token if params lengths mismatch", async () => {
        const signature = signHelper.signRemoveFeeToken(feeTokens, feeAmounts, 0);

        const tx = feeManager.removeFeeToken({ feeTokens: [], feeAmounts, signature });

        await expect(tx).to.be.revertedWith("FeeManager: params lengths mismatch");
      });

      it("should not remove fee token if wrong token address or amount", async () => {
        const signature = signHelper.signRemoveFeeToken([ETHEREUM_ADDRESS, ETHEREUM_ADDRESS], feeAmounts, 0);

        const tx = feeManager.removeFeeToken({
          feeTokens: [ETHEREUM_ADDRESS, ETHEREUM_ADDRESS],
          feeAmounts,
          signature,
        });

        await expect(tx).to.be.revertedWith("FeeManager: wrong token address or amount");
      });

      it("should not remove fee token if invalid signature", async () => {
        const signature = signHelper.signRemoveFeeToken(feeTokens, feeAmounts, 1);

        const tx = feeManager.removeFeeToken({ feeTokens, feeAmounts, signature });

        await expect(tx).to.be.revertedWith("Signers: invalid signature");
      });

      it("should remove fee token if all conditions are met", async () => {
        const signature = signHelper.signRemoveFeeToken(feeTokens, feeAmounts, 0);

        const tx = feeManager.removeFeeToken({ feeTokens, feeAmounts, signature });

        await expect(tx)
          .to.emit(feeManager, "RemovedFeeToken")
          .withArgs(feeTokens[0], feeAmounts[0])
          .to.emit(feeManager, "RemovedFeeToken")
          .withArgs(feeTokens[1], feeAmounts[1]);

        expect(await feeManager.getCommission(feeTokens[0])).to.eq(0);
        expect(await feeManager.getCommission(feeTokens[1])).to.eq(0);
      });
    });

    describe("#updateFeeToken", () => {
      const newAmounts = [wei("0.5"), wei("1")];

      it("should not update fee token if params lengths mismatch", async () => {
        const signature = signHelper.signUpdateFeeToken([], newAmounts, 0);

        const tx = feeManager.updateFeeToken({ feeTokens: [], feeAmounts: newAmounts, signature });

        await expect(tx).to.be.revertedWith("FeeManager: params lengths mismatch");
      });

      it("should not update fee token if token doesn't exist", async () => {
        const signature = signHelper.signUpdateFeeToken([COMMISSION_ADDRESS], [0], 0);

        const tx = feeManager.updateFeeToken({ feeTokens: [COMMISSION_ADDRESS], feeAmounts: [0], signature });

        await expect(tx).to.be.revertedWith("FeeManager: token doesn't exist");
      });

      it("should not update fee token if invalid signature", async () => {
        const signature = signHelper.signUpdateFeeToken(feeTokens, newAmounts, 1);

        const tx = feeManager.updateFeeToken({ feeTokens, feeAmounts: newAmounts, signature });

        await expect(tx).to.be.revertedWith("Signers: invalid signature");
      });

      it("should update fee token if all conditions are met", async () => {
        const signature = signHelper.signUpdateFeeToken(feeTokens, newAmounts, 0);

        const tx = feeManager.updateFeeToken({ feeTokens, feeAmounts: newAmounts, signature });

        await expect(tx)
          .to.emit(feeManager, "UpdatedFeeToken")
          .withArgs(feeTokens[0], newAmounts[0])
          .to.emit(feeManager, "UpdatedFeeToken")
          .withArgs(feeTokens[1], newAmounts[1]);

        expect(await feeManager.getCommission(feeTokens[0])).to.eq(newAmounts[0]);
        expect(await feeManager.getCommission(feeTokens[1])).to.eq(newAmounts[1]);
      });
    });

    describe("#withdrawFeeToken", () => {
      beforeEach(async () => {
        await setBalance(feeManager.address, feeAmounts[0]);
        await feeToken.transfer(feeManager.address, feeAmounts[1]);
      });

      it("should not withdraw fee token if params lengths mismatch", async () => {
        const signature = signHelper.signWithdrawFeeToken(OWNER.address, [], feeAmounts, 0);

        const tx = feeManager.withdrawFeeToken({
          receiver: OWNER.address,
          feeTokens: [],
          amounts: feeAmounts,
          signature,
        });

        await expect(tx).to.be.revertedWith("FeeManager: params lengths mismatch");
      });

      it("should not withdraw fee token if commission token", async () => {
        const signature = signHelper.signWithdrawFeeToken(OWNER.address, [COMMISSION_ADDRESS], [0], 0);

        const tx = feeManager.withdrawFeeToken({
          receiver: OWNER.address,
          feeTokens: [COMMISSION_ADDRESS],
          amounts: [0],
          signature,
        });

        await expect(tx).to.be.revertedWith("FeeManager: commission token");
      });

      it("should not withdraw fee token if failed to withdraw native", async () => {
        const signature = signHelper.signWithdrawFeeToken(feeToken.address, feeTokens, feeAmounts, 0);

        const tx = feeManager.withdrawFeeToken({
          receiver: feeToken.address,
          feeTokens,
          amounts: feeAmounts,
          signature,
        });

        await expect(tx).to.be.revertedWith("FeeManager: failed to withdraw native");
      });

      it("should not withdraw fee token if invalid signature", async () => {
        const signature = signHelper.signWithdrawFeeToken(OWNER.address, feeTokens, feeAmounts, 1);

        const tx = feeManager.withdrawFeeToken({ receiver: OWNER.address, feeTokens, amounts: feeAmounts, signature });

        await expect(tx).to.be.revertedWith("Signers: invalid signature");
      });

      it("should withdraw fee token if all conditions are met", async () => {
        const signature = signHelper.signWithdrawFeeToken(OWNER.address, feeTokens, feeAmounts, 0);

        const tx = feeManager.withdrawFeeToken({ receiver: OWNER.address, feeTokens, amounts: feeAmounts, signature });

        await expect(tx)
          .to.emit(feeManager, "WithdrawnFeeToken")
          .withArgs(OWNER.address, feeTokens[0], feeAmounts[0])
          .to.emit(feeManager, "WithdrawnFeeToken")
          .withArgs(OWNER.address, feeTokens[1], feeAmounts[1]);

        await expect(tx).to.changeEtherBalances([feeManager, OWNER], [-feeAmounts[0], feeAmounts[0]]);
        await expect(tx).to.changeTokenBalances(feeToken, [feeManager, OWNER], [-feeAmounts[1], feeAmounts[1]]);
      });
    });
  });
});
