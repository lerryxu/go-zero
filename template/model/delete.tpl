func (m *default{{.upperStartCamelObject}}Model) Delete(ctx context.Context, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) error {
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, {{.lowerStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}

{{end}}	{{.keys}}
    _, err {{if .containsIndexCache}}={{else}}:={{end}} m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table)
		return conn.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("delete from %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table)
		_,err:=m.conn.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}}){{end}}
	return err
}

func (m *default{{.upperStartCamelObject}}Model) DeleteBy(ctx context.Context, {{.lowerStartCamelPrimaryKey}} map[string]any) error {
	keys := []string{}
	values := []any{}
	for k, v := range {{.lowerStartCamelPrimaryKey}} {
		keys = append(keys, fmt.Sprintf("`%s` = ?", k))
		values = append(values, v)
	}
	query := fmt.Sprintf("delete from %s where %s", m.table, strings.Join(keys, " and "))
	{{if .withCache}}_, err:= m.ExecNoCacheCtx(ctx, query, values...){{else}}}
	_, err:= m.conn.ExecCtx(ctx, query, values...){{end}}
	return err
}
