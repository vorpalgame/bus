package bus

type TextEventListener interface {
	OnTextEvent(textChannel <-chan TextEvent)
}

// ////////////////////////////////////////////////
// // textEvent
// ////////////////////////////////////////////////

type TextEvent interface {
	Reinitialize() TextEvent
	GetLocation() (x, y int32)
	SetLocation(x, y int32) TextEvent
}
type textEventData struct {
	X, Y int32
}

func (ted *textEventData) GetLocation() (x, y int32) {
	return ted.X, ted.Y
}
func (ted *textEventData) SetLocation(x, y int32) TextEvent {
	ted.X = x
	ted.Y = y
	return ted
}

func (ted *textEventData) Reinitialize() TextEvent {
	//More specific events can override in whatever way
	//make sense. At this level there isn't much to do
	//nil'ing the location is potentially dangerous.
	return ted
}

//////////////////////////////////////////////////
//// MultilineTextEvent
//////////////////////////////////////////////////

type MultilineTextEvent interface {
	TextEvent
	Font
	GetText() []TextLine
	AddTextLine(TextLine) MultilineTextEvent
	AddText(string) MultilineTextEvent
	GetId() int32
}

// ////
// Text event can have a font specified for defaults.
// TODO switch to using Point for location.
type multilineTextEventData struct {
	textEventData
	Font
	text []TextLine
	id   int32
}

var nextTextEventId = int32(0)

func NewMultilineTextEvent(font string, fontSize, x, y int32) MultilineTextEvent {
	nextTextEventId++ //Oddly can't do this in the struct
	return &multilineTextEventData{textEventData{x, y}, NewFont(font, fontSize), make([]TextLine, 0), nextTextEventId}

}

// Reuse the text event with font/size information but clear the slice and update the id
// This is problematic for the builder pattern.

func (e *multilineTextEventData) Reinitialize() TextEvent {
	e.text = make([]TextLine, 0)
	e.id = e.id + 1
	return e
}

func (e *multilineTextEventData) AddText(text string) MultilineTextEvent {
	e.text = append(e.text, NewTextLine(text, e.Font))
	return e
}

// If one wishes to specify diferent font or font size.

func (e *multilineTextEventData) AddTextLine(text TextLine) MultilineTextEvent {
	e.text = append(e.text, text)
	return e
}
func (e *multilineTextEventData) GetText() []TextLine {
	return e.text
}

func (p *multilineTextEventData) GetId() int32 {
	return p.id
}
