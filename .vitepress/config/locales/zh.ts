import {type DefaultTheme, defineConfig} from 'vitepress'

export const zh = defineConfig({
    lang: 'zh-Hans',
    title: 'carbon',
    description: '一个轻量级、语义化、对开发者友好的 golang 时间处理库',
    head: [
        ['meta', { name: 'keywords', content: 'golang, carbon, 时间处理, 星座, 季节, 格里历, 农历, 儒略日, 简化儒略日, 波斯历, 伊朗历' }],
    ],
    themeConfig: {
        nav: nav(),

        sidebar: {
            '/zh/': { base: '/zh/', items: sidebarGuide() },
        },

        editLink: {
            pattern: 'https://github.com/dromara/carbon/edit/docs/src/:path',
            text: '在 GitHub 上编辑此页面'
        },

        footer: {
            message: '基于 MIT 许可发布，未经许可禁止任何形式的转载',
            copyright: `版权所有 © 2020-${new Date().getFullYear()} carbon team <a href="https://beian.miit.gov.cn" target="_blank">京ICP备19041346号-7</a>`
        },

        docFooter: {
            prev: '上一页',
            next: '下一页'
        },

        outline: {
            level: [2, 6],
            label: '当前页面'
        },

        lastUpdated: {
            text: '最后更新于',
            formatOptions: {
                dateStyle: 'short',
                timeStyle: 'medium'
            }
        },

        langMenuLabel: '多语言',
        returnToTopLabel: '回到顶部',
        sidebarMenuLabel: '菜单',
        darkModeSwitchLabel: '主题',
        lightModeSwitchTitle: '切换到浅色模式',
        darkModeSwitchTitle: '切换到深色模式'
    }
})

function nav(): DefaultTheme.NavItem[] {
    return [
        {
            text: '首页',
            link: '/zh'
        },
        {
            text: '使用文档',
            link: '/zh/overview',
            activeMatch: '/zh/overview'
        },
        {
            text: '常见问题',
            link: '/zh/faq',
            activeMatch: '/zh/faq'
        },
        {
            text: '更新日志',
            link: '/zh/change-log',
            activeMatch: '/zh/change-log'
        },
        {
            text: '赞助支持',
            link: '/zh/sponsor',
            activeMatch: '/zh/sponsor'
        },
        {
            text: 'Dongle',
            link: 'https://dongle.go-pkg.com/',
        },
    ]
}

function sidebarGuide(): DefaultTheme.SidebarItem[] {
    return [
        {
            text: '入门指引',
            collapsed: false,
            items: [
                {text: '项目简介', link: 'overview',},
                {text: '快速开始', link: 'getting-started',},
            ]
        },
        {
            text: '基础用法',
            collapsed: false,
            items: [
                {text: '设置默认值', link: 'usage/default'},
                {text: 'carbon/time.Time 互转', link: 'usage/conversion'},
                {text: '昨天/今天/明天', link: 'usage/yesterday-today-tomorrow'},
                {text: '创建 Carbon 实例', link: 'usage/creator'},
                {text: '时间解析', link: 'usage/parser'},
                {text: '时间旅行', link: 'usage/traveler'},
                {text: '时间设置', link: 'usage/setter'},
                {text: '时间获取', link: 'usage/getter'},
                {text: '时间输出', link: 'usage/output'},
                {text: '时间边界', link: 'usage/boundary'},
                {text: '时间差值', link: 'usage/difference'},
                {text: '时间极值', link: 'usage/extremum'},
                {text: '时间判断', link: 'usage/comparison'},
                {text: '错误处理', link: 'usage/error'},
            ]
        },
        {
            text: '扩展用法',
            collapsed: false,
            items: [
                {text: '星座', link: 'usage/constellation'},
                {text: '季节', link: 'usage/season'},
                {text: '日历', link: 'usage/calendar'},
                {text: '国际化', link: 'usage/i18n'},
                {text: 'JSON', link: 'usage/json'},
                {text: '测试', link: 'usage/test'},
            ]
        },
        {
            text: '附录',
            collapsed: false,
            items: [
                {text: '内置常量', link: 'appendix/constants'},
                {text: '格式符号', link: 'appendix/format-signs'},
                {base:'https://', text: '测试报告', link: 'github.com/dromara/carbon/blob/master/docs/BENCHMARK.cn.md', target: '_blank'},
            ]
        },
    ]
}

export const search: DefaultTheme.AlgoliaSearchOptions['locales'] = {
    zh: {
        placeholder: '搜索文档',
        translations: {
            button: {
                buttonText: '搜索文档',
                buttonAriaLabel: '搜索文档'
            },
            modal: {
                searchBox: {
                    resetButtonTitle: '清除查询条件',
                    resetButtonAriaLabel: '清除查询条件',
                    cancelButtonText: '取消',
                    cancelButtonAriaLabel: '取消'
                },
                startScreen: {
                    recentSearchesTitle: '搜索历史',
                    noRecentSearchesText: '没有搜索历史',
                    saveRecentSearchButtonTitle: '保存至搜索历史',
                    removeRecentSearchButtonTitle: '从搜索历史中移除',
                    favoriteSearchesTitle: '收藏',
                    removeFavoriteSearchButtonTitle: '从收藏中移除'
                },
                errorScreen: {
                    titleText: '无法获取结果',
                    helpText: '你可能需要检查你的网络连接'
                },
                footer: {
                    selectText: '选择',
                    navigateText: '切换',
                    closeText: '关闭',
                    searchByText: '搜索提供者'
                },
                noResultsScreen: {
                    noResultsText: '无法找到相关结果',
                    suggestedQueryText: '你可以尝试查询',
                    reportMissingResultsText: '你认为该查询应该有结果？',
                    reportMissingResultsLinkText: '点击反馈'
                }
            }
        }
    }
}
