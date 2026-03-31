package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/libs/acl/openai"
	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	ctx := context.Background()
	model, err := openai.NewClient(ctx, &openai.Config{APIKey: os.Getenv("API_KEY"),
		Model:   "gpt-5.1",
		BaseURL: os.Getenv("URL"),
	})
	if err != nil {
		panic(err)
	}
	// 准备消息
	messages := []*schema.Message{
		schema.SystemMessage("你是MyGO中的千早爱音"),
		schema.UserMessage("你是谁？"),
	}

	// 生成回答
	response, err := model.Generate(ctx, messages)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Content)
}
