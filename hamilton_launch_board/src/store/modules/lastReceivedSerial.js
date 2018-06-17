const lastReceivedSerial = {
  state: {
    lastReceived: '-'
  },
  mutations: {
    setLastReceivedSerial (state, newLastReceivedSerial) {
      console.log("bro biiiitch, ", newLastReceivedSerial)
      state.lastReceived = newLastReceivedSerial.lastReceived.toFixed(2)
    }
  }
}

export default lastReceivedSerial
