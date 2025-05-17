package main

import (
	"context"

	"github.com/cloudwego/eino-ext/components/model/ark"
)

func (r *RAGEngine) newChatModel(ctx context.Context) {
	m, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: r.config.ApiKey,
		Model:  r.config.ChatModel,
	})
	if err != nil {
		r.Err = err
		return
	}
	r.ChatModel = m
}
