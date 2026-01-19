package main

import "slices"

func addBinary(a string, b string) string {
	// a と b の末尾（最下位ビット）を指すインデックス
	i, j := len(a)-1, len(b)-1
	// 繰り上がり（carry）
	carry := 0
	// 結果を格納するスライス（byteで保持）
	// 下位ビットから追加していく
	res := []byte{}

	// a または b にまだ桁が残っている、または繰り上がりがある間ループ
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0') // '0' or '1' → 数値に変換
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0') // '0' or '1' → 数値に変換
			j--
		}

		// sum % 2 が現在の桁の値（0 or 1）
		// byte にして '0' を足し、文字として追加
		res = append(res, byte(sum%2)+'0')
		carry = sum / 2
	}

	slices.Reverse(res)
	return string(res)
}
