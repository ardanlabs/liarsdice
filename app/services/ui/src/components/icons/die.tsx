import React from 'react'
import { die } from '../../types/index.d'
import { Die1, Die2, Die3, Die4, Die5, Die6, DieQuestion } from './dice/index'

interface DieProps {
  width?: string
  height?: string
  dieNumber?: die
  fill?: string
  style?: React.CSSProperties
}

// Die component.
function Die(props: DieProps) {
  // Extracts props.
  const { dieNumber, style, width, height, fill } = props

  const dice: JSX.Element[] = [
    <Die1 />,
    <Die2 />,
    <Die3 />,
    <Die4 />,
    <Die5 />,
    <Die6 />,
  ]

  // If there's a dieNumber returns it's svg, otherwise returns a question die.
  return dieNumber ? (
    <svg
      role="img"
      className="die"
      fill={fill ? fill : 'var(--secondary-color)'}
      style={{
        width: width ? width : '59px',
        height: height ? height : '60px',
        ...style,
      }}
      xmlns="http://www.w3.org/2000/svg"
      viewBox="3 3 18 18"
    >
      {dice[dieNumber - 1]}
    </svg>
  ) : (
    <DieQuestion />
  )
}

export default Die
