package config

import (
	"encoding/json"
	"fiber-sample-project/app/responses"
	"fiber-sample-project/pb/commonpb"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"google.golang.org/protobuf/proto"
	"gopkg.in/natefinch/lumberjack.v2"
)

func AccessLogger() logger.Config {
	logFile := &lumberjack.Logger{
		Filename:   "logs/access.log", // base name
		MaxSize:    100,               // MB (per file)
		MaxBackups: 30,                // keep 30 days
		MaxAge:     30,                // days
		Compress:   true,
	}

	return logger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
		Stream:     logFile,
	}
}

var ErrorLog = log.New(&lumberjack.Logger{
	Filename:   "logs/error.log",
	MaxSize:    50,
	MaxBackups: 30,
	MaxAge:     30,
	Compress:   true,
}, "", log.LstdFlags)

func ErrorHandler(c fiber.Ctx, err error) error {
	ErrorLog.Printf(
		"ERROR %d %s %s | %v",
		c.Response().StatusCode(),
		c.Method(),
		c.Path(),
		err,
	)

	b, err := json.Marshal(responses.Response{
		Code:    fiber.StatusInternalServerError,
		Message: err.Error(),
		Datas:   map[string]interface{}{},
	})

	resp := &commonpb.Response{
		Code:    int32(fiber.StatusInternalServerError), // Convert int → int32 for protobuf
		Message: "Internal Server Error",
		Datas:   b,
	}

	bytes, errM := proto.Marshal(resp)
	if errM != nil {
		return errM
	}

	c.Set("Content-Type", "application/x-protobuf")
	return c.Status(fiber.StatusInternalServerError).Send(bytes)

}
