import React, { useCallback, useEffect, useState } from 'react'

import axios, { AxiosError, AxiosResponse } from 'axios'
import { apiUrl, axiosConfig } from '../utils/axiosConfig'
import Transaction from './transaction'
import useEthersConnection from './hooks/useEthersConnection'

function PlayerBalance() {
  // Extracts account from useEthersConnection.
  const { account } = useEthersConnection()

  // Sets balance state.
  const [balance, setBalance] = useState(0)

  const updateBalanceUCFn = (balance: number = -1) => {
    if (balance !== -2 && account) {
      axios
        .get(`http://${apiUrl}/balance`, axiosConfig)
        .then((balanceResponse: AxiosResponse) => {
          setBalance(parseFloat(balanceResponse.data.balance))
        })
        .catch((error: AxiosError) => {
          console.error(error)
        })
    }
  }

  // updateBalance creates a callback that fetches the player's balance.
  const updateBalance = useCallback(updateBalanceUCFn, [account])

  // toggle opens and closes the modal.
  function toggle() {
    const dropdown = document.querySelector('.dropdown-menu') as HTMLElement
    dropdown.style.display =
      dropdown.style.display === 'none' ? 'block' : 'none'
  }

  // We use an effect to trigger balance updates.
  useEffect(() => {
    updateBalance()
  }, [updateBalance])

  // renders the final markup if there's an account.
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
