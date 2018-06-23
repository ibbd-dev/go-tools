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
	c.ctx = context.Background()
	exists, err := c.es.IndexExists(c.IndexName).Do(c.ctx)
	if err != nil {
		return fmt.Errorf("check index exists error: %v", err.Error())
	}

	if !exists {
		return fmt.Errorf("index %s is not exists", c.IndexName)
	}

	searchResult, err := c.es.Search(c.IndexName).Query(query).Size(1).Do(c.ctx)
	if err != nil {
		return fmt.Errorf("search do error: %v", err.Error())
	}

	c.sResTotal = searchResult.Hits.TotalHits
	if c.Debug {
		fmt.Printf("search research total: %d\n", c.sResTotal)
	}
	return nil
}

// ReadRows 读取一个批量的数据
func (c *Client) ReadRows() (rows []map[string]interface{}, err error) {
	if int64(c.count) > c.sResTotal {
		if c.Debug {
			fmt.Printf("count > res total: %+v\n", c)
		}
		return rows, io.EOF
	}
	if c.Limit > 0 && c.count >= c.Limit {
		return rows, io.EOF
	}

	rows, err = c.bulkRead()
	c.count += len(rows)
	return rows, err
}

// Read 读取一行数据
func (c *Client) Read() (row map[string]interface{}, err error) {
	if int64(c.count) > c.sResTotal {
		if c.Debug {
			c.rows = nil
			fmt.Printf("count > res total: %+v\n", c)
		}
		return nil, io.EOF
	}
	if c.Limit > 0 && c.count >= c.Limit {
		if c.Debug {
			c.rows = nil
			fmt.Printf("limit over: %+v\n", c)
		}
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
	c.count++
	return c.rows[c.cursor-1], nil
}

func (c *Client) bulkRead() (rows []map[string]interface{}, err error) {
	// 分页获取数据
	var scrollId string
	if c.page > 0 {
		scrollId = c.searchResult.ScrollId
	}
	c.searchResult, err = func(scrollId string) (*elastic.SearchResult, error) {
		if c.Debug {
			fmt.Printf("search scroll next page: %d\n", c.page)
		}
		return c.es.Scroll(c.IndexName).ScrollId(scrollId).Query(c.query).Size(c.BulkSize).Do(c.ctx)
	}(scrollId)

	if err == io.EOF { // 没有数据了
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("next scroll error: %s", err.Error())
	}

	c.page++
	if c.Debug {
		fmt.Printf("search result page count: %d\n", len(c.searchResult.Hits.Hits))
	}

	// 解释数据
	for i, hit := range c.searchResult.Hits.Hits {
		if i == 0 && c.Debug {
			fmt.Printf("[debug]row[0] = %s\n", string(*hit.Source))
		}

		var row = make(map[string]interface{})
		err = json.Unmarshal(*hit.Source, &row)
		if err != nil {
			return rows, fmt.Errorf("search %d: json unmarshal error: %v", i, err.Error())
		}

		rows = append(rows, row)
	}

	return rows, nil
}
