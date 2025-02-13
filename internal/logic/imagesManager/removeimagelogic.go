package imagesManager

import (
	"context"
	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/utiles"

	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/svc"
	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveImageLogic {
	return &RemoveImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveImageLogic) RemoveImage(req *types.RemoveImageReq) (resp *types.MsgResp, err error) {
	// todo: add your logic here and delete this line
	msg, err := utiles.RemoveImage(l.svcCtx, req.ImageID, req.Force)
	return &msg, err
}
