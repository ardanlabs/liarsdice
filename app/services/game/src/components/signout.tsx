import React from 'react'
import Button from './button'
import LogOutIcon from './icons/logout'
import useEthersConnection from './hooks/useEthersConnection'
import { useNavigate } from 'react-router-dom'
import { SignOutProps } from '../types/props.d'

const SignOut = (props: SignOutProps) => {
  const { disabled } = props
  const { account, setAccount } = useEthersConnection()
  const navigate = useNavigate()
  // ===========================================================================

  // handleDisconnectAccount disconnects the user and deletes the token.
  function handleDisconnectAccount() {
    setAccount(undefined)
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
