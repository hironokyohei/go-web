package meander

import "strings"

type j struct {
	Name       string
	PlaceTypes []string
}

// Journeys ...
var Journeys = []interface{}{
	&j{Name: "ロマンティック", PlaceTypes: []string{"park", "bar",
		"movie_theater", "restaurant", "florist", "taxi_stand"}},
	&j{Name: "ショッピング", PlaceTypes: []string{"department_store",
		"cafe", "clothing_store", "jewelry_store", "shoe_store"}},
	&j{Name: "ナイトライフ", PlaceTypes: []string{"bar", "casion",
		"food", "bar", "night_club", "bar", "bar", "hospital"}},
	&j{Name: "カルチャー", PlaceTypes: []string{"museum", "cafe",
		"ceretery", "library", "art_gallery"}},
	&j{Name: "リラックス", PlaceTypes: []string{"hair_care",
		"beauty_salon", "cafe", "spa"}},
}

// Public ...
func (j *j) Public() interface{} {
	return map[string]interface{}{
		"name":    j.Name,
		"journey": strings.Join(j.PlaceTypes, "|"),
	}
}
