# Pruebas funcionales de servicios web

Ejecutar el servidor:

```bash
go run ./cmd/api
```

En otra terminal ejecutar:

```bash
curl http://localhost:8080/api/salud
curl http://localhost:8080/api/productos
curl http://localhost:8080/api/productos/PROD-001
curl -X POST http://localhost:8080/api/clientes -H "Content-Type: application/json" -d '{"id":"CLI-002","nombre":"Ana Perez","correo":"ana@example.com","direccion":"Quito","telefono":"0988888888"}'
curl -X POST http://localhost:8080/api/carrito/items -H "Content-Type: application/json" -d '{"clienteId":"CLI-002","productoId":"PROD-001","cantidad":1}'
curl http://localhost:8080/api/carrito/CLI-002
curl -X POST http://localhost:8080/api/pedidos -H "Content-Type: application/json" -d '{"clienteId":"CLI-002"}'
curl -X POST http://localhost:8080/api/pagos -H "Content-Type: application/json" -d '{"pedidoId":"PED-002","metodo":"transferencia"}'
curl http://localhost:8080/api/reportes/ventas
curl http://localhost:8080/api/reportes/inventario
```

Nota: el identificador exacto del pedido depende de la secuencia generada durante la ejecucion. Si cambia, copiar el id obtenido en la respuesta de `POST /api/pedidos`.
