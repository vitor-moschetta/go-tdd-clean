package controllers

import (
	"context"
	"fmt"

	"github.com/mercadolibre/fury_go-core/pkg/log"
)

func ErrorHandler(ctx context.Context, err error) {
	if err != nil {
		if customErr, ok := err.(*CustomError); ok {
			fmt.Println(customErr.Error())
		}
		log.Error(ctx, err.Error())
	}
}
