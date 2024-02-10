package app

import (
    "log/slog"
    "time"

    grpcapp "github.com/mmeow0/sso/internal/app/grpc"
    "github.com/mmeow0/sso/internal/services/auth"
    "github.com/mmeow0/sso/internal/storage/sqlite"
)

type App struct {
    GRPCServer *grpcapp.App
}

func New(
    log *slog.Logger,
    grpcPort int,
    storagePath string,
    tokenTTL time.Duration,
) *App {
    storage, err := sqlite.New(storagePath)
    if err != nil {
        panic(err)
    }

    authService := auth.New(log, storage, storage, storage, tokenTTL)

    grpcApp := grpcapp.New(log, authService, grpcPort)

    return &App{
        GRPCServer: grpcApp,
    }
}