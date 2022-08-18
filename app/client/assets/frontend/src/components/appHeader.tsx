import { useEthers } from '@usedapp/core'
import axios, { AxiosResponse } from 'axios'
import React, { FC, useCallback, useEffect, useState } from 'react'
import Transaction from './transaction'

interface AppHeaderProps {
  show?: boolean
}

const AppHeader: FC<AppHeaderProps> = (AppHeaderProps) => {
  const { show } = AppHeaderProps
  const { account } = useEthers()
  const [balance, setBalance] = useState(0)

  const updateBalance = useCallback(() => {
    if (account)
      axios
        .get(`http://${process.env.REACT_APP_GO_HOST}/balance/${account}`)
        .then((response: AxiosResponse) => {
          setBalance(response.data.balance)
        })
  }, [account])

  useEffect(() => {
    updateBalance()
  }, [updateBalance])

  if (!show) {
    return null
  }
  return (
    <header className="App-header">
      <h1>Ardan's Liar's Dice</h1>
      {account ? (
        <>
          <div
            style={{
              position: 'absolute',
              right: '18px',
            }}
          >
            <div>Current Balance: {balance}</div>
            <Transaction
              {...{ buttonText: 'Deposit', action: 'Deposit', updateBalance }}
            />
          </div>
        </>
      ) : (
        ''
      )}
    </header>
  )
}
export default AppHeader
