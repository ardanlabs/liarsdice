import React, { useContext } from 'react'
import Button from './button'
import LogOutIcon from './icons/logout'
import { useEthers } from '@usedapp/core'
import { GameContext } from '../gameContext'
import axios, { AxiosError, AxiosResponse } from 'axios'
import { game } from '../types/index.d'

function Footer() {
  const { account } = useEthers()
  const { deactivate } = useEthers()
  function handleDisconnectWallet() {
    deactivate()
  }
  const { game, setGame } = useContext(GameContext)

  const setNewGame = (data: game) => {
    let newGame = data
    newGame = newGame.players ? newGame : { ...newGame, players: [] }
    newGame = newGame.player_order
      ? newGame
      : { ...newGame, player_order: [] }
    setGame(newGame)
  }

  const sendClaim = () => {
    axios
      .post(`http://${process.env.REACT_APP_GO_HOST}/claim/${account}`)
      // Add the data that is sent to the backend this will be done with the handleInput method
      // This is the only step left to make the game functional, the next step after that will be adding the timer.
      .then(function (response: AxiosResponse) {
        console.log('New round!')
        if (response.data) {
          setNewGame(response.data)
        }
      })
      .catch(function (error: AxiosError) {
        console.log(error)
      })
  }
  const callLiar = () => {
    axios
      .post(`http://${process.env.REACT_APP_GO_HOST}/callliar/${account}`)  
      .then(function (response: AxiosResponse) {
        console.log('New round!')
      })
      .catch(function (error: AxiosError) {
        console.log(error)
      })
  }

  return account ? (
    <footer
      style={{
        backgroundColor: 'var(--modals)',
        position: 'fixed',
        bottom: '0',
        height: '70px',
        width: '100%',
        display: 'flex',
        justifyContent: 'start',
        alignItems: 'center',
      }}
    >
      <Button
        {...{
          id: 'metamask__wrapper',
          clickHandler: handleDisconnectWallet,
          classes: 'd-flex align-items-center pa-4',
        }}
      >
        <LogOutIcon />
      </Button>
      {game.current_player === account ? (
        <div
          style={{
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
          }}
        >
          <strong>My Claim: </strong>
          <div className="form-group mx-2 mt-2">
            <input
              type="number"
              min={1}
              className="form-control"
              id="claim__number"
              placeholder="1"
            />
          </div>
          <div className="form-group mx-2 mt-2">
            <select className="form-control" id="claim__suite">
              <option selected value="1">
                1
              </option>
              <option value="2">2</option>
              <option value="3">3</option>
              <option value="4">4</option>
              <option value="5">5</option>
              <option value="6">6</option>
            </select>
          </div>
          <Button
            {...{
              id: 'metamask__wrapper',
              clickHandler: sendClaim,
              classes: 'd-flex align-items-center pa-4',
            }}
          >
            <>
              Make Claim
            </>
          </Button>
          <Button
            {...{
              id: 'metamask__wrapper',
              clickHandler: callLiar,
              classes: 'd-flex align-items-center pa-4',
            }}
          >
            <>
              Call Liar
            </>
          </Button>
        </div>
      ) : (
        ''
      )}
    </footer>
  ) : null
}

export default Footer
