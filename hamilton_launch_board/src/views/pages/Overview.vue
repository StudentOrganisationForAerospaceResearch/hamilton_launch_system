<template>
  <v-container>
    <v-layout>
      <v-flex>
        <avionics></avionics>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>

import Avionics from '@/views/components/Avionics'

export default {
  name: 'Overview',
  components: {
    Avionics
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
