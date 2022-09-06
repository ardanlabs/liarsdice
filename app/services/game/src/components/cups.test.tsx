import { render } from '@testing-library/react'
import Cups from './cups'
import { game, bet, die } from '../types/index.d'
import { GameContext } from '../contexts/gameContext'

interface providerValueInterface {
  game: game
  setGame: React.Dispatch<React.SetStateAction<game>>
}

const renderComponent = (providerValue: providerValueInterface) => {
  return render(
    <GameContext.Provider value={providerValue}>
      <Cups />
    </GameContext.Provider>,
  )
}

test('renders all cups', () => {
  const providerValueWithCups = {
    game: {
      status: 'playing',
      lastOut: '',
      lastWin: '',
      currentPlayer: '',
      currentCup: 0,
      round: 1,
      cups: [
        {
          account: '0x8E113078ADF6888B7ba84967F299F29AeCe24c55',
          outs: 0,
          bet: {
            account: '0x8E113078ADF6888B7ba84967F299F29AeCe24c55',
            number: 3,
            suite: 5 as die,
          },
          dice: [1, 2, 3, 4, 5],
        },
        {
          account: '0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7',
          outs: 0,
          bet: {
            account: '0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7',
            number: 5,
            suite: 6 as die,
          },
          dice: [1, 2, 3, 4, 5],
        },
      ],
      playerOrder: [],
      bets: [] as bet[],
      anteUsd: 0,
    } as game,
    setGame: (() => {}) as React.Dispatch<React.SetStateAction<game>>,
  }

  const { queryAllByTestId } = renderComponent(providerValueWithCups)
  const playerCup = queryAllByTestId('player__ui')
  expect(playerCup).not.toBeNull()
})

test('renders all cups', () => {
  const providerValueWithCups = {
    game: {
      status: 'gameover',
      lastOut: '',
      lastWin: '',
      currentPlayer: '',
      currentCup: 0,
      round: 1,
      cups: [],
      playerOrder: [],
      bets: [] as bet[],
      anteUsd: 0,
    } as game,
    setGame: (() => {}) as React.Dispatch<React.SetStateAction<game>>,
  }

  const { queryAllByTestId } = renderComponent(providerValueWithCups)
  const playerCup = queryAllByTestId('player__ui')
  expect(playerCup.length).toBe(0)
})
