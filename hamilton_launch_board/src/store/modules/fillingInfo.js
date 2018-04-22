const fillingInfo = {
  state: {
    totalMass: '-',
    ventValveOpen: false,
    fillValveOpen: false
  },
  mutations: {
    setFillingInfo (state, newFillingInfo) {
      state.totalMass = newFillingInfo.totalMass.toFixed(1)
      state.ventValveOpen = newFillingInfo.ventValveOpen
      state.fillValveOpen = newFillingInfo.fillValveOpen
    }
  }
}

export default fillingInfo
