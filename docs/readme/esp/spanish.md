# Detalle de Ítem de MELI

- [Más detalles sobre el proyecto](architecture.md)
- [Prompts utilizados](prompts.md)

## Descripción

Este proyecto es una API simple para obtener detalles de ítems de Mercado Libre. Utiliza Go y Redis como cache.

## Instalación

Para instalar el proyecto, necesitas tener Go instalado en la máquina. Luego, podes ejecutar los siguientes comandos:

```sh
git clone https://github.com/juy13/MELI-API
cd meli-item-detail
go mod tidy
go build -o meli-api
```

## Ejecución

Para ejecutar el proyecto, necesitas tener Redis corriendo en la máquina. Se puede ejecutar usando Docker:

```sh
docker compose -f docker-compose-services.yml up -d
```

Asegúrate de que Redis esté en ejecución antes de iniciar la API. Debes redefinir la variable de entorno `REDIS_PASSWORD` en el archivo `deployment/docker-compose-services.yml` con tu contraseña. Agregar la ruta del dispositivo `/var/data/redis` a las `driver_opts` del volumen en el archivo `docker-compose-services.yml` permite que los datos persistan entre reinicios del contenedor.

## Configuración

El proyecto utiliza `config.yaml` para la configuración. Podés modificar  para cambiar el puerto de la API y otros ajustes. El archivo de configuración se encuentra en `configs/config.yaml`.

## Base de Datos

El proyecto utiliza un archivo JSON como base de datos. La ruta al archivo de la base de datos se especifica en el archivo `config.yaml` en la sección `database`.

## Métricas

El proyecto utiliza Prometheus para la recolección de métricas. El servidor de métricas se ejecuta en el puerto 9090 por defecto. Podés cambiar este puerto en el archivo `config.yaml` en la sección `metrics`.

Adicionalmente, existen métricas personalizadas:

  - http\_request\_duration\_seconds
  - http\_requests\_total

Se pueden encontrar visitando el endpoint `/metrics` del servidor de la API.

## Pruebas

El proyecto incluye pruebas unitarias para los endpoints de la API. Puedes ejecutar las pruebas usando el siguiente comando:

```sh
go test ./...
```

## Despliegue

Se proporcionan archivos de docker compose en la carpeta `deployment`. Podés usarlos para desplegar la aplicación localmente o en un entorno de producción.

`data-storage` es un volumen que se utiliza para almacenar datos. Se monta en `/data/storage` dentro del contenedor. Externamente, toma una carpeta con archivos JSON como base de datos.

Para ejecutar la aplicación usando docker compose, navega al directorio `deployment` y ejecuta:

```sh
docker compose -f docker-compose.yml up --build
```

## Swagger

El proyecto incluye documentación Swagger para los endpoints de la API. Puedes acceder a ella visitando el endpoint `/swagger/index.html` del servidor de la API.