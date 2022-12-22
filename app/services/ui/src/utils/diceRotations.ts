const DEFAULT_WIDTH = document.documentElement.clientWidth ?? 1280
const DEFAULT_HEIGHT = document.documentElement.clientHeight ?? 720
const DICE_SPACING = 51.6
const center = { x: DEFAULT_WIDTH / 2, y: DEFAULT_HEIGHT / 2 }

function PlaceOnACircle(
  circle: Phaser.Geom.Circle,
  angle: number,
  angleStep: number,
): {
  angle: number
  position: { rotation: number; x: number; y: number }
} {
  let angleVar = angle
  const x = circle.x + circle.radius * Math.cos(angle)
  const y = circle.y + circle.radius * Math.sin(angle)
  const angleDeg = (Math.atan2(y - center.y, x - center.x) * 180) / Math.PI
  const rotation = angleDeg + 90

  angleVar += angleStep

  return { angle: angleVar, position: { rotation, x, y } }
}

export default PlaceOnACircle
