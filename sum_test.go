package main

import "testing"

func TestSum(t *testing.T) {

	result := sum(2, 2)
	if result != 4 {
		t.Errorf("Resultado da soma é inválido: Resuldado %d. Esperado %d.", result, 4);
	}
}

func TestSub(t *testing.T) {
	
	result := sub(10, 5)
	if result != 5 {
		t.Errorf("Resultado da subtração é inválido: Resuldado %d. Esperado %d.", result, 5);
	}
}

func TestTiles(t *testing.T) {
	
	result := tiles(3, 3)
	if result != 9 {
		t.Errorf("Resultado da subtração é inválido: Resuldado %d. Esperado %d.", result, 9);
	}
}

func TestSumx(t *testing.T) {
	
	result := sumX(10, 5)
	if result != 25 {
		t.Errorf("Resultado da subtração é inválido: Resuldado %d. Esperado %d.", result, 25);
	}
}