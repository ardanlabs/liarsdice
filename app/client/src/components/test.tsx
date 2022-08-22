import React from 'react'
import { utils } from 'ethers'
import { useEthers } from '@usedapp/core'
import axios, { AxiosError } from 'axios'

const Test = () => {
  const { library } = useEthers()

  const signMessage = () => {
    var doc = { name: 'bill', status: 'ok' }

    const signer = library?.getSigner()

    // Marshal the transaction to a string and convert the string to bytes.
    const marshal = JSON.stringify(doc)
    const marshalBytes = utils.toUtf8Bytes(marshal)

    // Hash the transaction data into a 32 byte array. This will provide
    // a data length consistency with all transactions.
    const txHash = utils.keccak256(marshalBytes)
    const bytes = utils.arrayify(txHash)

    // Now sign the data. The underlying code will apply the Ardan stamp and
    // ID to the signature thanks to changes made to the ether.js api.

    signer?.signMessage(bytes).then((response: any) => {
      console.log(response)
      const data = { ...doc, sig: response }
      axios
        .post('http://localhost:3000/v1/game/test', data)
        .then((response) => {
          console.log(response, 'in')
        })
        .catch((error: AxiosError) => {
          console.log(error)
        })
    })
  }

  return <button onClick={() => signMessage()}>Test signing</button>
}

export default Test
