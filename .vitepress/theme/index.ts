// noinspection TypeScriptUnresolvedReference

import DefaultTheme from 'vitepress/theme'
import './vars.css'
import './custom.css'
import AsideAd from './components/AsideAd.vue'
import HomeBanner from './components/HomeBanner.vue'
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
            'aside-outline-after': () => h(AsideAd),
            'home-hero-info-after': () => h(HomeBanner)
        })
    }
}