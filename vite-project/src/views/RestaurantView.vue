<template>
   
<div>
  <p>This is restaurant page</p>
  <ul>
    
    <li v-for="item in items" :key="item.id" class="restaurant-block">{{ item.ID }} {{ item.Name }} {{ item.Price }}</li>
  </ul>
</div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Restaurant',
  // 把api拿到的資料，放進items
  data() {
    return {
      items: []
    };
  },
  beforeRouteEnter(to, from, next) {
    next(vm => {
      vm.fetchData();
    });
  },
  methods: {
    fetchData() {
      axios.get('/restaurant-list')
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
  .restaurant-block{
    margin: 20px auto;
    padding: 10px;
    background: lightblue;
  }
</style>

