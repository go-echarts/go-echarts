{{- define "script_attributes" -}}{{- end -}}
{{ define "header" }}
<head>
    <meta charset="utf-8">
    <title>{{ .PageTitle }}</title>
{{- range .JSAssets.Values }}
    <script{{ template "script_attributes" . }} src="{{ . }}"></script>
{{- end }}
{{- range .CustomizedJSAssets.Values }}
    <script{{ template "script_attributes" . }} src="{{ . }}"></script>
{{- end }}
{{- range .CSSAssets.Values }}
    <link href="{{ . }}" rel="stylesheet">
{{- end }}
{{- range .CustomizedCSSAssets.Values }}
    <link href="{{ . }}" rel="stylesheet">
{{- end }}
{{- range .CustomizedHeaders.Values }}
    {{ . | safeHTML }}
{{- end }}
</head>
{{ end }}
