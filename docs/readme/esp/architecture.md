# Resumen de Arquitectura

## Estructura del Código

El código base está organizado en varios directorios y archivos para mantener una Arquitectura Limpia. Los componentes principales son:
- `cmd`: Contiene los puntos de entrada de la aplicación.
- `models`: Define los modelos de datos utilizados en la aplicación.
- `packages`: Incluye diferentes paquetes que encapsulan funcionalidades específicas:
    - `cache`: Gestiona los mecanismos de caché.
    - `database`: Maneja las operaciones de la base de datos.
    - `service`: Contiene los servicios de lógica de negocio.
    - `server`: Implementación del servidor API y métricas.
- `storage`: Almacena archivos de datos como archivos JSON.

La Arquitectura Limpia se utiliza porque es la forma más idiomática de organizar código en Go, promueve una separación clara de responsabilidades y facilita las pruebas y el mantenimiento del código. El código debe organizarse de la manera más simple posible (no somos desarrolladores de Java! :) )

## Estructura del Proyecto

Básicamente, hay 3 componentes ahora:
1. Servidor API, que se encarga de manejar las solicitudes y respuestas HTTP.
2. Caché de Redis, que se usa para almacenar datos accedidos frecuentemente.
3. Base de datos, que se usa para almacenar datos persistentes.

Sin embargo, según la tarea, la base de datos debe reemplazarse con un almacenamiento de archivos JSON. Aun así, se mantiene la funcionalidad de la interfaz para futuros cambios.
