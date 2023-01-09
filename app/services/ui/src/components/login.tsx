import React, { useContext, useEffect, useState } from 'react'
import Button from './button'
import MetamaskLogo from './icons/metamask'
import { toast } from 'react-toastify'
import getNowDate from '../utils/getNowDate'
import { useNavigate } from 'react-router-dom'
import { getAppConfig } from '..'
import { token } from '../utils/axiosConfig'
import useEthersConnection from './hooks/useEthersConnection'
import useGame from './hooks/useGame'
import { WalletContext } from '@viaprotocol/web3-wallets'
import LogOutIcon from './icons/logout'
import CoinBaseLogo from './icons/coinbase'
import WalletConnectLogo from './icons/walletconnect'
import { Web3Provider } from '@ethersproject/providers'
import Wallets from '../utils/wallets'

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
  const { account, setAccount, setSigner } = useEthersConnection()

  const { connect, isConnected, disconnect, address, signMessage, provider } =
    useContext(WalletContext)

  // Sets local state to trigger re-render when load is complete.
  const [loading, setLoading] = useState(true)

  // Prompts user to connect to metamask usign the provider
  // and registers it to useEthersConnection hook
  async function init() {
    if (provider instanceof Web3Provider) {
      const signer = provider.getSigner()
      setSigner(signer)
    }

    setAccount(address)
  }

  // ===========================================================================

  function accountsChangeUEFn() {
    // handleAccountsChanged handles what happen when metamask account changes
    // If the array of accounts is non-empty, you're already connected.
    function handleAccountsChanged(accounts: string[]) {
      if (accounts.length === 0) {
        // MetaMask is locked or the user has not connected any accounts
        setLoading(true)
        return
      }
    }

    function connectFn() {
      init().then(() => setLoading(false))
    }
    // Note that this event is emitted on page load.
    window.ethereum.on('accountsChanged', handleAccountsChanged)
    // This event checks when the provider becomes able to submit
    // RPC requests to a chain.
    window.ethereum.on('connect', connectFn)
    init().then(() => setLoading(false))
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
        navigate('/mainRoom', {
          state: { ...response, reload: true },
          replace: true,
        })
      })
    }
  }

  // Next line is disabled so eslint doens't complain about missing dependencies.

  // eslint-disable-next-line
  useEffect(loggedUEFn, [account, isConnected])

  function connectPageOnLoadUEFn() {
    const connectWalletOnPageLoad = async () => {
      const connectedWallet: string | null =
        localStorage.getItem('connectedWallet')
      if (connectedWallet) {
        const isWalletConnected =
          JSON.parse(connectedWallet).name === `${Wallets.Metamask.name}` ||
          JSON.parse(connectedWallet).name === `${Wallets.Coinbase.name}` ||
          JSON.parse(connectedWallet).name === `${Wallets.WalletConnect.name}`

        if (isWalletConnected) {
          try {
            await connect(
              JSON.parse(localStorage.getItem('connectedWallet') as string),
            ).then(() => {
              init().then(() => setLoading(false))
            })
          } catch (ex) {
            console.log(ex)
          }
        }
      }
    }
    connectWalletOnPageLoad()
  }

  useEffect(connectPageOnLoadUEFn, [])

  // ===========================================================================
  //
  // End of hooks.
  //
  // ===========================================================================

  // handleConnectAccount takes care of the connection to the browser wallet.
  async function handleConnectAccount(wallet: { name: any; chainId: any }) {
    await connect(wallet).then(() => {
      init().then(() => {
        window.localStorage.setItem('connectedWallet', JSON.stringify(wallet))
        setLoading(false)
      })
    })
  }

  // signTransaction handles click on sign transaction
  // Creates a document to sign, signs the document and connects to game engine.
  function signTransaction() {
    toast.info('Connecting to game engine')

    const date = getNowDate()

    const doc = { address: address as string, dateTime: date }

    const parsedDoc = JSON.stringify(doc)

    // signerFn connects to the game Engine sending the signature and the signed document.
    const signerFn = (signerResponse: any) => {
      const data = { ...doc, sig: signerResponse }
      connectToGameEngine(data)
    }

    // signer.signmessage signs the data. The underlying code will apply the Ardan stamp and
    // ID to the signature thanks to changes made to the ether.js api.

    signMessage(parsedDoc).then(signerFn)
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
          {isConnected && !loading ? (
            <div className="d-flex justify-content-evenly">
              <span className="ml-2">Wallet {address} connected</span>
              <div
                onClick={() => disconnect()}
                className="mx-2"
                style={{ cursor: 'pointer' }}
              >
                <LogOutIcon />
              </div>
            </div>
          ) : (
            <div>
              <Button
                {...{
                  id: 'metamask__wrapper',
                  clickHandler: () => handleConnectAccount(Wallets.Metamask),
                  classes: 'd-flex align-items-center',
                }}
              >
                <MetamaskLogo {...{ width: '50px', height: '50px' }} />
                <span className="ml-4"> Metamask </span>
              </Button>
              <Button
                {...{
                  id: 'coinbase__wrapper',
                  clickHandler: () => handleConnectAccount(Wallets.Coinbase),
                  classes: 'd-flex align-items-center',
                }}
              >
                <CoinBaseLogo {...{ width: '50px', height: '50px' }} />
                <span className="ml-4"> Coinbase </span>
              </Button>
              <Button
                {...{
                  id: 'walletConnect__wrapper',
                  clickHandler: () =>
                    handleConnectAccount(Wallets.WalletConnect),
                  classes: 'd-flex align-items-center',
                }}
              >
                <WalletConnectLogo />
                <span className="ml-4"> Wallet Connect</span>
              </Button>
            </div>
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
