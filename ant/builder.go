package ant

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"
)

func build(gen *GenOut) {
	OutGoRPCsStr := buildFromTemplate("rpc.tgo", gen)
	writeOutput("pb__gen_ant.go", OutGoRPCsStr)

    OutGoRPCsEmptyStr := buildFromTemplate("rpc_empty_imple.tgo", gen)
	writeOutput("pb__gen_ant_empty.go", OutGoRPCsEmptyStr)

	writeOutput("pb__gen_enum.proto", buildFromTemplate("enums.proto", gen))
	writeOutput("RPC_INTERFACES.java", buildFromTemplate("RPC_INTERFACES.java", gen))
	writeOutput("PBFlatTypes.java", buildFromTemplate("PBFlatTypes.java", gen))
	writeOutput("flat.go", buildFromTemplate("flat.tgo", gen))

    //////////////// For Android /////////////
    writeOutputAndroid("RPC_INTERFACES.java", buildFromTemplate("RPC_INTERFACES.java", gen))
    writeOutputAndroid("PBFlatTypes.java", buildFromTemplate("PBFlatTypes.java", gen))

	/////// Enums /////////////////
}

func buildFromTemplate(tplName string, gen *GenOut) string {
	tpl := template.New("go_interface" + tplName)
	tpl.Funcs(fns)
	tplGoInterface, err := ioutil.ReadFile(TEMPLATES_DIR + tplName)
	noErr(err)
	tpl, err = tpl.Parse(string(tplGoInterface))
	noErr(err)

	buffer := bytes.NewBufferString("")
	err = tpl.Execute(buffer, gen)
	noErr(err)

	return buffer.String()
}

func writeOutput(fileName, output string) {
	ioutil.WriteFile(OUTPUT_DIR+fileName, []byte(output), os.ModeType)
}

func writeOutputAndroid(fileName, output string) {
    ioutil.WriteFile(OUTPUT_ANDROID_DIR_+fileName, []byte(output), os.ModeType)
}