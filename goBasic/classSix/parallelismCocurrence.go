/*
Concurrencia:
- el acto de iniciar ejecutar y completar 2 o mas
tareas en periodos de tiempos intarcalados, no necesariamente
al mismo tiempo.

Paralelismo:
- 2 o mas tareas se ejecutan exactamente al mismo tiempo.

GO RUTINES
-> implementacion de concurrencia
-> no son HILOS!
-> son gestionadas por el GO Runtime y su scheduler, no por el SO.
-> su ejecucion no bloqueara la continuacion de la funcion que la invoco
porque corre de forma concurrente.

*/

package main

import (
	"fmt"
	"time"
)

func process(i int) {
	fmt.Println(i, " Starting...")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, " Stoping...")
}
func processTwo(i int, ch chan int) {
	fmt.Println(i, " Starting...")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(i, " Stoping...")
	ch <- i
}

func main() {

	now := time.Now()
	/*
		fmt.Println(" Starting...")
		process(1)
		fmt.Println(" Stoping...")
		fmt.Printf("the function program took %d ms\n", time.Now().Sub(now).Milliseconds())

		for i := 0; i <= 10; i++ {
			process(i)
		}
		fmt.Printf("the bucle program took %d ms\n", time.Now().Sub(now).Milliseconds())

		for i := 0; i <= 10; i++ {
			go process(i)
			fmt.Printf("Go rutine %d\n", i)
		}
		time.Sleep(2000 * time.Millisecond)
		fmt.Printf("the go rutine program took %d ms\n", time.Now().Sub(now).Milliseconds())
	*/
	// Chanels
	// medio para enviar y recibir variables/valores
	chanel := make(chan int)
	for i := 0; i <= 10; i++ {
		go processTwo(i, chanel)
	}
	//para esperar todas las go rutines
	for i := 0; i <= 10; i++ {
		reading := <-chanel
		fmt.Println("Reading chanel ", reading)
	}
	fmt.Printf("the go rutine with chanels program took %d ms\n", time.Now().Sub(now).Milliseconds())

}
