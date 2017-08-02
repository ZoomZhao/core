package gitlab

import (
	"gitlab.ricebook.net/platform/core/source/common"
	"gitlab.ricebook.net/platform/core/types"
)

func New(config types.GitConfig) *common.GitScm {
	authheaders := map[string]string{}
	authheaders["PRIVATE-TOKEN"] = config.Token
	return &common.GitScm{Config: config, AuthHeaders: authheaders}
}
