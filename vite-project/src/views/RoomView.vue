<template>
   
<div class="w-max-2xl block">
  <ul class="flex mt-5 ml-96 mr-96 flex-wrap justify-center items-center" >
	<li class="mx-6 mt-5" v-for="item in items" :key="item.id" >
		<div class="
            py-8
            px-8
            max-w-sm
            mx-auto
            bg-white
            rounded-xl
            shadow-md
            space-y-2
            sm:py-4
            sm:flex
            sm:items-center
            sm:space-y-0 sm:space-x-6
          "> 
			<img class="block mx-auto h-24 rounded-full sm:mx-0 sm:flex-shrink-0" src="https://picsum.photos/300/300?ranodm=1" alt="Woman's Face" />
			<div class="text-center space-y-2 sm:text-left">
				<div class="space-y-0.5">
					<p class="text-lg text-black font-semibold">{{ item.ID }} {{ item.RoomName }} </p>
					<p class="text-gray-500 font-medium">{{ item.Price }} {{ item.Size }} {{ item.Floor }}</p>
				</div>
				<button class="
                px-4
                py-1
                text-sm text-purple-600
                font-semibold
                rounded-full
                border border-purple-200
                hover:text-white
                hover:bg-purple-600
                hover:border-transparent
                focus:outline-none
                focus:ring-2 focus:ring-purple-600 focus:ring-offset-2
              " @click="startbooking">
					訂房
				</button>
			</div>
		</div>
	</li>
  </ul>
  <BookingForm v-show="showbookingform"/>
  
</div>
</template>

<script>
import axios from 'axios';
import BookingForm from '../components/BookingForm.vue';

export default {
  components: { BookingForm },
  name: 'Room',
  // 把api拿到的資料，放進items
  data() {
    return {
      items: [],
      showbookingform : false
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
    },
    startbooking() {
      this.showbookingform = !this.showbookingform
    }

  }
};
</script>

<style>


</style>

