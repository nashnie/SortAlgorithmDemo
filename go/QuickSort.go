package SortAlgorithmDemo

func quicksort(data []int, start int, end int) {
    if start > end {
        return
    }
    var temp = data[start]
    var i = start + 1
    var j = end
    for true {
        for data[j] > temp {
            j--
        }
        for data[i] < temp && i < j {
            i++
        }
        if i >= j {
            break
        }
        swap(data, i, j)
        i++
        j--
    }
    if j != start {
        swap(data, start, j)
    }
    quicksort(data, j + 1, end)
    quicksort(data, start, j - 1)
}

func swap(data []int, a int, b int){
    var temp int = data[a];
    data[a] = data[b];
    data[b] = temp;
}