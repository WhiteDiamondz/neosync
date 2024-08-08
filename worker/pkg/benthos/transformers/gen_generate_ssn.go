
// Code generated by Neosync neosync_transformer_generator.go. DO NOT EDIT.
// source: generate_ssn.go

package transformers

import (
	"fmt"
	
	transformer_utils "github.com/nucleuscloud/neosync/worker/pkg/benthos/transformers/utils"
	"github.com/nucleuscloud/neosync/worker/pkg/rng"
	
)

type GenerateSSN struct{}

type GenerateSSNOpts struct {
	randomizer     rng.Rand
	
}

func NewGenerateSSN() *GenerateSSN {
	return &GenerateSSN{}
}

func NewGenerateSSNOpts(
  seedArg *int64,
) (*GenerateSSNOpts, error) {
	seed, err := transformer_utils.GetSeedOrDefault(seedArg)
  if err != nil {
    return nil, fmt.Errorf("unable to generate seed: %w", err)
	}
	
	return &GenerateSSNOpts{
		randomizer: rng.New(seed),	
	}, nil
}

func (t *GenerateSSN) GetJsTemplateData() (*TemplateData, error) {
	return &TemplateData{
		Name: "generateSSN",
		Description: "Generates a completely random social security numbers including the hyphens in the format xxx-xx-xxxx.",
		Example: "",
	}, nil
}

func (t *GenerateSSN) ParseOptions(opts map[string]any) (any, error) {
	transformerOpts := &GenerateSSNOpts{}

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
