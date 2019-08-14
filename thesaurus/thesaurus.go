package thesaurus

// Thesaurus torima
type Thesaurus interface {
	Synonyms(term string) ([]string, error)
}
