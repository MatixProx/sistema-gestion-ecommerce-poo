# Sistema de Gestion de E-commerce

## Descripcion
Proyecto academico desarrollado para la asignatura Programacion Orientada a Objetos. El sistema tiene como finalidad planificar y estructurar un Sistema de Gestion de e-commerce que permita administrar usuarios, productos, carrito de compras, pedidos, pagos, inventario y reportes administrativos.

## Objetivo general
Disenar la planeacion de un sistema de gestion de e-commerce aplicando principios de Programacion Orientada a Objetos, programacion funcional y conceptos de gestion empresarial.

## Objetivos especificos
- Identificar los modulos principales del sistema.
- Definir funcionalidades por modulo y su alcance inicial.
- Representar entidades del negocio mediante clases y objetos.
- Aplicar abstraccion, encapsulamiento, herencia y polimorfismo.
- Incorporar operaciones funcionales para filtrar, transformar y resumir informacion.
- Documentar el proyecto y organizarlo en un repositorio de GitHub.

## Modulos principales
1. Gestion de usuarios
2. Gestion de productos
3. Carrito de compras
4. Gestion de pedidos
5. Gestion de pagos
6. Gestion de inventario
7. Reportes administrativos

## Principios de Programacion Orientada a Objetos aplicados
- **Abstraccion:** clases como `Producto`, `Cliente`, `Pedido` y `Pago` representan elementos reales del negocio.
- **Encapsulamiento:** los atributos se protegen mediante acceso privado y metodos publicos.
- **Herencia:** `Cliente` y `Administrador` heredan de `Usuario`.
- **Polimorfismo:** el procesamiento de pagos puede variar segun el metodo seleccionado.

## Programacion funcional aplicada
El proyecto incorpora operaciones funcionales mediante expresiones lambda, `streams`, filtros, mapeos y reducciones para calcular totales, consultar productos con bajo stock y generar reportes.

## Tecnologias propuestas
- Java 17
- Maven
- JUnit 5
- MySQL
- MySQL Connector/J
- Git
- GitHub

## Estructura del repositorio
```text
sistema-gestion-ecommerce-poo/
|
|-- README.md
|-- .gitignore
|-- pom.xml
|
|-- docs/
|   |-- planeacion-software.pdf
|   |-- planeacion-software.docx
|   |-- alcance-proyecto.md
|   |-- diagrama-clases.png
|   |-- diagrama-clases.mmd
|
|-- src/
|   |-- main/java/com/ecommerce/
|   |   |-- modelo/
|   |   |-- servicio/
|   |   |-- util/
|   |   |-- Main.java
|   |
|   |-- test/java/com/ecommerce/
```

## Estado del proyecto
Etapa 1: Planeacion del software.

## Integrantes
- Nombre del integrante 1
- Nombre del integrante 2
- Nombre del integrante 3
- Nombre del integrante 4
