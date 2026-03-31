package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/cloudwego/eino-ext/libs/acl/openai"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	ctx := context.Background()
	embedder, err := openai.NewEmbeddingClient(ctx, &openai.EmbeddingConfig{
		APIKey:     os.Getenv("EMBEDDING_API_KEY"),
		Model:      "Qwen/Qwen3-Embedding-8B",
		BaseURL:    os.Getenv("EM_URL"),
		HTTPClient: http.DefaultClient,
	})
	if err != nil {
		panic(err)
	}
	inputs := []string{
		"我喜欢青木阳菜",
		"我喜欢千早爱音",
	}
	response, err := embedder.EmbedStrings(ctx, inputs)
	if err != nil {
		panic(err)
	}
	for i, x := range response {
		fmt.Printf("输入: %s\n", inputs[i])
		fmt.Printf("向量: %v\n", len(x))
	}
}
