package ecommerce

import "errors"

var (
	ErrCampoVacio         = errors.New("el campo no puede estar vacio")
	ErrCorreoInvalido     = errors.New("el correo no tiene un formato valido")
	ErrContrasenaInvalida = errors.New("la contrasena debe tener al menos 6 caracteres")
	ErrPrecioInvalido     = errors.New("el precio debe ser mayor a cero")
	ErrStockInvalido      = errors.New("el stock no puede ser negativo")
	ErrProductoNoExiste   = errors.New("el producto no existe en el inventario")
	ErrStockInsuficiente  = errors.New("stock insuficiente para completar la operacion")
	ErrCantidadInvalida   = errors.New("la cantidad debe ser mayor a cero")
	ErrCarritoVacio       = errors.New("el carrito esta vacio")
	ErrPagoInvalido       = errors.New("el metodo de pago no pudo procesarse")
	ErrEstadoInvalido     = errors.New("el estado del pedido no es valido")
)
