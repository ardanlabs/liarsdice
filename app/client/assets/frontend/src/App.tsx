import React from 'react'
import './App.css'
import Login from './components/login'

export function App(){
  return (
    <div className="App">
      <header className="App-header">Ardan's Liar's Dice</header>
      <div className="container-fluid d-flex align-items-center justify-content-center" style={{height: '100vh'}}>
        <Login />
      </div>
    </div>
  )
}

export default App
