# Cumplimiento de rubrica - Aprendizaje Autonomo 2

## Evidencia Paso 1
- El informe continua el plan del Aprendizaje Autonomo 1 y conecta el codigo con los modulos ya definidos.
- Se usan fuentes tecnicas, academicas y del proyecto, pero tambien se reconocen sus limites.
- Se cuestiona que la planeacion sola no demuestra funcionamiento; por eso se agrega codigo, pruebas y demostracion.

## Posicion del estudiante Paso 1
La postura del estudiante es construir primero la logica de dominio del e-commerce antes de integrar pantallas, base de datos o pasarelas reales. Esta decision se justifica porque productos, inventario, carrito, pedidos y pagos deben funcionar correctamente antes de crecer hacia una version comercial.

## Conclusiones y logros Paso 1
- El avance transforma la planeacion en codigo funcional.
- Se implementan usuarios, productos, inventario, carrito, pedidos, pagos y reportes.
- Se reconocen limites: aun falta base de datos, interfaz grafica y servicios externos.

## Encapsulacion Paso 2
- Los campos de Usuario, Producto, Categoria, Carrito, Pedido y Pago son privados.
- Se usan getters y setters para proteger la informacion.
- Los setters validan precio, stock, correo, campos vacios y cantidades.

## Manejo de errores e interfaces Paso 2
- Los errores estan centralizados en internal/ecommerce/errors.go.
- El sistema valida carrito vacio, stock insuficiente, pago invalido y estado incorrecto.
- La interfaz MetodoPago permite procesar tarjeta, transferencia y contra entrega de forma polimorfica.

## Comentarios en funcionalidades complejas
- ReporteService.GenerarResumen esta comentado porque usa goroutines y canales.
- Tambien hay comentarios en tipos importantes como Usuario, Producto, Inventario y MetodoPago.
