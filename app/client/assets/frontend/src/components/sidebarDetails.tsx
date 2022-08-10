import React from 'react'

const SidebarDetails = () => {
  return (
    <div
    className="details"
    style={{
      display: 'flex',
      alignItems: 'start',
      flexDirection: 'column',
      backgroundColor: 'var(--modals)',
      borderRadius: '0 8px 8px 0',
      margin: '42px 0 42px 0',
      padding: '12px',
      width: '80%',
    }}
    >
    <div className="d-flex">
      <strong className="details__title mr-6">Round:</strong>
      1:20 Dice
    </div>
    <div className="d-flex">
      <strong className="details__title mr-6">Ante:</strong>
      0.1 ETH
    </div>
    <div className="d-flex">
      <strong className="details__title mr-6">Pot:</strong>
      0.1 ETH
    </div>
    </div>
  )
}
export default SidebarDetails