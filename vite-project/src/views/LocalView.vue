<template>
    <ul>
      <li v-for="item in items" :key="item.id">{{ item.name }}</li>
    </ul>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Home',
  data() {
    return {
      items: []
    };
  },
  mounted() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      axios.get('https://api.example.com/items')
        .then(response => {
          this.items = response.data;
        })
        .catch(error => {
          console.error(error);
        });
    }
  }
};
</script>

<style>

</style>