

<template>
  <div class="h-screen flex justify-center items-center">
    <div class="flex justify-end bg-local bg-image" :style="{ backgroundImage: `url(${imagePath})` }">
      <div class="block w-1/3 mr-40">
         <div class="font-mono pt-12 text-zinc-600 text-2xl leading-normal">


          
          歡迎來到我們的可愛小浣熊風格飯店！我們的飯店設計靈感來自於迷人的小浣熊，旨在為您提供一個溫馨、舒適且充滿趣味的住宿經驗。
        </div>
        <div class="font-mono pt-12 text-zinc-500 text-2xl leading-normal">
          當您步入我們的大堂時，您將被迎接著一個生動可愛的浣熊雕塑，這個友善的小家夥將帶給您微笑與歡樂。整個飯店充滿著浣熊的元素，從牆壁上的壁畫到房間裡的配飾，處處都彰顯著可愛與童趣。
        </div>
        <div class="font-mono pt-12 text-zinc-600 text-2xl leading-normal">
          我們的客房設計獨具特色，每個房間都以浣熊為主題，配備了舒適的床鋪、柔軟的枕頭和豪華的床單，讓您在温馨的環境中享受舒適的睡眠。每間客房都配有現代化的設施，包括高速無線網絡、平面電視和迷你冰箱，以滿足您的需求。
        </div>
        <div class="font-mono pt-12 text-zinc-500 text-2xl leading-normal">
          我們的飯店還設有一個充滿浣熊主題的餐廳，供應各種美味的料理。您可以品嚐到經過精心製作的美食，同時欣賞著飯店的獨特裝潢。我們的員工都非常友善且熱情，將竭誠為您提供優質的服務，確保您有一個難忘的入住體驗。
        </div> 
        <div class="font-mono pt-12 text-zinc-600 text-2xl leading-normal">
          除了舒適的住宿和美食，我們的飯店還提供各種娛樂設施，如健身房、SPA中心和游泳池，讓您在旅途中放鬆身心。此外，我們的飯店還位於自然環境優美的地區，您可以參加由我們提供的導覽活動，探索周圍的大自然和野生動物。
        </div>
      </div>
    </div>
  </div>
</template>



<script setup>
import imagePath from '@/assets/home-bg-3.jpg'
</script>

<style>
.bg-image {
  width: 100%;
  height: 100%;
  background-size: cover;
}
</style>