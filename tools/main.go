package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		showUsage()
		os.Exit(1)
	}

	mode := os.Args[1]

	// Create engine instance
	engine := NewEngine("../dist/sitemap.xml", "urls.txt")

	// Load configuration
	if err := engine.LoadConfig(); err != nil {
		fmt.Printf("❌ 加载配置失败: %v\n", err)
		os.Exit(1)
	}

	switch mode {
	case "--validate":
		if err := engine.ValidateConfig(); err != nil {
			fmt.Printf("❌ 配置验证失败: %v\n", err)
			fmt.Println("")
			fmt.Println("💡 常见问题解决方案:")
			fmt.Println("1. 确保 keyLocation 的域名与 host 完全匹配")
			fmt.Println("2. 确保密钥文件在指定位置可以公开访问")
			fmt.Println("3. 检查密钥文件内容是否正确 (只包含密钥字符串)")
			fmt.Println("4. 确保 host 不包含协议 (http:// 或 https://)")
			fmt.Println("5. 确保 BAIDU_SITE 包含协议 (http:// 或 https://)")
			os.Exit(1)
		}
		fmt.Println("")
		fmt.Println("🚀 配置验证成功！现在可以运行:")
		fmt.Println("  go run main.go engine.go --preview")
		fmt.Println("  go run main.go engine.go --indexnow")
		fmt.Println("  go run main.go engine.go --baidu")

	case "--preview":
		fmt.Println("🔍 预览模式 - 生成 URL 但不提交")
		if err := engine.GenerateURLs(); err != nil {
			fmt.Printf("❌ 生成 URL 失败: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("🎉 你可以继续用 --indexnow 或 --baidu 推送这些 URL")

	case "--indexnow":
		fmt.Println("🚀 开始提交 URL 到 IndexNow...")

		// Generate URLs from sitemap
		urls, err := engine.ProcessSitemap()
		if err != nil {
			fmt.Printf("❌ 从 sitemap 读取 URL 失败: %v\n", err)
			fmt.Println("💡 请确保已运行 `npm run build` 生成 sitemap.xml")
			os.Exit(1)
		}

		fmt.Printf("✅ 从 sitemap.xml 读取了 %d 个 URL\n", len(urls))

		// Save URLs to file for review
		if err := engine.SaveURLsToFile(urls); err != nil {
			fmt.Printf("❌ 保存 URL 失败: %v\n", err)
			os.Exit(1)
		}

		// Submit to IndexNow
		if err := engine.SubmitToIndexNow(urls); err != nil {
			fmt.Printf("❌ 提交到 IndexNow 失败: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("🎉 成功提交所有 URL 到 IndexNow！")

	case "--baidu":
		fmt.Println("🚀 开始提交 URL 到百度搜索...")

		// Generate URLs from sitemap
		urls, err := engine.ProcessSitemap()
		if err != nil {
			fmt.Printf("❌ 从 sitemap 读取 URL 失败: %v\n", err)
			fmt.Println("💡 请确保已运行 `npm run build` 生成 sitemap.xml")
			os.Exit(1)
		}

		fmt.Printf("✅ 从 sitemap.xml 读取了 %d 个 URL\n", len(urls))

		// Save URLs to file for review
		if err := engine.SaveURLsToFile(urls); err != nil {
			fmt.Printf("❌ 保存 URL 失败: %v\n", err)
			os.Exit(1)
		}

		// Submit to Baidu
		if err := engine.SubmitToBaidu(urls); err != nil {
			fmt.Printf("❌ 提交到百度搜索失败: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("🎉 成功提交所有 URL 到百度搜索！")

	default:
		fmt.Printf("❌ 未知模式: %s\n", mode)
		showUsage()
		os.Exit(1)
	}
}

func showUsage() {
	fmt.Println("用法: go run main.go engine.go [--validate|--preview|--indexnow|--baidu]")
	fmt.Println("")
	fmt.Println("模式:")
	fmt.Println("  --validate   验证配置文件")
	fmt.Println("  --preview    预览生成的 URL (不提交)")
	fmt.Println("  --indexnow   生成 URL 并提交到 IndexNow")
	fmt.Println("  --baidu      生成 URL 并提交到百度搜索")
	fmt.Println("")
	fmt.Println("环境变量文件 (.env) 应包含:")
	fmt.Println("INDEXNOW_KEY=your-indexnow-key")
	fmt.Println("INDEXNOW_BASE_URL=https://carbon.go-pkg.com")
	fmt.Println("SITEMAP_PATH=../dist/sitemap.xml")
	fmt.Println("BAIDU_SITE=https://carbon.go-pkg.com")
	fmt.Println("BAIDU_TOKEN=your-baidu-token")
	fmt.Println("")
	fmt.Println("示例:")
	fmt.Println("  go run main.go engine.go --validate")
	fmt.Println("  go run main.go engine.go --preview")
	fmt.Println("  go run main.go engine.go --indexnow")
	fmt.Println("  go run main.go engine.go --baidu")
}
