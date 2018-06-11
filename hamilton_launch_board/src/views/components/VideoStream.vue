<template>
    <v-container>
      <v-layout>
        <img v-bind:src="'http://' + local_url + ':8090/' + streamFile" class="img-responsive" style="width:100%">

        <!-- change localhost to the host's ip-->
      </v-layout>
    </v-container>
</template>

<script>
export default {
  name: 'VideoStream',
  data () {
    return {
      local_url: document.location.hostname
    }
  },
  props: ['streamFile']
}
</script>
