import Vuex from 'vuex'
import Vue from 'vue'

import weather from './modules/weather'
import socket from './modules/socket'

Vue.use(Vuex)

/**
 * The Vuex store
 */
const store = new Vuex.Store({
  modules: {
    weather,
    socket
  }
})

export default store
