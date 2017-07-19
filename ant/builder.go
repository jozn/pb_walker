package ant

import (
    "io/ioutil"
    "fmt"
    "os"
    "text/template"
    "bytes"
)

func templSerivces(services []ServiceView) {
    tpl := template.New("go_interface")
    tplGoInterface,err := ioutil.ReadFile(TEMPLATES_DIR+ "rpc.go")
    noErr(err)
    tpl, err = tpl.Parse(string(tplGoInterface))
    noErr(err)

    s := struct {
        Services []ServiceView
    }{
        Services: services,
    }

    bts := bytes.NewBufferString("")
    err = tpl.Execute(bts, s)
    noErr(err)

    fmt.Println(bts.String())
    ioutil.WriteFile(DIR_OUTPUT, []byte(bts.String()), os.ModeType)
}

