import React, { FC } from 'react'

interface CounterProps {
  timer: number
  show: boolean
}

const Counter: FC<CounterProps> = (CounterProps) => {
  const { timer, show } = CounterProps

  if (show) {
    return (
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
    )
  }

  return null
}
export default Counter
