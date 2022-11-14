import { utils } from 'ethers'

// Parsed a document into Uint8Array
export default function docToUint8Array(doc: object): Uint8Array {
  // Marshal the transaction to a string and convert the string to bytes.
  const marshal = JSON.stringify(doc)
  const marshalBytes = utils.toUtf8Bytes(marshal)

  // Hash the transaction data into a 32 byte array. This will provide
  // a data length consistency with all transactions.
  const txHash = utils.keccak256(marshalBytes)
  const bytes = utils.arrayify(txHash)

  return bytes
}
