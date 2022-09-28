package templates

var BaseTpl = `
{{- define "base" }}
<div class="container">
    <div class="item" id="{{ .ChartID }}" style="width:{{ .Initialization.Width }};height:{{ .Initialization.Height }};"></div>
</div>

<script type="text/javascript">
    "use strict";
    let goecharts_{{ .ChartID | safeJS }} = echarts.init(document.getElementById('{{ .ChartID | safeJS }}'), "{{ .Theme }}");
    let option_{{ .ChartID | safeJS }} = {{ .JSONNotEscaped | safeJS }};
	let action_{{ .ChartID | safeJS }} = {{ .JSONNotEscapedAction | safeJS }};
    goecharts_{{ .ChartID | safeJS }}.setOption(option_{{ .ChartID | safeJS }});
 	goecharts_{{ .ChartID | safeJS }}.dispatchAction(action_{{ .ChartID | safeJS }});

    {{if .UpdaterConfig}}
    function UpdateOption() {
        var ip_addr = document.location.host;
    var ws = new WebSocket("ws://"+ip_addr+"/ws/{{ .ChartID| safeJS }}");
    ws.onopen = function () {
        {
           
            console.log("sending...");
        };

    }
    ws.onmessage = function (evt) {
        var received_msg = evt.data;
        console.log("data received...",received_msg);
        var opt=JSON.parse(evt.data)
        goecharts_{{ .ChartID | safeJS }}.setOption(opt)
    };
    ws.onclose = function () {
        // close websocket
        console.log("connection closed...");
    };
}
UpdateOption();
    {{end}}
    {{- range .JSFunctions.Fns }}
    {{ . | safeJS }}
    {{- end }}
</script>
{{ end }}
`
