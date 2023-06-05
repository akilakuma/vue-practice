<script setup>
import { RouterLink, RouterView } from 'vue-router'
import NavbarItem from './components/NavbarItem.vue'

</script>

<template>
  <div class="h-full">
    <NavbarItem/>
    <router-view></router-view>
    <footer>

    </footer>
  </div>

</template>


<script >

export default {
  name: 'Home',
  data() {
    return {
    };
  },

}

</script>



<style scoped>

</style>
