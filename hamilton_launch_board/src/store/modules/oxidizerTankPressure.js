const oxidizerTankPressure = {
  state: {
    pressure: '-'
  },
  mutations: {
    setOxidizerTankPressure (state, newOxidizerTankPressure) {
      state.pressure = newOxidizerTankPressure.pressure.toFixed(1)
    }
  }
}

export default oxidizerTankPressure
