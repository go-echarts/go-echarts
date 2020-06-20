package templates

var RoutersTpl = `
{{- define "routers" }}
{{ $routersLen := len .Routers }}
{{- if gt $routersLen 0 }}
<div class="select" style="margin-right:10px; margin-top:10px; position:fixed; right:10px;">
    <select onchange="window.location.href=this.options[this.selectedIndex].value">
    {{- range .Routers }}
        <option value="{{ .URL }}">{{ .Text }}</option>
    {{- end }}
    </select>
{{- end -}}
</div>
{{ end }}
`
