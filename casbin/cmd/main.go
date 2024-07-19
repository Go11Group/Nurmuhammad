package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

// Custom data structures for attribute-based request
type Subject struct {
	Name string
	Role string
}

type Object struct {
	Name string
}

type Environment struct {
	Time string
}

func main() {
	// Load the ABAC model and policy from files
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		panic(err)
	}

	// Define test cases
	testCases := []struct {
		sub      Subject
		obj      Object
		act      string
		expected bool
	}{
		{Subject{Name: "Alice", Role: "Doctor"}, Object{Name: "MedicalRecord"}, "read", true},
		{Subject{Name: "Bob", Role: "Nurse"}, Object{Name: "MedicalRecord"}, "read", true},
		{Subject{Name: "Charlie", Role: "Admin"}, Object{Name: "AnyResource"}, "read", true},
		{Subject{Name: "Charlie", Role: "Admin"}, Object{Name: "AnyResource"}, "write", true},
		{Subject{Name: "Alice", Role: "Doctor"}, Object{Name: "MedicalRecord"}, "write", false},
	}

	// Test the cases
	for _, tc := range testCases {
		result, err := e.Enforce(tc.sub, tc.obj, tc.act)
		if err != nil {
			fmt.Printf("Error in enforcement: %v \n", err)
		} else {
			if result == tc.expected {
				fmt.Printf("PASS: %s with role %s can %s at %s: %v \n", tc.sub.Name, tc.sub.Role, tc.act, tc.obj.Name, result)
			} else {
				fmt.Printf("FAIL: %s with role %s can %s at %s: %v \n", tc.sub.Name, tc.sub.Role, tc.act, tc.obj.Name, result)
			}
		}
	}
}
