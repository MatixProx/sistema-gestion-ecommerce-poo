package com.ecommerce;

import com.ecommerce.modelo.Categoria;
import com.ecommerce.modelo.Cliente;
import com.ecommerce.modelo.Producto;
import com.ecommerce.servicio.ProductoServicio;

public class Main {
    public static void main(String[] args) {
        Categoria tecnologia = new Categoria(1, "Tecnologia");
        Producto audifonos = new Producto(1, "Audifonos Bluetooth", 29.99, 15, "Audifonos inalambricos", tecnologia);
        ProductoServicio productoServicio = new ProductoServicio();
        productoServicio.registrarProducto(audifonos);

        Cliente cliente = new Cliente(1, "Cliente de prueba", "cliente@email.com", "1234", "Quito");
        cliente.getCarrito().agregarProducto(audifonos, 2);

        System.out.println("Sistema de Gestion de E-commerce");
        System.out.println("Total del carrito: $" + cliente.getCarrito().calcularTotal());
    }
}
