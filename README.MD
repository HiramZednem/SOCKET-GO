# Proyecto de Sockets Go con Angular

Este proyecto combina un servidor en Go con el framework Socket.IO y una aplicación de frontend en Angular para la comunicación en tiempo real a través de WebSockets.

## Instrucciones de ejecución

Para ejecutar el proyecto, sigue los siguientes pasos:

1. En la carpeta `./sockets-go`, asegúrate de tener instalado Go en tu sistema. Luego, ejecuta el siguiente comando para obtener las dependencias necesarias:

   ```shell
   go get github.com/googollee/go-socket.io

2. Ejecuta este comando para que sirva la app de angular
   ```shell
   go run './main.go'
   
## Notas
Todo lo que trabajes, tienes que crearle un .dist para que el servido de go lo pueda levantar, el problema actual esta en que no me puedo conectar al socket, si encuentra una mejor manera de resolver esta problematica adelante
