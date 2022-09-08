/* ************useWebSocketHook************

  This hook is in charge of providing the websocket connection.
  When you call connect you are connecting directly to the backend websocket endpoint. /events
  On every message send by the backend we will notify all browsers about the changes, aswell as notify them about the actions taken.
  The connect function has a default state for the game (the instance of gameContext) for when the socket closes.

  **************************************** */

import { shortenIfAddress } from '../../utils/address'
import { useContext, useState } from 'react'
import { toast } from 'react-toastify'
import { GameContext } from '../../contexts/gameContext'
import { apiUrl } from '../../utils/axiosConfig'
import useGame from './useGame'

function useWebSocket(restart: () => void) {
  // Creates a state that keeps track of the websocket status.
  let [wsStatus, setWsStatus] = useState('closed')

  // Extracts setGame from useContext hook.
  let { setGame } = useContext(GameContext)

  // Extracts updateStatus and rolldice from useGame hook.
  const { updateStatus, rolldice } = useGame()

  // Connects to the webscoket.
  function connect() {
    const ws = new WebSocket(`ws://${apiUrl}/events`)

    // If the status isn't open or attemptingConnection we trigger the ws bindings.
    if (wsStatus !== 'open' && wsStatus !== 'attemptingConnection') {
      setWsStatus('attemptingConnection')

      // ws.onopen binds an event listener that triggers with the "open" event.
      ws.onopen = () => {
        toast('Connection established')
        updateStatus()
        setWsStatus('open')
      }

      // ws.onmessage binds an event listener that triggers with "message" event.
      ws.onmessage = (evt: MessageEvent) => {
        updateStatus()
        if (evt.data) {
          let message = JSON.parse(evt.data)
          const messageAccount = shortenIfAddress(message.address)
          // We force a switch in order to check for every type of message
          switch (message.type) {
            // Message received when a new game has been created
            case 'newgame':
              toast(`New game created ${messageAccount}!`)
              break
            // Message received when the game starts
            case 'start':
              toast(`Game has been started by ${messageAccount}!`)
              rolldice()
              break
            // Message received when dices are rolled
            case 'rolldice':
              toast(`Rolling dice's`)
              break
            // Message received when a player joins the game
            case 'join':
              toast(`Account ${messageAccount} just joined`)
              break
            // Message received when bet is maded
            case 'bet':
              toast(`${messageAccount} made a bet`)
              restart()
              break
            // Message received when new round starts
            case 'newround':
              toast('Next Round!')
              break
            // Message received when next turn is started
            case 'nextturn':
              toast('Next Turn!')
              restart()
              break
            // Message received when player gets an out
            case 'outs':
              toast(`Player ${messageAccount} timed out and got striked`)
              break
            // Message received when a player gets called a liar
            case 'callliar':
              toast(`${messageAccount} was called a liar and lost!`)
              restart()
              break
          }
        }
        return
      }

      // ws.onclose binds an event listener that triggers with "close" event.
      // If the socket closes we show the user an error and set the game to
      // it's initial state.
      ws.onclose = (evt: CloseEvent) => {
        toast(`Connection is closed. Reconnect will be attempted in 1 second.`)
        setWsStatus('closed')
        setTimeout(function () {
          setGame({
            status: 'nogame',
            lastOut: '',
            lastWin: '',
            currentPlayer: '',
            currentCup: 0,
            round: 1,
            cups: [],
            playerOrder: [],
            bets: [],
            anteUSD: 0,
            currentID: '',
            balances: [],
          })
        }, 1000)
      }

      // ws.onerror binds an event listener that triggers with "error" event.
      ws.onerror = function (err) {
        toast(`Socket encountered error. Closing socket.`)
        console.error('Socket encountered error: ', err, 'Closing socket')
        ws.close()
        setWsStatus('close')
      }
    }
  }
  return { connect, wsStatus, setWsStatus }
}

export default useWebSocket
