import './assets/main.css'

import {createApp} from 'vue'
import App from './App.vue'
import {createPinia} from "pinia"
import "tailwindcss/tailwind.css"
import {subConfig} from "@/store/config.js";

createApp(App).use(createPinia()).mount('#app')

subConfig()