import {defineConfig} from 'vitepress'
import {La51Plugin} from 'vitepress-plugin-51la'
import {themeConfig} from "./theme";
import {markdown} from "./markdown";
import {locales} from "./locales";
import {head} from "./head";

export default defineConfig({
    lang: 'en-US',
    title: 'carbon',
    lastUpdated: true,
    metaChunk: true,
    base: '/',
    srcDir: './src/',
    outDir: './dist',
    appearance: 'dark',
    locales: locales,
    head: head,
    markdown: markdown,
    vite: {
        plugins: [
            La51Plugin({
                id: '3MTTmzFngjK76EdH',
                ck: '3MTTmzFngjK76EdH'
            })
        ]
    },
    themeConfig: themeConfig,
    rewrites: {
        'en/:rest*': ':rest*'
    },
    sitemap: {
        hostname: 'https://carbon.go-pkg.com'
    }
})
