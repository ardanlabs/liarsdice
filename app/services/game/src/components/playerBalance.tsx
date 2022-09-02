import React, { useCallback, useEffect, useState } from 'react'

import axios, { AxiosResponse } from 'axios'
import { axiosConfig } from '../utils/axiosConfig'
import Transaction from './transaction'
import useEthersConnection from './hooks/useEthersConnection'

const PlayerBalance = () => {
  const { account } = useEthersConnection()
  const [balance, setBalance] = useState(0)
  const apiUrl = process.env.REACT_APP_GO_HOST
    ? process.env.REACT_APP_GO_HOST
    : 'localhost:3000/v1/game'

  const updateBalance = useCallback(
    (balance: number = -1) => {
      if (balance !== -2 && account) {
        axios
          .get(`http://${apiUrl}/balance`, axiosConfig)
          .then((balanceResponse: AxiosResponse) => {
            setBalance(parseFloat(balanceResponse.data.balance))
          })
      }
    },
    [account, apiUrl],
  )

  const toggle = () => {
    const dropdown = document.querySelector('.dropdown-menu') as HTMLElement
    if (dropdown.style.display === 'none') {
      dropdown.style.display = 'block'
    } else {
      dropdown.style.display = 'none'
    }
  }

  useEffect(() => {
    updateBalance()
  }, [updateBalance])

  return account ? (
    <div
      className="dropdown dropleft dropdown-content"
      style={{
        position: 'absolute',
        right: '18px',
      }}
    >
      <button
        className="btn btn-secondary dropdown-toggle dropdown-content"
        type="button"
        id="dropdownMenuButton"
        data-toggle="dropdown"
        aria-haspopup="true"
        aria-expanded="false"
        style={{
          backgroundColor: 'transparent',
          border: '1px solid var(--secondary-color)',
        }}
        onClick={toggle}
      >
        Balance
      </button>
      <div
        className="dropdown-menu dropdown-content"
        aria-labelledby="dropdownMenuButton"
        style={{ display: 'none', backgroundColor: 'var(--modals)' }}
      >
        <div
          className="dropdown-item dropdown-content"
          style={{
            display: 'flex',
            justifyContent: 'flex-end',
            flexDirection: 'column',
            color: 'var(--secondary-color)',
          }}
        >
          <div
            style={{ display: 'flex', margin: '10px 0' }}
            className="dropdown-content"
          >
            Current Balance: {balance.toFixed(2)} U$D
            <Transaction
              {...{
                buttonText: 'Withdraw',
                action: 'Withdraw',
                updateBalance,
              }}
            />
          </div>
          <Transaction
            {...{
              buttonText: 'Deposit',
              action: 'Deposit',
              updateBalance,
            }}
          />
        </div>
      </div>
    </div>
  ) : null
}

export default PlayerBalance
