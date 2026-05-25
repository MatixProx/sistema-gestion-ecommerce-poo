package com.ecommerce.modelo;

public class Producto {
    private final int id;
    private String nombre;
    private double precio;
    private int stock;
    private String descripcion;
    private Categoria categoria;

    public Producto(int id, String nombre, double precio, int stock, String descripcion, Categoria categoria) {
        if (precio < 0) {
            throw new IllegalArgumentException("El precio no puede ser negativo.");
        }
        if (stock < 0) {
            throw new IllegalArgumentException("El stock no puede ser negativo.");
        }
        this.id = id;
        this.nombre = nombre;
        this.precio = precio;
        this.stock = stock;
        this.descripcion = descripcion;
        this.categoria = categoria;
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

    public double getPrecio() {
        return precio;
    }

    public void setPrecio(double precio) {
        if (precio < 0) {
            throw new IllegalArgumentException("El precio no puede ser negativo.");
        }
        this.precio = precio;
    }

    public int getStock() {
        return stock;
    }

    public void actualizarStock(int nuevoStock) {
        if (nuevoStock < 0) {
            throw new IllegalArgumentException("El stock no puede ser negativo.");
        }
        this.stock = nuevoStock;
    }

    public void disminuirStock(int cantidad) {
        if (cantidad <= 0) {
            throw new IllegalArgumentException("La cantidad debe ser mayor que cero.");
        }
        if (cantidad > stock) {
            throw new IllegalStateException("No existe stock suficiente.");
        }
        this.stock -= cantidad;
    }

    public String getDescripcion() {
        return descripcion;
    }

    public Categoria getCategoria() {
        return categoria;
    }
}
