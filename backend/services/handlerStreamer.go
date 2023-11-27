package services

import (
	"SLtwitchApi/db"
	"SLtwitchApi/models"
	"SLtwitchApi/twitch"
	"errors"
	"log"
	"os"
	"sync"

	"time"

	"github.com/joho/godotenv"
)

var (
	twitchClient *twitch.TwitchClient
)
// var debug = false

func InitStreamerCall() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
	twitchClientID := os.Getenv("TWITCH_CLIENT_ID")
	twitchBearerToken := os.Getenv("TWITCH_BEARER_TOKEN")
	twitchClient = twitch.NewTwitchClient(twitchClientID, twitchBearerToken)

	// fmt.Println(twitchClientID, twitchBearerToken)

	go CheckStreamerCall()
	go IsStreamerLive()
}

func CheckStreamerCall() {
	log.Printf("Checking streamers…")

	// Erstellen Sie einen Channel, um die Streamer-IDs zu übergeben
	streamerIDs := make(chan string)

	// Starten Sie mehrere Worker, um die Streamer-IDs zu verarbeiten
	var wg sync.WaitGroup
	const numWorkers = 10
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			for batch := range batchStreamerIDs(streamerIDs) {
				if err := updateStreamers(batch); err != nil {
					log.Printf("Failed to update streamers: %v", err)
				}
			}
		}()
	}

	// Lese alle Benutzer aus der Datenbank, die aktualisiert werden müssen
	var streamers []models.Streamer
	if err := db.SQL.Where("dataFetched = ?", false).Find(&streamers).Error; err != nil {
		log.Fatalf("Failed to retrieve streamers from the database: %v", err)
	}

	// Fügen Sie die Streamer-IDs dem Channel hinzu
	for _, streamer := range streamers {
		streamerIDs <- streamer.Username
	}

	// Schließen Sie den Channel, um die Worker zu beenden
	close(streamerIDs)

	// Warten Sie, bis alle Worker beendet sind
	wg.Wait()
}

func batchStreamerIDs(streamerIDs chan string) chan []string {
	const batchSize = 100
	batches := make(chan []string)
	go func() {
		defer close(batches)
		var batch []string
		for id := range streamerIDs {
			batch = append(batch, id)
			if len(batch) == batchSize {
				batches <- batch
				batch = nil
			}
		}
		if len(batch) > 0 {
			batches <- batch
		}
	}()
	return batches
}

