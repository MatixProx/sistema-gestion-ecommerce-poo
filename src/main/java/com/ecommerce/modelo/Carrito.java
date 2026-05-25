package com.ecommerce.modelo;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class Carrito {
    private final List<ItemCarrito> items = new ArrayList<>();

    public void agregarProducto(Producto producto, int cantidad) {
        items.add(new ItemCarrito(producto, cantidad));
    }

    public void eliminarProducto(int productoId) {
        items.removeIf(item -> item.getProducto().getId() == productoId);
    }

    public double calcularTotal() {
        return items.stream()
                .mapToDouble(ItemCarrito::calcularSubtotal)
                .sum();
    }

    public List<ItemCarrito> getItems() {
        return Collections.unmodifiableList(items);
    }
}
