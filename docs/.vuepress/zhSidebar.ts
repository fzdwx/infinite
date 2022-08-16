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
            prefix: "components/",
            children: [
                "input/",
            ],
        },
    ],
});