package configure

import (
	"github.com/lspaccatrosi16/aup/lib/types"
	"github.com/lspaccatrosi16/go-cli-tools/structconfig"
)

func Interactive(cfg *types.AUPData) error {
	cfgRunner := structconfig.NewConfig[types.AUPConfig]()

	cfgRunner.Run(cfg.Config)
	// cfgRunner.Debug()
	return nil
}
