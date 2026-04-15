# proxmox-monitor

> Herramienta personal de monitoreo de infraestructura para Proxmox VE con alertas en tiempo real vía Telegram.

Desarrollada para uso en homelab y entornos de trabajo. Consulta la API de Proxmox periódicamente, guarda el historial de estados en SQLite y envía alertas a Telegram cuando un nodo, VM o contenedor LXC cambia de estado.

---

## ¿Qué hace?

- Consulta el estado del nodo Proxmox y todas sus VMs y contenedores LXC
- Detecta cambios de estado (online → offline y viceversa)
- Envía alertas a Telegram solo cuando hay un cambio, sin spam
- Guarda un historial de 72 horas en SQLite
- Al arrancar, carga el último estado conocido desde la base de datos para evitar falsos positivos

---

## Stack

| Tecnología | Uso |
|---|---|
| Go | Lenguaje principal |
| SQLite (`modernc.org/sqlite`) | Historial de estados (sin dependencias externas) |
| Proxmox API REST | Fuente de datos |
| Telegram Bot API | Canal de alertas |

---

## Estructura del proyecto

```
proxmox-monitor/
├── cmd/
│   └── main.go              # Punto de entrada
├── internal/
│   ├── models/
│   │   ├── base.go          # Struct base compartido
│   │   ├── node.go          # Modelo de nodo Proxmox
│   │   └── vm.go            # Modelo de VM y LXC
│   ├── database/
│   │   ├── db.go            # Conexión a SQLite
│   │   └── schema.go        # Creación de tablas
│   ├── proxmox/
│   │   ├── client.go        # Cliente HTTP con SSL autofirmado
│   │   ├── nodes.go         # Consulta estado del nodo
│   │   └── vms.go           # Consulta VMs y LXC
│   ├── telegram/
│   │   └── notify.go        # Envío de alertas
│   └── scheduler/
│       └── scheduler.go     # Loop de chequeo periódico
├── .env.example
├── go.mod
└── go.sum
```

---

## Configuración

Copia `.env.example` a `.env` y completa los valores:

```env
# Proxmox
PROXMOX_HOST=https://192.168.1.100:8006
PROXMOX_TOKEN=usuario@pve!token-id=token-secret

# Telegram
TELEGRAM_BOT_TOKEN=tu-bot-token
TELEGRAM_CHAT_ID=tu-chat-id

# Base de datos
DB_PATH=./data/proxmox.db
```

### Crear el API Token en Proxmox

1. Ir a **Datacenter → Permissions → API Tokens**
2. Crear un token con rol `PVEAuditor` (solo lectura)
3. Copiar el valor del token al `.env`

---

## Ejemplo de alertas

```
🔴 ALERTA — VM caída
Nombre: ubuntu-web
ID: 101
Nodo: pve
Hora: 2025-04-14 03:42:11
```

```
🟢 RECUPERADO
Nombre: ubuntu-web
ID: 101
Tiempo caído: 5 minutos
```

---

## Estado del proyecto

🚧 En desarrollo activo.

- [x] Modelos de datos
- [x] Conexión a SQLite
- [x] Cliente HTTP Proxmox
- [x] Consulta de nodos
- [x] Consulta de VMs y LXC
- [x] Alertas Telegram
- [ ] Scheduler
- [ ] Main.go

---

## Autor

Creado por [Oriver Segura Vargas](https://github.com/oriversegura). Libre para usar, modificar y compartir bajo licencia MIT.

---

## Licencia

[MIT](./LICENSE) © 2026 Oriver Segura Vargas
