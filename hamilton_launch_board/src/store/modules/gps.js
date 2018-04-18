const gps = {
  state: {
    epochTimeMsec: '-',
    altitude: '-',
    latitude: '-',
    longitude: '-'
  },
  mutations: {
    setGps (state, newGps) {
      state.epochTimeMsec = newGps.epochTimeMsec.toFixed(1)
      state.altitude = newGps.altitude.toFixed(1)
      state.latitude = newGps.latitude.toFixed(1)
      state.longitude = newGps.longitude.toFixed(1)
    }
  }
}

export default gps
