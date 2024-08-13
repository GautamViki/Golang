package handler

type channel struct{}

func NewGoChannelHandler() *channel {
	return &channel{}
}

func (ch *channel) FibonacciBufferedChannels(n int, chn chan<- int) {
	first, second := 0, 1
	i := 0
	for i < n {
		chn <- first
		last := first + second
		first = second
		second = last
		i++
	}
	close(chn)
}
