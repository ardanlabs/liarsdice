import React, { useState, useMemo } from 'react'
import './App.css'
import Login from './components/login'
import Footer from './components/footer'
import { GameContext } from './gameContext'
import { game } from './types/index.d'
import AppHeader from './components/appHeader'
import { ToastContainer } from 'react-toastify'
import 'react-toastify/dist/ReactToastify.css'

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
      <GameContext.Provider value={providerGame}>
        <ToastContainer
          position="top-left"
          autoClose={2000}
          hideProgressBar={false}
          newestOnTop={false}
          closeOnClick
          rtl={false}
          draggable
        />
        <AppHeader show={true} />
        <div className="container-fluid d-flex align-items-center justify-content-center px-0">
          <Login />
        </div>
        <Footer />
      </GameContext.Provider>
    </div>
  )
}

export default App
