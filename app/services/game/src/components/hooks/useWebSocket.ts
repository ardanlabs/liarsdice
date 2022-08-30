/* ************useWebSocketHook************

  This hook is in charge of providing the websocket connection.
  When you call connect you are connecting directly to the backend websocket endpoint. /events
  On every message send by the backend we will notify all browsers about the changes, aswell as notify them about the actions taken.
  The connect function has a default state for the game (the instance of gameContext) for when the socket closes.

  **************************************** */

import { shortenIfAddress } from '@usedapp/core'
import { useContext, useRef } from 'react'
import { toast } from 'react-toastify'
import { GameContext } from '../../gameContext'
import { apiUrl } from '../../utils/axiosConfig'
import useGame from './useGame'

const useWebSocket = (resetTimer: Function) => {
  let wsStatus = useRef('closed')
  let { setGame } = useContext(GameContext)
  const { rolldice, updateStatus } = useGame()

  const connect = () => {
    const ws = new WebSocket(`ws://${apiUrl}/events`)
    if (
      wsStatus.current !== 'open' &&
      wsStatus.current !== 'attemptingConnection'
    ) {
      ws.onopen = () => {
        toast('Connection established')
        updateStatus()
        wsStatus.current = 'open'
      }
      ws.onmessage = (evt: MessageEvent) => {
        updateStatus()

        if (evt.data) {
          let message = JSON.parse(evt.data)
          const account = shortenIfAddress(message.address)
          // We force a switch in order to check for every type of message
          switch (message.type) {
            // Message received when the game starts
            case 'start':
              toast(`Game has been started by ${account}!`)
              // We roll the dices
              rolldice()
              break
            // Message received when dices are rolled
            case 'rolldice':
              toast(`Rolling dice's`)
              break
            // Message received when a player joins the game
            case 'join':
              toast(`Account ${account} just joined`)
              break
            // Message received when bet is maded
            case 'bet':
              toast(`${account} made a bet`)
              // We reset the timer because a new turn has started
              resetTimer()
              break
            // Message received when new round starts
            case 'newround':
              toast('Next Round!')
              rolldice()
              break
            // Message received when next turn is started
            case 'nextturn':
              toast('Next Turn!')
              // We reset the timer because a new turn has started
              resetTimer()
              break
            // Message received when player gets an out
            case 'outs':
              toast(`Player ${account} timed out and got striked`)
              break
            // Message received when a player gets called a liar
            case 'callliar':
              toast(`${account} was called a liar and lost!`)
              // We reset the timer because a new turn has started
              resetTimer()
              break
          }
        }
        return
      }
      ws.onclose = (evt: CloseEvent) => {
        // If the socket closes we show the user an error and set the game to it's initial state.
        toast(`Connection is closed. Reconnect will be attempted in 1 second.`)
        wsStatus.current = 'closed'
        setTimeout(function () {
          setGame({
            status: 'gameover',
            last_out: '',
            last_win: '',
            current_player: '',
            current_cup: 0,
            round: 1,
            cups: [],
            player_order: [],
            bets: [],
            ante_usd: 0,
          })
          connect()
        }, 1000)
      }
      ws.onerror = function (err) {
        toast(`Socket encountered error. Closing socket.`)
        console.error('Socket encountered error: ', err, 'Closing socket')
        ws.close()
        wsStatus.current = 'close'
      }
    }
  }
  return { connect, wsStatus }
}

export default useWebSocket
