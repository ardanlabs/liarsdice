import React, { FC } from 'react'

interface LiarsCallProps {}

const LiarsCall: FC<LiarsCallProps> = (LiarsCallProps) => {
  return (
    <div
      style={{
        display: 'flex',
        height: '70px',
        width: '60%',
        justifyContent: 'center',
        textAlign: 'center',
        alignItems: 'center',
        color: 'black',
        backgroundColor: 'var(--modals)',
        borderRadius: '8px',
        fontSize: '28px',
        fontWeight: '500',
        padding: '8px',
      }}
    >
      <span>Player two called Player one a liar and got striked</span>
    </div>
  )
}

export default LiarsCall
