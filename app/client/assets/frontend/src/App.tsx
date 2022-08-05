import React, { Component } from 'react'
import './App.css'
import Login from './components/login'
import { user } from '../types/index.d'

export type State = {
  users: Set<user>,
}

class App extends Component<{}, State> {
  constructor(props: any) {
    super(props)
    this.state = {
      users: new Set([{
        address: '',
      }]),
    }
    this.connectMetaMask = this.connectMetaMask.bind(this)
  }
  connectMetaMask() {
    if((window as any).ethereum){
      //check if Metamask wallet is installed
      (window as any).ethereum
        .request({
            method: "eth_requestAccounts",
        })
        .then((accounts : string[]) => {
          const modifiedUsers = this.state.users
          modifiedUsers.add({address: accounts[0]})
          this.setState({
            users: modifiedUsers,
            // Todo: install redux to manage live state
          })
        })
        .catch((error: any) => {
            alert(`Something went wrong: ${error}`);
        });
    }
  }
  render() {
    const { connectMetaMask } = this
    return (
      <div className="App">
        <header className="App-header">Ardan's Liar's Dice</header>
        <div className="container-fluid d-flex align-items-center justify-content-center" style={{height: '100vh'}}>
          <Login clickHandler={ connectMetaMask }/>
        </div>
      </div>
    )
  }
}

export default App
