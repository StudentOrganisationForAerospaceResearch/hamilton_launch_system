<template>
  <div class="avionics">
    <h1 class="display-1">Avionics Reporting</h1>
    <div class="avionics-content">

      <v-card class="avionics-card" raised>
        <v-card-title primary-title>
          <h4 class="headline">External Conditions</h4>
        </v-card-title>
        <v-divider></v-divider>
        <v-list dense>
          <v-list-tile>
            <v-list-tile-content class="subheading">Pressure</v-list-tile-content>
            <v-list-tile-content class="align-end subheading">{{ externalPressure }} Pa</v-list-tile-content>
          </v-list-tile>
          <v-list-tile>
            <v-list-tile-content class="subheading">Temperature</v-list-tile-content>
            <v-list-tile-content class="align-end subheading">{{ externalTemperature }} °C</v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-card>

      <v-card class="avionics-card" raised>
        <v-card-title primary-title>
          <h4 class="headline">Flight Board</h4>
        </v-card-title>
        <v-divider></v-divider>
        <v-list dense>
          <v-list-tile>
            <v-list-tile-content class="subheading">Flight Phase</v-list-tile-content>
            <v-list-tile-content class="align-end subheading">{{ flightPhase }}</v-list-tile-content>
          </v-list-tile>
          <v-list-tile>
            <v-list-tile-content class="subheading">Temperature</v-list-tile-content>
            <v-list-tile-content class="align-end subheading">{{ boardTemperature }} °C</v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-card>

      <v-card class="avionics-card" raised>
        <v-card-title primary-title>
          <h4 class="headline">GPS</h4>
        </v-card-title>
        <v-divider></v-divider>
        <v-list dense>
          <v-list-tile>
            <v-list-tile-content class="subheading">Time</v-list-tile-content>
            <v-list-tile-content class="align-end subheading">{{ gps.time }}</v-list-tile-content>
          </v-list-tile>
          <v-list-tile>
            <v-list-tile-content class="subheading">Altitude</v-list-tile-content>
            <v-list-tile-content class="align-end subheading">{{ gps.altitude }} m</v-list-tile-content>
          </v-list-tile>
          <v-list-tile>
            <v-list-tile-content class="subheading">Latitude</v-list-tile-content>
            <v-list-tile-content class="align-end subheading">{{ gps.latitude }}°</v-list-tile-content>
          </v-list-tile>
          <v-list-tile>
            <v-list-tile-content class="subheading">Longitude</v-list-tile-content>
            <v-list-tile-content class="align-end subheading">{{ gps.longitude }}°</v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-card>

      <inertial-measurement-unit></inertial-measurement-unit>
    </div>
  </div>
</template>

<script>
import InertialMeasurementUnit from '@/views/components/InertialMeasurementUnit'

export default {
  name: 'Avionics',
  components: {
    InertialMeasurementUnit
  },
  mounted () {
    setInterval(this.updateData, 100)
  },
  data () {
    return {
      externalPressure: 0,
      externalTemperature: 0,
      flightPhase: '',
      boardTemperature: 0,
      gps: {
        time: '',
        altitude: 0,
        latitude: 0,
        longitude: 0
      }
    }
  },
  methods: {
    updateData () {
      this.externalPressure = 100
      this.externalTemperature = 100
      this.flightPhase = 100
      this.boardTemperature = 100
      this.gps.time = 100
      this.gps.altitude = 100
      this.gps.latitude = 100
      this.gps.longitude = 100
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.avionics {
  max-width: 1500px;
  margin: auto;
  margin-top: 1em;
}

.avionics h1 {
  margin: 9px;
}

.avionics-content {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
}

.avionics-card {
  width: 30em;
  margin: 9px;
}

.avionics-card-content {
  display: flex;
  justify-content: center;
}

</style>
