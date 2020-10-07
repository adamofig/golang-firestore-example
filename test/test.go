package main

import (
	"fmt"
	"log"
	"math"
	"reflect"
)

type Persona struct {
	Nombre    string
	Edad      int
	EsSolteri bool
}

func (p Persona) Suscribirse() {
	println("Hola se inicia una suscripción", p.Nombre, p.Edad, p.EsSolteri)
}

func main() {
	dato := 129
	_ = dato
	var hola string = "hola ${dato}"
	log.Println(hola)
	a := fmt.Sprintln("Ok iniciando", "bien", dato)
	b := "hola " + string(dato)
	c := fmt.Sprintf("%s @ el dato es %d", a, dato)

	log.Println(a, b, c)
	//googleCloud.TestStorage()
	n := 24.0 / 12.0

	if n == math.Trunc(n) {
		log.Println("es un tentro completo ")
	} else {
		log.Println("No es completo")
	}

	log.Println("Mudular", reflect.TypeOf(n))
}
