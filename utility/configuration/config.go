package configuration

import (
	"os"
	"path/filepath"

	"github.com/cjexp/base/utility/environment"
	"github.com/cjexp/base/utility/loggers"
	"github.com/cjtoolkit/ctx/v2"
	"github.com/pelletier/go-toml"
)

type Base struct {
	Database     Database
	CsrfKey      string
	CookieKey    string
	PasswordSalt string
	HmacKey      string
}

type Database struct {
	MainSqlDsn string
	MongoDial  string
	MongoDb    string
	Redis      Redis
}

type Redis struct {
	Addr string
}

func GetConfig(context ctx.Context) Base {
	type c struct{}
	return context.Persist(c{}, func() (interface{}, error) {
		config := &Base{}
		ParseConfig(context, "base.toml", config)
		return *config, nil
	}).(Base)
}

func ParseConfig(context ctx.Context, fileName string, v interface{}) {
	location := environment.GetEnvironment(context).ParseConfigDirectory() + filepath.FromSlash("/"+fileName)
	errorService := loggers.GetBlankErrorService(context)

	file, err := os.Open(location)
	errorService.CheckErrorAndPanic(err)
	defer file.Close()

	err = toml.NewDecoder(file).Decode(v)
	errorService.CheckErrorAndPanic(err)
}
