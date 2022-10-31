package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type envtmp struct {
	key string
	val string
}

func LoadEnv(envPath string) {
	tmp := envtmp{}

	fl, err := os.Open(envPath)
	if err != nil {
		log.Fatalln(err)
	}

	reader := bufio.NewReader(fl)
	for {
		if c, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			str := StandardizeSpaces(string(c))
			if len(tmp.val) > 0 {
				if tmp.val == "=" {
					tmp.val = ""
				}
				tmp.val += str
			} else {
				if str == "=" {
					tmp.val += str
				} else {
					tmp.key += str
				}
			}
		}
	}

	fmt.Println(tmp)
}
