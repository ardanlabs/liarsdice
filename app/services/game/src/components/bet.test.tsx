import React from 'react'
import { render } from '@testing-library/react'
import Bet from './bet'
import { die } from '../types/index.d'

test('renders a bet', () => {
  const dieSize = '54px'
  const props = {
    bet: {
      account: '0x8E113078ADF6888B7ba84967F299F29AeCe24c55',
      number: 3,
      suite: 5 as die,
    },
    dieWidth: dieSize,
    dieHeight: dieSize,
    fill: '#fff',
  }
  const { container } = render(<Bet {...props} />)
  const svgElement = container.querySelector('.die') as SVGElement
  const pathElement = svgElement.querySelector('#die_path_5')

  expect(container.textContent).toBe(`${props.bet.number} X `)
  expect(svgElement.getAttribute('fill')).toBe('#fff')
  expect(svgElement).toBeInTheDocument()
  expect(pathElement).toBeInTheDocument()
  expect(svgElement.style.height).toBe(dieSize)
  expect(svgElement.style.width).toBe(dieSize)
})
