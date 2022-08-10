import React, { FC } from 'react'

interface CounterProps {}

const Counter: FC<CounterProps> = (CounterProps) => {
  return (
    <span
      style={{
        fontSize: '32px',
        fontWeight: '500',
        color: '#FFFF',
      }}
    > 00:30 </span>
  )
}

export default Counter
