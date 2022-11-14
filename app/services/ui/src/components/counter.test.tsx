import React from 'react'
import { render, screen } from '@testing-library/react'
import Counter from './counter'

test('renders the counter number', () => {
  const props = {
    timer: 5,
    show: true,
  }
  render(<Counter {...props} />)
  const spanElement = screen.getByTestId('counter-test')

  expect(spanElement).toHaveTextContent('5')
})
