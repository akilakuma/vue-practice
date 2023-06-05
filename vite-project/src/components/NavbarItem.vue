<script setup>
import { RouterLink, RouterView } from 'vue-router'
// import Vue from 'vue'
import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css'
</script>


<template>

    <nav class="bg-sky-100 p-4 space-x-3 mb-8 mx-auto md:mx-auto">
      <RouterLink to="/" class="inline-block p-2 text-white bg-blue-400 rounded md:p-0 dark:text-white md:dark:text-blue-500" active-class="bg-blue-600" >小浣熊大飯店</RouterLink>
      <RouterLink to="/room" class="inline-block p-2 text-white bg-blue-400 rounded md:p-0 dark:text-white md:dark:text-blue-500" active-class="bg-blue-600">訂房</RouterLink>
      <RouterLink to="/restaurant" class="inline-block p-2 text-white bg-blue-400 rounded md:p-0 dark:text-white md:dark:text-blue-500" active-class="bg-blue-600">餐廳</RouterLink>
      <RouterLink to="/activity" class="inline-block p-2 text-white bg-blue-400 rounded md:p-0 dark:text-white md:dark:text-blue-500" active-class="bg-blue-600">活動</RouterLink>
      <RouterLink to="/about" class="inline-block p-2 text-white bg-blue-400 rounded md:p-0 dark:text-white md:dark:text-blue-500" active-class="bg-blue-600">關於我們</RouterLink>
    </nav>

    
        

</template>

<style scoped>
nav RouterLink {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}
</style>