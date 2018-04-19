<template>
  <div class="control">
    <div class="control-content">

      <v-card class="control-card" raised>
        <v-stepper :value="launchProgress" alt-labels>
          <v-stepper-header color="yellow">
            <v-stepper-step step="1" complete>Prelaunch</v-stepper-step>
            <v-divider></v-divider>
            <v-stepper-step step="2">Armed</v-stepper-step>
            <v-divider></v-divider>
            <v-stepper-step step="3">Launched</v-stepper-step>
          </v-stepper-header>
        </v-stepper>
      </v-card>

      <v-card class="control-card" raised>
        <v-card-title primary-title>
          <h3 class="headline">Arm</h3>
        </v-card-title>
        <v-divider></v-divider>
        <div class="control-card-content">
          <div>
            <v-text-field
              name="input-1"
              label="Arm Code"
              color="yellow"
            ></v-text-field>
            <v-btn>ARM</v-btn>
          </div>
          <div class="subprogress-section">
            <div class="subprogress">
              <v-progress-circular
                :size="50"
                :width="12"
                :rotate="270"
                :value="armCounter"
              ></v-progress-circular>
              <h4 class="subheading">Software</h4>
            </div>
            <div class="subprogress">
              <v-progress-circular
                :size="50"
                :width="12"
                :rotate="270"
                :value="armCounter"
              ></v-progress-circular>
              <h4 class="subheading">Launch Systems</h4>
            </div>
            <div class="subprogress">
              <v-progress-circular
                :size="50"
                :width="12"
                :rotate="270"
                :value="armCounter"
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
            <h2 class="title disabled">ARMED</h2>
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
          <div>
            <v-text-field
              name="input-1"
              label="Launch Code"
              color="yellow"
            ></v-text-field>
            <v-btn>LAUNCH</v-btn>
          </div>
          <div class="subprogress-section">
            <div class="subprogress">
              <v-progress-circular
                :size="50"
                :width="12"
                :rotate="270"
                :value="armCounter"
              ></v-progress-circular>
              <h4 class="subheading">Software</h4>
            </div>
            <div class="subprogress">
              <v-progress-circular
                :size="50"
                :width="12"
                :rotate="270"
                :value="armCounter"
              ></v-progress-circular>
              <h4 class="subheading">Launch Systems</h4>
            </div>
            <div class="subprogress">
              <v-progress-circular
                :size="50"
                :width="12"
                :rotate="270"
                :value="armCounter"
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
            <h2 class="display-4 disabled">{{ countdown }}</h2>
            </v-progress-circular>
          </div>
        </div>
      </v-card>

    </div>
  </div>
</template>

<script>

export default {
  name: 'Control',
  data () {
    return {
      launchProgress: 2,
      armCounter: 20,
      countdown: 10
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
