
# Users Service

## Descripción
El servicio **Users Service** gestiona la interacción entre los usuarios dentro de la plataforma, permitiendo la funcionalidad de seguir a otros usuarios. Cada vez que un usuario sigue a otro, este servicio notifica al servicio **Timeline Service** para actualizar la línea de tiempo del usuario.

## Características principales
- Creación y gestión de usuarios.
- Seguimiento y des-seguimiento de usuarios.
- Integración con **Timeline Service** para actualizar la línea de tiempo en tiempo real.
- Exposición de APIs RESTful.
- Logs detallados para monitoreo y depuración.

## Tecnologías utilizadas
- **Golang**: Lenguaje de desarrollo principal.
- **PostgreSQL / MongoDB**: Base de datos para almacenamiento de usuarios y relaciones.
- **Docker**: Contenedorización para despliegue.
- **REST API**: Interfaz de comunicación con otros servicios.
- **Logrus**: Manejo de logs estructurados.

## Instalación y configuración
### Prerrequisitos
- Tener instalado **Go** (versión 1.24 o superior).
- Contar con **Docker** y **Docker Compose** (para pruebas locales).
- Configurar las variables de entorno adecuadas.

### Clonar el repositorio
```bash
  git clone https://github.com/uala-challenge/users-service.git
  cd users-service
```

### Configuración
El servicio utiliza un archivo de configuración en formato YAML. Antes de ejecutar, asegúrate de configurar correctamente `config.yaml`:
```yaml
api:
  port: 8081
  timeline_service_url: "http://timeline-service:8082"

logging:
  level: "info"
```

### Ejecutar en entorno local
#### Usando Go directamente
```bash
  go run main.go
```

#### Usando Docker
```bash
  docker-compose up --build
```

## API REST
El servicio expone los siguientes endpoints:

### **Seguir a un usuario**
```
PATCH /follow/{user_id}
```
**Payload enviado:**

```json
{
  "followed_user_id": "67890"
}
```
#### **Descripción:**
Permite a un usuario seguir a otro usuario.

#### **Respuesta:**
```json
{
  "message": "User followed successfully"
}
```

#### **Integración con Timeline Service:**
Cada vez que un usuario sigue a otro, el servicio realiza una petición REST a **Timeline Service** para actualizar la línea de tiempo del usuario:
```
PATCH /timeline/{user_id}
```
**Payload enviado:**
```json
{
  "follower_id": "user:123"
}

```

## Testing
Para ejecutar las pruebas unitarias:
```bash
  go test ./...
```
