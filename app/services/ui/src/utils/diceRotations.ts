const DEFAULT_WIDTH = document.documentElement.clientWidth ?? 1280
const DEFAULT_HEIGHT = document.documentElement.clientHeight ?? 720
const DICE_SPACING = 51.6
const center = { x: DEFAULT_WIDTH / 2, y: DEFAULT_HEIGHT / 2 }

function getDicePosition(
  initialYPosParam: number,
  xParam: number,
  yParam: number,
  rotationParam: number,
  i: number,
): {
  initialYPos: number
  position: { rotation: number; x: number; y: number }
} {
  var initialYPos = initialYPosParam
  var x = xParam
  var y = yParam
  var rotation = rotationParam

  switch (i) {
    case 0:
      return { initialYPos, position: { rotation, x, y } }
    case 1:
      x += DICE_SPACING
      y = initialYPos + 11
      rotation = 0.11
      break
    case 2:
      x += DICE_SPACING
      y = initialYPos + 13
      rotation = 0
      break
    case 3:
      x += DICE_SPACING
      y = initialYPos + 11
      rotation = -0.11
      break
    case 4:
      x += DICE_SPACING
      y = initialYPos
      rotation = -0.3
      break
    // Second to the left
    case 5:
      initialYPos = DEFAULT_HEIGHT / 2 - DICE_SPACING * 2
      x = DEFAULT_WIDTH / 2 - 300
      y = initialYPos
      rotation = 0.3
      return { initialYPos, position: { rotation, x, y } }
    case 6:
      y -= DICE_SPACING
      x = DEFAULT_WIDTH / 2 - 282
      rotation = -1.1
      break
    case 7:
      y -= DICE_SPACING
      x = DEFAULT_WIDTH / 2 - 250
      rotation = -0.9
      break
    case 8:
      y -= DICE_SPACING - 10
      x = DEFAULT_WIDTH / 2 - 210
      rotation = 0.9
      break
    case 9:
      y -= DICE_SPACING - 15
      x = DEFAULT_WIDTH / 2 - 160
      rotation = 1.1
      break

    // First to the left
    case 10:
      initialYPos = DEFAULT_WIDTH / 2 - 320
      x = initialYPos + 1
      y = center.y + 20
      rotation = -0.1
      return { initialYPos, position: { rotation, x, y } }
    case 11:
      y += DICE_SPACING
      x = initialYPos + 8
      rotation = -0.23
      break
    case 12:
      y += DICE_SPACING
      x = initialYPos + 20
      rotation = 1.23
      break
    case 13:
      y += DICE_SPACING
      x = initialYPos + 46
      rotation = -2.15
      break
    case 14:
      y += DICE_SPACING - 10
      x = initialYPos + 75
      rotation = 0.9
      break
    case 15:
      initialYPos = DEFAULT_WIDTH / 2 + 320
      x = initialYPos + 1
      y = center.y + 20
      rotation = 0.1
      return { initialYPos, position: { rotation, x, y } }
    case 16:
      y += DICE_SPACING
      x = initialYPos - 8
      rotation = 0.23
      break
    case 17:
      y += DICE_SPACING
      x = initialYPos - 20
      rotation = -1.23
      break
    case 18:
      y += DICE_SPACING
      x = initialYPos - 46
      rotation = 2.15
      break
    case 19:
      y += DICE_SPACING - 10
      x = initialYPos - 75
      rotation = -0.9
      break
    case 20:
      initialYPos = DEFAULT_HEIGHT / 2 - DICE_SPACING * 2
      x = DEFAULT_WIDTH / 2 + 300
      y = initialYPos
      rotation = -0.3
      return { initialYPos, position: { rotation, x, y } }
    case 21:
      y -= DICE_SPACING
      x = DEFAULT_WIDTH / 2 + 282
      rotation = 1.1
      break
    case 22:
      y -= DICE_SPACING
      x = DEFAULT_WIDTH / 2 + 250
      rotation = 0.9
      break
    case 23:
      y -= DICE_SPACING - 10
      x = DEFAULT_WIDTH / 2 + 210
      rotation = -0.9
      break
    case 24:
      y -= DICE_SPACING - 15
      x = DEFAULT_WIDTH / 2 + 160
      rotation = -1.1
      break
  }

  return { initialYPos, position: { rotation, x, y } }
}

// initialYPos = DEFAULT_HEIGHT / 2 + 200
// x = center.x - DICE_SPACING * 2
// y = initialYPos
// rotation = 0.41
// switch (i) {
//   case 1:
//     y = initialYPos + 20 * 1
//     rotation = 0.21
//     break
//   case 2:
//     y = initialYPos + 13 * 2
//     rotation = 0
//     break
//   case 3:
//     y = initialYPos + 20 * 1
//     rotation = -0.21
//     break
//   case 4:
//     y = initialYPos
//     rotation = -0.41
//     break
// }

// initialYPos = DEFAULT_HEIGHT / 2 - 200
// x = DEFAULT_WIDTH / 2 - DICE_SPACING * 2
// y = initialYPos
// rotation = -0.41
// switch (i) {
//   case 1:
//     y = initialYPos - 20 * 1
//     rotation = -0.21
//     break
//   case 2:
//     y = initialYPos - 13 * 2
//     rotation = 0
//     break
//   case 3:
//     y = initialYPos - 20 * 1
//     rotation = 0.21
//     break
//   case 4:
//     y = initialYPos
//     rotation = 0.41
//     break
// }

// initialYPos = DEFAULT_WIDTH / 2 - 300
// x = initialYPos
// y = center.y - DICE_SPACING * 2
// rotation = 1.98079633376
// switch (i) {
//   case 1:
//     x = initialYPos - 20 * 1
//     rotation = 1.7807962622799999863
//     break
//   case 2:
//     x = initialYPos - 13 * 2
//     rotation = 0
//     break
//   case 3:
//     x = initialYPos - 20 * 1
//     rotation = -1.7807962622799999863
//     break
//   case 4:
//     x = initialYPos
//     rotation = -1.98079633376
//     break
// }

// initialYPos = DEFAULT_WIDTH / 2 + 300
// x = initialYPos
// y = center.y - DICE_SPACING * 2
// rotation = -1.98079633376

// switch (i) {
//   case 1:
//     x = initialYPos + 20 * 1
//     rotation = -1.7807962622799999863
//     break
//   case 2:
//     x = initialYPos + 13 * 2
//     rotation = 0
//     break
//   case 3:
//     x = initialYPos + 20 * 1
//     rotation = 1.7807962622799999863
//     break
//   case 4:
//     x = initialYPos
//     rotation = 1.98079633376
//     break
// }
export default getDicePosition
