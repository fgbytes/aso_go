package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type AppProfileResponce struct {
	ResultCount int          `json:"resultCount"`
	Results     []AppProfile `json:"results"`
}

type AppProfile struct {
	ArtistViewURL                      string        `json:"artistViewUrl"`
	ArtworkURL60                       string        `json:"artworkUrl60"`
	ArtworkURL100                      string        `json:"artworkUrl100"`
	AppletvScreenshotUrls              []interface{} `json:"appletvScreenshotUrls"`
	IpadScreenshotUrls                 []string      `json:"ipadScreenshotUrls"`
	ScreenshotUrls                     []string      `json:"screenshotUrls"`
	ArtworkURL512                      string        `json:"artworkUrl512"`
	IsGameCenterEnabled                bool          `json:"isGameCenterEnabled"`
	Advisories                         []interface{} `json:"advisories"`
	SupportedDevices                   []string      `json:"supportedDevices"`
	Kind                               string        `json:"kind"`
	Features                           []string      `json:"features"`
	TrackCensoredName                  string        `json:"trackCensoredName"`
	TrackViewURL                       string        `json:"trackViewUrl"`
	ContentAdvisoryRating              string        `json:"contentAdvisoryRating"`
	AverageUserRatingForCurrentVersion float64       `json:"averageUserRatingForCurrentVersion"`
	LanguageCodesISO2A                 []string      `json:"languageCodesISO2A"`
	SellerURL                          string        `json:"sellerUrl"`
	UserRatingCountForCurrentVersion   int           `json:"userRatingCountForCurrentVersion"`
	TrackContentRating                 string        `json:"trackContentRating"`
	FileSizeBytes                      string        `json:"fileSizeBytes"`
	FormattedPrice                     string        `json:"formattedPrice"`
	CurrentVersionReleaseDate          time.Time     `json:"currentVersionReleaseDate"`
	PrimaryGenreID                     int           `json:"primaryGenreId"`
	ReleaseDate                        time.Time     `json:"releaseDate"`
	ReleaseNotes                       string        `json:"releaseNotes"`
	MinimumOsVersion                   string        `json:"minimumOsVersion"`
	SellerName                         string        `json:"sellerName"`
	Currency                           string        `json:"currency"`
	WrapperType                        string        `json:"wrapperType"`
	Version                            string        `json:"version"`
	TrackID                            int           `json:"trackId"`
	ArtistID                           int           `json:"artistId"`
	ArtistName                         string        `json:"artistName"`
	Genres                             []string      `json:"genres"`
	Price                              float64       `json:"price"`
	Description                        string        `json:"description"`
	TrackName                          string        `json:"trackName"`
	BundleID                           string        `json:"bundleId"`
	PrimaryGenreName                   string        `json:"primaryGenreName"`
	GenreIds                           []string      `json:"genreIds"`
	IsVppDeviceBasedLicensingEnabled   bool          `json:"isVppDeviceBasedLicensingEnabled"`
	UserRatingCount                    int           `json:"userRatingCount"`
	AverageUserRating                  float64       `json:"averageUserRating"`
}

// https://itunes.apple.com/lookup?id=368677368&country=en&entity=software
func LookUp(id int, country string) AppProfile {
	app := AppProfileResponce{}

	url := fmt.Sprintf("https://itunes.apple.com/lookup?id=%d&country=%s&entity=software", id, country)
	resp, err := getWebserviceResponse(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal()
	}
	err = json.Unmarshal(body, &app)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(app.Results[0])
	return app.Results[0]
}

func getWebserviceResponse(url string) (*http.Response, error) {
	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Spoof chrome user agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36")

	// Send the request via a client
	client := &http.Client{}
	return client.Do(req)
}
