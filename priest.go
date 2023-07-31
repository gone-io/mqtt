package mqtt

import (
	"github.com/gone-io/gone"
	"github.com/gone-io/gone/goner/config"
)

func Priest(cemetery gone.Cemetery) error {
	err := config.Priest(cemetery)
	if err != nil {
		return err
	}
	if nil == cemetery.GetTomById(IdGoneMqtt) {
		cemetery.Bury(NewClient())
	}
	return nil
}
