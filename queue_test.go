package queue

import (
			"testing"
		)

var queue = Queue{}

func TestSize(t *testing.T) {
	description := "Test queue size function"
	if queue.Size() != 0 {
		t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", description, 0, queue.Size())
	}
}

func TestEnqueue(t *testing.T) {
	description := "Test queue enqueue function"
	queue.Enqueue(10)
	value := queue.Dequeue()

	if queue.Size() == 1 || value != 10 {
		t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", description, 10, value)
	}
}

func TestDequeue(t *testing.T) {
	description := "Test queue dequeue function"
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Enqueue(50)

	value := queue.Dequeue()
	if queue.Size() == 3 || value != 20 {
		t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", description, 50, value)
	}
}

func TestIsEmpty(t *testing.T) {
	description := "Test queue isEmpty function"
	if queue.IsEmpty() {
		t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", description, false, queue.IsEmpty())
	}
}

