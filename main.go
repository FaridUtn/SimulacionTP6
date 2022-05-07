package main

import (
	"math"
	"math/rand"
)

func main() {

	var T float64 = 0
	var TPLL float64 = 0
	var TF float64 = 14400000
	var NH int = 1
	var ND int = 1
	var precioDocenaEmpanada int = 760
	var tiempoDocenaEmpanada int = 12
	var precioDocenaEmpanadaSurtida int = 830
	var tiempoDocenaEmpanadaSurtida int = 12
	var precioMediaDocenaEmpanada int = 400
	var tiempoMediaDocenaEmpanada int = 6
	var precioMediaDocenaEmpanadaSurtida = 440
	var tiempoMediaDocenaEmpanadaSurtida = 6
	var preciodocenaYMediaEmpanada int = 1080
	var tiempodocenaYMediaEmpanada int = 18
	var preciodocenaYMediaEmpanadaSurtida int = 1170
	var tiempodocenaYMediaEmpanadaSurtida int = 18
	var tiempoComprometidoDelivery = make([]float64, ND)
	var tiempoComprometidoHorno = make([]float64, NH)

	T = TPLL
	IA := GetIA()
	TPLL = T + IA

	categoria := GetCategoria()
	TAHorno := GetTAHorno(categoria)

	menorTC, i := GetMenorTC(tiempoComprometidoHorno)
	if T >= menorTC {
		tiempoComprometidoHorno[i] = T + TAHorno
	} else {
		tiempoComprometidoHorno[i] = menorTC + TAHorno
	}

	if T >= TF {
		return
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

func GetMenorTC(tiempoComprometidoHorno []float64) (float64, int) {
	if len(tiempoComprometidoHorno) < 2 {
		return tiempoComprometidoHorno[0], 0
	}
	var time = tiempoComprometidoHorno[0]
	var index = 0
	for i, t := range tiempoComprometidoHorno {
		if t < time {
			time = t
			index = i
		}
	}

	return time, index
}

func GetIA(x float64) float64 {
	return (-3.2109 * math.Log(1/x-1)) + 8.14

}

func GetTA(x float64) float64 {
	return -3.4753*math.Log(-math.Log(x)) + 9.1582
}
