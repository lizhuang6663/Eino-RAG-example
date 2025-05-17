package main

import (
	"context"
	"time"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
)

// 使用向量模型将文本向量化

func main() {
	ctx := context.Background()

	// 初始化嵌入器
	timeout := 30 * time.Second
	embedder, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIKey:  "7ce6c29b-359c-41fb-995b-72b91ca00317",
		Model:   "doubao-embedding-large-text-240915",
		Timeout: &timeout,
	})
	if err != nil {
		panic(err)
	}

	// 生成文本向量
	texts := []string{
		"这是第一段示例文本",
		"这是第二段示例文本",
	}

	embeddings, err := embedder.EmbedStrings(ctx, texts)
	if err != nil {
		panic(err)
	}

	// 使用生成的向量
	for i, embedding := range embeddings {
		// fmt.Println(embedding)
		println("文本", i+1, "的向量维度:", len(embedding))
	}

}
