// Returns a date string with the format yyyymmddhhmmss
export default function getNowDate(): string {
  const now = new Date()
  const dd = String(now.getDate()).padStart(2, '0')
  const mm = String(now.getMonth() + 1).padStart(2, '0') //January is 0!
  const yyyy = now.getFullYear()
  const hours = now.getHours()
  const minutes = now.getMinutes()
  const seconds = now.getSeconds()
  // We create the specific date to send to the signature
  return yyyy + mm + dd + hours + minutes + seconds
}
