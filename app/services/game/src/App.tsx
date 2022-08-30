import React, { useState, useMemo } from 'react'
import './App.css'
import Login from './components/login'
import { GameContext } from './gameContext'
import { game } from './types/index.d'
import { ToastContainer } from 'react-toastify'
import 'react-toastify/ReactToastify.min.css'
import { Route, Routes } from 'react-router-dom'
import MainRoom from './components/mainRoom'

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
    ante_usd: 0,
  } as game)

  const providerGame = useMemo(() => ({ game, setGame }), [game, setGame])

  const hideDropdowns = (event: React.MouseEvent<HTMLElement>) => {
    if (!(event.target as HTMLElement).classList.contains('dropdown-content')) {
      const dropdown = document.querySelector('.dropdown-menu') as HTMLElement
      if (dropdown) {
        dropdown.style.display = 'none'
      }
    }
  }

  return (
    <div
      className="App"
      style={{ scrollSnapType: 'y mandatory' }}
      onClick={hideDropdowns}
    >
      <ToastContainer />
      <GameContext.Provider value={providerGame}>
        <Routes>
          <Route path="/" element={<Login />}></Route>
          <Route path="/mainroom" element={<MainRoom />}></Route>
        </Routes>
      </GameContext.Provider>
    </div>
  )
}

export default App
