# Guion para video de demostracion de funcionalidades

Duracion sugerida: 4 a 6 minutos.

## 1. Abrir terminal

Explicar que primero se ejecutaran las pruebas unitarias para comprobar que las funcionalidades principales trabajan correctamente.

```bash
go test ./...
```

Mencionar que las pruebas validan producto, carrito, pedido, pago e informes.

## 2. Ejecutar demo

```bash
go run ./cmd/demo
```

## 3. Explicar salida

La salida muestra el flujo completo del sistema:

1. Se crea un cliente.
2. Se agregan productos al inventario.
3. Se agregan productos al carrito.
4. Se calcula el total.
5. Se crea un pedido.
6. Se procesa un pago con tarjeta.
7. Se genera un reporte administrativo.

## 4. Mostrar encapsulacion

Abrir `producto.go` y explicar que los campos del producto son privados. Mostrar los metodos `SetPrecio`, `SetStock` y `ReducirStock`.

## 5. Mostrar errores

Abrir `errors.go` y explicar que los errores centralizados ayudan a controlar entradas incorrectas, stock insuficiente y pagos invalidos.

## 6. Mostrar interfaces

Abrir `pago.go` y explicar que `MetodoPago` es una interfaz. Luego mostrar `TarjetaPago`, `TransferenciaPago` y `ContraEntregaPago` como implementaciones polimorficas.

## 7. Mostrar funcionalidad compleja

Abrir `reporte.go` y explicar que el reporte usa goroutines y canales para calcular ventas, estados y productos vendidos en paralelo.

## 8. Cierre

Indicar que el codigo queda cargado en GitHub y que esta version representa un avance funcional del Sistema de Gestion de E-commerce.
