<script setup>
import {NModal} from "naive-ui";

const show = defineModel()
</script>

<template>
  <n-modal :z-index="9999" v-model:show="show" class="h-[30vh] w-[50vw]" :show="true" transform-origin="center" preset="card">
    ceshi
  </n-modal>
</template>

<style scoped>

</style>