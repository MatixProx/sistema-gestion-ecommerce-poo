# Guion para video explicativo del avance

Duracion sugerida: 3 a 5 minutos.

## 1. Presentacion

Hola, mi nombre es Matias Castro Guerra. En este video presento el avance del Aprendizaje Autonomo 2 de Programacion Orientada a Objetos. El proyecto continua el Sistema de Gestion de E-commerce planificado en la entrega anterior.

## 2. Continuacion del plan

En el primer trabajo se definieron los modulos principales: usuarios, productos, carrito, pedidos, pagos, inventario y reportes. En esta entrega se comenzo la codificacion de esos modulos en Go, aplicando estructuras de datos, structs, metodos, encapsulacion, manejo de errores e interfaces.

## 3. Estructura del repositorio

Aqui se observa la estructura del proyecto. El archivo `go.mod` define el modulo. En `cmd/demo` esta la demostracion principal. En `internal/ecommerce` se encuentran las entidades y servicios del sistema. En `docs` esta la documentacion del avance y los guiones de video.

## 4. Clases o tipos implementados

Se implementaron los tipos `Usuario`, `Cliente`, `Administrador`, `Producto`, `Categoria`, `Inventario`, `Carrito`, `Pedido`, `Pago` y `ReporteService`. Cada tipo tiene una responsabilidad concreta para mantener el codigo organizado.

## 5. Encapsulacion

Los atributos se mantienen privados y se accede a ellos mediante metodos publicos. Por ejemplo, el producto tiene precio y stock privados, y solo se modifican mediante metodos como `SetPrecio`, `SetStock`, `ReducirStock` y `AumentarStock`. Esto evita datos invalidos.

## 6. Manejo de errores

El archivo `errors.go` centraliza errores como correo invalido, precio invalido, stock insuficiente y carrito vacio. Estos errores permiten controlar situaciones incorrectas sin que el programa falle de forma inesperada.

## 7. Interfaces

La interfaz `MetodoPago` permite aplicar polimorfismo. Actualmente existen tres implementaciones: pago con tarjeta, transferencia y contra entrega. Todas procesan un pedido, pero cada una lo hace de manera diferente.

## 8. Cierre

Con este avance el sistema ya no queda solo en planeacion, sino que tiene una primera version funcional y probada. Muchas gracias.
