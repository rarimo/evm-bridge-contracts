const secp256k1 = require("secp256k1");

function rawSign(hash, hexKey) {
  const privateKey = Buffer.from(hexKey.slice(2), "hex");
  const rawSignature = secp256k1.ecdsaSign(
    Uint8Array.from(Buffer.from(hash.slice(2), "hex")),
    Uint8Array.from(privateKey)
  );

  return web3.utils.bytesToHex(rawSignature.signature) + (rawSignature.recid + 27).toString(16);
}

module.exports = {
  rawSign,
};
