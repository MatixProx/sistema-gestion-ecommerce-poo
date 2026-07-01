package ecommerce

import (
	"errors"
	"fmt"
	"time"
)

// Producto modela un articulo del catalogo. Sus atributos internos permanecen
// encapsulados y solo se exponen mediante metodos controlados o DTOs.
type Producto struct {
	id          string
	nombre      string
	categoria   string
	descripcion string
	precio      float64
	stock       int
}

type ProductoDTO struct {
	ID          string  `json:"id"`
	Nombre      string  `json:"nombre"`
	Categoria   string  `json:"categoria"`
	Descripcion string  `json:"descripcion"`
	Precio      float64 `json:"precio"`
	Stock       int     `json:"stock"`
}

func NuevoProducto(id, nombre, categoria, descripcion string, precio float64, stock int) (*Producto, error) {
	if id == "" || nombre == "" {
		return nil, errors.New("el producto debe tener id y nombre")
	}
	if precio <= 0 {
		return nil, errors.New("el precio debe ser mayor a cero")
	}
	if stock < 0 {
		return nil, errors.New("el stock no puede ser negativo")
	}
	return &Producto{id: id, nombre: nombre, categoria: categoria, descripcion: descripcion, precio: precio, stock: stock}, nil
}

func (p *Producto) ID() string        { return p.id }
func (p *Producto) Nombre() string    { return p.nombre }
func (p *Producto) Precio() float64   { return p.precio }
func (p *Producto) Stock() int        { return p.stock }
func (p *Producto) Categoria() string { return p.categoria }

func (p *Producto) DescontarStock(cantidad int) error {
	if cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor a cero")
	}
	if cantidad > p.stock {
		return fmt.Errorf("stock insuficiente para %s: disponible %d, solicitado %d", p.nombre, p.stock, cantidad)
	}
	p.stock -= cantidad
	return nil
}

func (p *Producto) AumentarStock(cantidad int) error {
	if cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor a cero")
	}
	p.stock += cantidad
	return nil
}

func (p *Producto) DTO() ProductoDTO {
	return ProductoDTO{
		ID: p.id, Nombre: p.nombre, Categoria: p.categoria,
		Descripcion: p.descripcion, Precio: p.precio, Stock: p.stock,
	}
}

type Usuario struct {
	ID     string `json:"id"`
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
	Rol    string `json:"rol"`
}

type Cliente struct {
	Usuario
	Direccion string `json:"direccion"`
	Telefono  string `json:"telefono"`
}

type Administrador struct {
	Usuario
	Cargo string `json:"cargo"`
}

type DetalleCarrito struct {
	Producto ProductoDTO `json:"producto"`
	Cantidad int         `json:"cantidad"`
	Subtotal float64     `json:"subtotal"`
}

type Carrito struct {
	ClienteID string `json:"clienteId"`
	items     map[string]DetalleCarrito
}

func NuevoCarrito(clienteID string) *Carrito {
	return &Carrito{ClienteID: clienteID, items: make(map[string]DetalleCarrito)}
}

func (c *Carrito) AgregarProducto(producto *Producto, cantidad int) error {
	if producto == nil {
		return errors.New("producto inexistente")
	}
	if cantidad <= 0 {
		return errors.New("la cantidad debe ser mayor a cero")
	}
	if cantidad > producto.Stock() {
		return errors.New("stock insuficiente para agregar al carrito")
	}

	actual := c.items[producto.ID()]
	nuevaCantidad := actual.Cantidad + cantidad
	if nuevaCantidad > producto.Stock() {
		return errors.New("la cantidad total del carrito supera el stock disponible")
	}
	c.items[producto.ID()] = DetalleCarrito{
		Producto: producto.DTO(),
		Cantidad: nuevaCantidad,
		Subtotal: producto.Precio() * float64(nuevaCantidad),
	}
	return nil
}

func (c *Carrito) EliminarProducto(productoID string) {
	delete(c.items, productoID)
}

func (c *Carrito) Items() []DetalleCarrito {
	detalles := make([]DetalleCarrito, 0, len(c.items))
	for _, item := range c.items {
		detalles = append(detalles, item)
	}
	return detalles
}

