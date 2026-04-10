package internal

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

type AIEngine struct {
	Client *openai.Client
}

func NewAIEngine(apiKey, baseURL string) *AIEngine {
	config := openai.DefaultConfig(apiKey)
	config.BaseURL = baseURL
	return &AIEngine{
		Client: openai.NewClientWithConfig(config),
	}
}

// 模拟识别逻辑
func (e *AIEngine) SpeechToText(path string) (string, error) {
	/*f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: path,
	}
	ctx := context.Background()
	resp, err := e.Client.CreateTranscription(ctx, req)
	if err != nil {
		return "", fmt.Errorf("语音识别调用API失败：%v", err)
	}
	return resp.Text, nil
	*/
	testSentence := "大家好，我是 AI 助手，现在是一次测试。"

	fmt.Printf(">> [测试模式] 模拟识别内容: %s\n", testSentence)
	return testSentence, nil
}
func (e *AIEngine) Translate(text, from, to string) (string, error) {
	fmt.Printf(">> 正在向 AI 发送请求: [%s]\n", text)
	resp, err := e.Client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "qwen-plus",
			//Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: fmt.Sprintf("你是一个专业的翻译官。请直接将这段%s内容翻译成%s，不要解释，不要保留原文，只输出翻译后的文本：", from, to),
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: text,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
