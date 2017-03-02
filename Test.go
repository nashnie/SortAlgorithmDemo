package main

import (
	"fmt"
)

func selectsort(data []int) {
    for i := 0; i < len(data) - 1; i++ {
        var min int = i;
        var temp int = data[i];
        for j := i + 1; j < len(data); j++ {
            if (data[j] < temp){
                min = j;
                temp = data[j];
            }
        }
        if min != i {
            swap(data, min, i);
        }
    }
}

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
            break;
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

func main() {
	s := []int{1, 4, 9, 4, 2, 8, 7}
	bubblecocktailsort(s)
	fmt.Println(s)
}
