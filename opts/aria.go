package opts

import "github.com/go-echarts/go-echarts/v2/types"

// Aria options refer https: //echarts.apache.org/en/option.html#aria
type Aria struct {
	Enabled types.Bool `json:"enabled,omitempty"`
	// need to verify the functionalities
	Label *AriaLabel `json:"label,omitempty"`
	Decal *AriaDecal `json:"decal,omitempty"`
}

type AriaLabel struct {
	// Enabled Whether to enable label generation for accessibility.
	Enabled types.Bool `json:"enabled,omitempty"`

	// Description By default, an algorithm is used to automatically generate a chart description.
	// If you want to fully customize it, you can set this value to a description.
	Description string `json:"description,omitempty"`

	// General For the overall description of the chart.
	General *AriaGeneral `json:"general,omitempty"`

	// Series series-related configuration items.
	Series *AriaSeries `json:"series,omitempty"`
}
type AriaData struct {
	// MaxCount The maximum number of data occurrences per series in the description.
	MaxCount types.Int `json:"maxCount,omitempty"`

	// AllData Description to be used when all data is displayed.
	AllData string `json:"allData,omitempty"`

	// PartialData Descriptions used when only partial data is displayed. This includes template variables.
	// {displayCnt}: the number of data that will be replaced with the number of displays.
	PartialData string `json:"partialData,omitempty"`

	// WithName This description is used if the data has the name attribute. This includes the template variable.
	// {name}: name that will be replaced with the data.
	// {value}: the value that will be replaced with the data.
	WithName string `json:"withName,omitempty"`

	// WithoutName This description is used if the data does not have the name attribute. This includes the template variable.
	// {value}: the value that will be replaced with the data.
	WithoutName string `json:"withoutName,omitempty"`

	// Separator The separator between data and data description.
	Separator *AriaSeparator `json:"separator,omitempty"`
}
type AriaDecal struct {
	// Show Whether to display the decal pattern is not shown by default.
	Show types.Bool `json:"show,omitempty"`
	// Decals The style of the decal pattern.
	// An array, each item in the array will have one style and the data will be looped through the array in order.
	Decals []AriaDecal `json:"decals,omitempty"`
}

type AriaDecals struct {
	// Symbol of node of this category.
	// Icon types provided by ECharts includes
	// 'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
	// It can be set to an image with 'image://url' , in which URL is the link to an image, or dataURI of an image.
	Symbol string `json:"symbol,omitempty"`

	// SymbolSize Range of values: 0 to 1, representing the size of symbol relative to decal.
	SymbolSize types.Float `json:"symbolSize,omitempty"`
}

type AriaGeneral struct {
	// WithTitle If the chart exists title.text, then withTitle is used. This includes the template variable.
	// {title}: will be replaced with title.text of the chart.
	WithTitle string `json:"withTitle,omitempty"`

	// WithoutTitle If the chart does not have title.text, then withoutTitle is used.
	WithoutTitle string `json:"withoutTitle,omitempty"`
}

type AriaSeries struct {
	// MaxCount The maximum number of series in the description.
	MaxCount types.Int `json:"maxCount,omitempty"`
	// SingleSeries The description used when the chart contains only one series.
	SingleSeries *AriaSingle `json:"single,omitempty"`
	// MultipleSeries The description used when the chart contains multiple series.
	MultipleSeries *AriaMultiple `json:"multiple,omitempty"`
}

type AriaSingle struct {
	// Prefix Holistic descriptions for all series are shown before each series description. This includes template variables.
	// {seriesCount}: will be replaced with the number of series, where it is always 1.
	Prefix string `json:"prefix,omitempty"`

	// WithName This description is used if the series has the name attribute. This includes the template variable.
	// {seriesName}: will be replaced with name of the series.
	// {seriesType}: the name of the type that will be replaced with the series, e.g. 'Bar chart', 'Line chart', etc.
	WithName string `json:"withName,omitempty"`

	// WithoutName This description is used if the series has no name attribute. This includes the template variable.
	// {seriesType}: the name of the type that will be replaced with the series, e.g. 'Bar chart', 'Line chart', etc.v
	WithoutName string `json:"withoutName,omitempty"`
}

type AriaMultiple struct {
	// Prefix Holistic descriptions for all series are shown before each series description. This includes template variables.
	// {seriesCount}: will be replaced with the number of series, where it is always 1.
	Prefix string `json:"prefix,omitempty"`

	// WithName This description is used if the series has the name attribute. This includes the template variable.
	// {seriesName}: will be replaced with name of the series.
	// {seriesType}: the name of the type that will be replaced with the series, e.g. 'Bar chart', 'Line chart', etc.
	WithName string `json:"withName,omitempty"`

	// WithoutName This description is used if the series has no name attribute. This includes the template variable.
	// {seriesType}: the name of the type that will be replaced with the series, e.g. 'Bar chart', 'Line chart', etc.v
	WithoutName string `json:"withoutName,omitempty"`
}

type AriaSeparator struct {
	// Middle The delimiter of the data except the last one.
	Middle string `json:"middle,omitempty"`

	// End The delimiter after the last data.
	End string `json:"end,omitempty"`
}
