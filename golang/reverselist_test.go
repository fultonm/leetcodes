package main

import (
	"golang/helper"
	"testing"
)

func TestReverseList(t *testing.T) {
	input := &ListNode{
		1, &ListNode{
			2, &ListNode{
				3, &ListNode{
					4, &ListNode{
						5, nil,
					},
				},
			},
		},
	}
	expect := &ListNode{
		5, &ListNode{
			4, &ListNode{
				3, &ListNode{
					2, &ListNode{
						1, nil,
					},
				},
			},
		},
	}
	AssertLinkedList(reverseList(input), expect, t)
}

func AssertLinkedList(result *ListNode, expect *ListNode, t *testing.T) {
	resultCur := result
	expectCur := expect
	for resultCur != nil {
		if expectCur == nil {
			t.Fatalf("Result list is longer than expected list")
		}
		helper.AssertInt(resultCur.Val, expectCur.Val, t)
		resultCur = resultCur.Next
		expectCur = expectCur.Next
	}
	if expectCur != nil {
		t.Fatalf("Expected list is longer than result list")
	}
}
