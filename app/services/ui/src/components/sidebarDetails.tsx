import React, { useContext } from 'react'
import { GameContext } from '../contexts/gameContext'
import { shortenIfAddress } from '../utils/address'
import useEthersConnection from './hooks/useEthersConnection'
import useGame from './hooks/useGame'
import Players from './players'

// SideBarDetails component
function SidebarDetails() {
  // Extracts properties from useGame Hook.
  const { gamePot } = useGame()

  // Extracts game from useContext hook
  const { game } = useContext(GameContext)

  // Deconstructs props from game
  const { status, anteUSD, round, lastWin, lastOut } = game

  const { account } = useEthersConnection()

  // Renders this markup
  return (
    <div
      style={{
        display: 'flex',
        alignItems: 'start',
        flexDirection: 'column',
        backgroundColor: 'var(--modals)',
        width: '100%',
        flexGrow: '1',
      }}
    >
      <div
        className="details"
        style={{
          padding: '16px 10px',
        }}
      >
        <div className="d-flex">
          <strong className="details__title mr-6">Round:</strong>
          {round ? round : '-'}
        </div>
        <div className="d-flex">
          {anteUSD ? (
            <>
              <strong className="details__title mr-6">Ante:</strong>
              {anteUSD} U$D
            </>
          ) : (
            ''
          )}
        </div>
        <div className="d-flex">
          {gamePot ? (
            <>
              <strong className="details__title mr-6">Pot:</strong>
              {gamePot} U$D
            </>
          ) : (
            ''
          )}
        </div>
        <div className="d-flex">
          {status ? (
            <>
              <strong className="details__title mr-6">Status:</strong>
              {status}
            </>
          ) : (
            ''
          )}
        </div>
        <div className="d-flex">
          {lastWin ? (
            <>
              <strong className="details__title mr-6">Last win:</strong>
              {shortenIfAddress(lastWin)}
            </>
          ) : (
            ''
          )}
        </div>
        <div className="d-flex">
          {lastOut ? (
            <>
              <strong className="details__title mr-6">My address:</strong>
              {shortenIfAddress(lastOut)}
            </>
          ) : (
            ''
          )}
        </div>
        <div className="d-flex">
          {account ? (
            <>
              <strong className="details__title mr-6">My address:</strong>
              {shortenIfAddress(account)}
            </>
          ) : (
            ''
          )}
        </div>
      </div>
      <Players />
    </div>
  )
}
export default SidebarDetails
