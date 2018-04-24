<template>
  <v-container>
    <v-layout>
      <v-flex>
        <h1 class="display-3">{{ title }}</h1>
        <videostream
          stream-file="test.mjpg">
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
