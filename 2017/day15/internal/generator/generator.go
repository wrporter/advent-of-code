package generator

const iterations = 500_000

func CountSynchronously() int {
	numEqual := 0

	a := 0
	b := 0

	for i := 0; i < iterations; i++ {
		for {
			a += 1
			if a%4 == 0 {
				break
			}
		}

		for {
			b += 1
			if b%8 == 0 {
				break
			}
		}

		if a&0xff == b&0xff {
			numEqual++
		}
	}

	return numEqual
}

// Why does using channels slow this down rather than running things synchronously?
func CountWithChannels() int {
	abortA := make(chan int)
	generatorA := generate(abortA, 4)
	abortB := make(chan int)
	generatorB := generate(abortB, 8)

	numEqual := 0
	for i := 0; i < iterations; i++ {
		a := <-generatorA
		b := <-generatorB

		if a&0xff == b&0xff {
			numEqual++
		}
	}

	close(abortA)
	close(abortB)

	return numEqual
}

func generate(abort <-chan int, modulus int) <-chan int {
	channel := make(chan int)
	value := 0

	go func() {
		defer close(channel)
		for i := 0; ; i++ {
			select {
			case <-abort:
				return
			default:
				value += 1
				if (value % modulus) == 0 {
					channel <- value
				}
			}
		}
	}()

	return channel
}
