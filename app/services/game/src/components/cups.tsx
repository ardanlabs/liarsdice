import React, { useContext } from 'react'
import { bet, user } from '../types/index.d'
import Star from './icons/star'
import Bet from './bet'
import { GameContext } from '../contexts/gameContext'
import { shortenIfAddress } from '../utils/address'
import { motion, AnimatePresence } from 'framer-motion'

// Cups Component
// Renders the players ui
function Cups() {
  // Extracts the game using the useContext Hook
  const { game } = useContext(GameContext)

  // Extracts properties from the game
  const { cups, currentID, status } = game

  // Initialize the cups array.
  const cupsElements: JSX.Element[] = []

  // Iterates the game cups to create each part of the UI.
  ;(cups as user[]).forEach((player: user, i: number) => {
    const bets = game.bets
      ? game.bets.filter((bet: bet) => bet.account === player.account)
      : []

    const isPlayerTurn = currentID === player.account && status === 'playing'

    const isPlayerActive = player.outs < 3

    const stars: JSX.Element[] = []

    for (let i = 1; i <= 3; i++) {
      stars.push(
        <Star
          key={Math.random()}
          fill={i > player.outs ? '#F0EAD6' : 'var(--primary-color)'}
        />,
      )
    }

    const cupAnimationsVariants = {
      initial: {
        x: '50%',
        opacity: 0,
        scale: 0.5,
      },
      animate: {
        x: 0,
        opacity: 1,
        scale: 1,
      },
      transition: { staggerChildren: 0.07, delayChildren: 0.2, duration: 0.2 },
    }

    // getCup returns the player cup ui
    function getCup() {
      return (
        <motion.div
          initial={cupAnimationsVariants.initial}
          animate={cupAnimationsVariants.animate}
          transition={cupAnimationsVariants.transition}
          key={player.account}
          className={`player__ui ${isPlayerActive ? 'active' : ''}`}
          data-testid="player__ui"
        >
          <div className="d-flex">{stars}</div>
          <p
            style={{ fontWeight: '600' }}
            className={isPlayerTurn ? 'own_turn' : ''}
          >{`Player ${shortenIfAddress(player.account)}`}</p>
          {isPlayerActive ? (
            <div className={`bet`}>
              {bets[0] ? 'Bet: ' : ''}
              <Bet
                bet={bets[0]}
                dieWidth="27"
                dieHeight="27"
                fill="var(--modals)"
              />
            </div>
          ) : null}
        </motion.div>
      )
    }

    cupsElements.push(getCup())
  })

  // Renders this markup
  return (
    <div
      style={{
        display: 'flex',
        width: '100%',
        alignItems: 'center',
        justifyContent: 'center',
        flexWrap: 'wrap',
      }}
      id="cupsContainer"
    >
      {/* <AnimatePresence> */}
      {cupsElements}
      {/* </AnimatePresence> */}
    </div>
  )
}

export default Cups
