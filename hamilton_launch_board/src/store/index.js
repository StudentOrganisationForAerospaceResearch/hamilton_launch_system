import Vuex from 'vuex'
import Vue from 'vue'

import accelGyroMagnetism from './modules/accelGyroMagnetism'
import barometer from './modules/barometer'
import gps from './modules/gps'
import oxidizerTankPressure from './modules/oxidizerTankPressure'
import combustionChamberPressure from './modules/combustionChamberPressure'
import flightPhase from './modules/flightPhase'
import fillingInfo from './modules/fillingInfo'
import launchControlInfo from './modules/launchControlInfo'
import socket from './modules/socket'

Vue.use(Vuex)

/**
 * The Vuex store
 */
const store = new Vuex.Store({
  modules: {
    accelGyroMagnetism,
    barometer,
    gps,
    oxidizerTankPressure,
    combustionChamberPressure,
    flightPhase,
    fillingInfo,
    launchControlInfo,
    socket
  }
})

export default store
