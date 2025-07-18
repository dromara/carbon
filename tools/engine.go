package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// SitemapURL represents a URL in sitemap.xml
type SitemapURL struct {
	Loc string `xml:"loc"`
}

// Sitemap represents the structure of sitemap.xml
type Sitemap struct {
	XMLName xml.Name     `xml:"urlset"`
	URLs    []SitemapURL `xml:"url"`
}

// Config holds all configuration
type Config struct {
	IndexNowKey  string `json:"indexnowKey"`
	IndexNowSite string `json:"indexnowSite"`
	BaseURL      string `json:"baseURL"`
	SitemapPath  string `json:"sitemapPath"`
	BaiduSite    string `json:"baiduSite"`
	BaiduToken   string `json:"baiduToken"`
}

// IndexNowRequest represents the IndexNow API request structure
type IndexNowRequest struct {
	Host        string   `json:"host"`
	Key         string   `json:"key"`
	KeyLocation string   `json:"keyLocation"`
	URLList     []string `json:"urlList"`
}

// Engine handles the core logic for processing sitemap and generating URL files
type Engine struct {
	config      Config
	sitemapPath string
	outputPath  string
}

// NewEngine creates a new Engine instance
func NewEngine(sitemapPath, outputPath string) *Engine {
	return &Engine{
		sitemapPath: sitemapPath,
		outputPath:  outputPath,
	}
}

// LoadConfig loads configuration from .env file
func (e *Engine) LoadConfig() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}
	e.config.IndexNowKey = os.Getenv("INDEXNOW_KEY")
	e.config.IndexNowSite = os.Getenv("INDEXNOW_SITE")
	e.config.BaseURL = os.Getenv("BASE_URL")
	e.config.SitemapPath = os.Getenv("SITEMAP_PATH")
	e.config.BaiduSite = os.Getenv("BAIDU_SITE")
	e.config.BaiduToken = os.Getenv("BAIDU_TOKEN")

	if e.config.IndexNowKey == "" {
		return fmt.Errorf("INDEXNOW_KEY environment variable is required")
	}
	if e.config.IndexNowSite == "" {
		return fmt.Errorf("INDEXNOW_SITE environment variable is required")
	}
	if e.config.SitemapPath == "" {
		return fmt.Errorf("SITEMAP_PATH environment variable is required")
	}
	if e.config.BaiduSite == "" {
		return fmt.Errorf("BAIDU_SITE environment variable is required")
	}
	if e.config.BaiduToken == "" {
		return fmt.Errorf("BAIDU_TOKEN environment variable is required")
	}
	return nil
}

// ValidateConfig validates the configuration
func (e *Engine) ValidateConfig() error {
	fmt.Println("🔍 验证配置...")
	if e.config.IndexNowKey == "" {
		return fmt.Errorf("❌ INDEXNOW_KEY 字段不能为空")
	}
	if e.config.IndexNowSite == "" {
		return fmt.Errorf("❌ INDEXNOW_SITE 字段不能为空")
	}
	if e.config.BaiduSite == "" {
		return fmt.Errorf("❌ BAIDU_SITE 字段不能为空")
	}
	if e.config.BaiduToken == "" {
		return fmt.Errorf("❌ BAIDU_TOKEN 字段不能为空")
	}
	fmt.Printf("✅ 所有必需字段都已填写\n")
	if !strings.HasPrefix(e.config.IndexNowSite, "http://") && !strings.HasPrefix(e.config.IndexNowSite, "https://") {
		return fmt.Errorf("❌ INDEXNOW_SITE 必须包含协议 (http:// 或 https://)")
	}
	fmt.Printf("✅ INDEXNOW_SITE 格式正确: %s\n", e.config.IndexNowSite)
	// 校验 keyLocation 拼接和可访问性
	keyLocation := e.getIndexNowKeyLocation()
	fmt.Printf("🔍 检查密钥文件可访问性和内容...\n")
	if err := e.validateKeyFile(keyLocation, e.config.IndexNowKey); err != nil {
		return fmt.Errorf("❌ 密钥文件验证失败: %w", err)
	}
	fmt.Printf("✅ 密钥文件可以正常访问且内容正确\n")
	if !strings.HasPrefix(e.config.BaiduSite, "http://") && !strings.HasPrefix(e.config.BaiduSite, "https://") {
		return fmt.Errorf("❌ BAIDU_SITE 必须包含协议 (http:// 或 https://)，当前值: %s", e.config.BaiduSite)
	}
	fmt.Printf("✅ BAIDU_SITE 格式正确: %s\n", e.config.BaiduSite)
	fmt.Println("🎉 配置验证通过！")
	return nil
}

