const barometer = {
  state: {
    pressure: '-',
    temperature: '-'
  },
  mutations: {
    setBarometer (state, newBarometer) {
      state.pressure = newBarometer.pressure.toFixed(1)
      state.temperature = newBarometer.temperature.toFixed(1)
    }
  }
}

export default barometer
