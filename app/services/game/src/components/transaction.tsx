import React, { useState } from 'react'
import axios, { AxiosError } from 'axios'
import Button from './button'
// Contract Abi
import contractAbi from '../abi/Contract.json'
import { ethers, utils } from 'ethers'
// Another utils
import { toast } from 'react-toastify'
import { apiUrl } from '../utils/axiosConfig'
import { useLocation } from 'react-router-dom'
import { appConfig } from '../types/index.d'
import useEthersConnection from './hooks/useEthersConnection'

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
  const { account, signer, provider } = useEthersConnection()

  const send = async (value?: string) => {
    // Creates the interface with the contract aby
    const contractInterface = new utils.Interface(contractAbi)

    const contractAddress = (state as appConfig).contract_id

    // Creates a new contract object and connects it to the signer
    const contract = new ethers.Contract(
      contractAddress,
      contractInterface,
      signer,
    )
    const tx = {
      gasPrice: provider.getGasPrice(),
      gasLimit: '10000000',
    }
    if (action === 'Withdraw') {
      return contract.Withdraw(tx)
    }

    return contract.Deposit({ ...tx, value })
  }

  const [inputValue, setInputValue] = useState('')
  const handleAmountChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(event.target.value)
    setTransactionAmount(parseFloat(event.target.value))
  }
  const sendTransaction = () => {
    axios
      .get(`http://${apiUrl}/usd2wei/${transactionAmount}`)
      .then(async (response: usd2weiResponse) => {
        await send(`${response.data.wei}`)
          .then(async (txResponse: any) => {
            const transaction = await txResponse.wait(0)

            if (transaction.status !== 1) {
              toast.error(`${action} failed`)
            } else {
              updateBalance()
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
