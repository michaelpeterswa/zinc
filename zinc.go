package zinc

import (
	"context"
	"errors"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"go.uber.org/zap"
)

func InitLogger(logLevel string) (*zap.Logger, error) {
	var (
		logger *zap.Logger
		err    error
	)

	if logLevel == "dev" {
		logger, err = zap.NewDevelopment()
		if err != nil {
			return nil, err
		}
	} else if logLevel == "prod" {
		logger, err = zap.NewProduction()
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("invalid log level")
	}

	return logger, nil
}

func InitInfluxDB(ctx context.Context, url string, token string) (influxdb2.Client, error) {
	influx := influxdb2.NewClient(url, token)
	ping, err := influx.Ping(ctx)
	if err != nil {
		return nil, err
	}
	if !ping {
		return nil, errors.New("influx ping failed")
	}
	return influx, nil
}
