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
      store.commit('setWeather', message)
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
