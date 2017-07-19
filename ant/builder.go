package ant

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"
)

func build(gen *GenOut) {
    OutGoRPCsStr := getTemplateFull("rpc.tgo", gen)
    writeOutput("pb__gen_ant.go", OutGoRPCsStr )

    /////// Enums /////////////////

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
