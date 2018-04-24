<template>
  <v-jumbotron>
    <v-container fill-height>
      <h1 class="display-3">{{ msg }}</h1>
      <videostream
        message="Hello">
      </videostream>
    </v-container>
  </v-jumbotron>
</template>

<script>
import VideoStream from '@/views/components/VideoStream'

export default {
  name: 'VideoFeed',
  components: {
    VideoStream
  },
  data () {
    return {
      msg: 'Video Feed Test'
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
