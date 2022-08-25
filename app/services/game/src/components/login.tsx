import React, { useState } from 'react'
import Button from './button'
import MetamaskLogo from './icons/metamask'
import { useEthers } from '@usedapp/core'
import MainRoom from './mainRoom'
import { utils } from 'ethers'
import axios, { AxiosError } from 'axios'
import { toast } from 'react-toastify'
import { capitalize } from '../utils/capitalize'
import getNowDate from '../utils/getNowDate'
import { token } from '../utils/axiosConfig'

export default function Login() {
  const { account, library, activateBrowserWallet } = useEthers()
  const [isTokenSet, setIsTokenSet] = useState(false)
  function handleConnectAccount() {
    activateBrowserWallet()
  }
  const signTransaction = () => {
    toast.info('Connecting to game engine')
    const date = getNowDate()

    let doc = { date_time: date }

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
        .post('http://localhost:3000/v1/game/connect', data)
        .then((response) => {
          toast.info('Connected to game engine')
          window.sessionStorage.setItem(
            'token',
            `bearer ${response.data.token}`,
          )
          setIsTokenSet(true)
        })
        .catch((error: AxiosError) => {
          setIsTokenSet(false)
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
  return account && token() ? (
    <MainRoom />
  ) : (
    <div
      style={{
        display: 'flex',
        alignItems: 'center',
        height: 'calc(100vh - 70px)',
      }}
    >
      <div
        id="login__wrapper"
        className="d-flex align-items-start justify-content-center flex-column mt-10"
      >
        <h2>
          <strong> Connect your wallet </strong>
        </h2>
        Or you can also select a provider to create one.
        <div id="wallets__wrapper" className="mt-4">
          {account ? (
            <div className="d-flex">
              <span className="ml-2">Wallet {account} connected</span>
            </div>
          ) : (
            <Button
              {...{
                id: 'metamask__wrapper',
                clickHandler: handleConnectAccount,
                classes: 'd-flex align-items-center',
              }}
            >
              <MetamaskLogo {...{ width: '50px', height: '50px' }} />
              <span className="ml-4"> Metamask </span>
            </Button>
          )}
        </div>
        <div id="wallets__wrapper" className="mt-4">
          <Button
            {...{
              id: 'metamask__wrapper',
              clickHandler: signTransaction,
              classes: 'd-flex align-items-center',
            }}
          >
            <>Sign into app</>
          </Button>
        </div>
      </div>
    </div>
  )
}