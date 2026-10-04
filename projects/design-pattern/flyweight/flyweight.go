package flyweight

type Flyweight interface {
	Do() string
	Set(name string)
}

type FlyweightFactory struct {
	data map[string]Flyweight
}

func NewFlyweightFactory() FlyweightFactory {
	return FlyweightFactory{data: make(map[string]Flyweight)}
}

func (f *FlyweightFactory) Get(name string) Flyweight {
	if v, ok := f.data[name]; ok {
		return v
	}

	f.data[name] = &FlyweightImpl{}
	return f.data[name]
}

type FlyweightImpl struct {
	name string
}

func (f *FlyweightImpl) Do() string {
	return f.name
}

func (f *FlyweightImpl) Set(name string) {
	f.name = name
}
