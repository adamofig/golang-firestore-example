## Aplicación de go retomada 

actualización, estaba intentando que fuera mi backend para conectar con dialogflow pero ha sido algo complicado. 

## Instrucciones de ejecución, 

Tengo que hacerlo desde 0 pero parece que en general es

go mod install : para instalar los modulos,

Compila go y los archivos necesarios para funcionar. 
    go run main.go * 

## Ejecutar en windows
Para ejecutar podemos usar 

    go run main.go

Pasa que como Go Run borra los binarios y crea una nueva carpeta siempre pide activar el firewall y como es  mejor no desactivarlo
La mejor manera es usar go build y luego correr el ejecutable, -o para darle un nombre 
    
    go build -o app.exe
    app.exe 

## Entender que se hizo con ls go modules para iniciar

Para crear un modulo, es posible utilizar el siguiente comando, toda la aplicación será como un modulo

    go mod init dataclouder 
    

Se creará el archivo go.mod 

aunque la mejor estructura, es repo-server/username/packagename
github.com/adamofig/dataclouder sin embargo no parece ser muy importante 

## Instalar otros paquetes

Se sigue la estructura go get repo/packages
    go get cloud.google.com/go/firestore

Cuando ya existe nuestro go mod se descargan los paquetes aquí. 

Alguna de las formas de probar si estan todas las dependencias es 

    go test

    o 

    go mod download

Automaticamente descarga las dependencias y las pone en go.mod

ademas go.sum guarda los hash del paquete para asegurarse que nadie los ha modificado. 




Aqui me dió una idea de inicar con una estructura 
https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/

# Estructura de archivos Backend

***/internal*** Código interno de la aplicación, clases y paquetes

- ***internal/firestore*** Servicios para obtener información de firestore

- ***internal/entities*** Clases generales entidade de firestore

- ***internal/dialogflow*** Clases relacionadas a Dialog 


***/docs*** Documentación adicional

***/resources*** Archivos estaticos que requiera la aplicación 

- ***resources/algom*** 

## Documentar la api 
Se utilza swagger, aún estoy definiendo las mejores practicas, espero hacer un video de esto, pero siempre es muy costoso,

Primero es importante idetificar bien que librerías se estan usaudno. 
hay 3 importaciones adicionales en main 

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "dataclouder/docs" // docs is generated by Swag CLI, you have to import it.

Puede audar mucho la un ejemplo de la documentación. 

https://github.com/swaggo/swag/tree/master/example/celler
doc
// @Param id path int true "Account ID"
// @Param q query string false "name search by q" Format(email)
// @Param account body model.AddAccount true "Add account"

Es complicado la documentación pero es lo mejor que tengo para probar el back 

El objetivo es hacer un servicio  universal, que registre todo en firebase. 
pudiera validar si quisiera, que si lo haré para los registros de ingles. 
y otros pueden ser libre pero tengo que documentar todo