func (e *Engine) getIndexNowHost() string {
	u, err := url.Parse(e.config.IndexNowSite)
	if err != nil {
		return ""
	}
	return u.Host
}

func (e *Engine) getIndexNowKeyLocation() string {
	return strings.TrimRight(e.config.IndexNowSite, "/") + "/" + e.config.IndexNowKey + ".txt"
}

// validateKeyFile checks if the key file is accessible and contains the correct key
func (e *Engine) validateKeyFile(keyLocation, expectedKey string) error {
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

// ProcessSitemap reads sitemap.xml and extracts URLs
func (e *Engine) ProcessSitemap() ([]string, error) {
	// Open sitemap file
	file, err := os.Open(e.sitemapPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open sitemap file: %w", err)
	}
	defer file.Close()

	// Parse XML
	var sitemap Sitemap
	decoder := xml.NewDecoder(file)
	if err := decoder.Decode(&sitemap); err != nil {
		return nil, fmt.Errorf("failed to parse sitemap XML: %w", err)
	}

	// Extract URLs
	var urls []string
	for _, url := range sitemap.URLs {
		if url.Loc != "" {
			urls = append(urls, strings.TrimSpace(url.Loc))
		}
	}

	return urls, nil
}

// SaveURLsToFile saves URLs to the specified output file
func (e *Engine) SaveURLsToFile(urls []string) error {
	file, err := os.Create(e.outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	for _, url := range urls {
		if _, err := file.WriteString(url + "\n"); err != nil {
			return fmt.Errorf("failed to write URL to file: %w", err)
		}
	}

	return nil
}

// GenerateURLs is the main function that processes sitemap and generates urls.txt
func (e *Engine) GenerateURLs() error {
	// Process sitemap
	urls, err := e.ProcessSitemap()
	if err != nil {
		return fmt.Errorf("failed to process sitemap: %w", err)
	}

	if len(urls) == 0 {
		return fmt.Errorf("no URLs found in sitemap")
	}

	// Save URLs to file
	if err := e.SaveURLsToFile(urls); err != nil {
		return fmt.Errorf("failed to save URLs to file: %w", err)
	}

	fmt.Printf("✅ 已从 %s 读取 %d 个 URL，已写入 %s\n", e.sitemapPath, len(urls), e.outputPath)
	return nil
}

// SubmitToIndexNow submits URLs to IndexNow API
func (e *Engine) SubmitToIndexNow(urls []string) error {
	if err := e.ValidateConfig(); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}
	const maxURLsPerRequest = 10000
	for i := 0; i < len(urls); i += maxURLsPerRequest {
		end := i + maxURLsPerRequest
		if end > len(urls) {
			end = len(urls)
		}
		urlChunk := urls[i:end]
		req := IndexNowRequest{
			Host:        e.getIndexNowHost(),
			Key:         e.config.IndexNowKey,
			KeyLocation: e.getIndexNowKeyLocation(),
			URLList:     urlChunk,
		}
		jsonData, err := json.Marshal(req)
		if err != nil {
			return fmt.Errorf("failed to marshal request: %w", err)
		}
		fmt.Printf("Submitting batch %d/%d (%d URLs) to IndexNow...\n", (i/maxURLsPerRequest)+1, (len(urls)-1)/maxURLsPerRequest+1, len(urlChunk))
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

// SubmitToBaidu submits URLs to Baidu search API
func (e *Engine) SubmitToBaidu(urls []string) error {
	// Validate configuration first
	if err := e.ValidateConfig(); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	// Join URLs with newlines
	urlsText := strings.Join(urls, "\n")

	// Create the API URL
	apiURL := fmt.Sprintf("http://data.zz.baidu.com/urls?site=%s&token=%s", e.config.BaiduSite, e.config.BaiduToken)

	fmt.Printf("🚀 开始提交 %d 个 URL 到百度搜索...\n", len(urls))
	fmt.Printf("📡 API URL: %s\n", apiURL)

	// Create HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(urlsText))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "text/plain")

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Submit to Baidu API
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to submit to Baidu: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	fmt.Printf("📊 百度 API 响应状态码: %d\n", resp.StatusCode)
	fmt.Printf("📄 响应内容: %s\n", string(body))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Baidu API returned status %d: %s", resp.StatusCode, string(body))
	}

	fmt.Printf("✅ 成功提交 %d 个 URL 到百度搜索\n", len(urls))
	return nil
}
