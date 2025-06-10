import type {LocaleConfig} from 'vitepress';
import {en} from "./en";
import {zh} from "./zh";
import {jp} from "./jp";

export const locales: LocaleConfig = {
  root: {label: 'English', ...en},
  zh: {label: '简体中文', ...zh},
  jp: {label: '日本語', ...jp},
}