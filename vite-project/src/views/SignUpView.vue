<template>
  <div class="flex justify-center items-center min-h-screen bg-gray-100">
    <div class="w-full max-w-md p-4 bg-white rounded shadow">
      <h2 class="text-2xl font-bold mb-4">會員註冊</h2>
      
      <!-- 註冊表單 -->
      <form @submit.prevent="register">
        <div class="mb-4">
          <label for="email" class="block mb-1">Email</label>
          <input v-model="email" id="email" type="email" class="w-full px-3 py-2 rounded border border-gray-300">
        </div>
        <div class="mb-4">
          <label for="password" class="block mb-1">密碼</label>
          <input v-model="password" id="password" type="password" class="w-full px-3 py-2 rounded border border-gray-300">
        </div>
        <button type="submit" class="w-full py-2 px-4 bg-blue-500 text-white rounded hover:bg-blue-600">註冊</button>
      </form>
      
      <!-- 或者使用 Google 帳號註冊 -->
      <p class="mt-4 text-center text-gray-600">或者使用 Google 帳號註冊</p>
      <button @click="registerWithGoogle" class="w-full py-2 px-4 mt-2 bg-red-500 text-white rounded hover:bg-red-600">使用 Google 帳號註冊</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      email: '',
      password: '',
    };
  },
  methods: {
    async register() {
      // 將密碼使用 SHA-256 加密
      const passwordHash = await this.hashPassword(this.password);
      
      // 使用加密後的密碼進行註冊邏輯
      console.log('註冊', this.email, passwordHash);

      
      try {
        const url = 'https://localhost/register'; // 替換為您的 API 端點
        const data = {
          email: this.email,
          password: passwordHash
        };

        const response = await axios.post(url, data);
        if (response.status === 200) {
          // 請求成功
          const responseData = response.data;
          console.log('註冊成功', responseData);
        } else {
          // 請求失敗
          console.log('註冊失敗', response.status, response.statusText);
        }
      } catch (error) {
        console.error('發生錯誤', error);
      }

    },
    registerWithGoogle() {
      // 使用 Google 帳號註冊邏輯
      console.log('使用 Google 帳號註冊');
    },
    async hashPassword(password) {
      const encoder = new TextEncoder();
      const data = encoder.encode(password);
      const hashBuffer = await crypto.subtle.digest('SHA-256', data);
      const hashArray = Array.from(new Uint8Array(hashBuffer));
      const passwordHash = hashArray.map(byte => byte.toString(16).padStart(2, '0')).join('');
      return passwordHash;
    },
  },
};
</script>

<style>
/* 可以添加自訂的樣式 */
</style>