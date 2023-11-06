package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"net/http/httptest"
	"test-ozon/internal/service"
	"test-ozon/internal/service/mocks"
	"test-ozon/internal/service/response"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandlers_SaveUrl(t *testing.T) {
	type mockBehavior func(m *mocks.Api, url string)

	tests := []struct {
		name                 string
		inputBody            string
		request              Request
		mockBehavior         mockBehavior
		alias                string
		expectedError        error
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"url":"https://google.com"}`,
			request: Request{
				Url: "https://google.com",
			},

			expectedStatusCode:   200,
			expectedResponseBody: `{"status":"OK"}`,
		},
		{
			name:                 "Empty fields",
			inputBody:            `{}`,
			request:              Request{},
			expectedStatusCode:   400,
			expectedError:        errors.New("field Url is a required field"),
			expectedResponseBody: `{"status":"Error","error":"field Url is a required field"}`,
		},
		{
			name:      "Empty fields",
			inputBody: `{"urd":"abcd"}`,
			request: Request{
				Url: "",
			},
			expectedStatusCode:   400,
			expectedError:        errors.New("field Url is a required field"),
			expectedResponseBody: `{"status":"Error","error":"field Url is a required field"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Api := mocks.NewApi(t)

			if tt.expectedStatusCode == 200 {
				Api.On("SaveUrl", tt.request.Url).Return(tt.alias, tt.expectedError).Once()
			}

			s := &service.Service{
				Api: Api,
			}
			h := NewHandlers(s)

			r := gin.Default()
			r.POST("/", h.SaveUrl)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/", bytes.NewBufferString(tt.inputBody))

			r.ServeHTTP(w, req)

			// assert.Equal(t, tt.expectedStatusCode, w.Result())
			// assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.Equal(t, tt.expectedResponseBody, w.Body.String())
		})
	}

}

func TestHandler_GetUrl(t *testing.T) {
	type args struct {
		url      string
		response response.Response
	}

	tests := []struct {
		name                 string
		inputBody            string
		alias                string
		expectedResponseBody string
		response             args
		expectedStatusCode   int
		expectedError        error
	}{
		{
			name:                 "OK",
			alias:                "google",
			expectedResponseBody: `{"url":"https://google.com","status":"OK"}`,
			response: args{
				url: "https://google.com",
				response: response.Response{
					Status: "OK",
				},
			},
			expectedStatusCode: 200,
		},
		{
			name:                 "404",
			alias:                "",
			expectedResponseBody: "404 page not found",
			expectedStatusCode:   404,
		},
		{
			name:  "no results",
			alias: "0123456789A",
			response: args{
				response: response.Response{
					Status: "Error",
					Error:  "no results",
				},
			},
			expectedResponseBody: `{"status":"Error","error":"no results"}`,
			expectedStatusCode:   400,
			expectedError:        errors.New("no rezsults"),
		},
	}

	for _, tt := range tests {
		Api := mocks.NewApi(t)

		if tt.expectedStatusCode != 404 {
			Api.On("GetUrl", tt.alias).Return(tt.response.url, tt.expectedError).Once()
		}

		s := &service.Service{
			Api: Api,
		}

		h := NewHandlers(s)

		r := gin.Default()
		endPoint := fmt.Sprintf("/%s", tt.alias)

		r.GET("/:alias", h.GetUrl)

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", endPoint, bytes.NewBufferString(tt.inputBody))

		r.ServeHTTP(w, req)

		assert.Equal(t, tt.expectedResponseBody, w.Body.String())
	}
}
