package raw

import (
	"github.com/giongto35/cloud-game/v2/pkg/worker/emulator/libretro"
)

type Repo struct {
	Address     string
	Compression string
}

// NewRawRepo defines a simple zip file containing
// all the cores that will be extracted as is.
func NewRawRepo(address string) Repo {
	return Repo{Address: address, Compression: "zip"}
}

func (r Repo) GetCoreUrl(_ string, _ libretro.ArchInfo) string {
	return r.Address
}
