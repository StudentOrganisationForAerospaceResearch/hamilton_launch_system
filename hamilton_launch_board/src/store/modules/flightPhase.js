const flightPhase = {
  state: {
    phase: '-'
  },
  mutations: {
    setFlightPhase (state, newFlightPhase) {
      switch (newFlightPhase.flightPhase) {
        case 0:
          state.phase = 'PRELAUNCH'
          break
        case 1:
          state.phase = 'BURN'
          break
        case 2:
          state.phase = 'COAST'
          break
        case 3:
          state.phase = 'DROGUE_DESCENT'
          break
        case 4:
          state.phase = 'MAIN_DESCENT'
          break
        case 5:
          state.phase = 'ABORT'
          break
        default:
          state.phase = 'unknown'
      }
    }
  }
}

export default flightPhase
