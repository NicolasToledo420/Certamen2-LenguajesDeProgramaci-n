package main

import (
	"fmt"
	"time"
)

func main() {
	var control string

	canalControl := make(chan string)
	go corrutina(canalControl)

	canalControl <- "Control Corrutina"
	control = <-canalControl
	fmt.Println(control)
	abc()

	canalControl <- "Control Corrutina"
	control = <-canalControl
	fmt.Println(control)
	abc()

	fmt.Println("Fin ejecucion")
}

func corrutina(CanalControl chan string) {
	var control string

	control = <-CanalControl
	fmt.Println(control)
	_123()

	CanalControl <- "Control Main"

	control = <-CanalControl
	fmt.Println(control)
	_123()

	CanalControl <- "Fin corrutina, Control Main"

}

func abc() {
	lista := []string{"a", "b", "c", "d"}
	for i := 0; i < 4; i++ {
		fmt.Println(lista[i])
		time.Sleep(time.Duration(500 * int(time.Millisecond)))
	}
}
func _123() {
	for i := 0; i < 4; i++ {
		fmt.Println(i)
		time.Sleep(time.Duration(500 * int(time.Millisecond)))
	}
}
