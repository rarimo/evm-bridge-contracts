import { expect } from "chai";
import { Bridge } from "@ethers-v5";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "hardhat";
import { CHAIN_NAME, EMPTY_BUNDLE, RANDOM_ORIGIN_HASH } from "@/test/utils/constants";
import { Reverter } from "@/test/helpers/reverter";
import { SignHelper } from "@/test/utils/signature";
import { Wallet } from "ethers";
import { ZERO_ADDR, ZERO_BYTES32 } from "@/scripts/utils/constants";
import { wei } from "@/scripts/utils/utils";
import { MerkleTreeHelper } from "@/test/utils/merkletree";
import { WithdrawERC20Parameters, WithdrawNativeParameters } from "@/test/utils/types";

describe("Bridge", () => {
  const reverter = new Reverter();

  let signHelper: SignHelper;
  let merkleHelper: MerkleTreeHelper;

  let OWNER: SignerWithAddress;
  let SECOND: SignerWithAddress;
  let SIGNER: Wallet;

  let bridge: Bridge;

  before("setup", async () => {
    [OWNER, SECOND] = await ethers.getSigners();
    SIGNER = ethers.Wallet.createRandom();

    const ERC1967ProxyMock = await ethers.getContractFactory("ERC1967ProxyMock");
    const Bridge = await ethers.getContractFactory("Bridge");

    bridge = await Bridge.deploy();

    const erc1967 = await ERC1967ProxyMock.deploy(bridge.address, "0x");

    bridge = Bridge.attach(erc1967.address);
    await bridge.__Bridge_init(SIGNER.address, SECOND.address, CHAIN_NAME, OWNER.address);

    signHelper = new SignHelper(SIGNER, CHAIN_NAME, bridge.address);
    merkleHelper = new MerkleTreeHelper(signHelper, CHAIN_NAME, bridge.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("signers", () => {
    it("only facade should call these methods", async () => {
      await expect(
        bridge.connect(SECOND).checkSignatureAndIncrementNonce(0, ZERO_ADDR, ZERO_BYTES32, "0x")
      ).to.be.revertedWith("Bundler: not a facade");

      await expect(
        bridge.connect(SECOND).validateChangeAddressSignature(0, ZERO_ADDR, ZERO_ADDR, "0x")
      ).to.be.revertedWith("Bundler: not a facade");
    });
  });

  describe("proxy", () => {
    let bridgeV2: Bridge;

    beforeEach(async () => {
      const Bridge = await ethers.getContractFactory("Bridge");

      bridgeV2 = await Bridge.deploy();
    });

    it("should not upgrade if wrong method", async () => {
      const tx = bridge.upgradeTo(bridgeV2.address);

      await expect(tx).to.be.revertedWith("Bridge: this upgrade method is off");
    });

    it("should not upgrade if zero address", async () => {
      const signature = signHelper.signBridgeAuthorizeUpgrade(ZERO_ADDR, 0);

      const tx = bridge.upgradeToWithSig(ZERO_ADDR, signature);

      await expect(tx).to.be.revertedWith("Bridge: zero address");
    });

    it("should not upgrade if invalid signature", async () => {
      const signature = signHelper.signBridgeAuthorizeUpgrade(bridgeV2.address, 1);

      const tx = bridge.upgradeToWithSig(bridgeV2.address, signature);

      await expect(tx).to.be.revertedWith("Signers: invalid signature");
    });

    it("should not call upgrade method on implementation", async () => {
      const signature = signHelper.signBridgeAuthorizeUpgrade(bridge.address, 0);

      const tx = bridgeV2.upgradeToWithSig(bridge.address, signature);

      await expect(tx).to.be.revertedWith("Function must be called through delegatecall");
    });

    it("should not initialize Bundler parent contract", async () => {
      const tx = bridge.__Bundler_init(ZERO_ADDR);

      await expect(tx).to.be.revertedWith("Initializable: contract is not initializing");
    });

    it("should not initialize Signer parent contract", async () => {
      const tx = bridge.__Signers_init(ZERO_ADDR, ZERO_ADDR, CHAIN_NAME);

      await expect(tx).to.be.revertedWith("Initializable: contract is not initializing");
    });

    it("should not initialize twice", async () => {
      const tx = bridge.__Bridge_init(SIGNER.address, ZERO_ADDR, CHAIN_NAME, ZERO_ADDR);

      await expect(tx).to.be.revertedWith("Initializable: contract is already initialized");
    });

    it("should upgrade if all conditions are met", async () => {
      const signature = signHelper.signBridgeAuthorizeUpgrade(bridgeV2.address, 0);

      await bridge.upgradeToWithSig(bridgeV2.address, signature);

      const ERC1967Proxy = await ethers.getContractFactory("ERC1967ProxyMock");

      expect(await (await ERC1967Proxy.attach(bridge.address)).implementation()).to.eq(bridgeV2.address);
    });
  });

  describe("#verifyMerkleLeaf", () => {
    let withdrawERC20Params: Partial<WithdrawERC20Parameters>;
    let withdrawNativeParams: Partial<WithdrawNativeParameters>;

    beforeEach(async () => {
      withdrawERC20Params = {
        token: SECOND.address,
        amount: wei("10"),
        bundle: EMPTY_BUNDLE,
        originHash: RANDOM_ORIGIN_HASH,
        receiver: OWNER.address,
        isWrapped: true,
      };
      withdrawNativeParams = {
        amount: wei("10"),
        bundle: EMPTY_BUNDLE,
        originHash: RANDOM_ORIGIN_HASH,
        receiver: OWNER.address,
      };
    });

    it("should not verify merkle leaf if the caller is not a facade", async () => {
      const nativeDataLeaf = ethers.utils.solidityPack(["uint256"], [wei("10")]);
      const nativeProof = merkleHelper.getProof(merkleHelper.encodeLeaf(withdrawNativeParams));

      const tx = bridge
        .connect(SECOND)
        .verifyMerkleLeaf(
          nativeDataLeaf,
          withdrawNativeParams.bundle!,
          withdrawNativeParams.originHash!,
          OWNER.address,
          nativeProof
        );

      await expect(tx).to.be.revertedWith("Bundler: not a facade");
    });

    it("should not verify merkle leaf if the hash nonce is used", async () => {
      const erc20DataLeaf = ethers.utils.solidityPack(["address", "uint256"], [SECOND.address, wei("10")]);
      const nativeDataLeaf = ethers.utils.solidityPack(["uint256"], [wei("10")]);

      const erc20Proof = merkleHelper.getProof(merkleHelper.encodeLeaf(withdrawERC20Params));
      const nativeProof = merkleHelper.getProof(merkleHelper.encodeLeaf(withdrawNativeParams));

      await bridge.verifyMerkleLeaf(
        erc20DataLeaf,
        withdrawERC20Params.bundle!,
        withdrawERC20Params.originHash!,
        OWNER.address,
        erc20Proof
      );
      const tx = bridge.verifyMerkleLeaf(
        nativeDataLeaf,
        withdrawNativeParams.bundle!,
        withdrawNativeParams.originHash!,
        OWNER.address,
        nativeProof
      );

      await expect(tx).to.be.revertedWith("Hashes: the hash nonce is used");
    });

    it("should verify merkle leaf if all conditions are met", async () => {
      const erc20DataLeaf = ethers.utils.solidityPack(["address", "uint256"], [SECOND.address, wei("10")]);
      const nativeDataLeaf = ethers.utils.solidityPack(["uint256"], [wei("10")]);

      const erc20Proof = merkleHelper.getProof(merkleHelper.encodeLeaf(withdrawERC20Params));

      withdrawNativeParams.originHash = ethers.utils.randomBytes(32);
      const nativeProof = merkleHelper.getProof(merkleHelper.encodeLeaf(withdrawNativeParams));

      expect(await bridge.usedHashes(withdrawERC20Params.originHash!)).to.be.false;
      expect(await bridge.usedHashes(withdrawNativeParams.originHash!)).to.be.false;

      await bridge.verifyMerkleLeaf(
        erc20DataLeaf,
        withdrawERC20Params.bundle!,
        withdrawERC20Params.originHash!,
        OWNER.address,
        erc20Proof
      );

      expect(await bridge.usedHashes(withdrawERC20Params.originHash!)).to.be.true;
      expect(await bridge.usedHashes(withdrawNativeParams.originHash!)).to.be.false;

      await bridge.verifyMerkleLeaf(
        nativeDataLeaf,
        withdrawNativeParams.bundle!,
        withdrawNativeParams.originHash!,
        OWNER.address,
        nativeProof
      );

      expect(await bridge.usedHashes(withdrawERC20Params.originHash!)).to.be.true;
      expect(await bridge.usedHashes(withdrawNativeParams.originHash!)).to.be.true;
    });
  });

  describe("#changeSigner", () => {
    const newSigner = ethers.Wallet.createRandom();
    const publicKey = "0x" + newSigner.publicKey.slice(4);

    it("should not change signer if invalid signature", async () => {
      const signature = signHelper.signChangeSinger(newSigner.privateKey);

      const tx = bridge.changeSigner(publicKey, signature);

      await expect(tx).to.be.revertedWith("Signers: invalid signature");
    });

    it("should not change signer if wrong pubKey length", async () => {
      const signature = signHelper.signChangeSinger(newSigner.privateKey);

      const tx = bridge.changeSigner(newSigner.privateKey, signature);

      await expect(tx).to.be.revertedWith("Signers: wrong pubKey length");
    });

    it("should not change signer if zero pubKey", async () => {
      const pBytes = "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F";
      const zeroBytes = "0000000000000000000000000000000000000000000000000000000000000000";
      const notNull = "0000000000000000000000000000000000000000000000000000000000000001";

      const zeroPubKeys = [
        "0x" + zeroBytes + notNull,
        "0x" + pBytes + notNull,
        "0x" + notNull + zeroBytes,
        "0x" + notNull + pBytes,
      ];

      for (const pubKey of zeroPubKeys) {
        const signature = signHelper.signChangeSinger(pubKey);

        const tx = bridge.changeSigner(pubKey, signature);

        await expect(tx).to.be.revertedWith("Signers: zero pubKey");
      }
    });

    it("should not change signer if pubKey not on the curve", async () => {
      const wrongPubKey =
        "0x10101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010";

      const signature = signHelper.signChangeSinger(wrongPubKey);

      const tx = bridge.changeSigner(wrongPubKey, signature);

      await expect(tx).to.be.revertedWith("Signers: pubKey not on the curve");
    });

    it("should change signer if all conditions are met", async () => {
      expect(await bridge.callStatic.signer()).to.eq(SIGNER.address);

      const signature = signHelper.signChangeSinger(publicKey);

      await bridge.changeSigner(publicKey, signature);

      expect(await bridge.callStatic.signer()).to.eq(newSigner.address);
    });
  });

  describe("#changeBundleExecutorImplementation", () => {
    it("should not change bundle executor implementation if zero address", async () => {
      const signature = signHelper.signChangeBundleExecutorImplementation(ZERO_ADDR, 0);

      const tx = bridge.changeBundleExecutorImplementation(ZERO_ADDR, signature);

      await expect(tx).to.be.revertedWith("Bridge: zero address");
    });

    it("should not change bundle executor implementation if invalid signature", async () => {
      const signature = signHelper.signChangeBundleExecutorImplementation(OWNER.address, 1);

      const tx = bridge.changeBundleExecutorImplementation(OWNER.address, signature);

      await expect(tx).to.be.revertedWith("Signers: invalid signature");
    });

    it("should change bundle executor implementation if all conditions are met", async () => {
      expect(await bridge.bundleExecutorImplementation()).to.eq(SECOND.address);

      const signature = signHelper.signChangeBundleExecutorImplementation(OWNER.address, 0);

      await bridge.changeBundleExecutorImplementation(OWNER.address, signature);

      expect(await bridge.bundleExecutorImplementation()).to.eq(OWNER.address);
    });
  });

  describe("#changeFacade", () => {
    it("should not change facade if zero address", async () => {
      const signature = signHelper.signChangeFacade(ZERO_ADDR, 0);

      const tx = bridge.changeFacade(ZERO_ADDR, signature);

      await expect(tx).to.be.revertedWith("Bridge: zero address");
    });

    it("should not change facade if invalid signature", async () => {
      const signature = signHelper.signChangeFacade(SECOND.address, 1);

      const tx = bridge.changeBundleExecutorImplementation(SECOND.address, signature);

      await expect(tx).to.be.revertedWith("Signers: invalid signature");
    });

    it("should change facade if all conditions are met", async () => {
      expect(await bridge.facade()).to.eq(OWNER.address);

      const signature = signHelper.signChangeFacade(SECOND.address, 0);

      await bridge.changeFacade(SECOND.address, signature);

      expect(await bridge.facade()).to.eq(SECOND.address);
    });
  });
});
