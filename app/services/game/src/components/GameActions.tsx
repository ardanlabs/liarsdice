import { BaseSyntheticEvent, useContext, useState } from 'react'
import { GameContext } from '../contexts/gameContext'
import { die } from '../types/index.d'
import Button from './button'
import useEthersConnection from './hooks/useEthersConnection'
import useGame from './hooks/useGame'

//  GameActions component (betting and call liar)
function GameActions() {
  // Extracts account from ethersConnection Hook
  const { account } = useEthersConnection()

  // Extracts game from the gameContext using useContext Hook
  const { game } = useContext(GameContext)

  // Extracts functions that handle betting and calling liars from the useGame hook.
  const { callLiar, sendBet } = useGame()

  // Creates two states to handle the forms.
  // One for the number the other for the suite (die number).
  const [number, setNumber] = useState(1)
  const [suite, setSuite] = useState(1 as die)

  // ===========================================================================

  const isPlayersTurn = game.currentID === account
  const isGamePlaying = game.status === 'playing'

  // A handler to add parameters to the sendBet function
  function betButtonHandler() {
    sendBet(number, suite)
  }

  // handleForm recieves an event from the form input and sets the state.
  function handleForm(event: BaseSyntheticEvent) {
    switch (event.target.id) {
      case 'bet__number':
        setNumber(event.target.value)
        break
      case 'bet__suite':
        setSuite(event.target.value)
        break
    }
  }

  // If it's the player's turn, and the game it's playing, renders the actions.
  return isPlayersTurn && isGamePlaying ? (
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
          clickHandler: betButtonHandler,
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
  )
}

export default GameActions
