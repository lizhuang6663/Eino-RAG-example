// 使用字节跳动开源的大语言模型应用框架 Eino 构建一个基于 Redis Stack 向量数据库 的简单 RAG 应用。

package main

import (
	"context"
	"fmt"
	"io"

	"github.com/cloudwego/eino/components/document"
	uuid2 "github.com/google/uuid"
)

const (
	prefix = "OuterCyrex:"
	index  = "OuterIndex"
)

func main() {
	ctx := context.Background()

	// 初始化RAG：组装RAGEngine
	r, err := InitRAGEngine(ctx, index, prefix)
	if err != nil {
		panic(err)
	}

	// 读取对应URI文件的内容，将其以[]*schema.Document形式返回
	doc, err := r.Loader.Load(ctx, document.Source{
		URI: "./test_txt/mysql-1.md",
	})
	if err != nil {
		panic(err)
	}

	// 按照Splitter初始化时的分割策略进行分割，仍然返回[]*schema.Document
	// 此前我们设置 "#":"title"，则其会以 # 符号为分隔符来进行文本切割。
	docs, err := r.Splitter.Transform(ctx, doc)
	if err != nil {
		panic(err)
	}

	for _, d := range docs {
		uuid, _ := uuid2.NewUUID()
		d.ID = uuid.String()
	}

	// 向量嵌入(Embedding Vector)
	// 向量嵌入是一种将复杂、高维或稀疏的数据（如文本、图像、分类特征等）通过嵌入函数映射到低维、稠密向量空间的技术。
	// 我们将这个过程交给embedder即可，而我们实际要做的是将embedder返回的向量存储在本地向量数据库中，便于后续进行检索。

	// 建立索引
	err = r.InitVectorIndex(ctx)
	if err != nil {
		panic(err)
	}

	// 将文档向量化后存储在向量数据库中，方便后续查询
	_, err = r.Indexer.Store(ctx, docs)
	if err != nil {
		panic(err)
	}

	var query string

	for {
		_, _ = fmt.Scan(&query)
		output, err := r.Stream(ctx, query)
		if err != nil {
			panic(err)
		}
		for {
			o, err := output.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			fmt.Print(o.Content)
		}
	}
}
