package env

import (
	"errors"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func MustLoadEnvDefault() {
	MustLoadEnv(".env",
		".env.production",
		".env.real",
		".env.beta",
		".env.development",
		".env.dev",
		".env.local",
	)
}

func MustLoadEnv(files ...string) {
	loaded := make([]string, 0, len(files))
	for _, file := range files {
		i, e := os.Stat(file)
		if errors.Is(e, os.ErrNotExist) {
			continue
		}
		if e != nil {
			//stderr.Log(stderr.WRN, "skip", "env", file, "error", e)
			log.Println("skip", "env", file, "error", e)
			continue
		}

		if err := godotenv.Overload(file); err != nil {
			//stderr.Log(stderr.ERR, "overload", "name", i.Name(), "size", i.Size(), "mode", i.Mode(), "error", err)
			log.Println("overload name", i.Name(), "size", i.Size(), "mode", i.Mode(), "error", err)
			panic(err)
		}
		loaded = append(loaded, file)
	}

	//stderr.Log(stderr.DBG, "load env", "files", loaded)
	log.Println("load env files", loaded)
}

func MustFetchEnv[V any](out *V, prefix ...string) *V {
	opts := env.Options{
		Prefix:                strings.Join(prefix, "_"),
		UseFieldNameByDefault: true,
		RequiredIfNoDef:       true,
	}
	if err := env.ParseWithOptions(out, opts); err != nil {
		panic(err)
	}
	return out
}
