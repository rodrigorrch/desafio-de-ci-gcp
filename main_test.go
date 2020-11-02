package main

import "testing"

func TestSum(t *testing.T){
	total := sum(5,5)
	if total != 10 {
		t.Errorf("Soma está incorreto. O resultado deveria ser: %d, e o valor retornado foi: %d", 10, total);
	}
}