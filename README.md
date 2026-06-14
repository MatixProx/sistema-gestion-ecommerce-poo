# Sistema de Gestion de E-commerce - Aprendizaje Autonomo 2

Proyecto academico de Programacion Orientada a Objetos desarrollado en Go. Esta entrega evoluciona la planeacion del Aprendizaje Autonomo 1 hacia una primera version codificada del sistema de gestion de e-commerce.

## Objetivo

Implementar un avance significativo del sistema aplicando estructuras de datos, encapsulacion, manejo de errores e interfaces dentro de un entorno de programacion orientado a objetos.

## Relacion con la tarea

La tarea solicita continuar el plan, iniciar la codificacion del sistema, identificar las clases o tipos que se desarrollan y aplicar temas vistos en las unidades, especialmente encapsulacion, errores e interfaces. Este repositorio responde con una version funcional de consola que modela usuarios, productos, inventario, carrito, pedidos, pagos y reportes.

## Funcionalidades implementadas

- Registro conceptual de usuarios mediante `Usuario`, `Cliente` y `Administrador`.
- Encapsulacion real con campos privados y metodos getter/setter.
- Validaciones de correo, contrasena, nombre, precio, stock y cantidades.
- Manejo de errores con errores predefinidos y errores formateados.
- Inventario basado en `map[string]*Producto` para busqueda eficiente por ID.
- Carrito basado en `[]DetalleCarrito` para agregar, eliminar y modificar productos.
- Pedido con estados controlados: pendiente, pagado, enviado y cancelado.
- Interfaz `MetodoPago` para aplicar polimorfismo en tarjeta, transferencia y contra entrega.
- Reportes administrativos usando slices y concurrencia basica con goroutines y canales.
- Pruebas unitarias con `go test`.

## Estructura del proyecto

```text
sistema-gestion-ecommerce-poo/
|-- go.mod
|-- README.md
|-- .gitignore
|-- cmd/demo/main.go
|-- internal/ecommerce/
|   |-- usuario.go
|   |-- producto.go
|   |-- inventario.go
|   |-- carrito.go
|   |-- pedido.go
|   |-- pago.go
|   |-- reporte.go
|   |-- errors.go
|   |-- ecommerce_test.go
|-- docs/
|   |-- avance-autonomo-2.md
|   |-- guion-video-avance.md
|   |-- guion-video-demostracion.md
|   |-- comandos-git.md
```

## Como ejecutar

```bash
go run ./cmd/demo
```

## Como probar

```bash
go test ./...
```

## Evidencia para video

1. Mostrar el README y explicar que el proyecto continua el e-commerce planificado.
2. Mostrar los archivos de `internal/ecommerce` y explicar las entidades principales.
3. Ejecutar `go test ./...` para demostrar pruebas.
4. Ejecutar `go run ./cmd/demo` para demostrar inventario, carrito, pedido, pago y reporte.
5. Mostrar el repositorio en GitHub con los commits del avance.

## Autor

Matias Castro Guerra

## Estado

Etapa 2: Desarrollo del sistema con avance funcional.
