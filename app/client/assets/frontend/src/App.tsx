import React, { useState, useMemo } from 'react'
import './App.css'
import Login from './components/login'
import Footer from './components/footer'
import { GameContext } from './gameContext'
import { game } from './types/index.d'

export function App(){
  const [ game, setGame ] = useState({
    status: 'open',
    round: 0,
    current_player: '',
    player_order: [],
    players: [],
  } as game)

  const providerGame = useMemo(() => ({game, setGame}), [game, setGame])

  return (
    <div className="App" style={{scrollSnapType: 'y mandatory'}}>
      <GameContext.Provider
        value={providerGame}
      >
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
