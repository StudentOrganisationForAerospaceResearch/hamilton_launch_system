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
    setTotalMass (state, newTotalMass) {
      state.totalMass = newTotalMass.totalMass.toFixed(1)
    }
  }
}

export default fillingInfo
