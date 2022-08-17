import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import App from './App'
import reportWebVitals from './reportWebVitals'
import { Localhost } from '@usedapp/core'
// Connect dApp
import { getDefaultProvider } from 'ethers'
import { Mainnet, DAppProvider, Config } from '@usedapp/core'

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement)
const mainnetConfig: Config = {
  readOnlyChainId: Mainnet.chainId,
  readOnlyUrls: {
    [Mainnet.chainId]: getDefaultProvider('mainnet'),
  },
}

const ardansLocalHostConfig = {
  readOnlyChainId: Localhost.chainId,
  readOnlyUrls: {
    [Localhost.chainId]: 'http://127.0.0.1:8545/',
  },
}

const chainsConfig = {
  Ardans: ardansLocalHostConfig,
  EthMainnet: mainnetConfig,
}

// TODO: ADD dinamic .env support
root.render(
  <React.StrictMode>
    <DAppProvider config={chainsConfig['Ardans']}>
      <App />
    </DAppProvider>
  </React.StrictMode>,
)

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
