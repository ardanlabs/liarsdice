import React from 'react'
import { utils } from 'ethers'
import { useEthers } from '@usedapp/core'
import axios, { AxiosError } from 'axios'
import Button from './button'
import { toast } from 'react-toastify'
import { capitalize } from '../utils/capitalize'
import { axiosConfig } from '../utils/axiosConfig'

interface JoinProps {
  disabled: boolean
}

const Join = (props: JoinProps) => {
  const { disabled } = props

  const signMessage = () => {
    toast.info('Joining game...')
    axios
      .get('http://localhost:3000/v1/game/join', axiosConfig)
      .then((response) => {
        toast.info('Welcome to the game')
        window.localStorage.setItem('token', `bearer ${response.data.token}`)
      })
      .catch((error: AxiosError) => {
        let errorMessage = (error as any).response.data.error.replace(
          / \[.+\]/gm,
          '',
        )
        toast.error(
          <div style={{ textAlign: 'start' }}>{capitalize(errorMessage)}</div>,
        )
        console.group()
        console.error('Error:', (error as any).response.data.error)
        console.groupEnd()
      })
  }
  return (
    <Button
      disabled={disabled}
      classes="join__buton"
      clickHandler={() => signMessage()}
    >
      <span>Join Game</span>
    </Button>
  )
}

export default Join
