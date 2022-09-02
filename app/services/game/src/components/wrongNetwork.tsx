import React, { useState } from 'react'
import { getAppConfig } from '..'
import { appConfig } from '../types/index.d'

const WrongNetwork = () => {
  const [appConfig, setAppConfig] = useState<appConfig>({} as appConfig)

  getAppConfig.then((response) => {
    setAppConfig(response)
  })
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
              {`${appConfig.Network} (ChainID ${appConfig.ChainID})`}
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
