const fillingInfo = {
  state: {
    totalMass: '-',
    ventValveOpen: false,
    fillValveOpen: false
  },
  mutations: {
    setVentValveStatus (state, newVentValveStatus) {
      state.ventValveOpen = newVentValveStatus.ventValveOpen
    },
    setFillValveStatus (state, newFillValveStatus) {
      state.fillValveOpen = newFillValveStatus.fillValveOpen
    },
    setRocketMass (state, newRocketMass) {
      state.totalMass = newRocketMass.rocketMass
    }
  }
}

export default fillingInfo
