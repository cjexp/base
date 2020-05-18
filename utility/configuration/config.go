package configuration

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/cjtoolkit/ctx"
	"github.com/cjexp/base/utility/environment"
	"github.com/cjexp/base/utility/loggers"
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

func GetConfig(context ctx.BackgroundContext) Base {
	type c struct{}
	return context.Persist(c{}, func() (interface{}, error) {
		config := &Base{}
		ParseConfig(context, "base.json", config)
		return *config, nil
	}).(Base)
}

func ParseConfig(context ctx.BackgroundContext, fileName string, v interface{}) {
	location := environment.GetEnvironment(context).ParseConfigDirectory() + filepath.FromSlash("/"+fileName)
	errorService := loggers.GetBlankErrorService(context)

	file, err := os.Open(location)
	errorService.CheckErrorAndPanic(err)
	defer file.Close()

	err = json.NewDecoder(file).Decode(v)
	errorService.CheckErrorAndPanic(err)
}