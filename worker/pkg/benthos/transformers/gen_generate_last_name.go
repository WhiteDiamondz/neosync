
// Code generated by Neosync neosync_transformer_generator.go. DO NOT EDIT.
// source: generate_last_name.go

package transformers

import (
	"fmt"
	
	transformer_utils "github.com/nucleuscloud/neosync/worker/pkg/benthos/transformers/utils"
	"github.com/nucleuscloud/neosync/worker/pkg/rng"
	
)

type GenerateLastName struct{}

type GenerateLastNameOpts struct {
	randomizer     rng.Rand
	
	maxLength int64
}

func NewGenerateLastName() *GenerateLastName {
	return &GenerateLastName{}
}

func NewGenerateLastNameOpts(
	maxLengthArg *int64,
  seedArg *int64,
) (*GenerateLastNameOpts, error) {
	maxLength := int64(10000) 
	if maxLengthArg != nil {
		maxLength = *maxLengthArg
	}
	
	seed, err := transformer_utils.GetSeedOrDefault(seedArg)
  if err != nil {
    return nil, fmt.Errorf("unable to generate seed: %w", err)
	}
	
	return &GenerateLastNameOpts{
		maxLength: maxLength,
		randomizer: rng.New(seed),	
	}, nil
}

func (t *GenerateLastName) GetJsTemplateData() (*TemplateData, error) {
	return &TemplateData{
		Name: "generateLastName",
		Description: "Generates a random last name.",
		Example: "",
	}, nil
}

func (t *GenerateLastName) ParseOptions(opts map[string]any) (any, error) {
	transformerOpts := &GenerateLastNameOpts{}

	maxLength, ok := opts["maxLength"].(int64)
	if !ok {
		maxLength = 10000
	}
	transformerOpts.maxLength = maxLength

	var seedArg *int64
	if seedValue, ok := opts["seed"].(int64); ok {
			seedArg = &seedValue
	}
	seed, err := transformer_utils.GetSeedOrDefault(seedArg)
	if err != nil {
		return nil, fmt.Errorf("unable to generate seed: %w", err)
	}
	transformerOpts.randomizer = rng.New(seed)

	return transformerOpts, nil
}
