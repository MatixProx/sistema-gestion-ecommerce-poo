package com.ecommerce.modelo;

import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public class Pedido {
    private final int id;
    private final Cliente cliente;
    private final LocalDateTime fechaCreacion;
    private EstadoPedido estado;
    private final List<DetallePedido> detalles;

    public Pedido(int id, Cliente cliente, List<DetallePedido> detalles) {
        this.id = id;
        this.cliente = cliente;
        this.fechaCreacion = LocalDateTime.now();
        this.estado = EstadoPedido.PENDIENTE;
        this.detalles = new ArrayList<>(detalles);
    }

    public int getId() {
        return id;
    }

    public Cliente getCliente() {
        return cliente;
    }

    public EstadoPedido getEstado() {
        return estado;
    }

    public void cambiarEstado(EstadoPedido estado) {
        this.estado = estado;
    }

    public double calcularTotal() {
        return detalles.stream()
                .mapToDouble(DetallePedido::calcularSubtotal)
                .sum();
    }

    public LocalDateTime getFechaCreacion() {
        return fechaCreacion;
    }

    public List<DetallePedido> getDetalles() {
        return Collections.unmodifiableList(detalles);
    }
}
