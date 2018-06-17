const oxidizerTankPressure = {
  state: {
    pressure: '-'
  },
  mutations: {
    setOxidizerTankPressure (state, newOxidizerTankPressure) {
      state.pressure = (newOxidizerTankPressure.pressure / 10).toFixed(1)
    }
  }
}

export default oxidizerTankPressure
