package ant

import "log"

func noErr(err error) {
    if err != nil {
        log.Panic(err)
    }
}

