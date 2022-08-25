import React from 'react'
import { render, screen } from '@testing-library/react'
import AppHeader from './appHeader'

test('renders appHeader h1', () => {
  // Tests what happens if you send a number to the Die component.
  // Checks if that number is rendered
  render(<AppHeader show={true} />)

  const headerElement = screen.getByTestId('app-header')
  const h1Element = headerElement.querySelector('h1')
  expect(h1Element).toHaveTextContent(`Ardan's Liar's Dice`)
})
