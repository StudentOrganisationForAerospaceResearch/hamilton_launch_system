<template>
  <div class="inertial-measurement-unit">

    <v-card class="imu-card" raised>
      <v-card-title primary-title>
        <h4 class="headline">Accelerometer</h4>
      </v-card-title>
      <v-divider></v-divider>
      <v-list dense>
        <v-list-tile>
          <v-list-tile-content class="subheading">X</v-list-tile-content>
          <v-list-tile-content class="align-end subheading" >
            {{ accelGyroMagnetism.accel.x }} mG
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile>
          <v-list-tile-content class="subheading">Y</v-list-tile-content>
          <v-list-tile-content class="align-end subheading" >
            {{ accelGyroMagnetism.accel.y }} mG
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile>
          <v-list-tile-content class="subheading">Z</v-list-tile-content>
          <v-list-tile-content class="align-end subheading" >
            {{ accelGyroMagnetism.accel.z }} mG
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-card>

    <v-card class="imu-card" raised>
      <v-card-title primary-title>
        <h4 class="headline">Gyroscope</h4>
      </v-card-title>
      <v-divider></v-divider>
      <v-list dense>
        <v-list-tile>
          <v-list-tile-content class="subheading">X</v-list-tile-content>
          <v-list-tile-content class="align-end subheading" >
            {{ accelGyroMagnetism.gyro.x }} mDps
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile>
          <v-list-tile-content class="subheading">Y</v-list-tile-content>
          <v-list-tile-content class="align-end subheading" >
            {{ accelGyroMagnetism.gyro.y }} mDps
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile>
          <v-list-tile-content class="subheading">Z</v-list-tile-content>
          <v-list-tile-content class="align-end subheading" >
            {{ accelGyroMagnetism.gyro.z }} mDps
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-card>

    <v-card class="imu-card" raised>
      <v-card-title primary-title>
        <h4 class="headline">Magnetometer</h4>
      </v-card-title>
      <v-divider></v-divider>
      <v-list dense>
        <v-list-tile>
          <v-list-tile-content class="subheading">X</v-list-tile-content>
          <v-list-tile-content class="align-end subheading" >
            {{ accelGyroMagnetism.magneto.x }} mGauss
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile>
          <v-list-tile-content class="subheading">Y</v-list-tile-content>
          <v-list-tile-content class="align-end subheading" >
            {{ accelGyroMagnetism.magneto.y }} mGauss
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile>
          <v-list-tile-content class="subheading">Z</v-list-tile-content>
          <v-list-tile-content class="align-end subheading" >
            {{ accelGyroMagnetism.magneto.z }} mGauss
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-card>

  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'InertialMeasurementUnit',
  computed: {
    ...mapState({
      accelGyroMagnetism: state => state.accelGyroMagnetism
    })
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.inertial-measurement-unit {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
}

.imu-card {
  width: 250px;
  margin: 9px;
}

</style>
