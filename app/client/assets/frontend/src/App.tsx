import React from 'react'
import './App.css'
import Login from './components/login'
import Footer from './components/footer'

export function App(){
  return (
    <div className="App" style={{scrollSnapType: 'y mandatory'}}>
      <header className="App-header">Ardan's Liar's Dice</header>
      <div className="container-fluid d-flex align-items-center justify-content-center px-0">
        <Login />
      </div>
      <Footer />
    </div>
  )
}

export default App
