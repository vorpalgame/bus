package bus

type TextEvent interface{}

type TextEventListener interface {
	OnTextEvent(textChannel <-chan TextEvent)
}

type DrawEvent interface {
}

type DrawEventListener interface {
	OnDrawEvent(drawChannel <-chan DrawEvent)
}

type AudioEvent interface{}

type AudioEventListener interface {
	OnAudioEvent(audioChannel <-chan AudioEvent)
}

type ControlEvent interface{}

type ControlEventListener interface {
	OnControlEvent(controlChannel <-chan ControlEvent)
}

type MouseEvent interface{}

type MouseEventListener interface {
	OnMouseEvent(mouseChannel <-chan MouseEvent)
}
type KeyEvent interface{}

type KeyEventListener interface {
	OnKeyEvent(keyChannel <-chan KeyEvent)
}
