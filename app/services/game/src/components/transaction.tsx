import React, { useState } from 'react'
import axios, { AxiosError } from 'axios'
import Button from './button'
// Contract Abi
import contractAbi from '../abi/Contract.json'
// Another utils
import { ethers, utils } from 'ethers'
import { toast } from 'react-toastify'
import { apiUrl } from '../utils/axiosConfig'
import { useLocation } from 'react-router-dom'
import { appConfig } from '../types/index.d'
import useEthersConnection from './hooks/useEthersConnection'
import { transactionProps } from '../types/props.d'
import { usd2weiResponse } from '../types/responses.d'

// Transaction component.
function Transaction(props: transactionProps) {
  // Extracts router state from useLocation hook.
  const { state } = useLocation()

  // Extracts props.
  const { buttonText, action, updateBalance } = props

  // Sets local transactionAmount state
  const [transactionAmount, setTransactionAmount] = useState(0)

  // Sets local inputValue state.
  const [inputValue, setInputValue] = useState('')

  // Extracts properties from useEthersConnection hook.
  const { account, signer, provider } = useEthersConnection()

  // ===========================================================================

  // Send handles the web3 transaction send.
  // Takes an optional amount of wei to send.
  async function send(value?: string) {
    // Creates the interface with the contract aby
    const contractInterface = new utils.Interface(contractAbi)

    // Gets the contract address from the state.
    const contractAddress = (state as appConfig).contractId

    // Creates a new contract object and connects it to the signer
    const contract = new ethers.Contract(
      contractAddress,
      contractInterface,
      signer,
    )

    // Sets the transaction to send.
    const tx = {
      gasPrice: provider.getGasPrice(),
      gasLimit: '10000000',
    }

    // Withdraws all the money if it's a withdraw transaction.
    if (action === 'Withdraw') {
      return contract.Withdraw(tx)
    }

    // Adds the value to the deposit if it's a deposit transaction.
    return contract.Deposit({ ...tx, value })
  }

  // If action === 'Deposit'
  // sendTransaction calls to the backend to convert usd amount to wei and then
  // calls send() with that value.
  // If actions === 'Withdraw'
  // sendTransaction only calls to send()
  function sendTransaction() {
    if (action === 'Withdraw') {
      send()
      return
    }
    axios
      .get(`http://${apiUrl}/usd2wei/${transactionAmount}`)
      .then(async (response: usd2weiResponse) => {
        const sendTransactionSendFn = async (txResponse: any) => {
          console.log(txResponse)

          const transaction = await txResponse.wait(0)

          if (transaction.status !== 1) {
            toast.error(`${action} failed`)
            return
          }

          updateBalance()
          setInputValue('')
          toast.info(`${action} successful`)
        }

        await send(`${response.data.wei}`)
          .then(sendTransactionSendFn)
          .catch((error: AxiosError) => {
            console.error(error)
          })
        return response
      })
  }

  // ===========================================================================

  // handleAmountChange keeps track of the input value change in the local state.
  function handleAmountChange(event: React.ChangeEvent<HTMLInputElement>) {
    setInputValue(event.target.value)
    setTransactionAmount(parseFloat(event.target.value))
  }

  // Renders this markup if there's an account connected.
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
