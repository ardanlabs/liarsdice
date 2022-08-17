import React from 'react'
import Button from './button'
import MetamaskLogo from './icons/metamask'
import { useEthers } from '@usedapp/core'
import MainRoom from './mainRoom'

export default function Login() {
  const { account, activateBrowserWallet } = useEthers()

  function handleConnectAccount() {
    activateBrowserWallet()
  }
  return account?.length ? (
    <MainRoom />
  ) : (
    <div
      id="login__wrapper"
      className="d-flex align-items-start justify-content-center flex-column mt-10"
    >
      <h2>
        <strong> Connect your wallet </strong>
      </h2>
      Or you can also select a provider to create one.
      <div id="wallets__wrapper" className="mt-4">
        <Button
          {...{
            id: 'metamask__wrapper',
            clickHandler: handleConnectAccount,
            classes: 'd-flex align-items-center pa-4',
          }}
        >
          <MetamaskLogo {...{ width: '50px', height: '50px' }} />
          <span className="ml-4"> Metamask </span>
        </Button>
      </div>
    </div>
  )
}
