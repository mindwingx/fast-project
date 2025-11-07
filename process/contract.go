package process

type IFlow interface {
	Process(item Items)
	Next(next IFlow)
}
