package database

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UPDDoc struct {
	Id              int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ErpDoctypeId    int64                  `protobuf:"varint,2,opt,name=erp_doctype_id,json=erpDoctypeId,proto3" json:"erp_doctype_id,omitempty"`
	ErpDocId        int64                  `protobuf:"varint,3,opt,name=erp_doc_id,json=erpDocId,proto3" json:"erp_doc_id,omitempty"`
	SupplierId      int64                  `protobuf:"varint,4,opt,name=supplier_id,json=supplierId,proto3" json:"supplier_id,omitempty"`
	FileDisplayName string                 `protobuf:"bytes,5,opt,name=file_display_name,json=fileDisplayName,proto3" json:"file_display_name,omitempty"`
	IsDelete        bool                   `protobuf:"varint,6,opt,name=is_delete,json=isDelete,proto3" json:"is_delete,omitempty"`
	FilePath        string                 `protobuf:"bytes,7,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
	CreatedAt       *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt       *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

type testStruct struct {
	ID               int64      `json:"id" db:"id" col:"id"`
	CreatedAt        time.Time  `json:"createdAt" col:"created_at"`
	UpdatedAt        time.Time  `json:"updatedAt" col:"updated_at"`
	DeletedAt        *time.Time `json:"deletedAt" col:"deleted_at"`
	FirstName        string     `json:"firstName" col:"first_name"`
	LastName         string     `json:"lastName" col:"last_name"`
	MiddleName       string     `json:"middleName" col:"middle_name"`
	Phone            string     `json:"phone" col:"phone"`
	WBUserID         int64      `json:"wbUserId" col:"wb_user_id"`
	WBIsActive       bool       `json:"wbIsActive" col:"wb_is_active"`
	HasPhoto         bool       `json:"hasPhoto" col:"-"`
	Email            *string    `json:"email" col:"email"`
	Country          *string    `json:"country" col:"country"`
	Gender           string     `json:"gender" db:"-" col:"-"`
	IsActiveEmployee bool       `json:"isActiveEmployee" db:"is_active_employee" col:"is_active_employee"`
	Avatars          []string   `json:"avatar_uri" db:"avatar_uri" col:"avatar_uri"`
}

var testCols = []string{"id", "created_at", "updated_at", "deleted_at", "first_name", "last_name", "middle_name", "phone", "wb_user_id", "wb_is_active", "email", "country", "is_active_employee", "avatar_uri"}

func Test_GetEntityDBFilds(t *testing.T) {
	entitySlice := make([]*testStruct, 0)
	entitySingle := new(testStruct)
	entitySlice2 := make([]*UPDDoc, 0)

	cols := GetEntityDBFilds(entitySlice)
	t.Log(cols)
	assert.Equal(t, testCols, cols)
	cols = GetEntityDBFilds(entitySingle)
	t.Log(cols)
	assert.Equal(t, testCols, cols)
	cols = GetEntityDBFilds(&entitySlice2)
	t.Log(cols)
}

// struct, ptr
