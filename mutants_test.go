package mutants

import "testing"
import "github.com/ceoln/expressions"

func TestAccurateCopy(t *testing.T) {
	var m = map[string]expressions.Float{"foo": 7, "bar": 10.5, "baz": 3.14159}
	m1 := NewVariableRef("foo")
	m2 := NewConstant(200)
	m3 := NewConstant(5)
	m4 := NewBinaryOperation('*', m3, m2)
	m5 := NewBinaryOperation('+', m4, m1)
	mut, okay := m5.RoughCopy(1.0, m)
	if !okay {
		t.Errorf("%v copy result was not okay", m5)
	}
	if !m5.Equal(mut) {
		t.Errorf("%v accurate-copy result was %v, want %v", m5, mut, m5)
	}
}

func TestInaccurateCopy(t *testing.T) {
	var m = map[string]expressions.Float{"foo": 7, "bar": 10.5, "baz": 3.14159}
	m1 := NewVariableRef("foo")
	m2 := NewConstant(200)
	m3 := NewConstant(5)
	m4 := NewBinaryOperation('*', m3, m2)
	m5 := NewBinaryOperation('+', m4, m1)
	mut, okay := m5.RoughCopy(0.0, m)
	if !okay {
		t.Errorf("%v copy result was not okay", m5)
	}
	if m5.Equal(mut) {
		t.Errorf("%v inaccurate-copy result was %v, want anything but %v", m5, mut, m5)
	}
}
