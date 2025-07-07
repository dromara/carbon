package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// IndexNowRequest represents the IndexNow API request structure
type IndexNowRequest struct {
	Host        string   `json:"host"`
	Key         string   `json:"key"`
	KeyLocation string   `json:"keyLocation"`
	URLList     []string `json:"urlList"`
}

// IndexNowResponse represents the IndexNow API response
type IndexNowResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Config holds the IndexNow configuration
type Config struct {
	Host        string `json:"host"`
	Key         string `json:"key"`
	KeyLocation string `json:"keyLocation"`
	BaseURL     string `json:"baseURL"`
	SitemapPath string `json:"sitemapPath"`
}

// Sitemap XML structures
type Urlset struct {
	XMLName xml.Name `xml:"urlset"`
	URLs    []URL    `xml:"url"`
}

type URL struct {
	Loc string `xml:"loc"`
}

// generateURLsFromSitemap reads URLs from the sitemap.xml file
func generateURLsFromSitemap(sitemapPath string) ([]string, error) {
	// Read sitemap file
	file, err := os.Open(sitemapPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open sitemap file: %w", err)
	}
	defer file.Close()

	// Parse XML
	var urlset Urlset
	decoder := xml.NewDecoder(file)
	if err := decoder.Decode(&urlset); err != nil {
		return nil, fmt.Errorf("failed to parse sitemap XML: %w", err)
	}

	// Extract URLs
	var urls []string
	for _, url := range urlset.URLs {
		if url.Loc != "" {
			urls = append(urls, url.Loc)
		}
	}

	return urls, nil
}

// validateKeyFile checks if the key file is accessible and contains the correct key
func validateKeyFile(keyLocation, expectedKey string) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(keyLocation)
	if err != nil {
		return fmt.Errorf("failed to access key file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("key file returned status %d", resp.StatusCode)
	}

	// Read the key file content
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read key file content: %w", err)
	}

	// Trim whitespace and compare with expected key
	actualKey := strings.TrimSpace(string(body))
	if actualKey != expectedKey {
		return fmt.Errorf("key file content mismatch. Expected: %s, Got: %s", expectedKey, actualKey)
	}

	return nil
}

// validateConfig validates the IndexNow configuration
func validateConfig(config Config) error {
	fmt.Println("🔍 验证 IndexNow 配置...")

	// Check required fields
	if config.Host == "" {
		return fmt.Errorf("❌ host 字段不能为空")
	}
	if config.Key == "" {
		return fmt.Errorf("❌ key 字段不能为空")
	}
	if config.KeyLocation == "" {
		return fmt.Errorf("❌ keyLocation 字段不能为空")
	}
	if config.BaseURL == "" {
		return fmt.Errorf("❌ baseURL 字段不能为空")
	}

	fmt.Printf("✅ 所有必需字段都已填写\n")

	// Validate host format
	if strings.Contains(config.Host, "http://") || strings.Contains(config.Host, "https://") {
		return fmt.Errorf("❌ host 不应包含协议，应该是: %s", strings.TrimPrefix(strings.TrimPrefix(config.Host, "https://"), "http://"))
	}
	fmt.Printf("✅ host 格式正确: %s\n", config.Host)

	// Validate keyLocation belongs to host
	if !strings.Contains(config.KeyLocation, config.Host) {
		return fmt.Errorf("❌ keyLocation URL 必须属于 host 域名\n   当前 host: %s\n   当前 keyLocation: %s", config.Host, config.KeyLocation)
	}
	fmt.Printf("✅ keyLocation 域名匹配: %s\n", config.KeyLocation)

	// Validate key file accessibility and content
	fmt.Printf("🔍 检查密钥文件可访问性和内容...\n")
	if err := validateKeyFile(config.KeyLocation, config.Key); err != nil {
		return fmt.Errorf("❌ 密钥文件验证失败: %w", err)
	}
	fmt.Printf("✅ 密钥文件可以正常访问且内容正确\n")

	// Validate baseURL format
	if !strings.HasPrefix(config.BaseURL, "http://") && !strings.HasPrefix(config.BaseURL, "https://") {
		return fmt.Errorf("❌ baseURL 必须包含协议 (http:// 或 https://)")
	}
	fmt.Printf("✅ baseURL 格式正确: %s\n", config.BaseURL)

	fmt.Println("🎉 配置验证通过！")
	return nil
}

