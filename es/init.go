package es

import (
	"fmt"
	"time"

	"gopkg.in/olivere/elastic.v5"
)

type Client struct {
	Debug    bool
	Limit    int // 导入导出的数量限制，0表示不限
	BulkSize int // 批量导出导入时没批次的数量，默认为1000

	// es config
	Host      string
	Port      int
	IndexName string
	DocType   string

	// import
	DeleteIndex bool                   // 是否删除原索引
	Mapping     map[string]interface{} // mapping 文件

	// private
	page         int // 当前页数
	count        int // 计数变量
	es           *elastic.Client
	bulk         *elastic.BulkService
	searchResult *elastic.SearchResult
	sResTotal    int64 // 搜索结果的总量
	cursor       int   // 游标
	rows         []map[string]interface{}
}

func NewClient(host string, port int, indexName, docType string) (c *Client, err error) {
	c = &Client{
		Host:     host,
		Port:     port,
		BulkSize: 1000,
	}

	c.es, err = elastic.NewClient(
		elastic.SetURL(fmt.Sprintf("http://%s:%d", host, port)),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetMaxRetries(5),
	)

	return c, err
}
func (c *Client) SetLimit(limit int) {
	c.Limit = limit
}
func (c *Client) SetBulkSize(size int) {
	c.BulkSize = size
}
