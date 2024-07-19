package middleware

import (
	"fmt"
	"net/http"
	"new/api/auth"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func Check(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")

	if accessToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization is required",
		})
		return
	}

	_, err := auth.ValidateAccessToken(accessToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token provided",
		})
		return
	}

	c.Next()
}

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

func Role(c *gin.Context) {
	e, err := casbin.NewEnforcer("api/casbin/model.conf", "api/casbin/policy.csv")
	if err != nil {
		panic(err)
	}

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

	for _, tc := range testCases {
		result, err := e.Enforce(tc.sub, tc.obj, tc.act)
		if err != nil {
			fmt.Printf("Error in enforcement: %v \n", err)
		} else {
			if result == tc.expected {
				fmt.Printf("PASS: %s with role %s can %s at %s: %v \n", tc.sub.Name, tc.sub.Role, tc.act, tc.obj.Name, result)
			} else {
				fmt.Printf("FAIL: %s with role %s can %s at %s: %v \n", tc.sub.Name, tc.sub.Role, tc.act, tc.obj.Name, result)
				return
			}
		}
	}

	c.Next()
}
