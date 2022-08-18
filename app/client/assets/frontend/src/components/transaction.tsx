import React, { useState } from 'react'
import axios from 'axios'
import Button from './button'
// Contract and contract Abi
import { contractAddress } from '../contracts'
import contractAbi from '../abi/Contract.json'
// Contract utils from DApp library
import { useContractFunction, useEthers } from '@usedapp/core'
import { Contract } from '@ethersproject/contracts'
import { utils, BigNumber } from 'ethers'

type getExchangeRateResponse = {
  data: {
    amount: string
    base: 'ETH'
    currency: 'USD'
  }
}

type transactionProps = {
  buttonText: string
  action: 'Deposit' | 'Withdraw'
  updateBalance: Function
}

const Transaction = (props: transactionProps) => {
  async function getExchangeRate() {
    try {
      const { data } = await axios.get<getExchangeRateResponse>(
        'https://api.coinbase.com/v2/prices/ETH-USD/spot',
      )

      return data
    } catch (error) {
      if (axios.isAxiosError(error)) {
        console.error('error message: ', error.message)
        return error.message
      } else {
        console.error('unexpected error: ', error)
        return 'An unexpected error occurred'
      }
    }
  }
  const { buttonText, action, updateBalance } = props
  // Sets local state
  const [transactionAmount, setTransactionAmount] = useState(0)
  // Creates the interface with the contract aby
  const contractInterface = new utils.Interface(contractAbi)
  // Creates a new contract object
  const contract = new Contract(contractAddress, contractInterface)
  // Extracts the functions from the contract
  const { send } = useContractFunction(contract, action, {
    gasLimitBufferPercentage: 100,
  })
  const { account } = useEthers()
  const handleAmountChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setTransactionAmount(parseFloat(event.target.value))
  }
  const sendTransaction = () => {
    getExchangeRate().then((response) => {
      let responseEth = response as getExchangeRateResponse
      if (responseEth.data.amount) {
        const priceInWei =
          transactionAmount /
          parseInt(responseEth.data.amount) /
          0.000000000000000001
        send({ value: BigNumber.from(`${Math.round(priceInWei)}`) }).then(
          (response) => {
            updateBalance()
          },
        )
      } else {
        console.error(response)
      }
    })
  }

  return !account ? (
    <p>Please connect your wallet account to proceed.</p>
  ) : (
    <div
      style={{
        height: '100%',
        color: 'black',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'flex-end',
      }}
    >
      <div
        style={{
          display: 'flex',
          alignItems: 'center',
        }}
      >
        <span className="mr-3" style={{ color: 'var(--modals)' }}>
          U$D
        </span>
        <input
          className="form-control"
          type="number"
          onChange={handleAmountChange}
        />
        <Button
          {...{
            id: 'transaction',
            clickHandler: sendTransaction,
            classes: 'd-flex align-items-center pa-4 justify-content-center',
          }}
        >
          <span style={{ color: 'var(--modals)', textAlign: 'center' }}>
            {buttonText || action}
          </span>
        </Button>
      </div>
    </div>
  )
}

export default Transaction
