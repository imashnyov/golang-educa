# Define variable
```
s := "" // practically use
var s string // practically use

var s = ""
var s string = ""
```
# Reserved names
```
break, case, chan, const, continue, for, import, interface, map, package, range, return, default, func, defer, go, else, goto, fallthrough, if, select, struct switch, type, var;
```
# Predefined names
#### Constants
```
true, false, iota, nil;
```
#### Types
```
int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, complex128, complex64, bool, byte, rune, string, error;
```
#### Functions
```
make, len, cap, new, append, copy, close, delete, complex, real, imag, panic, recover;
```
# Formatting strings
```
name := "John"
var age int = 30
value := 1000.476
a := true

fmt.Printf("My name is %s.\nMy age is %d.\nI have %f$.\n%t", name, age, value, a)
```
# Logical operators
```
a := 3
b := 10

if a > 1 && b > 9 {
	fmt.Println("Right!")
}
if a > 2 || b > 9 {
	fmt.Println("Right!")
}
if a != 2 {
	fmt.Println("Right!")
}
```
# Loops
```
for i := 1; i <= 5; i++ {
	fmt.Printf("Hello %d\n", i)
}
```
#### ∞
```
for {
	fmt.Printf("Loop")
}
```
#### more loop's options
```
for i := 1; i <= 5; i++ {
	if i%2 != 0 {
		continue
	}
	fmt.Println(i)
}

for i := 1; i <= 5; i++ {
	if i%2 != 0 {
		break
	}
	fmt.Println(i)
}
```
#### while loop
```
i := 0
for i < 5 {
	fmt.Println(i)
	i++
}

nums := []int{1, 2, 3, 4, 5}
for i := 0; i < len(nums); i++ {
	fmt.Println(nums[i])
}
```
# Switches
```
name := "John"

switch name {
case "Jordan":
	fmt.Println("Wrong")
case "Kate":
	fmt.Println("Wrong")
case "John":
	fmt.Println("Hello John!")
default:
	fmt.Println("Something else wrong!")
}
```
#### Switch's fallthrough 
```
number := 10

switch {
case number > 1:
	fmt.Println("Number is greater than 1")
	fallthrough
case number > 11:
	fmt.Println("Number < 11")
}
```
# Arrays
```
array := [3]string{"Denis", "Kate", "John"}
for i := 0; i < len(array); i++ {
	fmt.Println(names[i])
}

// array := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
```
# Slices
```
slice := []int{3, 1, 2, 5, 4, 7}
slice = append(slice, 0)
slice[0] = 0
sort.Ints(slice)
fmt.Println(slice)

//range slice
slice := []int{1, 6, 3, 8, 9, 0, 1, 2, 7}

for _, element := range slice {
	fmt.Printf("%d\n", element)
}
```
# Maps
```
/*
money = map[string]int{
	"dollars":         800,
	"euros":           300,
	"uah":             36600,
	"a_first_element": 1,
}
*/
var money map[string]int = map[string]int{
	"dollars":         800,
	"euros":           300,
	"uah":             36600,
	"a_first_element": 1,
}
fmt.Println(money)
fmt.Println(money["dollars"])
money["dollars"] = 820
delete(money, "a_first_element")
fmt.Println(money)

// Find in map
element, status := money["dollars"]
fmt.Println(element, status)
```
# Functions
```
func sum(a int, b int) int {
	result := a + b
	return result
}

func main() {
	value := sum(5, 9)
	fmt.Println(value)
}
```

```
func math_func(a int, b int) (int, int, int, int) {
	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b

	return sum, sub, mul, div
}

func main() {
	sum, sub, mul, div := math_func(3, 5)
	fmt.Println(sum, sub, mul, div)
}

// Better return format
func math_func(a int, b int) (sum int, sub int, mul int, div int) {
	sum = a + b
	sub = a - b
	mul = a * b
	div = a / b

	return
}

func main() {
	sum, sub, mul, div := math_func(3, 5)
	fmt.Println(sum, sub, mul, div)
}

func main() {
	a := func(x string, y string) string {
		return x + y
	}
	result := a("Hi ", "there!")
	fmt.Println(result)
}
```
# Pointers
```
func changer(str *string) {
	*str = "Bye!"
}

func main() {
	s := "Hello!"
	fmt.Println(s)
	// var pointer *string = &s
	// changer(pointer)
	changer(&s)
	fmt.Println(s)
}
```
# Structures
```
type User struct {
	name     string
	age      int8
	password string
}

func changer(u *User) {
	u.name = "Kate"
}

func main() {
	//var John User = User{name: "John", age: 87, password: "123456"}
	John := User{"John", 87, "123456"}
	fmt.Println(John)
	John.name = "Mark"
	fmt.Println(John.name)
	changer(&John)
	fmt.Println(John)
}
```
#### Structure's methods
```
type User struct {
	name     string
	age      int8
	password string
}

func (u User) getName() string {
	return u.name
}

func (u *User) setName(name1 string) {
	u.name = name1
}

func (u User) isElder() bool {
	a := u.age
	isTrue := false

	if a >= 18 {
		isTrue = true
	} else if a < 18 {
		isTrue = false
	}

	return isTrue
}

func main() {
	user := User{"John", 17, "123456"}
	user.setName("Den")
	fmt.Println(user.getName())
	fmt.Println(user.isElder())
}
```
# Interfaces
```
type Numbers struct {
	num1 int
	num2 int
}

type NumbersInterface interface {
	Sum() int
	Multiply() int
	Division() float64
	Substract() int
}

func (n Numbers) Sum() int {
	return n.num1 + n.num2
}

func (n Numbers) Multiply() int {
	return n.num1 * n.num2
}

func (n Numbers) Division() float64 {
	return float64(n.num1) / float64(n.num1)
}

func (n Numbers) Substract() int {
	return n.num1 - n.num2
}

func main() {
	var i NumbersInterface
	numbers := Numbers{9, 8}
	i = numbers
	fmt.Printf("Sum: %d\n", i.Sum())
	fmt.Printf("Sum: %d\n", i.Multiply())
	fmt.Printf("Sum: %f\n", i.Division())
	fmt.Printf("Sum: %d\n", i.Substract())
}
```
# Goroutines
```
func main() {
	ch := make(chan string)
	
	go say("Hello Go!", ch)
	fmt.Println(<-ch)
}

func say(greet string, ch chan string) {
	ch <- greet
}
```

```
func main() {
	ch := make(chan int)

	go say("Hello Go!", ch)

	for x := range ch {
		fmt.Println(x)
	}
}

func say(greet string, ch chan int) {
	for i := 0; i <= 5; i++ {
		ch <- i
	}

	close(ch)
}
```
# Goto
```
func main() {
    learnGoTo()
}

func learnGoTo() {
    fmt.Println("a")
    goto FINISH
    fmt.Println("b")
FINISH:
    fmt.Println("c")
}

// a
// c
```
