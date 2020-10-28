package main

import (
	"fmt"
)

type Array []int

func main() {
	fmt.Println("Hello World")

	f := Array([]int{1, 2, 4, 8, 16, 32, 64, 128})
	j := Array([]int{1, 10, 15, 20, 25, 30})

	fmt.Println(f.merge(j,true))

	//r := f.binarySearch(8)
	//r := f.linearSearch(8)
	//i,j := f.max()
	//fmt.Printf("\nfound at Index: %d, %v\n", j,i)
	//f.reverse()
	//fmt.Println()
	//r :=f.isSorted()

}

func (a Array) linearSearch(key int) int {

	for i, v := range a {

		if key == v {
			return i
		}
	}
	return -1
}

/**
Binary Search assumes sorted array
*/
func (a Array) binarySearch(key int) int {
	low := 0
	high := len(a) - 1
	mid := low + high

	for low < high {
		mid = (low + high) / 2
		if key == a[mid] {
			return mid
		} else if key < a[mid] {
			high = mid - 1
		} else if key > a[mid] {
			low = mid + 1
		}
	}
	return -1
}

func (a Array) max() (max int, index int) {
	max = a[0]
	index = 0

	for i, v := range a {
		if max < v {
			max = v
			index = i
		}
	}
	return max, index
}

func (a Array) avg() int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum / (len(a))
}

func (a Array) sum() int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func (a Array) reverse() {

	j := len(a) - 1
	for i, v := range a {
		if i < j {
			tmp := a[j]
			a[i] = tmp
			a[j] = v
			j--
		}
	}
}

func (a Array) leftShift(rotate bool) {
	j := len(a) - 1
	tmp := a[0]
	for i, _ := range a {
		if i < j {
			a[i] = a[i+1]
		} else {
			if rotate {
				a[j] = tmp
			} else {
				a[j] = 0
			}
			return
		}
	}
}

func (a Array) isSorted() bool {
	j := len(a) - 1
	for i, _ := range a {
		if i < j {
			if a[i] > a[i+1] {
				return false
			}
		}
	}
	return true
}

func (a Array) rightShift(rotate bool) {

	i := len(a) - 1
	tmp := a[i]
	for i > 0 {
		a[i] = a[i-1]
		i--
	}
	if rotate {
		a[0] = tmp
	} else {
		a[0] = 0
	}

	fmt.Println(a)
}


func (a Array) merge(b Array, removeDuplicates bool ) Array {
	i := 0; j := 0 ; k := 0
	results := make(Array, len(a)+len(b))
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			results[k] = a[i]
			i++
		} else {
			if removeDuplicates && a[i] == b[j] {
				results[k] = b[j]
				i++
			}
			results[k] = b[j]
			j++
		}
		k++
	}
	for i < len(a) {
		results[k] = a[i]
		i++
		k++
	}

	for j < len(b) {
		results[k] = b[j]
		j++
		k++
	}

	return results
}
