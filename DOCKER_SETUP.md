# 🐳 DOCKER_SETUP.md

## Cómo levantar el proyecto con Docker

Sigue estos pasos para construir y ejecutar el contenedor Docker de este proyecto en Windows 10:

### 1. Requisitos previos
- Tener Docker Desktop instalado y en ejecución.

### 2. Construir la imagen Docker
Abre una terminal en la carpeta del proyecto y ejecuta:

```bash
docker build -t bugonautas-app .
```
Esto creará una imagen llamada `bugonautas-app`.

### 3. Levantar el contenedor
Ejecuta el siguiente comando:

```bash
docker run -d -p 8080:8080 --name bugonautas bugonautas-app
```
- `-d`: Ejecuta el contenedor en segundo plano.
- `-p 8080:8080`: Mapea el puerto 8080 del contenedor al 8080 de tu máquina.
- `--name bugonautas`: Asigna el nombre `bugonautas` al contenedor.

### 4. Acceder a la aplicación
Abre tu navegador y visita:
```
http://localhost:8080
```

---

## Preguntas frecuentes

### ¿Qué pasa si corro la imagen sin flags?
Si ejecutas:
```bash
docker run bugonautas-app
```
El contenedor se ejecuta en primer plano y no puedes usar la misma terminal para otros comandos hasta que lo detengas.

### ¿Cómo detengo el contenedor?
Si le diste un nombre:
```bash
docker stop bugonautas
```
Si no le diste nombre, usa el ID del contenedor que puedes ver con:
```bash
docker ps
```
