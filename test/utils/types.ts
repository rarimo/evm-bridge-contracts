import { IERC1155Handler, IERC20Handler, IERC721Handler, INativeHandler, ISBTHandler } from "@ethers-v5";
import { BytesLike } from "ethers";

export type WithdrawERC20Parameters = IERC20Handler.WithdrawERC20ParametersStruct;
export type WithdrawERC721Parameters = IERC721Handler.WithdrawERC721ParametersStruct;
export type WithdrawSBTParameters = ISBTHandler.WithdrawSBTParametersStruct;
export type WithdrawERC1155Parameters = IERC1155Handler.WithdrawERC1155ParametersStruct;
export type WithdrawNativeParameters = INativeHandler.WithdrawNativeParametersStruct;
export type Bundle = { bundle: BytesLike; salt: BytesLike };
