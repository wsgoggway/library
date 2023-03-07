package database

// import (
// 	"context"
// 	"database/sql"
// 	"testing"
// 	"time"

// 	cache "github.com/wsgoggway/library/cache"

// 	"github.com/go-redis/redis/v8"
// 	"github.com/stretchr/testify/suite"
// 	"go.uber.org/zap"
// )

// type DatabaseTestSuite struct {
// 	suite.Suite

// 	db IDatabase
// }

// func (suite *DatabaseTestSuite) SetupTest() {
// 	zapLogger, _ := zap.NewDevelopment()
// 	log := zapLogger.Sugar()

// 	redisClient := redis.NewClient(&redis.Options{
// 		Addr:         ":6738",
// 		Password:     "",
// 		DialTimeout:  time.Second * 3,
// 		ReadTimeout:  time.Second * 3,
// 		WriteTimeout: time.Second * 3,
// 	})

// 	cacheIplm := new(cache.Cache)
// 	cacheIplm.SetCacheImplementation(redisClient)
// 	cacheIplm.SetLogger(log)

// 	var err error
// 	suite.db, err = New(cacheIplm, log)
// 	if err != nil {
// 	}
// }

// type ContentFile struct {
// 	ContentType string `json:"contentType"`
// 	URI         string `json:"uri"`
// 	Size        int64  `json:"size"`
// }

// type Content struct {
// 	ID          int64          `db:"id" col:"id"`
// 	Created     time.Time      `db:"created" col:"created"`
// 	Updated     time.Time      `db:"updated" col:"updated"`
// 	Deleted     *time.Time     `db:"deleted" col:"deleted"`
// 	Title       sql.NullString `db:"title" col:"title"`
// 	URI         sql.NullString `db:"uri" col:"uri"`
// 	AuthorID    int64          `db:"author_id" col:"author_id"`
// 	CategoryID  int64          `db:"category_id" col:"category_id"`
// 	Status      int32          `db:"status" col:"status"`
// 	Meta        []byte         `db:"meta" col:"meta"`
// 	Description sql.NullString `db:"description" col:"description"`
// 	Search      sql.NullString `db:"search" col:"search"`
// 	UploadToken sql.NullString `db:"upload_token" col:"upload_token"`
// 	ContentType sql.NullString `db:"content_type" col:"content_type"`
// 	Playlist    sql.NullString `db:"playlist" col:"playlist"`
// 	Files       []*ContentFile `db:"files" col:"files"`
// 	AvStatus    int32          `db:"-"`
// 	AvMessage   sql.NullString `db:"-"`
// 	Test        sql.NullString `db:"-"`
// }

// func (suite *DatabaseTestSuite) TestExample() {
// 	sql := `INSERT INTO public.wbdigital_content
// 	(created, updated, deleted, title, uri, author_id, category_id, description, "search", upload_token, content_type, files)
// 	VALUES( '2022-11-23 21:35:47.000', '2022-11-23 21:37:17.000', null, '''Guardians'' Inferno'' _ Marvel Studios’ Guardians of the Galaxy Vol', '648/9e2949447f68bbde5b068b07b82ed55a', 9, 1, 'qweqweqwe', NULL, 'eyJzMy1kbC53aWxkYmVycmllcy5ydSI6IjJ+RHF2cUs3RXBQUnlSczhWLUtwRHotQU9lRjN6TFRDTCJ9', 'video/mp4', NULL)
// 	RETURNING *;`

// 	d := new(Content)
// 	err := suite.db.InsertWithReplace(context.Background(), d, sql)

// 	suite.T().Log(err)
// 	suite.T().Log(d)
// }

// func TestSuite(t *testing.T) {
// 	suite.Run(t, new(DatabaseTestSuite))
// }
