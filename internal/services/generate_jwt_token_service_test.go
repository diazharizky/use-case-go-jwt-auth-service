package services

import (
	"fmt"
	"testing"

	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGenerateJWTTokenServiceTestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GenerateJWTTokenService")
}

var _ = Describe("Test generate JWT token service", func() {
	service := NewGenerateJWTTokenService()

	type testCase struct {
		name       string
		payload    models.UserData
		assertions func(tokenString *string, err error)
	}

	contexts := []struct {
		name      string
		testCases []testCase
	}{
		{
			name: "A context",
			testCases: []testCase{
				{
					name: "Must be succeed",
					payload: models.UserData{
						Username: "foo",
					},
					assertions: func(tokenString *string, err error) {
						if err != nil {
							fmt.Println("err", err)
						}

						Expect(*tokenString).ShouldNot(Equal(""))
					},
				},
			},
		},
	}

	for _, ctx := range contexts {
		Context(ctx.name, func() {
			for _, tc := range ctx.testCases {
				It(tc.name, func() {
					tokenString, err := service.Call(tc.payload)
					tc.assertions(tokenString, err)
				})
			}
		})
	}
})
