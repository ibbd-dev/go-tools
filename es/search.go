package es

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"gopkg.in/olivere/elastic.v5"
)

// SearchInit 搜索初始化
func (c *Client) SearchInit(query elastic.Query) error {
	ctx := context.Background()
	exists, err := c.es.IndexExists(c.IndexName).Do(ctx)
	if err != nil {
		return fmt.Errorf("check index exists error: %v", err.Error())
	}

	if !exists {
		return fmt.Errorf("index %s is not exists", c.IndexName)
	}

	c.searchResult, err = c.es.Search(c.IndexName).Query(query).Size(c.BulkSize).Do(ctx)
	if err != nil {
		return fmt.Errorf("search do error: %v", err.Error())
	}

	c.sResTotal = c.searchResult.Hits.TotalHits
	fmt.Printf("search research total: %d\n", c.sResTotal)
	return nil
}

// ReadRows 读取一个批量的数据
func (c *Client) ReadRows() (rows []map[string]interface{}, err error) {
	if int64(c.count) > c.sResTotal {
		return rows, io.EOF
	}
	if c.Limit > 0 && c.count >= c.Limit {
		return rows, io.EOF
	}

	return c.bulkRead()
}

// Read 读取一行数据
func (c *Client) Read() (row map[string]interface{}, err error) {
	if int64(c.count) > c.sResTotal {
		return nil, io.EOF
	}
	if c.Limit > 0 && c.count >= c.Limit {
		return nil, io.EOF
	}

	if len(c.rows) == 0 || c.cursor >= len(c.rows) {
		// 还没有数据，或者游标已经超越了当前数组的下标
		c.rows, err = c.bulkRead()
		if err != nil {
			return nil, err
		}
		c.cursor = 0
	}

	c.cursor++
	return c.rows[c.cursor-1], nil
}

func (c *Client) bulkRead() (rows []map[string]interface{}, err error) {
	if c.page > 0 {
		// 下一页
		c.searchResult, err = func(scrollId string) (*elastic.SearchResult, error) {
			c.page++
			fmt.Printf("search page: %d\n", c.page)
			ctx := context.Background()
			return c.es.Scroll(c.IndexName).ScrollId(scrollId).Do(ctx)
		}(c.searchResult.ScrollId)
		if err != nil {
			return nil, fmt.Errorf("next scroll error: %s", err.Error())
		}
	}

	for i, hit := range c.searchResult.Hits.Hits {
		if i == 0 && c.Debug {
			fmt.Printf("[debug]row[0] = %s\n", string(*hit.Source))
		}
		if c.Limit > 0 && c.count > c.Limit {
			return rows, io.EOF
		}

		var row = make(map[string]interface{})
		err = json.Unmarshal(*hit.Source, &row)
		if err != nil {
			return rows, fmt.Errorf("search %d: json unmarshal error: %v", i, err.Error())
		}

		rows = append(rows, row)
		c.count += 1
	}

	return rows, nil
}
