package irwave

type IRWave interface {
	One(gpio uint32, duration uint32)
	Zero(gpio uint32, duration uint32)
}
