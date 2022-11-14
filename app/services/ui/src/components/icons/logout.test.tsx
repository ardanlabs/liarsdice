import React from 'react'
import { render } from '@testing-library/react'
import Logout from './logout'

test('renders logout icon', () => {
  // Tests what happens if you send a number to the Die component.
  // Checks if that number is rendered
  const { container } = render(<Logout width={'24px'} height={'24px'} />)

  const svgElement = container.querySelector('#logout_svg') as SVGElement
  const pathElement = svgElement.querySelector('path')
  expect(svgElement).toBeInTheDocument()
  expect(pathElement).toBeInTheDocument()
  expect(svgElement.style.height).toBe('24px')
  expect(svgElement.style.width).toBe('24px')
  expect(pathElement?.getAttribute('fill')).toBe('#00284d')
})
