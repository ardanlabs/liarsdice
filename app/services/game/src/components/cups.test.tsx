import { render } from '@testing-library/react'
import Cups from './cups'
import { game, dice, CupsProps, bet, die } from '../types/index.d'
import { GameContext } from '../gameContext'

let props = {
  playerDice: [4, 3, 5, 6, 2] as dice,
}

interface providerValueInterface {
  game: game
  setGame: React.Dispatch<React.SetStateAction<game>>
}

const renderComponent = (
  props: CupsProps,
  providerValue: providerValueInterface,
) => {
  return render(
    <GameContext.Provider value={providerValue}>
      <Cups {...props} />
    </GameContext.Provider>,
  )
}

test('renders all cups', () => {
  const providerValueWithCups = {
    game: {
      status: 'playing',
      last_out: '',
      last_win: '',
      current_player: '',
      current_cup: 0,
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
        },
        {
          account: '0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7',
          outs: 0,
          bet: {
            account: '0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7',
            number: 5,
            suite: 6 as die,
          },
        },
      ],
      player_order: [],
      bets: [] as bet[],
      ante_usd: 0,
    } as game,
    setGame: (() => {}) as React.Dispatch<React.SetStateAction<game>>,
  }

  const { queryAllByTestId } = renderComponent(props, providerValueWithCups)
  const playerCup = queryAllByTestId('player__ui')
  expect(playerCup).not.toBeNull()
})

test('renders all cups', () => {
  const providerValueWithCups = {
    game: {
      status: 'gameover',
      last_out: '',
      last_win: '',
      current_player: '',
      current_cup: 0,
      round: 1,
      cups: [],
      player_order: [],
      bets: [] as bet[],
      ante_usd: 0,
    } as game,
    setGame: (() => {}) as React.Dispatch<React.SetStateAction<game>>,
  }

  const { queryAllByTestId } = renderComponent(props, providerValueWithCups)
  const playerCup = queryAllByTestId('player__ui')
  expect(playerCup.length).toBe(0)
})
