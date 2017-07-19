package ant

import (
	"fmt"
	"github.com/dsymonds/gotoc/parser"
	"io/ioutil"
	"ms/sun/helper"
)

const OUTPUT_RPC = `C:\Go\_gopath\src\ms\sun\models\x\pb__gen_ant.go` //"./play/gen_sample_out.go"
const OUTPUT_DIR = `C:\Go\_gopath\src\ms\sun\models\x\` //"./play/gen_sample_out.go"
const TEMPLATES_DIR = "./templates/"                                  //relative to main func of parent directory

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

	fmt.Println(ast, err)
	fmt.Println(err)
	for _, f := range ast.Files {
		//helper.PertyPrint(f)

		_ = f
	}

	helper.PertyPrint(ExtractAllServicesViews(ast))
	helper.PertyPrint(ExtractAllMessagesViews(ast))
	helper.PertyPrint(ExtractAllEnumsViews(ast))

	templSerivces(ExtractAllServicesViews(ast))

}

func Run2() {
    const DIR_PROTOS = `C:\Go\_gopath\src\ms\sun\models\protos`
    files, err := ioutil.ReadDir(DIR_PROTOS)
    noErr(err)
    protos := make([]string, len(files))
    for i, f := range files {
        protos[i] = f.Name()
    }

    //ast, err := parser.ParseFiles([]string{"1.proto"}, []string{"./play/"})
    ast, err := parser.ParseFiles(protos, []string{DIR_PROTOS})

    gen := &GenOut{
        Messages: ExtractAllMessagesViews(ast),
        Services:ExtractAllServicesViews(ast),
        Enums: ExtractAllEnumsViews(ast),
    }

    build(gen)
}