import {defineConfig} from 'vitepress'
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
    appearance: true,
    locales: locales,
    head: head,
    markdown: markdown,
    themeConfig: themeConfig,
    rewrites: {
        'en/:rest*': ':rest*'
    },
})
