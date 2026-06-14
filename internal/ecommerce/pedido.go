package ecommerce

import (
	"fmt"
	"time"
)

type EstadoPedido string

const (
	EstadoPendiente EstadoPedido = "pendiente"
	EstadoPagado    EstadoPedido = "pagado"
	EstadoEnviado   EstadoPedido = "enviado"
	EstadoCancelado EstadoPedido = "cancelado"
)

type DetallePedido struct {
	productoID string
	nombre     string
	cantidad   int
	precio     float64
}

func (d DetallePedido) ProductoID() string { return d.productoID }
func (d DetallePedido) Nombre() string     { return d.nombre }
func (d DetallePedido) Cantidad() int      { return d.cantidad }
func (d DetallePedido) Precio() float64    { return d.precio }
func (d DetallePedido) Subtotal() float64  { return d.precio * float64(d.cantidad) }

type Pedido struct {
	id        string
	clienteID string
	estado    EstadoPedido
	total     float64
	fecha     time.Time
	detalles  []DetallePedido
}

func NewPedido(id string, cliente Cliente, carrito Carrito, inventario *Inventario) (*Pedido, error) {
	if carrito.Vacio() {
		return nil, ErrCarritoVacio
	}
	detalles := make([]DetallePedido, 0, len(carrito.Items()))
	for _, item := range carrito.Items() {
		producto, err := inventario.BuscarProducto(item.Producto().ID())
		if err != nil {
			return nil, err
		}
		if err := producto.ReducirStock(item.Cantidad()); err != nil {
			return nil, err
		}
		detalles = append(detalles, DetallePedido{
			productoID: producto.ID(),
			nombre:     producto.Nombre(),
			cantidad:   item.Cantidad(),
			precio:     producto.Precio(),
		})
	}
	return &Pedido{
		id:        id,
		clienteID: cliente.ID(),
		estado:    EstadoPendiente,
		total:     carrito.CalcularTotal(),
		fecha:     time.Now(),
		detalles:  detalles,
	}, nil
}

func (p Pedido) ID() string                { return p.id }
func (p Pedido) ClienteID() string         { return p.clienteID }
func (p Pedido) Estado() EstadoPedido      { return p.estado }
func (p Pedido) Total() float64            { return p.total }
func (p Pedido) Fecha() time.Time          { return p.fecha }
func (p Pedido) Detalles() []DetallePedido { return append([]DetallePedido{}, p.detalles...) }

func (p *Pedido) CambiarEstado(estado EstadoPedido) error {
	switch estado {
	case EstadoPendiente, EstadoPagado, EstadoEnviado, EstadoCancelado:
		p.estado = estado
		return nil
	default:
		return fmt.Errorf("estado %s: %w", estado, ErrEstadoInvalido)
	}
}

func (p *Pedido) Cancelar() error {
	if p.estado == EstadoEnviado {
		return fmt.Errorf("pedido %s enviado: no se puede cancelar", p.id)
	}
	p.estado = EstadoCancelado
	return nil
}
