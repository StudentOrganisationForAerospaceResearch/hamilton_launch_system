const launchControlInfo = {
  state: {
    softwareArmCounter: 0,
    launchSystemsArmCounter: 0,
    vpRocketsArmCounter: 0,
    armCounter: 0,
    softareLaunchCounter: 0,
    launchSystemsLaunchCounter: 0,
    vpRocketsLaunchCounter: 0,
    launchCounter: 0,
    countdown: 10
  },
  mutations: {
    setLaunchControlInfo (state, newLaunchControlInfo) {
      state.softwareArmCounter = newLaunchControlInfo.softwareArmCounter
      state.launchSystemsArmCounter = newLaunchControlInfo.launchSystemsArmCounter
      state.vpRocketsArmCounter = newLaunchControlInfo.vpRocketsArmCounter
      state.armCounter = newLaunchControlInfo.armCounter
      state.softareLaunchCounter = newLaunchControlInfo.softareLaunchCounter
      state.launchSystemsLaunchCounter = newLaunchControlInfo.launchSystemsLaunchCounter
      state.vpRocketsLaunchCounter = newLaunchControlInfo.vpRocketsLaunchCounter
      state.launchCounter = newLaunchControlInfo.launchCounter
      state.countdown = newLaunchControlInfo.launchCounter
    }
  }
}

export default launchControlInfo
