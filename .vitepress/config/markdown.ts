import type {MarkdownOptions} from 'vitepress';

export const markdown: MarkdownOptions = {
  // theme: 'github-dark',
  math: true,
  codeTransformers: [
    {
      postprocess(code) {
        return code.replace(/\[\!\!code/g, '[!code')
      }
    }
  ],
  config(md) {
    const fence = md.renderer.rules.fence!
    md.renderer.rules.fence = function (tokens, idx, options, env, self) {
      const { localeIndex = 'root' } = env
      const codeCopyButtonTitle = (() => {
        switch (localeIndex) {
          case 'en':
            return 'Copy code'
          case 'jp':
            return 'コードをコピー'
          case 'zh':
            return '复制代码'
          default:
            return 'Copy code'
        }
      })()
      return fence(tokens, idx, options, env, self).replace(
          '<button title="Copy Code" class="copy"></button>',
          `<button title="${codeCopyButtonTitle}" class="copy"></button>`
      )
    }
  }
}