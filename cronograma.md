openapi: 3.0.3
info:
  title: Sistema de Gestion de E-commerce - API Final POO
  version: 1.0.0
  description: Servicios web JSON para clientes, productos, carrito, pedidos, pagos y reportes.
servers:
  - url: http://localhost:8080
paths:
  /api/salud:
    get:
      summary: Verifica el estado del servicio
      responses:
        '200': { description: Servicio disponible }
  /api/productos:
    get:
      summary: Lista productos
      responses:
        '200': { description: Catalogo en formato JSON }
    post:
      summary: Crea producto
      responses:
        '201': { description: Producto creado }
  /api/productos/{id}:
    get:
      summary: Consulta producto por identificador
      parameters:
        - in: path
          name: id
          required: true
          schema: { type: string }
      responses:
        '200': { description: Producto encontrado }
  /api/clientes:
    post:
      summary: Registra un cliente
      responses:
        '201': { description: Cliente registrado }
  /api/carrito/items:
    post:
      summary: Agrega item al carrito
      responses:
        '200': { description: Carrito actualizado }
  /api/carrito/{clienteId}:
    get:
      summary: Consulta carrito de un cliente
      parameters:
        - in: path
          name: clienteId
          required: true
          schema: { type: string }
      responses:
        '200': { description: Carrito consultado }
  /api/pedidos:
    post:
      summary: Crea pedido desde carrito
      responses:
        '201': { description: Pedido creado }
  /api/pedidos/{id}:
    get:
      summary: Consulta un pedido
      parameters:
        - in: path
          name: id
          required: true
          schema: { type: string }
      responses:
        '200': { description: Pedido encontrado }
  /api/pagos:
    post:
      summary: Procesa pago de pedido
      responses:
        '200': { description: Pago procesado }
  /api/reportes/ventas:
    get:
      summary: Reporte de ventas
      responses:
        '200': { description: Indicadores de ventas }
  /api/reportes/inventario:
    get:
      summary: Reporte de inventario
      responses:
        '200': { description: Productos y bajo stock }
