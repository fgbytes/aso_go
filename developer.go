//"https://itunes.apple.com/lookup?id=284882218&country=us&entity=software"
package main

//DeveloperProfile with all app profiles included (only in the target country)
type DeveloperProfile struct {
	Apps          []Profile
	DeveloperName string
	DeveloperID   int
	DeveloperURL  string
	CountApps     int
}

// DeveloperLookup returns developer info + apps publishefd in the specific countries, [0] is developer, all further items are apps published in that country
func DeveloperLookup(id int, country string) DeveloperProfile {
	devProfile := DeveloperProfile{}
	dr := LookUp(id, country)
	devProfile.DeveloperName = dr[0].ArtistName
	devProfile.DeveloperURL = dr[0].DeveloperStoreURL
	devProfile.DeveloperID = dr[0].DeveloperID
	devProfile.CountApps = len(dr) - 1
	devProfile.Apps = dr[1:]

	return devProfile
}
