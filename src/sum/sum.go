package sum

// sum the numbers in a and send the result on res.
func sum(array []int, result chan<- int) {
	value := 0
	for _, x := range array {
		value += x
	}
	result <- value
}

// concurrently sum the array a.
func ConcurrentSum(array []int, threads int) int {
	len := len(array)
	if threads > len {
		panic("too many threads")
	}
	concurrency_length := len / threads
	rest := len % threads
	channel := make(chan int)
	for x := 0; x < threads; x++ {
		go sum(array[(concurrency_length*x):(concurrency_length*(x+1))], channel)
	}
	if rest != 0 {
		go sum(array[(concurrency_length*threads):((concurrency_length*threads)+rest)], channel)
	}

	sum := 0
	for x := 0; x < threads; x++ {
		sum += <-channel
	}
	if rest != 0 {
		sum += <-channel
	}
	// TODO Get the subtotals from the channel and return their sum
	return sum
}

func main() {
	println("Hello world!")
}
