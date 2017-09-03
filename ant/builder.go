package ant

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func build(gen *GenOut) {
	OutGoRPCsStr := buildFromTemplate("rpc.tgo", gen)
	writeOutput("pb__gen_ant.go", OutGoRPCsStr)

	OutGoRPCsEmptyStr := buildFromTemplate("rpc_empty_imple.tgo", gen)
	writeOutput("pb__gen_ant_empty.go", OutGoRPCsEmptyStr)

	writeOutput("pb__gen_enum.proto", buildFromTemplate("enums.proto", gen))
	writeOutput("RPC_HANDLERS.java", buildFromTemplate("RPC_HANDLERS.java", gen))
	writeOutput("PBFlatTypes.java", buildFromTemplate("PBFlatTypes.java", gen))
	writeOutput("flat.go", buildFromTemplate("flat.tgo", gen))

	//////////////// For Android /////////////
	writeOutputAndroid("RPC_HANDLERS.java", buildFromTemplate("RPC_HANDLERS.java", gen))
	writeOutputAndroid("PBFlatTypes.java", buildFromTemplate("PBFlatTypes.java", gen))
	writeOutputAndroid("RPC.java", buildFromTemplate("RPC.java", gen))
	writeOutputAndroid("RPC_ResponseBase.java", buildFromTemplate("RPC_ResponseBase.java", gen))

	/////// Enums /////////////////

    ////////// Realm /////////////
    buildForRealms(gen.Realms)
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

func buildForRealms(msgs []MessageView) {
	buff := make(map[string]string)

	for _, realmMsgView := range msgs {
		realmMsgView.RealmClass = pbToRealmName(realmMsgView.Name)
		tpl := template.New("realm_____")
		tpl.Funcs(fns)
		tplGoInterface, err := ioutil.ReadFile(TEMPLATES_DIR + "realm.java")
		noErr(err)
		tpl, err = tpl.Parse(string(tplGoInterface))
		noErr(err)

		buffer := bytes.NewBufferString("")
		err = tpl.Execute(buffer, &realmMsgView)
		noErr(err)

		buff[realmMsgView.RealmClass] = buffer.String()
	}

	for klass, out := range buff {
		ioutil.WriteFile(OUTPUT_ANDROID_REALM_DIR_+klass + ".java", []byte(out), os.ModeType)
	}
}

func pbToRealmName(pbName string) string {
	s := strings.Replace(pbName, "PB_", "Realm", -1)
	/*if s[0:6] != "Realm" {
		s += "Realm" + s
	}*/
	return s
}
