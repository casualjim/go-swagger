package generator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"text/template/parse"
	"unicode"

	"log"

	"github.com/go-openapi/inflect"
	"github.com/go-openapi/swag"
	"github.com/kr/pretty"
)

var (
	assets             map[string][]byte
	protectedTemplates map[string]bool

	// FuncMapFunc yields a map with all functions for templates
	FuncMapFunc func(*LanguageOpts) template.FuncMap

	templates *Repository
)

func initTemplateRepo() {
	FuncMapFunc = DefaultFuncMap

	// this makes the ToGoName func behave with the special
	// prefixing rule above
	swag.GoNamePrefixFunc = prefixForName

	assets = defaultAssets()
	protectedTemplates = defaultProtectedTemplates()
	templates = NewRepository(FuncMapFunc(DefaultLanguageFunc()))
}

// DefaultFuncMap yields a map with default functions for use n the templates.
// These are available in every template
func DefaultFuncMap(lang *LanguageOpts) template.FuncMap {
	return template.FuncMap(map[string]interface{}{
		"pascalize": pascalize,
		"camelize":  swag.ToJSONName,
		"varname":   lang.MangleVarName,
		"humanize":  swag.ToHumanNameLower,
		"snakize":   lang.MangleFileName,
		"toPackagePath": func(name string) string {
			return filepath.FromSlash(lang.ManglePackagePath(name, ""))
		},
		"toPackage": func(name string) string {
			return lang.ManglePackagePath(name, "")
		},
		"toPackageName": func(name string) string {
			return lang.ManglePackageName(name, "")
		},
		"dasherize":          swag.ToCommandName,
		"pluralizeFirstWord": pluralizeFirstWord,
		"json":               asJSON,
		"prettyjson":         asPrettyJSON,
		"hasInsecure": func(arg []string) bool {
			return swag.ContainsStringsCI(arg, "http") || swag.ContainsStringsCI(arg, "ws")
		},
		"hasSecure": func(arg []string) bool {
			return swag.ContainsStringsCI(arg, "https") || swag.ContainsStringsCI(arg, "wss")
		},
		"dropPackage":      dropPackage,
		"upper":            strings.ToUpper,
		"contains":         swag.ContainsStrings,
		"padSurround":      padSurround,
		"joinFilePath":     filepath.Join,
		"comment":          padComment,
		"blockcomment":     blockComment,
		"inspect":          pretty.Sprint,
		"cleanPath":        path.Clean,
		"mediaTypeName":    mediaMime,
		"arrayInitializer": lang.arrayInitializer,
		"hasPrefix":        strings.HasPrefix,
		"stringContains":   strings.Contains,
		"imports":          lang.imports,
	})
}

