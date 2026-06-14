package ecommerce

// ResumenReporte concentra indicadores administrativos del e-commerce.
type ResumenReporte struct {
	TotalVentas          float64
	PedidosPorEstado     map[EstadoPedido]int
	ProductosMasVendidos map[string]int
	ProductosBajoStock   []Producto
}

type ReporteService struct {
	inventario *Inventario
}

func NewReporteService(inventario *Inventario) *ReporteService {
	return &ReporteService{inventario: inventario}
}

// GenerarResumen usa goroutines y canales para calcular partes del reporte en paralelo.
// Se comenta porque es una funcionalidad mas compleja: separa ventas, estados y productos vendidos
// para que cada tarea se resuelva sin bloquear a las demas.
func (r ReporteService) GenerarResumen(pedidos []Pedido) ResumenReporte {
	ventasCh := make(chan float64)
	estadosCh := make(chan map[EstadoPedido]int)
	vendidosCh := make(chan map[string]int)

	go func() {
		total := 0.0
		for _, pedido := range pedidos {
			if pedido.Estado() == EstadoPagado || pedido.Estado() == EstadoEnviado {
				total += pedido.Total()
			}
		}
		ventasCh <- total
	}()

	go func() {
		conteo := make(map[EstadoPedido]int)
		for _, pedido := range pedidos {
			conteo[pedido.Estado()]++
		}
		estadosCh <- conteo
	}()

	go func() {
		vendidos := make(map[string]int)
		for _, pedido := range pedidos {
			for _, detalle := range pedido.Detalles() {
				vendidos[detalle.Nombre()] += detalle.Cantidad()
			}
		}
		vendidosCh <- vendidos
	}()

	return ResumenReporte{
		TotalVentas:          <-ventasCh,
		PedidosPorEstado:     <-estadosCh,
		ProductosMasVendidos: <-vendidosCh,
		ProductosBajoStock:   r.inventario.ProductosBajoStock(3),
	}
}
