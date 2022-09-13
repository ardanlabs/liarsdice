// Returns a date string with the format yyyymmddhhmmss
export default function getNowDate(): string {
  const now = new Date()
  const dd = String(now.getUTCDate()).padStart(2, '0')
  const mm = String(now.getUTCMonth() + 1).padStart(2, '0') //January is 0!
  const yyyy = now.getUTCFullYear()
  const hours = now.getUTCHours()
  const minutes = now.getUTCMinutes()
  let seconds: string | number = now.getUTCSeconds()
  seconds = seconds.toString().length === 1 ? `0${seconds}` : seconds

  // We create the specific date to send to the signature
  return yyyy + mm + dd + hours + minutes + seconds
}
