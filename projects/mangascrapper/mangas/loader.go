package mangas

import (
	"encoding/json"
	"os"
)

type Chapter struct {
	ChapUrl string   `json:"chap-url"`
	ChapNo  string   `json:"chap-no"`
	Images  []string `json:"images"`
}

type Manga struct {
	Name     string
	Chapters []Chapter
}

func LoadMangas() (map[string]Manga, error) {
	mangas := make(map[string]Manga)
	var err error

	mangas["thebeginningaftertheendmanga"], err = loadManga("./mangas/thebeginningaftertheendmanga.json", "The beginning after the end")
	if err != nil {
		return nil, err
	}
	mangas["talesofdemonsandgods"], err = loadManga("./mangas/talesofdeamonsandgods.json", "Tales of demons and gods")
	if err != nil {
		return nil, err
	}

	return mangas, nil
}

func loadManga(path, name string) (Manga, error) {
	file, err := os.Open(path)
	if err != nil {
		return Manga{}, err
	}
	dec := json.NewDecoder(file)
	var chapters []Chapter
	err = dec.Decode(&chapters)
	if err != nil {
		return Manga{}, err
	}

	return Manga{
		Name:     name,
		Chapters: chapters,
	}, nil
}
