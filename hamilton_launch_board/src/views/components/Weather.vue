<template>
  <div class="weather">
    <h1 class="display-1">Weather</h1>
    <div class="weather-content">
      <v-card class="weather-card" raised>
        <v-card-title primary-title>
          <h3 class="headline">Air Temperature</h3>
        </v-card-title>
        <!-- <v-icon>mdi-thermometer-lines</v-icon> -->
        <v-card-text class="weather-card-content display-1">
          <thermometer-lines-icon class="weather-icon" />
          <p>{{ airTemperature }} °C</p>
        </v-card-text>
      </v-card>
      <v-card class="weather-card" raised>
        <v-card-title primary-title>
          <h3 class="headline">Wind Speed</h3>
        </v-card-title>
        <!-- <v-icon>mdi-weather-windy</v-icon> -->
        <v-card-text class="weather-card-content display-1">
          <weather-windy-icon class="weather-icon" />
          <p>{{ windSpeed }} km/h</p>
        </v-card-text>
      </v-card>
      <v-card class="weather-card" raised>
        <v-card-title primary-title>
          <h3 class="headline">Wind Direction</h3>
        </v-card-title>
        <!-- <v-icon>mdi-sign-direction</v-icon> -->
        <v-card-text class="weather-card-content display-1">
          <sign-direction-icon class="weather-icon" />
          <p>{{ windDirection }}°</p>
        </v-card-text>
      </v-card>
      <v-card class="weather-card" raised>
        <v-card-title primary-title>
          <h3 class="headline">Relative Humidity</h3>
        </v-card-title>
        <!-- <v-icon>mdi-water</v-icon> -->
        <v-card-text class="weather-card-content display-1">
          <water-icon class="weather-icon" />
          <p>{{ relativeHumidity }}%</p>
        </v-card-text>
      </v-card>
    </div>
  </div>
</template>

<script>
import ThermometerLinesIcon from 'vue-material-design-icons/thermometer-lines.vue'
import WeatherWindyIcon from 'vue-material-design-icons/weather-windy.vue'
import SignDirectionIcon from 'vue-material-design-icons/sign-direction.vue'
import WaterIcon from 'vue-material-design-icons/water.vue'

export default {
  name: 'Weather',
  components: {
    ThermometerLinesIcon,
    WeatherWindyIcon,
    SignDirectionIcon,
    WaterIcon
  },
  computed: {
    airTemperature () {
      return 15.2
    },
    windSpeed () {
      return 22
    },
    windDirection () {
      return 264
    },
    relativeHumidity () {
      return 78
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.weather {
  max-width: 1200px;
  margin: auto;
  margin-top: 2em;
}

.weather h1 {
  margin: 9px;
}

.weather-content {
  display: flex;
  flex-direction: row;
  justify-content: center;
  flex-wrap: wrap;
}

.weather-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 250px;
  margin: 9px;
}

.weather-card-content {
  display: flex;
  justify-content: center;
  padding-bottom: 1em;
}

.weather-card-content p {
  margin-bottom: 0;
  margin-left: 9px;
  padding-top: 9px;
}

</style>
