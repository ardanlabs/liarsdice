import React from 'react'
import { ButtonProps } from '../types/props.d'

// Button component
function Button(props: ButtonProps) {
  // Extracts props.
  const { clickHandler, id, classes, disabled, children, style, tooltip } =
    props

  return (
    <button
      title={tooltip}
      type="button"
      style={{ cursor: 'pointer', ...style }}
      id={id}
      className={`${classes} btn btn-block`}
      disabled={disabled}
      onClick={() => clickHandler(id)}
    >
      {children}
    </button>
  )
}

export default Button
