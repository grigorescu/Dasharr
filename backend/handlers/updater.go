package handlers

import (
	"backend/helpers"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/tidwall/gjson"
)

func Update(c echo.Context) error {
	db, _ := sql.Open("sqlite3", "database.db")

	configFile, _ := os.ReadFile("config/config.json")
	trackerConfigs := gjson.Parse(string(configFile))

	trackerConfigs.ForEach(func(_, trackerConfig gjson.Result) bool {
		trackerName := trackerConfig.Get("tracker_name").String()
		enabled := trackerConfig.Get("enabled").Bool()

		if enabled {
			fmt.Printf("Updating %s's stats\n", trackerName)

			trackerStats := helpers.GetUserData(trackerConfig)

			insertSQL := `INSERT INTO user_stats (
			tracker_id, uploaded_torrents, uploaded_amount, snatched, seeding, leeching, ratio
			) VALUES (?, ?, ?, ?, ?, ?, ?);`
			_, err := db.Exec(insertSQL,
				trackerConfig.Get("tracker_id").String(),
				trackerStats["uploaded_torrents"],
				trackerStats["uploaded_amount"],
				trackerStats["snatched"],
				trackerStats["seeding"],
				trackerStats["leeching"],
				trackerStats["ratio"],
			)
			if err != nil {
				fmt.Println("Failed to insert data:", err.Error())
			}
		}

		return true
	})

	return c.String(http.StatusOK, "Data inserted successfully!")
}
