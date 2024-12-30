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
			tracker_id, uploaded_torrents, uploaded_amount, downloaded_amount, snatched, seeding, leeching, 
			ratio, required_ratio, last_access, torrent_comments, invited, forum_posts, warned, class, 
			donor, uploaded_rank, downloaded_rank, uploads_rank, requests_rank, bounty_rank, posts_rank, 
			artists_rank, overall_rank
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

			_, err := db.Exec(insertSQL,
				trackerConfig.Get("tracker_id").Int(),
				trackerStats["uploaded_torrents"],
				trackerStats["uploaded_amount"],
				trackerStats["downloaded_amount"],
				trackerStats["snatched"],
				trackerStats["seeding"],
				trackerStats["leeching"],
				trackerStats["ratio"],
				trackerStats["required_ratio"],
				trackerStats["last_access"],
				trackerStats["torrent_comments"],
				trackerStats["invited"],
				trackerStats["forum_posts"],
				trackerStats["warned"],
				trackerStats["class"],
				trackerStats["donor"],
				trackerStats["uploaded_rank"],
				trackerStats["downloaded_rank"],
				trackerStats["uploads_rank"],
				trackerStats["requests_rank"],
				trackerStats["bounty_rank"],
				trackerStats["posts_rank"],
				trackerStats["artists_rank"],
				trackerStats["overall_rank"],
			)

			if err != nil {
				fmt.Println("Failed to insert data:", err.Error())
			}
		}

		return true
	})

	return c.String(http.StatusOK, "Data inserted successfully!")
}
