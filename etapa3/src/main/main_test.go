package main

import "testing"

func TestGreeting(t *testing.T) {
	resultado := greeting("mensagem")
	esperado := "<b>mensagem</b>"

	if resultado != esperado {
		t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
	}
}
