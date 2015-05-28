package main

import (
	"fmt"       // A package in the Go standard library.
	"io/ioutil" // Implements some I/O utility functions.
	m "math"    // Math library with local alias m.
	"net/http"  // Yes, a web server!
	"strconv"   // String conversions.
)

func main() {
	fmt.Println("Hello world!")

	beyondHello()
}

func beyondHello() {
	var x int
	x = 3
	y := 4
	sum, prod := learnMultiple(x, y)
	fmt.Println("sum:", sum, "prod:", prod)
	learnTypes()
}

func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y
}

func learnTypes() {
	str := "Learn Go!"
	s2 := `A "raw" string literal
can include line breaks.`

	g := 'Σ'
	f := 3.14195
	c := 3 + 4i

	var u uint = 7
	var pi float32 = 22. / 7

	n := byte('\n')

	var a4 [4]int
	a3 := [...]int{3, 1, 5}

	s3 := []int{4, 5, 9}
	s4 := make([]int, 4)
	var d2 [][]float64
	bs := []byte("a slice")

	s := []int{1, 2, 3}
	s = append(s, 4, 5, 6)
	fmt.Println(s)

	s = append(s, []int{7, 8, 9}...)
	fmt.Println(s)

	p, q := learnMemory()
	fmt.Println(*p, *q)

	m := map[string]int{"three": 3, "four": 4}
	m["one"] = 1

	_, _, _, _, _, _, _, _, _, _ = str, s2, g, f, u, pi, n, a3, s4, bs
	fmt.Println(s, c, a4, s3, d2, m)

	learnFlowControl()
}

func learnNamedReturns(x, y int) (z int) {
	z = x * y
	return
}

func learnMemory() (p, q *int) {
	p = new(int)

	s := make([]int, 20)
	s[3] = 7
	r := -2
	return &s[3], &r
}

func expensiveComputation() float64 {
	return m.Exp(10)
}

func learnFlowControl() {
	if true {
		fmt.Println("told ya")
	}

	if false {
		// Pout.
	} else {
		// Gloat.
	}

	x := 42.0
	switch x {
	case 0:
	case 1:
	case 42:
	case 43:

	default:
	}

	for x := 0; x < 3; x++ {
		fmt.Println("iteration", x)
	}

	for {
		break
		continue
	}

	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
		fmt.Printf("key=%s, value=%d\n", key, value)
	}

	if y := expensiveComputation(); y > x {
		x = y
	}

	xBig := func() bool {
		return x > 10000
	}
	fmt.Println("xBig:", xBig())
	x = 1.3e3
	fmt.Println("xBig:", xBig())

	fmt.Println("Add + double two numbers: ",
		func(a, b int) int {
			return (a + b) * 2
		}(10, 2))

	goto love
love:

	learnFunctionFactory() // func returning func is fun(3)(3)
	learnDefer()      // A quick detour to an important keyword.
	learnInterfaces() // Good stuff coming up!
}

func learnFunctionFactory() {

	fmt.Println(sentenceFactory("summer")("A beautiful", "day!"))

	d := sentenceFactory("summer")
	fmt.Println(d("A beautiful", "day!"))
	fmt.Println(d("A lazy", "afternoon!"))
}

func sentenceFactory(mystring string) func(before, after string) string {
	return func(before, after string) string {
		return fmt.Sprintf("%s %s %s", before, mystring, after) // new string
	}
}

func learnDefer() (ok bool) {
	// Deferred statements are executed just before the function returns.
	defer fmt.Println("deferred statements execute in reverse (LIFO) order.")
	defer fmt.Println("\nThis line is being printed first because")
	// Defer is commonly used to close a file, so the function closing the
	// file stays close to the function opening the file.
	return true
}

type Stringer interface {
	String() string
}

type pair struct {
	x, y int
}

func (p pair) String() string { // p is called the "receiver"
	// Sprintf is another public function in package fmt.
	// Dot syntax references fields of p.
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func learnInterfaces() {
	// Brace syntax is a "struct literal". It evaluates to an initialized
	// struct. The := syntax declares and initializes p to this struct.
	p := pair{3, 4}
	fmt.Println(p.String()) // Call String method of p, of type pair.
	var i Stringer          // Declare i of interface type Stringer.
	i = p                   // Valid because pair implements Stringer
	// Call String method of i, of type Stringer. Output same as above.
	fmt.Println(i.String())

	// Functions in the fmt package call the String method to ask an object
	// for a printable representation of itself.
	fmt.Println(p) // Output same as above. Println calls String method.
	fmt.Println(i) // Output same as above.

	learnVariadicParams("great", "learning", "here!")
}

// Functions can have variadic parameters.
func learnVariadicParams(myStrings ...interface{}) {
	// Iterate each value of the variadic.
	// The underbar here is ignoring the index argument of the array.
	for _, param := range myStrings {
		fmt.Println("param:", param)
	}

	// Pass variadic value as a variadic parameter.
	fmt.Println("params:", fmt.Sprintln(myStrings...))

	learnErrorHandling()
}

func learnErrorHandling() {
	m := map[int]string{3: "three", 4: "four"}
	if x, ok := m[1]; !ok { // ok will be false because 1 is not in the map.
		fmt.Println("no one there")
	} else {
		fmt.Print(x) // x would be the value, if it were in the map.
	}
	// An error value communicates not just "ok" but more about the problem.
	if _, err := strconv.Atoi("non-int"); err != nil { // _ discards value
		// prints 'strconv.ParseInt: parsing "non-int": invalid syntax'
		fmt.Println(err)
	}
	learnConcurrency()
}

func inc(i int, c chan int) {
	c <- i + 1 // <- is the "send" operator when a channel appears on the left.
}

func learnConcurrency() {
	c := make(chan int)

	go inc(0, c) // go is a statement that starts a new goroutine.
	go inc(10, c)
	go inc(-805, c)

	fmt.Println(<-c, <-c, <-c) // channel on right, <- is "receive" operator.

	cs := make(chan string)       // Another channel, this one handles strings.
	ccs := make(chan chan string) // A channel of string channels.
	go func() { c <- 84 }()       // Start a new goroutine just to send a value.
	go func() { cs <- "wordy" }() // Again, for cs this time.

	select {
	case i := <-c: // The value received can be assigned to a variable,
		fmt.Printf("it's a %T", i)
	case <-cs: // or the value received can be discarded.
		fmt.Println("it's a string")
	case <-ccs: // Empty channel, not ready for communication.
		fmt.Println("didn't happen.")
	}

	learnWebProgramming() // Go does it. You want to do it too.
}

func learnWebProgramming() {

	go func() {
		err := http.ListenAndServe(":8080", pair{})
		fmt.Println(err) // don't ignore errors
	}()

	requestServer()
}

func (p pair) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You learned Go in Y minutes!"))
}

func requestServer() {
	resp, err := http.Get("http://localhost:8080")
	fmt.Println(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("\nWebserver said: `%s`", string(body))
}
