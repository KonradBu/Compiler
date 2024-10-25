package main

import (
	//"fmt"
	"errors"
	"strconv"
)

type List[T comparable] struct {
	value T
	next  *List[T]
	prev  *List[T]
}

func makeList[T comparable](values []T) List[T] {
	var prev *List[T]
	var head *List[T]
	for i, val := range values {
		node := new(List[T])
		node.value = val
		node.prev = prev
		prev = node
		if node.prev != nil {
			node.prev.next = node
		}
		if i == 0 {
			head = node
		}
	}
	return *head
}

func (node *List[T]) append(val T) {
	var prev *List[T]
	for node != nil {
		prev = node
		node = node.next
	}
	newNode := new(List[T])
	newNode.value = val
	newNode.prev = prev
	prev.next = newNode
}

func (node *List[T]) delete(deleteVal T) (List[T], error) {
	InvalidArgumentError := errors.New("object not found in list")
	head := node
	for node != nil && node.value != deleteVal {
		node = node.next
	}
	if node == nil || node.prev == nil {
		return *head, InvalidArgumentError
	} else {
		if node.prev != nil {
			node.prev.next = node.next
		} else {
			head = node.next
			node.next.prev = nil
		}
		if node.next != nil {
			node.next.prev = node.prev
		}
	}
	return *head, nil
}

func (node *List[T]) toString() (string, error) {
	CantParseError := errors.New("cant parse value")
	retString := ""
	for node != nil {
		var tmp string
		switch val := any(node.value).(type) {
		case int:
			tmp = strconv.Itoa(val)
		case float64:
			tmp = strconv.FormatFloat(val, 'f', -1, 64)
		case bool:
			tmp = strconv.FormatBool(val)
		case string:
			tmp = val
		default:
			return "", CantParseError
		}
		retString = retString + " " + tmp
		node = node.next
	}
	return retString, nil
}