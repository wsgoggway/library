package search

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type SearchTestSuite struct {
	suite.Suite
	es   SearchEngine
	zlog *zap.SugaredLogger
}

func (s *SearchTestSuite) SetupTest() {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "ts",
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "trace",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		// FunctionKey:    zapcore.OmitKey,
		LineEnding: zapcore.DefaultLineEnding,
	}
	zlogger := zap.New(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, zap.DebugLevel),
		zap.AddCaller(),
		zap.AddStacktrace(zap.WarnLevel),
	)
	log := zlogger.Sugar()
	s.zlog = log

	var err error
	if s.es, err = New(log, 0); err != nil {
		log.Fatal(err)
	}
}

func (s *SearchTestSuite) Test_GetSearchResult() {
	ctx := context.Background()
	result, err := s.es.Search(ctx, "джереми кларксон", 100, 0, "test.offers", []string{"title", "meta.meta.author", "tags"})

	s.Suite.NoError(err)
	s.Suite.NotEmpty(result)

	s.zlog.Info(result)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(SearchTestSuite))
}
