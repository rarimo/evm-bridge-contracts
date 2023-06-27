import { ethers } from "hardhat";

import { BigNumberish, BytesLike, Wallet } from "ethers";
import { BridgeMethodId, FacadeMethodId } from "@/test/utils/constants";

export class SignHelper {
  constructor(public signer: Wallet, public chainName: string, public bridge: string, public facade?: string) {}

  public signAddFeeToken(feeTokens: Array<string>, feeAmounts: Array<BigNumberish>, nonce: BigNumberish): string {
    const hash = ethers.utils.solidityKeccak256(
      ["uint256", "address", "string", "uint8", "address[]", "uint256[]"],
      [nonce, this.facade, this.chainName, FacadeMethodId.AddFeeToken, feeTokens, feeAmounts]
    );

    return this.sign(hash);
  }

  public signRemoveFeeToken(feeTokens: Array<string>, feeAmounts: Array<BigNumberish>, nonce: BigNumberish): string {
    const hash = ethers.utils.solidityKeccak256(
      ["uint256", "address", "string", "uint8", "address[]", "uint256[]"],
      [nonce, this.facade, this.chainName, FacadeMethodId.RemoveFeeToken, feeTokens, feeAmounts]
    );

    return this.sign(hash);
  }

  public signUpdateFeeToken(feeTokens: Array<string>, feeAmounts: Array<BigNumberish>, nonce: BigNumberish): string {
    const hash = ethers.utils.solidityKeccak256(
      ["uint256", "address", "string", "uint8", "address[]", "uint256[]"],
      [nonce, this.facade, this.chainName, FacadeMethodId.UpdateFeeToken, feeTokens, feeAmounts]
    );

    return this.sign(hash);
  }

  public signWithdrawFeeToken(
    receiver: string,
    feeTokens: Array<string>,
    amounts: Array<BigNumberish>,
    nonce: BigNumberish
  ): string {
    const hash = ethers.utils.solidityKeccak256(
      ["uint256", "address", "address", "string", "uint8", "address[]", "uint256[]"],
      [nonce, receiver, this.facade, this.chainName, FacadeMethodId.WithdrawFeeToken, feeTokens, amounts]
    );

    return this.sign(hash);
  }

  public signBridgeAuthorizeUpgrade(newAddress: string, nonce: BigNumberish): string {
    const hash = ethers.utils.solidityKeccak256(
      ["uint8", "address", "string", "uint256", "address"],
      [BridgeMethodId.AuthorizeUpgrade, newAddress, this.chainName, nonce, this.bridge]
    );

    return this.sign(hash);
  }

  public signFacadeAuthorizeUpgrade(newAddress: string, nonce: BigNumberish): string {
    const hash = ethers.utils.solidityKeccak256(
      ["uint8", "address", "string", "uint256", "address"],
      [FacadeMethodId.AuthorizeUpgrade, newAddress, this.chainName, nonce, this.facade]
    );

    return this.sign(hash);
  }

  public signChangeBundleExecutorImplementation(newAddress: string, nonce: BigNumberish): string {
    const hash = ethers.utils.solidityKeccak256(
      ["uint8", "address", "string", "uint256", "address"],
      [BridgeMethodId.ChangeBundleExecutorImplementation, newAddress, this.chainName, nonce, this.bridge]
    );

    return this.sign(hash);
  }

  public signChangeFacade(newAddress: string, nonce: BigNumberish): string {
    const hash = ethers.utils.solidityKeccak256(
      ["uint8", "address", "string", "uint256", "address"],
      [BridgeMethodId.ChangeFacade, newAddress, this.chainName, nonce, this.bridge]
    );

    return this.sign(hash);
  }

  public signChangeSinger(newPubKey: string): string {
    const hash = ethers.utils.solidityKeccak256(["bytes"], [newPubKey]);

    return this.sign(hash);
  }

  public sign(hash: BytesLike) {
    return ethers.utils.joinSignature(this.signer._signingKey().signDigest(hash));
  }
}
