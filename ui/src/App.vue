<style scoped>
.n-card {
  min-height: calc(100vh - 30px);
}
</style>
<template>
  <!--<n-message-provider>-->
  <!--  <content/>-->
  <!--</n-message-provider>-->
  <n-config-provider :theme="darkTheme">
    <n-drawer v-model:show="columnMapActive" :width="502" placement="left">
      <n-drawer-content>
        <n-scrollbar style="max-height: 100vh;padding: 20px;">
          <n-flex vertical>
            <n-form-item v-for="(item,key) in options.column_map" :label="key">
              <n-input v-model:value="options.column_map[key]"/>
            </n-form-item>
          </n-flex>
        </n-scrollbar>
      </n-drawer-content>
    </n-drawer>
    <n-drawer v-model:show="columnNullMapActive" :width="502" placement="left">
      <n-drawer-content>
        <n-scrollbar style="max-height: 100vh;padding: 20px;">
          <n-flex vertical>
            <n-form-item v-for="(item,key) in options.column_null_map" :label="key">
              <n-input v-model:value="options.column_null_map[key]"/>
            </n-form-item>
          </n-flex>
        </n-scrollbar>
      </n-drawer-content>
    </n-drawer>
    <!--<n-drawer v-model:show="tagActive" :width="502" placement="left">-->
    <!--  <n-drawer-content>-->
    <!--    <n-scrollbar style="max-height: 100vh;padding: 20px;">-->
    <!--      <n-flex vertical>-->
    <!--        <n-form-item v-for="(item,key) in options.column_null_map" :label="key">-->
    <!--          <n-input v-model:value="options.column_null_map[key]"/>-->
    <!--        </n-form-item>-->
    <!--      </n-flex>-->
    <!--    </n-scrollbar>-->
    <!--  </n-drawer-content>-->
    <!--</n-drawer>-->
    <n-flex :wrap="false">
      <n-card style="flex:2">
        <n-flex vertical>
          <n-flex align="center" style="margin-bottom: 20px">
            <n-button strong secondary type="primary" @click="columnMapActive=true">
              <template #icon>
                <n-icon>
                  <Construct/>
                </n-icon>
              </template>
              类型映射
            </n-button>
            <n-button strong secondary type="primary" @click="columnNullMapActive=true">
              <template #icon>
                <n-icon>
                  <Construct/>
                </n-icon>
              </template>
              NULL类型映射
            </n-button>
            <!--<n-button strong secondary type="primary" @click="tagActive=true">-->
            <!--  <template #icon>-->
            <!--    <n-icon>-->
            <!--      <Construct/>-->
            <!--    </n-icon>-->
            <!--  </template>-->
            <!--  Tag-->
            <!--</n-button>-->

            <n-select style="width: 150px" v-model:value="options.style" :options="styleOptions"/>

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

          <n-input
              placeholder="请填写需要解析的sql"
              type="textarea"
              :autosize="{  minRows: 30}"
          />
        </n-flex>
      </n-card>
      <n-card style="flex:3">
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
import {ref, onMounted, watch, reactive} from "vue"
import {
  darkTheme,
  NCard,
  NConfigProvider,
  NFlex,
  NInput,
  NCode,
  NButton,
  NCheckbox,
  NDrawer,
  NIcon,
  NSelect,
  NFormItem,
  NScrollbar,
  NMessageProvider,
  useMessage,
} from 'naive-ui'
import hljs from 'highlight.js/lib/core'
import go from 'highlight.js/lib/languages/go'
import axios from "axios"
import qs from "qs"
import {Construct} from "@vicons/ionicons5"


// Content-Type 响应头
hljs.registerLanguage('go', go)
const code = ref("")
const sql = ref("")
const styleOptions = ref([
  {label: "大驼峰命名风格", value: 0},
  {label: "小驼峰命名风格", value: 1},
  {label: "中横线命名风格", value: 2},
  {label: "蛇形命名风格", value: 3},
])
const columnMapActive = ref(false)
const columnNullMapActive = ref(false)
const tagActive = ref(false)
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
  axios.get("http://localhost:7788/api/config").then(resp => {
    if (resp.data.code === 200) {
      options.value = resp.data.data
    }
  })
}

const saveConf = function () {
  axios.patch("http://localhost:7788/api/config", {
    ...options.value,
    tag:JSON.stringify(options.value.tag),
    column_null_map:JSON.stringify(options.value.column_null_map),
    column_map:JSON.stringify(options.value.column_map)
  }).then(resp => {
    if (resp.data.code === 200) {
      // window.$message.success(resp.data.message)
      return
    }
    // window.$message.error(resp.data.message)
  })
}

watch(options, (n, o) => {
  if (JSON.stringify(n) === JSON.stringify(o)) {
    saveConf()
  }
}, {deep: true})

onMounted(() => {
  // 获取配置项
  getConf()
  window.$message = useMessage()
  window.$message.success("test")
})

</script>