
// Code generated by Neosync neosync_transformer_generator.go. DO NOT EDIT.
// source: generate_username.go

package transformers

import (
	"fmt"
	
	transformer_utils "github.com/nucleuscloud/neosync/worker/pkg/benthos/transformers/utils"
	"github.com/nucleuscloud/neosync/worker/pkg/rng"
	
)

type GenerateUsername struct{}

type GenerateUsernameOpts struct {
	randomizer     rng.Rand
	
	maxLength int64
}

func NewGenerateUsername() *GenerateUsername {
	return &GenerateUsername{}
}

func NewGenerateUsernameOpts(
	maxLengthArg *int64,
  seedArg *int64,
) (*GenerateUsernameOpts, error) {
	maxLength := int64(10000) 
	if maxLengthArg != nil {
		maxLength = *maxLengthArg
	}
	
	seed, err := transformer_utils.GetSeedOrDefault(seedArg)
  if err != nil {
    return nil, fmt.Errorf("unable to generate seed: %w", err)
	}
	
	return &GenerateUsernameOpts{
		maxLength: maxLength,
		randomizer: rng.New(seed),	
	}, nil
}

func (t *GenerateUsername) GetJsTemplateData() (*TemplateData, error) {
	return &TemplateData{
		Name: "generateUsername",
		Description: "Randomly generates a username",
		Example: "",
	}, nil
}

func (t *GenerateUsername) ParseOptions(opts map[string]any) (any, error) {
	transformerOpts := &GenerateUsernameOpts{}

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
