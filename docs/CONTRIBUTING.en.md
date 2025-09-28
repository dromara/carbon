# Carbon Language Contribution Guide

## How to Add New Language Support to Carbon

### Step 1: Copy Language Template File

```bash
# Copy from lang/en.json as template
cp lang/en.json lang/xx.json
```
Where `xx` is the `ISO 639-1` language code for the language you want to add (e.g., `zh-CN`, `ja`, `ko`, etc.)

### Step 2: Update Template File Content

Edit the newly created `lang/xx.json` file and translate all English content to the target language. Here is a complete example of a `English` language file:

```json
{
  "name": "English",
  "author": "https://github.com/gouguoyin",
  "months": "January|February|March|April|May|June|July|August|September|October|November|December",
  "short_months": "Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec",
  "weeks": "Sunday|Monday|Tuesday|Wednesday|Thursday|Friday|Saturday",
  "short_weeks": "Sun|Mon|Tue|Wed|Thu|Fri|Sat",
  "seasons": "Spring|Summer|Autumn|Winter",
  "constellations": "Aries|Taurus|Gemini|Cancer|Leo|Virgo|Libra|Scorpio|Sagittarius|Capricorn|Aquarius|Pisces",
  "year": "1 year|%d years",
  "month": "1 month|%d months",
  "week": "1 week|%d weeks",
  "day": "1 day|%d days",
  "hour": "1 hour|%d hours",
  "minute": "1 minute|%d minutes",
  "second": "1 second|%d seconds",
  "now": "just now",
  "ago": "%s ago",
  "from_now": "%s from now",
  "before": "%s before",
  "after": "%s after"
}
```

#### Field Description

| Field | Description | Example |
|-------|-------------|---------|
| `name` | Language name (written in that language) | "English" |
| `author` | Contributor's GitHub link | "https://github.com/gouguoyin" |
| `months` | Full month names, separated by `\|` | "January\|February\|March..." |
| `short_months` | Short month names, separated by `\|` | "Jan\|Feb\|Mar..." |
| `weeks` | Full week names, separated by `\|` | "Sunday\|Monday\|Tuesday..." |
| `short_weeks` | Short week names, separated by `\|` | "Sun\|Mon\|Tue..." |
| `seasons` | Season names, separated by `\|` | "Spring\|Summer\|Autumn\|Winter" |
| `constellations` | Constellation names, separated by `\|` | "Aries\|Taurus\|Gemini..." |
| `year` | Year format, supports singular/plural | "1 year\|%d years" |
| `month` | Month format, supports singular/plural | "1 month\|%d months" |
| `week` | Week format, supports singular/plural | "1 week\|%d weeks" |
| `day` | Day format, supports singular/plural | "1 day\|%d days" |
| `hour` | Hour format, supports singular/plural | "1 hour\|%d hours" |
| `minute` | Minute format, supports singular/plural | "1 minute\|%d minutes" |
| `second` | Second format, supports singular/plural | "1 second\|%d seconds" |
| `now` | Translation of "now" | "just now" |
| `ago` | Translation of "ago" | "%s ago" |
| `from_now` | Translation of "from now" | "%s from now" |
| `before` | Translation of "before" | "%s before" |
| `after` | Translation of "after" | "%s after" |

#### Singular/Plural Handling

Different languages handle singular/plural forms differently:

1. **East Asian languages (Chinese, Japanese, Korean, etc.)**: Usually use only one format
   ```json
   "year": "%d 年",
   "month": "%d 个月"
   ```

2. **Indo-European languages (English, French, German, etc.)**: Need to distinguish singular/plural
   ```json
   "year": "1 year|%d years",
   "month": "1 month|%d months"
   ```

3. **Slavic languages (Russian, Ukrainian, etc.)**: May have more complex plural rules
   ```json
   "year": "1 год|2 года|3 года|4 года|%d лет"
   ```

### Step 3: Submit Pull Request

1. **Create Branch**
   ```bash
   git checkout -b add-xx-language-support
   ```

2. **Commit Changes**
   ```bash
   git add lang/xx.json
   git commit -m "add XX language support #39"
   ```

3. **Push and Create Pull Request**
   ```bash
   git push origin add-xx-language-support
   ```

4. **Pull Request Title Format**
   ```
   Add XX Language Support #39
   ```

### Step 4: Test Verification

Before submitting, please ensure:

1. **Correct JSON format**: Use `JSON` validation tools to check syntax
2. **Complete fields**: Ensure all required `20` fields are included
3. **Correct separators**: Use `|` as array separator
4. **Correct placeholders**: Use `%d` as number placeholder, `%s` as string placeholder
5. **Maintain consistency**: Ensure translation style is consistent with existing language files
6. **Cultural adaptation**: Consider the cultural background and expression habits of the target language

Thank you for contributing new language support to the Carbon project!
