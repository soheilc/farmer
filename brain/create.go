package brain

import (
	"os"

	"errors"
	"github.com/farmer-project/farmer/api/request"
	"github.com/farmer-project/farmer/box"
	"github.com/farmer-project/farmer/utils/farmerFile"
	"github.com/farmer-project/farmer/utils/git"
)

var GREEN_HOUSE = os.Getenv("GREENHOUSE_VOLUME")

// TODO: Add a method to brain that init remotely
func Create(req request.CreateSeedRequest) error {
	codeDest := GREEN_HOUSE + "/" + req.Name

	// 1. Clone code
	if err := git.Clone(req.Repo, req.PathSpec, codeDest); err != nil {
		return err
	}

	// 2. Read .farmer.yml and fetch it's data
	ff, err := farmerFile.Parse(codeDest)
	if err != nil {
		os.RemoveAll(codeDest)
		return err
	}

	// 3. Init box configuration
	box := box.Box{
		Name: req.Name,
		Git: &box.GitConfig{
			Repo:     req.Repo,
			PathSpec: req.PathSpec,
		},
	}

	// 4. Create an container
	err = box.Run(boxConfigure(ff, codeDest))
	if err != nil {
		os.RemoveAll(codeDest)
		return err
	}

	return errors.New("")
}

func boxConfigure(conf farmerFile.ConfigFile, codeFolderAddress string) box.BoxConfig {
	user := "level1"
	network := &box.ContainerNetworkSetting{
		Ports: conf.Ports,
	}
	return box.BoxConfig{
		Image:        conf.Image,
		CgroupParent: user,
		Hostname:     user,
		Binds:        []string{"/app:" + codeFolderAddress}, // Any container has one source code
		Network:      network,
	}
}
