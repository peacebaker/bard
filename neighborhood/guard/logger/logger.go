package logger

import (
	"log"
	"os"
)

const ERRlOGfILE = "/var/log/secretSocial/guard.err"

var ErrLog *os.File
var Log *log.Logger

func Open() error {

	// I don't really want the service to start if it can't open its error log
	var err error
	if ErrLog, err = os.OpenFile(ERRlOGfILE, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640); err != nil {
		return err
	}

	// make the logger
	Log = log.New(ErrLog, "", log.LstdFlags)
	return nil
}

func Close() {
	ErrLog.Close()
}
