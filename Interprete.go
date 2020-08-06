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
	inicio := "Esta es la consola de comandos, ingrese X para finalizar."
	fmt.Println(inicio)

	///Leyendo desde consola.
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese texto: ")
	comando, _ := reader.ReadString('\n')
	fmt.Println(comando)

	if comando == "x" {
		finalizar = 1
	} else {
		if comando != "" {
			LineaComando(comando)
		}
	}

	for finalizar != 1 {
		fmt.Println("Insertar Comando: ")
		reader := bufio.NewReader(os.Stdin)
		comando, _ := reader.ReadString('\n')
		if comando == "x" {
			finalizar = 1
		} else {
			if comando != "" {
				LineaComando(comando)
			}
		}

	}
}

//LineaComando : obtiene el texto de entrada.
func LineaComando(comando string) {
	var commandArray []string
	commandArray = strings.Split(comando, " ")
	fmt.Println(commandArray[1])
	EjecutarComando(commandArray) //Ejecutamos el comando.
}

//EjecutarComando : metodo para ejecutar el codigo de entrada.
func EjecutarComando(commandArray []string) {
	data := strings.ToLower(commandArray[2])
	if data == "file" {
		fmt.Println("ARCHIVO")

	} else if data == "crear" {

		fmt.Println("CREAR ARCHIVO")

	} else {
		fmt.Println("Otro Comando")
	}
}
