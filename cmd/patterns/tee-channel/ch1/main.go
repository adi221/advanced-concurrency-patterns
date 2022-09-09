package main

/*
We’re going to use one select statement so that writes to out1 and out2 don’t block each other.
To ensure both are written to, we’ll perform two iterations of the select statement: one for each outbound channel.
Once we’ve written to a channel,
we set its shadowed copy to nil so that further writes will block and the other channel may continue.
*/

func main() {
	// done := make(chan interface{})
	// defer close(done)

	// out1, out2 := concepts.Tee(done, take(done, repeat(done, 1, 2), 4))

	// for val1 := range out1 {
	// 	fmt.Printf("out1: %v, out2: %v\n", val1, <-out2)
	// }
}
