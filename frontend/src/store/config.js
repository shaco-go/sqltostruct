import {defineStore} from "pinia";

export const configStore = defineStore("config", {
    state: () => {
        return {

        }
    }
})

export const subConfig = () => {
    configStore().$subscribe((mutation, state) => {

    })
}