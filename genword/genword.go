package genword


// GenHelloWorld is function looping to generate word
func GenHelloWorld(word string, num int, split string) string {
	rs := ""
	for i := 0; i < num; i++ {
		rs += word + split
	}

	return rs
}
