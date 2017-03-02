package SortAlgorithmDemo

func insertionsort(data []int) {
    for i := 1; i < len(data); i++ {
        var temp int = data[i]
		var j int
        for j = i - 1; j >= 0 && data[j] > temp; j-- {
            data[j + 1] = data[j]
        }
        data[j + 1] = temp
		fmt.Println(j)
    }
}

func insertionsortwithbinarysearch(data []int) {
    for i := 1; i < len(data); i++ {
        var temp int = data[i]
		var j int
        var binarysearchIndex = binarysearch(data, 0, i, i)
		fmt.Println(binarysearchIndex)
		if binarysearchIndex >= 0 {
			for j = i - 1; j >= binarysearchIndex; j-- {
            	data[j + 1] = data[j]
        	}
        	data[j + 1] = temp
		}
    }
}

func binarysearch(data []int, start int, end int, key int) int {
	if start > end {
		return -1
	}
    var mid int = start + (end - start) / 2
	if mid == key {
		return mid
	}
	if data[mid] > data[key] {
		if mid - 1 < 0 {
			return 0
		}
		if data[mid - 1] < data[key] {
			return mid
		}
        return binarysearch(data, start, mid - 1, key)
	}
	if data[mid] < data[key] {
		if data[mid + 1] > data[key] {
			return mid + 1
		}
        return binarysearch(data, mid + 1, end, key)
	}
	return mid
}
