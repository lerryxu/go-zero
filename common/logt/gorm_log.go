package logt

import (
	"context"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/logger"
)

type GormLogger struct {
	SlowThreshold time.Duration
}

// debug,info,error,severe
func NewGormLogger(slow time.Duration) logger.Interface {
	return &GormLogger{
		SlowThreshold: slow,
	}
}

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	logx.WithContext(ctx).Debugf(msg, data)
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	logx.WithContext(ctx).Infof(msg, data)
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	logx.WithContext(ctx).Errorf(msg, data)
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	fidlds := []logx.LogField{
		logx.Field("rows", rows),
		logx.Field("sql", sql),
	}

	switch {
	case err != nil && (!errors.Is(err, logger.ErrRecordNotFound)):
		logx.WithContext(ctx).WithDuration(elapsed).WithFields(fidlds...).Error(err)
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0:
		logx.WithContext(ctx).WithDuration(elapsed).WithFields(fidlds...).Slowf("SLOW SQL >= %v", l.SlowThreshold)
	default:
		logx.WithContext(ctx).WithDuration(elapsed).WithFields(fidlds...).Debug()
	}
}
