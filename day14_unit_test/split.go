package split

import (
	"strings"
)

//Split 将 s 按照 sep 进行分割，返回一个字符串的切片

//返回子串str在字符串s中第一次出现的位置。
//如果找不到则返回-1；如果str为空，则返回0

func Split(s, sep string) (ret []string) {

	ret = make([]string,0,strings.Count(s,sep)+1)   //优化 提前分配内存
	idx := strings.Index(s, sep)
	for idx > -1 {
		ret = append(ret, s[:idx])

		s = s[idx+len(sep):]
		idx = strings.Index(s, sep)
	}

	ret = append(ret, s)
	return

}