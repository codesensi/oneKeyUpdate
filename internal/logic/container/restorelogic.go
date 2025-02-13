package container

import (
	"context"
	"github.com/google/uuid"
	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/utiles"
	"runtime/debug"

	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/svc"
	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RestoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRestoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RestoreLogic {
	return &RestoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RestoreLogic) Restore(req *types.ContainerRestoreReq) (resp *types.Resp, err error) {
	resp = &types.Resp{}
	taskID := uuid.New().String()
	go func() {
		// Catch any panic and log the error
		defer func() {
			if r := recover(); r != nil {
				logx.Errorf("Recovered from panic in restoreContainer: %v\n%s", r, debug.Stack())
			}
		}()
		err := utiles.RestoreContainer(l.svcCtx, req.Filename, taskID)
		if err != nil {
			logx.Errorf("Error in restoreContainer: %v", err)
		}
	}()
	resp.Code = 200
	resp.Msg = "success"
	resp.Data = taskID
	return resp, nil
}
