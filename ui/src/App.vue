<style scoped>
.n-card {
  min-height: calc(100vh - 30px);
}
</style>
<template>
  <n-config-provider :theme="darkTheme">
    <n-flex :wrap="false">
      <n-flex vertical>
        <n-flex align="center">
          <n-button strong secondary type="primary">
            类型映射
          </n-button>
          <n-button strong secondary type="primary">
            NULL类型映射
          </n-button>
          <n-checkbox v-model:checked="options.is_null">
            NULL类型
          </n-checkbox>
          <n-checkbox v-model:checked="options.is_comment">
            注释
          </n-checkbox>
          <n-checkbox v-model:checked="options.is_gorm">
            GORM
          </n-checkbox>
          <n-checkbox v-for="(value,key) in options.tag" v-model:checked="options.tag[key]">
            {{ key }}
          </n-checkbox>
        </n-flex>
        <n-card style="flex: 1">
          <n-input
              placeholder="请填写需要解析的sql"
              type="textarea"
              :autosize="{  minRows: 30}"
          />
        </n-card>
      </n-flex>


      <n-card style="flex: 2">
        <n-code
            :hljs="hljs"
            :code="code"
            language="go"
        />
      </n-card>
    </n-flex>
  </n-config-provider>

</template>
<script setup>
import {ref, onMounted} from "vue"
import {darkTheme, NCard, NConfigProvider, NFlex, NInput, NCode, NButton, NCheckbox} from 'naive-ui'
import hljs from 'highlight.js/lib/core'
import go from 'highlight.js/lib/languages/go'
import axios from "axios"

hljs.registerLanguage('go', go)

const code = ref("")
const sql = ref("")
const options = ref({
  is_null: false,
  is_default: false,
  is_unsigned: false,
  is_len: false,
  is_gorm: false,
  is_comment: false,
  style: 0,
  column_map: [],
  column_null_map: [],
  tag: []
})

const getConf = function () {
  axios.get("http://localhost:7899/config").then(resp => {
    if (resp.data.code === 200) {
      options.value = resp.data.data
    }
  })
}

onMounted(() => {
  // 获取配置项
  getConf()
})

</script>