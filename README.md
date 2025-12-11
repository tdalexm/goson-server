# GOSON-Server

Un servidor JSON RESTful escrito en Go, inspirado en [json-server](https://github.com/typicode/json-server).

!> [!IMPORTANT]
> v0.1.0 Esta versión en desarrollo activo y puede tener cambios sustanciales en futuras releases.

## Inicio rápido

### Instalación

```bash
# Descargar ejecutable (Windows/Linux)
# Ver Releases: https://github.com/tdalexm/goson-server/releases

# Compilado desde codigo fuente
git clone https://github.com/tdalexm/goson-server.git
cd goson-server
go build -o goson-server
```


### Uso Básico

1. Crear un archivo `db.json`:

```json
{
  "pokemon": [
    { "id": "1", "name": "Bulbasaur", "type": "Grass/Poison", "level": 5 },
    { "id": "2", "name": "Charmander", "type": "Fire", "level": 5 },
    { "id": "3", "name": "Squirtle", "type": "Water", "level": 5 }
  ],
  "digimon": [
    { "id": "1", "name": "Agumon", "stage": "Rookie", "attribute": "Vaccine" },
    { "id": "2", "name": "Gabumon", "stage": "Rookie", "attribute": "Data" },
    { "id": "3", "name": "Patamon", "stage": "Rookie", "attribute": "Data" }
  ]
}
```

!> [!NOTE]
> El archivo .zip contiene un db.json de ejemplo listo para hacer pruebas.

2. Iniciar el servidor:

```bash
./goson-server 
./goson-server --db=db.json --port=8080
```

3. ¡Listo! La API REST esta disponible en `http://localhost:8080`

## API Reference (v0.1.0)

### Endpoints

| Método HTTP | Ruta | Descripción | Parámetros Query | Códigos de Respuesta |
|-------------|------|-------------|------------------|----------------------|
| `GET` | `/:resource` | Listar todos los elementos | `field`, `value`, `contains` | 200 OK, 204 No Results, 404 Not Found |
| `GET` | `/:resource/:id` | Obtener un elemento específico por ID | - | 200 OK, 404 Not Found |
| `POST` | `/:resource` | Crear un nuevo elemento | - | 201 Created, 400 Bad Request, 404 Not Found |
| `POST` | `/:resource/:id` | Sobreescribir un elemento | - | 200 OK, 400 Bad Request, 404 Not Found |
| `PATCH` | `/:resource/:id` | Actualizar parcialmente un elemento | - | 200 OK, 400 Bad Request, 404 Not Found |
| `DELETE` | `/:resource/:id` | Eliminar un elemento | - | 200 OK, 404 Not Found |

### Query Params `GET /:resource`

| Parámetro | Tipo | Descripción | Ejemplo |
|-----------|------|-------------|---------|
| `field` | string | Campo por el cual filtrar | `?field=name` |
| `value` | string | Valor exacto para filtrar. | `?field=name&value=Charmander` |
| `contains` | string | Texto a buscar dentro del campo (búsqueda parcial) | `?field=type&contains=grass` |

!> [!NOTE]
> Las busquedas **no** son case-sensitive.

### Notas Importantes:
1. **Los parámetros `field` y `value`/`contains` funcionan juntos** - si especificas `field`, debes proporcionar `value` o `contains`
2. **PATCH solo actualiza campos especificados**, no reemplaza el objeto completo
3. **No se permite la modificación del campo ID**

## Configuración CLI

```bash
./goson-server --help

Uso:
  --db string       Ruta relativa al archivo JSON de base de datos (default "db.json")
  --port string     Puerto para el servidor (default "8080")
  --help            Mostrar ayuda

Ejemplos:
  ./goson-server --db=./data/mi-db.json --port=3030
  ./goson-server --port=8000
```

## Dependencias

- Go 1.21 o superior
- Gin Framework (`v.1.9`)

## Roadmap

- Paginación y mejora de las respuestas.
- Persistencia de cambios en el archivo `db.json`.
- Concurrencia.

## Feedback

Tu feedback es crucial para seguir mejorando.

- [Reportar un Bug](https://github.com/tdalexm/goson-server/issues)
- [Solicitar una Feature](https://github.com/tdalexm/goson-server/issues)

## Agradecimientos
[typicode/json-server](https://github.com/typicode/json-server) 
