<style scoped>
.n-card {
  min-height: calc(100vh - 30px);
}
</style>
<template>
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
      <n-card style="flex:2 0 0">
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
              v-model:value="sql"
              :autosize="{  minRows: 30,maxRows:30}"
          />
        </n-flex>

      </n-card>
      <n-card style="width: 60%" title="Code">
        <template #header-extra>
          <n-button strong secondary type="primary" @click="copy">复制</n-button>
        </template>
        <n-flex align="center" style="height: 100%" v-if="sql!=='' && code === ''">
          <n-empty style="margin: auto" size="large" description="解析sql失败">
            <template #icon>
              <n-icon>
                <ReaderOutline/>
              </n-icon>
            </template>
          </n-empty>
        </n-flex>
        <n-scrollbar style="width: 100%;height: 100%" x-scrollable v-else-if="sql!==''">
          <n-code
              :hljs="hljs"
              :code="code"
              language="go"
          />
        </n-scrollbar>

      </n-card>
    </n-flex>
  </n-config-provider>


</template>
<script setup>
import {ref, onMounted, watch} from "vue"
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
  NEmpty,
  createDiscreteApi,
} from 'naive-ui'
import hljs from 'highlight.js/lib/core'
import go from 'highlight.js/lib/languages/go'
import axios from "axios"
import {Construct, ReaderOutline} from "@vicons/ionicons5"
import debounce from "lodash/debounce"
import lo from "lodash"
import useClipboard  from "vue-clipboard3"

const {message} = createDiscreteApi(["message"])
const {toClipboard}  = useClipboard()
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
  axios.get("api/config").then(resp => {
    if (resp.data.code === 200) {
      options.value = resp.data.data
    }
  })
}

const saveConf = debounce(() => {
  axios.patch("api/config", {
    ...options.value,
    tag: JSON.stringify(options.value.tag),
    column_null_map: JSON.stringify(options.value.column_null_map),
    column_map: JSON.stringify(options.value.column_map)
  }).then(resp => {
    if (resp.data.code === 200) {
      message.success('配置已保存')
      return
    }
    message.error(resp.data.message)
  })
}, 1000)

const copy = async () => {
  await toClipboard(code.value)
  message.success("已复制到剪贴板")
}

const generateSql = debounce(() => {
  axios.post("api/generate", {sql: sql.value}).then(resp => {
    if (resp.data.code === 200) {
      code.value = resp.data.data.code
      return
    }
    code.value = ''
  })
}, 500)

watch(options, (n, o) => {
  if (JSON.stringify(n) === JSON.stringify(o)) {
    saveConf()
  }
}, {deep: true})

watch(sql, (n, o) => {
  console.log(n, o)
  if (lo.trim(n) !== "") {
    generateSql()
  }
})

onMounted(() => {
  // 获取配置项
  getConf()
})

</script>