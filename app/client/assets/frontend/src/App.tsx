import React, { Component } from 'react'
import './App.css'
import Login from './components/login'
import MainRoom from "./components/MainRoom";
import { user } from '../types/index.d'

export type State = {
  users: Set<user>,
  isLogged: boolean,
}

class App extends Component<{}, State> {
  constructor(props: any) {
    super(props)
    this.state = {
      users: new Set([{
        address: ''
      }]),
      isLogged: false,
    }
    this.connectMetaMask = this.connectMetaMask.bind(this)
    this.setUser = this.setUser.bind(this)
    this.checkIfConnected = this.checkIfConnected.bind(this)
  }
  componentDidMount(): void {
    this.checkIfConnected()
  }
  checkIfConnected() {
    window.ethereum.request(
      { method: 'eth_accounts' }
    ).then(
      (accounts : string[]) => {
        if(accounts.length) {
          this.setState({
            isLogged: true,
          })
        } else {
          this.setState({
            isLogged: false,
          })
        }
      }
    ).catch((error: any) => {
      console.log(error)
    })
  }
  setUser(accounts: string[], isLogged: boolean) {
    const modifiedUsers = this.state.users
    modifiedUsers.add({address: accounts[0]})
    this.setState({
      users: modifiedUsers,
      isLogged: isLogged,
    })
  }
  connectMetaMask() {
    if((window as any).ethereum) {      
      //check if Metamask wallet is installed
      (window as any).ethereum
        .request({
            method: "eth_requestAccounts",
        })
        .then((accounts : string[]) => {
          this.setUser(accounts, true)
        })
        .catch((error: any) => {
          console.log(error)
          alert(`Something went wrong: ${error.message}`);
        });
    }
  }
  render() {
    const { connectMetaMask } = this
    const { isLogged } = this.state
    return (
      <div className="App">
        <header className="App-header">Ardan's Liar's Dice</header>
        <div className="container-fluid d-flex align-items-center justify-content-center" style={{height: '100vh'}}>
          <MainRoom show={isLogged} />
          <Login show={!isLogged} clickHandler={ connectMetaMask }/>
        </div>
      </div>
    )
  }
}

export default App
