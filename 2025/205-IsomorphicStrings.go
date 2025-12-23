package main

// https://leetcode.com/problems/isomorphic-strings/?envType=study-plan-v2&envId=top-interview-150
func isIsomorphic(s string, t string) bool {
	// mappingST: s の文字インデックス -> t の対応するバイト値
	// mappingTS: t の文字インデックス -> s の対応するバイト値
	// インデックスは 'a' からのオフセット（小文字の前提）を使っている
	mappingST := make(map[int]int)
	mappingTS := make(map[int]int)

	for i := 0; i < len(s); i++ {
		// s[i], t[i] を 'a' からのオフセットで表現（0..25 の想定）
		idxST := int(s[i] - 'a')
		idxTS := int(t[i] - 'a')
		// s -> t の既存マッピングを検査
		if val, ok := mappingST[idxST]; ok {
			// 既にマップされているなら、今回の t[i] と一致するか確認
			if val != int(t[i]) {
				return false
			}
		} else {
			mappingST[idxST] = int(t[i])
		}
		// t -> s の既存マッピングを検査（単射を保証）
		if val, ok := mappingTS[idxTS]; ok {
			if val != int(s[i]) {
				return false
			}
		} else {
			mappingTS[idxTS] = int(s[i])
		}
	}
	// 全ての位置で一貫した双方向マッピングが保たれている -> 真
	return true
}
