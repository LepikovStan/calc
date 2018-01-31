package stack

import (
	"log"
	"testing"
)

func TestStack_Push(t *testing.T) {
	stk := New()
	stk.Push("2")

	if stk.Head() != "2" {
		log.Fatalf("Expected 2, result:%s", stk.Head())
	}
}

func TestStack_Head(t *testing.T) {
	stk := New()
	stk.Push("2")

	if stk.Head() != "2" {
		log.Fatalf("Expected 2, result:%s", stk.Head())
	}
}

func TestStack_Pop(t *testing.T) {
	stk := New()
	stk.Push("2")
	stk.Push("4")
	stk.Pop()

	if stk.Head() != "2" {
		log.Fatalf("Expected 2, result:%s", stk.Head())
	}
}
