import React from 'react'
import { CounterProps } from '../types/props.d'

// Counter Function
function Counter(CounterProps: CounterProps) {
  // Extracts props
  const { timer, show } = CounterProps

  // Renders if show is true
  return show ? (
    <span
      data-testid="counter-test"
      style={{
        fontSize: '32px',
        fontWeight: '500',
        color: '#FFFF',
      }}
    >
      {timer}
    </span>
  ) : null
}
export default Counter
