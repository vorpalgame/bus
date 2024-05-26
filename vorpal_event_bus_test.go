package bus

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/mouse"
	"testing"
	"time"
)

func TestVorpalBus(t *testing.T) {
	bus := GetVorpalBus()

	bus.AddAudioListener(&testVals)
	bus.AddKeyListener(&testVals)
	bus.AddMouseListener(&testVals)
	bus.AddTextListener(&testVals)

	assert.Nil(t, testVals.textEvent)
	assert.Nil(t, testVals.audioEvent)
	assert.Nil(t, testVals.keyEvent)
	assert.Nil(t, testVals.mouseEvent)
	//Test is not concerned with internal data but on
	//dispatch.
	bus.SendTextEvent(NewMultilineTextEvent("", 0, 0, 0))
	bus.SendAudioEvent(NewPlayAudioEvent(NewAudioState("somefile.mp3", false)))
	var eventKey = key.Event{}
	eventKey.Rune = 55
	eventKey.Direction = key.DirPress
	eventKey.Code = key.Code(55)

	bus.SendKeyEvent(NewKeyEvent(eventKey))
	bus.SendMouseEvent(NewMouseEvent(mouse.Event{
		X:         0,
		Y:         0,
		Button:    0,
		Modifiers: 0,
		Direction: 0,
	}))
	//TODO Need better asynchronous callback notification.
	time.Sleep(1000 * time.Millisecond)

	assert.NotNil(t, testVals.mouseEvent)
	assert.NotNil(t, testVals.keyEvent)
	assert.NotNil(t, testVals.textEvent)
	assert.NotNil(t, testVals.audioEvent)
}

type bus_test struct {
	bus        VorpalBus
	textEvent  TextEvent
	mouseEvent MouseEvent
	keyEvent   KeyEvent
	audioEvent AudioEvent
}

var testVals = bus_test{}

func (t *bus_test) OnAudioEvent(channel <-chan AudioEvent) {
	for evt := range channel {
		t.audioEvent = evt
	}
}

func (t *bus_test) OnTextEvent(channel <-chan TextEvent) {
	for evt := range channel {
		t.textEvent = evt
	}
}

func (t *bus_test) OnKeyEvent(channel <-chan KeyEvent) {
	for evt := range channel {
		t.keyEvent = evt
	}

}
func (t *bus_test) OnMouseEvent(channel <-chan MouseEvent) {
	for evt := range channel {
		t.mouseEvent = evt

	}
}
