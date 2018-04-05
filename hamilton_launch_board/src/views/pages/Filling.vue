<template>
  <v-jumbotron>
    <v-container fill-height>
      <h1 class="display-3">{{ msg }}</h1>
    </v-container>
  </v-jumbotron>
</template>

<script>
export default {
  name: 'Filling',
  data () {
    return {
      msg: 'Filling'
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
