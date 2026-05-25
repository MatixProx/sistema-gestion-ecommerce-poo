package com.ecommerce;

import com.ecommerce.modelo.Categoria;
import com.ecommerce.modelo.Producto;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertThrows;

class ProductoTest {
    @Test
    void noPermitePrecioNegativo() {
        Categoria categoria = new Categoria(1, "Tecnologia");
        assertThrows(IllegalArgumentException.class, () ->
                new Producto(1, "Producto", -1, 10, "Descripcion", categoria)
        );
    }
}
