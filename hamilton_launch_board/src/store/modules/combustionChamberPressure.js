const combustionChamberPressure = {
  state: {
    pressure: '-'
  },
  mutations: {
    setCombustionChamberPressure (state, newCombustionChamberPressure) {
      state.pressure = (newCombustionChamberPressure.pressure / 1000).toFixed(3)
    }
  }
}

export default combustionChamberPressure
