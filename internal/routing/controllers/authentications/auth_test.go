package authentications_test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/diazharizky/use-case-go-jwt-auth-service/internal/app"
	"github.com/diazharizky/use-case-go-jwt-auth-service/pkg/httptesthelpers"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Auth controller test", func() {
	type want struct {
		status int
		resp   map[string]interface{}
	}

	type testCase struct {
		want

		name       string
		assertions func(want)
	}

	contexts := []struct {
		name      string
		testCases []testCase
	}{
		{
			name: "Failed to make request",
			testCases: []testCase{
				{
					name: "When basic authorization is not attached on the headers",
					want: want{
						status: http.StatusBadRequest,
						resp: map[string]interface{}{
							"data": "Unable to parse username/password from the request headers",
						},
					},
					assertions: func(want want) {
						Describe("", func() {
							It("Should return status code of 400 and `Unable to parse username/password from the request headers` error message", func() {
								gotStatus, gotBody, err := httptesthelpers.SendRequest(http.MethodPost, "/", nil, testRouter, map[string]string{})
								if err != nil {
									panic(
										fmt.Errorf("Error unable to make request: %v", err),
									)
								}

								Expect(gotStatus).Should(Equal(want.status))

								var gotResp map[string]interface{}
								if err = json.Unmarshal(gotBody.Bytes(), &gotResp); err != nil {
									panic(
										fmt.Errorf("Error unable to parse response body: %v", err),
									)
								}

								Expect(gotResp).Should(Equal(want.resp))
							})

						})
					},
				},
				{
					name: "When username and password are empty",
					want: want{
						status: http.StatusBadRequest,
						resp: map[string]interface{}{
							"data": "Username/password combination is incorrect",
						},
					},
					assertions: func(want want) {
						It("Should return status code of 400 and `Username/password combination is incorrect` error message", func() {
							gotStatus, gotBody, err := httptesthelpers.SendRequest(http.MethodPost, "/", nil, testRouter, map[string]string{
								"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(":"))),
							})

							if err != nil {
								panic(
									fmt.Errorf("Error unable to make request: %v", err),
								)
							}

							Expect(gotStatus).Should(Equal(want.status))

							var gotResp map[string]interface{}
							if err = json.Unmarshal(gotBody.Bytes(), &gotResp); err != nil {
								panic(
									fmt.Errorf("Error unable to parse response body: %v", err),
								)
							}

							Expect(gotResp).Should(Equal(want.resp))
						})
					},
				},
			},
		},
	}

	for _, ctx := range contexts {
		Context(ctx.name, func() {
			for _, tc := range ctx.testCases {
				BeforeEach(func() {
					mockCtrl := gomock.NewController(GinkgoT())
					mockGenerateJWTTokenService := app.NewMockIGenerateJWTTokenService(mockCtrl)

					appCtx = &app.Context{
						GenerateJWTTokenService: mockGenerateJWTTokenService,
					}
				})

				tc.assertions(tc.want)
			}
		})
	}
})
