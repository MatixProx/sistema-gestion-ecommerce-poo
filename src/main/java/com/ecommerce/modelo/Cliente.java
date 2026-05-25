package com.ecommerce.modelo;

public class Cliente extends Usuario {
    private String direccion;
    private final Carrito carrito;

    public Cliente(int id, String nombre, String correo, String contrasena, String direccion) {
        super(id, nombre, correo, contrasena);
        this.direccion = direccion;
        this.carrito = new Carrito();
    }

    public String getDireccion() {
        return direccion;
    }

    public void setDireccion(String direccion) {
        this.direccion = direccion;
    }

    public Carrito getCarrito() {
        return carrito;
    }

    @Override
    public String obtenerRol() {
        return "CLIENTE";
    }
}
