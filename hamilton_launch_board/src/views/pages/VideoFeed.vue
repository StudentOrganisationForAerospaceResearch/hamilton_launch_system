<template>
<div class="video-feed">
  <h1 class="display-3">{{ title }}</h1>
  <videostream
    class="high-quality"
    stream-file="high_quality.mjpg">
  </videostream>
  <div class="low-quality-streams" v-bind:class="{ column: isMobile() }">
    <videostream
      v-bind:class="{ full_width: isMobile() }"
      class="low-quality"
      stream-file="low_quality1.mjpg">
    </videostream>
    <videostream
      v-bind:class="{ full_width: isMobile() }"
      class="low-quality"
      stream-file="low_quality2.mjpg">
    </videostream>
  </div>
</div>
</template>

<script>
import Videostream from '@/views/components/VideoStream'

export default {
  name: 'VideoFeed',
  components: {
    Videostream
  },
  methods: {
    isMobile: function () {
      return /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent)
    }
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
.video-feed {
  width: 80%;
  min-width: 40em;
  margin: auto;
  margin-top: 1em;
}

.high-quality {
  width: 100%;
}

.low-quality-streams {
  display: flex;
}

.column {
  flex-direction: column;
}

.low-quality {
  flex-wrap: wrap;
  width: 40%;
}

.full_width {
  width: 100%;
}
</style>
