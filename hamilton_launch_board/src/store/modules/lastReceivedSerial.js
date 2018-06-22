const lastReceivedSerial = {
  state: {
    avionics: '-',
    groundStation: '-'
  },
  mutations: {
    setLastReceivedSerial (state, newLastReceivedSerial) {
      state.avionics = newLastReceivedSerial.avionics.toFixed(2)
      state.groundStation = newLastReceivedSerial.groundStation.toFixed(2)
    }
  }
}

export default lastReceivedSerial
