# Carbon 语言贡献指南

## 如何为 Carbon 添加新的语言支持

### 一、复制语言模板文件

 ```bash
 # 从 lang/en.json 复制作为模板
 cp lang/en.json lang/xx.json
 ```
 其中 `xx` 是您要添加的语言的 `ISO 639-1` 语言代码（如 `zh-CN`、`ja`、`ko` 等）

### 二、更新模板文件内容

 编辑新创建的 `lang/xx.json` 文件，将所有英文内容翻译为目标语言，以下是一个完整的 `简体中文` 语言文件示例：

 ```json
 {
   "name": "简体中文",
   "author": "https://github.com/your-username",
   "months": "一月|二月|三月|四月|五月|六月|七月|八月|九月|十月|十一月|十二月",
   "short_months": "1月|2月|3月|4月|5月|6月|7月|8月|9月|10月|11月|12月",
   "weeks": "星期日|星期一|星期二|星期三|星期四|星期五|星期六",
   "short_weeks": "周日|周一|周二|周三|周四|周五|周六",
   "seasons": "春季|夏季|秋季|冬季",
   "constellations": "白羊座|金牛座|双子座|巨蟹座|狮子座|处女座|天秤座|天蝎座|射手座|摩羯座|水瓶座|双鱼座",
   "year": "%d 年",
   "month": "%d 个月",
   "week": "%d 周",
   "day": "%d 天",
   "hour": "%d 小时",
   "minute": "%d 分钟",
   "second": "%d 秒",
   "now": "刚刚",
   "ago": "%s前",
   "from_now": "%s后",
   "before": "%s前",
   "after": "%s后"
 }
 ```

#### 字段说明

| 字段 | 说明 | 示例 |
|------|------|------|
| `name` | 语言名称（使用该语言的写法） | "简体中文" |
| `author` | 贡献者的 GitHub 链接 | "https://github.com/your-username" |
| `months` | 完整月份名称，用 `\|` 分隔 | "一月\|二月\|三月..." |
| `short_months` | 简短月份名称，用 `\|` 分隔 | "1月\|2月\|3月..." |
| `weeks` | 完整星期名称，用 `\|` 分隔 | "星期日\|星期一\|星期二..." |
| `short_weeks` | 简短星期名称，用 `\|` 分隔 | "周日\|周一\|周二..." |
| `seasons` | 季节名称，用 `\|` 分隔 | "春季\|夏季\|秋季\|冬季" |
| `constellations` | 星座名称，用 `\|` 分隔 | "白羊座\|金牛座\|双子座..." |
| `year` | 年份格式，支持单复数 | "%d 年" |
| `month` | 月份格式，支持单复数 | "%d 个月" |
| `week` | 周格式，支持单复数 | "%d 周" |
| `day` | 天格式，支持单复数 | "%d 天" |
| `hour` | 小时格式，支持单复数 | "%d 小时" |
| `minute` | 分钟格式，支持单复数 | "%d 分钟" |
| `second` | 秒格式，支持单复数 | "%d 秒" |
| `now` | "现在" 的翻译 | "刚刚" |
| `ago` | "之前" 的翻译 | "%s前" |
| `from_now` | "之后" 的翻译 | "%s后" |
| `before` | "之前" 的翻译 | "%s前" |
| `after` | "之后" 的翻译 | "%s后" |

#### 单复数说明

1. **东亚语言（中文、日文、韩文等）**：通常只使用一种格式
   ```json
   "year": "%d 年",
   "month": "%d 个月"
   ```

2. **印欧语言（英文、法文、德文等）**：需要区分单复数
   ```json
   "year": "1 year|%d years",
   "month": "1 month|%d months"
   ```

3. **斯拉夫语言（俄文、乌克兰文等）**：可能有更复杂的复数规则
   ```json
   "year": "1 год|2 года|3 года|4 года|%d лет"
   ```

### 三、提交 Pull Request

1. **创建分支**
   ```bash
   git checkout -b add-xx-language-support
   ```

2. **提交更改**
   ```bash
   git add lang/xx.json
   git commit -m "add XX language support #39"
   ```

3. **推送并创建 Pull Request**
   ```bash
   git push origin add-xx-language-support
   ```

4. **Pull Request 标题格式**
   ```
   Add XX Language Support #39
   ```

### 四、测试验证

提交前请确保：

1. **JSON 格式正确**：使用 `JSON` 验证工具检查语法
2. **字段完整**：确保包含所有必需的 `20` 个字段
3. **分隔符正确**：使用 `|` 作为数组分隔符
4. **占位符正确**：使用 `%d` 作为数字占位符，`%s` 作为字符串占位符
5. **保持一致性**：确保翻译风格与现有语言文件保持一致
6. **文化适应性**：考虑目标语言的文化背景和表达习惯

感谢您为 Carbon 项目贡献新的语言支持！ 
