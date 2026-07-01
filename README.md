# Sistema de Gestion de E-commerce - Proyecto Final POO

**Autor:** Matias Castro Guerra  
**Asignatura:** Programacion Orientada a Objetos  
**Docente guia:** Palacios Morocho Milton Ricardo  
**Repositorio sugerido:** https://github.com/MatixProx/sistema-gestion-ecommerce-poo.git

## 1. Descripcion del proyecto

Este proyecto implementa un Sistema de Gestion de E-commerce como proyecto integrador de Programacion Orientada a Objetos. La solucion permite administrar productos, clientes, carrito de compras, pedidos, pagos, inventario y reportes mediante servicios web que serializan la informacion en formato JSON.

## 2. Objetivo

Desarrollar una version funcional del sistema de gestion seleccionado, integrando conceptos de POO, estructuras de datos, encapsulacion, interfaces, manejo de errores y servicios web para evidenciar la evolucion del proyecto desde la planeacion inicial hasta la entrega final.

## 3. Tecnologias

- Go 1.23
- Libreria estandar `net/http`
- Serializacion JSON con `encoding/json`
- Almacenamiento en memoria para fines academicos
- Git y GitHub para versionamiento

## 4. Conceptos de POO aplicados

| Concepto | Aplicacion en el proyecto |
|---|---|
| Abstraccion | Se modelan entidades del negocio como Producto, Cliente, Carrito, Pedido y Pago. |
| Encapsulacion | Producto protege stock, precio y datos internos mediante metodos como `DescontarStock`. |
| Interfaces | `MetodoPago` permite usar tarjeta, transferencia o pago contra entrega sin cambiar la logica del pedido. |
| Polimorfismo | Diferentes metodos de pago responden al mismo comportamiento `Procesar`. |
| Responsabilidad unica | Modelo, repositorio, servicio y API HTTP se separan en archivos distintos. |
| Manejo de errores | Se bloquean cantidades invalidas, stock insuficiente, carritos vacios y metodos de pago no soportados. |

## 5. Servicios web implementados

| # | Metodo | Ruta | Funcionalidad |
|---|---|---|---|
| 1 | GET | `/api/salud` | Verifica disponibilidad del servicio. |
| 2 | GET | `/api/productos` | Lista productos del catalogo. |
| 3 | POST | `/api/productos` | Crea un producto. |
| 4 | GET | `/api/productos/{id}` | Consulta producto por identificador. |
| 5 | POST | `/api/clientes` | Registra un cliente. |
| 6 | POST | `/api/carrito/items` | Agrega un producto al carrito. |
| 7 | GET | `/api/carrito/{clienteId}` | Consulta carrito del cliente. |
| 8 | POST | `/api/pedidos` | Crea pedido desde el carrito. |
| 9 | GET | `/api/pedidos/{id}` | Consulta pedido por identificador. |
| 10 | POST | `/api/pagos` | Procesa pago del pedido. |
| 11 | GET | `/api/reportes/ventas` | Genera reporte de ventas. |
| 12 | GET | `/api/reportes/inventario` | Genera reporte de inventario y bajo stock. |

## 6. Ejecucion

```bash
go test ./...
go run ./cmd/api
# Opcional, si el puerto 8080 esta ocupado:
PORT=18080 go run ./cmd/api
```

El servidor se ejecuta en:

```text
http://localhost:8080
```

## 7. Ejemplo de uso

```bash
curl http://localhost:8080/api/salud
curl http://localhost:8080/api/productos
curl -X POST http://localhost:8080/api/carrito/items -H "Content-Type: application/json" -d '{"clienteId":"CLI-001","productoId":"PROD-001","cantidad":1}'
curl -X POST http://localhost:8080/api/pedidos -H "Content-Type: application/json" -d '{"clienteId":"CLI-001"}'
```

## 8. Estructura del repositorio

```text
sistema-gestion-ecommerce-poo/
|-- README.md
|-- go.mod
|-- cmd/
|   `-- api/
|       `-- main.go
|-- internal/
|   `-- ecommerce/
|       |-- domain.go
|       |-- repository.go
|       |-- service.go
|       |-- http.go
|       `-- ecommerce_test.go
|-- docs/
|   |-- cronograma.md
|   |-- openapi.yml
|   |-- pruebas_funcionales.md
|   |-- guion_video_presentacion.md
|   `-- guion_video_demostracion.md
`-- tests/
```

## 9. Limitaciones reconocidas

La version final academica utiliza almacenamiento en memoria y pagos simulados. No integra pasarelas bancarias reales, facturacion electronica oficial, empresas de envio externas ni autenticacion productiva. Estas limitaciones son decisiones de alcance para concentrar la entrega en POO, servicios web, JSON, validaciones y demostracion funcional.

## 10. Visualizacion del futuro

El sistema puede evolucionar hacia una plataforma omnicanal donde los servicios web se conecten con inteligencia artificial para recomendaciones, sensores IoT para inventario, pagos digitales seguros, analitica de datos y canales moviles. Esta visualizacion muestra como la tecnologia puede mejorar la gestion empresarial, reducir errores operativos y apoyar decisiones comerciales basadas en informacion.
