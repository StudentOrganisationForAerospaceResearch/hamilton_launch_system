<template>
  <div class="filling">
    <div class="filling-content">
      <v-card class="filling-card" raised>

        <v-card-title primary-title>
          <h4 class="headline">Filling Progress</h4>
        </v-card-title>
        <v-divider></v-divider>
        <div class="filling-card-content filling-card-content-circular">
          <v-progress-circular
            :size="250"
            :width="40"
            :rotate="270"
            :value="fillProgress"
          >
            <p class="display-2">{{ fillProgress.toFixed(1) }}%</p>
          </v-progress-circular>
          <v-list dense>
            <v-list-tile class="mass-item">
              <v-list-tile-content class="subheading">Total Mass</v-list-tile-content>
              <v-list-tile-content class="align-end subheading mass">{{ totalMass }} kg</v-list-tile-content>
            </v-list-tile>
            <v-list-tile class="mass-item">
              <v-list-tile-content class="subheading">Rocket Mass</v-list-tile-content>
              <v-list-tile-content class="align-end subheading mass">{{ ROCKET_MASS }} kg</v-list-tile-content>
            </v-list-tile>
            <v-list-tile class="mass-item">
              <v-list-tile-content class="subheading">Oxidizer Mass</v-list-tile-content>
              <v-list-tile-content class="align-end subheading mass">{{ oxidizerMass }} kg</v-list-tile-content>
            </v-list-tile>
            <v-list-tile class="mass-item">
              <v-list-tile-content class="subheading">Oxidizer Target</v-list-tile-content>
              <v-list-tile-content class="align-end subheading mass">{{ TARGET_OXIDIZER_MASS }} kg</v-list-tile-content>
            </v-list-tile>
          </v-list>
          <div class="fill-status">
            <h1 class="headline">Vent Valve Status</h1>
            <div class="display-1 status-content" v-show="ventValveOpen">
              <h2 class="display-1">OPEN</h2>
              <plus-network-icon />
            </div>
            <div class="display-1 status-content closed-valve" v-show="!ventValveOpen">
              <h2 class=" display-1">CLOSED</h2>
              <close-network-icon/>
            </div>
            <h1 class="headline">Fill Valve Status</h1>
            <div class="display-1 status-content" v-show="fillValveOpen">
              <h2 class="display-1">OPEN</h2>
              <plus-network-icon />
            </div>
            <div class="display-1 status-content closed-valve" v-show="!fillValveOpen">
              <h2 class=" display-1">CLOSED</h2>
              <close-network-icon/>
            </div>
            <v-text-field
              name="input-1"
              label="Fill Control Code"
              color="yellow"
              v-model="fillControlCode"
            ></v-text-field>
            <div>
              <v-btn large @mouseup="sendOpenFillValveCommand">
                OPEN
              </v-btn>
              <v-btn large @mouseup="sendCloseFillValveCommand">
                CLOSE
              </v-btn>
            </div>
          </div>
        </div>
      </v-card>

      <v-card class="filling-card" raised>
        <v-card-title primary-title>
          <h4 class="headline">Oxidizer Tank Pressure</h4>
        </v-card-title>
        <v-divider></v-divider>
        <div class="filling-card-content filling-card-content-bar">
          <div class="conditions-label">
            <p class="title">{{ pressure }} kPa</p>
            <p class="title max">MAX {{ MAX_PRESSURE }} kPa</p>
          </div>
          <v-progress-linear
            :value="pressurePercentage"
            height="20"
            color="white">
          </v-progress-linear>
        </div>
      </v-card>

    </div>
  </div>
</template>

<script>
import PlusNetworkIcon from 'vue-material-design-icons/plus-network.vue'
import CloseNetworkIcon from 'vue-material-design-icons/close-network.vue'
import { mapState } from 'vuex'

export default {
  name: 'Filling',
  components: {
    PlusNetworkIcon,
    CloseNetworkIcon
  },
  methods: {
    sendOpenFillValveCommand: function (event) {
      this.$socket.sendObj({
        type: 'fillControl',
        command: 'openFillValve',
        code: this.fillControlCode
      })
    },
    sendCloseFillValveCommand: function (event) {
      this.$socket.sendObj({
        type: 'fillControl',
        command: 'closeFillValve',
        code: this.fillControlCode
      })
    }
  },
  data: function () {
    return {
      ROCKET_MASS: 80.12,
      TARGET_OXIDIZER_MASS: 21.2,
      MAX_PRESSURE: 50.2,
      MAX_TEMPERATURE: 32.2,
      fillControlCode: ''
    }
  },
  computed: {
    ...mapState({
      totalMass: state => state.fillingInfo.totalMass,
      ventValveOpen: state => state.fillingInfo.ventValveOpen,
      fillValveOpen: state => state.fillingInfo.fillValveOpen,
      pressure: state => state.oxidizerTankPressure.pressure
    }),
    oxidizerMass: function () {
      if (this.totalMass === '-') {
        return '-'
      } else if (this.totalMass < this.ROCKET_MASS) {
        return 0
      }
      return (this.totalMass - this.ROCKET_MASS).toFixed(2)
    },
    fillProgress: function () {
      if (this.totalMass === '-') {
        return 0
      } else if (this.totalMass < this.ROCKET_MASS) {
        return 0
      }
      return (this.oxidizerMass / this.TARGET_OXIDIZER_MASS) * 100
    },
    pressurePercentage: function () {
      if (this.pressure === '-') {
        return 0
      }
      return (this.pressure / this.MAX_PRESSURE) * 100
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.filling {
  max-width: 1500px;
  margin: auto;
  margin-top: 1em;
}

.filling-content {
  display: flex;
  align-items: center;
  flex-direction: column;
}

.filling-card {
  margin: 9px;
  max-width: 90vw;
  width: 1000px;
}

.filling-card-content {
  display: flex;
  align-items: center;
  padding: 2rem;
}

.filling-card-content-circular {
  justify-content: space-around;
  flex-wrap: wrap;
}

.mass-item {
  margin: 1rem;
}

.mass {
  margin-left: 2rem;
}

.filling-card-content-bar {
  flex-direction: column;
}

.conditions-label {
  width: 100%;
  display: flex;
  justify-content: space-between;
}

.max {
  color: #757575;
}

.fill-status {
  text-align-last: center;
}

.fill-status * {
  margin-bottom: 1rem;
}

.status-content {
  display: flex;
  justify-content: space-between;
}

.closed-valve {
  color: #757575;
}
</style>
