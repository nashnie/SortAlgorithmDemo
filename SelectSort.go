package SortAlgorithmDemo

import "fmt"

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

func swap(data []int, a int, b int){
     var temp int = data[a];
    data[a] = data[b];
    data[b] = temp;
}

func main() {
	fmt.Println("selectsort")
	s := []int{1, 3, 6, 9, 2, 3, 0}
	selectsort(s);
	fmt.Println(s)
}
