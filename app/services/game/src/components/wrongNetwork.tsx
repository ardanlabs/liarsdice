import React, { useState } from 'react'
import { getAppConfig } from '..'
import { appConfig } from '../types/index.d'

// Component that shows when your network doesn't match the backend config.
const WrongNetwork = () => {
  // We set a state to trigger a re-render when the configuration is loaded.
  // React re-renders the UI when you have a state change inside of it.
  const [appConfig, setAppConfig] = useState<appConfig>({} as appConfig)

  const getAppConfigFn = (response: appConfig) => {
    setAppConfig(response)
  }
  // Gets the backend config and sets it to the state.
  getAppConfig.then(getAppConfigFn)

  // ===========================================================================

  // Render
  return appConfig ? (
    <div
      className="container-fluid d-flex align-items-center justify-content-center px-0 flex-column"
      style={{
        display: 'flex',
        alignItems: 'center',
        height: 'calc(100vh - 70px)',
      }}
    >
      <div
        id="notNetwork__wrapper"
        className="d-flex align-items-center justify-content-center flex-column mt-10"
      >
        <h2>
          {appConfig ? (
            <strong>
              Please switch to network{' '}
              {`${appConfig.network} (ChainID ${appConfig.chainId})`}
            </strong>
          ) : (
            <strong> Loading configuration...</strong>
          )}
        </h2>
      </div>
    </div>
  ) : null
}

export default WrongNetwork
