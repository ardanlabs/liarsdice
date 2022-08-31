import React, { FC } from 'react'

interface ButtonProps {
  clickHandler: Function
  classes?: string
  id?: string
  disabled?: boolean
  children: JSX.Element[] | JSX.Element | string
  style?: React.CSSProperties
  tooltip?: string
}

const Button: FC<ButtonProps> = (ButtonProps) => {
  let { classes } = ButtonProps
  const { clickHandler, id, disabled, children, style, tooltip } = ButtonProps
  classes = classes ? `${classes} btn btn-block` : 'btn btn-block'
  return (
    <button
      title={tooltip}
      type="button"
      style={{ cursor: 'pointer', ...style }}
      id={id}
      className={classes}
      disabled={disabled}
      onClick={() => clickHandler(id)}
    >
      {children}
    </button>
  )
}

export default Button
