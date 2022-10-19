package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Config_AddDefaultData(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	testApp.Session.Put(ctx, "flash", "flash")
	testApp.Session.Put(ctx, "warning", "warning")
	testApp.Session.Put(ctx, "error", "error")

	td := testApp.AddDefaultData(&TemplateData{}, req)

	if td.Flash != "flash" {
		t.Error("failed to get the flash data")
	}

	if td.Warning != "warning" {
		t.Error("failed to get the warning data")
	}

	if td.Error != "error" {
		t.Error("failed to get the error data")
	}
}

func Test_Config_IsAuthenticated(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	auth := testApp.IsAuthenticated(req)

	if auth {
		t.Error("returns true for authenticated and should return false")
	}

	testApp.Session.Put(ctx, "userID", 1)

	auth = testApp.IsAuthenticated(req)

	if !auth {
		t.Error("returns false for authenticated and should return true")
	}
}

func Test_Config_render(t *testing.T) {
	pathToTemplates = "./templates"

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	testApp.render(rr, req, "home.page.gohtml", &TemplateData{})

	if rr.Code != http.StatusOK {
		t.Error("failed to render page")
	}
}
