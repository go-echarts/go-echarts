<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>{{ .Title }}</title>
{{- range .JSAssets.Values }}
    <script src="{{ . }}"></script>
{{- end }}
{{- range .CSSAssets.Values }}
    <link href="{{ . }}" rel="stylesheet">
{{- end }}
</head>
{{ $JSFunctions := .JSFunctions }}
{{- range .Containers }}
<div class="container">
    <div class="item" id="{{ .ChartID }}" style="width:{{ .Width }};height:{{ .Height }};"></div>
</div>

<script type="text/javascript">
    "use strict";
    let goecharts_{{ .ChartID | safeJS }} = echarts.init(document.getElementById('{{ .ChartID }}'), "{{ .Theme }}");
    let option_{{ .ChartID | safeJS }} = {{ .Chart | prettier }};
    goecharts_{{ .ChartID | safeJS }}.setOption(option_{{ .ChartID | safeJS }});
    {{ $ChartID := .ChartID }}
    {{- range .Events }}
    goecharts_{{ $ChartID | safeJS }}.{{ .EventBinder | safeJS }}(
    {{- if .EventType }}
       {{ .EventType }},
    {{- end }}
    {{- if .Params }}
       {{ .Params | safeJS }},
    {{- end }}
    {{- if .Action }}
       {{ .Action | safeJS }},
     {{- end }}
    );
    {{- end }}

{{- range $JSFunctions.Fns }}
    {{ . | safeJS }}
{{- end }}
</script>

<style>
    .container {margin-top:30px; display: flex;justify-content: center;align-items: center;}
    .item {margin: auto;}
</style>

{{- end }}
</html>