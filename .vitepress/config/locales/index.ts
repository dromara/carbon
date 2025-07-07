import type {LocaleConfig} from 'vitepress';
import {en} from "./en";
import {zh} from "./zh";
import {ja} from "./ja";
import {ko} from "./ko";

export const locales: LocaleConfig = {
  root: {label: 'English', ...en},
  zh: {label: '简体中文', ...zh},
  ja: {label: '日本語', ...ja},
  ko: {label: '한국어', ...ko},
}