{{- define "page" }}
<!DOCTYPE html>
<html>
    {{- template "header" . }}
<body>
{{ if eq .Layout "none" }}
    {{- range .Charts }} {{ template "base" . }} {{- end }}
{{ end }}

{{ if eq .Layout "center" }}
    <style{{ template "style_attributes" . }}> .container {display: flex;justify-content: center;align-items: center;} .item {margin: auto;} </style>
    {{- range .Charts }} {{ template "base" . }} {{- end }}
{{ end }}

{{ if eq .Layout "flex" }}
    <style{{ template "style_attributes" . }}> .box { justify-content:center; display:flex; flex-wrap:wrap } </style>
    <div class="box"> {{- range .Charts }} {{ template "base" . }} {{- end }} </div>
{{ end }}

{{ if eq .Layout "full" }}
    <style{{ template "style_attributes" . }}> .container {height: 100vh;} .item {margin: auto;} </style>
    {{- range .Charts }} {{ template "base" . }} {{- end }}
{{ end }}

</body>
</html>
{{ end }}
