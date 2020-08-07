package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	Interpretar()
}

//Interpretar : Funcion principal para interpretar el texto de entrada.
func Interpretar() {

	finalizar := 0
	inicio := "MANEJO E IMPLEMENTACION DE ARCHIVOS.\n"
	fmt.Println(inicio)

	///Leyendo desde consola.
	for finalizar != 1 {
		fmt.Println("INGRESAR COMANDO>>>: ")
		reader := bufio.NewReader(os.Stdin)
		comando, _ := reader.ReadString('\n')

		if comando == "x\n" {
			fmt.Println("FINALIZAR")
			finalizar = 1
			return
		} else if comando != "" {
			DividirArreglo(comando)

		}

	}
}

//DividirArreglo : Dividir el arreglo de entrada por los espacios en blanco.
func DividirArreglo(comando string) {
	var arreglo []byte = []byte(comando)
	fmt.Println("tam ", len(arreglo))

	bandera := 0

	for i := 0; i <= len(arreglo)-1; i++ {

		if string(arreglo[i]) == "\"" {
			bandera = 1
		}

		if string(arreglo[i]) == " " && bandera == 1 {
			arreglo[i] = 95
		}

	}

	newArreglo := string(arreglo)

	var commandArray []string
	commandArray = strings.Split(newArreglo, " ")
	ArregloDividido(commandArray) //Ejecutamos el comando.

}

//ArregloDividido : Arreglo dividido por medio de los espacios en blanco.
func ArregloDividido(arreglo []string) {

	tamArreglo := len(arreglo)
	fmt.Println("tamaño del arreglo:", tamArreglo)
	aux0 := strings.ToLower(arreglo[0])
	aux1 := strings.ToLower(arreglo[1])
	aux2 := strings.ToLower(arreglo[2])

	fmt.Println(aux0)
	fmt.Println(aux1)
	fmt.Println(aux2)

	if tamArreglo == 4 {
		aux3 := strings.ToLower(arreglo[3])
		fmt.Println(aux3)

	}

	if aux1 == "file" {
		fmt.Println("ARCHIVO")

	} else if aux1 == "crear" {

		fmt.Println("CREAR ARCHIVO")

	} else {
		fmt.Println("Otro Comando")
	}
}
