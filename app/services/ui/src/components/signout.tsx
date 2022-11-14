import React from 'react'
import Button from './button'
import LogOutIcon from './icons/logout'
import useEthersConnection from './hooks/useEthersConnection'
import { useNavigate } from 'react-router-dom'
import { SignOutProps } from '../types/props.d'

// SignOut component
function SignOut(props: SignOutProps) {
  // Extracts props.
  const { disabled } = props

  // Extracts account and setAccount from useEthersConnection hook.
  const { account, setAccount } = useEthersConnection()

  // Extracts router navigation functionality from useNavigate hook.
  const navigate = useNavigate()

  // ===========================================================================
  // handleDisconnectAccount disconnects the user and deletes the token.
  function handleDisconnectAccount() {
    setAccount(null)
    window.sessionStorage.removeItem('token')
    navigate('/')
  }

  // Renders if the user is logged.
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
