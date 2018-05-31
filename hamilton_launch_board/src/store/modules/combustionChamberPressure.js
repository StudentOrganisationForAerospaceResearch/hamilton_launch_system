const combustionChamberPressure = {
  state: {
    pressure: '-'
  },
  mutations: {
    setCombustionChamberPressure (state, newCombustionChamberPressure) {
      state.pressure = newCombustionChamberPressure.pressure.toFixed(1)
    }
  }
}

export default combustionChamberPressure
