import React, { useState, useEffect, useRef } from 'react'
import { user, claim, die } from '../types/index.d'
import SideBar from './sidebar'
import GameTable from './gameTable'
import axios, { AxiosError, AxiosResponse } from 'axios'
import { useEthers } from '@usedapp/core'

interface MainRoomProps {}
const MainRoom = (props: MainRoomProps) => {
  const currentClaim: { wallet: string; claim: claim } = {
    wallet: '0x39249126d90671284cd06495d19C04DD0e54d33',
    claim: { number: 1, suite: 4 },
  }
  const [currentPlayerWallet, setCurrentPlayerWallet] = useState('')
  const [activePlayers, setActivePlayers]= useState([] as user[])
  const [currentGameStatus, setCurrentGameStatus] = useState({})
  const { account } = useEthers()
  const [refresh, setRefresh] = useState(0)
  const timerId = setInterval(() => {
    setRefresh(refresh + 1)
  }, 30000);

  useEffect(
    () => {
      axios
        .get('http://localhost:3000/v1/game/status')
        .then(function (response: AxiosResponse) {
          if (Array.isArray(response.data.players)) {
            setCurrentPlayerWallet(response.data.current_player)
            setActivePlayers(response.data.players)
            setCurrentGameStatus(response.data)
          }
        })
        .catch(function (error: AxiosError) {
          console.log(error)
        })
    }, [refresh]
  )
  const joinGame = () => {
    console.log('Joining game...')
    axios
      .post('http://localhost:3000/v1/game/join', {
        wallet: account,
      })
      .then(function (response: AxiosResponse) {
        setRefresh(refresh + 1)
      })
      .catch(function (error: AxiosError) {
        console.log(error)
      })
  }

  
  return (
    <div
      style={{
        width: '100%',
        display: 'flex',
        justifyContent: 'start',
        alignItems: 'center',
        maxWidth: '100vw',
      }}
      id="mainRoom"
    >
      <SideBar
        activePlayers={activePlayers}
        joinGame={joinGame}
        currentGameStatus={currentGameStatus}
      />
      <GameTable
        activePlayers={activePlayers}
        currentPlayerWallet={currentPlayerWallet}
        currentClaim={currentClaim}
      />
    </div>
  )
}

export default MainRoom
