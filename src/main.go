package main

import (
	"fmt"
	"github.com/goriiin/bashnya-nomework0/pkg"
)

func main() {
	var (
		a, b, n, m int
	)
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		return
	}

	_, err = fmt.Scan(&n, &m)
	if err != nil {
		return
	}

	if a > b {
		pkg.Swap(&a, &b)
		pkg.Swap(&n, &m)
	}

	mas1 := pkg.FillArray(a, n)
	mas2 := pkg.FillArray(b, m)

	//fmt.Println(mas1, mas2)
	totMin := pkg.MaxNum(mas1[0], mas2[0])
	totMax := pkg.MinNum(mas1[1], mas2[1])
	if totMin > totMax {
		fmt.Println(-1)
	} else {
		fmt.Println(totMin, totMax)
	}
}
