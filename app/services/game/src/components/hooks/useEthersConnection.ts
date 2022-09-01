import { ethers, Signer } from 'ethers'
import { useContext } from 'react'
import { EthersContext } from '../../contexts/ethersContext'

const useEthersConnection = () => {
  const { ethersConnection, setEthersConnection } = useContext(EthersContext)
  const provider = new ethers.providers.Web3Provider(window.ethereum, 'any')

  /**
   * Signer setter.
   **/
  const setSigner = (signer: Signer) => {
    const newContext = ethersConnection
    newContext.signer = signer

    setEthersConnection(newContext)
  }
  /**
   * Signer getter.
   **/
  const signer = ethersConnection.signer
  /**
   * Signer account.
   **/
  const setAccount = (account: string | undefined) => {
    const newContext = ethersConnection
    newContext.account = account

    setEthersConnection(newContext)
  }
  /**
   * Account getter.
   **/
  const account = ethersConnection.account
    ? ethersConnection.account
    : undefined

  const setNetwork = (network: object) => {
    const newContext = ethersConnection
    newContext.network = network

    setEthersConnection(newContext)
  }
  const switchNetwork = (network: Partial<{ chainId: string }>) => {
    window.ethereum.request({
      method: 'wallet_switchEthereumChain',
      params: [network],
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
