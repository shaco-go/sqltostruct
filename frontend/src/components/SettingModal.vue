<script setup>
import {NButton, NCheckbox, NDynamicInput, NFlex, NIcon, NModal, NScrollbar, NSpin} from "naive-ui"
import {ref, watch} from "vue"
import {syncFuncLoad} from "@/utils.js"
import axios from "axios";
import IconConf1 from "@/components/icons/IconConf1.vue";
import TagsModal from "@/components/TagsModal.vue";


const show = defineModel()
const tagsShow = ref(false)
const mapping = ref([])
const tags = ref([
  {name: "gorm", enable: true, is_default: true},
  {name: "form", enable: false, is_default: true},
  {name: "json", enable: true, is_default: true},
])
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
  <n-modal title="配置" style="width: 80vw;" v-model:show="show" transform-origin="center" preset="card">
    <TagsModal v-model="tagsShow"/>
    <n-spin :show="load">
      <n-scrollbar style="height: 70vh;padding: 0 20px 0 0 ">
        <n-flex class="mb-5" align="center">
          <n-checkbox v-for="item in tags" v-model:checked="item.enable" :label="item.name"/>
          <n-button quaternary circle type="primary" @click="tagsShow = true">
            <template #icon>
              <n-icon size="18px">
                <IconConf1/>
              </n-icon>
            </template>
          </n-button>

        </n-flex>
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