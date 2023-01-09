/* ************useGameHook************

  This hook provides all functions needed to run the game.
  It only exports the methods needed to run the game from another component.
  Apart from the gameflow methods (line 36) we have the gamepot set in here.
  Inside ~/src/components/mainRoom.tsx you can see an implementation of how all of this is working.
  Might be helpfull to see ~/src/hooks/useWebSocket.ts to understand the events that run the game notifications/updating system.
  
  **************************************** */
import axios, { AxiosError } from 'axios'
import { apiUrl } from '../../utils/axiosConfig'
import { connectResponse } from '../../types/responses.d'
import { getAppConfig } from '../..'
import { useNavigate } from 'react-router-dom'

// Create an axios instance to keep the token updated
const axiosInstance = axios.create({
  headers: {
    authorization: window.sessionStorage.getItem('token') as string,
  },
})

function useGame() {
  const navigate = useNavigate()

  // connectToGameEngine connects to the game engine, and stores the token
  // in the sessionStorage. Takes an object with the following type:
  // { dateTime: string; sig: string }
  function connectToGameEngine(data: {
    address: string
    dateTime: string
    sig: string
  }) {
    const getAppConfigFn = () => {
      navigate('/mainroom')
    }
    if (window.sessionStorage.getItem('token')) {
      window.localStorage.setItem('account', data.address)
      getAppConfig.then(getAppConfigFn)
    }
    const axiosConnectFn = (connectResponse: connectResponse) => {
      window.sessionStorage.setItem(
        'token',
        `Bearer ${connectResponse.data.token}`,
      )

      window.localStorage.setItem('account', data.address)

      getAppConfig.then(getAppConfigFn)
    }

    const axiosConnectErrorFn = (error: AxiosError) => {
      const errorMessage = (error as any).response.data.error.replace(
        / \[.+\]/gm,
        '',
      )

      console.group()
      console.error('Error:', errorMessage)
      console.groupEnd()
    }

    axiosInstance
      .post(`http://${apiUrl}/connect`, { ...data })
      .then(axiosConnectFn)
      .catch(axiosConnectErrorFn)
  }

  return {
    connectToGameEngine,
  }
}

export default useGame
