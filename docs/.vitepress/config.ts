import type {UserConfig} from 'vitepress'

const config: UserConfig = {
    title: "infinite",
    description: '测试测试测试测试',
    base: "/infinite/",
    themeConfig: {
        
        editLink: {
            pattern: 'https://github.com/fzdwx/infinite/edit/page/docs/:path',
            text: 'Suggest changes to this page'
        },

        socialLinks: [
            { icon: 'github', link: 'https://github.com/fzdwx/infinite' }
        ],

        localeLinks: {
            text: '简体中文',
            items: [
                { text: 'English', link: '/en/' },
            ]
        },
        
        nav: [
            {
                text: '快速开始',
                link: '/zh/getting-started',
                activeMatch: '^/zh/'
            },
            {text: 'team', link: '/zh/team', activeMatch: '^/infinite/zh'},
            
            {text: 'team', link: '/en/team', activeMatch: '^/infinite/en'},
        ],

        sidebar: {

            '/zh/': [
                {
                    items: [
                        {text: '快速开始', link: 'zh/getting-started'},
                    ]
                },
                {
                    text: '组件',
                    collapsible: true,
                    items: [
                        {text: 'input', link: 'zh//components/input'},
                    ]
                }
            ],

            '/en/': []
        }
    }
}

export default config