import React from 'react'
import { fireEvent, render, screen } from '@testing-library/react'
import Button from './button'

test('handles onClick', () => {
  const onClick = jest.fn()
  // Tests what happens if you send a number to the Die component.
  // Checks if that number is rendered
  const props = {
    clickHandler: onClick,
    id: 'testing-id',
    disabled: false,
    style: { width: '100%' },
    classes: 'testing',
  }
  render(
    <Button {...props}>
      <>Testing</>
    </Button>,
  )
  const ButtonElement = screen.getByText('Testing')
  fireEvent.click(ButtonElement)
  expect(onClick).toHaveBeenCalledTimes(1)
  expect(ButtonElement.id).toEqual('testing-id')
  expect(ButtonElement.style.width).toEqual('100%')
  expect(ButtonElement.classList).toContain('testing')
  expect(ButtonElement.classList).toContain('btn')
  expect(ButtonElement.classList).toContain('btn-block')
})
