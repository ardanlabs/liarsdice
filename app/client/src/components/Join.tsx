import React from 'react'
import { utils } from 'ethers'
import { useEthers } from '@usedapp/core'
import axios, { AxiosError } from 'axios'
import Button from './button'
import { toast } from 'react-toastify'
import { capitalize } from '../utils/capitalize'

interface JoinProps {
  disabled: boolean
}

const Join = (props: JoinProps) => {
  const { disabled } = props
  const { library } = useEthers()

  const signMessage = () => {
    toast.info('Joining game...')
    const now = new Date()
    const dd = String(now.getDate()).padStart(2, '0')
    const mm = String(now.getMonth() + 1).padStart(2, '0') //January is 0!
    const yyyy = now.getFullYear()
    const hours = now.getHours()
    const minutes = now.getMinutes()
    const seconds = now.getSeconds()
    // We create the specific date to send to the signature
    const date = yyyy + mm + dd + hours + minutes + seconds

    var doc = { date_time: date }

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
      const data = { ...doc, sig: response }
      axios
        .post('http://localhost:3000/v1/game/join', data)
        .then((response) => {
          toast.info('Welcome to the game')
          window.sessionStorage.setItem(
            'token',
            `bearer ${response.data.token}`,
          )
        })
        .catch((error: AxiosError) => {
          let errorMessage = (error as any).response.data.error.replace(
            / \[.+\]/gm,
            '',
          )
          toast.error(
            <div style={{ textAlign: 'start' }}>
              {capitalize(errorMessage)}
            </div>,
          )
          console.group()
          console.error('Error:', (error as any).response.data.error)
          console.groupEnd()
        })
    })
  }
  return (
    <Button
      disabled={disabled}
      classes="join__buton"
      clickHandler={() => signMessage()}
    >
      <span>Join Game</span>
    </Button>
  )
}

export default Join
