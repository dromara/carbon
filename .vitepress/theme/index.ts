// noinspection TypeScriptUnresolvedReference

import DefaultTheme from 'vitepress/theme'
import './vars.css'
import AsideAd from './components/AsideAd.vue'
import TopBanner from './components/TopBanner.vue'
import { h } from 'vue'

declare var _hmt: any;

DefaultTheme.enhanceApp = ({router}) => {
    router.onBeforeRouteChange = (to) => {
        if (typeof _hmt !== 'undefined') {
            _hmt.push(['_trackPageview', to]);
        }
    };
}

export default {
    extends: DefaultTheme,
    Layout: () => {
        return h(DefaultTheme.Layout, null, {
            'doc-before': () => h(TopBanner),
            'aside-outline-after': () => h(AsideAd)
        })
    }
}