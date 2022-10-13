package deletedomain

import (
	"bytes"
	"errors"
	"github.com/ar-mokhtari/orginfo-generator/entity"
	"github.com/ar-mokhtari/orginfo-generator/env"
	"io/ioutil"
	"os"
	"strings"
)

func DeleteDomain(domain entity.Domain) error {
	//remove domain in layers (delivery / usecase / dto / entity)
	for _, layer := range []string{env.Delivery + "/http/V1", env.Usecase, env.Usecase + "/validation", env.DTO, env.Entity} {
		err := os.RemoveAll(env.MainPath + layer + "/" + domain.SnakeName)
		if err != nil {
			return err
		}
	}
	//remove from storage layer
	storageActionErr := os.RemoveAll(env.MainPath + env.Adapter + "/storage/" + "action_" + domain.SnakeName + ".go")
	if storageActionErr != nil {
		return storageActionErr
	}
	storageModelErr := os.RemoveAll(env.MainPath + env.Adapter + "/storage/models/" + domain.SnakeName + ".go")
	if storageModelErr != nil {
		return storageModelErr
	}
	//remove domain route init in env.MainPath + env.Delivery + "/http/V1/init.go"
	deliveryHttpV1Path := env.MainPath + env.Delivery + "/http/V1/init.go"
	if _, existErr := os.Stat(deliveryHttpV1Path); !errors.Is(existErr, os.ErrNotExist) {
		input, readErr := ioutil.ReadFile(deliveryHttpV1Path)
		if readErr != nil {
			return readErr
		}
		contentText := "\n\t" + env.Delivery + domain.UpperName + ".Routs(echo)"
		contentImportText := "\n\t" + env.Delivery + domain.UpperName + " " + `"` + env.MainRepositoryPath + "/delivery/http/V1/" + domain.SnakeName + `"`
		textExist := strings.Contains(string(input), contentText)
		if textExist {
			output := bytes.Replace(input, []byte(contentText), []byte(nil), -1)
			output = bytes.Replace(output, []byte(contentImportText), []byte(nil), -1)
			if writeErr := ioutil.WriteFile(deliveryHttpV1Path, output, 0666); writeErr != nil {
				return writeErr
			}
		}
	}

	//remove domain model for migrate in env.MainPath +"adapter/storage/setup.go"
	storagePath := env.MainPath + "adapter/storage/setup.go"
	if _, existErr := os.Stat(storagePath); !errors.Is(existErr, os.ErrNotExist) {
		storageInput, readStoreErr := ioutil.ReadFile(storagePath)
		if readStoreErr != nil {
			return readStoreErr
		}
		storeContentText := "storage." + domain.UpperName + "{}, "
		storeTextExist := strings.Contains(string(storageInput), storeContentText)
		if storeTextExist {
			output := bytes.Replace(storageInput, []byte(storeContentText), []byte(nil), -1)
			if writeErr := ioutil.WriteFile(storagePath, output, 0666); writeErr != nil {
				return writeErr
			}
		}
	}
	return nil
}
