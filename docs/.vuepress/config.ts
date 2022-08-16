import {defineUserConfig} from 'vuepress'

export default defineUserConfig({
    title: 'infinite docs',
    description: 'infinite 项目的文档',
    port: 7777,
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