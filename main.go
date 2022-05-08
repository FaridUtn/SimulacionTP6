package main

import (
	"fmt"
	"math"
	"math/rand"
)

const NH = 1
const ND = 1
const sueldos = 25000
const hornosCosto = 20000

func main() {
	var T float64 = 0
	var TPLL float64 = 0
	var TF float64 = 14400000
	var tiempoComprometidoDelivery = new([ND]float64)[0:ND]
	var tiempoComprometidoHorno = new([NH]float64)[0:NH]
	var STO = new([NH]float64)[0:NH]
	var STODelivery = new([NH]float64)[0:NH]
	var STE = new([NH]float64)[0:NH]
	var cantidadPedidos float64 = 0
	var cantidadPedidosGratis int = 0
	var precioPedidosGratis int = 0
	var beneficios int = 0

	for T <= TF {
		T = TPLL
		IA := GetIA()
		TPLL = T + IA

		TA := GetTA()
		cantidadPedidos++

		categoria := GetCategoria()
		TAHorno := GetTAHorno(categoria)
		precioPedido := GetPrecio(categoria)
		beneficios = beneficios + precioPedido
		menorTC, i := GetMenorTC(tiempoComprometidoHorno)
		menorTcDelivery, j := GetMenorTC(tiempoComprometidoDelivery)
		if T >= menorTC {
			STO[i] = STO[i] + (T - menorTC)
			tiempoComprometidoHorno[i] = T + TAHorno
			if T >= menorTcDelivery {
				STODelivery[j] = STODelivery[j] + (T - menorTcDelivery)
				tiempoComprometidoDelivery[j] = T + TAHorno + TA
			} else {
				tiempoComprometidoDelivery[j] = menorTcDelivery + TAHorno + TA
			}
		} else {
			STE[i] = STE[i] + (menorTC - T) + TAHorno
			if (menorTC - T) > 15 {
				cantidadPedidosGratis++
				precioPedidosGratis = precioPedidosGratis + precioPedido
			}
			tiempoComprometidoHorno[i] = menorTC + TAHorno
			if menorTcDelivery <= menorTC+TAHorno {
				tiempoComprometidoDelivery[j] = menorTC + TAHorno + TA
			} else {
				tiempoComprometidoDelivery[j] = menorTcDelivery + TA
			}
		}
	}

	CalculateResults(STO, STODelivery, T, cantidadPedidos, precioPedidosGratis, beneficios)

}

func CalculateResults(STO []float64, STODelivery []float64, T float64, cantidadPedidos float64, precioPedidosGratis int,
	beneficios int) {
	fmt.Println("Porcentaje tiempo ocioso de cada horno:")
	//Porcentaje tiempo ocioso de cada horno
	for i, t := range STO {
		fmt.Println("PTO Horno %+v: %+v", i+1, t/T)
	}

	fmt.Println("Porcentaje tiempo ocioso de cada delivery:")
	//Porcentaje tiempo ocioso de cada delivery
	for i, t := range STODelivery {
		fmt.Println("PTO Delivery %+v: %+v", i+1, t/T)
	}

	//Promedio de tiempo de espera de pedido
	fmt.Println("Promedio de tiempo de espera de pedido:")
	for i, t := range STODelivery {
		fmt.Println("PTO Delivery %+v: %+v", i+1, t/cantidadPedidos)
	}

	//Promedio Pérdida Mensual por pedidos gratis
	fmt.Println("Promedio Pérdida Mensual por pedidos gratis: %+v", precioPedidosGratis/333)

	//Costo de sueldos y hornos mensual
	fmt.Println("Costo de sueldos mensuales: %+v", ND*sueldos)
	fmt.Println("Costo de hornos mensuales: %+v", NH*hornosCosto)

	//Beneficios Mensuales
	fmt.Println("Beneficios Mensuales: %+v", (beneficios-precioPedidosGratis)/333)
}

func GetPrecio(categoria int) int {
	switch categoria {
	case 1:
		return 760
	case 2:
		return 830
	case 3:
		return 400
	case 4:
		return 440
	case 5:
		return 1080
	case 6:
		return 1170
	default:
		return 0
	}
}

func GetTAHorno(categoria int) float64 {
	switch categoria {
	case 1:
		return 12
	case 2:
		return 12
	case 3:
		return 6
	case 4:
		return 6
	case 5:
		return 18
	case 6:
		return 18
	default:
		return 0
	}
}

func GetCategoria() int {
	var r = rand.Float64()
	if r >= 0.52 {
		return 2
	} else if r >= 0.21 {
		return 1
	} else if r >= 0.07 {
		return 4
	} else if r >= 0.02 {
		return 3
	} else if r >= 0.006 {
		return 6
	} else {
		return 5
	}

}

func GetMenorTC(tc []float64) (float64, int) {
	if len(tc) < 2 {
		return tc[0], 0
	}
	var time = tc[0]
	var index = 0
	for i, t := range tc {
		if t < time {
			time = t
			index = i
		}
	}

	return time, index
}

func GetIA() float64 {
	x := rand.Float64()
	return (-3.2109 * math.Log(1/x-1)) + 8.14

}

func GetTA() float64 {
	x := rand.Float64()
	return -3.4753*math.Log(-math.Log(x)) + 9.1582
}
