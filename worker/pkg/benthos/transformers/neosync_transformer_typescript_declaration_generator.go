//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"go/token"
	"os"
	"text/template"

	transformers "github.com/nucleuscloud/neosync/worker/pkg/benthos/transformers"
)

func main() {
	args := os.Args
	if len(args) < 1 {
		panic("must provide necessary args")
	}

	docsPath := args[1]

	fileSet := token.NewFileSet()
	transformerSpecs, err := transformers.ExtractBenthosSpec(fileSet)
	if err != nil {
		fmt.Println("Error finding transformer bloblang specs:", err)
		return
	}

	for _, tf := range transformerSpecs {
		parsedSpec, err := transformers.ParseBloblangSpec(tf)
		if err != nil {
			fmt.Println("Error parsing bloblang params:", err)
		}
		tf.Params = parsedSpec.Params
		tf.Description = parsedSpec.SpecDescription
	}

	tsDeclarationSpec := []*tsDeclarationSpec{}
	for _, tf := range transformerSpecs {
		ts, err := toDeclarationSpec(tf)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		tsDeclarationSpec = append(tsDeclarationSpec, ts)
	}

	codeStr, err := generateTypescriptDeclaration(tsDeclarationSpec)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}

	outputFile, err := os.Create(docsPath)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}

	_, err = outputFile.WriteString(codeStr)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return
	}
	outputFile.Close()

}

type tsDeclarationSpec struct {
	Name            string
	SpecType        string
	Description     string
	Params          []*tsDeclarationParam
	InterfaceName   string
	TsReturnTypeStr string
}

type tsDeclarationParam struct {
	Name        string
	IsOptional  bool
	HasDefault  bool
	Default     string
	Description string
	TsTypeStr   string
}

func toDeclarationSpec(benthosSpec *transformers.BenthosSpec) (*tsDeclarationSpec, error) {
	tsParams := []*tsDeclarationParam{}
	for _, p := range benthosSpec.Params {
		tsTypeStr, err := convertToTypescriptType(p.TypeStr, p.Default)
		if err != nil {
			return nil, err
		}
		tsParams = append(tsParams, &tsDeclarationParam{
			Name:        p.Name,
			IsOptional:  p.IsOptional,
			HasDefault:  p.HasDefault,
			Default:     p.Default,
			Description: p.Description,
			TsTypeStr:   tsTypeStr,
		})
	}

	return &tsDeclarationSpec{
		Name:            benthosSpec.Name,
		SpecType:        benthosSpec.Type,
		Description:     benthosSpec.Description,
		Params:          tsParams,
		InterfaceName:   fmt.Sprintf("%sOptions", transformers.CapitalizeFirst(benthosSpec.Name)),
		TsReturnTypeStr: "any",
	}, nil
}

func convertToTypescriptType(typeStr, defaultStr string) (string, error) {
	switch typeStr {
	case "string":
		return "string", nil
	case "int64":
		return "number", nil
	case "float64":
		return "number", nil
	case "int32":
		return "number", nil
	case "bool":
		return "boolean", nil
	case "any":
		if defaultStr == "[]any{}" {
			return "any[]", nil
		}
		return "any", nil
	default:
		return "", fmt.Errorf("Error: unable to convert type to typescript type: %s", typeStr)
	}
}

const docTemplate = `
/* prettier-ignore */
/**
 * Code generated by Neosync neosync_transformer_typescript_declaration_generator.go. DO NOT EDIT.
 */

/* eslint @typescript-eslint/no-explicit-any: 0 */

declare namespace neosync {

  /**
	 * Transformers
   */
	{{ range $i, $spec := .TransformerSpecs }}
	export interface {{$spec.InterfaceName}} {
	{{- range $j, $param := $spec.Params }}
	{{- if eq $param.Name "value" -}}{{ continue }}{{- end }}
		/** {{$param.Description}} */
		{{$param.Name}}{{ if $param.IsOptional }}?{{ else if $param.HasDefault }}?{{end}}: {{$param.TsTypeStr}};
	{{- end }}
	}
 
  /**
   * {{$spec.Description}}
   */
	declare function {{$spec.Name}}(value: any, options: {{$spec.InterfaceName}}): {{$spec.TsReturnTypeStr}};

	{{ end }}

	
  /**
	 * Generators 
   */
	{{ range $i, $spec := .GeneratorSpecs }}
	export interface {{$spec.InterfaceName}} {
	{{- range $j, $param := $spec.Params}}
	{{- if eq $param.Name "value" -}}{{ continue }}{{- end }}
		/** {{$param.Description}} */
		{{$param.Name}}{{ if $param.IsOptional }}?{{ else if $param.HasDefault }}?{{end}}: {{$param.TsTypeStr}};
	{{- end }}
	}

  /**
   * {{$spec.Description}}
   */
	declare function {{$spec.Name}}(options: {{$spec.InterfaceName}}): {{$spec.TsReturnTypeStr}};
	
	{{ end }}
}`

type TemplateData struct {
	TransformerSpecs []*tsDeclarationSpec
	GeneratorSpecs   []*tsDeclarationSpec
}

func generateTypescriptDeclaration(specs []*tsDeclarationSpec) (string, error) {
	transformerSpecs := []*tsDeclarationSpec{}
	generatorSpecs := []*tsDeclarationSpec{}

	for _, spec := range specs {
		if spec.SpecType == "transform" {
			transformerSpecs = append(transformerSpecs, spec)
		} else if spec.SpecType == "generate" {
			generatorSpecs = append(generatorSpecs, spec)
		}
	}
	data := TemplateData{
		TransformerSpecs: transformerSpecs,
		GeneratorSpecs:   generatorSpecs,
	}
	t := template.Must(template.New("neosyncTransformerDocs").Parse(docTemplate))
	var out bytes.Buffer
	err := t.Execute(&out, data)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
