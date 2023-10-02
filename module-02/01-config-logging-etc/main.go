package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var listenPortFlag = flag.Int("port", 80, "Port to listen")
var listenAddrFlag string

func init() {
	flag.StringVar(&listenAddrFlag, "addr", "127.0.0.1", "Addr to listen")

	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	flag.Parse()

	readEnvSimple()

	useFlagsSimple()

	loggingExamples()

	yamlConfigExample()

	configFromEnv()

	configFromEnvAndFromFile()

	// metrics

	// handle signals

	// fast startup/graceful shutdown

	// toggles
}

func readEnvSimple() {
	listenPortStr := os.Getenv("PORT")
	if listenPortStr == "" {
		log.Fatal("No 'PORT' variable set")
	}
	listenPort, err := strconv.Atoi(listenPortStr)
	if err != nil {
		log.Fatal(err)
	}
	_ = listenPort

	err = os.Setenv("PORT", "80") // anti pattern
	if err != nil {
		log.Fatal(err)
	}
}

func useFlagsSimple() {
	// ***********
	// flags
	// ${BINARY_NAME} --help
	// Usage of ${BINARY_NAME}:
	//  -addr string
	//        Addr to listen (default "127.0.0.1")
	//  -port int
	//        Port to listen (default 80)

	// ${BINARY_NAME} -addr 10.10.10.10
	fmt.Println(listenAddrFlag) // 10.10.10.10
}

func loggingExamples() {
	log.Println("test") // 2023/01/30 10:57:42 test

	// since 1.21
	slog.Info("test", "param", "value") // 2023/01/30 11:03:53 INFO test param=value

	logger := slog.Default()
	logger.Info("test with default logger", "param", "value") // 2023/01/30 11:04:48 INFO test with default logger param=value

	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("test with default logger", "param", "value") // time=2023-01-30T11:06:10.464+04:00 level=INFO msg="test with default logger" param=value

	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("test with default logger", "param", "value") // {"time":"2023-01-30T11:06:10.464751+04:00","level":"INFO","msg":"test with default logger","param":"value"}

	logrus.WithFields(logrus.Fields{
		"param": "value",
	}).Info("Log string from logrus") // {"level":"info","msg":"Log string from logrus","param":"value","time":"2023-01-30T11:09:19+04:00"}

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	// creates copy of struct internally
	contextLogger := logrus.WithFields(logrus.Fields{
		"param":   "value",
		"userID":  12345,
		"traceId": uuid.New().String(),
	})

	contextLogger.Info("Log with fields")
	// {"level":"info","msg":"Log with fields","param":"value","time":"2023-01-30T11:11:35+04:00","traceId":"5ff1e41d-3236-4898-b92d-d98cd520eb2f","userID":12345}
	contextLogger.WithField("hello", "world").Info("Log with more fields")
	// {"hello":"world","level":"info","msg":"Log with more fields","param":"value","time":"2023-01-30T11:11:35+04:00","traceId":"5ff1e41d-3236-4898-b92d-d98cd520eb2f","userID":12345}

}

func yamlConfigExample() {
	viper.SetConfigType("yaml")
	var yamlExample = []byte(`
server:
  addr: "127.0.0.1"
  port: 8080
middlewares:
  - "csrf"
  - "rate-limiter"
apiKey: "secret"
`)

	type yamlConfig struct {
		Server struct {
			Addr string `yaml:"addr"`
			Port int    `yaml:"port"`
		} `yaml:"server"`
		Middlewares []string `yaml:"middlewares"`
		ApiKey      string   `yaml:"secret"`
	}

	err := viper.ReadConfig(bytes.NewBuffer(yamlExample))
	if err != nil {
		log.Fatal(err)
	}
	config := yamlConfig{}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config) // {{127.0.0.1 8080} [csrf rate-limiter] secret}
}

func configFromEnv() {
	v := viper.New()

	type envConfig struct {
		ServerAddr  string   `mapstructure:"SERVER_ADDR"`
		ServerPort  int      `mapstructure:"SERVER_PORT"`
		Middlewares []string `mapstructure:"MIDDLEWARES"`
		ApiKey      string   `mapstructure:"SECRET"`
	}

	// don't confuse
	// viper.SetConfigType("env")

	v.SetConfigType("env")
	v.SetEnvPrefix("MY_SERVICE")

	v.SetDefault("SECRET", "")
	v.SetDefault("MIDDLEWARES", []string{})
	v.SetDefault("SERVER_ADDR", "127.0.0.1")
	v.SetDefault("SERVER_PORT", "9090")

	v.AutomaticEnv()
	cfg := envConfig{}
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	// MY_SERVICE_SERVER_ADDR="10.0.0.1" MY_SERVICE_SERVER_PORT="9090" MY_SERVICE_SECRET=env-secret MY_SERVICE_MIDDLEWARES="rate-limiter ip-blacklist" go run main.go
	fmt.Println(cfg) // {10.0.0.1 9090 [rate-limiter ip-blacklist] env-secret}
}

func configFromEnvAndFromFile() {
	v := viper.New()

	type envAndFileConfig struct {
		ServerAddr  string   `mapstructure:"SERVER_ADDR" yaml:"server_addr"`
		ServerPort  int      `mapstructure:"SERVER_PORT" yaml:"server_port"`
		Middlewares []string `mapstructure:"MIDDLEWARES" yaml:"middlewares"`
		ApiKey      string   `mapstructure:"SECRET" yaml:"secret"`
	}

	v.SetConfigType("yaml")
	v.SetConfigFile("config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	v.AutomaticEnv()
	cfg := envAndFileConfig{}
	err = v.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg) // {10.0.0.1 8080 [csrf ip-blacklist rate-limiter] config-secret}

	// SECRET=other SERVER_ADDR=1.2.3.4 go run main.go
	// {1.2.3.4 8080 [csrf ip-blacklist rate-limiter] other}
}
