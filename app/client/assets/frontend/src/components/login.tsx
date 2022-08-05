import React, { Component } from 'react'
import Button from './button'
import MetamaskLogo from './icons/metamask'

interface loginProps {
  clickHandler: Function
  show: boolean
}

class Login extends Component<loginProps> {
  render() {
    const { clickHandler, show } = this.props
    if (!show) {
      return null
    }
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
