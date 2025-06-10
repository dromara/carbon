import {type DefaultTheme, defineConfig} from 'vitepress'

export const jp = defineConfig({
    lang: 'ja-JP',
    themeConfig: {
        nav: nav(),

        sidebar: {
            '/jp/': { base: '/jp/', items: sidebarGuide() },
        },

        editLink: {
            pattern: 'https://github.com/golang-module/carbon-docs/blob/main/src/:path',
            text: 'GitHubでこのページを編集する'
        },

        footer: {
            message: 'MITライセンスに基づいて公開されており、許可なく複製することは禁止されています。',
            copyright: `無断転載を禁じます © 2024-${new Date().getFullYear()} carbon <a href="https://beian.miit.gov.cn" target="_blank">京ICP备19041346号-6</a>`
        },

        docFooter: {
            prev: '前のページ',
            next: '次のページ'
        },

        outline: {
            level: [2, 6],
            label: '現在のディレクトリ'
        },

        lastUpdated: {
            text: '最終更新日',
            formatOptions: {
                dateStyle: 'short',
                timeStyle: 'medium'
            }
        },

        langMenuLabel: '多言語',
        returnToTopLabel: 'トップに戻る',
        sidebarMenuLabel: 'メニュー',
        darkModeSwitchLabel: 'トピック＃トピック＃',
        lightModeSwitchTitle: 'ライトモードに切り替え',
        darkModeSwitchTitle: 'ダークカラーモードに切り替え'
    }
})

function nav(): DefaultTheme.NavItem[] {
    return [
        {
            text: 'ホームページ',
            link: '/jp'
        },
        {
            text: 'ドキュメント',
            link: '/jp/overview',
            activeMatch: '/jp/overview'
        },
        {
            text: 'よくある質問',
            link: '/jp/faq',
            activeMatch: '/jp/faq'
        },
        {
            text: '変更履歴',
            link: '/jp/change-log',
            activeMatch: '/jp/change-log'
        },
        {
            text: '寄付する',
            link: '/jp/sponsor',
            activeMatch: '/jp/sponsor'
        },
    ]
}

function sidebarGuide(): DefaultTheme.SidebarItem[] {
    return [
        {
            text: 'はじめに',
            collapsed: false,
            items: [
                {text: 'プロジェクトの概要', link: 'overview',},
                {text: 'クイックスタート', link: 'getting-started',},
            ]
        },
        {
            text: '基本的な使い方',
            collapsed: false,
            items: [
                {text: 'グローバルのデフォルト値設定', link: 'usage/default'},
                {text: 'carbon/time.Time 間の変換', link: 'usage/conversion'},
                {text: '昨日/現在/明日', link: 'usage/yesterday-today-tomorrow'},
                {text: 'carbon インスタンスを作成する', link: 'usage/creator'},
                {text: '時間解析', link: 'usage/parser'},
                {text: 'タイムトラベル', link: 'usage/traveler'},
                {text: '時間設定', link: 'usage/setter'},
                {text: '時間取得', link: 'usage/getter'},
                {text: '時間出力', link: 'usage/output'},
                {text: '時間差分', link: 'usage/difference'},
                {text: '时间境界', link: 'usage/boundary'},
                {text: '時間極值', link: 'usage/extremum'},
                {text: '時間比較', link: 'usage/comparison'},
                {text: 'エラー処理', link: 'usage/error'},
            ]
        },
        {
            text: '拡張の使用方法',
            collapsed: false,
            items: [
                {text: '星座', link: 'usage/constellation'},
                {text: '季節', link: 'usage/season'},
                {text: 'カレンダ＃カレンダ＃', link: 'usage/calendar'},
                {text: '国際化', link: 'usage/i18n'},
                {text: 'JSON', link: 'usage/json'},
                {text: '時間凍結', link: 'usage/freeze'},
            ]
        },
        {
            text: 'ふろく',
            collapsed: false,
            items: [
                {text: 'コンスタント', link: 'appendix/constants'},
                {text: '書式設定記号', link: 'appendix/format-signs'},
            ]
        },
    ]
}


export const search: DefaultTheme.AlgoliaSearchOptions['locales'] = {
    jp: {
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
