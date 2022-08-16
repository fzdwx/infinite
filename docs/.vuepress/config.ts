import {defineUserConfig} from 'vuepress'
import {hopeTheme} from "vuepress-theme-hope";
import {zhSidebarConfig} from "./zhSidebar";

export default defineUserConfig({
    title: 'infinite docs',
    description: 'infinite 项目的文档',
    port: 7777,
    base: "/infinite/",
    theme: hopeTheme({
        iconAssets: "iconfont",
        
        footer: "<a href=\"https://github.com/fzdwx\">fzdwx</a>",
        repo: 'fzdwx/infinite',
        // 全屏
        fullscreen: true,

        // 插件
        plugins: {
            // mdEnhance: {
            //     enableAll: true,
            // },
        },

        locales: {
            '/': {
                sidebar: zhSidebarConfig
            },
            '/en/': {
                sidebar: []
            },

        }
    }),
    locales: {
        '/': {
            lang: 'zh-CN',
            title: 'infinite',
            description: '🧬 用于开发交互式 CLI(tui,terminal) 程序的组件库.',
        },
        '/en/': {
            lang: 'en-US',
            title: 'infinite',
            description: '🧬 A component library for developing interactive CLI (tui, terminal) programs.',
        },
    }
})