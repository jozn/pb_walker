package ant

import (
	"github.com/dsymonds/gotoc/parser"
	"github.com/emicklei/proto"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"ms/sun/helper"
	"os"
	_ "os"
	"os/exec"
	_ "os/exec"
	"path"
)

const OUTPUT_DIR = `C:\Go\_gopath\src\ms\sun\models\x\`                             //"./play/gen_sample_out.go"
const OUTPUT_ANDROID_DIR_ = `D:\dev_working2\MS_Native\app\src\main\java\ir\ms\pb\` //"./play/gen_sample_out.go"
//const TEMPLATES_DIR = "./templates/"                    //relative to main func of parent directory
const TEMPLATES_DIR = `C:\Go\_gopath\src\ms\ants/templates/` //relative to main func of parent directory
const DIR_PROTOS = `C:\Go\_gopath\src\ms\sun\models\protos`

func Run() {
	xxx()
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
	/*os.Chdir(`C:\Go\_gopath\src\ms\sun\scripts\`)
	err = exec.Command(`C:\Go\_gopath\src\ms\sun\scripts\gen_pb.exe`).Run()
	noErr(err)
	err = exec.Command("gofmt", "-w", OUTPUT_DIR).Run()*/
	//noErr(err)

	err = exec.Command("javafmt").Run()
	////
}

func RunV2() {
	//xxx()
	files, err := ioutil.ReadDir(DIR_PROTOS)
	noErr(err)
	protos := make([]string, len(files))
	var prtos []*proto.Proto
	for i, f := range files {
		protos[i] = f.Name()
		reader, err := os.Open(path.Join(DIR_PROTOS, f.Name()))
		helper.NoErr(err)
		defer reader.Close()
		parser := proto.NewParser(reader)
		def, err := parser.Parse()
		if err != nil {
			log.Panic("error parsing proto: ", f.Name(), " ", err, "\n")
		}
		prtos = append(prtos, def)
	}

	gen := &GenOut{
		Messages: ExtractAllMessagesViewsV2(prtos),
		Services: ExtractAllServicesViewsV2(prtos),
		Enums:    ExtractAllEnumsViewsV2(prtos),
	}

    print("===========================================")
    helper.PertyPrint(gen.Messages)
    //helper.PertyPrint(prtos)

	build(gen)

	/////////// commeant albe ///
	/*os.Chdir(`C:\Go\_gopath\src\ms\sun\scripts\`)
	  err = exec.Command(`C:\Go\_gopath\src\ms\sun\scripts\gen_pb.exe`).Run()
	  noErr(err)
	  err = exec.Command("gofmt", "-w", OUTPUT_DIR).Run()*/
	//noErr(err)

	err = exec.Command("javafmt").Run()
	////
}

func alaki() {
	exec.Command(``).Run()
	os.Chdir(``)
}
