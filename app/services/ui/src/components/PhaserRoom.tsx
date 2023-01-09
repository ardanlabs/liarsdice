import React, { useRef } from 'react'
import { config } from '../utils/config'
import AppHeader from './appHeader'
import usePhaser from './hooks/usePhaser'
import { PhaserTestProps } from '../types/props.d'

function PhaserTest(props: PhaserTestProps) {
  const parentEl = useRef<HTMLDivElement>(null)
  usePhaser(config, parentEl)

  return (
    <div className="container">
      <AppHeader show={false} />
      <div ref={parentEl} className="gameContainer" />
    </div>
  )
}

export default PhaserTest
