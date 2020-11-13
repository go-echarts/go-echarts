package templates

var HeaderTpl = `
{{ define "header" }}
<head>
    <meta charset="utf-8">
    <title>{{ .PageTitle }}</title>
{{- range .JSAssets.Values }}
    <script src="{{ . }}"></script>
{{- end }}
{{- range .CustomizedJSAssets.Values }}
    <script src="{{ . }}"></script>
{{- end }}
{{- range .CSSAssets.Values }}
    <link href="{{ . }}" rel="stylesheet">
{{- end }}
{{- range .CustomizedCSSAssets.Values }}
    <link href="{{ . }}" rel="stylesheet">
{{- end }}
</head>
{{ end }}
`
