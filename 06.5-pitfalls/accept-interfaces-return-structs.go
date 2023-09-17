package main

import (
	"fmt"
)

// bad design

// package event_emitter

type BadEventEmitterIface interface {
	Emit(event string)
}

type BadEventEmitter struct{}

func NewBadEventEmitter() BadEventEmitterIface {
	return &BadEventEmitter{}
}

func (e BadEventEmitter) Emit(event string) {
	fmt.Println("emit event", event)
}

// package handler

func BadHandler(e BadEventEmitterIface) {
	e.Emit("login")
}

// good design
// package event_emitter

type EventEmitter struct{}

func NewGoodEventEmitter() EventEmitter {
	return EventEmitter{}
}

func (e EventEmitter) Emit(event string) {
	fmt.Println("emit event", event)
}

// package handler

type GoodEventEmitterIface interface {
	Emit(event string)
}

func GoodHandler(e GoodEventEmitterIface) {
	e.Emit("login")
}

func acceptInterfacesReturnStructs() {
	GoodHandler(NewGoodEventEmitter()) // emit event login
}
