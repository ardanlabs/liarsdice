import React from 'react'
import { render } from '@testing-library/react'
import Die from './die'

test('renders a die with provided suite', () => {
  const dieSize = '54px'
  // Tests what happens if you send a number to the Die component.
  // Checks if that number is rendered
  const { container } = render(
    <Die width={dieSize} height={dieSize} dieNumber={5} fill="#fff" />,
  )

  const svgElement = container.querySelector('.die') as SVGElement
  const pathElement = svgElement.querySelector('#die_path_5')
  expect(svgElement.getAttribute('fill')).toBe('#fff')
  expect(svgElement).toBeInTheDocument()
  expect(pathElement).toBeInTheDocument()
  expect(svgElement.style.height).toBe(dieSize)
  expect(svgElement.style.width).toBe(dieSize)
})

test('renders a question die when no suite is provider', () => {
  // Tests what happens if you don't send a number to the Die component.
  // Expectes the die to be a question die
  const { container: questionContainer } = render(<Die />)
  const questionDieSvgElement = questionContainer.querySelector(
    '.question_die',
  ) as SVGElement
  const questionDiepathElement = questionDieSvgElement.querySelector('path')
  expect(questionDieSvgElement).toBeInTheDocument()
  expect(questionDiepathElement).toBeInTheDocument()
})
