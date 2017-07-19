package ant

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

func templSerivces(services []ServiceView) {
	tpl := template.New("go_interface")
	tplGoInterface, err := ioutil.ReadFile(TEMPLATES_DIR + "rpc.tgo")
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
