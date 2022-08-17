import { useEtherBalance, useEthers } from '@usedapp/core'
import axios, { AxiosResponse } from 'axios'
import React, { FC, useEffect, useState } from 'react'
import Transaction from './transaction'

interface AppHeaderProps {
  show?: boolean
}

const AppHeader: FC<AppHeaderProps> = (AppHeaderProps) => {
  const { show } = AppHeaderProps
  const { account } = useEthers()
  const [balance, setBalance] = useState(0)

  const updateBalance = () => {
    axios
      .get(`http://${process.env.REACT_APP_GO_HOST}/balance/${account}`)
      .then((response: AxiosResponse) => {
        setBalance(response.data.balance)
      })
  }

  useEffect(() => {
    updateBalance()
  }, [])

  if (!show) {
    return null
  }
  return (
    <header className="App-header">
      <span>Ardan's Liar's Dice</span>
      <div>Current Balance: {balance}</div>
      <Transaction
        {...{ buttonText: 'Deposit', action: 'Deposit', updateBalance }}
      />
    </header>
  )
}
export default AppHeader
