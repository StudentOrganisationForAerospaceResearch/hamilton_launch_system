<template>
  <v-container>
    <v-layout>
      <v-flex>
        <div class="weather">
          <h1 class="display-2">Weather</h1>
          <div class="weather-content">
            <v-card class="weather-card" raised>
              <v-card-title primary-title>
                <h3 class="headline mb-0">Air Temperature</h3>
              </v-card-title>
              <!-- <v-icon>mdi-thermometer-lines</v-icon> -->
              <v-card-text class="weather-card-content display-1">
                <thermometer-lines class="weather-icon" />
                <p>{{ airTemperature }}°C</p>
              </v-card-text>
            </v-card>
            <v-card class="weather-card" raised>
              <v-card-title primary-title>
                <h3 class="headline mb-0">Wind Speed</h3>
              </v-card-title>
              <!-- <v-icon>mdi-weather-windy</v-icon> -->
              <v-card-text class="weather-card-content display-1">
                <weather-windy class="weather-icon" />
                <p>{{ windSpeed }}km/h</p>
              </v-card-text>
            </v-card>
            <v-card class="weather-card" raised>
              <v-card-title primary-title>
                <h3 class="headline mb-0">Wind Direction</h3>
              </v-card-title>
              <!-- <v-icon>mdi-sign-direction</v-icon> -->
              <v-card-text class="weather-card-content display-1">
                <sign-direction class="weather-icon" />
                <p>{{ windDirection }}°</p>
              </v-card-text>
            </v-card>
            <v-card class="weather-card" raised>
              <v-card-title primary-title>
                <h3 class="headline mb-0">Relative Humidity</h3>
              </v-card-title>
              <!-- <v-icon>mdi-water</v-icon> -->
              <v-card-text class="weather-card-content display-1">
                <water class="weather-icon" />
                <p>{{ relativeHumidity }}%</p>
              </v-card-text>
            </v-card>
          </div>
        </div>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  name: 'Overview',
  data: function() {
    return {
      airTemperature: 15.2,
      windSpeed: 14,
      windDirection: 264,
      relativeHumidity: 78
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.weather {
  max-width: 1200px;
  margin: auto;
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
