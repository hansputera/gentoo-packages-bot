package utils

import (
	"bufio"
	"io"
	"log"
	"os"
)

type envtmp struct {
	key  string
	val  string
	stop bool
}

func LoadEnv(envPath string) {
	tmp := &envtmp{}

	fl, err := os.Open(envPath)
	if err != nil {
		log.Fatalln(err)
	}

	reader := bufio.NewReader(fl)
	for {
		if c, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				os.Setenv(tmp.key, tmp.val)
				tmp = nil
				break
			} else {
				log.Fatal(err)
			}
		} else {
			str := string(c)

			if str == "\n" {
				os.Setenv(tmp.key, tmp.val)

				tmp.key = ""
				tmp.val = ""
				tmp.stop = false
			} else if tmp.stop {
				continue
			}

			str = StandardizeSpaces(str)
			if str == "#" {
				tmp.stop = true
				continue
			} else if len(tmp.val) > 0 {
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
}
