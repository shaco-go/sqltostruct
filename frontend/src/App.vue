<script setup xmlns="">
import {ref} from 'vue'
import {
  darkTheme,
  NConfigProvider,
  NFlex,
  NFloatButton,
  NFloatButtonGroup,
  NGlobalStyle,
  NIcon,
  NMessageProvider,
  zhCN
} from 'naive-ui'
import Simple from "@/components/Simple.vue";
import Complex from "@/components/Complex.vue";
import IconConf from "@/components/icons/IconConf.vue";
import IconCode from "@/components/icons/IconCode.vue";
import IconProject from "@/components/icons/IconProject.vue";
import SettingModal from "@/components/SettingModal.vue";
import IconBright from "@/components/icons/IconBright.vue";

const settingModalShow = ref(false)
const mode = ref("code")
const theme = ref(darkTheme)

const switchTheme = () => {
  if (theme.value === null) {
    theme.value = darkTheme
  } else {
    theme.value = null
  }
}
</script>

<template>
  <n-config-provider class="h-full w-full" :theme="theme" :locale="zhCN">
    <n-global-style/>
    <n-message-provider>
      <SettingModal v-model="settingModalShow"/>
      <!--浮动菜单-->
      <n-flex align="flex-start" class="fixed right-0 z-[99]">
        <n-float-button-group :right="40" top="40vh" position="relative">
          <n-float-button :type="mode==='code'?'primary':'default'" @click="mode='code'">
            <n-icon>
              <IconCode/>
            </n-icon>
          </n-float-button>
          <n-float-button :type="mode==='project'?'primary':'default'" @click="mode='project'">
            <n-icon>
              <IconProject/>
            </n-icon>
          </n-float-button>
          <n-float-button @click="switchTheme">
            <n-icon>
              <IconBright/>
            </n-icon>
          </n-float-button>
          <n-float-button @click="settingModalShow = true">
            <n-icon>
              <IconConf/>
            </n-icon>
          </n-float-button>
        </n-float-button-group>
      </n-flex>
      <simple v-if="mode === 'code'"></simple>
      <complex v-else></complex>
    </n-message-provider>
  </n-config-provider>
</template>

<style scoped>
</style>
