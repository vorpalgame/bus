package bus

// The peer raylibPeerController is to mediate any impedance mismatch between the tight single
// threaded loop of the Raylib engine and the concurrent mechanisms of Golang.
// Note: Raylib is very fast but very intolerant of mutations on any data it is using.
// So the peer and MediaCache can keep off thread data for rendering, listening and
// sending. Raylib loop will query this raylibPeerController for data it needs.
func NewRaylibPeerController() RaylibPeerController {
	controller := raylibPeerController{}
	return &controller
}

type RaylibPeerController interface {
	DrawEventListener
	AudioEventListener
	TextEventListener
	KeysRegistrationListener
	ControlEventListener
	GetControlEvents() []ControlEvent
	GetDrawEvent() DrawEvent
	GetAudioEvent() AudioEvent //One event at at time...
	GetTextEvent() TextEvent
	GetKeysRegistrationEvent() KeysRegistrationEvent
}

type raylibPeerController struct {
	bus                   VorpalBus
	drawEvent             DrawEvent
	audioEvent            []AudioEvent //Different audio events for stop, start, etc. so they need to be kept in slice for processing.
	textEvent             TextEvent
	keysRegistrationEvent KeysRegistrationEvent //Only one set of keys to listen for at a time.
	controlEvents         []ControlEvent        //May need slice..
}

func (c *raylibPeerController) OnControlEvent(controlChannel <-chan ControlEvent) {
	for evt := range controlChannel {
		c.controlEvents = append(c.controlEvents, evt)
	}
}

func (c *raylibPeerController) OnDrawEvent(drawChannel <-chan DrawEvent) {
	for evt := range drawChannel {
		c.drawEvent = evt
	}

}

func (c *raylibPeerController) OnKeyRegistrationEvent(keyRegistrationChannel <-chan KeysRegistrationEvent) {
	for evt := range keyRegistrationChannel {
		c.keysRegistrationEvent = evt
	}
}

func (c *raylibPeerController) OnAudioEvent(audioChannel <-chan AudioEvent) {
	for evt := range audioChannel {
		c.audioEvent = append(c.audioEvent, evt)
	}
}

func (c *raylibPeerController) OnTextEvent(textChannel <-chan TextEvent) {
	for evt := range textChannel {
		c.textEvent = evt

	}
}

func (c *raylibPeerController) GetDrawEvent() DrawEvent {
	//evt := c.drawEvent
	//c.drawEvent = nil
	return c.drawEvent
}

func (c *raylibPeerController) GetAudioEvent() AudioEvent {
	var evt AudioEvent = nil
	if len(c.audioEvent) > 0 {
		evt, c.audioEvent = c.audioEvent[0], c.audioEvent[1:]
	}
	return evt
}

// Don't repeat process.
func (c *raylibPeerController) GetControlEvents() []ControlEvent {
	temp := c.controlEvents
	c.controlEvents = nil
	return temp
}

func (c *raylibPeerController) GetTextEvent() TextEvent {
	//temp := c.textEvent
	//c.textEvent = nil
	return c.textEvent
}

func (c *raylibPeerController) GetKeysRegistrationEvent() KeysRegistrationEvent {
	return c.keysRegistrationEvent
}
