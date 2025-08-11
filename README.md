# System Monitor

Este proyecto es una aplicación para monitorear métricas básicas del sistema, como el uso de CPU, memoria y disco.

Está construido como un monorepo con dos componentes principales:

1.  **Backend (Go):** Una API simple escrita en Go que recolecta y expone la información del sistema.
2.  **Frontend (Angular):** una interfaz web moderna construida con Angular que consume la API del backend y muestra los datos en tiempo real.

---

## Estructura del Proyecto

```
/
├── system-monitor-backend/   # Código fuente de la API en Go
└── system-monitor-frontend/  # Código fuente de la aplicación web en Angular
```

---

## Cómo Empezar

### Prerrequisitos

Asegúrate de tener instalado lo siguiente:

*   [Go](https://golang.org/dl/)
*   [Node.js](https://nodejs.org/) y npm
*   [Angular CLI](https://angular.io/cli) (`npm install -g @angular/cli`)

### Pasos para Ejecutar

Debes ejecutar el backend y el frontend en dos terminales separadas.

**1. Ejecutar el Backend:**

```bash
# Navega al directorio del backend
cd system-monitor-backend

# Ejecuta la aplicación de Go
go run main.go
```

El servidor de la API se iniciará, por lo general en `http://localhost:8080`.

**2. Ejecutar el Frontend:**

```bash
# En una nueva terminal, navega al directorio del frontend
cd system-monitor-frontend

# Instala las dependencias (solo la primera vez)
npm install

# Inicia el servidor de desarrollo de Angular
npm start
```

La aplicación web estará disponible en `http://localhost:4200`.

---

## Contribuciones

Las contribuciones son bienvenidas. Por favor, abre un "issue" para discutir cambios importantes o envía un "pull request" con tus mejoras.
