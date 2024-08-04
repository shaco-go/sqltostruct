<script setup>
import {ref} from 'vue'
import {NMessageProvider, NConfigProvider, NGlobalStyle, darkTheme, NButton, zhCN} from 'naive-ui'
import MappingSetting from "@/components/MappingSetting.vue";

let mappingSettingShow = ref(false)
</script>

<template>
  <n-config-provider style="height: 100%;width: 100%" :theme="darkTheme" :locale="zhCN">
    <n-global-style/>
    <n-message-provider>
      <n-card>
        <n-button @click="mappingSettingShow = true">按钮</n-button>
        <MappingSetting v-model="mappingSettingShow"/>
      </n-card>
    </n-message-provider>
  </n-config-provider>
</template>

<style scoped>
#app {
  min-width: 100%;
  min-height: 100%;
}
</style>
