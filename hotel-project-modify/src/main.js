import Vue from 'vue'
import { createApp } from 'vue';
import App from './App.vue'
import router from './router'
import { BootstrapVue } from 'bootstrap-vue'
import Vuelidate from 'vuelidate'



import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

// Custom Css
import "@/assets/scss/main.scss"

Vue.use(BootstrapVue)
Vue.use(Vuelidate)
Vue.use(router);

Vue.config.productionTip = false
createApp(App).mount('#app');
