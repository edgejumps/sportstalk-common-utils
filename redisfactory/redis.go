package redisfactory

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	ErrInvalidRedisClientConfig = errors.New("Invalid Redis client config")
	DefaultServerAddr           = ":6379"
	DefaultKey                  = "sportstalk"
)

func getFields(v reflect.Value, fieldNames ...string) (bool, map[string]string) {
	result := map[string]string{}
	for _, fieldName := range fieldNames {
		f := v.FieldByName(fieldName)
		if !f.IsValid() {
			fmt.Printf("%s field is not valid !\n", fieldName)
			return false, nil
		}
		if (f.Kind() == reflect.String || f.Kind() == reflect.Int32 || f.Kind() == reflect.Int64) && f.String() != "" {
			result[fieldName] = f.String()
		} else {
			return false, nil
		}
	}

	return true, result
}

// RedisClientFactory creates a redis client based on the config
func RedisClientFactory(ctx context.Context, config interface{}) (*redis.Client, error) {
	v := reflect.ValueOf(config)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Check if the value is a struct
	if v.Kind() != reflect.Struct {
		return nil, ErrInvalidRedisClientConfig
	}

	if ok, m := getFields(v, "SentinelMasterName", "SentinelAddrs", "SentinelUsername", "SentinelPassword", "Username", "Password"); ok {
		senAddrs := strings.Split(m["SentinelAddrs"], ",")
		if senAddrs == nil || len(senAddrs) == 0 {
			return nil, fmt.Errorf("invalid sentinel addrs: %s", m["SentinelAddrs"])
		}

		client := redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:       m["SentinelMasterName"],
			SentinelAddrs:    senAddrs,
			SentinelUsername: m["SentinelUsername"],
			SentinelPassword: m["SentinelPassword"],
			Username:         m["Username"],
			Password:         m["Password"],
			DB:               0, // Use default DB
		})

		return client, nil
	}

	if ok, m := getFields(v, "ConnectionString"); ok {
		opt, err := redis.ParseURL(m["ConnectionString"])
		if err != nil {
			return nil, err
		}
		return redis.NewClient(opt), nil
	}

	addr := DefaultServerAddr
	if ok, m := getFields(v, "Addr"); ok {
		addr = m["Addr"]

		opt := &redis.Options{
			Addr: addr,
		}

		if ok, m := getFields(v, "Password"); ok {
			opt.Password = m["Password"]
		}

		if ok, m := getFields(v, "Username"); ok {
			opt.Username = m["Username"]
		}

		if ok, m := getFields(v, "TlsX509CertFile", "TlsX509KeyFile", "TlsCACertFile"); ok {
			cert, err := tls.LoadX509KeyPair(m["TlsX509CertFile"], m["TlsX509KeyFile"])
			if err != nil {
				return nil, err
			}
			caCert, err := os.ReadFile(m["TlsCACertFile"])
			if err != nil {
				return nil, err
			}
			caCertPool := x509.NewCertPool()
			caCertPool.AppendCertsFromPEM(caCert)
			opt.TLSConfig = &tls.Config{
				MinVersion:   tls.VersionTLS12,
				Certificates: []tls.Certificate{cert},
				RootCAs:      caCertPool,
			}
		}

		if ok, m := getFields(v, "Timeout"); ok {
			t, err := strconv.Atoi(m["Timeout"])
			if err == nil {
				opt.DialTimeout = time.Duration(t) * time.Second
			}
		}
		return redis.NewClient(opt), nil
	}

	return nil, ErrInvalidRedisClientConfig
}

// UniversalRedisClient creates a redis client based on the config.
// todo: implement it
func UniversalRedisClient(ctx context.Context, config interface{}) (redis.UniversalClient, error) {
	switch c := config.(type) {
	case *redis.UniversalOptions:
		return redis.NewUniversalClient(c), nil
	case *redis.FailoverOptions:
		return redis.NewFailoverClient(c), nil
	}

	opts := &redis.UniversalOptions{
		Addrs: make([]string, 0),
	}

	v := reflect.ValueOf(config)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if ok, m := getFields(v, "Addrs"); ok {
		opts.Addrs = strings.Split(m["Addrs"], ",")
	}

	if ok, m := getFields(v, "Addr"); ok {
		opts.Addrs = append(opts.Addrs, m["Addr"])
	}

	return nil, ErrInvalidRedisClientConfig
}
