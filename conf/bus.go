package conf

// Threhold threhold
type Threhold struct {
	MaxTemp float32
	MinTemp float32
}

// Bus bus
type Bus struct {
	Threhold *Threhold
}
