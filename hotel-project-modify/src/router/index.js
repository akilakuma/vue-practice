// 1. Define route components.
// These can be imported from other files
// 引入Vue和Vue Router
import { createRouter, createWebHistory } from 'vue-router';

// 使用Vue Router插件


// 2. Define some routes
// Each route should map to a component.
// We'll talk about nested routes later.
const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'HomeView',
      component: () => import('@/views/HomeView.vue')
    },
    {
      path: '/hotel/:id?',
      name: 'HotelView',
      component: () => import('@/views/HotelView.vue')
    },
    {
      path: '/reservation',
      name: 'ReservationView',
      component: () => import('@/views/ReservationView.vue')
    
    }
  ]
});


export default router;