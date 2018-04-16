const weather = {
  state: {
    airTemperature: '-',
    windSpeed: '-',
    windDirection: '-',
    relativeHumidity: '-'
  },
  mutations: {
    setWeather (state, newWeather) {
      state.airTemperature = newWeather.airTemperature.toFixed(1)
      state.windSpeed = newWeather.windSpeed.toFixed(1)
      state.windDirection = newWeather.windDirection.toFixed(1)
      state.relativeHumidity = newWeather.relativeHumidity.toFixed(1)
    }
  }
}

export default weather
