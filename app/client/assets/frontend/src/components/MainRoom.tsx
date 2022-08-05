import React, { Component } from "react";

interface MainRoomProps {
  show: boolean
}

class MainRoom extends Component<MainRoomProps>{
  render() {
    const { show } = this.props
    if (!show) {
      return null
    }
    return (
      <div style={{background: 'red', height: '100%', width: '100%'}}></div>
    )
  }
}

export default MainRoom