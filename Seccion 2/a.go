package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	channel := make(chan int)

	fmt.Println("Funcion principal. Esta funcion imprimira las primeras 8 vocales del abecedario")
	for i := 0; i < 8; i++ {
		if i == 0 {
			fmt.Println("a")
		}
		if i == 1 {
			fmt.Println("b")
		}
		if i == 2 {
			fmt.Println("c")
		}
		if i == 3 {
			fmt.Println("d")
			fmt.Println("Reenviando el programa a la primera rutina, esta rutina imprimira los primeros 5 numeros primos")
			time.Sleep(1000 * time.Millisecond)
			wg.Add(2)
			go rutina1(wg, channel)
			go rutina2(wg, channel)
			wg.Wait()
		}
		if i == 4 {
			fmt.Println("e")
		}
		if i == 5 {
			fmt.Println("f")
		}
		if i == 6 {
			fmt.Println("g")
		}
		if i == 7 {
			fmt.Println("h")
		}
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Println("Ha terminado la ejecucion de las 2 gorutinas y de la funcion principal. Hasta pronto.")

}

func rutina1(wg *sync.WaitGroup, channel chan int) {
	var ultNum int
	for i := 0; i < 5; i++ {
		if i == 0 {
			fmt.Println(i + 2)
		}
		if i == 1 {
			fmt.Println(i + 2)
		}
		if i == 2 {
			fmt.Println(i + 3)
		}
		if i == 3 {
			fmt.Println(i + 4)
		}
		if i == 4 {
			ultNum = i + 5
			fmt.Println(ultNum)
		}
		time.Sleep(1000 * time.Millisecond)
	}
	channel <- ultNum
	fmt.Println("Terminando la ejecucion de la primera rutina, redirigiendo a la segunda corutina. Esta imprimira los 5 primeros multiplos del ultimo numero primo mostrado anteriormente")
	wg.Done()
}

func rutina2(wg *sync.WaitGroup, channel chan int) {
	lnum := <-channel
	var mult int
	for i := 0; i < 5; i++ {
		if i == 0 {
			mult = lnum * 1
			fmt.Println(mult)
		}
		if i == 1 {
			mult = lnum * 2
			fmt.Println(mult)
		}
		if i == 2 {
			mult = lnum * 3
			fmt.Println(mult)
		}
		if i == 3 {
			mult = lnum * 4
			fmt.Println(mult)
		}
		if i == 4 {
			mult = lnum * 5
			fmt.Println(mult)
		}
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Println("Terminando la ejecucion de la segunda corutina, redirigiendo a la funcion principal para terminar la impresion de las letras.")

	wg.Done()
}
