package make

import (
	"embed"
	"fmt"
	"gohub/pkg/console"
	"gohub/pkg/file"
	"gohub/pkg/str"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

// 整理好的数据：
// {
//     "TableName": "topic_comments",
//     "StructName": "TopicComment",
//     "StructNamePlural": "TopicComments"
//     "VariableName": "topicComment",
//     "VariableNamePlural": "topicComments",
//     "PackageName": "topic_comment"
// }
type Model struct {
	TableName          string
	StructName         string
	StructNamePlural   string
	VariableName       string
	VariableNamePlural string
	PackageName        string
}

//打包 .stub后缀的文件
var stubsFS embed.FS

var CmdMake = &cobra.Command{
	Use:   "make",
	Short: "Generate file and code",
}

func init() {
	CmdMake.AddCommand(
		CmdMakeCMD,
	)
}

func makeModelFromString(name string) Model {
	model := Model{}
	model.StructName = str.Singular(strcase.ToCamel(name))
	model.StructNamePlural = str.Plural(strcase.ToCamel(name))
	model.TableName = str.Snake(model.StructNamePlural)
	model.VariableName = str.LowerCamel(model.StructName)
	model.VariableNamePlural = str.LowerCamel(model.StructNamePlural)
	model.PackageName = str.Snake(model.StructName)
	return model
}

//读取 stub 文件并进行变量替换
func createFileFromStub(filePath string, stubName string, model Model, variables ...interface{}) {
	replaces := make(map[string]string)
	if len(variables) > 0 {
		replaces = variables[0].(map[string]string)
	}

	//目录文件已存在
	if file.Exits(filePath) {
		console.Exit(filePath + " already exists!")
	}
	stubFile := "stubs/" + stubName + ".stub"
	//读取 stub 模板文件
	modelData, err := stubsFS.ReadFile(stubFile)
	console.ExitIf(err)
	modelStub := string(modelData)

	//添加替换变量
	replaces["{{VariableName}}"] = model.VariableName
	replaces["{{VariableNamePlural}}"] = model.VariableNamePlural
	replaces["{{StructName}}"] = model.StructName
	replaces["{{StructNamePlural}}"] = model.StructNamePlural
	replaces["{{PackageName}}"] = model.PackageName
	replaces["{{TableName}}"] = model.TableName

	for search, replace := range replaces {
		modelStub = strings.ReplaceAll(modelStub, search, replace)
	}

	//存储到文件中
	err = file.Put([]byte(modelStub), filePath)
	console.ExitIf(err)

	console.Success(fmt.Sprintf("[%s] created.", filePath))
}
