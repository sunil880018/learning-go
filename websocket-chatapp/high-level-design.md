# Chat Application (Go + WebSockets)

## Used pure web socket, You can use Socket.io for better performance

## 🔹 High-Level Flow

### 1. HTTP Server Setup (`main.go`)

- Serves templates (`index.html`, `chat.html`).
- Serves static files (CSS, JS).
- Handles `/room` route → upgrades HTTP connection to a WebSocket connection.

### 2. Room Management (`room.go`)

- Each room keeps track of connected clients.
- Messages sent in the room are broadcast to **all connected clients**.

### 3. Client Handling (`client.go`)

- Each client represents a single user connected via WebSocket.
- `read()` → listens for incoming messages from this user.
- `write()` → sends outgoing messages to this user.

### https://second-viper-21a.notion.site/Gin-Package-16cf1fff91fb801cb536f1edc5b474c7
