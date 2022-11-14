import React from 'react'
import { render } from '@testing-library/react'
import MetamaskLogo from './metamask'

test('renders metamask logo', () => {
  // Tests what happens if you send a number to the Die component.
  // Checks if that number is rendered
  const { container } = render(<MetamaskLogo width={'50px'} height={'50px'} />)

  const svgElement = container.querySelector('#Layer_1') as SVGElement
  expect(svgElement).toBeInTheDocument()

  const polygons = svgElement.querySelectorAll('[class^="st"]')
  // The svg is composed of 29 polygons
  expect(polygons.length).toBe(29)
})
