package wavegenerator

type WaveGenerator interface {
	One(duration uint32)
	Zero(duration uint32)
}
