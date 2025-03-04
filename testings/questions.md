# Golang Technical Assessment

Total points: 50

## Basic Concepts (15 points)

1. What is the zero value of a pointer in Go? (1 point)

- [ ] undefined
- [x] nil
- [ ] 0
- [ ] void


2. Which of the following declarations are valid in Go? (1 point)

- [x] var x = 42
- [x] x := 42
- [ ] x =: 42
- [x] var x int = 42

2. What happens when you try to access a map key that doesn't exist? (2 points)

- [ ] The program panics
- [ ] An error is returned
- [x] The zero value of the value type is returned
- [ ] nil is returned

3. What happens in the following code (3 points)

```go
func main() {
    var m map[string]int
    m["1"] = 1
    fmt.Println(m["1"])
}
```

- [ ] The main function return without print
- [ ] The main function prints `1` to the standard output and returns 
- [ ] The main function panics cause it cannot print a number
- [x] The main function panics cause it cannot assign a value to a map that's not initialized

4. Which statements about goroutines are correct? (2 points)

- [x] They are lightweight threads managed by the Go runtime
- [x] Multiple goroutines can run concurrently
- [ ] They always execute in the order they are created
- [x] They share the same address space

5. What is the correct way to create an unbuffered channel (3 points)

- [ ] var ch chan int
- [ ] var ch = make(chan int, 1)
- [x] ch := make(chan int)
- [x] ch := make(chan int, 0)

5. What is the purpose of the defer statement? (1 point)

- [x] Delays the execution of a function until the surrounding function returns
- [ ] Permanently cancels the execution of a function
- [ ] Creates a new goroutine
- [ ] Pauses the execution of a function for a specified duration

6. How can a program recover from a panic (2 points)

- [ ] A panic cannot be recovered
- [ ] By using a catch in the main
- [ ] By using recover in the main
- [x] By using a defer function with a recover in one of the functions of the panic call stack.

## Intermediate Concepts (20 points)

7. Which of the following are true about interfaces in Go? (2 points)

- [x] A type implements an interface by implementing its methods
- [x] Interfaces can be implemented implicitly
- [x] An empty interface can hold values of any type
- [ ] Interfaces must be explicitly declared with the 'implements' keyword

8. What is the output of this code? (3 points)

```go
func main() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    close(ch)
    for v := range ch {
        fmt.Println(v)
    }
}
```

- [x] It prints 1 then 2
- [ ] It causes a deadlock
- [ ] It prints nothing
- [ ] It panics

9. Which statements about slices are correct? (2 points)

- [x] A slice is a reference to an array
- [x] The capacity of a slice can be larger than its length
- [ ] Slices are fixed-size data structures
- [x] append() may create a new underlying array

10. What happens in this code? (3 points)

```go
func main() {
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go func(n int) {
            fmt.Println(n)
            wg.Done()
        }(i)
    }
    wg.Wait()
}
```

- [x] Always prints 0, 1, 2 in any order
- [ ] Causes a race condition
- [ ] Might print 3 three times
- [x] The order of numbers is not guaranteed

11. What issues this statement can have `i := v.(int)` (2 points)

- [ ] The value `i` cannot be an int
- [ ] The cast to int from float may overflow
- [ ] The value of v may be undefined
- [x] The underline value hold by `v` be may not be an int

12. How can this code panic: (2 points)

```go
func f(x any, y any) int {
    return x.(int) / y.(int)
}
```

- [x] `x` may not be an int
- [x] `y` may be 0
- [ ] `any` cannot be casted to int due to unsafe code
- [ ] another part of the program can introduce a panic and affect this function

13. What will be the output of this code? (3 points)

```go
type Number interface {
    int64 | float64
}

func min[T Number](x, y T) T {
    if x < y {
        return x
    }
    return y
}

func main() {
    fmt.Println(min(2.5, 1.5))
    fmt.Println(min(int64(10), int64(5)))
}
```

- [x] It will compile and print 1.5 and 5
- [ ] It will fail to compile due to type constraints
- [ ] It will compile but panic at runtime
- [ ] It will print 2.5 and 10

14. Consider this code using mutexes: (3 points)

```go
type Counter struct {
    mu sync.Mutex
    count int
}

func (c *Counter) Increment() {
    go func() {
        c.mu.Lock()
        c.count++
        go func() {
            c.count++
            c.mu.Unlock()
        }()
    }()
}
```

- [ ] This is a thread-safe implementation
- [x] This can cause a deadlock
- [x] The mutex unlock should be in the same goroutine as the lock
- [x] This code has potential race conditions

### Advanced Concepts (15 points)

15. Which statements about reflection in Go are correct? (3 points)

- [x] It allows inspection of types at runtime
- [x] It can be used to modify struct fields dynamically
- [x] It may impact performance
- [ ] It's the recommended way to handle type conversion


16. What is true about the context package? (3 points)

- [x] It can be used to carry deadlines and cancellation signals
- [x] Context should be the first parameter of a request-handler function
- [x] It can propagate request-scoped values
- [ ] Contexts cannot be nested


17. Which statements about memory management in Go are correct? (3 points)

- [x] Go uses garbage collection
- [x] The runtime can run garbage collection concurrently
- [x] Memory is allocated on the heap or stack
- [ ] Developers must manually manage memory

18. What is the output of this code involving channels and select? (3 points)

```go
func main() {
    ch1 := make(chan int, 1)
    ch2 := make(chan int, 1)
    ch1 <- 1
    close(ch1)
    close(ch2)
    
    select {
    case v1, ok1 := <-ch1:
        fmt.Printf("v1=%d, ok1=%v\n", v1, ok1)
    case v2, ok2 := <-ch2:
        fmt.Printf("v2=%d, ok2=%v\n", v2, ok2)
    default:
        fmt.Println("default case")
    }
}
```

- [x] It will print "v1=1, ok1=true"
- [ ] It will print "v2=0, ok2=false"
- [ ] It will print "default case"
- [ ] It will panic

19. What will arr1 contain at the end of this code (3 points)

```go
func f() {
    arr1 := []int{1,2,3}
    arr2 := arr1[:1]
    arr2 = append(arr2, 4)
    arr2[0] = 10
}
```

- [ ] `[1,2,3]`
- [ ] `[1,2,4]`
- [ ] `[1,4,3]`
- [x] `[10,4,3]`