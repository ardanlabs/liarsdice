import { Signer } from 'ethers'
import React from 'react'

export interface ethersContextInterface {
  network: object
  signer: Signer
  account: string | undefined
}

export const EthersContext = React.createContext({
  ethersConnection: {} as ethersContextInterface,
  setEthersConnection: (() => {}) as React.Dispatch<
    React.SetStateAction<ethersContextInterface>
  >,
})
