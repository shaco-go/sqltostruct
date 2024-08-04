<script setup>
import {NButton, NDynamicInput, NFlex, NModal, NScrollbar, NSpin} from "naive-ui"
import {defineModel, ref, watch} from "vue"
import {syncFuncLoad} from "@/utils.js"
import axios from "axios";


const show = defineModel()
const mapping = ref([])
const load = ref(false)

// 获取配置
const getMapping = () => {
  syncFuncLoad(() => axios.get("http://127.0.0.1:8080/").then(res => {
    console.log(res)
  }), load)
}

// 设置配置
const setMapping = () => {
  syncFuncLoad(() => {
    mapping.value = [{key: "ceshi", value: "a"}]
  }, load)
}

const resetMapping = () => {
  syncFuncLoad(() => {
    mapping.value = [{key: "ceshi", value: "a"}]
  }, load)
}

watch(show, (newVal) => {
  if (newVal) {
    // 获取配置
    getMapping()
  }
})


</script>

<template>
  <n-modal title="字段映射设置" style="width: 80vw;" v-model:show="show" transform-origin="center" preset="card">
    <n-spin :show="load">
      <n-scrollbar style="height: 70vh;padding: 0 20px 0 0 ">
        <n-dynamic-input
            v-model:value="mapping"
            preset="pair"
            key-placeholder="类型"
            value-placeholder="映射值"
        />
      </n-scrollbar>
    </n-spin>
    <template #footer>
      <n-flex justify="end">
        <n-button secondary @click="show = false">取消</n-button>
        <n-button secondary type="info" @click="resetMapping" :disabled="load">重置默认配置</n-button>
        <n-button secondary type="primary" @click="setMapping" :disabled="load">保存</n-button>
      </n-flex>
    </template>
  </n-modal>
</template>

<style scoped>
</style>