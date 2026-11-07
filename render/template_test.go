package render_test

import (
	"bytes"
	"regexp"
	"strings"
	"testing"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/templates"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const attributeOverrides = `
{{ define "script_attributes" }}{{ with renderValue "nonce" }} nonce="{{ . }}"{{ end }}{{ end }}
{{ define "style_attributes" }}{{ with renderValue "nonce" }} nonce="{{ . }}"{{ end }}{{ end }}`

func templateOptions(nonce string) *opts.TemplateOptions {
	return &opts.TemplateOptions{
		Overrides: []string{attributeOverrides},
		Data:      map[string]interface{}{"nonce": nonce},
	}
}

func chartWithTemplates(options *opts.TemplateOptions) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithInitializationOpts(opts.Initialization{
		ChartID:   "template_test",
		Templates: options,
	}))
	return bar
}

func assertTagAttributes(t *testing.T, content, tag, attribute string, count int) {
	t.Helper()
	tags := regexp.MustCompile(`<`+tag+`\b[^>]*>`).FindAllString(content, -1)
	require.Len(t, tags, count)
	for _, opening := range tags {
		assert.Contains(t, opening, attribute)
	}
}

func TestDefaultTemplatesUnchanged(t *testing.T) {
	// Reconstruct the templates without the new attribute hooks and compare
	// byte-for-byte with the default render, including every page layout.
	withoutHooks := func(content string) string {
		content = strings.ReplaceAll(content, "{{- define \"script_attributes\" -}}{{- end -}}\n", "")
		content = strings.ReplaceAll(content, "{{- define \"style_attributes\" -}}{{- end -}}\n", "")
		content = strings.ReplaceAll(content, `{{ template "script_attributes" . }}`, "")
		return strings.ReplaceAll(content, `{{ template "style_attributes" . }}`, "")
	}
	t.Run("chart", func(t *testing.T) {
		bar := chartWithTemplates(nil)
		actual := bar.RenderContent()
		tpl := render.MustTemplate(render.ModChart, []string{
			withoutHooks(templates.HeaderTpl), withoutHooks(templates.BaseTpl),
			withoutHooks(templates.ChartTpl),
		})
		var expected bytes.Buffer
		require.NoError(t, tpl.ExecuteTemplate(&expected, render.ModChart, bar))
		assert.Equal(t, expected.Bytes(), actual)
		bar.Initialization.Templates = &opts.TemplateOptions{}
		assert.Equal(t, actual, bar.RenderContent())
	})
	for _, layout := range []components.Layout{
		components.PageNoneLayout, components.PageCenterLayout,
		components.PageFlexLayout, components.PageFullLayout,
	} {
		t.Run(string(layout), func(t *testing.T) {
			page := components.NewPage().SetLayout(layout).AddCharts(chartWithTemplates(nil))
			actual := page.RenderContent()
			tpl := render.MustTemplate(render.ModPage, []string{
				withoutHooks(templates.HeaderTpl), withoutHooks(templates.BaseTpl),
				withoutHooks(templates.PageTpl),
			})
			var expected bytes.Buffer
			require.NoError(t, tpl.ExecuteTemplate(&expected, render.ModPage, page))
			assert.Equal(t, expected.Bytes(), actual)
			page.SetTemplateOptions(&opts.TemplateOptions{})
			assert.Equal(t, actual, page.RenderContent())
		})
	}
}

func TestHeaderTemplateStandalone(t *testing.T) {
	bar := chartWithTemplates(nil)
	bar.Validate()
	tpl := render.MustTemplate("header", []string{templates.HeaderTpl})
	var buf bytes.Buffer
	require.NoError(t, tpl.ExecuteTemplate(&buf, "header", bar))
	assert.Contains(t, buf.String(), `<script src="`)
	assert.NotContains(t, buf.String(), "nonce=")
}

