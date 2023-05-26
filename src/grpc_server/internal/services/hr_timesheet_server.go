package services

import (
	"powergate.com/hr_timesheet/api/pb"
	"powergate.com/hr_timesheet/component"
)

type HrTimeSheetServer struct {
	appContext component.AppContext
	pb.UnimplementedHrTimeSheetServiceServer
}

func NewHrTimeSheet(appCtx component.AppContext) *HrTimeSheetServer {
	return &HrTimeSheetServer{
		appContext:                            appCtx,
		UnimplementedHrTimeSheetServiceServer: pb.UnimplementedHrTimeSheetServiceServer{},
	}
}
