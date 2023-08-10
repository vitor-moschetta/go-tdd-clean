package api

import (
	"context"
	"fmt"
	"go-tdd-clean/50/shared"

	"github.com/mercadolibre/fury_go-core/pkg/log"
)

func ErrorHandler(ctx context.Context, err error) {
	if err != nil {
		if customErr, ok := err.(*shared.Error); ok {
			fmt.Println(customErr.Error())
		}
		log.Error(ctx, err.Error())
	}
}
