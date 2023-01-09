import React from 'react'
import Button from './button'
import { JoinProps } from '../types/props.d'

// Join component
function Join(props: JoinProps) {
  // Extracts props.
  const { disabled } = props

  // Checks if button is disabled
  const isButtonDisabled = disabled

  // ===========================================================================
  function handleClick() {}

  // Renders this markup
  return (
    <Button
      disabled={isButtonDisabled}
      classes="join__buton"
      clickHandler={() => handleClick()}
      style={{
        backgroundColor: `${
          isButtonDisabled ? 'grey' : 'var(--primary-color)'
        }`,
      }}
    >
      <span> JOIN GAME</span>
    </Button>
  )
}

export default Join
