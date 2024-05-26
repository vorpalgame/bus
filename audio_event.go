package bus

//Having to return the struct instead of the interface is an unfortunate
//side effect of the YAML marshaling.

func NewAudioState(fileName string, loop bool) AudioState {
	return &AudioStateData{fileName, loop, 0}
}

type AudioState interface {
	SetAudioFile(fileName string) AudioState
	SetAudioLoop(shouldLoop bool) AudioState
	IsAudioOnLoop() bool
	GetAudioFile() string
	IncrementAudio() int32
	ResetAudioCount() AudioState
}

type AudioStateData struct {
	AudioFile string `yaml:"AudioFile"`
	LoopAudio bool   `yaml:"LoopAudio"`
	count     int32  `yaml:"-"`
}

func (s *AudioStateData) IncrementAudio() int32 {
	s.count++
	return s.count
}

func (s *AudioStateData) ResetAudioCount() AudioState {
	s.count = 0
	return s
}

func (s *AudioStateData) SetAudioLoop(shouldLoop bool) AudioState {
	s.LoopAudio = shouldLoop
	return s
}

func (s *AudioStateData) IsAudioOnLoop() bool {
	return s.LoopAudio
}

func (s *AudioStateData) GetAudioFile() string {
	return s.AudioFile
}
func (s *AudioStateData) SetAudioFile(fileName string) AudioState {
	s.AudioFile = fileName
	return s
}

func NewPlayAudioEvent(state AudioState) PlayAudioEvent {
	return &playAudioEventData{state.GetAudioFile(), state.IsAudioOnLoop()}
}

func NewStopAudioEvent(state AudioState) StopAudioEvent {
	return &stopAudioEventData{state.GetAudioFile()}
}

/////PlayAudioEvent
//Need to ensure there is asymmetery in events/implementations they can be
//distinguished by the case switch.

type PlayAudioEvent interface {
	AudioEvent
	GetAudioFile() string
	IsLoop() bool
}

type playAudioEventData struct {
	fileName string
	isLoop   bool
}

func (p *playAudioEventData) GetAudioFile() string {
	return p.fileName
}

func (p *playAudioEventData) IsLoop() bool {
	return p.isLoop
}

/////StopAudioEvent

type StopAudioEvent interface {
	AudioEvent
}

type stopAudioEventData struct {
	fileName string
}

func (s *stopAudioEventData) GetAudioFile() string {
	return s.fileName
}
