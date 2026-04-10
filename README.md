# 	Mini-TMK-Agent: 语音同传翻译助手
本项目是一个基于GO语言开发的AI翻译Agent示例。它模拟了同传设备的核心逻辑：通过**ASR（语音获取文字）** 获取音频内容，并利用 **LLM（大语言模型）** 进行实时多语言翻译。支持文件转录、流式同传模拟以及 Web 交互界面。

## 功能特性
-    **流式同传模式**：通过麦克风实时录音（3 秒）并翻译。
-    **文件转录模式**：解析 WAV 音频文件并完成翻译。
-    **多语言支持**：支持中文翻译至英语、日语、韩语（可扩展）。
-    **Web 可视化界面**：简洁的网页操作界面，无需命令行操作。
-    **测试模式**：内置模拟识别逻辑，无需真实音频文件即可测试。

## 环境要求
-   Go 1.18+
-   操作系统：Windows
-   网络环境：可访问阿里云通义千问 API

##	快速开始
### 1. 配置环境变量

file.env 文件，填写阿里云通义千问的 API Key 和 Base URL：
```
AI_API_KEY=你的通义千问API密钥 
AI_BASE_URL=https://dashscope.aliyuncs.com/compatible-mode/v1
```
### 2. 安装依赖
具体看go.mod文档。

### 3.启动服务
```
go run main.go web
```
## 使用指南 
### 文件转录模式:
在 Web 界面选择「文件转录模式」 输入待翻译的 WAV 音频文件名（默认：test.wav） 选择目标语言（英语 / 日语 / 韩语） 点击「开始翻译」按钮，等待结果展示 。
### 流式同传模式:
 在 Web 界面选择「流式同传模式」 选择目标语言 点击「开启麦克风录音」按钮 程序会录制 3 秒音频并自动完成识别与翻译（当前为测试模式，使用固定文本）。
 
## 项目结构
```
mini-tmk-agent/
├── main.go             # 程序入口，处理路由分发与 CLI 逻辑
├── file.env            # 环境变量配置 (API Key & URL)
├── static/
│   └── index.html      # 前端 Web 交互界面
└── internal/
    ├── ai.go           # AI 引擎封装 (ASR 与 Chat 翻译)
    ├── audio.go        # 音频文件辅助工具
    └── recorder.go     # 本地麦克风录音逻辑 (基于 PortAudio)
 ```

 ## 扩展开发:
  1. 启用真实语音识别 取消 internal/ai.go 中 SpeechToText 方法的注释，启用真实的 OpenAI Whisper 语音识别：
  ```
  // 取消以下代码注释
   f, err := os.Open(path) 
   if err != nil { return "", err } 
   defer f.Close()
   req := openai.AudioRequest{ Model: openai.Whisper1, FilePath: path, } 
   ctx := context.Background() 
   resp, err := e.Client.CreateTranscription(ctx, req) 
   if err != nil { return "", fmt.Errorf("语音识别调用API失败：%v", err) } 
   return resp.Text, nil
 ```
### 2. 启用真实麦克风录音
取消 `internal/recorder.go` 的注释，并安装 PortAudio 依赖。

## 注意事项： 
1.API Key 需妥善保管，请勿泄露或提交到代码仓库 。
2.真实录音功能需要安装 PortAudio 依赖，不同系统安装方式不同 。
3.测试模式下不会调用真实 API，仅返回模拟文本。
4. 音频文件仅支持 WAV 格式，且需符合 Whisper 模型要求（16kHz 采样率）。

## 常见问题：
 Q1: 启动服务提示 AI_API_KEY 未设置？
A1: ```file.env.env文件是否存在，且AI_API_KEY和AI_BASE_URL配置正确。```
  Q2: Web 界面访问失败？ 
  A2: ```确认服务已启动（go run main.go web），且端口 8080 未被占用。```
   Q3: 翻译接口返回错误？ 
   A3: ```检查 API Key 有效性、网络连通性，以及 Base URL 是否正确（阿里云通义千问需使用兼容模式地址）。```