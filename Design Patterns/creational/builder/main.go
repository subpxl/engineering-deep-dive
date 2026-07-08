package main

import (
	"fmt"
	"strings"
)

// sql query builder
func main() {

	query1 := NewSqlQueryBuilder("users").
		Select("name", "email").
		Where("age > 10").
		Where("active = true").
		OrderBy("name", "ASC").
		Limit(10).
		Build()

	query2 := NewSqlQueryBuilder("orders").
		Select("id", "total", "created_at").
		Where("status = 'completed'").
		Where("total > 100").
		OrderBy("created_at", "DESC").
		Limit(20).
		Offset(40).
		Build()

	fmt.Println(query1)
	fmt.Println(query2)
}

type SqlQuery struct {
	table          string
	columns        []string
	conditions     []string
	orderBy        string
	orderDirection string
	limit          int
	offset         int
}

func (q *SqlQuery) String() string {
	return q.toSql()
}

func (q *SqlQuery) toSql() string {
	var sql strings.Builder
	sql.WriteString("SELECT ")
	if len(q.columns) == 0 {
		sql.WriteString("*")
	} else {
		sql.WriteString(strings.Join(q.columns, ", "))
	}
	sql.WriteString(" FROM ")
	sql.WriteString(q.table)
	if len(q.conditions) > 0 {
		sql.WriteString(" WHERE ")
		sql.WriteString(strings.Join(q.conditions, " AND "))
	}
	if q.orderBy != "" {
		sql.WriteString(" ORDER BY ")
		sql.WriteString(q.orderBy)
		sql.WriteString(" ")
		sql.WriteString(q.orderDirection)
	}
	if q.limit > 0 {
		sql.WriteString(" LIMIT ")
		sql.WriteString(fmt.Sprint(q.limit))
	}
	if q.offset > 0 {
		sql.WriteString(" OFFSET ")
		sql.WriteString(fmt.Sprint(q.offset))
	}
	return sql.String()
}

type SqlQueryBuilder struct {
	table          string
	columns        []string
	conditions     []string
	orderBy        string
	orderDirection string
	limit          int
	offset         int
}

func NewSqlQueryBuilder(table string) *SqlQueryBuilder {
	return &SqlQueryBuilder{
		table:          table,
		columns:        []string{},
		conditions:     []string{},
		orderDirection: "ASC",
	}
}

func (b *SqlQueryBuilder) Select(cols ...string) *SqlQueryBuilder {
	b.columns = append(b.columns, cols...)
	return b
}

func (b *SqlQueryBuilder) Where(condition string) *SqlQueryBuilder {
	b.conditions = append(b.conditions, condition)
	return b
}

func (b *SqlQueryBuilder) OrderBy(column, direction string) *SqlQueryBuilder {
	b.orderBy = column
	b.orderDirection = direction
	return b
}

func (b *SqlQueryBuilder) Limit(limit int) *SqlQueryBuilder {
	b.limit = limit
	return b
}

func (b *SqlQueryBuilder) Offset(offset int) *SqlQueryBuilder {
	b.offset = offset
	return b
}

func (b *SqlQueryBuilder) Build() *SqlQuery {
	return &SqlQuery{
		table:          b.table,
		columns:        append([]string(nil), b.columns...),
		conditions:     append([]string(nil), b.conditions...),
		orderBy:        b.orderBy,
		orderDirection: b.orderDirection,
		limit:          b.limit,
		offset:         b.offset,
	}
}
