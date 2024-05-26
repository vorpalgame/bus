package bus

// Common functionality as a usable example but the user can creat their own
// draw events by simply adding the interface DrawEvent so that the bus can distinguish.
// The receiver can switch on the type.
type DrawLayersEvent interface {
	DrawEvent
	Reset()
	GetImageLayers() []ImageLayerData
	AddImageLayer(imgLayer ImageLayerData) DrawEvent
}

type drawEvent struct{}

///// Constructors //////////////////

func NewDrawLayersEvent() DrawLayersEvent {
	evt := drawLayerEvent{}
	evt.imageLayers = make([]ImageLayerData, 0)
	return &evt
}
func NewDrawEvent() DrawEvent {
	return &drawEvent{}
}

/////////////////////////////////////

type ImageLayerData struct {
	LayerMetadata []*ImageMetadata `yaml:"LayerMetadata"`
}

/////////////////////////////////////////////////////////////////////////

type ImageMetadata struct {
	ImageFileName  string `yaml:"ImageFileName"`
	X              int32  `yaml:"X"`
	Y              int32  `yaml:"Y"`
	Width          int32  `yaml:"Width"`
	Height         int32  `yaml:"Height"`
	HorizontalFlip bool   `yaml:"HorizontalFlip"`
}

/////////////////////////////////////////////////

type drawLayerEvent struct {
	imageLayers []ImageLayerData
}

func (evt *drawLayerEvent) Reset() {
	evt.imageLayers = make([]ImageLayerData, 0)
}

func (evt *drawLayerEvent) AddImageLayer(img ImageLayerData) DrawEvent {
	evt.imageLayers = append(evt.imageLayers, img)
	return evt
}

func (evt *drawLayerEvent) GetImageLayers() []ImageLayerData {
	return evt.imageLayers
}
