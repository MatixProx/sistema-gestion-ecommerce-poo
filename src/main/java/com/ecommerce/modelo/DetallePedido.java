package com.ecommerce.modelo;

public class DetallePedido {
    private final Producto producto;
    private final int cantidad;
    private final double precioUnitario;

    public DetallePedido(Producto producto, int cantidad) {
        this.producto = producto;
        this.cantidad = cantidad;
        this.precioUnitario = producto.getPrecio();
    }

    public Producto getProducto() {
        return producto;
    }

    public int getCantidad() {
        return cantidad;
    }

    public double calcularSubtotal() {
        return precioUnitario * cantidad;
    }
}
