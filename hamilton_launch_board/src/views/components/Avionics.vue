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
            <v-list-tile-content class="align-end subheading"> {{ externalPressure }} Pa</v-list-tile-content>
          </v-list-tile>
          <v-list-tile>
            <v-list-tile-content class="subheading">Temperature</v-list-tile-content>
            <v-list-tile-content class="align-end subheading"> {{ externalTemperature }} °C</v-list-tile-content>
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
            <v-list-tile-content class="align-end subheading"> {{ flightPhase }}</v-list-tile-content>
          </v-list-tile>
          <v-list-tile>
            <v-list-tile-content class="subheading">Temperature</v-list-tile-content>
            <v-list-tile-content class="align-end subheading"> {{ boardTemperature }} °C</v-list-tile-content>
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
            <v-list-tile-content class="align-end subheading"> {{ gps.time }}</v-list-tile-content>
          </v-list-tile>
          <v-list-tile>
            <v-list-tile-content class="subheading">Altitude</v-list-tile-content>
            <v-list-tile-content class="align-end subheading"> {{ gps.altitude }} m</v-list-tile-content>
          </v-list-tile>
          <v-list-tile>
            <v-list-tile-content class="subheading">Latitude</v-list-tile-content>
            <v-list-tile-content class="align-end subheading"> {{ gps.latitude }}°</v-list-tile-content>
          </v-list-tile>
          <v-list-tile>
            <v-list-tile-content class="subheading">Longitude</v-list-tile-content>
            <v-list-tile-content class="align-end subheading"> {{ gps.longitude }}°</v-list-tile-content>
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
  computed: {
    externalPressure () {
      return 15.2
    },
    externalTemperature () {
      return 22
    },
    flightPhase () {
      return 'PRELAUNCH'
    },
    boardTemperature () {
      return 264
    },
    gps () {
      return {
        time: 'Thu, 05 Apr 2018 08:27:13 GMT',
        altitude: 2345,
        latitude: 936.42,
        longitude: 1249.2
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.avionics {
  max-width: 1500px;
  margin: auto;
  margin-top: 2em;
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
