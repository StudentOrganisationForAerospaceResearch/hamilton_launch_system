import store from '@/store'

const socket = {
  state: {
    isConnected: false,
    reconnectError: false
  },
  mutations: {
    SOCKET_ONOPEN (state, event) {
      state.isConnected = true
    },
    SOCKET_ONCLOSE (state, event) {
      state.isConnected = false
    },
    SOCKET_ONERROR (state, event) {
      console.error(state, event)
    },
    // default handler called for all methods
    SOCKET_ONMESSAGE (state, message) {
      switch (message.type) {
        case 'accelGyroMagnetism':
          store.commit('setAccelGyroMagnetism', message)
          break
        case 'barometer':
          store.commit('setBarometer', message)
          break
        case 'gps':
          store.commit('setGps', message)
          break
        case 'oxidizerTankPressure':
          store.commit('setOxidizerTankPressure', message)
          break
        case 'combustionChamberPressure':
          store.commit('setCombustionChamberPressure', message)
          break
        case 'flightPhase':
          store.commit('setFlightPhase', message)
          break
        case 'launchControlInfo':
          store.commit('setLaunchControlInfo', message)
          break
        case 'ventStatus':
          store.commit('setVentValveStatus', message)
          break
        case 'fillValveStatus':
          store.commit('setFillValveStatus', message)
          break
        case 'loadCellData':
          store.commit('setTotalMass', message)
          break
      }
    },
    // mutations for reconnect methods
    SOCKET_RECONNECT (state, count) {
      console.info(state, count)
    },
    SOCKET_RECONNECT_ERROR (state) {
      state.reconnectError = true
    }
  }
}

export default socket
