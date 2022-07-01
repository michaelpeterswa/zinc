package zinc

import (
	"context"
	"errors"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"go.uber.org/zap"
)

func InitLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
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
