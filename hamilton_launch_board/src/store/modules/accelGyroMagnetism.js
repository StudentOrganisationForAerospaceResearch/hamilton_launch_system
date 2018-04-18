const accelGyroMagnetism = {
  state: {
    accel: {
      x: '-',
      y: '-',
      z: '-'
    },
    gyro: {
      x: '-',
      y: '-',
      z: '-'
    },
    magneto: {
      x: '-',
      y: '-',
      z: '-'
    }
  },
  mutations: {
    setAccelGyroMagnetism (state, newAccelGyroMagnetism) {
      state.accel = newAccelGyroMagnetism.accel
      state.gyro = newAccelGyroMagnetism.gyro
      state.magneto = newAccelGyroMagnetism.magneto
    }
  }
}

export default accelGyroMagnetism
