// noinspection TypeScriptUnresolvedReference

import DefaultTheme from 'vitepress/theme'
import './vars.css'

declare var _hmt: any;

DefaultTheme.enhanceApp = ({router}) => {
    router.onBeforeRouteChange = (to) => {
        if (typeof _hmt !== 'undefined') {
            _hmt.push(['_trackPageview', to]);
        }
    };
}
export default { ...DefaultTheme }