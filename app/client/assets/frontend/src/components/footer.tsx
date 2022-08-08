import React from "react"
import Button from "./button"
import LogOutIcon from './icons/logout'
import { useEthers } from "@usedapp/core";

function Footer () {
  const { deactivate } = useEthers()
  function handleDisconnectWallet() {
    deactivate()
  }

  return (
    <footer style={{backgroundColor: 'var(--modals)', height: '70px', width: '100%', display: 'flex', justifyContent: 'space-around'}}>
      <Button {...{ id: 'metamask__wrapper', clickHandler: handleDisconnectWallet, classes: 'd-flex align-items-center pa-4'}}>
        <LogOutIcon />
      </Button>
    </footer>
  )
}

export default Footer