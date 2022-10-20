package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/dsolerh/examples/concurrency.course/finalProject/pkg/data"
)

func Test_Pages(t *testing.T) {
	testCases := []struct {
		desc           string
		url            string
		statusExpected int
		handler        http.HandlerFunc
		sessionData    map[string]any
		expectedHTML   string
	}{
		{
			desc:           "Home page test",
			url:            "/",
			statusExpected: http.StatusOK,
			handler:        testApp.HomePage,
		},
		{
			desc:           "Login page test",
			url:            "/login",
			statusExpected: http.StatusOK,
			handler:        testApp.LoginPage,
		},
		{
			desc:           "Register page test",
			url:            "/register",
			statusExpected: http.StatusOK,
			handler:        testApp.RegisterPage,
		},
		{
			desc:           "Logout page test",
			url:            "/logout",
			statusExpected: http.StatusSeeOther,
			handler:        testApp.Logout,
			sessionData: map[string]any{
				"userID": 1,
				"user":   data.User{},
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			rr := httptest.NewRecorder()

			req, _ := http.NewRequest("GET", tC.url, nil)
			ctx := getCtx(req)
			req = req.WithContext(ctx)

			if len(tC.sessionData) > 0 {
				for key, value := range tC.sessionData {
					testApp.Session.Put(ctx, key, value)
				}
			}

			tC.handler.ServeHTTP(rr, req)

			if rr.Code != tC.statusExpected {
				t.Errorf("Failed: expected %d, but got %d", tC.statusExpected, rr.Code)
			}

			if len(tC.expectedHTML) > 0 {
				html := rr.Body.String()
				if !strings.Contains(html, tC.expectedHTML) {
					t.Errorf("Failed: expected to find %s, but did not", tC.expectedHTML)
				}
			}
		})
	}
}

func Test_Config_Login(t *testing.T) {

	postedData := url.Values{
		"email":    {"admin@example.com"},
		"password": {"abc123"},
	}

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/login", strings.NewReader(postedData.Encode()))
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	handler := http.HandlerFunc(testApp.Login)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Errorf("Failed: expected %d, but got %d", http.StatusSeeOther, rr.Code)
	}

	if !testApp.Session.Exists(ctx, "userID") {
		t.Error("did not find userID in session")
	}
}

func Test_Config_SubscribeToPlan(t *testing.T) {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/subscribe?id=1", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	testApp.Session.Put(ctx, "user", data.User{
		ID:        1,
		Email:     "admin@example.com",
		FirstName: "Admin",
		LastName:  "Admin",
		Active:    1,
	})

	handler := http.HandlerFunc(testApp.Login)

	handler.ServeHTTP(rr, req)
	testApp.Wait.Wait()

	if rr.Code != http.StatusSeeOther {
		t.Errorf("Failed: expected %d, but got %d", http.StatusSeeOther, rr.Code)
	}

}
