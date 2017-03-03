package SortAlgorithmDemo

func mergesort(data []int) []int {
    if len(data) <= 1 {
        return data
    }
    var mid int = len(data) / 2
    var left []int = []int{}
    var right []int = []int{}

    for i := 0; i < mid; i++ {
        left = append(left, data[i])
    } 
    for j := mid; j < len(data); j++ {
        right = append(right, data[j])
    }
	
    left = mergesort(left)
    right = mergesort(right)
	
	fmt.Println("mergesort........................................")
	fmt.Println("mergesort left")
	fmt.Println(left)
	fmt.Println("mergesort right")
	fmt.Println(right)	
    return merge(left, right)
}

func merge(left []int, right []int) []int {
    var temp []int = []int{}
    for len(left) > 0 && len(right) > 0 {
        if left[0] <= right[0] {
            temp = append(temp, left[0])
			left = append(left[:0], left[1:]...)
			fmt.Println("merge left")
			fmt.Println(left)
			fmt.Println("temp")
			fmt.Println(temp)
        } else {
            temp = append(temp, right[0])
			right = append(right[:0], right[1:]...)
			fmt.Println("merge right")
			fmt.Println(right)
			fmt.Println("temp")
			fmt.Println(temp)
        }
    }
    if len(left) > 0 {
        for i := 0; i < len(left); i++ {
            temp = append(temp, left[i])
        }
    }

    if len(right) > 0 {
        for i := 0; i < len(right); i++ {
            temp = append(temp, right[i])
        }
    }
	fmt.Println("temp finaly")
	fmt.Println(temp)
    return temp
}
