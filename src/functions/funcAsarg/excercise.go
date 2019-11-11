package main

func main() {
	// passing func1 as arg to func2
	// this is call callback
	println(func2(func1))
}

func func1() string {
	return "i am func1"

}

//passing a func as arg

func func2(f func() string) string {
	return f()

}
