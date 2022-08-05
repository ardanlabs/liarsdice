import React, { Component } from 'react'
import { user } from '../../types/index.d'
import Button from './button'
import MetamaskLogo from './icons/metamask'

export type State = {
  user: user
}

interface userProps {
  clickHandler: Function
}

class Login extends Component<userProps, State> {
  constructor(props: userProps) {
    super(props)
    this.state = {
      user: {
        address: '',
      },
    }
  }
  render() {
    const { clickHandler } = this.props
    return (
      <div
        id="login__wrapper"
        className="d-flex align-items-start justify-content-center flex-column"
      >
        <h2>
          <strong> Connect your wallet </strong>
        </h2>
        Or you can also select a provider to create one.
        <div id="wallets__wrapper" className="mt-4">
          <Button {...{ id: 'metamask__wrapper', clickHandler: clickHandler, classes: 'd-flex align-items-center pa-4'}}>
            <MetamaskLogo {...{ width: '50px', height: '50px' }} />
            <span className="ml-4"> Metamask </span>
          </Button>
        </div>
      </div>
    )
  }
}
export default Login
