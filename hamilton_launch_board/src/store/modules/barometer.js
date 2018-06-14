const barometer = {
  state: {
    pressure: '-',
    temperature: '-'
  },
  mutations: {
    setBarometer (state, newBarometer) {
      state.pressure = newBarometer.pressure / 100
      state.temperature = newBarometer.temperature / 100
    }
  }
}

export default barometer
