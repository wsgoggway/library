package cache

// func Test_lpush_lpop(t *testing.T) {

// 	encoderCfg := zapcore.EncoderConfig{
// 		TimeKey:        "ts",
// 		MessageKey:     "msg",
// 		LevelKey:       "level",
// 		NameKey:        "logger",
// 		CallerKey:      "caller",
// 		StacktraceKey:  "trace",
// 		EncodeLevel:    zapcore.LowercaseLevelEncoder,
// 		EncodeTime:     zapcore.RFC3339TimeEncoder,
// 		EncodeDuration: zapcore.StringDurationEncoder,
// 		EncodeCaller:   zapcore.ShortCallerEncoder,
// 		FunctionKey:    zapcore.OmitKey,
// 		LineEnding:     zapcore.DefaultLineEnding,
// 	}
// 	zlogger := zap.New(
// 		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, zap.DebugLevel),
// 		zap.AddCaller(),
// 		zap.AddStacktrace(zap.WarnLevel),
// 	)
// 	log := zlogger.Sugar()

// 	redisClient := redis.NewClient(
// 		&redis.Options{
// 			Addr:         "",
// 			Password:     "",
// 			DialTimeout:  time.Second * 3,
// 			ReadTimeout:  time.Second * 3,
// 			WriteTimeout: time.Second * 3,
// 		},
// 	)
// 	defer redisClient.Close()

// 	cacheIplm := new(Cache)
// 	cacheIplm.SetCacheImplementation(redisClient)
// 	cacheIplm.SetLogger(log)

// 	var (
// 		err  error
// 		test AccountingIncreaseRequest
// 	)

// 	req := &IncreaseRequest{
// 		AuthorID:   9,
// 		Amount:     decimal.NewFromInt(10),
// 		TxUUID:     uuid.NewString(),
// 		PurchaseID: 1234,
// 		GUID:       uuid.New(),
// 	}
// 	data, err := req.MarshalJSON()
// 	if err != nil {
// 		log.Error(err)
// 	}

// 	err = cacheIplm.Push("test_q", data)
// 	assert.NoError(t, err)

// 	err = cacheIplm.Pop("test_q", &test)
// 	assert.NoError(t, err)

// 	log.Infof("%#v", test)

// 	assert.Equal(t, int64(1234), test.PurchaseID)
// }
