import {type DefaultTheme, defineConfig} from 'vitepress'

export const en = defineConfig({
  lang: 'en-US',

  themeConfig: {
    nav: nav(),

    sidebar: {
      '/': { base: '/', items: sidebarGuide() },
    },

    editLink: {
      pattern: 'https://github.com/golang-module/carbon-docs/blob/main/src/:path',
      text: 'Edit this page on GitHub'
    },

    footer: {
      message: 'Released under the MIT License, unauthorized reproduction is prohibited in any form',
      copyright: `Copyright © 2024-${new Date().getFullYear()} carbon <a href="https://beian.miit.gov.cn" target="_blank">京ICP备19041346号-6</a>`
    },

    outline: {
      level: [2, 6],
      label: 'on this page'
    },

    lastUpdated: {
      text: 'Last updated',
      formatOptions: {
        dateStyle: 'short',
        timeStyle: 'medium'
      }
    },
  }
})

function nav(): DefaultTheme.NavItem[] {
  return [
    {
      text: 'Home',
      link: '/'
    },
    {
      text: 'Doc',
      link: '/overview',
      activeMatch: '/overview'
    },
    {
      text: 'FAQ',
      link: '/faq',
      activeMatch: '/faq'
    },
    {
      text: 'CHANGELOG',
      link: '/change-log',
      activeMatch: '/change-log'
    },
    {
      text: 'Sponsor',
      link: '/sponsor',
      activeMatch: '/sponsor'
    },
  ]
}

function sidebarGuide(): DefaultTheme.SidebarItem[] {
  return [
    {
      text: 'Getting-started',
      collapsed: false,
      items: [
        {text: 'overview', link: 'overview',},
        {text: 'installation', link: 'getting-started',},
      ]
    },
    {
      text: 'Basic usage',
      collapsed: false,
      items: [
        {text: 'default', link: 'usage/default'},
        {text: 'conversion', link: 'usage/conversion'},
        {text: 'yesterday/today/tomorrow', link: 'usage/yesterday-today-tomorrow'},
        {text: 'creator', link: 'usage/creator'},
        {text: 'parser', link: 'usage/parser'},
        {text: 'traveler', link: 'usage/traveler'},
        {text: 'setter', link: 'usage/setter'},
        {text: 'getter', link: 'usage/getter'},
        {text: 'output', link: 'usage/output'},
        {text: 'difference', link: 'usage/difference'},
        {text: 'boundary', link: 'usage/boundary'},
        {text: 'extremum', link: 'usage/extremum'},
        {text: 'comparison', link: 'usage/comparison'},
        {text: 'error handling', link: 'usage/error'},
      ]
    },
    {
      text: 'Extended usage',
      collapsed: false,
      items: [
        {text: 'constellation', link: 'usage/constellation'},
        {text: 'season', link: 'usage/season'},
        {text: 'calendar', link: 'usage/calendar'},
        {text: 'i18n', link: 'usage/i18n'},
        {text: 'json', link: 'usage/json'},
        {text: 'freeze test', link: 'usage/freeze'},
      ]
    },
    {
      text: 'Appendix',
      collapsed: false,
      items: [
        {text: 'constants', link: 'appendix/constants'},
        {text: 'Format signs', link: 'appendix/format-signs'},
      ]
    },
  ]
}

export const search: DefaultTheme.AlgoliaSearchOptions['locales'] = {
  en: {
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

