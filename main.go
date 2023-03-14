package main

import (
	"encoding/json"
	"net/http"
	_ "net/http/pprof"
	"prof/pkg"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

func main() {
	go simpleMap()

	go simpleStruct()

	go easyJson()

	http.HandleFunc("/", hiHandler)
	http.ListenAndServe(":8080", nil)
}

func simpleMap() {
	for {
		res1 := map[string]any{}
		_ = json.Unmarshal([]byte(pkg.Example), &res1)
		//fmt.Print(res1)
	}
}

func simpleStruct() {
	for {
		res2 := pkg.Message{}
		_ = json.Unmarshal([]byte(pkg.Example), &res2)
		//fmt.Print(res2)
	}
}

func easyJson() func() {
	for {
		res3 := pkg.Message{}
		_ = res3.UnmarshalJSON([]byte(pkg.Example))
		//fmt.Print(res3)
	}
}
