package handlers

import (
	"backend/database"
	"backend/helpers"
	"backend/indexers"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
)

func SaveCredentials(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	jsonBody := gjson.Parse(string(body))

	prowlarrIds := getProwlarrIndexerIdsFromDB()

	indexer := jsonBody.Get("indexer").Str
	indexerId := ""
	for key, value := range prowlarrIds {
		if strings.Contains(value, indexer) {
			indexerId = key
			break
		}
	}
	username := jsonBody.Get("username").Str
	password := jsonBody.Get("password").Str
	cookies := loginAndGetCookies(indexer, username, password)

	insertSQL := `INSERT OR REPLACE INTO credentials (
	indexer_id, username, password, cookies, api_key
	) VALUES (?, ?, ?, ?, ?);
	`

	args := []interface{}{indexerId, username, password, cookies, jsonBody.Get("api_key").Str}

	database.ExecuteQuery(insertSQL, args)

	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}

func loginAndGetCookies(indexer string, username string, password string) string {
	indexerInfo := helpers.GetIndexerInfo(indexer)
	if !indexerInfo.Exists() {
		return ""
	}
	indexerType := indexers.DetermineIndexerType(indexer)

	loginURL := indexerInfo.Get("login.url").String()

	if indexerType == "unit3d" {
		return indexers.LoginAndGetCookiesUnit3d(username, password, loginURL, indexerInfo.Get("domain").Str)
	} else if indexerType == "anthelion" {
		return indexers.LoginAndGetCookiesAnthelion(username, password, loginURL, indexerInfo)
	}

	return ""

}

// returns the indexers that have their credentials registered in Dasharr's db by the user
func SavedCredentials(c echo.Context) error {
	sql := `SELECT indexer_id from credentials`
	results := database.ExecuteQuery(sql, []interface{}{})
	indexerNames := getProwlarrIndexerIdsFromDB()

	for _, obj := range results {
		if id, ok := obj["indexer_id"].(int64); ok {
			obj["indexer_name"] = strings.TrimSuffix(indexerNames[fmt.Sprint(id)], " (API)")
		}
	}
	return c.JSON(http.StatusOK, results)
}
