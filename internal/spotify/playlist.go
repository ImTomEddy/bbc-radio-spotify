package spotify

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const createPlaylistEndpoint = "https://api.spotify.com/v1/users/%s/playlists"

//Playlist is a data structure for playlists
type Playlist struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Owner struct {
		ID string `json:"id"`
	} `json:"owner"`
}

//GetPlaylist gets a playlist with the given name
func GetPlaylist(name string) (p *Playlist, err error) {
	u, err := url.Parse(searchEndpoint)

	if err != nil {
		return nil, err
	}

	q := url.Values{}

	q.Add("q", name)
	q.Add("type", "playlist")
	q.Add("limit", "1")

	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)

	req.Header.Add("Authorization", "Bearer "+token)

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	var r struct {
		Playlists struct {
			Items []Playlist `json:"items"`
		} `json:"playlists"`
	}

	err = json.Unmarshal(b, &r)

	if err != nil {
		return nil, err
	}

	if len(r.Playlists.Items) == 0 || r.Playlists.Items[0].Owner.ID != userID {
		return nil, errors.New("can't find playlist")
	}

	p = &r.Playlists.Items[0]

	if p.Name != name {
		return nil, errors.New(fmt.Sprint("can't find playlist, closest was ", p.Name))
	}

	return p, err
}

//CreatePlaylist creates a playlist with the given name and description
func CreatePlaylist(name string, desc string) (p *Playlist, err error) {
	u, err := url.Parse(fmt.Sprintf(createPlaylistEndpoint, userID))

	if err != nil {
		return nil, err
	}

	b := map[string]string{
		"name":        name,
		"description": desc,
	}

	body, err := json.Marshal(b)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	rb, err := ioutil.ReadAll(resp.Body)

	p = &Playlist{}

	err = json.Unmarshal(rb, p)

	if err != nil {
		return nil, err
	}

	return p, err
}

//GetOrCreatePlaylist gets a playlist or creates it
func GetOrCreatePlaylist(name string, desc string) (p *Playlist, err error) {
	p, err = GetPlaylist(name)

	if p != nil && err == nil {
		return p, err
	}

	return CreatePlaylist(name, desc)
}
