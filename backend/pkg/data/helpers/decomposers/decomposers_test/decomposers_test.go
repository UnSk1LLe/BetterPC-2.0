package decomposers_test

import (
	"BetterPC_2.0/pkg/data/helpers/decomposers"
	productRequests "BetterPC_2.0/pkg/data/models/products/requests"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type Inputs struct {
	FullInput                     productRequests.UpdateCpuRequest `json:"full_input"`
	NilPointerInput               productRequests.UpdateCpuRequest `json:"nil_pointer_input"`
	EmbedStructureNilPointerInput productRequests.UpdateCpuRequest `json:"embed_structure_nil_pointer_input"`
	NilValueInput                 productRequests.UpdateCpuRequest `json:"nil_value_input"`
}

func parseJsonData() (Inputs, error) {
	var inputs Inputs
	filePath := "./test_inputs.json"

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return inputs, err
	}

	// Unmarshal JSON into UpdateCpuRequest structure
	err = json.Unmarshal(jsonData, &inputs)
	if err != nil {
		return inputs, err
	}

	return inputs, err
}

func TestDecomposeWithTag(t *testing.T) {
	inputs, err := parseJsonData()
	if err != nil {
		panic(err)
	}

	t.Run("Test nil pointer", func(t *testing.T) {

		result, err := decomposers.DecomposeWithTag(&inputs.NilPointerInput, "json")

		assert.NoError(t, err)

		expected := map[string]interface{}{
			"ram.channels":      5,
			"ram.max_frequency": []int{5555, 0},
			"ram.max_capacity":  128,
		}

		assert.Equal(t, expected, result)
	})

	t.Run("Test embed struct nil pointer", func(t *testing.T) {

		result, err := decomposers.DecomposeWithTag(&inputs.EmbedStructureNilPointerInput, "json")

		assert.NoError(t, err)

		expected := map[string]interface{}{
			"ram.channels":     5,
			"ram.max_capacity": 128,
			"pci_e":            5,
		}

		assert.Equal(t, expected, result)
	})

	t.Run("Test nil value", func(t *testing.T) {

		result, err := decomposers.DecomposeWithTag(&inputs.NilValueInput, "json")

		assert.NoError(t, err)

		expected := map[string]interface{}{
			"general.discount": 0,
			"ram.channels":     5,
			"ram.max_capacity": 128,
			"graphics":         "",
			"pci_e":            5,
		}

		assert.Equal(t, expected, result)
	})

	t.Run("Test full input", func(t *testing.T) {

		result, err := decomposers.DecomposeWithTag(&inputs.FullInput, "json")
		assert.NoError(t, err)

		expected := map[string]interface{}{
			"general.manufacturer": "Intel",
			"general.model":        "OTHER MODEL",
			"general.price":        555555,
			"general.discount":     0,
			"general.amount":       5,
			"general.image":        "/assets/img/default/cpu.jpg",

			"main.category":   "Intel Core i7",
			"main.generation": "Intel 10th",
			"main.socket":     "LGA 1200",
			"main.year":       2020,

			"clock_frequency.p-cores":         []float64{5.5, 5.5},
			"clock_frequency.e-cores":         []float64{5, 5},
			"clock_frequency.free_multiplier": true,

			"ram.channels":      5,
			"ram.max_frequency": []int{5555, 0},
			"ram.max_capacity":  128,

			"tdp":             555,
			"graphics":        "",
			"pci_e":           5,
			"max_temperature": 100,
		}

		assert.Equal(t, expected, result)
	})
}
