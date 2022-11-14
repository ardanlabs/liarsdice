import React, { useState, useMemo, useEffect } from 'react'
import './App.css'
import Login from './components/login'
import { GameContext, gameContextInterface } from './contexts/gameContext'
import { appConfig, game } from './types/index.d'
import { ToastContainer } from 'react-toastify'
import 'react-toastify/ReactToastify.min.css'
import { Route, Routes, useNavigate } from 'react-router-dom'
import MainRoom from './components/mainRoom'
import { getAppConfig } from '.'
import { utils } from 'ethers'
import {
  ethersConnectionInterface,
  EthersContext,
  ethersContextInterface,
} from './contexts/ethersContext'
import useEthersConnection from './components/hooks/useEthersConnection'
import WrongNetwork from './components/wrongNetwork'
import { Network } from '@ethersproject/networks'
import { WalletProvider } from '@viaprotocol/web3-wallets'

// =============================================================================

// App component. First React component after Index.ts
function App() {
  const [game, setGame] = useState({
    status: 'gameover',
    lastOut: '',
    lastWin: '',
    currentPlayer: '',
    currentCup: 0,
    round: 1,
    cups: [],
    playerOrder: [],
    bets: [],
    anteUSD: 0,
    currentID: '',
    balances: [],
  } as game)

  // Extracts provider from useEthersConnection hook
  const { provider, switchNetwork } = useEthersConnection()

  // Extracts router navigation functionality
  const navigate = useNavigate()

  // Sets state for the ethersConnection context
  const [ethersConnection, setEthersConnection] =
    useState<ethersConnectionInterface>({} as ethersConnectionInterface)

  // getProviderGame returns a memoized instance of the game state. This is used
  // to set the game provider context. The game provider context is used for
  // creating a global instance of the game, accesible from all files wrap around
  // that context.
  const memoFn = (): gameContextInterface => ({ game, setGame })
  const getProviderGame = useMemo(memoFn, [game, setGame])

  // ===========================================================================

  // hideDropdowns handles the clicks outside a dropdown.
  function hideDropdowns(event: React.MouseEvent<HTMLDivElement, MouseEvent>) {
    const dropdown = document.querySelector('.dropdown-menu') as HTMLElement
    const isMenu = (event.target as HTMLElement).classList.contains(
      'dropdown-content',
    )

    if (!isMenu && dropdown) {
      dropdown.style.display = 'none'
    }
  }

  // getEthersContextDefaultValue returns the default context value for the
  // ethers.js support. This context is used for creating global support for
  // ethers.js connection.
  function getEthersContextDefaultValue(): ethersContextInterface {
    return { ethersConnection, setEthersConnection }
  }

  // handleNetworkChange changes of network. When a Provider makes its initial
  // connection,it emits a "network" event with a null oldNetwork along with the
  // newNetwork. So, if the oldNetwork exists, it represents a changing network.
  function handleNetworkChange(newNetwork: Network): void {
    const fn = async (getConfigResponse: appConfig) => {
      if (newNetwork.chainId !== getConfigResponse.chainId) {
        window.sessionStorage.removeItem('token')
        // It navigates to the wrongNetwork component, and switches the network.
        // If the network isn't switched, it will stay at that component until it does.
        navigate('/wrongNetwork', { state: { ...getConfigResponse } })
        switchNetwork({
          chainId: utils.hexValue(getConfigResponse.chainId),
        })
        return
      }
      navigate('/')
    }
    getAppConfig.then(fn)
  }

  // ===========================================================================

  // effectFn handles network changes.
  const effectFn = () => {
    provider.on('network', (newNetwork, _) => handleNetworkChange(newNetwork))
  }
  // We disable the next line so eslint doens't complain about missing dependencies.
  // eslint-disable-next-line
  useEffect(effectFn, [])

  // ===========================================================================
  // Renders this final markup.
  return (
    <div
      className="App"
      style={{ scrollSnapType: 'y mandatory' }}
      onClick={hideDropdowns}
    >
      <WalletProvider>
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
      </WalletProvider>
    </div>
  )
}

export default App
