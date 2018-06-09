<template>
    <v-container fluid grid-list-md>
        <h1 class="display-3">{{ title }}</h1>
        <v-layout row wrap>
            <v-flex lg6>
                <videostream stream-file="high_quality.mjpg">
                </videostream>
            </v-flex>
            <v-flex lg6>
                <v-layout row wrap>
                    <v-flex md7>
                        <videostream
                          stream-file="low_quality1.mjpg">
                        </videostream>
                    </v-flex>
                    <v-flex md7>
                        <videostream
                          stream-file="low_quality2.mjpg">
                        </videostream>
                    </v-flex>
                </v-layout>
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
