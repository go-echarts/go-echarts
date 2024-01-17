{{- define "base" }}
<div class="container">
    <div class="item" id="{{ .ChartID }}" style="width:{{ .Initialization.Width }};height:{{ .Initialization.Height }};"></div>
</div>

<script type="text/javascript">
    "use strict";
    let goecharts_{{ .ChartID | safeJS }} = echarts.init(document.getElementById('{{ .ChartID | safeJS }}'), "{{ .Theme }}", { renderer: "{{  .Initialization.Renderer }}" });
    let option_{{ .ChartID | safeJS }} = {{ .JSONNotEscaped | safeJS }};
    {{ if isSet  "BaseActions" . }}
	let action_{{ .ChartID | safeJS }} = {{ .JSONNotEscapedAction | safeJS }};
    {{ end }}
    goecharts_{{ .ChartID | safeJS }}.setOption(option_{{ .ChartID | safeJS }});
 	goecharts_{{ .ChartID | safeJS }}.dispatchAction(action_{{ .ChartID | safeJS }});

  {{- range  $listener := .EventListeners }}
    {{if .Query  }}
    goecharts_{{ $.ChartID | safeJS }}.on({{ $listener.EventName }}, {{ $listener.Query | safeJS }},{{ $listener.Handler | safeJS }});
    {{ else }}
    goecharts_{{ $.ChartID | safeJS }}.on({{ $listener.EventName }},{{ $listener.Handler | safeJS }})
    {{ end }}
  {{- end }}

    {{- range .JSFunctions.Fns }}
    {{ . | safeJS }}
    {{- end }}
</script>
{{ end }}