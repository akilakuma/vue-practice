<script setup>
import { RouterLink, RouterView } from 'vue-router'
// import Vue from 'vue'
import 'bootstrap'
import 'bootstrap/dist/css/bootstrap.min.css'
</script>


<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-light">
    <div class="container">
      <a class="navbar-brand" href="#">敦謙國際</a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav"
        aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNav">
        <ul class="navbar-nav">
          <li class="nav-item">
            <RouterLink to="/" class="nav-link">首頁</RouterLink>
        
          </li>
          <li class="nav-item">
            <RouterLink to="/booking" class="nav-link">訂房</RouterLink>
          </li>
          <li class="nav-item">
            <nav>
            <RouterLink to="/resturant" class="nav-link">餐廳</RouterLink>
            </nav>
          </li>
          <li class="nav-item">
            <RouterLink to="/activity" class="nav-link">活動</RouterLink>
          </li>
          <li class="nav-item">
            <RouterLink to="/about" class="nav-link">關於我們</RouterLink>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<style scoped>
nav RouterLink {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}
</style>