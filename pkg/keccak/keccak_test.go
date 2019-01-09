package keccak

import (
	"bytes"
	"testing"
)

func TestKeccak(t *testing.T) {
	type data struct {
		d int
		M []byte
	}
	inputs := []data{
		data{0, make([]byte, 0)},
	}
	expected := [][]byte{
		make([]byte, 0),
	}
	for i, input := range inputs {
		hash, err := Keccak(input.d, input.M)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(hash, expected[i]) {
			t.Errorf("[TestKeccak] Case %d failed: result '%x', but expected '%x'\n", i, hash, expected[i])
		}
	}
}

func TestTheta(t *testing.T) {
	inputs := []State{
		make(State, 0),
	}
	expected := []State{
		make(State, 0),
	}
	for i, input := range inputs {
		output, err := Theta(input)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(output, expected[i]) {
			t.Errorf("[TestTheta] Case %d failed: result '%x', but expected '%x'\n", i, output, expected[i])
		}
	}
}

func TestRho(t *testing.T) {
	inputs := []State{
		make(State, 0),
	}
	expected := []State{
		make(State, 0),
	}
	for i, input := range inputs {
		output, err := Rho(input)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(output, expected[i]) {
			t.Errorf("[TestRho] Case %d failed: result '%x', but expected '%x'\n", i, output, expected[i])
		}
	}
}

func TestPi(t *testing.T) {
	inputs := []State{
		make(State, 0),
	}
	expected := []State{
		make(State, 0),
	}
	for i, input := range inputs {
		output, err := Pi(input)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(output, expected[i]) {
			t.Errorf("[TestPi] Case %d failed: result '%x', but expected '%x'\n", i, output, expected[i])
		}
	}
}

func TestChi(t *testing.T) {
	inputs := []State{
		make(State, 0),
	}
	expected := []State{
		make(State, 0),
	}
	for i, input := range inputs {
		output, err := Chi(input)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(output, expected[i]) {
			t.Errorf("[TestChi] Case %d failed: result '%x', but expected '%x'\n", i, output, expected[i])
		}
	}
}

func TestIota(t *testing.T) {
	inputs := []State{
		make(State, 0),
	}
	expected := []State{
		make(State, 0),
	}
	for i, input := range inputs {
		output, err := Iota(input)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(output, expected[i]) {
			t.Errorf("[TestIota] Case %d failed: result '%x', but expected '%x'\n", i, output, expected[i])
		}
	}
}

// Rc calculates a round constant.it use in Iota.
func TestRc(t *testing.T) {
	inputs := []int{
		0,
	}
	expected := []byte{
		0,
	}
	for i, input := range inputs {
		output, err := Rc(input)
		if err != nil {
			t.Error(err)
		}
		if output != expected[i] {
			t.Errorf("[TestRc] Case %d failed: result '%x', but expected '%x'\n", i, output, expected[i])
		}
	}
}

// Pad makes padding represented by a regular expression of 10*1.
func TestPad(t *testing.T) {
	type data struct {
		x int
		m int
	}
	inputs := []data{
		data{0, 0},
	}
	expected := [][]byte{
		make([]byte, 0),
	}
	for i, input := range inputs {
		output, err := Pad(input.x, input.m)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(output, expected[i]) {
			t.Errorf("[TestPad] Case %d failed: result '%x', but expected '%x'\n", i, output, expected[i])
		}
	}
}
