package str

import (
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

//转为复数 user->users
func Plural(word string) string {
	return pluralize.NewClient().Plural(word)
}

//复数转单数 users->user
func Singular(word string) string {
	return pluralize.NewClient().Singular(word)
}

//转为 snake_case
func Snake(s string) string {
	return strcase.ToSnake(s)
}

//转为 CamelCase
func Camel(s string) string {
	return strcase.ToCamel(s)
}

//转为camelCase
func LowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}
