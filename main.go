package main

import (
	"encoding/json"
	"fmt"
	"prof/pkg"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			res1 := map[string]any{}
			_ = json.Unmarshal([]byte(pkg.Example), &res1)
			fmt.Print(res1)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			res2 := pkg.Message{}
			_ = json.Unmarshal([]byte(pkg.Example), &res2)
			fmt.Print(res2)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 1000; i++ {
			res3 := pkg.Message{}
			_ = res3.UnmarshalJSON([]byte(pkg.Example))
			fmt.Print(res3)
		}
		wg.Done()
	}()

	wg.Wait()
}
