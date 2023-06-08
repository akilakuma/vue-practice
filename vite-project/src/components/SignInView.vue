<template>
  <div class="flex justify-center items-center min-h-screen bg-gray-100">
    <div class="w-full max-w-md p-4 bg-white rounded shadow">
      <h2 class="text-2xl font-bold mb-4">會員登入</h2>
      
      <!-- 登入 -->
      <form @submit.prevent="login">
        <div class="mb-4">
          <label for="email" class="block mb-1">Email</label>
          <input v-model="email" id="email" type="email" class="w-full px-3 py-2 rounded border border-gray-300">
        </div>
        <div class="mb-4">
          <label for="password" class="block mb-1">密碼</label>
          <input v-model="password" id="password" type="password" class="w-full px-3 py-2 rounded border border-gray-300">
        </div>
        <button type="submit" class="w-full py-2 px-4 bg-blue-500 text-white rounded hover:bg-blue-600">登入</button>
        <div class="mt-10 mb-4 text-center">
         <p>註冊會員</p>
         <button @click="gotoRegister" class="w-1/2 py-2 px-4 bg-green-600 text-white rounded ">註冊</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import UserDataView from '../components/UserDataView.vue'

export default {
  data() {
    return {
      email: "",
      password: "",
      user: {},
    };
  },
  methods:{
    gotoRegister(status, u) {
      if (status === undefined) {
        status = "SignUp"
      }
      this.$emit('update', status, u)
    },
    async login() {
      const passwordHash = await this.hashPassword(this.password);
      axios.get('/login', {
        params: {
          email: this.email,
          password: passwordHash
        }
      })
      .then(response => {
        this.user = response.data;
        this.gotoRegister("ShowData", this.user)

      })
      .catch(error => {
        console.error(error);
      });
    },
    async hashPassword(password) {
      const encoder = new TextEncoder();
      const data = encoder.encode(password);
      const hashBuffer = await crypto.subtle.digest('SHA-256', data);
      const hashArray = Array.from(new Uint8Array(hashBuffer));
      const passwordHash = hashArray.map(byte => byte.toString(16).padStart(2, '0')).join('');
      return passwordHash;
    },
  }
}
</script>

<style>
/* 可以添加自訂的樣式 */
</style>