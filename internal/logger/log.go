package logger

// para mejorar este paquete debería ser una clase, entrar a La librería Debug para darnos una idea.

import (
	"encoding/json"

	"github.com/gookit/color"
)

func Debug(args ...interface{}) {
	//puede recibir multiples parametros,
	//args es un arreglo por lo que podemos iterar
	// para volver a pasar parámetres como si fueran , args...
	color.Debug.Println(args...)
}

func Info(args ...interface{}) {
	color.Info.Println(args...)
}

func Warn(args ...interface{}) {
	color.Warn.Println(args...)
}

func Success(args ...interface{}) {
	color.Success.Println(args...)
}

func Log(args ...interface{}) {
	color.Light.Println(args...)
}

func Primary(args ...interface{}) {
	color.Primary.Println(args...)

}

func Error(args ...interface{}) {
	color.Error.Println(args...)
}

func DebugJson(jsonType interface{}) {

	prettyJSON, err := json.MarshalIndent(jsonType, "", "  ")

	if err != nil {
		color.Error.Println("Failed to generate json", err)
	}

	color.Debug.Printf("%s\n", string(prettyJSON))
}
