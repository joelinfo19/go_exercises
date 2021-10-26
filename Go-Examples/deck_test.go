package main

import "testing"

func TestNewDeck(t *testing.T){
	d:=newDeck()
	v:=12
	if len(d) != v{
		t.Errorf("Expected deck length of %v, but got %v",len(d),v)
	}
}