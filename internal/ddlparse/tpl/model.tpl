{{- $conf := .Conf -}}
// {{ .Table.TableName | pascalCase }} {{.Table.Comment}}
type {{ .Table.TableName | pascalCase }} struct {
    {{- range .Table.Columns }}
    {{ .Field | pascalCase }}  {{ getType . $conf }} {{ getTag . $conf }} {{ getComment . $conf }}
    {{- end }}
}

func ({{ .Table.TableName | getFirstLetter}} *{{ .Table.TableName | pascalCase }}) TableName () string {
    return "{{ .Table.TableName }}"
}