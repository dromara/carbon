import type {HeadConfig} from 'vitepress';

export const head: HeadConfig[] = [
  ['link', { rel: 'icon', href: '/favicon.ico' }],
  ['meta', { name: 'theme-color', content: '#5f67ee' }],
  ['meta', { property: 'og:type', content: 'website' }],
  ['meta', { property: 'og:site_name', content: 'carbon' }],
  ['meta', { property: 'og:url', content: 'https://carbon.go-pkg.com' }],
  ['meta', { name: "baidu-site-verification", content: "codeva-E2ORyxQl2d" }],
  ['meta', { name: "google-site-verification", content: "BbBdC4u76yDtOYRHjo5z__72et0bp8UN6x5ZgTlZ2fY" }],
  ['meta', { name: "viewport", content: "width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no,shrink-to-fit=no" }],
  ['script', {}, `window._hmt = window._hmt||[];(function(){var hm=document.createElement("script");hm.src="https://hm.baidu.com/hm.js?88639be4b8ad56e48e7f1279e8ed2a0f";var s=document.getElementsByTagName("script")[0];s.parentNode.insertBefore(hm,s)})();`]
]