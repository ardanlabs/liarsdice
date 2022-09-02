import React, { useEffect, useState } from 'react'
import Button from './button'
import MetamaskLogo from './icons/metamask'
import { utils } from 'ethers'
import axios, { AxiosError } from 'axios'
import { toast } from 'react-toastify'
import { capitalize } from '../utils/capitalize'
import getNowDate from '../utils/getNowDate'
import { useNavigate } from 'react-router-dom'
import { getAppConfig } from '..'
import { token } from '../utils/axiosConfig'
import useEthersConnection from './hooks/useEthersConnection'

// Login component.
// Doesn't receives parameters
// Outputs the Html output for login in.
// Handles logic for login in and also connecting to the game engine.
function Login() {
  // extracts navigate from useNavigate Hook
  const navigate = useNavigate()

  /* Extracts functions from useEthersConnection Hook
   *  This hook handles all connections to ethers.js
   */
  const { setSigner, account, signer, setAccount, provider } =
    useEthersConnection()
  // Sets local state to trigger re-render when load is complete.
  const [loading, setLoading] = useState(true)

  // Prompts user to connect to metamask usign the provider
  // and registers it to UseEthersConnection hook
  const init = async () => {
    const signer = provider.getSigner()
    const signerAddress = await signer.getAddress()

    setAccount(signerAddress)
    setSigner(signer)
  }

  // The Effect Hook lets you perform side effects in function components
  // In this case we use it to handle what happens when metamask accounts change.
  useEffect(() => {
    // handles what happen
    function handleAccountsChanged(accounts: string[]) {
      if (accounts.length === 0) {
        // MetaMask is locked or the user has not connected any accounts
        setLoading(true)
      } else if (accounts[0] !== account) {
        init().then(() => {
          setLoading(false)
        })
      }
    }

    // Note that this event is emitted on page load.
    // If the array of accounts is non-empty, you're already connected.
    window.ethereum.on('accountsChanged', handleAccountsChanged)

    init().then(() => {
      setLoading(false)
    })
    // An empty dependecies array triggers useEffect only on the first render of the component
    // We disable the next line so eslint doens't complain about missing dependencies.
    // eslint-disable-next-line
  }, [])

  // The Effect Hook lets you perform side effects in function components
  // In this case we handle what happen if you're already log when you enter the app
  useEffect(() => {
    function handleIfLogged() {
      if (token() && account) {
        getAppConfig.then((response) => {
          navigate('/mainRoom', { state: { ...response } })
        })
      }
    }

    handleIfLogged()

    // UseEffect will trigger every time account changes values.
    // We disable the next line so eslint doens't complain about missing dependencies.
    // eslint-disable-next-line
  }, [account])

  // Parsed a document into Uint8Array
  function parseDocToUint8Array(doc: object): Uint8Array {
    // Marshal the transaction to a string and convert the string to bytes.
    const marshal = JSON.stringify(doc)
    const marshalBytes = utils.toUtf8Bytes(marshal)

    // Hash the transaction data into a 32 byte array. This will provide
    // a data length consistency with all transactions.
    const txHash = utils.keccak256(marshalBytes)
    const bytes = utils.arrayify(txHash)

    return bytes
  }

  // Connects to the game engine, and stores the token in the sessionStorage.
  // Takes an object of a date_time document and a signature
  function connectToGameEngine(data: { date_time: string; sig: string }) {
    axios
      .post('http://localhost:3000/v1/game/connect', data)
      .then((connectResponse) => {
        window.sessionStorage.setItem(
          'token',
          `bearer ${connectResponse.data.token}`,
        )
        getAppConfig.then((getConfigResponse) => {
          navigate('/mainRoom', { state: { ...getConfigResponse } })
        })
      })
      .catch((error: AxiosError) => {
        let errorMessage = (error as any).response.data.error.replace(
          / \[.+\]/gm,
          '',
        )

        toast(
          <div style={{ textAlign: 'start' }}>{capitalize(errorMessage)}</div>,
        )

        console.group()
        console.error('Error:', (error as any).response.data.error)
        console.groupEnd()
      })
  }

  // Handles click on Metamask button is clicked
  // We trigger the connection to the browser wallet
  async function handleConnectAccount() {
    await provider.send('eth_requestAccounts', [])

    const signer = provider.getSigner()

    const signerAddress = await signer.getAddress()

    setAccount(signerAddress)

    setSigner(signer)
  }

  // Handles click on sign transaction
  // Performs the following actions:
  // Creates a document to sign.
  // Signes the document.
  // Connects to game engine.
  const signTransaction = () => {
    toast.info('Connecting to game engine')

    const date = getNowDate()

    const doc = { date_time: date }

    const parsedDoc = parseDocToUint8Array(doc)

    // We sign the data. The underlying code will apply the Ardan stamp and
    // ID to the signature thanks to changes made to the ether.js api.
    signer?.signMessage(parsedDoc).then((signerResponse: any) => {
      const data = { ...doc, sig: signerResponse }
      // After the data is signed we connect to the game Engine sending the signature and the signed document.
      connectToGameEngine(data)
    })
  }
  // Renders
  return (
    <div
      className="container-fluid d-flex align-items-center justify-content-center px-0 flex-column"
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
          {account && !loading ? (
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

export default Login
