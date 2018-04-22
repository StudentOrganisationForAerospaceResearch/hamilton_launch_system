<template>
  <v-app dark>
    <v-toolbar>
      <v-toolbar-title>Hamilton Launch Board</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-spacer></v-spacer>
      <v-toolbar-title >
        <lan-connect-icon v-show="connected" />
        <lan-disconnect-icon v-show="!connected" class="disconnected" />
      </v-toolbar-title>
    </v-toolbar>
    <v-tabs fixed-tabs height=70 slider-color="yellow">
      <v-tab ripple :to="{name:'Overview'}">
        Overview
      </v-tab>
      <v-tab ripple :to="{name:'VideoFeed'}">
        Video Feed
      </v-tab>
      <v-tab ripple :to="{name:'Filling'}">
        Filling
      </v-tab>
      <v-tab ripple :to="{name:'Control'}">
        Control
      </v-tab>
    </v-tabs>
    <router-view></router-view>
  </v-app>
</template>

<script>
import LanConnectIcon from 'vue-material-design-icons/lan-connect.vue'
import LanDisconnectIcon from 'vue-material-design-icons/lan-disconnect.vue'

export default {
  name: 'App',
  components: {
    LanConnectIcon,
    LanDisconnectIcon
  },
  computed: {
    connected: function () {
      return this.$store.state.socket.isConnected
    }
  }
}
</script>

<style scoped>
.disconnected {
  color: red;
}
</style>
