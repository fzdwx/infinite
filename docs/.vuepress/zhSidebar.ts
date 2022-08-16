import {sidebar} from "vuepress-theme-hope";

export const zhSidebarConfig = sidebar({
    "/": [
        {
            text: "README",
            link: "/"
        },
        {
            text: "组件",
            icon: "creative",
            collapsable: true,
            link: "components/",
            prefix: "components/",
            children: [
                "input/",
            ],
        },
    ],
});