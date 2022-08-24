import React from 'react'
import { useEthers } from '@usedapp/core'
import Button from './button'
import LogOutIcon from './icons/logout'

interface SignOutProps {
  disabled: boolean
}

const SignOut = (props: SignOutProps) => {
  const { disabled } = props
  const { account, deactivate } = useEthers()

  function handleDisconnectAccount() {
    deactivate()
    window.sessionStorage.removeItem('token')
  }
  return account ? (
    <Button
      {...{
        disabled: disabled,
        id: 'metamask__wrapper',
        clickHandler: handleDisconnectAccount,
        classes: 'd-flex align-items-center pa-4',
      }}
    >
      <LogOutIcon />
    </Button>
  ) : null
}

export default SignOut
