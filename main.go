package main

import (
	"fmt"
	"log"
	"mini-tmk-agent/internal"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	_ = godotenv.Load("file.env")
	apiKey := os.Getenv("AI_API_KEY")
	baseURL := os.Getenv("AI_BASE_URL")

	if apiKey == "" || baseURL == "" {
		log.Fatal("错误：AI_API_KEY 或 AI_BASE_URL 未设置，请检查 file.env")
	}

	engine := internal.NewAIEngine(apiKey, baseURL)

	var rootCmd = &cobra.Command{Use: "mini-tmk-agent"}

	// Web UI 命令
	var webCmd = &cobra.Command{
		Use: "web",
		Run: func(cmd *cobra.Command, args []string) {
			http.Handle("/", http.FileServer(http.Dir("./static")))

			http.HandleFunc("/api/translate", func(w http.ResponseWriter, r *http.Request) {
				mode := r.URL.Query().Get("mode")     // 获取模式: file 或 stream
				fileName := r.URL.Query().Get("file") // 文件名
				target := r.URL.Query().Get("target") // 目标语言

				var originalText string
				var err error

				if mode == "stream" {
					// 流式同传逻辑：调用录音并识别
					/*
						tempFile := "web_capture.wav"
						fmt.Println(">> 正在通过网页触发麦克风录音...")
						err = internal.RecordToFile(3, tempFile)
						if err != nil {
							sendError(w, fmt.Sprintf("录音失败: %v (请检查设备权限或 PortAudio 安装)", err))
							return
						}
						originalText, err = engine.SpeechToText(tempFile)
					*/
					// 测试模式：直接使用固定文本模拟识别结果
					originalText, err = engine.SpeechToText("")
				} else {
					// 2. 文件转录逻辑：识别指定文件
					if fileName == "" {
						fileName = "test.wav"
					}
					fmt.Printf(">> 正在处理网页请求文件: %s\n", fileName)
					originalText, err = engine.SpeechToText(fileName)
				}

				if err != nil {
					sendError(w, fmt.Sprintf("识别失败: %v", err))
					return
				}

				//执行翻译
				translated, err := engine.Translate(originalText, "中文", target)
				if err != nil {
					sendError(w, fmt.Sprintf("翻译失败: %v", err))
					return
				}

				// 返回 JSON
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, `{"original": "%s", "translated": "%s"}`, originalText, translated)
			})

			fmt.Println(">> Web 服务已启动: http://localhost:8080")
			log.Fatal(http.ListenAndServe(":8080", nil))
		},
	}

	rootCmd.AddCommand(webCmd)
	rootCmd.Execute()
}

func sendError(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	fmt.Fprintf(w, `{"error": "%s"}`, msg)
}
