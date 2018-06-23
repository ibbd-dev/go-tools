# Elasticsearch

实现功能：

- [x] 批量导入
- [x] 搜索

## init
new客户端对象，及一些基础参数

```go
func NewClient(host string, port int, indexName, docType string) (c *Client, err error) 

func (c *Client) SetDebug(debug bool) 
func (c *Client) SetLimit(limit int) 
func (c *Client) SetBulkSize(size int) 
```

## import
批量导入数据

```go
// ImportInit 数据导入初始化
func (c *Client) ImportInit(deleteIndex bool, mapping map[string]interface{}) error 

// BulkAdd 增加一行记录
func (c *Client) BulkAdd(row interface{}) error 

// BulkImport 批量导入
// 注意：外部最后需要调用一次这个方法
func (c *Client) BulkImport() error 
```

## search

```go
// SearchInit 搜索初始化
func (c *Client) SearchInit(query elastic.Query) error 

// ReadRows 读取一个批量的数据
func (c *Client) ReadRows() (rows []map[string]interface{}, err error) 

// Read 读取一行数据
func (c *Client) Read() (row map[string]interface{}, err error) 
```

