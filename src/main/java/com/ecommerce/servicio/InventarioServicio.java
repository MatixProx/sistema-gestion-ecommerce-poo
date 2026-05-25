package com.ecommerce.servicio;

import com.ecommerce.modelo.Producto;

import java.util.List;
import java.util.stream.Collectors;

public class InventarioServicio {
    public List<Producto> obtenerProductosConBajoStock(List<Producto> productos, int limiteStock) {
        return productos.stream()
                .filter(producto -> producto.getStock() <= limiteStock)
                .collect(Collectors.toList());
    }

    public void descontarStock(Producto producto, int cantidad) {
        producto.disminuirStock(cantidad);
    }
}
