# Carbon 文档

Carbon 文档网站使用 VitePress 构建。

## 本地开发

安装依赖并启动 VitePress 开发服务器：

```bash
yarn install
yarn docs:dev
```

## 构建

在本地构建文档网站：

```bash
yarn docs:build
```

构建生成的文件位于 `dist` 目录。

## 部署

文档通过 GitHub Pages Actions 部署。`.github/workflows/deploy-docs.yml` 工作流
会在 `docs` 分支有代码推送时自动运行，也可以手动触发。工作流会构建 VitePress
网站，并将生成的 `dist` 目录发布到 GitHub Pages。

网站地址：https://dromara.github.io/carbon
