package repository

import (
	"context"

	"github.com/yagikota/gRPC_with_go/pkg/domain/model"
)

// IHogeHoge represents interface of HogeHoge
type IStudentRepository interface {
	SelectAllStudents(ctx context.Context) (model.StudentSlice, error)
	SelectStudentByID(ctx context.Context, id int) (*model.Student, error)
}
