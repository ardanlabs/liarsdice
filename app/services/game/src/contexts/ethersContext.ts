import { Signer } from 'ethers'
import React from 'react'

export interface ethersConnectionInterface {
  network: object
  signer: Signer
  account: string | null
}

export interface ethersContextInterface {
  ethersConnection: ethersConnectionInterface
  setEthersConnection: React.Dispatch<
    React.SetStateAction<ethersConnectionInterface>
  >
}

export const EthersContext = React.createContext({
  ethersConnection: {} as ethersConnectionInterface,
  setEthersConnection: (() => {}) as React.Dispatch<
    React.SetStateAction<ethersConnectionInterface>
  >,
})
