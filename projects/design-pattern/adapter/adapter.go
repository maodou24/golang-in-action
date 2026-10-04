package adapter

type Target interface {
	Play() string
}

type Adaptee interface {
	PlayMP3() string
}

func NewAdaptee() Adaptee {
	return &MP3Player{}
}

type MP3Player struct{}

func (m MP3Player) PlayMP3() string {
	return "PlayMP3"
}

type Adapter struct {
	adaptee Adaptee
}

func NewAdapter(a Adaptee) *Adapter {
	return &Adapter{adaptee: a}
}

func (a Adapter) Play() string {
	return a.adaptee.PlayMP3()
}
