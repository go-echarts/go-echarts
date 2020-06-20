package templates

var BaseTpl = `
{{- define "base" }}
<div class="container">
    <div class="item" id="{{ .ChartID }}"
         style="width:{{ .Initialization.Width }};height:{{ .Initialization.Height }};"></div>
</div>
  <script type="text/javascript">
     "use strict";
     var myChart___x__{{ .ChartID }}__x__ = echarts.init(document.getElementById('{{ .ChartID }}'), "{{ .Theme }}");
     var option___x__{{ .ChartID }}__x__ = {
         title: {{ .Title  }},
         tooltip: {{ .Tooltip }},
         legend: {{ .Legend }},
     {{- if .HasGeo }}
         geo: {{ .GeoComponent }},
     {{- end }}
 {{- if .HasRadar }}
         radar: {{ .RadarComponent }},
     {{- end }}
     {{- if .HasParallel }}
         parallel: {{ .ParallelComponent }},
         parallelAxis: {{ .ParallelAxis }},
     {{- end }}
     {{- if .HasSingleAxis }}
         singleAxis: {{ .SingleAxis }},
     {{- end }}
     {{- if .Toolbox.Show }}
         toolbox: {{ .Toolbox }},
     {{- end }}
     {{ $datazoomLen := len .DataZoomList }}
     {{- if gt $datazoomLen 0 }}
         dataZoom:{{ .DataZoomList }},
     {{- end }}
     {{ $visualmapLen := len .VisualMapList }}
     {{- if gt $visualmapLen 0 }}
         visualMap:{{ .VisualMapList }},
     {{- end }}
     {{- if .HasXYAxis }}
         xAxis: {{ .XAxisList }},
         yAxis: {{ .YAxisList }},
     {{- end }}
     {{- if .Has3DAxis }}
         xAxis3D: {{ .XAxis3D }},
         yAxis3D: {{ .YAxis3D }},
         zAxis3D: {{ .ZAxis3D }},
         grid3D: {{ .Grid3D }},
     {{- end }}
         series: [
         {{ range .MultiSeries }}
         {{- . }},
         {{ end -}}
         ],
     {{- if eq .Theme "white" }}
         color: {{ .Colors }},
     {{- end }}
     {{- if ne .BackgroundColor "" }}
         backgroundColor: {{ .BackgroundColor }}
     {{- end }}
     };
     myChart___x__{{ .ChartID }}__x__.setOption(option___x__{{ .ChartID }}__x__);

     {{- range .JSFunctions.Fns }}
     {{ . }}
   {{- end }}
 </script>
{{ end }}
`
