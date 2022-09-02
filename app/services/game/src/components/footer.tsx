import React, { BaseSyntheticEvent, useContext, useState } from 'react'
import Button from './button'
import { shortenIfAddress } from '../utils/address'
import { GameContext } from '../contexts/gameContext'
import axios, { AxiosError, AxiosResponse } from 'axios'
import { game } from '../types/index.d'
import { axiosConfig } from '../utils/axiosConfig'
import { toast } from 'react-toastify'
import SignOut from './signout'
import getActivePlayersLength from '../utils/getActivePlayers'
import useEthersConnection from './hooks/useEthersConnection'

function Footer() {
  const { account } = useEthersConnection()
  const apiUrl = process.env.REACT_APP_GO_HOST
    ? process.env.REACT_APP_GO_HOST
    : 'localhost:3000/v1/game'

  const { game, setGame } = useContext(GameContext)
  const [number, setNumber] = useState(1)
  const [suite, setSuite] = useState(1)

  const setNewGame = (data: game) => {
    let newGame = data
    newGame = newGame.cups ? newGame : { ...newGame, cups: [] }
    newGame = newGame.player_order ? newGame : { ...newGame, player_order: [] }
    setGame(newGame)
  }

  const sendBet = () => {
    axios
      .get(`http://${apiUrl}/bet/${number}/${suite}`, axiosConfig)
      .then(function (response: AxiosResponse) {
        if (response.data) {
          setNewGame(response.data)
        }
      })
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }
  const callLiar = () => {
    axios
      .get(`http://${apiUrl}/liar`, axiosConfig)
      .then(function (response: AxiosResponse) {
        if (getActivePlayersLength(game.player_order) === 1) {
          toast.info(
            `Game finished! Winner is ${shortenIfAddress(
              response.data.cups[0].account,
            )}`,
          )
        }
      })
      .catch(function (error: AxiosError) {
        console.error(error)
      })
  }

  const handleForm = (event: BaseSyntheticEvent) => {
    if (event.target.id === 'bet__number') {
      setNumber(event.target.value)
    }
    if (event.target.id === 'bet__suite') {
      setSuite(event.target.value)
    }
  }

  return account ? (
    <footer
      id="footer"
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
      <div
        style={{
          width: 'fit-content',
        }}
      >
        <SignOut disabled={!account} />
      </div>
      {(game.player_order as string[])[game.current_cup] === account &&
      game.status === 'playing' ? (
        <div
          style={{
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            width: '100%',
          }}
        >
          <strong
            style={{
              fontSize: '24px',
              color: 'var(--secondary-color)',
            }}
          >
            My Bet:{' '}
          </strong>
          <div className="form-group mx-2 my-2">
            <input
              type="number"
              min={1}
              className="form-control"
              id="bet__number"
              onChange={handleForm}
              style={{
                backgroundColor: 'transparent',
                borderColor: '1px solid var(--secondary-color)',
              }}
            />
          </div>
          <div className="form-group mx-2 my-2">
            <select
              defaultValue="1"
              className="form-control"
              id="bet__suite"
              onChange={handleForm}
              style={{
                backgroundColor: 'transparent',
                borderColor: '1px solid var(--secondary-color)',
              }}
            >
              <option value="1">1</option>
              <option value="2">2</option>
              <option value="3">3</option>
              <option value="4">4</option>
              <option value="5">5</option>
              <option value="6">6</option>
            </select>
          </div>
          <Button
            {...{
              style: {
                margin: '0 8px',
                width: 'fit-content',
                backgroundColor: 'var(--primary-color)',
                color: 'white',
                fontWeight: '600',
              },
              clickHandler: sendBet,
              classes: 'd-flex align-items-center pa-4',
            }}
          >
            <>Make Bet</>
          </Button>
          <Button
            {...{
              style: {
                width: 'fit-content',
                margin: '0 8px',
                backgroundColor: 'var(--primary-color)',
                color: 'white',
                fontWeight: '600',
              },
              clickHandler: callLiar,
              classes: 'd-flex align-items-center pa-4',
            }}
          >
            <>Call Liar</>
          </Button>
        </div>
      ) : (
        <div
          style={{
            display: 'flex',
            flexGrow: '1',
          }}
        ></div>
      )}
    </footer>
  ) : null
}

export default Footer
