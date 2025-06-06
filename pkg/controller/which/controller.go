package which

import (
	"context"
	"io"
	"os"

	"github.com/aquaproj/aqua/v2/pkg/checksum"
	"github.com/aquaproj/aqua/v2/pkg/config"
	"github.com/aquaproj/aqua/v2/pkg/config/aqua"
	"github.com/aquaproj/aqua/v2/pkg/config/registry"
	"github.com/aquaproj/aqua/v2/pkg/runtime"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/go-osenv/osenv"
)

type Controller struct {
	stdout            io.Writer
	rootDir           string
	configFinder      ConfigFinder
	configReader      ConfigReader
	registryInstaller RegistryInstaller
	runtime           *runtime.Runtime
	osenv             osenv.OSEnv
	fs                afero.Fs
	linker            Linker
}

func New(param *config.Param, configFinder ConfigFinder, configReader ConfigReader, registryInstaller RegistryInstaller, rt *runtime.Runtime, osEnv osenv.OSEnv, fs afero.Fs, linker Linker) *Controller {
	return &Controller{
		stdout:            os.Stdout,
		rootDir:           param.RootDir,
		configFinder:      configFinder,
		configReader:      configReader,
		registryInstaller: registryInstaller,
		runtime:           rt,
		osenv:             osEnv,
		fs:                fs,
		linker:            linker,
	}
}

type Linker interface {
	Lstat(s string) (os.FileInfo, error)
	Symlink(dest, src string) error
	Readlink(src string) (string, error)
}

type ConfigReader interface {
	Read(logE *logrus.Entry, configFilePath string, cfg *aqua.Config) error
}

type MockController struct {
	FindResult *FindResult
	Err        error
}

type RegistryInstaller interface {
	InstallRegistries(ctx context.Context, logE *logrus.Entry, cfg *aqua.Config, cfgFilePath string, checksums *checksum.Checksums) (map[string]*registry.Config, error)
	InstallRegistry(ctx context.Context, logE *logrus.Entry, regist *aqua.Registry, cfgFilePath string, checksums *checksum.Checksums) (*registry.Config, error)
}