// submitToIndexNow submits URLs to IndexNow API
func submitToIndexNow(config Config, urls []string) error {
	// Validate configuration first
	if err := validateConfig(config); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	// Split URLs into chunks of 10,000 (IndexNow limit)
	const maxURLsPerRequest = 10000

	for i := 0; i < len(urls); i += maxURLsPerRequest {
		end := i + maxURLsPerRequest
		if end > len(urls) {
			end = len(urls)
		}

		urlChunk := urls[i:end]

		request := IndexNowRequest{
			Host:        config.Host,
			Key:         config.Key,
			KeyLocation: config.KeyLocation,
			URLList:     urlChunk,
		}

		jsonData, err := json.Marshal(request)
		if err != nil {
			return fmt.Errorf("failed to marshal request: %w", err)
		}

		fmt.Printf("Submitting batch %d/%d (%d URLs) to IndexNow...\n", (i/maxURLsPerRequest)+1, (len(urls)-1)/maxURLsPerRequest+1, len(urlChunk))

		// Submit to IndexNow API
		resp, err := http.Post("https://api.indexnow.org/IndexNow", "application/json; charset=utf-8", bytes.NewBuffer(jsonData))
		if err != nil {
			return fmt.Errorf("failed to submit to IndexNow: %w", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("IndexNow API returned status %d: %s", resp.StatusCode, string(body))
		}

		fmt.Printf("Successfully submitted %d URLs to IndexNow (batch %d/%d)\n", len(urlChunk), (i/maxURLsPerRequest)+1, (len(urls)-1)/maxURLsPerRequest+1)
	}

	return nil
}

// loadConfig loads configuration from .env file
func loadConfig(envPath string) (Config, error) {
	var config Config

	// Load .env file
	if err := godotenv.Load(envPath); err != nil {
		return config, fmt.Errorf("failed to load .env file: %w", err)
	}

	// Read environment variables
	config.Host = os.Getenv("HOST")
	config.Key = os.Getenv("KEY")
	config.KeyLocation = os.Getenv("KEY_LOCATION")
	config.BaseURL = os.Getenv("BASE_URL")
	config.SitemapPath = os.Getenv("SITEMAP_PATH")

	// Validate that all required fields are present
	if config.Host == "" {
		return config, fmt.Errorf("HOST environment variable is required")
	}
	if config.Key == "" {
		return config, fmt.Errorf("KEY environment variable is required")
	}
	if config.KeyLocation == "" {
		return config, fmt.Errorf("KEY_LOCATION environment variable is required")
	}
	if config.BaseURL == "" {
		return config, fmt.Errorf("BASE_URL environment variable is required")
	}
	if config.SitemapPath == "" {
		return config, fmt.Errorf("SITEMAP_PATH environment variable is required")
	}

	return config, nil
}

// saveURLsToFile saves generated URLs to a file for review
func saveURLsToFile(urls []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	for _, url := range urls {
		if _, err := file.WriteString(url + "\n"); err != nil {
			return fmt.Errorf("failed to write URL: %w", err)
		}
	}

	fmt.Printf("Saved %d URLs to %s\n", len(urls), filename)
	return nil
}

func main() {
	var envPath string
	var mode string

	// Parse command line arguments
	if len(os.Args) < 2 {
		// No arguments provided, use default .env file
		envPath = ".env"
	} else if strings.HasPrefix(os.Args[1], "--") {
		// First argument is a mode flag, use default .env file
		envPath = ".env"
		mode = os.Args[1]
	} else {
		// First argument is env file path
		envPath = os.Args[1]
		if len(os.Args) > 2 {
			mode = os.Args[2]
		}
	}

	// Show usage if no mode is specified
	if mode == "" {
		fmt.Println("用法: go run indexnow.go [.env] [--preview|--validate|--push]")
		fmt.Println("")
		fmt.Println("模式:")
		fmt.Println("  --preview    预览生成的 URL (不提交)")
		fmt.Println("  --validate   验证配置文件")
		fmt.Println("  --push       生成 URL 并提交到 IndexNow")
		fmt.Println("")
		fmt.Println("环境变量文件 (.env) 应包含:")
		fmt.Println("HOST=carbon.go-pkg.com")
		fmt.Println("KEY=your-indexnow-key")
		fmt.Println("KEY_LOCATION=https://carbon.go-pkg.com/key.txt")
		fmt.Println("BASE_URL=https://carbon.go-pkg.com")
		fmt.Println("SITEMAP_PATH=../dist/sitemap.xml")
		fmt.Println("")
		fmt.Println("示例:")
		fmt.Println("  go run indexnow.go --validate")
		fmt.Println("  go run indexnow.go --preview")
		fmt.Println("  go run indexnow.go --push")
		fmt.Println("  go run indexnow.go .env --validate")
		os.Exit(1)
	}

	// Load configuration
	config, err := loadConfig(envPath)
	if err != nil {
		fmt.Printf("❌ 加载环境变量文件失败: %v\n", err)
		os.Exit(1)
	}

	// Handle different modes
	switch mode {
	case "--validate":
		if err := validateConfig(config); err != nil {
			fmt.Printf("❌ 配置验证失败: %v\n", err)
			fmt.Println("")
			fmt.Println("💡 常见问题解决方案:")
			fmt.Println("1. 确保 keyLocation 的域名与 host 完全匹配")
			fmt.Println("2. 确保密钥文件在指定位置可以公开访问")
			fmt.Println("3. 检查密钥文件内容是否正确 (只包含密钥字符串)")
			fmt.Println("4. 确保 host 不包含协议 (http:// 或 https://)")
			os.Exit(1)
		}
		fmt.Println("")
		fmt.Println("🚀 配置验证成功！现在可以运行:")
		fmt.Printf("  go run indexnow.go %s --preview\n", envPath)
		fmt.Printf("  go run indexnow.go %s --push\n", envPath)
		return

	case "--preview":
		fmt.Println("🔍 预览模式 - 生成 URL 但不提交")

		// Generate URLs from sitemap.xml
		urls, err := generateURLsFromSitemap(config.SitemapPath)
		if err != nil {
			fmt.Printf("❌ 从 sitemap 读取 URL 失败: %v\n", err)
			fmt.Printf("💡 请确保已运行 `npm run build` 生成 sitemap.xml 或检查 SITEMAP_PATH 配置: %s\n", config.SitemapPath)
			os.Exit(1)
		}

		fmt.Printf("✅ 从 %s 读取了 %d 个 URL\n", config.SitemapPath, len(urls))

		// Save URLs to file for review
		if err := saveURLsToFile(urls, "preview_urls.txt"); err != nil {
			fmt.Printf("❌ 保存 URL 失败: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("📄 URL 已保存到 preview_urls.txt")
		fmt.Println("🔍 请检查生成的 URL，确认无误后运行:")
		fmt.Printf("  go run indexnow.go %s --push\n", envPath)
		return

	case "--push":
		// Normal submission mode
		fmt.Println("🚀 开始提交 URL 到 IndexNow...")

		// Generate URLs from sitemap.xml
		urls, err := generateURLsFromSitemap(config.SitemapPath)
		if err != nil {
			fmt.Printf("❌ 从 sitemap 读取 URL 失败: %v\n", err)
			fmt.Printf("💡 请确保已运行 `npm run build` 生成 sitemap.xml 或检查 SITEMAP_PATH 配置: %s\n", config.SitemapPath)
			os.Exit(1)
		}

		fmt.Printf("✅ 从 %s 读取了 %d 个 URL\n", config.SitemapPath, len(urls))

		// Save URLs to file for review
		if err := saveURLsToFile(urls, "preview_urls.txt"); err != nil {
			fmt.Printf("❌ 保存 URL 失败: %v\n", err)
			os.Exit(1)
		}

		// Submit to IndexNow
		if err := submitToIndexNow(config, urls); err != nil {
			fmt.Printf("❌ 提交到 IndexNow 失败: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("🎉 成功提交所有 URL 到 IndexNow！")

	case "":
		fmt.Println("❌ 请指定运行模式")
		fmt.Println("支持的模式: --preview, --validate, --push")
		os.Exit(1)

	default:
		fmt.Printf("❌ 未知模式: %s\n", mode)
		fmt.Println("支持的模式: --preview, --validate, --push")
		os.Exit(1)
	}
}
