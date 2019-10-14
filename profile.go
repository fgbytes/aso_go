package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type LookupResponce struct {
	ResultCount int       `json:"resultCount"`
	Results     []Profile `json:"results"`
}

type Profile struct {
	DeveloperAppstoreURL             string        `json:"artistViewUrl,omitempty"`
	DeveloperStoreURL                string        `json:"artistLinkUrl,omitempty"`
	Icon60                           string        `json:"artworkUrl60,omitempty"`
	Icon100                          string        `json:"artworkUrl100,omitempty"`
	AppletvScreenshot                []interface{} `json:"appletvScreenshotUrls,omitempty"`
	IpadScreenshot                   []string      `json:"ipadScreenshotUrls,omitempty"`
	Screenshot                       []string      `json:"screenshotUrls,omitempty"`
	Icon512                          string        `json:"artworkUrl512,omitempty"`
	IsGameCenterEnabled              bool          `json:"isGameCenterEnabled,omitempty"`
	SupportedDevices                 []string      `json:"supportedDevices,omitempty"`
	Kind                             string        `json:"kind,omitempty"`
	Features                         []string      `json:"features,omitempty"`
	URL                              string        `json:"trackViewUrl,omitempty"`
	ContentRating                    string        `json:"contentAdvisoryRating,omitempty"`
	RatingForCurrentVersion          float64       `json:"averageUserRatingForCurrentVersion,omitempty"`
	Languages                        []string      `json:"languageCodesISO2A,omitempty"`
	DeveloperWebsite                 string        `json:"sellerUrl,omitempty"`
	ReviewsForCurrentVersion         int           `json:"userRatingCountForCurrentVersion,omitempty"`
	AppSize                          string        `json:"fileSizeBytes,omitempty"`
	FormattedPrice                   string        `json:"formattedPrice,omitempty"`
	UpdatedLast                      time.Time     `json:"currentVersionReleaseDate,omitempty"`
	CategoryID                       int           `json:"primaryGenreId,omitempty"`
	FirstReleaseDate                 time.Time     `json:"releaseDate,omitempty"`
	ReleaseNotes                     string        `json:"releaseNotes,omitempty"`
	MinimumOsVersion                 string        `json:"minimumOsVersion,omitempty"`
	DeveloperName                    string        `json:"sellerName,omitempty"`
	Currency                         string        `json:"currency,omitempty"`
	Version                          string        `json:"version,omitempty"`
	AppID                            int           `json:"trackId,omitempty"`
	DeveloperID                      int           `json:"artistId,omitempty"`
	ArtistName                       string        `json:"artistName,omitempty"`
	CategoriesLocalised              []string      `json:"genres,omitempty"`
	Price                            float64       `json:"price,omitempty"`
	Description                      string        `json:"description,omitempty"`
	Title                            string        `json:"trackName,omitempty"`
	BundleID                         string        `json:"bundleId,omitempty"`
	PrimaryCategoryName              string        `json:"primaryGenreName,omitempty"`
	GenreIds                         []string      `json:"genreIds,omitempty"`
	IsVppDeviceBasedLicensingEnabled bool          `json:"isVppDeviceBasedLicensingEnabled,omitempty"`
	UserRatingCount                  int           `json:"userRatingCount,omitempty"`
	AverageUserRating                float64       `json:"averageUserRating,omitempty"`
}

func LookUp(id int, country string) []Profile {
	app := LookupResponce{}

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
	//	log.Println(app.Results[0])
	return app.Results
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
