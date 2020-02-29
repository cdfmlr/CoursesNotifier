package data

import (
	"database/sql"
	"example.com/CoursesNotifier/models"
	"reflect"
	"testing"
)

func TestNewStudentCourseRelationshipDatabase(t *testing.T) {
	type args struct {
		dataSourceName string
	}
	tests := []struct {
		name string
		args args
		want *StudentCourseRelationshipDatabase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStudentCourseRelationshipDatabase(tt.args.dataSourceName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStudentCourseRelationshipDatabase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudentCourseRelationshipDatabase_Delete(t *testing.T) {
	type fields struct {
		dataSourceName string
	}
	type args struct {
		r models.Relationship
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantRowsAffected int64
		wantErr          bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rdb := &StudentCourseRelationshipDatabase{
				dataSourceName: tt.fields.dataSourceName,
			}
			gotRowsAffected, err := rdb.Delete(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRowsAffected != tt.wantRowsAffected {
				t.Errorf("Delete() gotRowsAffected = %v, want %v", gotRowsAffected, tt.wantRowsAffected)
			}
		})
	}
}

func TestStudentCourseRelationshipDatabase_GetRelationship(t *testing.T) {
	type fields struct {
		dataSourceName string
	}
	type args struct {
		sid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Relationship
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rdb := &StudentCourseRelationshipDatabase{
				dataSourceName: tt.fields.dataSourceName,
			}
			got, err := rdb.GetRelationshipsOfStudent(tt.args.sid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRelationshipsOfStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRelationshipsOfStudent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudentCourseRelationshipDatabase_GetRelationships(t *testing.T) {
	type fields struct {
		dataSourceName string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []models.Relationship
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rdb := &StudentCourseRelationshipDatabase{
				dataSourceName: tt.fields.dataSourceName,
			}
			got, err := rdb.GetAllRelationships()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllRelationships() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllRelationships() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudentCourseRelationshipDatabase_Insert(t *testing.T) {
	type fields struct {
		dataSourceName string
	}
	type args struct {
		relationship models.Relationship
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantRowsAffected int64
		wantErr          bool
	}{
		{
			name: "insert Foo",
			fields:fields{dataSourceName:"c:000123@/test?charset=utf8"},
			args:args{relationship:*models.NewRelationship("201810000999", "ee7d85e6d90a1981141e7f50f72d9a63")},
			wantRowsAffected: 1,
			wantErr: false,
		},
		{
			name: "insert Foo Again",
			fields:fields{dataSourceName:"c:000123@/test?charset=utf8"},
			args:args{relationship:*models.NewRelationship("201810000999", "ee7d85e6d90a1981141e7f50f72d9a63")},
			wantRowsAffected: 0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rdb := &StudentCourseRelationshipDatabase{
				dataSourceName: tt.fields.dataSourceName,
			}
			gotRowsAffected, err := rdb.Insert(tt.args.relationship)
			if (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRowsAffected != tt.wantRowsAffected {
				t.Errorf("Insert() gotRowsAffected = %v, want %v", gotRowsAffected, tt.wantRowsAffected)
			}
		})
	}
}

func TestStudentCourseRelationshipDatabase_Update(t *testing.T) {
	type fields struct {
		dataSourceName string
	}
	type args struct {
		sid          string
		relationship models.Relationship
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantRowsAffected int64
		wantErr          bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rdb := &StudentCourseRelationshipDatabase{
				dataSourceName: tt.fields.dataSourceName,
			}
			gotRowsAffected, err := rdb.Update(tt.args.sid, tt.args.relationship)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRowsAffected != tt.wantRowsAffected {
				t.Errorf("Update() gotRowsAffected = %v, want %v", gotRowsAffected, tt.wantRowsAffected)
			}
		})
	}
}

func Test_getRelationship(t *testing.T) {
	type args struct {
		db  *sql.DB
		sid string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Relationship
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getRelationshipsOfStudent(tt.args.db, tt.args.sid)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRelationshipsOfStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRelationshipsOfStudent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRelationships(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Relationship
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getAllRelationships(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAllRelationships() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllRelationships() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertRelationship(t *testing.T) {
	type args struct {
		db           *sql.DB
		relationship models.Relationship
	}
	tests := []struct {
		name             string
		args             args
		wantRowsAffected int64
		wantErr          bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRowsAffected, err := insertRelationship(tt.args.db, tt.args.relationship)
			if (err != nil) != tt.wantErr {
				t.Errorf("insertRelationship() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRowsAffected != tt.wantRowsAffected {
				t.Errorf("insertRelationship() gotRowsAffected = %v, want %v", gotRowsAffected, tt.wantRowsAffected)
			}
		})
	}
}

func Test_updateRelationship(t *testing.T) {
	type args struct {
		db           *sql.DB
		sid          string
		relationship models.Relationship
	}
	tests := []struct {
		name             string
		args             args
		wantRowsAffected int64
		wantErr          bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRowsAffected, err := updateRelationship(tt.args.db, tt.args.sid, tt.args.relationship)
			if (err != nil) != tt.wantErr {
				t.Errorf("updateRelationship() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRowsAffected != tt.wantRowsAffected {
				t.Errorf("updateRelationship() gotRowsAffected = %v, want %v", gotRowsAffected, tt.wantRowsAffected)
			}
		})
	}
}