package query

import (
	"database/sql"
	"fmt"
	"math"
)

type Paging struct {
	db           *sql.DB
	total        int    // 记录总数
	perPage      int    // 每页显示的记录数
	totalSql     string // 查询总数的sql语句
	querySql     string // 查询数据的sql语句
	pageQuerySql string // 查询数据的sql语句（分页查询sql）
	currentPage  int    // 当前页码
	pagesTotal   int    // 总页数
}

// NewPaging 获取分页读取数据表对象
// @param totalSql 获取记录总数的sql, 如：select count(*) as cnt from tablename
// @param querySql 获取查询数据的sql，不需要分页的limit部分, 如：select id from tablename
func NewPaging(db *sql.DB, totalSql, querySql string, perPage int) (p *Paging, err error) {
	p = &Paging{
		perPage:  perPage,
		db:       db,
		totalSql: totalSql,
		querySql: querySql,
	}

	// 获取记录总数
	rows, err := db.Query(totalSql)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		if err = rows.Scan(&p.total); err != nil {
			rows.Close()
			return nil, err
		}
	}
	rows.Close()

	// 计算总页数
	p.pagesTotal = int(math.Ceil(float64(p.total) / float64(perPage)))
	return p, nil
}

func (p *Paging) Reset() {
	p.currentPage = 0
}

// HasNextPage 判断是否存在下一页
func (p *Paging) HasNextPage() bool {
	return p.currentPage < p.pagesTotal
}

// Next 获取下一页
func (p *Paging) Next() (*sql.Rows, error) {
	offset := p.currentPage * p.perPage
	p.currentPage++
	p.pageQuerySql = fmt.Sprintf("%s LIMIT %d, %d", p.querySql, offset, p.perPage)
	return p.db.Query(p.pageQuerySql)
}

// GetTotal 获取记录总数
func (p *Paging) GetTotal() int {
	return p.total
}

// GetPagesTotal 获取分页总数
func (p *Paging) GetPagesTotal() int {
	return p.pagesTotal
}

// GetCurrentPage 获取当前页码
func (p *Paging) GetCurrentPage() int {
	return p.currentPage
}

// GetPageSql 获取分页查询的sql
func (p *Paging) GetPageSql() string {
	return p.pageQuerySql
}
