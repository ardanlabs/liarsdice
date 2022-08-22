import React from 'react'
import { utils } from 'ethers'
import { useEthers } from '@usedapp/core'

const Test = () => {
  const { library } = useEthers()

  const signMessage = (message: string) => {
    const signer = library?.getSigner()
    // Marshal the transaction to a string and convert the string to bytes.
    const marshal = JSON.stringify(message)
    const marshalBytes = utils.toUtf8Bytes(marshal)
    // Hash the transaction data into a 32 byte array. This will provide
    // a data length consistency with all transactions.
    const txHash = utils.keccak256(marshalBytes)
    const bytes = utils.arrayify(txHash)

    // Now sign the data. The underlying code will apply the Ardan stamp and
    // ID to the signature thanks to changes made to the ether.js api.

    signer?.signMessage(bytes).then((response: any) => {
      console.log(response)
    })
  }
  const message = JSON.stringify({ name: 'bill', status: 'ok' })

  return <button onClick={() => signMessage(message)}>Test signing</button>
}

export default Test
