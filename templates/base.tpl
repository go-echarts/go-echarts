{{- define "base" }}
<div class="container">
    <div class="item" id="{{ .ChartID }}" style="width:{{ .Initialization.Width }};height:{{ .Initialization.Height }};"></div>
</div>

<script type="text/javascript">
    "use strict";
    let goecharts_{{ .ChartID | safeJS }} = echarts.init(document.getElementById('{{ .ChartID | safeJS }}'), "{{ .Theme }}", { renderer: "{{  .Initialization.Renderer }}" });
    let option_{{ .ChartID | safeJS }} = {{ .JSONNotEscaped | safeJS }};
    {{ if isSetAction .BaseActions.Type }}
	let action_{{ .ChartID | safeJS }} = {{ .JSONNotEscapedAction | safeJS }};
 	goecharts_{{ .ChartID | safeJS }}.dispatchAction(action_{{ .ChartID | safeJS }});
    {{ end }}
    goecharts_{{ .ChartID | safeJS }}.setOption(option_{{ .ChartID | safeJS }});

  {{- range  $listener := .EventListeners }}
    {{if .Query  }}
    goecharts_{{ $.ChartID | safeJS }}.on({{ $listener.EventName }}, {{ $listener.Query | safeJS }}, {{ injectInstance $listener.Handler "%MY_ECHARTS%"  $.ChartID | safeJS }});
    {{ else }}
    goecharts_{{ $.ChartID | safeJS }}.on({{ $listener.EventName }}, {{ injectInstance $listener.Handler "%MY_ECHARTS%"  $.ChartID | safeJS }})
    {{ end }}
  {{- end }}

    {{- range .JSFunctions.Fns }}
    {{ injectInstance . "%MY_ECHARTS%"  $.ChartID  | safeJS }}
    {{- end }}
</script>
{{ end }}