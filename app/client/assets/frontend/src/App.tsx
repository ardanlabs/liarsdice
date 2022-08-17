import React, { useState, useMemo } from 'react'
import './App.css'
import Login from './components/login'
import Footer from './components/footer'
import { GameContext } from './gameContext'
import { game } from './types/index.d'

export function App() {
  const [game, setGame] = useState({
    status: 'gameover',
    last_out: '',
    last_win: '',
    current_player: '',
    current_cup: 0,
    round: 1,
    cups: [],
    player_order: [],
    claims: [],
  } as game)

  const providerGame = useMemo(() => ({ game, setGame }), [game, setGame])

  return (
    <div className="App" style={{ scrollSnapType: 'y mandatory' }}>
      <GameContext.Provider value={providerGame}>
        <header className="App-header">Ardan's Liar's Dice</header>
        <div className="container-fluid d-flex align-items-center justify-content-center px-0">
          <Login />
        </div>
        <Footer />
      </GameContext.Provider>
    </div>
  )
}

export default App
