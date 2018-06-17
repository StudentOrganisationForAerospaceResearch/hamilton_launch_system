const oxidizerTankPressure = {
  state: {
    pressure: '-'
  },
  mutations: {
    setOxidizerTankPressure (state, newOxidizerTankPressure) {
      state.pressure = newOxidizerTankPressure.pressure.toFixed(0)
    }
  }
}

export default oxidizerTankPressure
