# Carbon 文档 URL 推送工具

这是一个统一的 Go 工具，用于将 Carbon 文档网站的所有页面 URL 推送到搜索引擎，支持 IndexNow 和百度搜索提交。

## 📋 功能特性

- ✅ **统一架构**: 基于 `engine.go` 核心引擎和 `main.go` 程序入口的模块化设计
- ✅ **多平台支持**: 同时支持 IndexNow 和百度搜索提交
- ✅ **自动读取 URL**: 从 `dist/sitemap.xml` 自动读取所有构建后的页面 URL
- ✅ **配置验证**: 完整的配置验证，包括密钥文件可访问性检查
- ✅ **预览模式**: 预览将要提交的 URL 而不实际提交
- ✅ **批量提交**: 自动处理大量 URL 的批量提交 (IndexNow 每批最多 10,000 个)
- ✅ **错误处理**: 详细的错误信息和故障排除建议
- ✅ **中文界面**: 完全中文的用户界面和错误信息

## 🛠️ 使用方法

### 步骤 1: 构建文档站点

在使用工具之前，需要先构建文档站点以生成 sitemap.xml：

```bash
# 在项目根目录运行
npm run build
```

这将在 `dist/` 目录下生成 `sitemap.xml` 文件，包含所有页面的 URL。

### 步骤 2: 准备环境变量文件

进入工具目录并创建环境变量文件：

```bash
cd tools
cp .env.example .env
# 编辑 .env 填入你的配置信息
```

### 步骤 3: 验证配置

```bash
go run main.go engine.go --validate
```

### 步骤 4: 预览 URL

```bash
go run main.go engine.go --preview
```

### 步骤 5: 推送到搜索引擎

```bash
# 推送到 IndexNow
go run main.go engine.go --indexnow

# 推送到百度搜索
go run main.go engine.go --baidu
```

## 📝 环境变量配置

环境变量文件 `.env` 需要包含以下变量：

```env
INDEXNOW_KEY=your-indexnow-key
INDEXNOW_SITE=https://carbon.go-pkg.com
SITEMAP_PATH=../dist/sitemap.xml
BAIDU_SITE=https://carbon.go-pkg.com
BAIDU_TOKEN=your-baidu-token
```

### 配置说明

- **INDEXNOW_KEY**: IndexNow 密钥字符串
- **INDEXNOW_SITE**: 网站的完整 URL (包含协议)
- **SITEMAP_PATH**: sitemap.xml 文件的相对路径
- **BAIDU_SITE**: 百度搜索提交的网站 URL (必须包含协议)
- **BAIDU_TOKEN**: 百度搜索提交的 token

## 🔧 命令说明

| 命令 | 描述 |
|------|------|
| `go run main.go engine.go --validate` | 验证配置文件 |
| `go run main.go engine.go --preview` | 预览 URL (不提交) |
| `go run main.go engine.go --indexnow` | 提交 URL 到 IndexNow |
| `go run main.go engine.go --baidu` | 提交 URL 到百度搜索 |

## 📁 文件结构

```
tools/
├── engine.go          # 核心逻辑引擎
├── main.go            # 程序入口
├── .env               # 环境变量配置文件
├── .env.example       # 环境变量示例文件
├── .gitignore         # Git 忽略文件
├── go.mod             # Go 模块文件
├── go.sum             # Go 依赖校验文件
├── README.md          # 说明文档
└── urls.txt           # 生成的 URL 文件 (运行时生成)
```

## 🚨 故障排除

### 常见错误

1. **"从 sitemap 读取 URL 失败"**
   - 确保已运行 `npm run build` 生成 sitemap.xml
   - 检查 `dist/sitemap.xml` 文件是否存在

2. **"密钥文件验证失败"**
   - 检查密钥文件是否可以公开访问
   - 确认文件内容只包含密钥字符串

3. **"BAIDU_SITE 必须包含协议"**
   - 确保 BAIDU_SITE 包含 http:// 或 https:// 协议

4. **"over quota" (百度搜索)**
   - 百度搜索提交配额已用完，等待配额重置或联系百度站长平台

## 🔍 技术细节

### 架构设计

- **Engine 结构体**: 封装所有核心逻辑
  - `LoadConfig()`: 加载环境变量配置
  - `ValidateConfig()`: 验证配置完整性
  - `ProcessSitemap()`: 解析 sitemap.xml
  - `GenerateURLs()`: 生成 URL 文件
  - `SubmitToIndexNow()`: 提交到 IndexNow
  - `SubmitToBaidu()`: 提交到百度搜索

- **Main 程序**: 处理命令行参数和用户交互

### API 限制

- **IndexNow**: 每次请求最多 10,000 个 URL (自动分批处理)
- **百度搜索**: 根据百度站长平台配额限制

### 文件格式

- **输入**: XML sitemap 格式
- **输出**: 纯文本格式，每行一个 URL
- **配置**: 环境变量格式

## 📖 相关链接

- [IndexNow 官方文档](https://www.indexnow.org/)
- [百度站长平台](https://ziyuan.baidu.com/linksubmit/index)
- [Carbon 文档网站](https://carbon.go-pkg.com/)

## 🚀 快速开始

```bash
# 1. 构建文档
npm run build

# 2. 进入工具目录
cd tools

# 3. 配置环境变量
cp .env.example .env
# 编辑 .env 文件

# 4. 验证配置
go run main.go engine.go --validate

# 5. 预览 URL
go run main.go engine.go --preview

# 6. 推送到搜索引擎
go run main.go engine.go --indexnow
go run main.go engine.go --baidu
``` 