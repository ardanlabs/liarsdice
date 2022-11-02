import { ReactNode, useState } from 'react'

import { motion, AnimatePresence, MotionStyle } from 'framer-motion'

import { useNotificationCenter } from 'react-toastify/addons/use-notification-center'
import Trigger from './trigger'
import Button from '../button'

// contains framer-motion variants to animate different parts of the UI
// when the notification center is visible or not
// https://www.framer.com/docs/examples/#variants
const variants = {
  container: {
    open: {
      y: 0,
      x: 291,
      opacity: 1,
    },
    closed: {
      y: -10,
      x: 291,
      opacity: 0,
    },
  },
  // used to stagger item animation when switching from closed to open and vice versa
  content: {
    open: {
      transition: { staggerChildren: 0.07, delayChildren: 0.2 },
    },
    closed: {
      transition: { staggerChildren: 0.05, staggerDirection: -1 },
    },
  },
  item: {
    open: {
      y: 0,
      opacity: 1,
      transition: {
        y: { stiffness: 1000, velocity: -100 },
      },
    },
    closed: {
      y: 50,
      opacity: 0,
      transition: {
        y: { stiffness: 1000 },
      },
    },
  },
}

interface NotificationCenterProps {
  trigger?: boolean
  mainContainerStyle?: React.CSSProperties
  asideContainerStyle?: MotionStyle
  notificationCenterWidth: string
}

// NotificationCenter component
function NotificationCenter(props: NotificationCenterProps) {
  // Extracts constant props.
  const { trigger, asideContainerStyle, notificationCenterWidth } = props

  // Extracts variable props.
  let { mainContainerStyle } = props

  // Extracts functions and constants from the useNotificationCenter hook.
  const { notifications, clear, unreadCount } = useNotificationCenter()

  // Sets state to handle if the component is open.
  const [isOpen, setIsOpen] = useState(false)

  // If the component has a trigger sets the style object properly.
  mainContainerStyle =
    trigger && !mainContainerStyle
      ? {
          width: 'min(60ch, 100ch)',
          borderRadius: '8px',
          overflow: 'hidden',
          border: '1px inset var(--secondary-color)',
        }
      : mainContainerStyle

  // Renders this markup.
  return (
    <div style={mainContainerStyle}>
      {trigger ? (
        <Trigger onClick={() => setIsOpen(!isOpen)} count={unreadCount} />
      ) : null}
      <motion.aside
        initial={!trigger}
        variants={variants.container}
        style={asideContainerStyle}
      >
        <header
          style={{
            background: 'var(--primary-color)',
            color: 'var(--modals)',
            margin: '0',
            padding: '5px 1rem',
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
          }}
        >
          <h3>Notifications</h3>
        </header>
        <AnimatePresence>
          <motion.section
            variants={variants.content}
            animate={isOpen ? 'open' : 'closed'}
            style={{
              background: 'var(--modals)',
              height: `${trigger ? '400px' : '100%'}`,
              overflowY: 'scroll',
              overflowX: 'hidden',
              color: '#000',
              padding: '0.2rem',
              position: 'relative',
            }}
          >
            {!notifications.length && (
              <motion.h4
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                exit={{ opacity: 0 }}
                style={{
                  margin: '0',
                  textAlign: 'center',
                  padding: '2rem',
                  color: 'var(--secondary-color)',
                }}
              >
                Your queue is empty! you are all set{' '}
                <span role="img" aria-label="dunno what to put">
                  🎉
                </span>
              </motion.h4>
            )}
            <AnimatePresence>
              {notifications.map((notification) => {
                return (
                  <motion.div
                    key={notification.id}
                    layout
                    initial={{ scale: 0.4, opacity: 0, y: 50 }}
                    exit={{
                      scale: 0,
                      opacity: 0,
                      transition: { duration: 0.2 },
                    }}
                    animate={{ scale: 1, opacity: 1, y: 0 }}
                    style={{
                      padding: '0.3rem',
                      textAlign: 'left',
                      fontWeight: '600',
                      borderBottom: '1px solid #9f9f9f',
                    }}
                  >
                    {notification.content as ReactNode}
                  </motion.div>
                )
              })}
            </AnimatePresence>
          </motion.section>
        </AnimatePresence>
        <footer
          style={{
            background: 'var(--primary-color)',
            color: 'var(--modals)',
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'center',
            position: 'absolute',
            bottom: '0',
            width: `${notificationCenterWidth}`,
          }}
        >
          <Button
            style={{
              color: 'var(--modals)',
              fontSize: '20px',
              fontWeight: '500',
            }}
            clickHandler={clear}
          >
            Clear All
          </Button>
        </footer>
      </motion.aside>
    </div>
  )
}

export default NotificationCenter
