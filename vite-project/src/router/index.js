import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import RoomView from '../views/RoomView.vue'
import RestaurantView from '../views/RestaurantView.vue'
import MemberView from '../views/MemberView.vue'
import AboutView from '../views/AboutView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/room',
      name: 'room',
      component: RoomView
    },
    {
      path: '/restaurant',
      name: 'restaurant',
      component: RestaurantView
    },
    {
      path: '/member',
      name: 'member',
      component: MemberView
    },
    {
      path: '/about',
      name: 'about',
      component: AboutView
    }
  ]
})

export default router
