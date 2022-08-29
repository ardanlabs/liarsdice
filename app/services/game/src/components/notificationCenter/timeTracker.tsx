import dayjs from 'dayjs'
import duration from 'dayjs/plugin/duration'
import relativeTime from 'dayjs/plugin/relativeTime'
import { useEffect, useReducer, useRef } from 'react'

dayjs.extend(duration)
dayjs.extend(relativeTime)

interface Props {
  createdAt: number
}

export function TimeTracker({ createdAt }: Props) {
  const [, forceUpdate] = useReducer((x) => x + 1, 0)
  const intervalRef = useRef<NodeJS.Timer>()

  // refresh value of `createdAt` every ~ 1 minute
  useEffect(() => {
    intervalRef.current = setInterval(() => {
      forceUpdate()
    }, 1000)

    return () => {
      clearInterval(intervalRef.current)
    }
  }, [])

  return (
    <div style={{ color: '#666' }}>
      <span>{dayjs(createdAt).fromNow()}</span>
    </div>
  )
}
