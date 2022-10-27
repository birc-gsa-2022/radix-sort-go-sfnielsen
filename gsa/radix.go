package gsa

import (
	"fmt"
	"sort"
	"strings"
)

// Sort the string x using a count sort
func CountSort(x string) string {
	var sb strings.Builder
	counts := make(map[byte]int)
	keys := []int{}
	//count occourences
	for i := range x {
		if _, ok := counts[x[i]]; !ok {
			keys = append(keys, int(x[i]))
		}
		counts[x[i]]++
	}

	//IS THIS OKAY ??? k log k. Should still be constant or order 'len x' in normal cases?
	sort.Ints(keys)

	//create ordered string
	for _, k := range keys {
		v := counts[byte(k)]
		for i := 0; i < v; i++ {
			sb.WriteByte(byte(k))
		}
	}
	return sb.String()
}

// Sort the indices in idx according to the letters in x
// using a bucket sort
func BucketSort(x string, idx []int) []int {
	xs := CountSort(x)

	buckets := make(map[rune]int) //create first bucket beggining at 0
	//create buckets with accumulated values
	for i, v := range xs {
		if i == 0 {
			buckets[v] = 0
			continue
		}
		if v != rune(xs[i-1]) {
			buckets[v] = i
		}
	}
	idx_s := make([]int, len(idx))
	for _, v := range idx {
		key := rune(x[v])
		new_i := buckets[key]
		buckets[key]++
		idx_s[new_i] = v
	}
	return idx_s
}

// Compute the suffix array for x using a least-significant digits
// first radix sort
func LsdRadixSort(x string) []int {
	var sb strings.Builder

	//add sentinel
	if x[len(x)-1] != '$' {
		sb.WriteString(x)
		sb.WriteByte('$')
	}
	x = sb.String()
	sb.Reset()

	//create initial sorting
	idx := make([]int, len(x))
	for i := range idx {
		idx[i] = i
	}

	for col := len(x) - 1; col >= 0; col-- {
		sb.Reset()
		for suf := range x {
			sb.WriteByte(byte(x[(suf+col)%len(x)]))
		}
		idx = BucketSort(sb.String(), idx)
	}
	fmt.Println(idx)

	return idx
}

// Compute the suffix array for x using a most-significant digits
// first radix sort
func MsdRadixSort(x string) []int {
	return []int{}
}
