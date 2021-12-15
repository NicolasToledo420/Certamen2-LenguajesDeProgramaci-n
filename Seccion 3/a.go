package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	//Captura información ingresada por terminal
	ingreso := flag.Int("tiempo", 2, "Ingrese Tiempo de espera")
	flag.Parse()

	//Creacion de las listas de los clientes, para tanto su ID como su timepo de llegada.
	var idfila []int
	var tiempofila []int

	//Creacion de la gorutina y canales que genera los clientes en la fila que deben realizar.
	canalCliente := make(chan int)
	canalControl := make(chan int)
	go genRandom(canalCliente, canalControl, *ingreso)

	//Se ingresa en las listas, los clientes que fueron generados por la gorutina.
	for i := 0; i < 10; i++ {
		canalControl <- 1
		idCliente := <-canalCliente
		idfila = append(idfila, idCliente)
		tiempo := <-canalCliente
		tiempofila = append(tiempofila, tiempo)
	}

	//Se le envia a la gorutina un valor que le indica que debe terminar su proceso (este es el canal de control).
	canalControl <- 0

	//Imprimir las listas de ID y tiempo, para tener un control del funcionamiento de la fila.
	fmt.Println(idfila)
	fmt.Println(tiempofila)
	wg := &sync.WaitGroup{}

	//Creacion de las cajas con sus canales correspondientes.
	canalIn := make(chan int)
	canalOut := make(chan int)
	go cajas(wg, canalIn, canalOut, "1")

	canalIn2 := make(chan int)
	canalOut2 := make(chan int)
	go cajas(wg, canalIn2, canalOut2, "2")

	//Variables que definen los tiempos en que terminan las gortinas.
	tiempoCaja1 := 0
	tiempoCaja2 := 0

	//Este es el for principal donde se realiza la simulacion de las cajas.
	Condition := true
	for Condition {
		//Variable que tendra el valor que la gorutina envia.
		var valor int
		//Si se tiene que la fila de los tiempos de llegada vacia se terminara la ejecucion.
		if len(tiempofila) == 0 {
			//Se le envia el valor 100, que le da a entender a la gorutina su fin de ejecucion.
			canalIn <- 100
			//Retorno del mismo valor para consistencia.
			valor = <-canalOut
			//Se le envia el valor 100, que le da a entender a la gorutina su fin de ejecucion.
			canalIn2 <- 100
			//Retorno del mismo valor para consistencia.
			valor = <-canalOut2
			println("Fin Ejecucion")
			Condition = false
		} else {
			//Se toma el primer valor de la fila tanto su ID como su timepo de llegada
			cliente := idfila[0]
			tiempo := tiempofila[0]
			//Si el tiempo de llegada es mayor o igual al tiempo fin de caja 1, esta atendera al cliente
			if tiempo >= tiempoCaja1 {
				canalIn <- cliente
				valor = <-canalOut
				tiempoCaja1 = tiempo + valor
				fmt.Println(tiempoCaja1)
				//Si el tiempo de llegada es mayor o igual al timepo fin de caja 2, esta atendera al cliente
			} else if tiempo >= tiempoCaja2 {
				canalIn2 <- cliente
				valor = <-canalOut2
				tiempoCaja2 = tiempo + valor
				fmt.Println(tiempoCaja2)
				//Si el tiempo de llegada es menor que cualquier tiempo de anterior
			} else if tiempoCaja1 > tiempoCaja2 {
				canalIn2 <- cliente
				valor = <-canalOut2
				tiempoCaja2 = tiempoCaja2 + valor
				fmt.Println(tiempoCaja2)
			} else {
				canalIn <- cliente
				valor = <-canalOut
				tiempoCaja1 = tiempoCaja1 + valor
				fmt.Println(tiempoCaja1)
			}
			idfila = RemoveIndex(idfila, 0)
			tiempofila = RemoveIndex(tiempofila, 0)
			println(valor)
			if valor == 100 {
				println("Funca")
				Condition = false
			}
		}
	}

}

func cajas(wg *sync.WaitGroup, in chan int, out chan int, id string) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	Condition := true
	for Condition {
		cliente := <-in
		if cliente == 100 {
			Condition = false
			out <- 100
		} else {
			fmt.Println("Caja" + id)
			fmt.Println("Atendiendo al cliente: ", cliente)
			//Generar tiempo random de la gente que va a estar en la caja
			tiempo := r1.Intn(4) + 1
			time.Sleep(time.Duration(tiempo * int(time.Second)))
			out <- tiempo
		}
	}
}

func genRandom(out chan int, in chan int, parametro int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	tiempo := 0
	incremento := parametro
	Condition := true
	for Condition {
		control := <-in
		if control == 0 {
			Condition = false
		} else {
			out <- r1.Intn(50) + 1
			tiempo = tiempo + incremento
			out <- tiempo
		}
	}
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
