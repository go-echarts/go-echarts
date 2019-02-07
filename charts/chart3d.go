package charts

type Chart3D struct {
	BaseOpts
	Series
	Grid3dOpacity float32
	Grid3dShading string
	XAxis3dType   string
	YAxis3dType   string
	ZAxis3dType   string
}
