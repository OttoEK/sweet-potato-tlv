package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	s := "11AB398765UJ1A050200N23"
	m, err := MapearTlv([]byte(s))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(m)
	}

}

func MapearTlv(value []byte) (map[interface{}]string, error) {
	if value == nil {
		return nil, errors.New("sin parametros")
	}
	pos := 0
	m := make(map[interface{}]string)
	for {
		size := value[pos : pos+2]
		pos = pos + 2
		nsize := string(size[0]) + string(size[1])
		largo, err := strconv.Atoi(nsize)
		if err != nil {
			return nil, err
		}
		valor := string(value[pos : pos+largo])
		pos = pos + largo
		tipo := string(value[pos])
		pos = pos + 1
		nCampo, err := strconv.Atoi(string(value[pos : pos+2]))

		if tipo == "N" {
			m[nCampo] = valor
		} else if tipo == "A" {
			m[string(value[pos:pos+2])] = valor
		}

		pos = pos + 2
		if pos == len(value) {
			break
		}
	}

	return m, nil

}
