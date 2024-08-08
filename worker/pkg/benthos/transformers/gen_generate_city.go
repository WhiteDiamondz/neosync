
// Code generated by Neosync neosync_transformer_generator.go. DO NOT EDIT.
// source: generate_city.go

package transformers

import (
	"fmt"
	
	transformer_utils "github.com/nucleuscloud/neosync/worker/pkg/benthos/transformers/utils"
	"github.com/nucleuscloud/neosync/worker/pkg/rng"
	
)

type GenerateCity struct{}

type GenerateCityOpts struct {
	randomizer     rng.Rand
	
	maxLength int64
}

func NewGenerateCity() *GenerateCity {
	return &GenerateCity{}
}

func NewGenerateCityOpts(
	maxLength int64,
  seedArg *int64,
) (*GenerateCityOpts, error) {
	seed, err := transformer_utils.GetSeedOrDefault(seedArg)
  if err != nil {
    return nil, fmt.Errorf("unable to generate seed: %w", err)
	}
	
	return &GenerateCityOpts{
		maxLength: maxLength,
		randomizer: rng.New(seed),	
	}, nil
}

func (t *GenerateCity) GetJsTemplateData() (*TemplateData, error) {
	return &TemplateData{
		Name: "generateCity",
		Description: "Randomly selects a city from a list of predefined US cities.",
		Example: "",
	}, nil
}

func (t *GenerateCity) ParseOptions(opts map[string]any) (any, error) {
	transformerOpts := &GenerateCityOpts{}

	if _, ok := opts["maxLength"].(int64); !ok {
		return nil, fmt.Errorf("missing required argument. function: %s argument: %s", "generateCity", "maxLength")
	}
	maxLength := opts["maxLength"].(int64)
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
