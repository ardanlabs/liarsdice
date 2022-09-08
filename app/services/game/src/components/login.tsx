import React, { useEffect, useState } from 'react'
import Button from './button'
import MetamaskLogo from './icons/metamask'
import { toast } from 'react-toastify'
import getNowDate from '../utils/getNowDate'
import { useNavigate } from 'react-router-dom'
import { getAppConfig } from '..'
import { token } from '../utils/axiosConfig'
import useEthersConnection from './hooks/useEthersConnection'
import docToUint8Array from '../utils/docToUint8Array'
import useGame from './hooks/useGame'

// Login component.
// Doesn't receives parameters
// Outputs the Html output for login in.
// Handles logic for login in and also connecting to the game engine.
function Login() {
  // ===========================================================================
  // Hooks setup.

  // Extract connectToGameEngine from useGame hook
  const { connectToGameEngine } = useGame()

  // Extracts navigate from useNavigate Hook
  const navigate = useNavigate()

  // Extracts functions from useEthersConnection Hook
  // useEthersConnection hook handles all connections to ethers.js
  const { setSigner, account, signer, setAccount, provider } =
    useEthersConnection()

  // Sets local state to trigger re-render when load is complete.
  const [loading, setLoading] = useState(true)

  // ===========================================================================

  const accountsChangeUEFn = () => {
    // Prompts user to connect to metamask usign the provider
    // and registers it to UseEthersConnection hook
    async function init() {
      const signer = provider.getSigner()
      if (signer._address) {
        const signerAddress = await signer.getAddress()
        setAccount(signerAddress)
      }

      setSigner(signer)
    }

    // handleAccountsChanged handles what happen when metamask account changes
    // If the array of accounts is non-empty, you're already connected.
    function handleAccountsChanged(accounts: string[]) {
      if (accounts.length === 0) {
        // MetaMask is locked or the user has not connected any accounts
        setLoading(true)
        return
      }
    }

    function connectFn(connectInfo: string) {
      init().then(() => setLoading(false))
    }
    // Note that this event is emitted on page load.
    window.ethereum.on('accountsChanged', handleAccountsChanged)
    // This event checks when the provider becomes able to submit
    // RPC requests to a chain.
    window.ethereum.on('connect', connectFn)
  }

  // The Effect Hook lets you perform side effects in function components
  // In this case we use it to handle what happens when metamask accounts change.
  // An empty dependecies array triggers useEffect only on the first render
  // of the component. We disable the next line so eslint doens't complain about
  // missing dependencies.

  // eslint-disable-next-line
  useEffect(accountsChangeUEFn, [])

  // ===========================================================================

  // loggedUEFn handles what happen if you're already log when you enter the app
  const loggedUEFn = () => {
    if (token() && account) {
      getAppConfig.then((response) => {
        navigate('/mainRoom', { state: { ...response }, replace: true })
      })
    }
  }

  // Next line is disabled so eslint doens't complain about missing dependencies.

  // eslint-disable-next-line
  useEffect(loggedUEFn, [account])

  // ===========================================================================
  //
  // End of hooks.
  //
  // ===========================================================================

  // handleConnectAccount takes care of the connection to the browser wallet.
  async function handleConnectAccount() {
    await provider.send('eth_requestAccounts', []).then((accounts: string) => {
      const signer = provider.getSigner()
      setAccount(accounts[0])

      setSigner(signer)
      setLoading(false)
    })
  }

  // signTransaction handles click on sign transaction
  // Creates a document to sign, signs the document and connects to game engine.
  function signTransaction() {
    toast.info('Connecting to game engine')

    const date = getNowDate()

    const doc = { dateTime: date }

    const parsedDoc = docToUint8Array(doc)

    // signerFn connects to the game Engine sending the signature and the signed document.
    const signerFn = (signerResponse: any) => {
      const data = { ...doc, sig: signerResponse }
      connectToGameEngine(data)
    }

    // signer.signmessage signs the data. The underlying code will apply the Ardan stamp and
    // ID to the signature thanks to changes made to the ether.js api.
    signer?.signMessage(parsedDoc).then(signerFn)
  }

  // ===========================================================================

  // Renders this final markup.
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