func TestChartTemplateAttributes(t *testing.T) {
	bar := chartWithTemplates(templateOptions("chart-nonce"))
	bar.AddCustomizedJSAssets("https://example.com/custom.js")
	bar.AddCustomizedHeaders(`<script src="raw.js"></script>`)
	var buf bytes.Buffer
	require.NoError(t, bar.Render(&buf))
	assert.Equal(t, buf.Bytes(), bar.RenderContent())
	content := strings.ReplaceAll(buf.String(), `<script src="raw.js"></script>`, "")
	assertTagAttributes(t, content, "script", `nonce="chart-nonce"`, 3)
	assertTagAttributes(t, content, "style", `nonce="chart-nonce"`, 1)
	assert.Contains(t, buf.String(), `<script src="raw.js"></script>`)

	snippet := bar.RenderSnippet()
	assertTagAttributes(t, snippet.Script, "script", `nonce="chart-nonce"`, 1)
	assert.Contains(t, content, snippet.Script)
	assert.NotContains(t, snippet.Element, "nonce")
	assert.NotContains(t, snippet.Option, "nonce")
	assert.NotContains(t, snippet.Option, "Templates")
}

func TestTemplateOverrides(t *testing.T) {
	tests := []struct {
		name     string
		override string
		want     string
	}{
		{"element", `{{ define "base_element" }}<section id="{{ .ChartID }}">{{ renderValue "label" }}</section>{{ end }}`, `<section id="template_test">custom-label</section>`},
		{"script", `{{ define "base_script" }}<script>console.log({{ renderValue "label" }});</script>{{ end }}`, `<script>console.log("custom-label");</script>`},
		{"option", `{{ define "base_option" }}{"custom": true}{{ end }}`, `{"custom": true}`},
		{"header", `{{ define "header" }}<head><title>{{ renderValue "label" }}</title></head>{{ end }}`, `<head><title>custom-label</title></head>`},
		{"chart", `{{ define "chart" }}<main>{{ renderValue "label" }}</main>{{ end }}`, `<main>custom-label</main>`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bar := chartWithTemplates(&opts.TemplateOptions{
				Overrides: []string{tt.override},
				Data:      map[string]interface{}{"label": "custom-label"},
			})
			assert.Contains(t, string(bar.RenderContent()), tt.want)
			snippet := bar.RenderSnippet()
			switch tt.name {
			case "element":
				assert.Equal(t, tt.want, strings.TrimSpace(snippet.Element))
			case "script":
				assert.Equal(t, tt.want, strings.TrimSpace(snippet.Script))
			case "option":
				assert.Contains(t, snippet.Script, tt.want)
				assert.Equal(t, tt.want, strings.TrimSpace(snippet.Option))
			}
		})
	}
}

func TestPageTemplateOptions(t *testing.T) {
	for _, layout := range []components.Layout{
		components.PageNoneLayout, components.PageCenterLayout,
		components.PageFlexLayout, components.PageFullLayout,
	} {
		t.Run(string(layout), func(t *testing.T) {
			childOptions := templateOptions("child-nonce")
			childOptions.Overrides = append(childOptions.Overrides,
				`{{ define "base_element" }}<aside>child-only</aside>{{ end }}`)
			bar := chartWithTemplates(childOptions)
			line := charts.NewLine()
			page := components.NewPage().SetLayout(layout).
				SetTemplateOptions(templateOptions("page-nonce")).AddCharts(bar, line)
			page.AddCustomizedJSAssets("https://example.com/page.js")
			var buf bytes.Buffer
			require.NoError(t, page.Render(&buf))
			content := buf.String()
			assertTagAttributes(t, content, "script", `nonce="page-nonce"`, 4)
			styles := 1
			if layout == components.PageNoneLayout {
				styles = 0
			}
			assertTagAttributes(t, content, "style", `nonce="page-nonce"`, styles)
			assert.NotContains(t, content, "child-nonce")
			assert.NotContains(t, content, "child-only")
			assert.Same(t, childOptions, bar.GetTemplateOptions())
			assert.Equal(t, "child-nonce", childOptions.Data["nonce"])
			assert.Contains(t, string(bar.RenderContent()), "child-only")

			// An unconfigured page must not inherit child template options.
			page.SetTemplateOptions(nil)
			content = string(page.RenderContent())
			assert.NotContains(t, content, "nonce=")
			assert.NotContains(t, content, "child-only")
		})
	}
}

