package initlize

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"github.com/sunp13/dbtool"
)

func initDB() error {
	err := dbtool.Init("./conf/datasource.yml")
	if err != nil {
		return err
	}
	log.Info().Msg("init_db succ")
	return nil
}
