package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "11AB398765UJ1A050200N23"
	m, err := mapearTlv([]byte(s))
	if(err!=nil){
		fmt.Println(err)
	} else {
		fmt.Println(m)
	}

}

func mapearTlv(value []byte) (m map[interface{}]string, err error) {
	pos := 0
	m = make(map[interface{}]string)
	for {
		size := value[pos : pos+2]
		pos = pos + 2
		nsize := string(size[0]) + string(size[1])
		largo, err := strconv.Atoi(nsize)
		if err != nil {
			return m, err
		}
		valor := string(value[pos : pos+largo])
		pos = pos + largo
		tipo := string(value[pos]) //2+largo+3
		pos = pos + 1
		nCampo, err := strconv.Atoi(string(value[pos : pos+2]))
		pos = pos + 2

		if tipo == "N" {
			m[nCampo] = valor
		} else if tipo == "A" {
			m[strconv.Itoa(nCampo)] = valor
		}
		if pos == len(value) {
			break
		}
	}
	return m, err

}
