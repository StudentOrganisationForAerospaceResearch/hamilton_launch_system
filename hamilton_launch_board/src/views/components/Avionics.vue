<template>
  <div class="avionics">
    <h1 class="display-1">Avionics Reporting</h1>
    <div class="avionics-content">

      <v-card class="avionics-card" raised>
        <v-card-title primary-title>
          <h4 class="headline">Rocket Conditions</h4>
        </v-card-title>
        <v-divider></v-divider>
        <v-list dense>
          <v-list-tile>
            <v-list-tile-content class="subheading">External Pressure</v-list-tile-content>
            <v-list-tile-content class="align-end subheading">{{ barometer.pressure }} Pa</v-list-tile-content>
          </v-list-tile>
          <v-list-tile>
            <v-list-tile-content class="subheading">External Temperature</v-list-tile-content>
            <v-list-tile-content class="align-end subheading">{{ barometer.temperature }} °C</v-list-tile-content>
          </v-list-tile>
          <v-list-tile>
            <v-list-tile-content class="subheading">Combustion Chamber Pressure</v-list-tile-content>
            <v-list-tile-content class="align-end subheading">{{ combustionChamberPressure.pressure }} Pa</v-list-tile-content>
          </v-list-tile>
          <v-list-tile>
            <v-list-tile-content class="subheading">Flight Phase</v-list-tile-content>
            <v-list-tile-content class="align-end subheading">{{ flightPhase.phase }}</v-list-tile-content>
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
            <v-list-tile-content class="align-end subheading">{{ gps.epochTimeMsec }}</v-list-tile-content>
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
import { mapState } from 'vuex'
import InertialMeasurementUnit from '@/views/components/InertialMeasurementUnit'

export default {
  name: 'Avionics',
  components: {
    InertialMeasurementUnit
  },
  computed: {
    ...mapState({
      barometer: state => state.barometer,
      combustionChamberPressure: state => state.combustionChamberPressure,
      flightPhase: state => state.flightPhase,
      gps: state => state.gps
    })
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

</style>
