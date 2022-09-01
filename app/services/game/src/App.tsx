import React, { useState, useMemo, useEffect } from 'react'
import './App.css'
import Login from './components/login'
import { GameContext } from './contexts/gameContext'
import { game } from './types/index.d'
import { ToastContainer } from 'react-toastify'
import 'react-toastify/ReactToastify.min.css'
import { Route, Routes } from 'react-router-dom'
import MainRoom from './components/mainRoom'
import { getAppConfig } from '.'
import { utils } from 'ethers'
import { EthersContext, ethersContextInterface } from './contexts/ethersContext'
import useEthersConnection from './components/hooks/useEthersConnection'

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
    useState<ethersContextInterface>({} as ethersContextInterface)
  const { provider, switchNetwork } = useEthersConnection()

  const providerGame = useMemo(() => ({ game, setGame }), [game, setGame])

  const hideDropdowns = (event: React.MouseEvent<HTMLElement>) => {
    if (!(event.target as HTMLElement).classList.contains('dropdown-content')) {
      const dropdown = document.querySelector('.dropdown-menu') as HTMLElement
      if (dropdown) {
        dropdown.style.display = 'none'
      }
    }
  }

  const ethersContextDefaultValue = {
    ethersConnection,
    setEthersConnection,
  }

  useEffect(() => {
    provider.on('network', (newNetwork, oldNetwork) => {
      // When a Provider makes its initial connection, it emits a "network"
      // event with a null oldNetwork along with the newNetwork. So, if the
      // oldNetwork exists, it represents a changing network
      if (oldNetwork) {
        window.location.reload()
      }
      getAppConfig.then(async (getConfigResponse) => {
        if (newNetwork.chainId !== getConfigResponse.ChainID) {
          try {
            await switchNetwork({
              chainId: utils.hexValue(getConfigResponse.ChainID), // A 0x-prefixed hexadecimal string
            })
          } catch (error) {
            console.log(error, 'error')
          }
        }
      })
    })
    // eslint-disable-next-line
  }, [])
  return (
    <div
      className="App"
      style={{ scrollSnapType: 'y mandatory' }}
      onClick={hideDropdowns}
    >
      <>
        <EthersContext.Provider value={ethersContextDefaultValue}>
          <ToastContainer />
          <GameContext.Provider value={providerGame}>
            <Routes>
              <Route path="/" element={<Login />}></Route>
              <Route path="/mainroom" element={<MainRoom />}></Route>
            </Routes>
          </GameContext.Provider>
        </EthersContext.Provider>
      </>
    </div>
  )
}

export default App