func TestPageTemplateOverride(t *testing.T) {
	page := components.NewPage().SetTemplateOptions(&opts.TemplateOptions{
		Overrides: []string{`{{ define "page" }}<main>{{ renderValue "label" }}</main>{{ end }}`},
		Data:      map[string]interface{}{"label": "custom-page"},
	})
	assert.Equal(t, `<main>custom-page</main>`, string(page.RenderContent()))
}

func TestTemplateDataUpdatedBetweenRenders(t *testing.T) {
	options := templateOptions("first")
	bar := chartWithTemplates(options)
	assert.Contains(t, string(bar.RenderContent()), `nonce="first"`)
	options.Data["nonce"] = "second"
	content := string(bar.RenderContent())
	assert.Contains(t, content, `nonce="second"`)
	assert.NotContains(t, content, `nonce="first"`)
	assert.Contains(t, bar.RenderSnippet().Script, `nonce="second"`)
	delete(options.Data, "nonce")
	assert.NotContains(t, string(bar.RenderContent()), "nonce=")
	bar.Initialization.Templates = templateOptions("replacement")
	assert.Contains(t, string(bar.RenderContent()), `nonce="replacement"`)
	bar.Initialization.Templates = nil
	assert.NotContains(t, bar.RenderSnippet().Script, "nonce=")
}

func TestTemplateAttributeEscaping(t *testing.T) {
	bar := chartWithTemplates(templateOptions(`<&"'`))
	content := string(bar.RenderContent())
	assertTagAttributes(t, content, "script", `nonce="&lt;&amp;&#34;&#39;"`, 2)
	assertTagAttributes(t, content, "style", `nonce="&lt;&amp;&#34;&#39;"`, 1)
	assert.NotContains(t, content, `<&"'`)
}

func TestCustomScriptAndStyleAttributes(t *testing.T) {
	bar := chartWithTemplates(&opts.TemplateOptions{
		Overrides: []string{`
			{{ define "script_attributes" }} data-tenant="{{ renderValue "tenant" }}"{{ end }}
			{{ define "style_attributes" }} data-layout="{{ renderValue "layout" }}"{{ end }}
		`},
		Data: map[string]interface{}{"tenant": "tenant-a", "layout": "compact"},
	})
	content := string(bar.RenderContent())
	assertTagAttributes(t, content, "script", `data-tenant="tenant-a"`, 2)
	assertTagAttributes(t, content, "style", `data-layout="compact"`, 1)
}

func TestTemplateOverrideOrder(t *testing.T) {
	bar := chartWithTemplates(&opts.TemplateOptions{Overrides: []string{
		`{{ define "base_element" }}<div>first</div>{{ end }}`,
		`{{ define "base_element" }}<div>second</div>{{ end }}`,
	}})
	assert.Equal(t, `<div>second</div>`, strings.TrimSpace(bar.RenderSnippet().Element))
}

func TestTemplateOptionsIsolation(t *testing.T) {
	for _, nonce := range []string{"one", "two", "three"} {
		nonce := nonce
		t.Run(nonce, func(t *testing.T) {
			t.Parallel()
			bar := chartWithTemplates(templateOptions(nonce))
			assertTagAttributes(t, string(bar.RenderContent()), "script", `nonce="`+nonce+`"`, 2)
		})
	}
}

func TestInvalidTemplateOverride(t *testing.T) {
	bar := chartWithTemplates(&opts.TemplateOptions{Overrides: []string{"{{"}})
	assert.Panics(t, func() { bar.RenderContent() })
}
