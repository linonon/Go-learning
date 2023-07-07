package codewar

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func rangeExtraction(list []int) string {
	if len(list) == 0 {
		return ""
	}
	if len(list) == 1 {
		return strconv.Itoa(list[0])
	}
	if len(list) == 2 {
		return strconv.Itoa(list[0]) + "," + strconv.Itoa(list[1])
	}

	res := []string{}
	tmp := []int{list[0]}
	for i, v := range list[1:] {
		// continue
		if v-tmp[len(tmp)-1] == 1 {
			tmp = append(tmp, v)
			if i+1 == len(list[1:]) {
				if len(tmp) >= 3 {
					res = append(res, strconv.Itoa(tmp[0])+"-"+strconv.Itoa(tmp[len(tmp)-1]))
				} else {
					for _, vt := range tmp {
						res = append(res, strconv.Itoa(vt))
					}
				}
			}
			continue
		}
		// not continue
		if len(tmp) <= 2 {
			for _, vt := range tmp {
				res = append(res, strconv.Itoa(vt))
			}
			tmp = tmp[:1]
			tmp[0] = v
			continue
		}
		res = append(res, strconv.Itoa(tmp[0])+"-"+strconv.Itoa(tmp[len(tmp)-1]))
		tmp = tmp[:1]
		tmp[0] = v
		if i+1 == len(list[1:]) {
			res = append(res, strconv.Itoa(v))
		}
	}

	return strings.Join(res, ",")
}

func TestRangeExtraction(t *testing.T) {

	res := rangeExtraction([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20, 22})
	// res := rangeExtraction([]int{-6, -3})
	fmt.Println(res)
	fmt.Println(res == "-6,-3-1,3-5,7-11,14,15,17-20")

}
