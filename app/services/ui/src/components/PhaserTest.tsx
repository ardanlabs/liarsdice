import React, { useEffect, useRef } from 'react'
import { useNavigate } from 'react-router-dom'
import { config } from '../utils/config'
import AppHeader from './appHeader'
import usePhaser from './hooks/usePhaser'
export interface PhaserTestProps {}

function PhaserTest(props: PhaserTestProps) {
  // Extracts navigate from useNavigate Hook
  const navigate = useNavigate()

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
