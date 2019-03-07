package helpers

import "log"

func PanicOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func RecoverWithLog() {
	if rec := recover(); rec != nil {
		err := rec.(error)
		msg := err.Error()
		log.Print(msg)
	}
}
