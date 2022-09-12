/* ************useGameHook************

  This hook provides all functions needed to connect to the web3 provider.

**************************************** */

import { ethers, Signer } from 'ethers'
import { useContext } from 'react'
import { EthersContext } from '../../contexts/ethersContext'

function useEthersConnection() {
  const { ethersConnection, setEthersConnection } = useContext(EthersContext)
  const provider = new ethers.providers.Web3Provider(window.ethereum, 'any')

  // Signer setter.
  function setSigner(signer: Signer) {
    const newContext = ethersConnection
    newContext.signer = signer

    setEthersConnection(newContext)
  }

  // Signer getter.
  const signer = ethersConnection.signer

  // Signer account
  function setAccount(account: string | undefined) {
    const newContext = ethersConnection
    newContext.account = account

    setEthersConnection(newContext)
  }
  // Account getter.
  const account = ethersConnection.account
    ? ethersConnection.account
    : undefined

  // Network setter
  function setNetwork(network: object) {
    const newContext = ethersConnection
    newContext.network = network

    setEthersConnection(newContext)
  }

  // Switches the current network that metamask is connected to
  function switchNetwork(network: Partial<{ chainId: string }>) {
    window.ethereum
      .request({
        method: 'wallet_switchEthereumChain',
        params: [network],
      })
      .catch((response: any) => {
        console.error(response)
      })
  }

  return {
    setSigner,
    signer,
    setAccount,
    account,
    setNetwork,
    switchNetwork,
    provider,
  }
}

export default useEthersConnection
