package main

import "fmt"

func main() {
	st := NewSegmentTree()
	for i := 0; i < 10; i++ {
		st.Add(0, 0, 9, i, i, 1)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(st.Query(0, 0, 9, i, i))
	}
	fmt.Println(st.Query(0, 0, 9, 0, 9))

	for i := 0; i < 10; i++ {
		st.Add(0, 0, 9, i, i, 1)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(st.Query(0, 0, 9, i, i))
	}
	fmt.Println(st.Query(0, 0, 9, 0, 9))

	st.Add(0, 0, 9, 2, 2, 100)
	fmt.Println(st.Query(0, 0, 9, 0, 9))
	fmt.Printf("max: %d, min: %d\n", st.QueryMax(0, 0, 9, 0, 9), st.QueryMin(0, 0, 9, 0, 9))

	st.Set(0, 0, 9, 2, 2, 0)
	fmt.Println(st.Query(0, 0, 9, 0, 9))

	fmt.Printf("max: %d, min: %d\n", st.QueryMax(0, 0, 9, 0, 9), st.QueryMin(0, 0, 9, 0, 9))

	st.Set(0, 0, 9, 0, 9, 0)
	fmt.Println(st.Query(0, 0, 9, 0, 9))
	fmt.Printf("max: %d, min: %d\n", st.QueryMax(0, 0, 9, 0, 9), st.QueryMin(0, 0, 9, 0, 9))
	/*
		a := []int{1, 2, 3, -1, -2, -3, 11, -22, 233}
		tree := NewSegmentTree(len(a))

		b := make([]int, len(a))
		for i := range a {
			b[i] = i
		}
		sort.Slice(b, func(i, j int) bool {
			return a[b[i]] < a[b[j]]
		})

		for _, x := range a {
			idx := sort.Search(len(b), func(j int) bool {
				return a[b[j]] >= x
			})
			trie.query()
		} */
}
