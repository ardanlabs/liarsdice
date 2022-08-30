import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import App from './App'
import { BrowserRouter } from 'react-router-dom'
import reportWebVitals from './reportWebVitals'
// Connect dApp
import { Localhost } from '@usedapp/core'
import { getDefaultProvider } from 'ethers'
import { Mainnet, DAppProvider, Config } from '@usedapp/core'
import axios, { AxiosResponse } from 'axios'
import { apiUrl } from './utils/axiosConfig'

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement)
const mainnetConfig: Config = {
  readOnlyChainId: Mainnet.chainId,
  readOnlyUrls: {
    [Mainnet.chainId]: getDefaultProvider('mainnet'),
  },
}

let providerConfig: Config = {}
export const getAppConfig = axios
  .get(`http://${apiUrl}/config`)
  .then((response: AxiosResponse) => {
    const data = response.data
    providerConfig = {
      readOnlyChainId: data.ChainID,
      readOnlyUrls: {
        [data.ChainID]: data.Network,
      },
    }
    return data
  })

const ardansLocalHostConfig: Config = {
  readOnlyChainId: Localhost.chainId,
  readOnlyUrls: {
    [Localhost.chainId]: 'http://127.0.0.1:8545/',
  },
}

const chainsConfig = {
  Ardans: ardansLocalHostConfig,
  EthMainnet: mainnetConfig,
  default: providerConfig,
}

// TODO: ADD dinamic .env support
root.render(
  <React.StrictMode>
    <BrowserRouter>
      <DAppProvider config={chainsConfig['default']}>
        <App />
      </DAppProvider>
    </BrowserRouter>
  </React.StrictMode>,
)

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals()
