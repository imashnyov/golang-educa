# Formatting strings
```
name := "John"
var age int = 30
value := 1000.476
a := true

fmt.Printf("My name is %s.\nMy age is %d.\nI have %f$.\n%t", name, age, value, a)
```
# Loop
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

for i := 1; i <= 5; i++ {
	if i%2 != 0 {
		break
	}
	fmt.Println(i)
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
# Switch
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
#### Switch fallthrough 
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
# Array
```
array := [3]string{"Denis", "Kate", "John"}
for i := 0; i < len(array); i++ {
	fmt.Println(names[i])
}

// array := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
```
# Slice
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