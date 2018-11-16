package episodes

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

// EpisodeList is a simple collection of Episodes.
type EpisodeList []*Episode

// LoadFromJSON loads all episode data from a JSON file.
func (el *EpisodeList) LoadFromJSON(filename string) {
	jsn, err := ioutil.ReadFile(filename)
	panicIf(err)

	err = json.Unmarshal(jsn, &el)
	panicIf(err)
}

// Random selects an Episode randomly from the EpisodeList.
func (el *EpisodeList) Random() *Episode {
	rand.Seed(time.Now().Unix())
	return (*el)[rand.Intn(len(*el))]
}

// Add appends an Episode to the EpisodeList.
func (el *EpisodeList) Add(e Episode) {
	(*el) = append((*el), &e)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
