package main

//LookUp https://itunes.apple.com/lookup?id=368677368&country=en&entity=software

func AppLookup(id int, country string) Profile {
	return LookUp(id, country)[0]
}
