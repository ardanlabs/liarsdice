import React, { FC } from 'react'

interface ButtonProps {
  clickHandler: Function
  classes?: string
  id?: string
  disabled?: boolean
  children: JSX.Element[] | JSX.Element
  style?: React.CSSProperties
}

const Button: FC<ButtonProps> = (ButtonProps) => {
  let { classes } = ButtonProps
  const { clickHandler, id, disabled, children, style } = ButtonProps
  classes = classes ? `${classes} btn btn-block` : 'btn btn-block'
  return (
    <button
      type="button"
      style={style}
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
