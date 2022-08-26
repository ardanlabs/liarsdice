import { render, screen } from '@testing-library/react'
import CurrentClaim from './currentClaim'
import { die, claim } from '../types/index.d'
import { shortenIfAddress } from '@usedapp/core'

let currentClaim = {
  account: '0x8E113078ADF6888B7ba84967F299F29AeCe24c55',
  number: 3,
  suite: 5 as die,
}

const renderComponent = (props: claim) => {
  return render(<CurrentClaim currentClaim={props} />)
}

test(`happy path if there's a claim made`, () => {
  renderComponent(currentClaim)

  const currentClaimTextContainer = screen.getByTestId(
    'current_claim_text_container',
  )
  expect(currentClaimTextContainer).toHaveTextContent(
    `Current claim by Player ${shortenIfAddress(currentClaim.account)}`,
  )
})
