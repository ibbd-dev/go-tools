package es

import (
	"context"
	"fmt"

	"gopkg.in/olivere/elastic.v5"
)

// ImportInit 数据导入初始化
func (c *Client) ImportInit(deleteIndex bool, mapping map[string]interface{}) error {
	ctx := context.Background()
	exists, err := c.es.IndexExists(c.IndexName).Do(ctx)
	if err != nil {
		panic(fmt.Errorf("check index exists error: %v", err.Error()))
	}

	if exists {
		if !deleteIndex {
			// 索引已经存在，而且又不删除索引
			return nil
		}

		// 删除同名索引
		if c.Debug {
			fmt.Printf("begin to delete index: %s\n", c.IndexName)
		}
		if _, err = c.es.DeleteIndex(c.IndexName).Do(ctx); err != nil {
			return fmt.Errorf("delete index: %s, error: %s", c.IndexName, err.Error())
		}
	}

	// 创建索引
	if c.Debug {
		fmt.Printf("begin to create index: %s\n", c.IndexName)
	}
	if _, err = c.es.CreateIndex(c.IndexName).Do(ctx); err != nil {
		return err
	}

	if mapping != nil {
		if c.Debug {
			fmt.Printf("begin to put index mapping: %s\n", c.IndexName)
		}
		if _, err := c.es.PutMapping().Index(c.IndexName).Type(c.DocType).BodyJson(mapping).Do(ctx); err != nil {
			return err
		}
	}

	c.bulk = c.es.Bulk()
	return nil
}

// BulkAdd 增加一行记录
func (c *Client) BulkAdd(row map[string]interface{}) error {
	req := elastic.NewBulkIndexRequest().Index(c.IndexName).Type(c.DocType).Doc(row)
	c.bulk.Add(req)

	c.count++
	if c.count%c.BulkSize == c.BulkSize-1 {
		return c.BulkImport()
	}
	return nil
}

// BulkImport 批量导入
func (c *Client) BulkImport() error {
	// 执行导入
	ctx := context.Background()
	resp, err := c.bulk.Do(ctx)
	if err != nil {
		return fmt.Errorf("index %v 批量导入数据出错: %v", c.IndexName, err.Error())
	}

	// 统计写入状态
	var errCount int
	indexed := resp.Indexed()
	for i, res := range indexed {
		if res.Error != nil {
			fmt.Printf("ERROR: %d, %+v\n", i, res.Error)
			errCount++
		}
	}
	fmt.Printf("向%s写入的数据量：%d，异常：%d\n", c.IndexName, c.count, errCount)
	return nil
}
