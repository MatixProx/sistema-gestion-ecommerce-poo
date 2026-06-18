# Aprendizaje Autonomo 2: Desarrollo del Sistema de Gestion de E-commerce

## 1. Introduccion

El presente avance continua la planeacion del Sistema de Gestion de E-commerce realizada en el Aprendizaje Autonomo 1. En esta segunda etapa se inicia la codificacion del sistema y se implementan entidades principales del negocio: usuarios, clientes, administradores, productos, categorias, inventario, carrito, pedidos, pagos y reportes administrativos.

El avance se desarrolla en Go porque las unidades de la asignatura trabajan sintaxis, paquetes, estructuras de datos, structs, metodos, constructores, encapsulacion, manejo de errores, interfaces, polimorfismo, concurrencia, servicios web, serializacion y testing. La propuesta busca demostrar que el proyecto no solo esta planeado, sino que ya cuenta con una base funcional, organizada y comprobable.

## 2. Objetivo del avance

Implementar un avance significativo del Sistema de Gestion de E-commerce aplicando estructuras de datos, encapsulacion, manejo de errores e interfaces, de manera coherente con la planeacion del trabajo autonomo anterior y con los temas revisados en clase.

## 3. Continuacion del plan

En la primera entrega se definieron los modulos del sistema: gestion de usuarios, gestion de productos, carrito de compras, pedidos, pagos, inventario y reportes administrativos. Para esta entrega se priorizo una primera version funcional de consola que permite demostrar el flujo principal de una compra:

1. Crear categorias y productos.
2. Registrar un cliente.
3. Agregar productos al inventario.
4. Agregar productos al carrito.
5. Crear un pedido y reducir el stock.
6. Procesar el pago con una interfaz polimorfica.
7. Generar reportes administrativos.

## 4. Clases o tipos desarrollados

| Tipo / struct | Responsabilidad | Concepto aplicado |
|---|---|---|
| `Usuario` | Datos base de usuarios del sistema | Encapsulacion, getters y setters |
| `Cliente` | Usuario que realiza compras | Composicion de structs |
| `Administrador` | Usuario con permisos de gestion | Composicion y roles |
| `Categoria` | Clasifica productos | Struct y validaciones |
| `Producto` | Representa articulo vendido | Encapsulacion, validacion de precio y stock |
| `Inventario` | Administra productos por ID | Map como estructura de datos |
| `Carrito` | Administra productos seleccionados | Slice y calculo de total |
| `Pedido` | Formaliza la compra | Estados y control de flujo |
| `Pago` | Registra el pago de un pedido | Struct y metodos |
| `MetodoPago` | Contrato para formas de pago | Interface y polimorfismo |
| `ReporteService` | Genera indicadores administrativos | Slices, maps, goroutines y canales |

## 5. Encapsulacion implementada

Los atributos de las estructuras principales se declararon privados, por ejemplo `nombre`, `correo`, `precio`, `stock`, `items` y `estado`. Para acceder a ellos se usan metodos publicos como `Nombre()`, `Correo()`, `Precio()`, `Stock()` y `CalcularTotal()`. Para modificar datos sensibles se usan setters como `SetCorreo`, `SetPrecio` y `SetStock`, que validan la informacion antes de guardar cambios.

Esta decision protege la integridad del sistema porque evita que desde otras partes del programa se asignen valores incorrectos, como precio cero, stock negativo, correo invalido o cantidades menores a uno.

## 6. Manejo de errores

El sistema incorpora errores reutilizables en el archivo `errors.go`, por ejemplo:

- `ErrCampoVacio`
- `ErrCorreoInvalido`
- `ErrPrecioInvalido`
- `ErrStockInsuficiente`
- `ErrCarritoVacio`
- `ErrPagoInvalido`
- `ErrEstadoInvalido`

Ademas, se usan errores formateados con contexto, por ejemplo cuando un producto no tiene stock suficiente o cuando un campo especifico esta vacio. Esto facilita la depuracion y permite entender la causa del problema sin revisar todo el programa.

## 7. Interfaces y polimorfismo

La interfaz principal implementada es `MetodoPago`:

```go
type MetodoPago interface {
    Nombre() string
    Procesar(pedido *Pedido) (*Pago, error)
}
```

Esta interfaz permite que diferentes metodos de pago se comporten de forma distinta, pero mantengan el mismo contrato. Las implementaciones actuales son:

- `TarjetaPago`
- `TransferenciaPago`
- `ContraEntregaPago`

Gracias a esto, el sistema puede procesar pagos sin depender de una clase concreta. Si en el futuro se agrega PayPal, PayPhone o una pasarela real, solo se debe crear un nuevo struct que implemente la misma interfaz.

## 8. Estructuras de datos aplicadas

Se aplicaron estructuras de datos revisadas en clase:

- `map[string]*Producto` en inventario para busqueda eficiente por identificador.
- `[]DetalleCarrito` en carrito para manejar dinamicamente los productos seleccionados.
- `[]Pedido` en reportes para recorrer pedidos y calcular estadisticas.
- `map[EstadoPedido]int` para agrupar pedidos por estado.
- `map[string]int` para calcular productos mas vendidos.

## 9. Funcionalidades complejas comentadas

La funcion `GenerarResumen` contiene comentarios explicativos porque integra concurrencia con goroutines y canales. Esta funcion calcula simultaneamente ventas totales, pedidos por estado y productos mas vendidos. Esto evidencia el uso de temas avanzados sin perder claridad para quien revise el codigo.

## 10. Pruebas realizadas

Se agregaron pruebas unitarias con `go test` para validar:

- Validacion de precio de producto.
- Calculo total del carrito.
- Reduccion de stock al crear pedido.
- Procesamiento polimorfico de pagos.
- Generacion concurrente de reportes.

## 11. Evidencia de ejecucion

Comando para ejecutar pruebas:

```bash
go test ./...
```

Comando para ejecutar demostracion:

```bash
go run ./cmd/demo
```

La demostracion muestra cliente, total del carrito, estado del pedido, pago confirmado, ventas confirmadas, pedidos por estado, productos mas vendidos y productos con bajo stock.

## 12. Conclusion

El avance desarrollado cumple con la continuacion del plan del Sistema de Gestion de E-commerce porque transforma la planeacion inicial en una base funcional de codigo. La implementacion evidencia el uso de structs, metodos, composicion, estructuras de datos, encapsulacion, manejo de errores e interfaces. Ademas, se incorporan pruebas y reportes para mostrar que el sistema puede crecer de forma ordenada.

La solucion mantiene coherencia con el contexto empresarial del e-commerce, ya que permite representar clientes, productos, inventario, carrito, pedidos, pagos y reportes. Por ello, el trabajo demuestra transferencia de los temas de la asignatura hacia un problema real de gestion digital.
