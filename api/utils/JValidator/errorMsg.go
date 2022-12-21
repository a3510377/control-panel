package JValidator

// {0} Field
//
// {1} Param
//
// {2} Tag
//
// {3} Value
//
// {4} Kind
//
// {5} Type
//
// {6} Namespace
//
// {7} StructNamespace
//
// {8} StructField
//
// {9} ActualTag
var errorMsg = map[string]string{ // add else https://github.com/go-playground/validator#fields
	"required": "{0} is required",
	"max":      "{0} must be at most {1} long",
	"min":      "{0} must be at least {1} long",
	"email":    "{0} must be a valid email address",
}
