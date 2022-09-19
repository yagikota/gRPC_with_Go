package http

import (
	"context"

	"github.com/yagikota/gRPC_with_go/pkg/adapter/proto"
	"github.com/yagikota/gRPC_with_go/pkg/usecase"
	"google.golang.org/grpc"
)

type studentServer struct {
	usecase usecase.IStudentUsecase
	// https://github.com/grpc/grpc-go/issues/3794#issuecomment-781863019
	proto.UnimplementedStudentServiceServer
}

func NewStudentServer(gs *grpc.Server, su usecase.IStudentUsecase) *studentServer {
	ss := &studentServer{
		usecase: su,
	}
	return ss
}

func (sh *studentServer) FindAllStudents(ctx context.Context, req *proto.AllStudentsRequest) (*proto.StudentsResponse, error) {
	sSlice, err := sh.usecase.FindAllStudents(ctx)
	if err != nil {
		return nil, err
	}
	students := make([]*proto.StudentResponse, len(sSlice))
	for i, s := range sSlice {
		students[i] = &proto.StudentResponse{
			Id:    int64(s.ID),
			Name:  s.Name,
			Age:   int32(s.Age),
			Class: int32(s.Class),
		}
	}
	return &proto.StudentsResponse{
		Students: students,
	}, nil
}

func (sh *studentServer) FindStudentByID(ctx context.Context, req *proto.StudentByIDRequest) (*proto.StudentResponse, error) {
	student, err := sh.usecase.FindStudentByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &proto.StudentResponse{
		Id:    int64(student.ID),
		Name:  student.Name,
		Age:   int32(student.Age),
		Class: int32(student.Class),
	}, nil
}
