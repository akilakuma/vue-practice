<template>
    <div class="container mx-auto px-4 py-8">
      <h2 class="text-2xl font-bold mb-4">會員資料</h2>
      <div class="bg-white p-4 shadow rounded">
        <div class="mb-4">
          <label class="text-gray-600">帳號名稱</label>
          <div class="mt-1">{{ username }}</div>
        </div>
        <div class="mb-4">
          <label class="text-gray-600">EMAIL</label>
          <div class="mt-1">{{  account }}</div>
        </div>
        <div class="mb-4">
          <label class="text-gray-600">電話號碼</label>
          <div class="mt-1">{{ phone }}</div>
        </div>
      </div>
    </div>
  </template>

<script>
export default {
  name: 'UserDataView',
  props: {
    username: String,
    account: String,
    phone: Number
  }
}
</script>

<style>

</style>