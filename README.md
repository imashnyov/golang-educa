# Loop
```
func loop1() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Hello %d\n", i)
	}
}
```
#### ∞
func loop2() {
	for {
		fmt.Printf("Loop")
	}
}

func loop3() {
	for i := 1; i <= 5; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}

func loop4() {
	for i := 1; i <= 5; i++ {
		if i%2 != 0 {
			break
		}
		fmt.Println(i)
	}
}

#### while loop
func loop6() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}

func loop7() {
	nums := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}

# Switch
func main() {
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
}

#### Switch fallthrough 
func main() {
	number := 10

	switch {
	case number > 1:
		fmt.Println("Number is greater than 1")
		fallthrough
	case number > 11:
		fmt.Println("Number < 11")
	}
}