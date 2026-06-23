package storage

import (
	"encoding/json"
	"os"
	"path"
)

type Command struct {
	Label string `json:"label"`
	Cmd   string `json:"cmd"`
}

type Library struct {
	Description string    `json:"description"`
	Install     string    `json:"install"`
	Cmd         []Command `json:"cmd"`
	Name        string
}

type Data struct {
	Libraries map[string]Library `json:"libraries"`
}

func (l Library) Title() string       { return l.Name }
func (l Library) Desc() string        { return l.Description }
func (l Library) FilterValue() string { return l.Name }

func GetDataPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dir := path.Join(home, ".devstash")

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0750)
	}

	return path.Join(dir, "data.json")
}

func ReadData() (Data, error) {
	var data Data
	data.Libraries = make(map[string]Library)
	bytes, err := os.ReadFile(GetDataPath())
	if err != nil {
		return data, err
	}

	if len(bytes) == 0 {
		return data, nil
	}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func WriteData(data Data) error {
	bytes, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(GetDataPath(), bytes, 0644)
}
func init() {
	dataFile := GetDataPath()
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		os.WriteFile(dataFile, []byte(`{"libraries": {}}`), 0644)
	}
}
