package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gitlab.smarthomecomputing.net/john/pokedexcli/internal/pokecache"
)

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Locations struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

func newLocations(data []byte) (*Locations, error) {
	ll := Locations{}
	jerr := json.Unmarshal(data, &ll)
	if jerr != nil {
		fmt.Println(jerr)
		return nil, jerr
	}
	return &ll, nil
}

type PokeAPIProxy struct {
	currentLocation int
	cache           *pokecache.Cache
}

func NewPokeAPIProxy() (*PokeAPIProxy, error) {
	c := pokecache.NewCache(5 * time.Minute)
	p := PokeAPIProxy{currentLocation: -1, cache: c}
	return &p, nil
}

func (p *PokeAPIProxy) checkCacheOrRetrieve(url string) ([]byte, error) {
	data, ok := p.cache.Get(url)
	if !ok {
		d, err := GetPokeData(url)
		if err != nil {
			fmt.Println("got error from pokeapi - ", err)
			return nil, err
		} else {
			// we got data from the api so store it in the cache
			p.cache.Add(url, d)
			data = d
		}
	}
	return data, nil
}

func (p *PokeAPIProxy) getLocation() ([]Location, error) {
	// so create url, then check if data already in cache
	// if not, call pokeapi to get location info, stick in cache, parse and return
	url := "https://pokeapi.co/api/v2/location/"
	if p.currentLocation > 0 {
		url = fmt.Sprintf("%s?offset=%d&limit=20", url, p.currentLocation)
	}

	data, err := p.checkCacheOrRetrieve(url)
	if err != nil {
		return nil, err
	}

	locs, err := newLocations(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return locs.Results, nil
	}
}

func (p *PokeAPIProxy) GetNextLocations() ([]Location, error) {
	// we do locations in blocks of 20
	if p.currentLocation < 0 {
		p.currentLocation = 0
	} else {
		p.currentLocation += 20
	}

	return p.getLocation()
}

func (p *PokeAPIProxy) GetPreviousLocations() ([]Location, error) {
	if p.currentLocation < 20 {
		return nil, errors.New("invalid location")
	}

	p.currentLocation -= 20

	return p.getLocation()
}

func newLocationArea(data []byte) (*LocationArea, error) {
	la := LocationArea{}
	jerr := json.Unmarshal(data, &la)
	if jerr != nil {
		fmt.Println(jerr)
		return nil, jerr
	}
	return &la, nil
}

func (p *PokeAPIProxy) GetLocationArea(area string) (*LocationArea, error) {
	url := "https://pokeapi.co/api/v2/location-area/" + area

	emptyLA := LocationArea{}

	data, err := p.checkCacheOrRetrieve(url)
	if err != nil {
		return &emptyLA, err
	}

	la, err := newLocationArea(data)
	if err != nil {
		fmt.Println("unable to unmarshal location area: " + area)
		return &emptyLA, err
	}

	return la, nil
}

func newPokemon(data []byte) (*Pokemon, error) {
	p := Pokemon{}
	err := json.Unmarshal(data, &p)
	return &p, err
}

func (p *PokeAPIProxy) GetPokemon(name string) (*Pokemon, error) {
	url := "https://pokeapi.co/api/v2/pokemon/" + name

	data, err := p.checkCacheOrRetrieve(url)
	if err != nil {
		fmt.Println("unable to unmarshal pokemon: " + name)
		return nil, err
	}

	pokemon, err := newPokemon(data)
	return pokemon, err

}
