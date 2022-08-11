import React from "react"
import Button from "./button"
import LogOutIcon from './icons/logout'
import { useEthers } from "@usedapp/core";


function Footer () {
  const { account } = useEthers()
  const { deactivate } = useEthers()
  function handleDisconnectWallet() {
    deactivate()
  }

  return account ? (
    <footer style={{backgroundColor: 'var(--modals)',position: 'fixed', bottom: '0', height: '70px', width: '100%', display: 'flex', justifyContent: 'space-around'}}>
      <Button {...{ id: 'metamask__wrapper', clickHandler: handleDisconnectWallet, classes: 'd-flex align-items-center pa-4'}}>
        <LogOutIcon />
      </Button>
    </footer>
  ) : null
}

export default Footer