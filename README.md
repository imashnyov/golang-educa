## Loops

func loop1() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Hello %d\n", i)
	}
}

# ∞
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

# while loop
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

##