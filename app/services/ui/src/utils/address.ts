import { utils } from 'ethers'
import { Falsy } from '../types/index.d'

export function shortenString(str: string) {
  return str.substring(0, 6) + '...' + str.substring(str.length - 4)
}

export function shortenAddress(address: string): string {
  try {
    const formattedAddress = utils.getAddress(address)
    return shortenString(formattedAddress)
  } catch {
    throw new TypeError("Invalid input, address can't be parsed")
  }
}

export function shortenIfAddress(address: string | Falsy): string {
  if (typeof address === 'string' && address.length > 0) {
    return shortenAddress(address)
  }
  return ''
}
