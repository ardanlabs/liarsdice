import React, { FC } from 'react'

interface ButtonProps {
  clickHandler: Function,
  classes?: string,
  id?: string,
  disabled?: boolean,
  children: JSX.Element[] | JSX.Element
}

const Button: FC<ButtonProps> = (ButtonProps) => {
  let { classes } = ButtonProps
  const { clickHandler, id, disabled, children } = ButtonProps
  classes = classes ? `${classes} btn btn-block` : 'btn btn-block'
  return (
    <button type="button" id={id} className={classes} disabled={disabled} onClick={() => clickHandler(id)}>
      { children }
    </button>
  )
}

export default Button
