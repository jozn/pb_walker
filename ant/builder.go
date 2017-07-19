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

func build(gen *GenOut) {
    /////// Go RPC generator ////////
   /* tplRpc:= getTemplate("rpc.tgo")
    btsRpc := bytes.NewBufferString("")
    err := tplRpc.Execute(btsRpc, gen)
    noErr(err)*/

    OutGoRPCsStr := getTemplateFull("rpc.tgo", gen)


    writeOutput("pb__gen_ant.go", OutGoRPCsStr )

    /////// Enums /////////////////


}

func getTemplate(tplName string) *template.Template {
	tpl := template.New("go_interface" + tplName)
	tplGoInterface, err := ioutil.ReadFile(TEMPLATES_DIR + tplName)
	noErr(err)
	tpl, err = tpl.Parse(string(tplGoInterface))
	noErr(err)
	return tpl
}

func getTemplateFull(tplName string, gen *GenOut) string {
    tpl := template.New("go_interface" + tplName)
    tplGoInterface, err := ioutil.ReadFile(TEMPLATES_DIR + tplName)
    noErr(err)
    tpl, err = tpl.Parse(string(tplGoInterface))
    noErr(err)

    buffer := bytes.NewBufferString("")
    err = tpl.Execute(buffer, gen)
    noErr(err)

    return buffer.String()
}

func writeOutput(fileName,output string)  {
    ioutil.WriteFile(OUTPUT_DIR + fileName, []byte(output) , os.ModeType)
}
