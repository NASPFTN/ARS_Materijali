package main // deklaracija paketa

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
)

func add(x int, y int) int {
	return x + y
}

func fact(n int) int {
	if n < 1 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

func isPrime(n int) bool {
	//6k + 1 || 6k -1 except 2 and 3
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func f() func() int {
	i := 0
	return func() int {
		i += 10
		return i
	}
}

type Person struct {
	firstName string
	lastName  string
	balance   float64
	personID  string
}

func pointersExample() {
	i := 42
	p := &i         // pokazivac na i
	fmt.Println(*p) // deferenciranje pokazivaca
	*p = 21         // postavljanje vrednosti i
	fmt.Println(*p)
	fmt.Println(i) // ispisi i
}

func multiply(a int) func(int) int {
	return func(i int) int {
		return a * i
	}
}

type Osoba interface {
	toString() string
}

type Student struct {
	ime, prz, brIndeksa string
}
type Radnik struct {
	ime, prz, jmbg string
}

func (s Student) toString() string {
	return "Student[" + s.ime + " , " + s.prz + " ," + s.brIndeksa + "]"
}
func (r Radnik) toString() string {
	return "Radnik[" + r.ime + " , " + r.prz + " ," + r.jmbg + "]"
}

func display(o Osoba) {
	fmt.Println(o.toString())
}

func hello(from string) {
	for i := 1; i < 100000000; i++ {
	}
	fmt.Println("Hello from : " + from)
}

func sumArrayIntoChannel(s []int, c chan int, ordNum int) {
	fmt.Printf("Started routine %d\n", ordNum)
	fmt.Println("Slice from routine", ordNum, ": ", s)
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
	fmt.Printf("Finished routine %d\n", ordNum)
}

// ulazna tacka programa
func main() {
	fmt.Println("Hello World from Go")

	var a int // neinicijalizovana int promenljiva ima vrednost 0

	// deklaracija i inicijalizacija
	var b = 5
	var c, d = 5, 6
	s := "Hello World"

	// alternativni nacin deklaracije
	/*var (
		b, c,d = 5,5,6
		s = "hello"
	)*/

	fmt.Print("Enter a number: ")
	if _, err := fmt.Scanf("%d", &a); err != nil { // citanje sa standardnog ulaza
		log.Fatal(err)
		return
	}
	fmt.Printf("Number is %d \n", b)
	fmt.Printf("Number c is: %d, number d: %d \n", c, d)
	fmt.Println(s)

	if a%2 == 0 {
		fmt.Printf("%d is an even number", a)
	} else { // } mora da bude u istom redu kao i else
		fmt.Printf("%d is an odd number", a)
	}

	fmt.Print("Go runs on ")
	// goos:=runtime.GOOS; se izvrši neposredno pre switch dela
	switch goos := runtime.GOOS; goos {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", goos)
	}

	for sum, i := 0, 0; i < 10; i++ {
		sum += i
	}

	var result = add(42, 13)
	var resultPlus5 = result + 5
	fmt.Printf("Result is %d, %d", result, resultPlus5)
	fmt.Println("42 is prime: ", isPrime(42))

	i := 42
	e := func() {
		j := i / 2 // ima pristup i
		fmt.Println(j)
	}
	e() // 21

	g := f()
	h := f()
	fmt.Println(g()) // 10
	fmt.Println(h()) // 10
	g()              // pozovi ponovo
	fmt.Println(g()) // 30
	fmt.Println(g()) // 20

	multiplyBy4 := multiply(4)
	fmt.Println("5 * 4: ", multiplyBy4(5))

	if _, err := os.Create("delete_me"); err != nil { // Može se otvoriti fajl u režimu koji ga kreira ako ne postoji
		log.Fatal(err)
		return
	}
	file, err := os.OpenFile("delete_me", os.O_RDWR, os.ModeAppend)
	defer file.Close()
	if err != nil {
		log.Printf("can’t openfile")
	}
	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString("1. row\n"); err != nil {
		log.Fatal(err)
		return
	}
	_ = writer.Flush()

	pointersExample()

	var names [2]string
	names[0] = "Marc"
	names[1] = "John "
	fmt.Println(names[0], names[1])
	fmt.Println(names)

	slice := make([]int, 5) // dinamički kreiraj slice
	fmt.Println(slice)
	slice = append(slice, 3)
	fmt.Println(slice)

	student := Student{"marko", "markovic", "ee-222/2012"}
	radnik := Radnik{"rastko", "raicevic", "055121312321312"}
	display(student)
	display(radnik)

	hello("program")
	go hello("Go routine")
	go func(caller string) {
		fmt.Println("Anonymous f: called by " + caller)
	}("Go routine")
	_, _ = fmt.Scanln()

	ss := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	cc := make(chan int)
	go sumArrayIntoChannel(ss[:5], cc, 1)
	go sumArrayIntoChannel(ss[5:10], cc, 2)
	go sumArrayIntoChannel(ss[10:15], cc, 3)
	go sumArrayIntoChannel(ss[15:], cc, 4)
	x, y, z, v := <-cc, <-cc, <-cc, <-cc // citaj iz kanala
	fmt.Println(x, y, z, v, x+y+z+v)
}
