package cfsm

import "errors"

var (
	// ErrStateUndef is an error when referencing a state that does not exist.
	ErrStateUndef = errors.New("undefined state")
	// ErrStateAlias is an error when reusing a non-free CFSM state in a
	// different CFSM (aliasing).
	ErrStateAlias = errors.New("state is already attached to a CFSM")
)
