import { render, screen } from '@testing-library/react'
import CurrentBet from './currentBet'
import { die, bet } from '../types/index.d'
import { shortenIfAddress } from '../utils/address'

let currentBet = {
  account: '0x8E113078ADF6888B7ba84967F299F29AeCe24c55',
  number: 3,
  suite: 5 as die,
}

const renderComponent = (props: bet) => {
  return render(<CurrentBet currentBet={props} />)
}

test(`happy path if there's a bet made`, () => {
  renderComponent(currentBet)

  const currentBetTextContainer = screen.getByTestId(
    'current_bet_text_container',
  )
  expect(currentBetTextContainer).toHaveTextContent(
    `Current bet by Player ${shortenIfAddress(currentBet.account)}`,
  )
})
