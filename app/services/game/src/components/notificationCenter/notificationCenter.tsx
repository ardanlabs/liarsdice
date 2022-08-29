import { useState } from 'react'

import { Icons, toast } from 'react-toastify'
import { motion, AnimatePresence } from 'framer-motion'

import { useNotificationCenter } from 'react-toastify/addons/use-notification-center'
import { Trigger } from './trigger'
import Button from '../button'
import { TimeTracker } from './timeTracker'

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

const NotificationCenter = () => {
  const { notifications, clear, remove, unreadCount } = useNotificationCenter()
  const [isOpen, setIsOpen] = useState(false)
  return (
    <div
      style={{
        position: 'absolute',
        top: '35px',
        left: '-275px',
        zIndex: '4',
      }}
    >
      <Trigger onClick={() => setIsOpen(!isOpen)} count={unreadCount} />
      <motion.aside
        initial={false}
        variants={variants.container}
        animate={isOpen ? 'open' : 'closed'}
        style={{
          width: 'min(60ch, 100ch)',
          borderRadius: '8px',
          overflow: 'hidden',
          border: '1px inset var(--secondary-color)',
        }}
      >
        <header
          style={{
            background: 'var(--primary-color)',
            color: 'var(--modals)',
            margin: '0',
            padding: '5px 1rem',
            display: 'flex',
            justifyContent: 'space-between',
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
              background: '#fff',
              height: '400px',
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
                    style={{ padding: '0.8rem' }}
                  >
                    <motion.article
                      key={notification.id}
                      variants={variants.item}
                      style={{
                        display: 'grid',
                        gridTemplateColumns: '40px 1fr 40px',
                        gap: '8px',
                        padding: '0.8rem',
                        background: 'rgba(0, 0, 0, 0.1)',
                        borderRadius: '8px',
                      }}
                    >
                      <div style={{ width: '32px' }}>
                        {notification.icon ||
                          Icons.info({
                            theme: notification.theme || 'light',
                            type: toast.TYPE.INFO,
                          })}
                      </div>
                      <div>
                        <div>{notification.content}</div>
                        <TimeTracker createdAt={notification.createdAt} />
                      </div>
                      <Button
                        clickHandler={() => remove(notification.id)}
                        tooltip="Archive"
                      >
                        <svg
                          style={{ width: '24px', height: '24px' }}
                          viewBox="0 0 24 24"
                        >
                          <path
                            fill="currentColor"
                            d="M6,19A2,2 0 0,0 8,21H16A2,2 0 0,0 18,19V7H6V19M8,9H16V19H8V9M15.5,4L14.5,3H9.5L8.5,4H5V6H19V4H15.5Z"
                          />
                        </svg>
                      </Button>
                    </motion.article>
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
            padding: '1rem',
            display: 'flex',
            justifyContent: 'space-between',
            alignItems: 'center',
          }}
        >
          <Button style={{ color: 'var(--modals)' }} clickHandler={clear}>
            Clear All
          </Button>
        </footer>
      </motion.aside>
    </div>
  )
}

export default NotificationCenter
