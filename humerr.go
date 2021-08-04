// Ported from code presented in "Root Cause: Human Errno" by Jan Schaumann. See
// https://www.netmeister.org/blog/humanerrno.html.
//
// This code is in the public domain.

// Package humerr defines the root cause of all errors.
package humerr

import (
	"fmt"
)

// Errno is a type of numeric error code.
type Errno uint8

// Errno codes.
const (
	ETOOLAZY        Errno = iota + 101 /* Operation too cumbersome */
	EBORING                            /* Operation timed out */
	ETLDR                              /* Information overload */
	EFORGOT                            /* Cache miss */
	ENOCLUE                            /* Persistent lookup failure */
	EBIKESHED                          /* Endless loop */
	EGROUPTHINK                        /* Erroneous consensus */
	ENOTIT                             /* Indirect assignment */
	EIKEA                              /* Can't delete my own work */
	ENIH                               /* Import failure */
	ELMGTFY                            /* Indirect cache miss */
	ESNAFU                             /* Unexpected status quo */
	ESEP                               /* Insufficient pay grade or deferred responsibility */
	ENOTENOUGHPI                       /* Insufficient amount of pie */
	ENOCOFFEE                          /* Energy starvation */
	EPEBKAC                            /* User error */
	EWUTANG                            /* Ruckus applied, unfuckwithably */
	ECOMPILING                         /* Interrupted distraction */
	EGITHUB                            /* Distributed SPOF */
	EGITPUB                            /* Information disclosure */
	EMEETING                           /* Inefficient interrupt */
	EMAIL                              /* Inefficient persistent interrupt */
	ETYOP                              /* Input error */
	EBOFH                              /* Access denied */
	EYOLO                              /* Failsafe disabled */
	EASLEEP                            /* Process temporarily suspended */
	EIANAL                             /* Unfunded speculation */
	EPLUSONE                           /* Pointlessly pending review */
	ESLOPPY                            /* Forced completion */
	EWELLACTUALLY                      /* Superflous commentary */
	EPERL                              /* Incorrect number of $@%{ */
	EJAVA                              /* Errno exception factory */
	EPYTHONG                           /* Universal Freudian slip */
	EPHP                               /* Fubar */
	ERMFR                              /* Self-inflicted DoS */
	EWIKI                              /* Information retrieval impossible */
	EBROOKS                            /* Too many developers */
	EBOBBYTABLES                       /* Command injection */
	EUNICORN                           /* Inflated self-evaluation */
	ESTACKOVERFLOWN                    /* Misunderstood copy and paste */
	ESLACKED                           /* Information disclosure as a Service */
	EDGAF                              /* Mental resource exhaustion */
	EGOLANG                            /* But no generics */

	EINCONVENIENT = ETOOLAZY /* Defined for portability */
)

// Error represents the ultimate type of error.
type Error struct {
	// No is the numeric error code. Expected to be ENOCLUE if unknown.
	No Errno
	// The underlying error that triggered this one, if any.
	Err error
}

// Error implements the error interface.
func (e *Error) Error() string {
	msg := "🤷"
	if e.Err != nil {
		msg = e.Err.Error()
	}
	return fmt.Sprintf("(%d) %s", e.No, msg)
}

// Unwrap the underlying error when using the %w verb.
func (e *Error) Unwrap() error {
	return e.Err
}

// Is allows to check if an error is of a specific type by evaluating its code.
//
//	if errors.Is(err, &Error{No: EBORING}) {
//		// Do something if err is boring.
//	}
func (e *Error) Is(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}
	return e.No == t.No
}
