import { expect } from "chai";
import {
  Bridge,
  BridgeFacade,
  ERC20MintableBurnable,
  ERC721MintableBurnable,
  SBTMintableBurnable,
  ERC1155MintableBurnable,
  WrappedNativeMock,
} from "@ethers-v5";
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { ethers } from "hardhat";
import { wei } from "@/scripts/utils/utils";
import {
  CHAIN_NAME,
  COMMISSION_ADDRESS,
  EMPTY_BUNDLE,
  ETHEREUM_ADDRESS,
  RANDOM_BUNDLE,
  RANDOM_ORIGIN_HASH,
  RANDOM_SALT,
  RECEIVER,
} from "@/test/utils/constants";
import { Reverter } from "@/test/helpers/reverter";
import { SignHelper } from "@/test/utils/signature";
import { Wallet } from "ethers";
import { MerkleTreeHelper } from "@/test/utils/merkletree";
import { ZERO_ADDR } from "@/scripts/utils/constants";
import { setBalance } from "@nomicfoundation/hardhat-network-helpers";
import { Bundle } from "@/test/utils/types";

describe("BridgeFacade", () => {
  const reverter = new Reverter();

  let signHelper: SignHelper;
  let merkleHelper: MerkleTreeHelper;

  let OWNER: SignerWithAddress;
  let SECOND: SignerWithAddress;
  let SIGNER: Wallet;

  let facade: BridgeFacade;
  let bridge: Bridge;
  let feeToken: ERC20MintableBurnable;
  let feeTokens: Array<string>;
  let feeAmounts: Array<bigint>;
  let erc20: ERC20MintableBurnable;
  let erc721: ERC721MintableBurnable;
  let sbt: SBTMintableBurnable;
  let erc1155: ERC1155MintableBurnable;
  let weth9: WrappedNativeMock;

  before("setup", async () => {
    [OWNER, SECOND] = await ethers.getSigners();
    SIGNER = ethers.Wallet.createRandom();

    const BridgeFacade = await ethers.getContractFactory("BridgeFacade");
    const Bridge = await ethers.getContractFactory("Bridge");
    const BundleExecutorImplementation = await ethers.getContractFactory("BundleExecutorImplementation");
    const ERC20MintableBurnable = await ethers.getContractFactory("ERC20MintableBurnable");
    const ERC721MintableBurnable = await ethers.getContractFactory("ERC721MintableBurnable");
    const ERC1155MintableBurnable = await ethers.getContractFactory("ERC1155MintableBurnable");
    const SBTMintableBurnable = await ethers.getContractFactory("SBTMintableBurnable");
    const WrappedNativeMock = await ethers.getContractFactory("WrappedNativeMock");

    facade = await BridgeFacade.deploy();
    bridge = await Bridge.deploy();
    weth9 = await WrappedNativeMock.deploy();
    const bundleImplementation = await BundleExecutorImplementation.deploy(weth9.address);

    feeToken = await ERC20MintableBurnable.deploy("ERC20FeeMock", "ERC20FeeMock", 18, OWNER.address);

    feeTokens = [ETHEREUM_ADDRESS, feeToken.address, COMMISSION_ADDRESS];
    feeAmounts = [wei("1"), wei("2"), 1n];

    await feeToken.mintTo(OWNER.address, feeAmounts[1]);
    await feeToken.approve(facade.address, feeAmounts[1]);

    erc20 = await ERC20MintableBurnable.deploy("ERC20Mock", "ERC20Mock", 18, OWNER.address);
    await erc20.mintTo(OWNER.address, wei("10"));
    await erc20.approve(facade.address, wei("10"));
    await erc20.approve(bridge.address, wei("10"));

    erc721 = await ERC721MintableBurnable.deploy("ERC721Mock", "ERC721Mock", OWNER.address, "");
    await erc721.mintTo(OWNER.address, 10, "URI");
    await erc721.approve(facade.address, 10);

    sbt = await SBTMintableBurnable.deploy("SBTMock", "SBTMock", OWNER.address, "");
    await sbt.attestTo(OWNER.address, 10, "URI");
    await sbt.attestTo(SECOND.address, 11, "URI");

    erc1155 = await ERC1155MintableBurnable.deploy(OWNER.address, "");
    await erc1155.mintTo(OWNER.address, 10, wei("10"), "URI");
    await erc1155.setApprovalForAll(facade.address, true);

    await bridge.__Bridge_init(SIGNER.address, bundleImplementation.address, CHAIN_NAME, facade.address);
    await facade.__BridgeFacade_init(bridge.address);

    await erc20.transferOwnership(bridge.address);
    await erc721.transferOwnership(bridge.address);
    await sbt.transferOwnership(bridge.address);
    await erc1155.transferOwnership(bridge.address);

    signHelper = new SignHelper(SIGNER, CHAIN_NAME, bridge.address, facade.address);
    merkleHelper = new MerkleTreeHelper(signHelper, CHAIN_NAME, bridge.address);

    await reverter.snapshot();
  });

  afterEach(reverter.revert);

  describe("proxy", () => {
    it("should not initialize twice", async () => {
      await expect(facade.__BridgeFacade_init(bridge.address)).to.be.revertedWith(
        "Initializable: contract is already initialized"
      );
    });
  });

  context("if fee tokens are added", () => {
    beforeEach(async () => {
      const signature = signHelper.signAddFeeToken(feeTokens, feeAmounts, 0);

      await facade.addFeeToken({ feeTokens, feeAmounts, signature });
    });

    describe("erc20", () => {
      describe("#depositERC20", () => {
        it("should not deposit if excessive native amount", async () => {
          const tx = facade.depositERC20(
            {
              feeToken: feeTokens[0],
            },
            {
              token: erc20.address,
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            },
            { value: feeAmounts[0] + 1n }
          );

          await expect(tx).to.be.revertedWith("BridgeFacade: excessive native amount");
        });

        it("should not deposit if zero token", async () => {
          const tx = facade.depositERC20(
            {
              feeToken: COMMISSION_ADDRESS,
            },
            {
              token: ZERO_ADDR,
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            }
          );

          await expect(tx).to.be.revertedWith("BridgeFacade: zero token");
        });

        it("should not deposit if amount is zero", async () => {
          const tx = facade.depositERC20(
            {
              feeToken: COMMISSION_ADDRESS,
            },
            {
              token: erc20.address,
              amount: 0,
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            }
          );

          await expect(tx).to.be.revertedWith("BridgeFacade: amount is zero");
        });

        it("should not deposit if wrong fee token", async () => {
          const tx = facade.depositERC20(
            {
              feeToken: OWNER.address,
            },
            {
              token: erc20.address,
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            }
          );

          await expect(tx).to.be.revertedWith("FeeManager: wrong fee token");
        });

        it("should not deposit if wrong native amount", async () => {
          const tx = facade.depositERC20(
            {
              feeToken: ETHEREUM_ADDRESS,
            },
            {
              token: erc20.address,
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            },
            { value: wei("0.5") }
          );

          await expect(tx).to.be.revertedWith("FeeManager: wrong native amount");
        });

        it("should deposit if isWrapped=true and no commission", async () => {
          const tx = facade.depositERC20(
            {
              feeToken: feeTokens[2],
            },
            {
              token: erc20.address,
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: true,
            }
          );

          await expect(tx).to.changeTokenBalances(erc20, [OWNER, facade, bridge], [-wei("10"), 0, 0]);
        });

        it("should deposit if isWrapped=false and native commission", async () => {
          const tx = facade.depositERC20(
            {
              feeToken: feeTokens[0],
            },
            {
              token: erc20.address,
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            },
            { value: feeAmounts[0] }
          );

          await expect(tx).to.changeEtherBalances([OWNER, facade, bridge], [-feeAmounts[0], feeAmounts[0], 0]);
          await expect(tx).to.changeTokenBalances(erc20, [OWNER, facade, bridge], [-wei("10"), 0, wei("10")]);
        });

        it("should deposit if isWrapped=false and erc20 commission", async () => {
          const tx = facade.depositERC20(
            {
              feeToken: feeTokens[1],
            },
            {
              token: erc20.address,
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            }
          );

          await expect(tx).to.changeTokenBalances(
            feeToken,
            [OWNER, facade, bridge],
            [-feeAmounts[1], feeAmounts[1], 0]
          );
          await expect(tx).to.changeTokenBalances(erc20, [OWNER, facade, bridge], [-wei("10"), 0, wei("10")]);
        });
      });

      describe("#withdrawERC20", () => {
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

          it("should not withdraw if invalid proof", async () => {
            const withdrawParams = {
              token: erc20.address,
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              originHash: RANDOM_ORIGIN_HASH,
              receiver: OWNER.address,
              isWrapped: false,
            };

            const leaf = merkleHelper.encodeLeaf(withdrawParams);

            merkleHelper.addLeaf(leaf);

            const path = merkleHelper.getPath(leaf);
            const signature = signHelper.sign(ethers.utils.randomBytes(32));

            const proof = ethers.utils.defaultAbiCoder.encode(["bytes32[]", "bytes"], [path, signature]);

            const tx = facade.withdrawERC20({ ...withdrawParams, proof });

            await expect(tx).to.be.revertedWith("Signers: invalid signature");
          });

          it("should withdraw if all conditions are met", async () => {
            const withdrawParams = {
              token: erc20.address,
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              originHash: RANDOM_ORIGIN_HASH,
              receiver: OWNER.address,
              isWrapped: false,
            };

            const leaf = merkleHelper.encodeLeaf(withdrawParams);
            const proof = merkleHelper.getProof(leaf);

            const tx = facade.withdrawERC20({ ...withdrawParams, proof });

            await expect(tx).to.changeTokenBalances(
              erc20,
              [OWNER, bridge, facade, proxyAddress],
              [wei("10"), -wei("10"), 0, 0]
            );
          });

          it("should withdraw if all conditions are met with wrong bundle", async () => {
            const withdrawParams = {
              token: erc20.address,
              amount: wei("10"),
              bundle: RANDOM_BUNDLE,
              originHash: RANDOM_ORIGIN_HASH,
              receiver: OWNER.address,
              isWrapped: false,
            };

            const leaf = merkleHelper.encodeLeaf(withdrawParams);
            const proof = merkleHelper.getProof(leaf);

            const tx = facade.withdrawERC20({ ...withdrawParams, proof });

            await expect(tx).to.changeTokenBalances(
              erc20,
              [OWNER, bridge, facade, proxyAddress],
              [wei("10"), -wei("10"), 0, 0]
            );
          });

          it("should withdraw bundle if all conditions are met", async () => {
            const withdrawParams = {
              token: erc20.address,
              amount: wei("10"),
              bundle: bundleApprove,
              originHash: RANDOM_ORIGIN_HASH,
              receiver: OWNER.address,
              isWrapped: false,
            };

            const leaf = merkleHelper.encodeLeaf(withdrawParams);
            const proof = merkleHelper.getProof(leaf);

            const tx = facade.withdrawERC20({ ...withdrawParams, proof });

            await expect(tx).to.changeTokenBalances(
              erc20,
              [OWNER, bridge, facade, proxyAddress],
              [0, -wei("10"), 0, wei("10")]
            );
            expect(await erc20.allowance(proxyAddress, OWNER.address)).to.eq(wei("5"));
          });
        });

        context("if burned", () => {
          beforeEach(async () => {
            await erc20.burnFrom(OWNER.address, wei("10"));
          });

          it("should withdraw if all conditions are met", async () => {
            const withdrawParams = {
              token: erc20.address,
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              originHash: RANDOM_ORIGIN_HASH,
              receiver: OWNER.address,
              isWrapped: true,
            };

            const leaf = merkleHelper.encodeLeaf(withdrawParams);
            const proof = merkleHelper.getProof(leaf);

            const tx = facade.withdrawERC20({ ...withdrawParams, proof });

            await expect(tx).to.changeTokenBalances(erc20, [OWNER, bridge, facade, proxyAddress], [wei("10"), 0, 0, 0]);
          });

          it("should withdraw bundle if all conditions are met", async () => {
            const withdrawParams = {
              token: erc20.address,
              amount: wei("10"),
              bundle: bundleApprove,
              originHash: RANDOM_ORIGIN_HASH,
              receiver: OWNER.address,
              isWrapped: true,
            };

            const leaf = merkleHelper.encodeLeaf(withdrawParams);
            const proof = merkleHelper.getProof(leaf);

            const tx = facade.withdrawERC20({ ...withdrawParams, proof });

            await expect(tx).to.changeTokenBalances(erc20, [OWNER, bridge, facade, proxyAddress], [0, 0, 0, wei("10")]);
            expect(await erc20.allowance(proxyAddress, OWNER.address)).to.eq(wei("5"));
          });
        });
      });
    });

    describe("erc721", () => {
      describe("#depositERC721", () => {
        it("should not deposit if excessive native amount", async () => {
          const tx = facade.depositERC721(
            {
              feeToken: feeTokens[0],
            },
            {
              token: erc721.address,
              tokenId: 10,
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            },
            { value: feeAmounts[0] + 1n }
          );

          await expect(tx).to.be.revertedWith("BridgeFacade: excessive native amount");
        });

        it("should not deposit if zero token", async () => {
          const tx = facade.depositERC721(
            {
              feeToken: COMMISSION_ADDRESS,
            },
            {
              token: ZERO_ADDR,
              tokenId: 10,
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            }
          );

          await expect(tx).to.be.revertedWith("BridgeFacade: zero token");
        });

        it("should not deposit if wrong fee token", async () => {
          const tx = facade.depositERC721(
            {
              feeToken: OWNER.address,
            },
            {
              token: erc721.address,
              tokenId: 10,
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            }
          );

          await expect(tx).to.be.revertedWith("FeeManager: wrong fee token");
        });

        it("should not deposit if wrong native amount", async () => {
          const tx = facade.depositERC721(
            {
              feeToken: ETHEREUM_ADDRESS,
            },
            {
              token: erc721.address,
              tokenId: 10,
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            },
            { value: wei("0.5") }
          );

          await expect(tx).to.be.revertedWith("FeeManager: wrong native amount");
        });

        it("should deposit if isWrapped=true and no commission", async () => {
          const tx = facade.depositERC721(
            {
              feeToken: feeTokens[2],
            },
            {
              token: erc721.address,
              tokenId: 10,
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: true,
            }
          );

          await expect(tx).to.changeTokenBalances(erc721, [OWNER, facade, bridge], [-1, 0, 0]);
        });

        it("should deposit if isWrapped=false and native commission", async () => {
          const tx = facade.depositERC721(
            {
              feeToken: feeTokens[0],
            },
            {
              token: erc721.address,
              tokenId: 10,
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            },
            { value: feeAmounts[0] }
          );

          await expect(tx).to.changeEtherBalances([OWNER, facade, bridge], [-feeAmounts[0], feeAmounts[0], 0]);
          await expect(tx).to.changeTokenBalances(erc721, [OWNER, facade, bridge], [-1, 0, 1]);
        });

        it("should deposit if isWrapped=false and erc20 commission", async () => {
          const tx = facade.depositERC721(
            {
              feeToken: feeTokens[1],
            },
            {
              token: erc721.address,
              tokenId: 10,
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            }
          );

          await expect(tx).to.changeTokenBalances(
            feeToken,
            [OWNER, facade, bridge],
            [-feeAmounts[1], feeAmounts[1], 0]
          );
          await expect(tx).to.changeTokenBalances(erc721, [OWNER, facade, bridge], [-1, 0, 1]);
        });
      });

      describe("#withdrawERC721", () => {
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

          it("should not withdraw if invalid proof", async () => {
            const withdrawParams = {
              token: erc721.address,
              tokenId: 10,
              tokenURI: "URI",
              bundle: EMPTY_BUNDLE,
              originHash: RANDOM_ORIGIN_HASH,
              receiver: OWNER.address,
              isWrapped: false,
            };

            const leaf = merkleHelper.encodeLeaf(withdrawParams);

            merkleHelper.addLeaf(leaf);

            const path = merkleHelper.getPath(leaf);
            const signature = signHelper.sign(ethers.utils.randomBytes(32));

            const proof = ethers.utils.defaultAbiCoder.encode(["bytes32[]", "bytes"], [path, signature]);

            const tx = facade.withdrawERC721({ ...withdrawParams, proof });

            await expect(tx).to.be.revertedWith("Signers: invalid signature");
          });

          it("should withdraw if all conditions are met", async () => {
            const withdrawParams = {
              token: erc721.address,
              tokenId: 10,
              tokenURI: "URI",
              bundle: EMPTY_BUNDLE,
              originHash: RANDOM_ORIGIN_HASH,
              receiver: OWNER.address,
              isWrapped: false,
            };

            const leaf = merkleHelper.encodeLeaf(withdrawParams);
            const proof = merkleHelper.getProof(leaf);

            const tx = facade.withdrawERC721({ ...withdrawParams, proof });

            await expect(tx).to.changeTokenBalances(erc721, [OWNER, bridge, facade, proxyAddress], [1, -1, 0, 0]);
          });

          it("should withdraw if all conditions are met with wrong bundle", async () => {
            const withdrawParams = {
              token: erc721.address,
              tokenId: 10,
              tokenURI: "URI",
              bundle: {
                bundle: ethers.utils.defaultAbiCoder.encode(
                  ["address[]", "uint256[]", "bytes[]"],
                  [[ZERO_ADDR], [], ["0x"]]
                ),
                salt: RANDOM_SALT,
              },
              originHash: RANDOM_ORIGIN_HASH,
              receiver: OWNER.address,
              isWrapped: false,
            };

            const leaf = merkleHelper.encodeLeaf(withdrawParams);
            const proof = merkleHelper.getProof(leaf);

            const tx = facade.withdrawERC721({ ...withdrawParams, proof });

            await expect(tx).to.changeTokenBalances(erc721, [OWNER, bridge, facade, proxyAddress], [1, -1, 0, 0]);
          });

          it("should withdraw bundle if all conditions are met", async () => {
            const withdrawParams = {
              token: erc721.address,
              tokenId: 10,
              tokenURI: "URI",
              bundle: bundleApprove,
              originHash: RANDOM_ORIGIN_HASH,
              receiver: OWNER.address,
              isWrapped: false,
            };

            const leaf = merkleHelper.encodeLeaf(withdrawParams);
            const proof = merkleHelper.getProof(leaf);

            const tx = facade.withdrawERC721({ ...withdrawParams, proof });

            await expect(tx).to.changeTokenBalances(erc721, [OWNER, bridge, facade, proxyAddress], [0, -1, 0, 1]);
            expect(await erc721.getApproved(10)).to.eq(OWNER.address);
          });
        });

        context("if burned", () => {
          beforeEach(async () => {
            await erc721.burnFrom(OWNER.address, 10);
          });

          it("should withdraw if all conditions are met", async () => {
            const withdrawParams = {
              token: erc721.address,
              tokenId: 10,
              tokenURI: "URI",
              bundle: EMPTY_BUNDLE,
              originHash: RANDOM_ORIGIN_HASH,
              receiver: OWNER.address,
              isWrapped: true,
            };

            const leaf = merkleHelper.encodeLeaf(withdrawParams);
            const proof = merkleHelper.getProof(leaf);

            const tx = facade.withdrawERC721({ ...withdrawParams, proof });

            await expect(tx).to.changeTokenBalances(erc721, [OWNER, bridge, facade, proxyAddress], [1, 0, 0, 0]);
            expect(await erc721.tokenURI(10)).to.eq("URI");
          });

          it("should withdraw bundle if all conditions are met", async () => {
            const withdrawParams = {
              token: erc721.address,
              tokenId: 10,
              tokenURI: "URI",
              bundle: bundleApprove,
              originHash: RANDOM_ORIGIN_HASH,
              receiver: OWNER.address,
              isWrapped: true,
            };

            const leaf = merkleHelper.encodeLeaf(withdrawParams);
            const proof = merkleHelper.getProof(leaf);

            const tx = facade.withdrawERC721({ ...withdrawParams, proof });

            await expect(tx).to.changeTokenBalances(erc721, [OWNER, bridge, facade, proxyAddress], [0, 0, 0, 1]);
            expect(await erc721.getApproved(10)).to.eq(OWNER.address);
            expect(await erc721.tokenURI(10)).to.eq("URI");
          });
        });
      });
    });

    describe("sbt", () => {
      describe("#depositSBT", () => {
        it("should not deposit if excessive native amount", async () => {
          const tx = facade.depositSBT(
            {
              feeToken: feeTokens[0],
            },
            {
              token: sbt.address,
              tokenId: 10,
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
            },
            { value: feeAmounts[0] + 1n }
          );

          await expect(tx).to.be.revertedWith("BridgeFacade: excessive native amount");
        });

        it("should not deposit if zero token", async () => {
          const tx = facade.depositSBT(
            {
              feeToken: COMMISSION_ADDRESS,
            },
            {
              token: ZERO_ADDR,
              tokenId: 10,
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
            }
          );

          await expect(tx).to.be.revertedWith("BridgeFacade: zero token");
        });

        it("should not deposit if invalid token id", async () => {
          const tx = facade.depositSBT(
            {
              feeToken: COMMISSION_ADDRESS,
            },
            {
              token: sbt.address,
              tokenId: 11,
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
            }
          );

          await expect(tx).to.be.revertedWith("BridgeFacade: invalid token id");
        });

        it("should not deposit if wrong fee token", async () => {
          const tx = facade.depositSBT(
            {
              feeToken: OWNER.address,
            },
            {
              token: sbt.address,
              tokenId: 10,
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
            }
          );

          await expect(tx).to.be.revertedWith("FeeManager: wrong fee token");
        });

        it("should not deposit if wrong native amount", async () => {
          const tx = facade.depositSBT(
            {
              feeToken: ETHEREUM_ADDRESS,
            },
            {
              token: sbt.address,
              tokenId: 10,
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
            },
            { value: wei("0.5") }
          );

          await expect(tx).to.be.revertedWith("FeeManager: wrong native amount");
        });

        it("should deposit if no commission", async () => {
          const tx = facade.depositSBT(
            {
              feeToken: feeTokens[2],
            },
            {
              token: sbt.address,
              tokenId: 10,
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
            }
          );

          await expect(tx).to.changeTokenBalances(sbt, [OWNER, facade, bridge], [0, 0, 0]);
        });

        it("should deposit if native commission", async () => {
          const tx = facade.depositSBT(
            {
              feeToken: feeTokens[0],
            },
            {
              token: sbt.address,
              tokenId: 10,
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
            },
            { value: feeAmounts[0] }
          );

          await expect(tx).to.changeEtherBalances([OWNER, facade, bridge], [-feeAmounts[0], feeAmounts[0], 0]);
          await expect(tx).to.changeTokenBalances(sbt, [OWNER, facade, bridge], [0, 0, 0]);
        });

        it("should deposit if erc20 commission", async () => {
          const tx = facade.depositSBT(
            {
              feeToken: feeTokens[1],
            },
            {
              token: sbt.address,
              tokenId: 10,
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
            }
          );

          await expect(tx).to.changeTokenBalances(
            feeToken,
            [OWNER, facade, bridge],
            [-feeAmounts[1], feeAmounts[1], 0]
          );
          await expect(tx).to.changeTokenBalances(sbt, [OWNER, facade, bridge], [0, 0, 0]);
        });
      });

      describe("#withdrawSBT", () => {
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

        it("should not withdraw if invalid proof", async () => {
          const withdrawParams = {
            token: sbt.address,
            tokenId: 10,
            tokenURI: "URI",
            bundle: EMPTY_BUNDLE,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
          };

          const leaf = merkleHelper.encodeLeaf(withdrawParams);

          merkleHelper.addLeaf(leaf);

          const path = merkleHelper.getPath(leaf);
          const signature = signHelper.sign(ethers.utils.randomBytes(32));

          const proof = ethers.utils.defaultAbiCoder.encode(["bytes32[]", "bytes"], [path, signature]);

          const tx = facade.withdrawSBT({ ...withdrawParams, proof });

          await expect(tx).to.be.revertedWith("Signers: invalid signature");
        });

        it("should withdraw if all conditions are met", async () => {
          const withdrawParams = {
            token: sbt.address,
            tokenId: 10,
            tokenURI: "URI",
            bundle: EMPTY_BUNDLE,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
          };

          const leaf = merkleHelper.encodeLeaf(withdrawParams);
          const proof = merkleHelper.getProof(leaf);

          const tx = facade.withdrawSBT({ ...withdrawParams, proof });

          await expect(tx).to.changeTokenBalances(sbt, [OWNER, bridge, facade, proxyAddress], [1, 0, 0, 0]);
          expect(await sbt.tokenURI(10)).to.eq("URI");
        });

        it("should withdraw if all conditions are met with wrong bundle", async () => {
          const withdrawParams = {
            token: sbt.address,
            tokenId: 10,
            tokenURI: "URI",
            bundle: {
              bundle: ethers.utils.defaultAbiCoder.encode(
                ["address[]", "uint256[]", "bytes[]"],
                [[ZERO_ADDR], [0], []]
              ),
              salt: RANDOM_SALT,
            },
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
          };

          const leaf = merkleHelper.encodeLeaf(withdrawParams);
          const proof = merkleHelper.getProof(leaf);

          const tx = facade.withdrawSBT({ ...withdrawParams, proof });

          await expect(tx).to.changeTokenBalances(sbt, [OWNER, bridge, facade, proxyAddress], [1, 0, 0, 0]);
          expect(await sbt.tokenURI(10)).to.eq("URI");
        });

        it("should withdraw bundle if all conditions are met", async () => {
          const withdrawParams = {
            token: sbt.address,
            tokenId: 10,
            tokenURI: "URI",
            bundle: bundleApprove,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
          };

          const leaf = merkleHelper.encodeLeaf(withdrawParams);
          const proof = merkleHelper.getProof(leaf);

          const tx = facade.withdrawSBT({ ...withdrawParams, proof });

          await expect(tx).to.changeTokenBalances(sbt, [OWNER, bridge, facade, proxyAddress], [0, 0, 0, 1]);
          expect(await sbt.getApproved(10)).to.eq(OWNER.address);
          expect(await sbt.tokenURI(10)).to.eq("URI");
        });
      });
    });

    describe("erc1155", () => {
      describe("#depositERC1155", () => {
        it("should not deposit if excessive native amount", async () => {
          const tx = facade.depositERC1155(
            {
              feeToken: feeTokens[0],
            },
            {
              token: erc1155.address,
              tokenId: 10,
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            },
            { value: feeAmounts[0] + 1n }
          );

          await expect(tx).to.be.revertedWith("BridgeFacade: excessive native amount");
        });

        it("should not deposit if zero token", async () => {
          const tx = facade.depositERC1155(
            {
              feeToken: COMMISSION_ADDRESS,
            },
            {
              token: ZERO_ADDR,
              tokenId: 10,
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            }
          );

          await expect(tx).to.be.revertedWith("BridgeFacade: zero token");
        });

        it("should not deposit if amount is zero", async () => {
          const tx = facade.depositERC1155(
            {
              feeToken: COMMISSION_ADDRESS,
            },
            {
              token: erc1155.address,
              tokenId: 10,
              amount: 0,
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            }
          );

          await expect(tx).to.be.revertedWith("BridgeFacade: amount is zero");
        });

        it("should not deposit if wrong fee token", async () => {
          const tx = facade.depositERC1155(
            {
              feeToken: OWNER.address,
            },
            {
              token: erc1155.address,
              tokenId: 10,
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            }
          );

          await expect(tx).to.be.revertedWith("FeeManager: wrong fee token");
        });

        it("should not deposit if wrong native amount", async () => {
          const tx = facade.depositERC1155(
            {
              feeToken: ETHEREUM_ADDRESS,
            },
            {
              token: erc1155.address,
              tokenId: 10,
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            },
            { value: wei("0.5") }
          );

          await expect(tx).to.be.revertedWith("FeeManager: wrong native amount");
        });

        it("should deposit if isWrapped=true and no commission", async () => {
          await facade.depositERC1155(
            {
              feeToken: feeTokens[2],
            },
            {
              token: erc1155.address,
              tokenId: 10,
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: true,
            }
          );

          expect(await erc1155.balanceOf(OWNER.address, 10)).to.eq(0);
          expect(await erc1155.balanceOf(facade.address, 10)).to.eq(0);
          expect(await erc1155.balanceOf(bridge.address, 10)).to.eq(0);
        });

        it("should deposit if isWrapped=false and native commission", async () => {
          const tx = facade.depositERC1155(
            {
              feeToken: feeTokens[0],
            },
            {
              token: erc1155.address,
              tokenId: 10,
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            },
            { value: feeAmounts[0] }
          );

          await expect(tx).to.changeEtherBalances([OWNER, facade, bridge], [-feeAmounts[0], feeAmounts[0], 0]);
          expect(await erc1155.balanceOf(OWNER.address, 10)).to.eq(0);
          expect(await erc1155.balanceOf(facade.address, 10)).to.eq(0);
          expect(await erc1155.balanceOf(bridge.address, 10)).to.eq(wei("10"));
        });

        it("should deposit if isWrapped=false and erc20 commission", async () => {
          const tx = facade.depositERC1155(
            {
              feeToken: feeTokens[1],
            },
            {
              token: erc1155.address,
              tokenId: 10,
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
              isWrapped: false,
            }
          );

          await expect(tx).to.changeTokenBalances(
            feeToken,
            [OWNER, facade, bridge],
            [-feeAmounts[1], feeAmounts[1], 0]
          );

          expect(await erc1155.balanceOf(OWNER.address, 10)).to.eq(0);
          expect(await erc1155.balanceOf(facade.address, 10)).to.eq(0);
          expect(await erc1155.balanceOf(bridge.address, 10)).to.eq(wei("10"));
        });
      });

      describe("#withdrawERC1155", () => {
        let bundleApprove: Bundle;
        let proxyAddress: string;

        beforeEach(async () => {
          bundleApprove = {
            bundle: ethers.utils.defaultAbiCoder.encode(
              ["address[]", "uint256[]", "bytes[]"],
              [
                [erc1155.address],
                [0],
                [erc1155.interface.encodeFunctionData("setApprovalForAll", [OWNER.address, true])],
              ]
            ),
            salt: RANDOM_SALT,
          };
          proxyAddress = await bridge.determineProxyAddress(bundleApprove.salt);
        });

        context("if deposited", () => {
          beforeEach(async () => {
            await erc1155.safeTransferFrom(OWNER.address, bridge.address, 10, wei("10"), "0x");
          });

          it("should not withdraw if invalid proof", async () => {
            const withdrawParams = {
              token: erc1155.address,
              tokenId: 10,
              tokenURI: "URI",
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              originHash: RANDOM_ORIGIN_HASH,
              receiver: OWNER.address,
              isWrapped: false,
            };

            const leaf = merkleHelper.encodeLeaf(withdrawParams);

            merkleHelper.addLeaf(leaf);

            const path = merkleHelper.getPath(leaf);
            const signature = signHelper.sign(ethers.utils.randomBytes(32));

            const proof = ethers.utils.defaultAbiCoder.encode(["bytes32[]", "bytes"], [path, signature]);

            const tx = facade.withdrawERC1155({ ...withdrawParams, proof });

            await expect(tx).to.be.revertedWith("Signers: invalid signature");
          });

          it("should withdraw if all conditions are met", async () => {
            const withdrawParams = {
              token: erc1155.address,
              tokenId: 10,
              tokenURI: "URI",
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              originHash: RANDOM_ORIGIN_HASH,
              receiver: OWNER.address,
              isWrapped: false,
            };

            const leaf = merkleHelper.encodeLeaf(withdrawParams);
            const proof = merkleHelper.getProof(leaf);

            await facade.withdrawERC1155({ ...withdrawParams, proof });

            expect(await erc1155.balanceOf(OWNER.address, 10)).to.eq(wei("10"));
            expect(await erc1155.balanceOf(facade.address, 10)).to.eq(0);
            expect(await erc1155.balanceOf(bridge.address, 10)).to.eq(0);
            expect(await erc1155.balanceOf(proxyAddress, 10)).to.eq(0);
          });

          it("should withdraw if all conditions are met with wrong bundle", async () => {
            const withdrawParams = {
              token: erc1155.address,
              tokenId: 10,
              tokenURI: "URI",
              amount: wei("10"),
              bundle: {
                bundle: ethers.utils.defaultAbiCoder.encode(["address[]", "uint256[]", "bytes[]"], [[], [], []]),
                salt: RANDOM_SALT,
              },
              originHash: RANDOM_ORIGIN_HASH,
              receiver: OWNER.address,
              isWrapped: false,
            };

            const leaf = merkleHelper.encodeLeaf(withdrawParams);
            const proof = merkleHelper.getProof(leaf);

            await facade.withdrawERC1155({ ...withdrawParams, proof });

            expect(await erc1155.balanceOf(OWNER.address, 10)).to.eq(wei("10"));
            expect(await erc1155.balanceOf(facade.address, 10)).to.eq(0);
            expect(await erc1155.balanceOf(bridge.address, 10)).to.eq(0);
            expect(await erc1155.balanceOf(proxyAddress, 10)).to.eq(0);
          });

          it("should withdraw bundle if all conditions are met", async () => {
            const withdrawParams = {
              token: erc1155.address,
              tokenId: 10,
              tokenURI: "URI",
              amount: wei("10"),
              bundle: bundleApprove,
              originHash: RANDOM_ORIGIN_HASH,
              receiver: OWNER.address,
              isWrapped: false,
            };

            const leaf = merkleHelper.encodeLeaf(withdrawParams);
            const proof = merkleHelper.getProof(leaf);

            await facade.withdrawERC1155({ ...withdrawParams, proof });

            expect(await erc1155.balanceOf(OWNER.address, 10)).to.eq(0);
            expect(await erc1155.balanceOf(facade.address, 10)).to.eq(0);
            expect(await erc1155.balanceOf(bridge.address, 10)).to.eq(0);
            expect(await erc1155.balanceOf(proxyAddress, 10)).to.eq(wei("10"));
            expect(await erc1155.isApprovedForAll(proxyAddress, OWNER.address)).to.be.true;
          });
        });

        context("if burned", () => {
          beforeEach(async () => {
            await erc1155.burnFrom(OWNER.address, 10, wei("10"));
          });

          it("should withdraw if all conditions are met", async () => {
            const withdrawParams = {
              token: erc1155.address,
              tokenId: 10,
              tokenURI: "URI",
              amount: wei("10"),
              bundle: EMPTY_BUNDLE,
              originHash: RANDOM_ORIGIN_HASH,
              receiver: OWNER.address,
              isWrapped: true,
            };

            const leaf = merkleHelper.encodeLeaf(withdrawParams);
            const proof = merkleHelper.getProof(leaf);

            await facade.withdrawERC1155({ ...withdrawParams, proof });

            expect(await erc1155.balanceOf(OWNER.address, 10)).to.eq(wei("10"));
            expect(await erc1155.balanceOf(facade.address, 10)).to.eq(0);
            expect(await erc1155.balanceOf(bridge.address, 10)).to.eq(0);
            expect(await erc1155.balanceOf(proxyAddress, 10)).to.eq(0);
            expect(await erc1155.uri(10)).to.eq("URI");
          });

          it("should withdraw bundle if all conditions are met", async () => {
            const withdrawParams = {
              token: erc1155.address,
              tokenId: 10,
              amount: wei("10"),
              tokenURI: "URI",
              bundle: bundleApprove,
              originHash: RANDOM_ORIGIN_HASH,
              receiver: OWNER.address,
              isWrapped: true,
            };

            const leaf = merkleHelper.encodeLeaf(withdrawParams);
            const proof = merkleHelper.getProof(leaf);

            await facade.withdrawERC1155({ ...withdrawParams, proof });

            expect(await erc1155.balanceOf(OWNER.address, 10)).to.eq(0);
            expect(await erc1155.balanceOf(facade.address, 10)).to.eq(0);
            expect(await erc1155.balanceOf(bridge.address, 10)).to.eq(0);
            expect(await erc1155.balanceOf(proxyAddress, 10)).to.eq(wei("10"));
            expect(await erc1155.isApprovedForAll(proxyAddress, OWNER.address)).to.be.true;
            expect(await erc1155.uri(10)).to.eq("URI");
          });
        });
      });
    });

    describe("native", () => {
      describe("#depositNative", () => {
        it("should not deposit if no funds to deposit", async () => {
          const tx = facade.depositNative(
            {
              feeToken: feeTokens[0],
            },
            {
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
            },
            { value: feeAmounts[0] }
          );

          await expect(tx).to.be.revertedWith("BridgeFacade: no funds to deposit");
        });

        it("should not deposit if wrong fee token", async () => {
          const tx = facade.depositNative(
            {
              feeToken: OWNER.address,
            },
            {
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
            },
            { value: wei("10") }
          );

          await expect(tx).to.be.revertedWith("FeeManager: wrong fee token");
        });

        it("should not deposit if wrong native amount", async () => {
          const tx = facade.depositNative(
            {
              feeToken: ETHEREUM_ADDRESS,
            },
            {
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
            },
            { value: wei("0.5") }
          );

          await expect(tx).to.be.revertedWith("FeeManager: wrong native amount");
        });

        it("should deposit if no commission", async () => {
          const tx = facade.depositNative(
            {
              feeToken: feeTokens[2],
            },
            {
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
            },
            { value: wei("10") }
          );

          await expect(tx).to.changeEtherBalances([OWNER, facade, bridge], [-wei("10"), 0, wei("10")]);
        });

        it("should deposit if native commission", async () => {
          const tx = facade.depositNative(
            {
              feeToken: feeTokens[0],
            },
            {
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
            },
            { value: feeAmounts[0] + wei("10") }
          );

          await expect(tx).to.changeEtherBalances(
            [OWNER, facade, bridge],
            [-feeAmounts[0] - wei("10"), feeAmounts[0], wei("10")]
          );
        });

        it("should deposit if erc20 commission", async () => {
          const tx = facade.depositNative(
            {
              feeToken: feeTokens[1],
            },
            {
              bundle: EMPTY_BUNDLE,
              network: CHAIN_NAME,
              receiver: RECEIVER,
            },
            { value: wei("10") }
          );

          await expect(tx).to.changeTokenBalances(
            feeToken,
            [OWNER, facade, bridge],
            [-feeAmounts[1], feeAmounts[1], 0]
          );
          await expect(tx).to.changeEtherBalances([OWNER, facade, bridge], [-wei("10"), 0, wei("10")]);
        });
      });

      describe("#withdrawNative", () => {
        let bundleTransfer: Bundle;
        let proxyAddress: string;

        beforeEach(async () => {
          await setBalance(bridge.address, wei("10"));

          bundleTransfer = {
            bundle: ethers.utils.defaultAbiCoder.encode(
              ["address[]", "uint256[]", "bytes[]"],
              [[OWNER.address], [wei("5")], ["0x"]]
            ),
            salt: RANDOM_SALT,
          };
          proxyAddress = await bridge.determineProxyAddress(bundleTransfer.salt);
        });

        it("should not withdraw if invalid proof", async () => {
          const withdrawParams = {
            amount: wei("10"),
            bundle: EMPTY_BUNDLE,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
          };

          const leaf = merkleHelper.encodeLeaf(withdrawParams);

          merkleHelper.addLeaf(leaf);

          const path = merkleHelper.getPath(leaf);
          const signature = signHelper.sign(ethers.utils.randomBytes(32));

          const proof = ethers.utils.defaultAbiCoder.encode(["bytes32[]", "bytes"], [path, signature]);

          const tx = facade.withdrawNative({ ...withdrawParams, proof });

          await expect(tx).to.be.revertedWith("Signers: invalid signature");
        });

        it("should withdraw if all conditions are met", async () => {
          const withdrawParams = {
            amount: wei("10"),
            bundle: EMPTY_BUNDLE,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
          };

          const leaf = merkleHelper.encodeLeaf(withdrawParams);
          const proof = merkleHelper.getProof(leaf);

          const tx = facade.withdrawNative({ ...withdrawParams, proof });

          await expect(tx).to.changeEtherBalances([OWNER, bridge, facade, proxyAddress], [wei("10"), -wei("10"), 0, 0]);
        });

        it("should withdraw if all conditions are met with wrong bundle", async () => {
          const withdrawParams = {
            amount: wei("10"),
            bundle: {
              bundle: ethers.utils.defaultAbiCoder.encode(
                ["address[]", "uint256[]", "bytes[]"],
                [[feeToken.address], [wei("5")], ["0x"]]
              ),
              salt: RANDOM_SALT,
            },
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
          };

          const leaf = merkleHelper.encodeLeaf(withdrawParams);
          const proof = merkleHelper.getProof(leaf);

          const tx = facade.withdrawNative({ ...withdrawParams, proof });

          await expect(tx).to.changeEtherBalances([OWNER, bridge, facade, proxyAddress], [wei("10"), -wei("10"), 0, 0]);
        });

        it("should withdraw bundle if all conditions are met", async () => {
          const withdrawParams = {
            amount: wei("10"),
            bundle: bundleTransfer,
            originHash: RANDOM_ORIGIN_HASH,
            receiver: OWNER.address,
          };

          const leaf = merkleHelper.encodeLeaf(withdrawParams);
          const proof = merkleHelper.getProof(leaf);

          const tx = facade.withdrawNative({ ...withdrawParams, proof });

          await expect(tx).to.changeEtherBalances([OWNER, bridge, facade, proxyAddress], [wei("5"), -wei("10"), 0, 0]);
          await expect(tx).to.changeTokenBalances(weth9, [OWNER, bridge, facade, proxyAddress], [0, 0, 0, wei("5")]);
        });
      });
    });
  });
});
