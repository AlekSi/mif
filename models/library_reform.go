package models

// generated with gopkg.in/reform.v1

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type libraryTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *libraryTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("library").
func (v *libraryTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *libraryTableType) Columns() []string {
	return []string{"id", "user_id", "volume_id", "created_at", "updated_at"}
}

// NewStruct makes a new struct for that view or table.
func (v *libraryTableType) NewStruct() reform.Struct {
	return new(Library)
}

// NewRecord makes a new record for that table.
func (v *libraryTableType) NewRecord() reform.Record {
	return new(Library)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *libraryTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// LibraryTable represents library view or table in SQL database.
var LibraryTable = &libraryTableType{
	s: parse.StructInfo{Type: "Library", SQLSchema: "", SQLName: "library", Fields: []parse.FieldInfo{{Name: "ID", PKType: "int32", Column: "id"}, {Name: "UserID", PKType: "", Column: "user_id"}, {Name: "VolumeID", PKType: "", Column: "volume_id"}, {Name: "CreatedAt", PKType: "", Column: "created_at"}, {Name: "UpdatedAt", PKType: "", Column: "updated_at"}}, PKFieldIndex: 0},
	z: new(Library).Values(),
}

// String returns a string representation of this struct or record.
func (s Library) String() string {
	res := make([]string, 5)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "UserID: " + reform.Inspect(s.UserID, true)
	res[2] = "VolumeID: " + reform.Inspect(s.VolumeID, true)
	res[3] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[4] = "UpdatedAt: " + reform.Inspect(s.UpdatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Library) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.UserID,
		s.VolumeID,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Library) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.UserID,
		&s.VolumeID,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}

// View returns View object for that struct.
func (s *Library) View() reform.View {
	return LibraryTable
}

// Table returns Table object for that record.
func (s *Library) Table() reform.Table {
	return LibraryTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Library) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Library) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Library) HasPK() bool {
	return s.ID != LibraryTable.z[LibraryTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *Library) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.ID = int32(i64)
	} else {
		s.ID = pk.(int32)
	}
}

// check interfaces
var (
	_ reform.View   = LibraryTable
	_ reform.Struct = new(Library)
	_ reform.Table  = LibraryTable
	_ reform.Record = new(Library)
	_ fmt.Stringer  = new(Library)
)

func init() {
	parse.AssertUpToDate(&LibraryTable.s, new(Library))
}
