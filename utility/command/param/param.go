package param

import (
	"os"

	"github.com/cjexp/base/utility/command"
	"github.com/cjtoolkit/ctx/v2"
)

type Param struct {
	Address    string
	Production bool
	TestRun    bool
}

func GetParam(context ctx.Context) Param {
	type c struct{}
	return context.Persist(c{}, func() (interface{}, error) {
		return initParam(), nil
	}).(Param)
}

func initParam() Param {
	param := &Param{
		Address:    ":8080",
		Production: false,
		TestRun:    false,
	}

	options := command.BuildOptions(os.Args[1:])

	options.ExecOption("address", func(strings []string) {
		param.Address = strings[0]
	})
	options.ExecOption("prod", func(_ []string) {
		param.Production = true
	})
	options.ExecOption("test-run", func(_ []string) {
		param.TestRun = true
	})

	return *param
}

func CheckIfTestRun(param Param) {
	if param.TestRun {
		os.Exit(0)
	}
}
