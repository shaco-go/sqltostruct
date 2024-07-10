<style scoped>
.n-card {
  min-height: calc(100vh - 30px);
}
</style>
<template>
  <n-config-provider :theme="darkTheme">
    <n-flex :wrap="false">
      <n-card style="flex: 1">
        <n-input
            placeholder="请填写需要解析的sql"
            type="textarea"
            :autosize="{  minRows: 30}"
        />
      </n-card>

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
import {ref} from "vue"
import {darkTheme, NCard, NConfigProvider, NFlex, NInput, NCode} from 'naive-ui'
import hljs from 'highlight.js/lib/core'
import go from 'highlight.js/lib/languages/go'

const code = ref(`type User struct {
\tId        int     \`gorm:"column:id"\`
\tName      varchar \`gorm:"column:name"\`
\tAge       tinyint \`gorm:"column:age"\`
\tPrice     decimal \`gorm:"column:price"\`
\tCreatedAt date    \`gorm:"column:created_at"\`
}`)

hljs.registerLanguage('go', go)
</script>