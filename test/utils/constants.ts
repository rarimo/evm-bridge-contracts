import { ZERO_BYTES32 } from "@/scripts/utils/constants";

export enum BridgeMethodId {
  None,
  AuthorizeUpgrade,
  ChangeBundleExecutorImplementation,
  ChangeFacade,
}

export enum FacadeMethodId {
  AddFeeToken,
  RemoveFeeToken,
  UpdateFeeToken,
  WithdrawFeeToken,
  AuthorizeUpgrade,
}

export const ETHEREUM_ADDRESS = "0x0000000000000000000000000000000000000000";
export const COMMISSION_ADDRESS = "0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF";

export const CHAIN_NAME = "ethereum";
export const RECEIVER = "receiver";

export const RANDOM_ORIGIN_HASH = "0x0202020202020202020202020202020202020202020202020202020202020202";
export const RANDOM_SALT = "0x0101010101010101010101010101010101010101010101010101010101010101";

export const EMPTY_BUNDLE = {
  bundle: "0x",
  salt: ZERO_BYTES32,
};
export const RANDOM_BUNDLE = {
  bundle: "0x01",
  salt: RANDOM_SALT,
};
