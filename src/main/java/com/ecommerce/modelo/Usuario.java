package com.ecommerce.modelo;

public abstract class Usuario {
    private final int id;
    private String nombre;
    private String correo;
    private String contrasena;

    protected Usuario(int id, String nombre, String correo, String contrasena) {
        this.id = id;
        this.nombre = nombre;
        this.correo = correo;
        this.contrasena = contrasena;
    }

    public int getId() {
        return id;
    }

    public String getNombre() {
        return nombre;
    }

    public void setNombre(String nombre) {
        this.nombre = nombre;
    }

    public String getCorreo() {
        return correo;
    }

    public void setCorreo(String correo) {
        this.correo = correo;
    }

    public boolean validarContrasena(String contrasenaIngresada) {
        return contrasena != null && contrasena.equals(contrasenaIngresada);
    }

    public abstract String obtenerRol();
}
