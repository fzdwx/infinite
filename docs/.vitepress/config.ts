import {defineConfig} from 'vitepress'

export default defineConfig(
    {
        head: [
            ['link', {rel: 'stylesheet', href: 'https://cdnjs.cloudflare.com/ajax/libs/tailwindcss/2.2.19/tailwind.min.css'}]
        ],
        title: "infinite",
        description: '测试测试测试测试',
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
                message: 'Released under the MIT License.',
                copyright: 'Copyright © 2022-present fzdwx'
            },

            sidebar: {

                '/zh/': [
                    {
                        text: 'Guide',
                        collapsible: true,
                        items: [
                            {text: '简介', link: 'zh/guide/'},
                            {text: '快速开始', link: 'zh/guide/getting-started'},
                        ]
                    },
                    {
                        text: '组件',
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