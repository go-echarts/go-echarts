package opts

// TemplateOptions customizes the HTML templates used by a chart or page.
// Templates are trusted code; do not populate Overrides from untrusted input.
type TemplateOptions struct {
	// Overrides are parsed after the built-in templates, in order. Define a
	// template with an existing name to replace that part of the output.
	Overrides []string

	// Data is available to all templates through the renderValue function.
	// Refresh request-specific values before rendering, and do not modify this
	// map or the options concurrently with rendering.
	Data map[string]interface{}
}

// GetTemplateOptions returns the template configuration for a single chart.
func (opt Initialization) GetTemplateOptions() *TemplateOptions {
	return opt.Templates
}

// GetTemplateOptions returns the template configuration for a page.
func (opt PageConfiguration) GetTemplateOptions() *TemplateOptions {
	return opt.Templates
}
