package main

import (
	"AxiaoA/go_project/func_c/func_14_time/myloger"
)

func main() {

	log := myloger.NewLog("Info")
	for i := 0; i < 10; i++ {
		log.Debug("debug")
		log.Info("info")
		log.Warning("Waring")
		log.Error("Error")
	}
}
