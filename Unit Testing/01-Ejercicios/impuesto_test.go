// archivo: impuesto_test.go

package inpuesto_test

import (
	"testing"
)

func TestImpuestoSalarioMenor50k(t *testing.T) {
	salario := 40000.0
	impuesto := impuestoSalario(salario)

	// Verificar que el impuesto sea 0 para salario menor a $50,000
	if impuesto != 0 {
		t.Errorf("Para salario de $%v, el impuesto debería ser $0, pero es $%v", salario, impuesto)
	}
}

func TestImpuestoSalarioEntre50kY150k(t *testing.T) {
	salario := 80000.0
	impuesto := impuestoSalario(salario)

	// Verificar que el impuesto sea calculado correctamente para salario entre $50,000 y $150,000
	impuestoEsperado := salario * 0.17
	if impuesto != impuestoEsperado {
		t.Errorf("Para salario de $%v, el impuesto debería ser $%v, pero es $%v", salario, impuestoEsperado, impuesto)
	}
}

func TestImpuestoSalarioMayor150k(t *testing.T) {
	salario := 160000.0
	impuesto := impuestoSalario(salario)

	// Verificar que el impuesto sea calculado correctamente para salario mayor a $150,000
	impuestoEsperado := salario * 0.27
	if impuesto != impuestoEsperado {
		t.Errorf("Para salario de $%v, el impuesto debería ser $%v, pero es $%v", salario, impuestoEsperado, impuesto)
	}
}
