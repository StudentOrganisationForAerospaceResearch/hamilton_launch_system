import Vue from 'vue'
import Vuetify from 'vuetify'
import Router from 'vue-router'
import Overview from '@/pages/Overview'
import VideoFeed from '@/pages/VideoFeed'
import Filling from '@/pages/Filling'

Vue.use(Router)
Vue.use(Vuetify)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      redirect: '/overview'
    },
    {
      path: '/overview',
      name: 'Overview',
      component: Overview
    },
    {
      path: '/video-feed',
      name: 'VideoFeed',
      component: VideoFeed
    },
    {
      path: '/filling',
      name: 'Filling',
      component: Filling
    }
  ]
})
