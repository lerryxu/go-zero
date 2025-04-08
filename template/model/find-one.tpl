func (m *default{{.upperStartCamelObject}}Model) FindOne(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) (*{{.upperStartCamelObject}}, error) {
	{{if .withCache}}{{.cacheKey}}
	var resp {{.upperStartCamelObject}}
	err := m.QueryRowCtx(ctx, &resp, {{.cacheKeyVariable}}, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query :=  fmt.Sprintf("select %s from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} limit 1", {{.lowerStartCamelObject}}Rows, m.table)
		return conn.QueryRowCtx(ctx, v, query, {{.lowerStartCamelPrimaryKey}})
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}{{else}}query := fmt.Sprintf("select %s from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} limit 1", {{.lowerStartCamelObject}}Rows, m.table)
	var resp {{.upperStartCamelObject}}
	err := m.conn.QueryRowCtx(ctx, &resp, query, {{.lowerStartCamelPrimaryKey}})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}{{end}}
}

func (m *default{{.upperStartCamelObject}}Model) FindOneBy(ctx context.Context, {{.lowerStartCamelPrimaryKey}} map[string]any) (*{{.upperStartCamelObject}}, error) {
	keys := []string{}
	values := []any{}
	for k, v := range {{.lowerStartCamelPrimaryKey}} {
		keys = append(keys, fmt.Sprintf("`%s` = ?", k))
		values = append(values, v)
	}
	query :=  fmt.Sprintf("select %s from %s where %s limit 1", {{.lowerStartCamelObject}}Rows, m.table, strings.Join(keys, " and "))
	var resp {{.upperStartCamelObject}}
	{{if .withCache}}
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, values...){{else}}
	err := m.conn.QueryRowCtx(ctx, &resp, query, values...){{end}}
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
