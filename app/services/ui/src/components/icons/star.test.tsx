import React from 'react'
import { render } from '@testing-library/react'
import Star from './star'

test('renders star icon', () => {
  // Tests what happens if you send a number to the Die component.
  // Checks if that number is rendered
  const { container } = render(
    <Star width={'29px'} height={'26px'} fill="#0000" />,
  )

  const svgElement = container.querySelector('#star_svg') as SVGElement
  const pathElement = svgElement.querySelector('path')
  expect(svgElement).toBeInTheDocument()
  expect(pathElement).toBeInTheDocument()
  expect(svgElement.style.height).toBe('26px')
  expect(svgElement.style.width).toBe('29px')
  expect(pathElement?.getAttribute('fill')).toBe('#0000')
})
