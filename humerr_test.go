package humerr_test

import (
	"errors"
	"testing"

	"github.com/christgf/humerr"
)

func TestErrorError(t *testing.T) {
	var e humerr.Error

	e = humerr.Error{No: humerr.ETOOLAZY, Err: errors.New("too cumbersome")}
	if exp, got := "(101) too cumbersome", e.Error(); exp != got {
		t.Fatalf("expected %q but got %q", exp, got)
	}

	e = humerr.Error{No: humerr.EBORING}
	if exp, got := "(102) 🤷", e.Error(); exp != got {
		t.Fatalf("expected %q but got %q", exp, got)
	}
}

func TestErrorIs(t *testing.T) {
	e := &humerr.Error{No: humerr.ETOOLAZY}

	tt := []struct {
		err error
		exp bool
	}{
		{&humerr.Error{No: humerr.ETOOLAZY}, true},
		{&humerr.Error{No: humerr.ETOOLAZY, Err: errors.New("too cumbersome")}, true},
		{&humerr.Error{No: humerr.EINCONVENIENT}, true},
		{&humerr.Error{No: humerr.EBORING}, false},
		{errors.New("(101) operation too cumbersome"), false},
	}

	for _, tc := range tt {
		if exp, got := tc.exp, errors.Is(e, tc.err); exp != got {
			t.Errorf("errors.Is(%+v, %+v): expected %v", e, tc.err, exp)
		}
	}
}
