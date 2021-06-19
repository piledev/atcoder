package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var N string
	fmt.Scanf("%s", &N)
	anss := ""
	ansi := 0

	if len(N)%2 == 1 {
		anss = strings.Repeat("9", len(N)/2)
		ansi, _ = strconv.Atoi(anss)

	} else {
		n := strings.Split(N, "")
		befs := strings.Join(n[:len(N)/2], "")
		afts := strings.Join(n[len(N)/2:], "")
		befi, _ := strconv.Atoi(befs)
		afti, _ := strconv.Atoi(afts)

		if befi <= afti {
			ansi = befi
		} else {
			if len(befs) == len(strconv.Itoa(befi-1)) {
				ansi = befi - 1
			} else {
				anss = strings.Repeat("9", len(N)/2-1)
				ansi, _ = strconv.Atoi(anss)
			}
		}
	}

	fmt.Println(ansi)
}
