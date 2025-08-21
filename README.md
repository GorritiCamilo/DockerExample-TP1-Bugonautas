# 🐹 Hola Mundo con Go

## 📋 Descripción
Servidor web minimalista escrito en Go. Perfecto para aprender Docker con el lenguaje Go.

## 💡 Warning
Las siguientes `warnings` son para crear el `dockerfile`
 * Se recomienda empezar desde la siguiente imagen: `golang:1.19-alpine`
 * No olvidar copiar los files `go.mod` y `main.go` dentro del docker
 * No olvidar buildear el proyecto durante el building del `docker image`
````bash
# Ejecutar el comando
go build -o main .
````
 * Se tiene que exponer el `port` 8080. Podes hacerlo usando
````bash
# Exponer  el puerto 5000
EXPOSE 8080
````
 * Es fundamental que se ejecute el comando `CMD` al runnear la `docker image`
````bash
# Ejecutar el proyecto
./main
````

## 🚀 Actividades
Deben hacer el `DOCKER_SETUP.md` teniendo las siguientes consideraciones
 * ¿Qué pasa si corremos la `docker image` sin asignar ninguna flag a `docker run`? ¿Podemos usar la misma terminal para correr otros comandos?
 * El proyecto usa el usa el port `8080`. Intentar hacer `docker run` con y sin los parametros correspondientes. ¿Qué ocurre en cada caso?
 * Ejecutar `docker stop <container>`. ¿Qué pasa si al hacer `docker run` no le asigno un nombre al contenedor? ¿Qué debo poner en `<container>`para poder hacer `docker stop <container>`?
 * Si corro el contenedor en segundo plano, no veo información de la dirección IP que necesito para usar mi proyecto. Documentar qué se debe poner en el navegador
