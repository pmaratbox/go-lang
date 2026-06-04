package main

import "fmt"

func addStrings(a, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := 0
	res := []byte{}
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		res = append(res, byte('0'+sum%10))
		carry = sum / 10
	}
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return string(res)
}

func main() {
	fmt.Println(addStrings("999999999999", "1"))
}