func defaultAssets() map[string][]byte {
	return map[string][]byte{
		// schema validation templates
		"validation/primitive.gotmpl":    MustAsset("templates/validation/primitive.gotmpl"),
		"validation/customformat.gotmpl": MustAsset("templates/validation/customformat.gotmpl"),
		"validation/structfield.gotmpl":  MustAsset("templates/validation/structfield.gotmpl"),
		"modelvalidator.gotmpl":          MustAsset("templates/modelvalidator.gotmpl"),
		"structfield.gotmpl":             MustAsset("templates/structfield.gotmpl"),
		"schemavalidator.gotmpl":         MustAsset("templates/schemavalidator.gotmpl"),
		"schemapolymorphic.gotmpl":       MustAsset("templates/schemapolymorphic.gotmpl"),

		// schema serialization templates
		"additionalpropertiesserializer.gotmpl": MustAsset("templates/serializers/additionalpropertiesserializer.gotmpl"),
		"aliasedserializer.gotmpl":              MustAsset("templates/serializers/aliasedserializer.gotmpl"),
		"allofserializer.gotmpl":                MustAsset("templates/serializers/allofserializer.gotmpl"),
		"basetypeserializer.gotmpl":             MustAsset("templates/serializers/basetypeserializer.gotmpl"),
		"marshalbinaryserializer.gotmpl":        MustAsset("templates/serializers/marshalbinaryserializer.gotmpl"),
		"schemaserializer.gotmpl":               MustAsset("templates/serializers/schemaserializer.gotmpl"),
		"subtypeserializer.gotmpl":              MustAsset("templates/serializers/subtypeserializer.gotmpl"),
		"tupleserializer.gotmpl":                MustAsset("templates/serializers/tupleserializer.gotmpl"),

		// schema generation template
		"docstring.gotmpl":  MustAsset("templates/docstring.gotmpl"),
		"schematype.gotmpl": MustAsset("templates/schematype.gotmpl"),
		"schemabody.gotmpl": MustAsset("templates/schemabody.gotmpl"),
		"schema.gotmpl":     MustAsset("templates/schema.gotmpl"),
		"model.gotmpl":      MustAsset("templates/model.gotmpl"),
		"header.gotmpl":     MustAsset("templates/header.gotmpl"),

		"swagger_json_embed.gotmpl": MustAsset("templates/swagger_json_embed.gotmpl"),

		// server templates
		"server/parameter.gotmpl":    MustAsset("templates/server/parameter.gotmpl"),
		"server/urlbuilder.gotmpl":   MustAsset("templates/server/urlbuilder.gotmpl"),
		"server/responses.gotmpl":    MustAsset("templates/server/responses.gotmpl"),
		"server/operation.gotmpl":    MustAsset("templates/server/operation.gotmpl"),
		"server/builder.gotmpl":      MustAsset("templates/server/builder.gotmpl"),
		"server/server.gotmpl":       MustAsset("templates/server/server.gotmpl"),
		"server/configureapi.gotmpl": MustAsset("templates/server/configureapi.gotmpl"),
		"server/main.gotmpl":         MustAsset("templates/server/main.gotmpl"),
		"server/doc.gotmpl":          MustAsset("templates/server/doc.gotmpl"),

		// client templates
		"client/parameter.gotmpl": MustAsset("templates/client/parameter.gotmpl"),
		"client/response.gotmpl":  MustAsset("templates/client/response.gotmpl"),
		"client/client.gotmpl":    MustAsset("templates/client/client.gotmpl"),
		"client/facade.gotmpl":    MustAsset("templates/client/facade.gotmpl"),
	}
}

func defaultProtectedTemplates() map[string]bool {
	return map[string]bool{
		"dereffedSchemaType":          true,
		"docstring":                   true,
		"header":                      true,
		"mapvalidator":                true,
		"model":                       true,
		"modelvalidator":              true,
		"objectvalidator":             true,
		"primitivefieldvalidator":     true,
		"privstructfield":             true,
		"privtuplefield":              true,
		"propertyValidationDocString": true,
		"propertyvalidator":           true,
		"schema":                      true,
		"schemaBody":                  true,
		"schemaType":                  true,
		"schemabody":                  true,
		"schematype":                  true,
		"schemavalidator":             true,
		"serverDoc":                   true,
		"slicevalidator":              true,
		"structfield":                 true,
		"structfieldIface":            true,
		"subTypeBody":                 true,
		"swaggerJsonEmbed":            true,
		"tuplefield":                  true,
		"tuplefieldIface":             true,
		"typeSchemaType":              true,
		"validationCustomformat":      true,
		"validationPrimitive":         true,
		"validationStructfield":       true,
		"withBaseTypeBody":            true,
		"withoutBaseTypeBody":         true,

		// all serializers TODO(fred)
		"additionalPropertiesSerializer": true,
		"tupleSerializer":                true,
		"schemaSerializer":               true,
		"hasDiscriminatedSerializer":     true,
		"discriminatedSerializer":        true,
	}
}

// AddFile adds a file to the default repository. It will create a new template based on the filename.
// It trims the .gotmpl from the end and converts the name using swag.ToJSONName. This will strip
// directory separators and Camelcase the next letter.
// e.g validation/primitive.gotmpl will become validationPrimitive
//
// If the file contains a definition for a template that is protected the whole file will not be added
func AddFile(name, data string) error {
	return templates.addFile(name, data, false)
}

// NewRepository creates a new template repository with the provided functions defined
func NewRepository(funcs template.FuncMap) *Repository {
	repo := Repository{
		files:     make(map[string]string),
		templates: make(map[string]*template.Template),
		funcs:     funcs,
	}

	if repo.funcs == nil {
		repo.funcs = make(template.FuncMap)
	}

	return &repo
}

// Repository is the repository for the generator templates
type Repository struct {
	files         map[string]string
	templates     map[string]*template.Template
	funcs         template.FuncMap
	allowOverride bool
}

// LoadDefaults will load the embedded templates
func (t *Repository) LoadDefaults() {

	for name, asset := range assets {
		if err := t.addFile(name, string(asset), true); err != nil {
			log.Fatal(err)
		}
	}
}

