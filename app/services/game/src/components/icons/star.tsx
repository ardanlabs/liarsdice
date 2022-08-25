import React, { FC } from 'react'

interface StarProps {
  width?: string
  height?: string
  fill?: string
}

const Star: FC<StarProps> = (StarProps) => {
  let { width, height, fill } = StarProps
  width = width ? width : '29px'
  height = height ? height : '26px'
  fill = fill ? fill : '#F0EAD6'
  return (
    <svg
      id="star_svg"
      style={{ width: width, height: height }}
      width="29"
      height="26"
      viewBox="0 0 29 26"
    >
      <path
        d="M14.5559 1.62318L14.4141 1.21443L14.2724 1.62318L11.2957 10.2084H1.65653H1.1697L1.57212 10.4824L9.36182 15.7861L6.38778 24.3636L6.23932 24.7918L6.61392 24.5368L14.4141 19.226L22.2143 24.5368L22.5889 24.7918L22.4405 24.3636L19.4664 15.7861L27.2561 10.4824L27.6586 10.2084H27.1717H17.5325L14.5559 1.62318Z"
        fill={fill}
        stroke="#35495C"
        strokeWidth="0.3"
      />
    </svg>
  )
}

export default Star
