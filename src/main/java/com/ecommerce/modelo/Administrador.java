package com.ecommerce.modelo;

public class Administrador extends Usuario {
    public Administrador(int id, String nombre, String correo, String contrasena) {
        super(id, nombre, correo, contrasena);
    }

    @Override
    public String obtenerRol() {
        return "ADMINISTRADOR";
    }
}
