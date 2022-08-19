import {defineConfig} from 'vitepress'

export default defineConfig(
    {

        markdown: {},
        head: [
            ['link', {rel: 'icon', href: 'https://raw.githubusercontent.com/fzdwx/infinite/page/docs/asset/infinite.svg'}]
        ],
        title: "Infinite",
        base: "/infinite/",
        themeConfig: {
            editLink: {
                pattern: 'https://github.com/fzdwx/infinite/edit/page/docs/:path',
                text: 'Suggest changes to this page'
            },

            socialLinks: [
                {icon: 'github', link: 'https://github.com/fzdwx/infinite'}
            ],

            localeLinks: {
                text: '简体中文',
                items: [
                    {text: 'English', link: '/en/'},
                ]
            },

            nav: [
                // {
                //     text: '快速开始',
                //     link: '/zh/getting-started',
                // },
            ],

            footer: {
                message: 'Released under the <a target="_blank" style="color: #a5a3a3" href="https://github.com/fzdwx/infinite/blob/main/LICENSE">MIT License</a>.',
                copyright: 'Copyright © 2022-present <a target="_blank" style="color: #1BE5BB" href="https://github.com/fzdwx">fzdwx</a>'
            },

            sidebar: {

                '/zh/': [
                    {
                        text: 'Guide',
                        collapsible: true,
                        items: [
                            {text: '简介', link: 'zh/guide/'},
                            {text: '快速开始', link: 'zh/guide/getting-started'},
                            {text: '示例', link: 'zh/guide/examples'},
                        ]
                    },
                    {
                        text: 'Components',
                        collapsible: true,
                        items: [
                            {text: 'input', link: 'zh/components/input'},
                        ]
                    }
                ],

                '/en/': [
                    {
                        text: 'Guide',
                        collapsible: true,
                        items: [
                            {text: 'Introduction', link: 'en/guide/'},
                            {text: 'Getting started', link: 'en/guide/getting-started'},
                            {text: 'Examples', link: 'en/guide/examples'},
                        ]
                    },
                    {
                        text: 'Components',
                        collapsible: true,
                        items: [
                            {text: 'input', link: 'en/components/input'},
                        ]
                    }
                ]
            }
        }
    }
)