// LoadDir will walk the specified path and add each .gotmpl file it finds to the repository
func (t *Repository) LoadDir(templatePath string) error {
	err := filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {

		if strings.HasSuffix(path, ".gotmpl") {
			if assetName, e := filepath.Rel(templatePath, path); e == nil {
				if data, e := ioutil.ReadFile(path); e == nil {
					if ee := t.AddFile(assetName, string(data)); ee != nil {
						return fmt.Errorf("could not add template: %v", ee)
					}
				}
				// Non-readable files are skipped
			}
		}
		if err != nil {
			return err
		}
		// Non-template files are skipped
		return nil
	})
	if err != nil {
		return fmt.Errorf("could not complete template processing in directory \"%s\": %v", templatePath, err)
	}
	return nil
}

// LoadContrib loads template from contrib directory
func (t *Repository) LoadContrib(name string) error {
	log.Printf("loading contrib %s", name)
	const pathPrefix = "templates/contrib/"
	basePath := pathPrefix + name
	filesAdded := 0
	for _, aname := range AssetNames() {
		if !strings.HasSuffix(aname, ".gotmpl") {
			continue
		}
		if strings.HasPrefix(aname, basePath) {
			target := aname[len(basePath)+1:]
			err := t.addFile(target, string(MustAsset(aname)), true)
			if err != nil {
				return err
			}
			log.Printf("added contributed template %s from %s", target, aname)
			filesAdded++
		}
	}
	if filesAdded == 0 {
		return fmt.Errorf("no files added from template: %s", name)
	}
	return nil
}

func (t *Repository) addFile(name, data string, allowOverride bool) error {
	fileName := name
	name = swag.ToJSONName(strings.TrimSuffix(name, ".gotmpl"))

	templ, err := template.New(name).Funcs(t.funcs).Parse(data)

	if err != nil {
		return fmt.Errorf("failed to load template %s: %v", name, err)
	}

	// check if any protected templates are defined
	if !allowOverride && !t.allowOverride {
		for _, template := range templ.Templates() {
			if protectedTemplates[template.Name()] {
				return fmt.Errorf("cannot overwrite protected template %s", template.Name())
			}
		}
	}

	// Add each defined template into the cache
	for _, template := range templ.Templates() {

		t.files[template.Name()] = fileName
		t.templates[template.Name()] = template.Lookup(template.Name())
	}

	return nil
}

// MustGet a template by name, panics when fails
func (t *Repository) MustGet(name string) *template.Template {
	tpl, err := t.Get(name)
	if err != nil {
		panic(err)
	}
	return tpl
}

// AddFile adds a file to the repository. It will create a new template based on the filename.
// It trims the .gotmpl from the end and converts the name using swag.ToJSONName. This will strip
// directory separators and Camelcase the next letter.
// e.g validation/primitive.gotmpl will become validationPrimitive
//
// If the file contains a definition for a template that is protected the whole file will not be added
func (t *Repository) AddFile(name, data string) error {
	return t.addFile(name, data, false)
}

// SetAllowOverride allows setting allowOverride after the Repository was initialized
func (t *Repository) SetAllowOverride(value bool) {
	t.allowOverride = value
}

func findDependencies(n parse.Node) []string {

	var deps []string
	depMap := make(map[string]bool)

	if n == nil {
		return deps
	}

	switch node := n.(type) {
	case *parse.ListNode:
		if node != nil && node.Nodes != nil {
			for _, nn := range node.Nodes {
				for _, dep := range findDependencies(nn) {
					depMap[dep] = true
				}
			}
		}
	case *parse.IfNode:
		for _, dep := range findDependencies(node.BranchNode.List) {
			depMap[dep] = true
		}
		for _, dep := range findDependencies(node.BranchNode.ElseList) {
			depMap[dep] = true
		}

	case *parse.RangeNode:
		for _, dep := range findDependencies(node.BranchNode.List) {
			depMap[dep] = true
		}
		for _, dep := range findDependencies(node.BranchNode.ElseList) {
			depMap[dep] = true
		}

	case *parse.WithNode:
		for _, dep := range findDependencies(node.BranchNode.List) {
			depMap[dep] = true
		}
		for _, dep := range findDependencies(node.BranchNode.ElseList) {
			depMap[dep] = true
		}

	case *parse.TemplateNode:
		depMap[node.Name] = true
	}

	for dep := range depMap {
		deps = append(deps, dep)
	}

	return deps

}

