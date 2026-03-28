package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "go_service/docs"

	"github.com/stretchr/testify/assert"
	"github.com/swaggo/swag"
)

func TestSwaggerDocumentation(t *testing.T) {
	router := SetupRouter()

	t.Run("Positive: Swagger Spec Is Registered", func(t *testing.T) {
		doc, err := swag.ReadDoc("swagger")
		assert.NoError(t, err)
		assert.NotEmpty(t, doc)
	})

	t.Run("Positive: Spec Contains All API Paths", func(t *testing.T) {
		doc, err := swag.ReadDoc("swagger")
		assert.NoError(t, err)

		var spec map[string]interface{}
		err = json.Unmarshal([]byte(doc), &spec)
		assert.NoError(t, err)

		paths, ok := spec["paths"].(map[string]interface{})
		assert.True(t, ok, "spec must contain paths")

		assert.Contains(t, paths, "/data")
		assert.Contains(t, paths, "/status")
		assert.Contains(t, paths, "/info")
	})

	t.Run("Positive: Spec Has Correct Title And Version", func(t *testing.T) {
		doc, _ := swag.ReadDoc("swagger")

		var spec map[string]interface{}
		json.Unmarshal([]byte(doc), &spec)

		info := spec["info"].(map[string]interface{})
		assert.Equal(t, "Go Gin API", info["title"])
		assert.Equal(t, "1.0", info["version"])
	})

	t.Run("Positive: Data Endpoint Has POST Method", func(t *testing.T) {
		doc, _ := swag.ReadDoc("swagger")

		var spec map[string]interface{}
		json.Unmarshal([]byte(doc), &spec)

		paths := spec["paths"].(map[string]interface{})
		dataPath := paths["/data"].(map[string]interface{})
		assert.Contains(t, dataPath, "post")
	})

	t.Run("Positive: Swagger Route Is Registered", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/swagger/", nil)
		router.ServeHTTP(w, req)

		assert.NotEqual(t, http.StatusMethodNotAllowed, w.Code)
	})

	t.Run("Negative: Access Docs Without Swagger Prefix", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/doc.json", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Negative: Unregistered Spec Name Returns Error", func(t *testing.T) {
		_, err := swag.ReadDoc("nonexistent")
		assert.Error(t, err)
	})
}
