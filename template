// {{ pascalCase .Table.TableName }} {{.Table.Comment}}
type {{ pascalCase .Table.TableName }} struct {
    {{- range .Columns }}
    {{ . }}
    {{- end }}
}

func ({{ getWord .Table.TableName }} *{{pascalCase .Table.TableName}}) TableName () string {
    return "{{ .Table.TableName }}"
}