# 数据库查询相关工具

## 分页读取数据工具

### install

```
go get -u github.com/ibbd-dev/go-tools/db/query
```

### 使用

```
// 统计总数
totalString := fmt.Sprintf("SELECT COUNT(*) as cnt FROM `%s`", tableName)
queryString := fmt.Sprintf("SELECT `id`, `content` FROM `%s` ORDER BY `id`", tableName)
paging, err := query.NewPaging(db, totalString, queryString, perPage)
if err != nil {
    log.Println("分页读取数据初始化出错")
    return err
}
log.Printf("数据库记录总数为：%d, 分%d页处理。\n", paging.GetTotal(), paging.GetPagesTotal())

var count int
for paging.HasNextPage() {
    log.Printf("\n读取第%d页的数据...\n", paging.GetCurrentPage()+1)

    rows, err := paging.Next()
    log.Println("SQL: ", paging.GetPageSql())
    if err != nil {
        log.Println("paging query error")
        rows.Close()
        return err
    }

    // 读取数据
    contents = contents[0:0]
    ids = ids[0:0]
    for rows.Next() {
        var row = &Query{}
        err = rows.Scan(&row.Id, &row.Content)
        if err != nil {
            log.Printf("读取数据出错，当前已读取：%d\n", count+1)
            rows.Close()
            return err
        }

        count += 1
    }
    rows.Close()

}
```

