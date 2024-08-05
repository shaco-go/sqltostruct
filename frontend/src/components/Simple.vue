<script setup>

import {NCard, NFlex, NInput, NScrollbar} from "naive-ui";
import {onMounted, ref} from "vue";

const scrollContentRef = ref(null)
const inputCode = ref("")
const minRows = ref(20)

const resetRows = () => {
  let height = scrollContentRef.value?.scrollbarInstRef?.yBarSizePx
  if (height !== undefined) {
    minRows.value = Math.floor((parseInt(height) - 13) / 25.6)
  }
}

onMounted(() => {
  resetRows()
})
</script>

<template>
  <n-flex :vertical="false" :wrap="false">
    <n-card class="w-2/5 flex-grow-0">
      <n-scrollbar ref="scrollContentRef" style="height: calc(100vh - 82px)">
        <n-input
            class="h-full text-[16px] leading-tight"
            :autosize="{  minRows: minRows }"
            v-model:value="inputCode"
            type="textarea"
            placeholder="请输入DDL"
        />
      </n-scrollbar>
    </n-card>
    <n-card class="w-3/5 flex-shrink-0 flex-grow">
      <n-scrollbar style="height: calc(100vh - 82px)">

      </n-scrollbar>
    </n-card>
  </n-flex>
</template>

<style scoped>

</style>