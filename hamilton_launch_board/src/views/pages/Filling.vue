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
            :rotate="360"
            :value="progress"
          >
            <p class="display-2">{{ progress }}%</p>
          </v-progress-circular>
          <v-list dense>
            <v-list-tile class="mass-item">
              <v-list-tile-content class="subheading">Total Mass</v-list-tile-content>
              <v-list-tile-content class="align-end subheading mass">{{ totalMass }} kg</v-list-tile-content>
            </v-list-tile>
            <v-list-tile class="mass-item">
              <v-list-tile-content class="subheading">Rocket Mass</v-list-tile-content>
              <v-list-tile-content class="align-end subheading mass">{{ rocketMass }} kg</v-list-tile-content>
            </v-list-tile>
            <v-list-tile class="mass-item">
              <v-list-tile-content class="subheading">Oxidizer Mass</v-list-tile-content>
              <v-list-tile-content class="align-end subheading mass">{{ oxidizerMass }} kg</v-list-tile-content>
            </v-list-tile>
            <v-list-tile class="mass-item">
              <v-list-tile-content class="subheading">Oxidizer Target</v-list-tile-content>
              <v-list-tile-content class="align-end subheading mass">{{ targetOxidizerMass }} kg</v-list-tile-content>
            </v-list-tile>
          </v-list>
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
            <p class="title max">MAX 50 kPa</p>
          </div>
          <v-progress-linear
            value="75"
            height="20"
            color="white">
          </v-progress-linear>
        </div>
      </v-card>

      <v-card class="filling-card" raised>
        <v-card-title primary-title>
          <h4 class="headline">Oxidizer Tank Temperature</h4>
        </v-card-title>
        <v-divider></v-divider>
        <div class="filling-card-content filling-card-content-bar">
          <div class="conditions-label">
            <p class="title">{{ temperature }} °C</p>
            <p class="title max">MAX 50 °C</p>
          </div>
          <v-progress-linear
            value="75"
            height="20"
            color="white">
          </v-progress-linear>
        </div>
      </v-card>

    </div>
  </div>
</template>

<script>
export default {
  name: 'Filling',
  data () {
    return {
      progress: 20,
      pressure: 50,
      temperature: 30,
      totalMass: 300,
      rocketMass: 280,
      oxidizerMass: 20,
      targetOxidizerMass: 20
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
</style>
