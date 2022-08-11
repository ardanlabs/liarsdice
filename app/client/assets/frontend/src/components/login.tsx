import React, { useMemo, useState } from 'react'
import Button from './button'
import MetamaskLogo from './icons/metamask'
import { useEthers } from '@usedapp/core'
import MainRoom from './mainRoom'
import { GameContext } from '../gameContext'
import { game } from '../types/index.d'

export default function Login() {
  const { account, activateBrowserWallet } = useEthers()
  const [ game, setGame ] = useState({
    status: 'open',
    round: 0,
    current_player: '',
    player_order: [],
    players: [],
  } as game)

  const providerGame = useMemo(() => ({game, setGame}), [game, setGame])

  function handleConnectWallet() {
    activateBrowserWallet()
  }
  return account?.length ? (
    <GameContext.Provider
      value={providerGame}
    >
      <MainRoom />
    </GameContext.Provider>
  ) : (
    <div
      id="login__wrapper"
      className="d-flex align-items-start justify-content-center flex-column"
    >
      <h2>
        <strong> Connect your wallet </strong>
      </h2>
      Or you can also select a provider to create one.
      <div id="wallets__wrapper" className="mt-4">
        <Button
          {...{
            id: 'metamask__wrapper',
            clickHandler: handleConnectWallet,
            classes: 'd-flex align-items-center pa-4',
          }}
        >
          <MetamaskLogo {...{ width: '50px', height: '50px' }} />
          <span className="ml-4"> Metamask </span>
        </Button>
      </div>
    </div>
  )
}
