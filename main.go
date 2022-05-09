package main

import (
	"fmt"
	"math"
	"math/rand"
)

const NH = 2
const ND = 5
const sueldos = 25000
const hornosCosto = 20000

func main() {
	var T float64 = 0
	var TPLL float64 = 0
	var TF float64 = 72000000
	var tiempoComprometidoDelivery = new([ND]float64)[0:ND]
	var tiempoComprometidoHorno = new([NH]float64)[0:NH]
	var STO = new([NH]float64)[0:NH]
	var STODelivery = new([ND]float64)[0:ND]
	var STEH float64 = 0
	var STED float64 = 0
	var contador float64 = 0
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
			if (T + TAHorno) >= menorTcDelivery {
				STODelivery[j] = STODelivery[j] + ((T + TAHorno) - menorTcDelivery)
				tiempoComprometidoDelivery[j] = T + TAHorno + TA
			} else {
				STED = STED + (menorTcDelivery - (T + TAHorno))
				contador = contador + (menorTcDelivery - (T + TAHorno))
				if (menorTcDelivery - (T + TAHorno)) > 20 {
					cantidadPedidosGratis++
					precioPedidosGratis = precioPedidosGratis + precioPedido
				}
				tiempoComprometidoDelivery[j] = menorTcDelivery + TAHorno + TA
			}
		} else {
			STEH = STEH + (menorTC - T) + TAHorno
			contador = contador + (menorTC - T)
			if (menorTC - T) > 20 {
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

	CalculateResults(STO, STODelivery, STEH, STED, T, cantidadPedidos, precioPedidosGratis, beneficios, contador)

}

func CalculateResults(STO []float64, STODelivery []float64, STEH float64, STED float64, T float64, cantidadPedidos float64, precioPedidosGratis int,
	beneficios int, contador float64) {
	fmt.Println("Porcentaje tiempo ocioso de cada horno:")
	//Porcentaje tiempo ocioso de cada horno
	for i, t := range STO {
		fmt.Println(fmt.Sprintf("PTO Horno %d: %f", i+1, t/T*100))
	}

	fmt.Println("Porcentaje tiempo ocioso de cada delivery:")
	//Porcentaje tiempo ocioso de cada delivery
	for i, t := range STODelivery {
		fmt.Println(fmt.Sprintf("PTO Delivery %d: %f", i+1, t/T*100))
	}

	//Promedio de tiempo de espera de pedido por horno
	fmt.Println("Promedio de tiempo de espera de pedido:")
	fmt.Println(fmt.Sprintf("STEH: %f", STEH/cantidadPedidos))

	//Promedio de tiempo de espera de pedido por delivery
	fmt.Println("Promedio de tiempo de espera de pedido:")
	fmt.Println(fmt.Sprintf("STED: %f", STED/cantidadPedidos))

	//Promedio Pérdida Mensual por pedidos gratis
	fmt.Println(fmt.Sprintf("Promedio Pérdida Mensual por pedidos gratis: %d", precioPedidosGratis/1666))

	//Costo de sueldos y hornos mensual
	fmt.Println(fmt.Sprintf("Costo de sueldos mensuales: %d", ND*sueldos))
	fmt.Println(fmt.Sprintf("Costo de hornos mensuales: %d", NH*hornosCosto))

	//Beneficios Mensuales
	fmt.Println(fmt.Sprintf("Beneficios Mensuales: %d", ((beneficios-precioPedidosGratis)/1666)-ND*sueldos-NH*hornosCosto))

	fmt.Println(fmt.Sprintf("Promedio tiempo espera: %f", contador/cantidadPedidos))
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
	result := (-3.2109 * math.Log(1/x-1)) + 8.14
	if result < 0 {
		return GetIA() / 1.3
	} else {
		return result / 1.3
	}

}

func GetTA() float64 {
	x := rand.Float64()
	result := -3.4753*math.Log(-math.Log(x)) + 9.1582
	if result < 0 {
		return 2 * GetTA()
	} else {
		return 2 * result
	}
}
