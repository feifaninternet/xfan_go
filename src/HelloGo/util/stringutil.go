package util

/*
@Time : 2019/1/22 17:58
@Author : xfan
@Desc: Reverse String Util
*/
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i],r[j] = r[j],r[i]
	}
	return string(r)
}

func MiddleChar(s string) string {
	i := len(s)
	if i % 2 != 0 {
		return string(s[(i + 1) / 2 - 1])
	}
	return string("")
}

