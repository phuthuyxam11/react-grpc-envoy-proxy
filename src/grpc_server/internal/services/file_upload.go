package services

import (
	"fmt"
	"powergate.com/hr_timesheet/api/pb"
)

func (server *HrTimeSheetServer) UploadFile(req *pb.FileUploadRequest, stream pb.HrTimeSheetService_UploadFileServer) error {
	fmt.Println("hello")
	//fileStream, err := stream.Recv()
	//if err != nil {
	//	log.Fatalln("Can't read file info csv: ", err)
	//	return err
	//}
	//fileInfo := fileStream.GetName()
	//count := 0
	//
	//fmt.Println("file info: ", fileInfo)
	//
	////header, err := stream.Recv()
	//if err != nil {
	//	log.Fatalln("Can't read header csv: ", err)
	//	return err
	//}
	//for {
	//	count++
	//	err := contextError(stream.Context())
	//	if err == io.EOF {
	//		// end of one request
	//		break
	//	}
	//	if err != nil {
	//		log.Fatalln("Error of context: ", err)
	//		return err
	//	}
	//}
	return nil

}
