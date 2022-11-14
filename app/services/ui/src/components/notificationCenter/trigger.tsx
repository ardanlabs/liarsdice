interface Props {
  onClick: () => void
  count: number
}

// Trigger component
function Trigger({ count, onClick }: Props) {
  // Renders this markup.
  return (
    <div
      onClick={onClick}
      className="notification__trigger"
      style={{
        position: 'relative',
        display: 'inline-block',
        color: '#fff',
        marginBottom: '10px',
        cursor: 'pointer',
      }}
    >
      <svg style={{ width: '24px', height: '24px' }} viewBox="0 0 24 24">
        <path
          fill="currentColor"
          d="M21,19V20H3V19L5,17V11C5,7.9 7.03,5.17 10,4.29C10,4.19 10,4.1 10,4A2,2 0 0,1 12,2A2,2 0 0,1 14,4C14,4.1 14,4.19 14,4.29C16.97,5.17 19,7.9 19,11V17L21,19M14,21A2,2 0 0,1 12,23A2,2 0 0,1 10,21"
        />
      </svg>
      <span
        style={{
          position: 'absolute',
          top: '-8px',
          right: '-8px',
          background: 'red',
          borderRadius: '4px',
          width: '22px',
          height: '22px',
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
        }}
      >
        {count}
      </span>
    </div>
  )
}

export default Trigger
