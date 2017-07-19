package ant

import (
    "io/ioutil"
    "fmt"
    "ms/sun/helper"
    "github.com/dsymonds/gotoc/parser"
)

const DIR_OUTPUT = `C:\Go\_gopath\src\ms\sun\models\x\pb__gen_ant.go` //"./play/gen_sample_out.go"
const TEMPLATES_DIR  = "./templates/" //relative to main func of parent directory

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

