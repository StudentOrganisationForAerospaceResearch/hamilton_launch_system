import Vuex from 'vuex'
import Vue from 'vue'

import accelGyroMagnetism from './modules/accelGyroMagnetism'
import weather from './modules/weather'
import barometer from './modules/barometer'
import gps from './modules/gps'
import oxidizerTankConditions from './modules/oxidizerTankConditions'
import socket from './modules/socket'

Vue.use(Vuex)

/**
 * The Vuex store
 */
const store = new Vuex.Store({
  modules: {
    weather,
    accelGyroMagnetism,
    barometer,
    gps,
    oxidizerTankConditions,
    socket
  }
})

export default store
