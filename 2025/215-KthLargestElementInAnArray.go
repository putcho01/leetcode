package main

import "container/heap"

// --- container/heap 用の型定義 ---
// int32のスライスをヒープとして扱えるようにします
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// ヒープ
func (h IntHeap) Less(i, j int) bool {
	// 大小比較をした結果を返します。降順に
	return h[i] > h[j]
}

// ２つの要素を入れ替える関数です
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 新しい要素を末尾に追加します
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// まず先頭の要素を保存します。次に、末尾のデータを先頭に代入し、末尾の要素を削除します。
func (h *IntHeap) Pop() interface{} {
	old := *h

	n := len(old)

	x := old[n-1]

	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	// 1. 配列をヒープとして初期化
	h := IntHeap(nums)
	heap.Init(&h) // O(N) で整列されます

	for i := 0; i < k; i++ {
		if i == k-1 {
			return heap.Pop(&h).(int)
		}
		heap.Pop(&h)
	}
	return 0
}
