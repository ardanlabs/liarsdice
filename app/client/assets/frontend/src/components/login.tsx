import React from 'react'
import Button from './button'
import MetamaskLogo from './icons/metamask'
import { useEthers, useEtherBalance } from "@usedapp/core";
import MainRoom from './mainRoom'

export default function Login() {
  const { account, deactivate, activateBrowserWallet } = useEthers()
  const etherBalance = useEtherBalance(account)
  function handleConnectWallet() {
    activateBrowserWallet();
  }
  function handleDisconnectWallet() {
    deactivate();
  }

  return account ? (
    <div>
      <MainRoom etherBalance={etherBalance} show={Boolean(account)} />
      <Button {...{ id: 'metamask__wrapper', clickHandler: handleDisconnectWallet, classes: 'd-flex align-items-center pa-4'}}>
        <span className="ml-4"> Disconnect </span>
      </Button>
    </div>
  ) : (
    <div
      id="login__wrapper"
      className="d-flex align-items-start justify-content-center flex-column"
    >
      <h2>
        <strong> Connect your wallet </strong>
      </h2>
      Or you can also select a provider to create one.
      <div id="wallets__wrapper" className="mt-4">
        <Button {...{ id: 'metamask__wrapper', clickHandler: handleConnectWallet, classes: 'd-flex align-items-center pa-4'}}>
          <MetamaskLogo {...{ width: '50px', height: '50px' }} />
          <span className="ml-4"> Metamask </span>
        </Button>
      </div>
    </div>
  );
}