package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

type Users struct {
	Name string
	Surname string
	Age int
	EntityFunctions
}

type EntityFunctions interface {
	all() *EntityEntries
}


//type EntityCol
//type EntityRows interface{}

// map[column][]rows
type EntityEntries map[string][]string
type EntityRegisters map[string]reflect.Type

type EntityOptions struct {

}

const (
	T_ENT_FUNC_CALLER			=	0x2E
	T_ENT_FUNC_PARAM_SEP		=	0x2C
)

const (
	T_DATA_PTR			=	0x60
)

var entName string
var entData []byte

func main() {
	//q := "<entity>.<entity_functions>(<entity_field>, <params>)..."


	reader, _ := os.Open("./data.txt")
	data, _ := ioutil.ReadAll(reader)

	q := " Users.all()"
	for i, v := range q {
		switch v {
		case T_ENT_FUNC_CALLER:
			entName = strings.TrimSpace(q[:i])


		}
	}

	var betweenData []byte
	startIndex := 0
	endIndex := 0
	for i, v := range data {
		if v == T_DATA_PTR {
			// sadece belirli bir tablonun verilerini alacağımız için,
			// herhangi bir tablonun başlangıç işaretinin indexi
			// olup olmadığını kontrol ediyorum. eğer işaret indexi yoksa (== 0)
			// ve okunan verinin bulunan indeksine(i) kadar olanki veri,
			// entity adına eşitse(eg. data[:i] == "Users") başlangıç indeksini
			// bulunan noktaya işaretliyorum.
			if startIndex == 0 {
				beforeStartPart := data[:i]
				if string(beforeStartPart) == entName {
					startIndex = i
				}

			} else {
				// eğer başlangıç indeksi var ise artık
				// bitiş indeksini arayabilirim demek.
				// çünkü entity verileri, başlangıç ve bitiş işaretleri
				// arasında depolanacak. Bu durumda aradaki datayı da alabilirim.
				endIndex = i
				betweenData = data[startIndex + 1:endIndex]
				break
			}
		}
	}

	betweenData = bytes.ReplaceAll(betweenData, []byte("-{"), []byte("{"))

	//var result map[string]string
	d := normalizeData(betweenData)

	for _, value := range d {
		fmt.Println(value)
	}

	//d := ""
	//s := strings.Split(string(betweenData), ",")
	//for _, value := range s {
	//
	//	d = strings.Replace(value, "{", "", 1)
	//	d = strings.Replace(d, "}", "", 1)
	//	fmt.Println(d)
	//}



}


func normalizeData(data []byte) []string {
	var d []string
	s := strings.Split(string(data), ",")
	d = make([]string, len(s))
	for i, value := range s {

		if value != "" {
			d[i] = strings.TrimSpace(strings.Replace(value, "{", "", 1))
			d[i] = strings.TrimSpace(strings.Replace(d[i], "}", "", 1))
		}
	}

	return d
}

func delete_empty (s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}