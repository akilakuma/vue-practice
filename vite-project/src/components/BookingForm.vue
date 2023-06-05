<template>
    <div class="flex justify-center items-center flex-col mt-10">
        <div class="bg-gray-300">
            <div class="m-2 flex items-center">
                <label class="mr-3 w-20">Name</label>
                <div class="justify-start">
                    <input class="rounded p-1" type="text">
                </div>
            </div>
            <div class="m-2 flex items-center">
                <label class="mr-3 w-20">Date</label>
                <div class="justify-start">
                    <input class="rounded p-1" type="date">
                </div>
            </div>
            <div class="m-2 flex items-center">
                <label class="mr-3 w-20">Email</label>
                <div class="justify-start">
                    <input class="rounded p-1" type="email">
                </div>
            </div>
            <div class="m-2 flex items-center">
                <label class="mr-3 w-20">Phone</label>
                <div class="justify-start">
                    <input class="rounded p-1" type="tel">
                </div>
            </div>
        </div>
        <button class="rounded bg-zinc-200 text-red-500">立刻下訂</button>
    </div>
</template>
  
<script>
export default {

}
</script>
 