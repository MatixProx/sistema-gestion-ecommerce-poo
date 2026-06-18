package ecommerce

import (
	"fmt"
	"strings"
	"time"
)

// MetodoPago define un contrato comun. Cualquier struct que implemente
// Nombre y Procesar puede usarse como metodo de pago del sistema.
type MetodoPago interface {
	Nombre() string
	Procesar(pedido *Pedido) (*Pago, error)
}

type Pago struct {
	id       string
	metodo   string
	monto    float64
	estado   string
	fecha    time.Time
	pedidoID string
}

func (p Pago) ID() string       { return p.id }
func (p Pago) Metodo() string   { return p.metodo }
func (p Pago) Monto() float64   { return p.monto }
func (p Pago) Estado() string   { return p.estado }
func (p Pago) PedidoID() string { return p.pedidoID }

type TarjetaPago struct {
	numeroTarjeta string
	titular       string
}

func NewTarjetaPago(numeroTarjeta, titular string) TarjetaPago {
	return TarjetaPago{numeroTarjeta: numeroTarjeta, titular: titular}
}

func (t TarjetaPago) Nombre() string { return "tarjeta" }

func (t TarjetaPago) Procesar(pedido *Pedido) (*Pago, error) {
	if pedido == nil || pedido.Total() <= 0 {
		return nil, ErrPagoInvalido
	}
	if len(strings.TrimSpace(t.numeroTarjeta)) < 8 || strings.TrimSpace(t.titular) == "" {
		return nil, fmt.Errorf("tarjeta: %w", ErrPagoInvalido)
	}
	if err := pedido.CambiarEstado(EstadoPagado); err != nil {
		return nil, err
	}
	return newPago("PAG-TAR-"+pedido.ID(), t.Nombre(), pedido), nil
}

type TransferenciaPago struct {
	codigoOperacion string
}

func NewTransferenciaPago(codigoOperacion string) TransferenciaPago {
	return TransferenciaPago{codigoOperacion: codigoOperacion}
}

func (t TransferenciaPago) Nombre() string { return "transferencia" }

func (t TransferenciaPago) Procesar(pedido *Pedido) (*Pago, error) {
	if pedido == nil || pedido.Total() <= 0 || strings.TrimSpace(t.codigoOperacion) == "" {
		return nil, fmt.Errorf("transferencia: %w", ErrPagoInvalido)
	}
	if err := pedido.CambiarEstado(EstadoPagado); err != nil {
		return nil, err
	}
	return newPago("PAG-TRA-"+pedido.ID(), t.Nombre(), pedido), nil
}

type ContraEntregaPago struct{}

func (c ContraEntregaPago) Nombre() string { return "contra entrega" }

func (c ContraEntregaPago) Procesar(pedido *Pedido) (*Pago, error) {
	if pedido == nil || pedido.Total() <= 0 {
		return nil, ErrPagoInvalido
	}
	// En contra entrega el pedido queda pendiente hasta que el repartidor confirme cobro.
	return &Pago{
		id:       "PAG-CE-" + pedido.ID(),
		metodo:   c.Nombre(),
		monto:    pedido.Total(),
		estado:   "pendiente de cobro",
		fecha:    time.Now(),
		pedidoID: pedido.ID(),
	}, nil
}

func newPago(id, metodo string, pedido *Pedido) *Pago {
	return &Pago{
		id:       id,
		metodo:   metodo,
		monto:    pedido.Total(),
		estado:   "confirmado",
		fecha:    time.Now(),
		pedidoID: pedido.ID(),
	}
}
