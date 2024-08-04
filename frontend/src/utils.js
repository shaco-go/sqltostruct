export const syncFuncLoad = async (f, load) => {
    load.value = true
    try {
        await f()
    } finally {
        load.value = false
    }
}