import React from 'react'

interface SidebarDetailsProps {
  round: number
  ante?: number
  pot?: number
  diceAmount: number
}
const SidebarDetails = (props: SidebarDetailsProps) => {
  const { round, ante, pot, diceAmount } = props
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
      { round ? round : '-' }: { diceAmount } Dice
    </div>
    <div className="d-flex">
      <strong className="details__title mr-6">Ante:</strong>
      { ante } ETH
    </div>
    <div className="d-flex">
      <strong className="details__title mr-6">Pot:</strong>
      { pot } ETH
    </div>
    </div>
  )
}
export default SidebarDetails