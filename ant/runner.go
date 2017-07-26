package ant

import (
	"github.com/dsymonds/gotoc/parser"
	"io/ioutil"
	"os"
	_ "os"
	"os/exec"
	_ "os/exec"
)

const OUTPUT_DIR = `C:\Go\_gopath\src\ms\sun\models\x\` //"./play/gen_sample_out.go"
const TEMPLATES_DIR = "./templates/"                    //relative to main func of parent directory

func Run() {
	const DIR_PROTOS = `C:\Go\_gopath\src\ms\sun\models\protos`
	files, err := ioutil.ReadDir(DIR_PROTOS)
	noErr(err)
	protos := make([]string, len(files))
	for i, f := range files {
		protos[i] = f.Name()
	}

	//ast, err := parser.ParseFiles([]string{"1.proto"}, []string{"./play/"})
	ast, err := parser.ParseFiles(protos, []string{DIR_PROTOS})
	noErr(err)
	gen := &GenOut{
		Messages: ExtractAllMessagesViews(ast),
		Services: ExtractAllServicesViews(ast),
		Enums:    ExtractAllEnumsViews(ast),
	}

	build(gen)

	/////////// commeant albe ///
	os.Chdir(`C:\Go\_gopath\src\ms\sun\scripts\`)
	err = exec.Command(`C:\Go\_gopath\src\ms\sun\scripts\gen_pb.exe`).Run()
	noErr(err)
	err = exec.Command("gofmt", "-w", OUTPUT_DIR).Run()
	noErr(err)

	////
}
