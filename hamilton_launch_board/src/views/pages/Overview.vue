<template>
  <v-container>
    <v-layout>
      <v-flex>
        <weather></weather>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>

import Weather from '@/views/components/Weather'

export default {
  name: 'Overview',
  components: {
    Weather
  },
  data: function () {
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
