import type {HeadConfig} from 'vitepress';

export const head: HeadConfig[] = [
  ['link', { rel: 'icon', href: '/favicon.ico' }],
  ['meta', { name: 'theme-color', content: '#5f67ee' }],
  ['script', {}, `window._hmt = window._hmt||[];(function(){var hm=document.createElement("script");hm.src="https://hm.baidu.com/hm.js?88639be4b8ad56e48e7f1279e8ed2a0f";var s=document.getElementsByTagName("script")[0];s.parentNode.insertBefore(hm,s)})();`]
]