const lastReceivedSerial = {
  state: {
    lastReceived: '-'
  },
  mutations: {
    setLastReceivedSerial (state, newLastReceivedSerial) {
      state.lastReceived = newLastReceivedSerial.lastReceived.toFixed(2)
    }
  }
}

export default lastReceivedSerial
