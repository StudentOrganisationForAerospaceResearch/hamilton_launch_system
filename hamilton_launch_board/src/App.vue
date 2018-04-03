<template>
  <v-app dark>
    <v-toolbar>
      <v-toolbar-title>Hamilton Launch Board</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-toolbar-title>Student Organization for Aerospace Research</v-toolbar-title>
    </v-toolbar>
    <v-tabs grow fixed-tabs height=70 slider-color="yellow">
      <v-tab ripple :to="{name:'Overview'}">
        Overview
      </v-tab>
      <v-tab ripple :to="{name:'VideoFeed'}">
        Video Feed
      </v-tab>
      <v-tab ripple :to="{name:'Filling'}">
        Filling
      </v-tab>
    </v-tabs>
    <router-view></router-view>
  </v-app>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
</style>
