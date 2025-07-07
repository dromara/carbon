# IndexNow URL 提交工具

这个工具可以自动将 Carbon 文档网站的所有页面 URL 提交到 IndexNow API，帮助搜索引擎快速发现和索引新内容。

## 📋 功能特性

- ✅ **自动读取 URL**: 从 `dist/sitemap.xml` 自动读取所有构建后的页面 URL
- ✅ **配置验证**: 验证 IndexNow 配置的正确性和密钥文件可访问性
- ✅ **预览模式**: 预览将要提交的 URL 而不实际提交
- ✅ **批量提交**: 自动处理大量 URL 的批量提交 (每批最多 10,000 个)
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
go run indexnow.go --validate
```

### 步骤 4: 预览 URL

```bash
go run indexnow.go --preview
```

### 步骤 5: 提交到 IndexNow

```bash
go run indexnow.go --push
```

## 📝 环境变量配置

环境变量文件 `.env` 需要包含以下变量：

```env
HOST=carbon.go-pkg.com
KEY=your-indexnow-key
KEY_LOCATION=https://carbon.go-pkg.com/your-key.txt
BASE_URL=https://carbon.go-pkg.com
SITEMAP_PATH=../dist/sitemap.xml
```

### 配置说明

- **HOST**: 网站域名 (不包含协议)
- **KEY**: IndexNow 密钥字符串
- **KEY_LOCATION**: 密钥文件的完整 URL (必须与 HOST 域名匹配)
- **BASE_URL**: 网站的完整 URL (包含协议)
- **SITEMAP_PATH**: sitemap.xml 文件的相对路径

## 🔧 命令说明

| 命令 | 描述 |
|------|------|
| `go run indexnow.go --validate` | 验证配置文件 |
| `go run indexnow.go --preview` | 预览 URL (不提交) |
| `go run indexnow.go --push` | 提交 URL 到 IndexNow |

## 📁 输出文件

- `tools/preview_urls.txt`: 包含所有从 sitemap.xml 读取的 URL
- `tools/.env`: IndexNow 环境变量配置文件

## 🚨 故障排除

### 常见错误

1. **"从 sitemap 读取 URL 失败"**
   - 确保已运行 `npm run build` 生成 sitemap.xml
   - 检查 `dist/sitemap.xml` 文件是否存在

2. **"keyLocation URL 必须属于 host 域名"**
   - 确保 keyLocation 的域名与 host 完全匹配

3. **"密钥文件验证失败"**
   - 检查密钥文件是否可以公开访问
   - 确认文件内容只包含密钥字符串

## 🔍 技术细节

- **URL 来源**: 从构建后的 `dist/sitemap.xml` 读取所有页面 URL
- **提交限制**: 每次请求最多 10,000 个 URL (自动分批处理)
- **重试机制**: 内置错误处理和详细错误信息
- **文件格式**: 生成的 URL 列表保存为纯文本格式，每行一个 URL

## 📖 相关链接

- [IndexNow 官方文档](https://www.indexnow.org/)
- [Carbon 文档网站](https://carbon.go-pkg.com/) 