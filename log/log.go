package log

import (
		
		"io/ioutil"
		"log"
		"os"
)

var (
	Trace   *log.Logger
    Info    *log.Logger
    Warning *log.Logger
	Error   *log.Logger
	Fatal	*log.Logger
)



func init() {

	
	Trace = log.New(ioutil.Discard,
        "TRACE: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Info = log.New(os.Stdout,
        "INFO: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Warning = log.New(os.Stdout,
        "WARNING: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    Error = log.New(os.Stdout,
        "ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	
	Fatal = log.New(os.Stdout,
        "FATAL: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}