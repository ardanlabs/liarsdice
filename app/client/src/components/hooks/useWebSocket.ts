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

const useWebSocket = () => {
  let wsStatus = useRef('closed')
  let { setGame } = useContext(GameContext)
  /* eslint-disable @typescript-eslint/no-unused-vars */
  const [
    rolldice,
    timer,
    resetTimer,
    gamePot,
    playerDice,
    managePlayerDice,
    updateStatus,
  ] = useGame()

  const connect = () => {
    const ws = new WebSocket(`ws://${apiUrl}/events`)
    if (
      wsStatus.current !== 'open' &&
      wsStatus.current !== 'attemptingConnection'
    ) {
      ws.onopen = () => {
        updateStatus()
        wsStatus.current = 'open'
        toast.info('Connection established')
      }
      ws.onmessage = (evt: MessageEvent) => {
        updateStatus()

        if (evt.data) {
          let message = evt.data
          // We force a switch in order to check for every type of message
          switch (true) {
            // Message received when the game starts
            case message.startsWith('start:'):
              // We cut the gameOwner account from the end of message
              // Message receive is like start:0x0342235k522j2234h41
              const gameStartStart = 'start:'
              const gameOwnerAccount = shortenIfAddress(
                message.substring(gameStartStart.length),
              )
              toast.info(`Game has been started by ${gameOwnerAccount}!`)
              // We roll the dices
              rolldice()
              break
            // Message received when dices are rolled
            case message === 'rolldice':
              toast.info(`Rolling dice's`)
              break
            // Message received when a player joins the game
            case message.startsWith('join:'):
              // We cut the account that joined from the end of message
              // Message receive is like join:0x0342235k522j2234h41
              const joinStart = 'join:'
              const joinAccount = shortenIfAddress(
                message.substring(joinStart.length),
              )
              toast.info(`Account ${joinAccount} just joined`)
              break
            // Message received when claim is maded
            case message === 'claim':
              // We reset the timer because a new turn has started
              window.sessionStorage.removeItem('round_timer')
              resetTimer()
              break
            // Message received when new round starts
            case message === 'newround':
              toast.info('New round!')
              break
            // Message received when next turn is started
            case message === 'nextturn':
              toast.info('Next turn!')
              break
            // Message received when player gets an out
            case message.startsWith('outs:'):
              // We cut the account that joined from the end of message
              // Message receive is like outs:0x0342235k522j2234h41
              const outsStart = 'outs:'
              const strikedAccount = shortenIfAddress(
                message.substring(outsStart.length),
              )
              toast.info(`Player ${strikedAccount} timed out and got striked`)
              break
            // Message received when a player gets called a liar
            case message.startsWith('callliar:'):
              toast.info('A player was called a liar and lost!')
              break
          }
        }
        return
      }
      ws.onclose = (evt: CloseEvent) => {
        // If the socket closes we show the user an error and set the game to it's initial state.
        toast.error(
          'Connection is closed. Reconnect will be attempted in 1 second. ' +
            evt.reason,
        )
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
            claims: [],
            ante_usd: 0,
          })
          connect()
        }, 1000)
      }

      ws.onerror = function (err) {
        console.error('Socket encountered error: ', err, 'Closing socket')
        ws.close()
        wsStatus.current = 'close'
      }
    }
  }
  return [connect, wsStatus] as const
}

export default useWebSocket
