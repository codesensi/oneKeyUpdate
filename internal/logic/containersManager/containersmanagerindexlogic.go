package containersManager

import (
	"context"
	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/svc"
	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/types"
	"github.com/onlyLTY/oneKeyUpdate/zspace/internal/utiles"
	"github.com/zeromicro/go-zero/core/logx"
)

type ContainersManagerIndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContainersManagerIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContainersManagerIndexLogic {
	return &ContainersManagerIndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContainersManagerIndexLogic) ContainersManagerIndex() (*[]types.Container, error) {
	list, err := utiles.GetContainerList(l.svcCtx)
	if err != nil {
		return nil, err
	}
	utiles.CheckImageUpdate(l.svcCtx, list)
	return &list, nil
}
