import {type DefaultTheme, defineConfig} from 'vitepress'

export const en = defineConfig({
  lang: 'en-US',
  title: 'carbon',
  description: 'A simple, semantic and developer-friendly time package for golang',
  head: [
    ['meta', { name: 'keywords', content: 'golang, carbon, datetime, date, time, calendar, Gregorian, Julian Day, Modified Julian Day, Lunar, Persian, Jalaali' }],
  ],
  themeConfig: {
    nav: nav(),

    sidebar: {
      '/': { items: sidebarGuide() },
    },

    editLink: {
      pattern: 'https://github.com/dromara/carbon/edit/docs/src/:path',
      text: 'Edit this page on GitHub'
    },

    footer: {
      message: 'Released under the MIT License, unauthorized reproduction is prohibited in any form',
      copyright: `Copyright © 2020-${new Date().getFullYear()} carbon team`
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
      text: 'ChangeLog',
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
        {text: 'boundary', link: 'usage/boundary'},
        {text: 'difference', link: 'usage/difference'},
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
        {text: 'test', link: 'usage/test'},
      ]
    },
    {
      text: 'Appendix',
      collapsed: false,
      items: [
        {text: 'constants', link: 'appendix/constants'},
        {text: 'Format signs', link: 'appendix/format-signs'},
        {text: 'Test reports', link: 'https://github.com/dromara/carbon/blob/master/test_report.en.md', target: '_blank'},
      ]
    },
  ]
}

export const search: DefaultTheme.AlgoliaSearchOptions['locales'] = {
  en: {}
}

