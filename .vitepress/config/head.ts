import type {HeadConfig} from 'vitepress';

export const head: HeadConfig[] = [
  ['link', { rel: 'icon', href: '/favicon.ico' }],
  ['meta', { name: 'theme-color', content: '#5f67ee' }],
  ['script', {}, `window._hmt = window._hmt||[];(function(){var hm=document.createElement("script");hm.src="https://hm.baidu.com/hm.js?5dc761ba104eb61c3443cbe14e98389e";var s=document.getElementsByTagName("script")[0];s.parentNode.insertBefore(hm,s)})();`]
]