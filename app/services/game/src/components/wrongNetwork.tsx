import React, { useState } from 'react'
import { getAppConfig } from '..'
import { appConfig } from '../types/index.d'

// Component that shows when your network doesn't match the backend config.
const WrongNetwork = () => {
  // We set a state to trigger a rerender when the configuration is loaded.
  const [appConfig, setAppConfig] = useState<appConfig>({} as appConfig)

  // Gets the backend config and sets it to the state.
  getAppConfig.then((response) => {
    setAppConfig(response)
  })

  // Render
  return (
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
              {`${appConfig.network} (ChainID ${appConfig.chain_id})`}
            </strong>
          ) : (
            <strong> Loading configuration...</strong>
          )}
        </h2>
      </div>
    </div>
  )
}

export default WrongNetwork
