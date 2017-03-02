package SortAlgorithmDemo

import "fmt"

func bubblesort(data []int) {
    for i := len(data) - 1; i > 0 ; i-- {
        for j := 0; j < i; j++ {
            if data[j] > data[j + 1] {
                swap(data, j, j + 1)
            }
        }
    }
}

func bubblesortwithflag(data []int) {
	var withoutswap bool = true
    for i := len(data) - 1; i > 0 ; i-- {
        for j := 0; j < i; j++ {
            if data[j] > data[j + 1] {
                swap(data, j, j + 1)
                withoutswap = false
            }
        }
        if withoutswap {
			fmt.Println(i)
            break
        }
    }
}

func bubblecocktailsort(data []int) {
	var withoutswap bool = true
    var m, n int
    for i := len(data) - 1; i > 0 ; i-- {
		if i % 2 == 0 {
			for j := n; j < len(data) - 1 - m; j++ {
				if data[j] > data[j + 1] {
					swap(data, j, j + 1)
					withoutswap = false
				}
			}
			if withoutswap {
				fmt.Println(i)
				break
			}
			m++
		}else{
			for j := len(data) - 1 - m; j > n; j-- {
				if data[j] < data[j - 1] {
					swap(data, j, j - 1)
					withoutswap = false
				}
			}
			if withoutswap {
				fmt.Println(i)
				break
			}
			n++
		}
	}
}

func swap(data []int, a int, b int){
     var temp int = data[a]
    data[a] = data[b]
    data[b] = temp
}