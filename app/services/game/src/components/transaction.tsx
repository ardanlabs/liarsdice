import React, { useState } from 'react'
import axios, { AxiosError } from 'axios'
import Button from './button'
// Contract and contract Abi
import contractAbi from '../abi/Contract.json'
// Contract utils from DApp library
import { useContractFunction, useEthers } from '@usedapp/core'
import { Contract } from '@ethersproject/contracts'
import { utils } from 'ethers'
// Another utils
import { toast } from 'react-toastify'
import { apiUrl } from '../utils/axiosConfig'
import { useLocation } from 'react-router-dom'
import { appConfig } from '../types/index.d'

type transactionProps = {
  buttonText: string
  action: 'Deposit' | 'Withdraw'
  updateBalance: Function
}

interface usd2weiResponse {
  data: {
    usd: number
    wei: number
  }
}

const Transaction = (props: transactionProps) => {
  const { state } = useLocation()
  const { buttonText, action, updateBalance } = props
  // Sets local state
  const [transactionAmount, setTransactionAmount] = useState(0)
  // Creates the interface with the contract aby
  const contractInterface = new utils.Interface(contractAbi)
  const contractAddress = (state as appConfig).config.ContractID
  const contract = new Contract(contractAddress, contractInterface)
  // Creates a new contract object
  // Extracts the functions from the contract
  const { send } = useContractFunction(contract, action, {
    gasLimitBufferPercentage: 100,
  })
  const { account } = useEthers()
  const [inputValue, setInputValue] = useState('')
  const handleAmountChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(event.target.value)
    setTransactionAmount(parseFloat(event.target.value))
  }
  const sendTransaction = () => {
    axios
      .get(`http://${apiUrl}/usd2wei/${transactionAmount}`)
      .then((response: usd2weiResponse) => {
        send({ value: `${response.data.wei}` })
          .then((response) => {
            if (response === undefined) {
              toast.error(`${action} failed`)
            } else {
              updateBalance(-1)
              setInputValue('')
              toast.info(`${action} successful`)
            }
          })
          .catch((error: AxiosError) => {
            console.error(error)
          })
        return response
      })
  }

  return !account ? (
    <p>Please connect your wallet account to proceed.</p>
  ) : (
    <div
      className="dropdown-content"
      style={{
        height: '100%',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'flex-end',
        marginBottom: '10px',
      }}
    >
      <div
        style={{
          display: 'flex',
          alignItems: 'center',
        }}
      >
        {action === 'Deposit' ? (
          <>
            <span
              className="mr-3 dropdown-content"
              style={{ color: 'var(--secondary-color)' }}
            >
              U$D
            </span>
            <input
              className="form-control dropdown-content"
              id="transaction-input"
              type="number"
              value={inputValue}
              onChange={handleAmountChange}
            />
          </>
        ) : (
          ''
        )}
        <Button
          {...{
            id: 'transaction',
            clickHandler: sendTransaction,
            classes:
              'd-flex align-items-center pa-4 justify-content-center dropdown-content',
            style: {
              ...{
                border: '1px solid var(--secondary-color)',
                margin: '0 10px',
              },
            },
          }}
        >
          <span
            style={{ color: 'var(--secondary-color)', textAlign: 'center' }}
          >
            {buttonText || action}
          </span>
        </Button>
      </div>
    </div>
  )
}

export default Transaction
