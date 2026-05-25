package com.ecommerce.servicio;

import com.ecommerce.modelo.Producto;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;

public class ProductoServicio {
    private final List<Producto> productos = new ArrayList<>();

    public void registrarProducto(Producto producto) {
        productos.add(producto);
    }

    public Optional<Producto> buscarPorId(int id) {
        return productos.stream()
                .filter(producto -> producto.getId() == id)
                .findFirst();
    }

    public List<Producto> listarPorCategoria(String categoria) {
        return productos.stream()
                .filter(producto -> producto.getCategoria().getNombre().equalsIgnoreCase(categoria))
                .collect(Collectors.toList());
    }

    public List<Producto> listarOrdenadosPorPrecio() {
        return productos.stream()
                .sorted(Comparator.comparingDouble(Producto::getPrecio))
                .collect(Collectors.toList());
    }

    public List<Producto> listarTodos() {
        return List.copyOf(productos);
    }
}
