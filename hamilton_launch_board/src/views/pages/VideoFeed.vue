<template>
  <v-container fluid>
    <v-layout row wrap>
      <v-flex d-flex xs12 sm6 md4>
        <h1 class="display-3">{{ title }}</h1>
        <v-card>
          <videostream
            stream-file="test1.mjpg">
          </videostream>
        </v-card>
        <videostream
          stream-file="test1.mjpg">
        </videostream>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import Videostream from '@/views/components/VideoStream'

export default {
  name: 'VideoFeed',
  components: {
    Videostream
  },
  data () {
    return {
      title: 'Video Feed'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
</style>
