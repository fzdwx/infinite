import DefaultTheme from 'vitepress/theme'
import Qwe from './vvv/Qwe.vue'

export default {
    ...DefaultTheme,
    enhanceApp({app}) {
        app.component('Qwe', Qwe)
    }
}