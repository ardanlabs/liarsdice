import React, { FC } from 'react'
import { die } from '../../types/index.d'
import { Die1, Die2, Die3, Die4, Die5, Die6, DieQuestion } from './dice/index'

interface DieProps {
  width?: string
  height?: string
  dieNumber?: die
  fill: string
}

const Die: FC<DieProps> = (DieProps) => {
  const { dieNumber } = DieProps
  let { width, height, fill } = DieProps
  const dice: JSX.Element[] = [<Die1 />, <Die2 />, <Die3 />, <Die4 />, <Die5 />, <Die6 />]
  width = width ? width : '59px'
  height = height ? height : '60px'
  fill = fill ? fill : 'var(--secondary-color)'
  if (dieNumber) {
    return (
      <svg className="die" fill={fill} style={{ width: width, height: height }} xmlns="http://www.w3.org/2000/svg" viewBox='3 3 18 18'>
        {dice[dieNumber - 1]}
      </svg>
    )
  }
  return (
    <DieQuestion />
  )
}

export default Die
