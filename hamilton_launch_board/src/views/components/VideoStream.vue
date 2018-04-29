<template>
    <v-container>
      <v-layout>
        <img :src="'http://localhost:8090/' + streamFile" class="img-responsive" style="width:100%">
      </v-layout>
    </v-container>
</template>

<script>
export default {
  name: 'VideoStream',
  props: ['streamFile']
}
</script>
