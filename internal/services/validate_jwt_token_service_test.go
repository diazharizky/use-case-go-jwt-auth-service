package services_test

import (
	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/models"
	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/services"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test validate JWT token service", func() {
	service := services.NewValidateJWTTokenService()

	type testCase struct {
		name        string
		tokenString string
		assertions  func(isValid bool, claims *models.JWTClaims, err error)
	}

	contexts := []struct {
		name      string
		testCases []testCase
	}{
		{
			name: "When failed to validate",
			testCases: []testCase{
				{
					name:        "Invalid token string",
					tokenString: "",
					assertions: func(isValid bool, _ *models.JWTClaims, _ error) {
						Expect(isValid).Should(Equal(false))
					},
				},
			},
		},
	}

	for _, ctx := range contexts {
		Context(ctx.name, func() {
			for _, tc := range ctx.testCases {
				It(tc.name, func() {
					isValid, claims, err := service.Call(tc.tokenString)
					tc.assertions(isValid, claims, err)
				})
			}
		})
	}
})
