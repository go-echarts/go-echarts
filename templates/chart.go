package templates

var ChartTpl = `
{{- define "chart" }}
<!DOCTYPE html>
<html>
    {{- template "header" . }}
<body>
    {{- template "base" . }}
<style>
    .container {margin-top:30px; display: flex;justify-content: center;align-items: center;}
    .item {margin: auto;}
</style>
</body>
</html>
{{ end }}
`
