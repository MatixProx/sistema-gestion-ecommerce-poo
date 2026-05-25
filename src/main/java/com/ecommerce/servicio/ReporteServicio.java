package com.ecommerce.servicio;

import com.ecommerce.modelo.EstadoPedido;
import com.ecommerce.modelo.Pedido;

import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

public class ReporteServicio {
    public double calcularVentasTotales(List<Pedido> pedidos) {
        return pedidos.stream()
                .filter(pedido -> pedido.getEstado() == EstadoPedido.PAGADO)
                .mapToDouble(Pedido::calcularTotal)
                .sum();
    }

    public Map<EstadoPedido, Long> contarPedidosPorEstado(List<Pedido> pedidos) {
        return pedidos.stream()
                .collect(Collectors.groupingBy(Pedido::getEstado, Collectors.counting()));
    }
}
