import React, { useState, useMemo, useEffect } from 'react'
import './App.css'
import Login from './components/login'
import { GameContext, gameContextInterface } from './contexts/gameContext'
import { game } from './types/index.d'
import { ToastContainer } from 'react-toastify'
import 'react-toastify/ReactToastify.min.css'
import { Route, Routes, useNavigate } from 'react-router-dom'
import MainRoom from './components/mainRoom'
import { getAppConfig } from '.'
import {
  ethersConnectionInterface,
  EthersContext,
  ethersContextInterface,
} from './contexts/ethersContext'
import useEthersConnection from './components/hooks/useEthersConnection'
import WrongNetwork from './components/wrongNetwork'
import { Network } from '@ethersproject/networks'

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
    bets: [],
    ante_usd: 0,
  } as game)

  const [ethersConnection, setEthersConnection] =
    useState<ethersConnectionInterface>({} as ethersConnectionInterface)

  const { provider } = useEthersConnection()

  const navigate = useNavigate()

  // Returns a memoized instance of the game state.
  // This is used to set the game provider context.
  // The game provider context is used for creating a global instance of the game, accesible from all files wrap around that context.
  const getProviderGame = useMemo(
    (): gameContextInterface => ({ game, setGame }),
    [game, setGame],
  )

  // Function to handle the clicks outside a dropdown.
  function hideDropdowns(event: React.MouseEvent<HTMLDivElement, MouseEvent>) {
    if (!(event.target as HTMLElement).classList.contains('dropdown-content')) {
      const dropdown = document.querySelector('.dropdown-menu') as HTMLElement
      if (dropdown) {
        dropdown.style.display = 'none'
      }
    }
  }

  // Returns the default context value for the ethers.js support.
  // This context is used for creating global support for ethers.js connection.
  function getEthersContextDefaultValue(): ethersContextInterface {
    return {
      ethersConnection,
      setEthersConnection,
    }
  }

  // Handles changes of network.
  // When a Provider makes its initial connection, it emits a "network"
  // event with a null oldNetwork along with the newNetwork. So, if the
  // oldNetwork exists, it represents a changing network
  function handleNetworkChange(newNetwork: Network, oldNetwork: Network): void {
    getAppConfig.then(async (getConfigResponse) => {
      if (newNetwork.chainId !== getConfigResponse.ChainID) {
        window.sessionStorage.removeItem('token')
        navigate('/wrongNetwork', { state: { ...getConfigResponse } })
      } else {
        navigate('/')
      }
    })
  }

  useEffect(() => {
    provider.on('network', (newNetwork, oldNetwork) => {
      handleNetworkChange(newNetwork, oldNetwork)
    })

    // eslint-disable-next-line
  }, [])
  // Render
  return (
    <div
      className="App"
      style={{ scrollSnapType: 'y mandatory' }}
      onClick={hideDropdowns}
    >
      <>
        <EthersContext.Provider value={getEthersContextDefaultValue()}>
          <ToastContainer />
          <GameContext.Provider value={getProviderGame}>
            <Routes>
              <Route path="/" element={<Login />}></Route>
              <Route path="/mainroom" element={<MainRoom />}></Route>
              <Route path="/wrongNetwork" element={<WrongNetwork />}></Route>
            </Routes>
          </GameContext.Provider>
        </EthersContext.Provider>
      </>
    </div>
  )
}

export default App
