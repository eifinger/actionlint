// Code generated by actionlint/scripts/generate-popular-actions. DO NOT EDIT.

package actionlint

// PopularActions is data set of known popular actions. Keys are specs (owner/repo@ref) of actions
// and values are their metadata.
var PopularActions = map[string]*ActionMetadata{
	"rhysd/action-setup-vim@v1.2.7": {
		Name: "Setup Vim",
		Inputs: ActionMetadataInputs{
			"neovim":  {"neovim", false},
			"token":   {"token", false},
			"version": {"version", false},
		},
		Outputs: ActionMetadataOutputs{
			"executable": {"executable"},
		},
	},
	"rhysd/changelog-from-release/action@v2.2.2": {
		Name: "Run changelog-from-release",
		Inputs: ActionMetadataInputs{
			"commit":       {"commit", false},
			"file":         {"file", true},
			"github_token": {"github_token", true},
			"push":         {"push", false},
			"version":      {"version", false},
		},
	},
}