func (t *Repository) flattenDependencies(templ *template.Template, dependencies map[string]bool) map[string]bool {
	if dependencies == nil {
		dependencies = make(map[string]bool)
	}

	deps := findDependencies(templ.Tree.Root)

	for _, d := range deps {
		if _, found := dependencies[d]; !found {

			dependencies[d] = true

			if tt := t.templates[d]; tt != nil {
				dependencies = t.flattenDependencies(tt, dependencies)
			}
		}

		dependencies[d] = true

	}

	return dependencies

}

func (t *Repository) addDependencies(templ *template.Template) (*template.Template, error) {

	name := templ.Name()

	deps := t.flattenDependencies(templ, nil)

	for dep := range deps {

		if dep == "" {
			continue
		}

		tt := templ.Lookup(dep)

		// Check if we have it
		if tt == nil {
			tt = t.templates[dep]

			// Still don't have it, return an error
			if tt == nil {
				return templ, fmt.Errorf("could not find template %s", dep)
			}
			var err error

			// Add it to the parse tree
			templ, err = templ.AddParseTree(dep, tt.Tree)

			if err != nil {
				return templ, fmt.Errorf("dependency error: %v", err)
			}

		}
	}
	return templ.Lookup(name), nil
}

// Get will return the named template from the repository, ensuring that all dependent templates are loaded.
// It will return an error if a dependent template is not defined in the repository.
func (t *Repository) Get(name string) (*template.Template, error) {
	templ, found := t.templates[name]

	if !found {
		return templ, fmt.Errorf("template doesn't exist %s", name)
	}

	return t.addDependencies(templ)
}

// DumpTemplates prints out a dump of all the defined templates, where they are defined and what their dependencies are.
func (t *Repository) DumpTemplates() {
	buf := bytes.NewBuffer(nil)
	fmt.Fprintln(buf, "\n# Templates")
	for name, templ := range t.templates {
		fmt.Fprintf(buf, "## %s\n", name)
		fmt.Fprintf(buf, "Defined in `%s`\n", t.files[name])

		if deps := findDependencies(templ.Tree.Root); len(deps) > 0 {

			fmt.Fprintf(buf, "####requires \n - %v\n\n\n", strings.Join(deps, "\n - "))
		}
		fmt.Fprintln(buf, "\n---")
	}
	log.Println(buf.String())
}

// FuncMap functions

func asJSON(data interface{}) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func asPrettyJSON(data interface{}) (string, error) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func pluralizeFirstWord(arg string) string {
	sentence := strings.Split(arg, " ")
	if len(sentence) == 1 {
		return inflect.Pluralize(arg)
	}

	return inflect.Pluralize(sentence[0]) + " " + strings.Join(sentence[1:], " ")
}

func dropPackage(str string) string {
	parts := strings.Split(str, ".")
	return parts[len(parts)-1]
}

func padSurround(entry, padWith string, i, ln int) string {
	var res []string
	if i > 0 {
		for j := 0; j < i; j++ {
			res = append(res, padWith)
		}
	}
	res = append(res, entry)
	tot := ln - i - 1
	for j := 0; j < tot; j++ {
		res = append(res, padWith)
	}
	return strings.Join(res, ",")
}

func padComment(str string, pads ...string) string {
	// pads specifes padding to indent multi line comments.Defaults to one space
	pad := " "
	lines := strings.Split(str, "\n")
	if len(pads) > 0 {
		pad = strings.Join(pads, "")
	}
	return (strings.Join(lines, "\n//"+pad))
}

func blockComment(str string) string {
	return strings.Replace(str, "*/", "[*]/", -1)
}

func pascalize(arg string) string {
	runes := []rune(arg)
	switch len(runes) {
	case 0:
		return "Empty"
	case 1: // handle special case when we have a single rune that is not handled by swag.ToGoName
		switch runes[0] {
		case '+', '-', '#', '_': // those cases are handled differently than swag utility
			return prefixForName(arg)
		}
	}
	return swag.ToGoName(swag.ToGoName(arg)) // want to remove spaces
}

func prefixForName(arg string) string {
	first := []rune(arg)[0]
	if len(arg) == 0 || unicode.IsLetter(first) {
		return ""
	}
	switch first {
	case '+':
		return "Plus"
	case '-':
		return "Minus"
	case '#':
		return "HashTag"
		// other cases ($,@ etc..) handled by swag.ToGoName
	}
	return "Nr"
}
