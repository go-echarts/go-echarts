package types

// Theme Type contains string representations of theme types.
const (
	ThemeChalk         = "chalk"
	ThemeEssos         = "essos"
	ThemeInfographic   = "infographic"
	ThemeMacarons      = "macarons"
	ThemePurplePassion = "purple-passion"
	ThemeRoma          = "roma"
	ThemeRomantic      = "romantic"
	ThemeShine         = "shine"
	ThemeVintage       = "vintage"
	ThemeWalden        = "walden"
	ThemeWesteros      = "westeros"
	ThemeWonderland    = "wonderland"
)

func PresetTheme(theme string) bool {
	switch theme {
	case ThemeChalk, ThemeEssos, ThemeInfographic, ThemeMacarons,
		ThemePurplePassion, ThemeRoma, ThemeRomantic, ThemeShine,
		ThemeVintage, ThemeWalden, ThemeWesteros, ThemeWonderland:
		return true
	default:
		return false
	}
}
