const oxidizerTankConditions = {
  state: {
    pressure: '-',
    temperature: '-'
  },
  mutations: {
    setOxidizerTankConditions (state, newOxidizerTankConditions) {
      state.pressure = newOxidizerTankConditions.pressure.toFixed(1)
      state.temperature = newOxidizerTankConditions.temperature.toFixed(1)
    }
  }
}

export default oxidizerTankConditions
