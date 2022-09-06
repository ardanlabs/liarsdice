import React from 'react'
import SignOut from './signout'
import useEthersConnection from './hooks/useEthersConnection'
import GameActions from './GameActions'

// Footer component.
// It's responsible for rendering the footer of the page.
// It contains the GameActions component.
function Footer() {
  // Extracts account from ethersConnection Hook
  const { account } = useEthersConnection()

  // Renders if there's an wallet connected.
  return account ? (
    <footer
      id="footer"
      style={{
        backgroundColor: 'var(--modals)',
        position: 'fixed',
        bottom: '0',
        height: '70px',
        width: '100%',
        display: 'flex',
        justifyContent: 'start',
        alignItems: 'center',
      }}
    >
      <div
        style={{
          width: 'fit-content',
        }}
      >
        <SignOut disabled={!account} />
      </div>
      <GameActions />
    </footer>
  ) : null
}

export default Footer