func (c *Carrito) Total() float64 {
	total := 0.0
	for _, item := range c.items {
		total += item.Subtotal
	}
	return total
}

func (c *Carrito) EstaVacio() bool {
	return len(c.items) == 0
}

func (c *Carrito) Limpiar() {
	c.items = make(map[string]DetalleCarrito)
}

type EstadoPedido string

const (
	EstadoPendiente EstadoPedido = "pendiente"
	EstadoPagado    EstadoPedido = "pagado"
	EstadoCancelado EstadoPedido = "cancelado"
)

type DetallePedido struct {
	ProductoID string  `json:"productoId"`
	Nombre     string  `json:"nombre"`
	Cantidad   int     `json:"cantidad"`
	Precio     float64 `json:"precio"`
	Subtotal   float64 `json:"subtotal"`
}

type Pedido struct {
	ID        string          `json:"id"`
	ClienteID string          `json:"clienteId"`
	Estado    EstadoPedido    `json:"estado"`
	Total     float64         `json:"total"`
	Fecha     time.Time       `json:"fecha"`
	Detalles  []DetallePedido `json:"detalles"`
}

func NuevoPedido(id string, carrito *Carrito) (*Pedido, error) {
	if carrito == nil || carrito.EstaVacio() {
		return nil, errors.New("no se puede crear un pedido con carrito vacio")
	}
	detalles := make([]DetallePedido, 0, len(carrito.items))
	for _, item := range carrito.items {
		detalles = append(detalles, DetallePedido{
			ProductoID: item.Producto.ID,
			Nombre:     item.Producto.Nombre,
			Cantidad:   item.Cantidad,
			Precio:     item.Producto.Precio,
			Subtotal:   item.Subtotal,
		})
	}
	return &Pedido{ID: id, ClienteID: carrito.ClienteID, Estado: EstadoPendiente, Total: carrito.Total(), Fecha: time.Now(), Detalles: detalles}, nil
}

func (p *Pedido) CambiarEstado(nuevoEstado EstadoPedido) error {
	switch nuevoEstado {
	case EstadoPendiente, EstadoPagado, EstadoCancelado:
		p.Estado = nuevoEstado
		return nil
	default:
		return errors.New("estado de pedido no valido")
	}
}

// MetodoPago es una interfaz que permite procesar diferentes formas de pago
// sin acoplar la logica de pedido a una implementacion concreta.
type MetodoPago interface {
	Nombre() string
	Procesar(monto float64) error
}

type PagoTarjeta struct{}

func (p PagoTarjeta) Nombre() string { return "tarjeta" }
func (p PagoTarjeta) Procesar(monto float64) error {
	if monto <= 0 {
		return errors.New("el monto de pago no es valido")
	}
	return nil
}

type PagoTransferencia struct{}

func (p PagoTransferencia) Nombre() string { return "transferencia" }
func (p PagoTransferencia) Procesar(monto float64) error {
	if monto <= 0 {
		return errors.New("el monto de pago no es valido")
	}
	return nil
}

type PagoContraEntrega struct{}

func (p PagoContraEntrega) Nombre() string { return "contra_entrega" }
func (p PagoContraEntrega) Procesar(monto float64) error {
	if monto <= 0 {
		return errors.New("el monto de pago no es valido")
	}
	return nil
}

type Pago struct {
	ID       string    `json:"id"`
	PedidoID string    `json:"pedidoId"`
	Metodo   string    `json:"metodo"`
	Monto    float64   `json:"monto"`
	Estado   string    `json:"estado"`
	Fecha    time.Time `json:"fecha"`
}

type ReporteVentas struct {
	TotalPedidos int            `json:"totalPedidos"`
	Ventas       float64        `json:"ventas"`
	PorEstado    map[string]int `json:"porEstado"`
	FechaReporte time.Time      `json:"fechaReporte"`
}

type ReporteInventario struct {
	TotalProductos int           `json:"totalProductos"`
	BajoStock      []ProductoDTO `json:"bajoStock"`
	FechaReporte   time.Time     `json:"fechaReporte"`
}
