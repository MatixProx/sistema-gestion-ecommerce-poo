package com.ecommerce.modelo;

public class Pago {
    private final int id;
    private final MetodoPago metodoPago;
    private final double monto;
    private boolean aprobado;

    public Pago(int id, MetodoPago metodoPago, double monto) {
        this.id = id;
        this.metodoPago = metodoPago;
        this.monto = monto;
    }

    public boolean procesarPago(Pedido pedido) {
        if (monto >= pedido.calcularTotal()) {
            aprobado = true;
            pedido.cambiarEstado(EstadoPedido.PAGADO);
        }
        return aprobado;
    }

    public MetodoPago getMetodoPago() {
        return metodoPago;
    }

    public boolean isAprobado() {
        return aprobado;
    }
}
