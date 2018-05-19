package main

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"time"

	termutil "github.com/japanoise/termbox-util"
	"github.com/kennygrant/sanitize"
	homedir "github.com/mitchellh/go-homedir"
)

func initDirs(subdir string) (string, string, error) {
	configdir := os.Getenv("XDG_CONFIG_HOME")
	datadir := os.Getenv("XDG_DATA_HOME")
	if configdir == "" {
		h, err := homedir.Dir()
		if err != nil {
			return configdir, datadir, err
		}
		configdir = h + string(os.PathSeparator) + ".config" + string(os.PathSeparator) + subdir
	} else {
		configdir += string(os.PathSeparator) + subdir
	}
	if datadir == "" {
		h, err := homedir.Dir()
		if err != nil {
			return configdir, datadir, err
		}
		datadir = h + string(os.PathSeparator) + ".local" + string(os.PathSeparator) + "share" + string(os.PathSeparator) + subdir
	} else {
		datadir += string(os.PathSeparator) + subdir
	}
	err := os.MkdirAll(datadir, 0755)
	if err != nil {
		return configdir, datadir, err
	}
	err = os.MkdirAll(configdir, 0755)
	return configdir, datadir, err
}

func generateFileName(g *Gamedata) string {
	return DataDir + string(os.PathSeparator) + sanitize.BaseName(g.Chars[g.PlayerId].GetNameString()) + "-" + time.Now().Format("2006-01-02-15-04-05") + ".json.gz"
}

func SaveGame(g *Gamedata) string {
	fn := generateFileName(g)

	file, err := os.Create(fn)
	if err != nil {
		return err.Error()
	}
	defer file.Close()

	zipper := gzip.NewWriter(file)
	defer zipper.Close()

	encoder := json.NewEncoder(zipper)
	err = encoder.Encode(g)
	if err != nil {
		return err.Error()
	}
	return fn
}

func LoadGame(fn string) (*Gamedata, error) {
	ret := &Gamedata{}

	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	unzipper, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer unzipper.Close()

	decoder := json.NewDecoder(unzipper)
	err = decoder.Decode(ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func LoadGameChoice() (*Gamedata, error) {
	savefiles, err := ioutil.ReadDir(DataDir)
	if err != nil {
		return nil, err
	} else if len(savefiles) == 0 {
		return nil, errors.New("No save files found")
	}

	fnchoice := make([]string, 0, len(savefiles))
	for _, filedata := range savefiles {
		if strings.HasSuffix(filedata.Name(), ".json.gz") && !filedata.IsDir() {
			fnchoice = append(fnchoice, filedata.Name())
		}
	}
	fn := termutil.ChoiceIndex("Load which save game?", fnchoice, 0)
	return LoadGame(DataDir + string(os.PathSeparator) + fnchoice[fn])
}
