import {type DefaultTheme, defineConfig} from 'vitepress'

export const ja = defineConfig({
    lang: 'ja-JP',
    title: 'carbon',
    description: '軽量、セマンティック、開発者に優しい golang 時間処理ライブラリ',
    head: [
        ['meta', { name: 'keywords', content: 'golang, carbon, 时间处理, 星座, 季节, グレゴリオ暦, 旧暦, 儒略の日, 簡略化ユリウスデー, ペルシア暦, イラン暦' }],
    ],
    themeConfig: {
        nav: nav(),

        sidebar: {
            '/ja/': { base: '/ja/',items: sidebarGuide() },
        },

        editLink: {
            pattern: 'https://github.com/dromara/carbon/edit/docs/src/:path',
            text: 'GitHubでこのページを編集する'
        },

        footer: {
            message: 'MITライセンスに基づいて公開されており、許可なく複製することは禁止されています',
            copyright: `無断転載を禁じます © 2020-${new Date().getFullYear()} carbon team`
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
            link: '/ja'
        },
        {
            text: 'ドキュメント',
            link: '/ja/overview',
            activeMatch: '/ja/overview'
        },
        {
            text: 'よくある質問',
            link: '/ja/faq',
            activeMatch: '/ja/faq'
        },
        {
            text: '変更履歴',
            link: '/ja/change-log',
            activeMatch: '/ja/change-log'
        },
        {
            text: '寄付する',
            link: '/ja/sponsor',
            activeMatch: '/ja/sponsor'
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
                {text: '时间境界', link: 'usage/boundary'},
                {text: '時間差值', link: 'usage/difference'},
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
                {text: 'テスト', link: 'usage/test'},
            ]
        },
        {
            text: 'ふろく',
            collapsed: false,
            items: [
                {text: 'コンスタント', link: 'appendix/constants'},
                {text: '書式設定記号', link: 'appendix/format-signs'},
                {base:'https://', text: 'テストレポート', link: 'github.com/dromara/carbon/blob/master/docs/BENCHMARK.ja.md', target: '_blank'},
            ]
        },
    ]
}


export const search: DefaultTheme.AlgoliaSearchOptions['locales'] = {
    ja: {
        placeholder: 'ドキュメントの検索',
        translations: {
            button: {
                buttonText: 'ドキュメントの検索',
                buttonAriaLabel: 'ドキュメントの検索'
            },
            modal: {
                searchBox: {
                    resetButtonTitle: 'クエリー条件の消去',
                    resetButtonAriaLabel: 'クエリー条件の消去',
                    cancelButtonText: 'キャンセル',
                    cancelButtonAriaLabel: 'キャンセル'
                },
                startScreen: {
                    recentSearchesTitle: '検索履歴',
                    noRecentSearchesText: '検索履歴がありません',
                    saveRecentSearchButtonTitle: '検索履歴に保存',
                    removeRecentSearchButtonTitle: '検索履歴から削除',
                    favoriteSearchesTitle: 'コレクション',
                    removeFavoriteSearchButtonTitle: 'コレクションから削除'
                },
                errorScreen: {
                    titleText: '結果を取得できませんでした',
                    helpText: 'ネットワーク接続をチェックする必要があるかもしれません'
                },
                footer: {
                    selectText: 'せんたく',
                    navigateText: '切り替え',
                    closeText: '閉じる',
                    searchByText: 'プロバイダの検索'
                },
                noResultsScreen: {
                    noResultsText: '関連する結果を見つけることができませんでした',
                    suggestedQueryText: '検索を試みることができます',
                    reportMissingResultsText: 'クエリに結果があると思いますか?',
                    reportMissingResultsLinkText: 'クリックフィードバック'
                }
            }
        }
    }
}
