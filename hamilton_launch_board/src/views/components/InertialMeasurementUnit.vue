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
            {{ accelerometer.x }} m/s
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile>
          <v-list-tile-content class="subheading">Y</v-list-tile-content>
          <v-list-tile-content class="align-end subheading" >
            {{ accelerometer.y }} m/s
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile>
          <v-list-tile-content class="subheading">Z</v-list-tile-content>
          <v-list-tile-content class="align-end subheading" >
            {{ accelerometer.z }} m/s
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
            {{ gyroscope.x }} m/s
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile>
          <v-list-tile-content class="subheading">Y</v-list-tile-content>
          <v-list-tile-content class="align-end subheading" >
            {{ gyroscope.y }} m/s
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile>
          <v-list-tile-content class="subheading">Z</v-list-tile-content>
          <v-list-tile-content class="align-end subheading" >
            {{ gyroscope.z }} m/s
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
            {{ magnetometer.x }} m/s
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile>
          <v-list-tile-content class="subheading">Y</v-list-tile-content>
          <v-list-tile-content class="align-end subheading" >
            {{ magnetometer.y }} m/s
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile>
          <v-list-tile-content class="subheading">Z</v-list-tile-content>
          <v-list-tile-content class="align-end subheading" >
            {{ magnetometer.z }} m/s
          </v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-card>

  </div>
</template>

<script>

export default {
  name: 'InertialMeasurementUnit',
  mounted () {
    setInterval(this.updateData, 100)
  },
  data () {
    return {
      accelerometer: {
        x: 0,
        y: 0,
        z: 0
      },
      gyroscope: {
        x: 0,
        y: 0,
        z: 0
      },
      magnetometer: {
        x: 0,
        y: 0,
        z: 0
      }
    }
  },
  methods: {
    updateData () {
      this.accelerometer.x = 13.04
      this.accelerometer.y = 75.38
      this.accelerometer.z = 96.46
      this.gyroscope.x = 41.15
      this.gyroscope.y = 8.05
      this.gyroscope.z = 40.37
      this.magnetometer.x = 26.62
      this.magnetometer.y = 0.22
      this.magnetometer.z = 73.24
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.inertial-measurement-unit {
  display: flex;
  justify-content: space-between;
}

.imu-card {
  width: 250px;
  margin: 9px;
}

</style>
