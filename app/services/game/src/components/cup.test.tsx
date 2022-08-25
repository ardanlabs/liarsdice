import { render } from '@testing-library/react'
import Cup from './cup'
import { die, CupProps, dice } from '../types/index.d'

let props = {
  player: {
    account: '0x8E113078ADF6888B7ba84967F299F29AeCe24c55',
    outs: 0,
    claim: {
      account: '0x8E113078ADF6888B7ba84967F299F29AeCe24c55',
      number: 3,
      suite: 5 as die,
    },
  },
  playerDice: [4, 3, 5, 6, 2] as dice,
}

const renderComponent = (props: CupProps) => {
  return render(<Cup {...props} />)
}

test('renders an active player cup', () => {
  const { queryByTestId } = renderComponent(props)

  const playerCup = queryByTestId('player__cup')

  expect(playerCup).not.toBeNull()
  expect(playerCup?.classList).toContain('active')
})

test('renders an striked out player cup', () => {
  const { queryByTestId } = renderComponent({
    ...props,
    player: { ...props.player, outs: 3 },
  })

  const playerCup = queryByTestId('player__cup')

  expect(playerCup).not.toBeNull()
  expect(playerCup?.classList).not.toContain('active')
})
