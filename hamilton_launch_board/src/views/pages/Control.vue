<template>
  <div class="control">
    <div class="control-content">

      <v-card class="control-card" raised>
        <v-card-title primary-title>
          <h3 class="headline">Arm</h3>
        </v-card-title>
        <v-divider></v-divider>
        <div class="control-card-content">
          <div v-show="!isMobile()">
            <v-text-field
              name="input-1"
              label="Arm Code"
              color="yellow"
              v-model="armCode"
            ></v-text-field>
            <v-btn large
              @mousedown="sendArmCommand"
              @mouseup="stopArmCommand"
              @mouseleave="stopArmCommand">ARM</v-btn>
            <v-btn large v-show="armCounter >= 100 "
              @mouseup="retryArmCommand">RETRY</v-btn>
          </div>
          <div class="subprogress-section">
            <div class="subprogress">
              <v-progress-circular
                :size="50"
                :width="12"
                :rotate="270"
                :value="softwareArmCounter"
              ></v-progress-circular>
              <h4 class="subheading">Software</h4>
            </div>
            <div class="subprogress">
              <v-progress-circular
                :size="50"
                :width="12"
                :rotate="270"
                :value="launchSystemsArmCounter"
              ></v-progress-circular>
              <h4 class="subheading">Launch Systems</h4>
            </div>
            <div class="subprogress">
              <v-progress-circular
                :size="50"
                :width="12"
                :rotate="270"
                :value="vpRocketsArmCounter"
              ></v-progress-circular>
              <h4 class="subheading">VP Rockets</h4>
            </div>
          </div>
          <div>
            <v-progress-circular
              :size="250"
              :width="40"
              :rotate="270"
              :value="armCounter"
            >
            <h2
              class="title"
              v-bind:class="{ disabled: armCounter < 100 }">
              ARMED
            </h2>
            </v-progress-circular>
          </div>
        </div>
      </v-card>

      <v-card class="control-card" raised>
        <v-card-title primary-title>
          <h3 class="headline">Launch</h3>
        </v-card-title>
        <v-divider></v-divider>
        <div class="control-card-content">
          <div v-show="!isMobile()">
            <v-text-field
              name="input-1"
              label="Launch Code"
              color="yellow"
              v-model="launchCode"
            ></v-text-field>
            <v-btn large
              @mousedown="sendLaunchCommand"
              @mouseup="stopLaunchCommand"
              @mouseleave="stopLaunchCommand">LAUNCH</v-btn>
          </div>
          <div class="subprogress-section">
            <div class="subprogress">
              <v-progress-circular
                :size="50"
                :width="12"
                :rotate="270"
                :value="softwareLaunchCounter"
              ></v-progress-circular>
              <h4 class="subheading">Software</h4>
            </div>
            <div class="subprogress">
              <v-progress-circular
                :size="50"
                :width="12"
                :rotate="270"
                :value="launchSystemsLaunchCounter"
              ></v-progress-circular>
              <h4 class="subheading">Launch Systems</h4>
            </div>
            <div class="subprogress">
              <v-progress-circular
                :size="50"
                :width="12"
                :rotate="270"
                :value="vpRocketsLaunchCounter"
              ></v-progress-circular>
              <h4 class="subheading">VP Rockets</h4>
            </div>
          </div>
          <div>
            <v-progress-circular
              :size="250"
              :width="40"
              :rotate="270"
              :value="launchCounter"
            >
            <h2
              class="display-4"
              v-bind:class="{ disabled: launchCounter < 100 }">
              {{ countdown }}
            </h2>
            </v-progress-circular>
          </div>
        </div>
      </v-card>

      <v-card class="control-card" raised>
        <v-card-title primary-title>
          <h3 class="headline">Abort</h3>
        </v-card-title>
        <v-divider></v-divider>
        <div class="control-card-content">
          <div>
            <v-text-field
              name="input-1"
              label="Abort Code"
              color="red"
              v-model="abortCode"
            ></v-text-field>
            <v-btn large
              color="red"
              v-on:click="sendAbortCommand">ABORT</v-btn>
          </div>
        </div>
      </v-card>

    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'Control',
  data: function () {
    return {
      armCode: '',
      launchCode: '',
      abortCode: ''
    }
  },
  computed: {
    ...mapState({
      softwareArmCounter: state => state.launchControlInfo.softwareArmCounter,
      launchSystemsArmCounter: state => state.launchControlInfo.launchSystemsArmCounter,
      vpRocketsArmCounter: state => state.launchControlInfo.vpRocketsArmCounter,
      armCounter: state => state.launchControlInfo.armCounter,
      softwareLaunchCounter: state => state.launchControlInfo.softwareLaunchCounter,
      launchSystemsLaunchCounter: state => state.launchControlInfo.launchSystemsLaunchCounter,
      vpRocketsLaunchCounter: state => state.launchControlInfo.vpRocketsLaunchCounter,
      launchCounter: state => state.launchControlInfo.launchCounter,
      countdown: state => state.launchControlInfo.countdown
    })
  },
  methods: {
    sendArmCommand: function (event) {
      this.$socket.sendObj({
        type: 'launchControl',
        command: 'arm',
        code: this.armCode
      })
    },
    stopArmCommand: function (event) {
      this.$socket.sendObj({
        type: 'launchControl',
        command: 'stopArm',
        code: this.armCode
      })
    },
    retryArmCommand: function (event) {
      this.$socket.sendObj({
        type: 'launchControl',
        command: 'retry',
        code: this.armCode
      })
    },
    sendLaunchCommand: function (event) {
      this.$socket.sendObj({
        type: 'launchControl',
        command: 'launch',
        code: this.launchCode
      })
    },
    stopLaunchCommand: function (event) {
      this.$socket.sendObj({
        type: 'launchControl',
        command: 'stopLaunch',
        code: this.launchCode
      })
    },
    sendAbortCommand: function (event) {
      this.$socket.sendObj({
        type: 'launchControl',
        command: 'abort',
        code: this.abortCode
      })
    },
    isMobile: function () {
      return /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.control {
  max-width: 1500px;
  margin: auto;
  margin-top: 1em;
}

.control-content {
  display: flex;
  align-items: center;
  flex-direction: column;
}

.control-card {
  margin: 9px;
  max-width: 90vw;
  width: 1000px;
}

.control-card-content {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
  align-items: center;
  padding: 2rem;
}

.subprogress-section {
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: left;
}

.subprogress {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 1rem;
}

.subprogress h4 {
  margin-left: 9px;
}

.disabled {
  color: #383838;
}
</style>
