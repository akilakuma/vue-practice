<template>
   
<div>
  <p>This is room page</p>
  <ul>
    
    <li v-for="item in items" :key="item.id" class="room-block">{{ item.ID }} {{ item.RoomName }} {{ item.Price }} {{ item.Size }} {{ item.Floor }}</li>
  </ul>
</div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Room',
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
      axios.get('/room-list')
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
  .room-block{
    margin: 20px auto;
    padding: 10px;
    background: #ddd;
  }

</style>