func updateStreamers(ids []string) error {
	tx := db.SQL.Begin()
	for _, id := range ids {
		streamer, err := CheckStreamHandler(id)
		if err != nil {
			log.Printf("Failed to update streamer %s: %v", id, err)
			if err.Error() == "user not found on twitch" {
				if err := tx.Delete(&streamer).Error; err != nil {
					tx.Rollback()
					return err
				}
			}
		} else {
			log.Printf("Updated streamer %s", id)
			RefreshLiveStreamerData()
			tx.Rollback()
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
func CheckStreamHandler(username string) (models.Streamer, error) {
	var streamer models.Streamer
	result := db.SQL.Where("user_name = ?", username).First(&streamer)
	if result.Error != nil {
		return streamer, result.Error
	}

	if streamer.DataFetched {
		return streamer, nil
	}

	twitchResponse, err := twitchClient.FindUser(streamer.Username)
	if err != nil {
		return streamer, err
	}

	if len(twitchResponse.Data) == 0 {
		return streamer, errors.New("user not found on twitch")
	}

	twitchUser := twitchResponse.Data[0]

	twitchFollowersResponse, err := twitchClient.GetFollowersForStreamer(twitchUser.ID)
	log.Print(twitchFollowersResponse);
	if err != nil {
		return streamer, err
	}

	tx := db.SQL.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	streamer.Viewers = nil
	streamer.ThumbnailImg = nil
	streamer.Title = nil
	streamer.DataFetched = true
	streamer.Followers = twitchFollowersResponse.Total
	streamer.ProfileImg = twitchUser.ProfileImage
	streamer.LastStreamed = nil
	if err := tx.Save(&streamer).Error; err != nil {
		tx.Rollback()
		return streamer, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return streamer, err
	}

	return streamer, nil
}

func IsStreamerLive() {
	log.Printf("Checking streamers for live status...")

	ticker := time.NewTicker(4 * time.Minute)
	defer ticker.Stop()

	RefreshLiveStreamerData()

	for range ticker.C {
		RefreshLiveStreamerData()
	}
}
func RefreshLiveStreamerData() {
	// Lese alle Benutzer aus der Datenbank, die aktualisiert werden müssen
	var streamers []models.Streamer
	if err := db.SQL.Where("dataFetched = ?", true).Find(&streamers).Error; err != nil {
		log.Fatalf("Fehler beim Abrufen der Streamer aus der Datenbank: %v", err)
	}
	log.Printf("Überprüfe Live-Status für %v Streamer", len(streamers))
	t1 := time.Now()
	for _, streamer := range streamers {
		twitchResponse, err := twitchClient.GetStreamInformation(streamer.Username)
		if err != nil {
			continue
		}
		if twitchResponse.Data != nil && len(twitchResponse.Data) > 0 {
			log.Println(twitchResponse)
		}

		if len(twitchResponse.Data) == 0 {
			// Der Streamer ist nicht live, setze den IsLive-Status auf 0 und lösche die Viewer und die ThumbnailImg
			if err := db.SQL.Model(&streamer).Updates(map[string]interface{}{"IsLive": 0, "Viewers": nil, "ThumbnailImg": nil, "Title": nil}).Error; err != nil {
				log.Printf("Fehler beim Aktualisieren der Daten für Streamer %s: %v", streamer.Username, err)
			}
			continue
		}
		// if debug {
		// 	if twitchResponse.Data[0].GameName != "Grand Theft Auto V" {
		// 		// Debug-Nachricht für Spiel nicht GTA V
		// 		log.Printf("DEBUG: Der Streamer %s spielt nicht GTA V. Überspringe die Aktualisierung.", streamer.Username)
		// 	}

		// 	if !strings.Contains(strings.ToLower(twitchResponse.Data[0].Title), "storyline") &&
		// 		!strings.Contains(strings.ToLower(twitchResponse.Data[0].Title), "[storyline]") {
		// 		// Debug-Nachricht für fehlende Storyline im Titel
		// 		log.Printf("DEBUG: Der Streamer %s hat kein 'Storyline' im Titel. Setze den IsLive-Status auf 0.", streamer.Username)
		// 	}
		// }

		// Der Streamer ist live und spielt GTA V, aktualisiere den IsLive-Status, Viewer und ThumbnailImg
		timestamp := twitchResponse.Data[0].StartedAt
		time, err := time.Parse(time.RFC3339, timestamp)
		if err != nil {
			log.Printf("Fehler beim Parsen des Zeitstempels für Streamer %s: %v", streamer.Username, err)
			continue
		}
		// if err := db.SQL.Model(&streamer).Exec("ALTER TABLE twitch MODIFY Title VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci").Error; err != nil {
		// 	log.Printf("Fehler beim Ändern des Zeichensatzes für den Title: %v", err)
		// 	continue
		// }

		if err := db.SQL.Model(&streamer).Updates(map[string]interface{}{"IsLive": 1, "Viewers": twitchResponse.Data[0].ViewerCount, "ThumbnailImg": twitchResponse.Data[0].ThumbnailUrl, "Title": twitchResponse.Data[0].Title}).Error; err != nil {
			log.Printf("Fehler beim Aktualisieren der Daten für Streamer %s: %v", streamer.Username, err)
			continue
		}
		if err := db.SQL.Model(&streamer).Update("LastStreamed", &time).Error; err != nil {
			log.Printf("Fehler beim Aktualisieren des LastStreamed-Werts für Streamer %s: %v", streamer.Username, err)
			continue
		}
	}
	log.Println("Überprüfung des Live-Status hat gedauert:", time.Since(t1))
}

func GetAllStreamerInformation() (streamer []models.Streamer) {
	db.SQL.Where("active = ?", true).Find(&streamer)
	return
}

func GetStreamerDetailsByUsername(username string) (streamer []models.StreamerDetails) {
	db.SQL.Table("streamer").Select("user_name, shortStory, ProfileImg, followers, fraktion, charname").Where("dataFetched = ? AND username = ? AND active = ?", true, username, true).Scan(&streamer)
	return
}
