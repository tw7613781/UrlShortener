# Study note of this project   

Golang is a really intersting language. At one hand, it has a fast computing ability, just like C language, only use struct, no class. Like a low-level programming language. At another hand, it has many features that facilitates fast coding and flexibility, such as built-in async call, package management, slicing. Like a high-level language. 

Below is my study note for the project.

## makefile  
- Makefile is a configure file of *make utility*, which has been pre-installed by almost every Linux OS or MacOS
- Simple syntax of makefile is:   
    ~~~~
    target: prerequisites      
    <TAB>recipe
    ~~~~
    For example
    ~~~~
    hello: main.o
        g++ -g -o hello main.c
    main.o: main.cpp
        g++ -c -g main.cpp
- Special keywords     
all: the default targets that need to make   
.PHONY: the targets that no need to be created physically

## Golang program structure   
- Program is made up of packages.
- Many different .go files can belong to one package, so the filename and package name are generally not the same   
- A program must have a "package main", in which there is a "func main", which is the entry point of the program.
- In the first line of your .go file, there should states the package name.
- Visibility rule: When the identifier starts with an *uppercase* letter, then the object with this identifier is visible in code outside the package, like Public in OO language


## defer keyword   
- A special keyword in Golang.   
- The English meaning of "defer" is "delay" or "postpone", so the defer keyword allows us to postpone the execution of a statement or a function until the end of the enclosing function.
- Usage Pattern:    
 ~~~~
 import "sync"

 mu := sync.Mutex
 mu.Lock()
 defer mu.Unlock()
 // continue with thread safe code
 ~~~~
The same code without defer keyword

 ~~~~
 import "sync"

 mu := sync.Mutex
 mu.Lock()
// continue with thread safe code
 mu.Unlock()
 ~~~~
So basically, the keyword protected you from forgetting the remaining half    
Another example    
~~~~
//Open a file
defer file.Close()
balalalal
~~~~


## new, make and initial a struct object
- new: new(T) returns a pointer to a newly allocated, *zeroed object T*, that is, a pointer to a nil value
- make: make(T, args) serves a purpose different from new(T). *It creates slices, maps, and channels only*, and it returns an initialized (not zeroed) value of type T (not *T)
- Example:   
~~~~
package main
import "fmt"
func main(){
    // declare a slice with size 10 and initial with value 0 
    test1 := make([]int, 10) 
    // declare a pointer to a slice, the value is nil
    test2 := new([]int)
    fmt.Println(test1[4])
    // if fmt.Println(test2[4]) will occur error
    fmt.Println(test2)
}
// output
max@max ~/code/UrlShortener (master)
$ go run test.go 
0
&[]
~~~~
- If need to initialize a struct object, need to use a special function (The same with construction function in other language). The function starts with New, concatenate the struct name.    
~~~~
type URLStore struct{
	urls map[string]string
	mu sync.RWMutex
}

func NewURLStore() *URLStore{
	return &URLStore{ urls: make(map[string]string)}
}

// will initial the URLStore object *store* with     
// a initialized map object inside.
var store = NewURLStore()
~~~~


## slice and array
- array: An array type definition specifies a length and an element type. For example, the type [4]int represents an array of four integers. An array's size is fixed; its length is part of its type ([4]int and [5]int are distinct, incompatible types).   
- slice: Since array is inflexible, you don't see them too often in Go code. Slices, though, are everywhere. The type specification for a slice is []T, where T is the type of the elements of the slice. Unlike an array type, a slice type has no specified length.
- Example:    
~~~~
package main
import "fmt"

func main() {
    // declare a array, type is [6]int
    // means a array with element type is int, and length is 6
    var arr1 [6]int
    // declare a slice, including arr1[2], arr1[3], arr1[4]
    var slice1 []int = arr1[2:5]
}
~~~~

## goroutine, channel, async
- What am I saying is not 100% precise, I just state my understanding.
- goroutine: it's a go abstract concept of Golang beyond OS thread. So basically, it can regards as thread. Start with keyword go + function name (The logic running inside the goroutine)
~~~~
go test()
func test() {
    //
}
~~~~
- channel: basically it's a message queue between goroutines. It's like a producer and consumer mode. One side push data into channel, another side pull data from channel.
~~~~
// type means the element type inside the queue.
var identifier chan type
such:
var save chan string
~~~~

- useage pattern   
Like the callback function in other language, you can define a function that running on goroutine. The function will actually process the data be passed in. In main logic, just send data to the goroutine by channel.
~~~~
package main

import (
	"fmt"
	"time"
)

func main() {
    // initial a channel with make func
	ch := make(chan string)

    // define two goroutines, one is producer, one is consumer
	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {
    // "<-" operator means push data into channel
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func getData(ch chan string) {
	var input string
	for {
        // "<-" operator also menas pull data from channel
        // The defference is the which side the channel object on
		input = <-ch
		fmt.Printf("%s ", input)
	}
}

output:
Washington Tripoli London Beijing tokyo
~~~~


## error handling pattern
In golang code, there is a commom pattern to handle error.    
Before introduce the pattern, we need to study a syntax first.
~~~~
// a special syntax for "if" in golang
if an-assignment; condition {
    //code
}
example
if x := 3; x>2 {
    //code
}
~~~~
For functions that the return value type is error, we can handle like below
~~~~
if err := func1(); err != nil {
    //error handle code
}
~~~~
