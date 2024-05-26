package bus

//Common control events like window size and title. These are not require. The user can
//create their own control events by adding the ControlEvent interface to their own.

// /////WindowTitleEvent//////
func NewWindowTitleEvent(data string) ControlEvent {
	return &windowTitleData{data}
}

type WindowTitleEvent interface {
	ControlEvent
	GetTitle() string
}

type windowTitleData struct {
	title string
}

func (w *windowTitleData) GetTitle() string {
	return w.title
}

// //WindowSizeEvent
func NewWindowSizeEvent(width, height int) WindowSizeEvent {
	return &windowSizeData{width, height}
}

type WindowSizeEvent interface {
	ControlEvent
	GetWindowSize() (int, int)
}

type windowSizeData struct {
	width, height int
}

func (w *windowSizeData) GetWindowSize() (int, int) {
	return w.width, w.height
}
