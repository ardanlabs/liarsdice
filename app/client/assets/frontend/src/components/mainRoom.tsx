import React, { Component } from "react";
import { formatEther } from '@ethersproject/units'

interface MainRoomProps {
  show: boolean
  etherBalance: any
}

class MainRoom extends Component<MainRoomProps>{
  render() {
    const { show, etherBalance } = this.props
    
    if (!show) {
      return null
    }
    return (
      <div style={{height: '100%', width: '100%', color: "black", display: 'flex', justifyContent: 'center', alignItems: 'center'}}>
        <div id="players-list" style={{display: 'flex', justifyContent: 'center', alignItems: 'center'}}>
          {formatEther(etherBalance)} ETH
        </div>
      </div>
    )
  }
}

export default MainRoom