package main

import (
	"encoding/json"
	"fmt"
)

// type _interface interface {
// 	I()
// }

// type _struct struct {
// 	V int `json:"v"`
// }

// type _outer struct {
// 	S _interface `json:"s"`
// }

// func (_v *_struct) I() {}

// type base struct {
// 	Prop1 int `json:"prop_1"`
// }

// type record map[string]*base

// type final struct {
// 	*base
// 	ByName record `json:"by_name"`
// }

// func readFromChan(ch chan int) {
// 	for i := range ch {
// 		fmt.Println("read: ", i)
// 	}
// 	fmt.Println("end of read")
// }

func main() {
	// name := "daniel"
	// last := "soler"
	// // dd := []*string{&name, nil, &last}
	// dd := make([]*string, 4)
	// dd[0] = &name
	// dd[2] = &last
	// bb, _ := json.Marshal(&dd)
	// fmt.Printf("bb: %s\n", bb)
	// last = "hedez"
	// bb, _ = json.Marshal(&dd)
	// fmt.Printf("bb: %s\n", bb)
	// data := `{"prop_1":1}`

	// a := map[string]int{
	// 	"a": 1,
	// 	"b": 2,
	// }
	// fmt.Printf("a: %v\n", a)
	// for key := range a {
	// 	a[key]--
	// }
	// fmt.Printf("a: %v\n", a)
	// fmt.Printf("(%d) binary: %016b, (%d) %016b\n", 0xE, 0xE, 0xE<<2, 0xE<<2)

	// data := `[1,2,3]`
	// var numbers []any
	// err := json.Unmarshal([]byte(data), &numbers)
	// fmt.Printf("err: %v\n", err)
	// fmt.Printf("numbers: %v\n", numbers)
	// v := int(numbers[2].(int64))
	// fmt.Printf("v: %v\n", v)
	//
	// var twonumbers [2]int
	// err = json.Unmarshal([]byte(data), &twonumbers)
	// fmt.Printf("err: %v\n", err)
	// fmt.Printf("twonumbers: %v\n", twonumbers)
	//
	// var fivenumbers [5]int
	// err = json.Unmarshal([]byte(data), &fivenumbers)
	// fmt.Printf("err: %v\n", err)
	// fmt.Printf("fivenumbers: %v\n", fivenumbers)

	// a := make([]int, 0, 3)
	// a = append(a, 0)
	// a = append(a, 1)
	// a = append(a, 2)
	// fmt.Printf("a: %v\n", a)
	// b := a[1:]
	// fmt.Printf("b: %v\n", b)
	// b = append(b, 3)
	// fmt.Printf("b: %v\n", b)
	// fmt.Printf("a: %v\n", a)
	testLabel()

	// val := &final{}

	// err := json.Unmarshal([]byte(data), val)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(val)
	// fmt.Println(val.ByName["susan"])
	// var value = _outer{S: &_struct{V: 12}}

	// data, err := json.Marshal(value)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%s\n", data)
	// nan := math.NaN()
	// n := rand.Float64()
	// fmt.Printf("%v > %v == %v\n", nan, nan, nan > nan)
	// fmt.Printf("%v < %v == %v\n", nan, nan, nan < nan)
	// fmt.Printf("%v == %v == %v\n", nan, nan, nan == nan)

	// fmt.Printf("%v == %v == %v\n", nan, n, nan > n)
	// fmt.Printf("%v == %v == %v\n", nan, n, nan < n)
	// fmt.Printf("%v == %v == %v\n", nan, n, nan == n)

	// type foo struct {
	// 	Timestamp time.Time
	// }

	// var f = &foo{time.Now().UTC()}
	// data, err := json.Marshal(f)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(string(data))
	// }

	// _f := new(foo)
	// err = json.Unmarshal(data, _f)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(_f)
	// }

	// ch := make(chan int)
	// go readFromChan(ch)

	// time.Sleep(2 * time.Second)
	// for i := 0; i < 10; i++ {
	// 	ch <- i
	// }
	// close(ch)

	// m := map[string][]int{}

	// m["da"] = append(m["da"], 1)

	// fmt.Println(m)
	// err := errors.Join(fmt.Errorf("this is the first error"), fmt.Errorf("this is the second error"))

	// fmt.Printf("err: %s\n", err.Error())

	// heapProf("./heap_before.pprof")
	// printAlloc()
	// foo := initFoo()

	// // printAlloc()
	// fmt.Println(foo)

	// // two := keepFirstTwoElementsOnly(foos)
	// heapProf("./heap_after.pprof")
	// printAlloc()
	// runtime.KeepAlive(two)

	// tt()
	// msg := make(chan byte, 1)

	// select {
	// case msg <- 1:
	// 	fmt
	// }

	// close(ch)
	// heapProf("./heap_after.pprof")
	// time.Sleep(2 * time.Second)
}

// func tt() {
// 	f := &Foo{
// 		ch: make(chan bool, 1),
// 	}
// 	f.ch <- true
// 	go f.routine()
// }

// type Foo struct {
// 	ch chan bool
// }

// func (f *Foo) routine() {
// 	f.ch <- true
// 	fmt.Println("finish the write to the channel")
// }

// func keepFirstTwoElementsOnly(foos []Foo) []Foo {
// 	return foos[2:]
// }

// func initFoo() *Foo {

// 	foo := &Foo{
// 		v: make(chan []byte, 1),
// 	}
// 	foo.v <- make([]byte, 1024*1024)

// 	return foo
// }

// func printAlloc() {
// 	var m runtime.MemStats
// 	runtime.ReadMemStats(&m)
// 	fmt.Printf("%d KB\n", m.Alloc/1024)
// }

// func heapProf(pprofName string) error {
// 	f, err := os.Create(pprofName)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close() // error handling omitted for example
// 	runtime.GC()
// 	// return nil
// 	return pprof.WriteHeapProfile(f)
// }

type label struct {
	L string `json:"l,omitempty"`
}

func testLabel() {
	l := label{L: "super"}

	data, err := json.Marshal(&l)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("data: %s\n", data)
}
