# This is a practice in doing ddd for my masters degree at La Salle.

# 🏢 Coworking App - DDD & Hexagonal Architecture

Este proyecto es una práctica de **Domain-Driven Design (DDD)** para mi máster en **La Salle**, utilizando la arquitectura **Ports and Adapters (Hexagonal Architecture)**.

## 🚀 Descripción

Esta aplicación de coworking permite gestionar **hotdesks, oficinas y salas de reuniones**, aplicando principios de **DDD** y un diseño modular y escalable.

## 🛠️ Tecnologías

- **Golang** - Lenguaje de programación principal.
- **Go Fiber** - Framework web rápido y minimalista.
- **Testify** - Framework de testing.
- **Go Validator** - Validación de DTOs.

## 📂 Estructura del Proyecto

Para el dia que este readme fue actualizado.

├── cmd
│ └── main.go
├── go.mod
├── go.sum
├── internal
│ ├── adapters
│ │ ├── http
│ │ │ ├── handlers
│ │ │ │ ├── hotdesk_handler.go
│ │ │ │ ├── meeting_room_handler.go
│ │ │ │ ├── office_handler.go
│ │ │ │ └── utils.go
│ │ │ ├── http_errors
│ │ │ │ └── errors.go
│ │ │ ├── models
│ │ │ │ ├── hotdesk_dto.go
│ │ │ │ ├── meeting_room_dto.go
│ │ │ │ └── office_dto.go
│ │ │ ├── routes.go
│ │ │ └── server.go
│ │ └── storage
│ │ ├── hotdesk_repository.go
│ │ ├── meeting_room_repository.go
│ │ └── office_repository.go
│ ├── app
│ │ ├── domain
│ │ │ ├── domain_errors
│ │ │ │ └── domain_errors.go
│ │ │ ├── entities
│ │ │ │ ├── hotdesk.go
│ │ │ │ ├── hotdesk_test.go
│ │ │ │ ├── meeting_room.go
│ │ │ │ ├── meeting_room_test.go
│ │ │ │ ├── office.go
│ │ │ │ └── office_test.go
│ │ │ └── vo
│ │ │ ├── hotdesk.go
│ │ │ ├── metting_room.go
│ │ │ ├── office.go
│ │ │ └── status.go
│ │ └── usecases
│ │ ├── command.go
│ │ └── commands
│ │ ├── register_hotdesk.go
│ │ ├── register_meeting_room.go
│ │ └── register_office.go
│ └── ports
│ ├── http_port.go
│ └── storage_port.go
└── README.md

## 📦 Instalación

### 1️⃣ Clonar el repositorio

```bash
git clone https://github.com/tu-usuario/coworking-app.git
cd coworking-app

```

### 2️⃣ Ejecutar la aplicación

```bash
go mod tidy
go run cmd/main.go
```
