package data

import (
	"example.com/CoursesNotifier/models"
	"reflect"
	"testing"
)

func TestNewStudentDatabase(t *testing.T) {
	type args struct {
		dataSourceName string
	}
	tests := []struct {
		name string
		args args
		want *StudentDatabase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStudentDatabase(tt.args.dataSourceName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStudentDatabase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudentDatabase_Delete(t *testing.T) {
	type fields struct {
		dataSourceName string
	}
	type args struct {
		sid string
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantRowsAffected int64
		wantErr          bool
	}{
		{
			name: "delete Foo",
			fields:fields{dataSourceName:"c:000123@/test?charset=utf8"},
			args:args{sid:"201810000999"},
			wantRowsAffected:1,
			wantErr:false,
		},
		{
			name: "delete not exist",
			fields:fields{dataSourceName:"c:000123@/test?charset=utf8"},
			args:args{sid:"不知道"},
			wantRowsAffected: 0,
			wantErr:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &StudentDatabase{
				dataSourceName: tt.fields.dataSourceName,
			}
			gotRowsAffected, err := sdb.Delete(tt.args.sid)
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

func TestStudentDatabase_GetStudent(t *testing.T) {
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
		want    *models.Student
		wantErr bool
	}{
		{
			name: "get Foo",
			fields:fields{dataSourceName:"c:000123@/test?charset=utf8"},
			args:args{sid:"201810000999"},
			want: &models.Student{"201810000999", "updated" ,"hahaha", []models.Course{}, 1582458848},
			wantErr:false,
		},
		{
			name: "get not exist",
			fields:fields{dataSourceName:"c:000123@/test?charset=utf8"},
			args:args{sid:"不知道"},
			want: &models.Student{},
			wantErr:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &StudentDatabase{
				dataSourceName: tt.fields.dataSourceName,
			}
			got, err := sdb.GetStudent(tt.args.sid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStudent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStudent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudentDatabase_GetStudents(t *testing.T) {
	type fields struct {
		dataSourceName string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []models.Student
		wantErr bool
	}{
		{
			name: "getstudents",
			fields:fields{dataSourceName:"c:000123@/test?charset=utf8"},
			want:[]models.Student{{"201810000999", "updated" ,"hahaha", []models.Course{}, 1582458848}},
			wantErr:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &StudentDatabase{
				dataSourceName: tt.fields.dataSourceName,
			}
			got, err := sdb.GetStudents()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStudents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStudents() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStudentDatabase_Insert(t *testing.T) {
	type fields struct {
		dataSourceName string
	}
	type args struct {
		student models.Student
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
			args:args{student: *models.NewStudent("201810000999", "hd000000", "hahaha")},
			wantRowsAffected: 1,
			wantErr: false,
		},
		{		// 重复添加相同 sid
			name: "insert Foo Again",
			fields:fields{dataSourceName:"c:000123@/test?charset=utf8"},
			args:args{student: *models.NewStudent("201810000999", "23333", "sdfsdfdf")},
			wantRowsAffected: 0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &StudentDatabase{
				dataSourceName: tt.fields.dataSourceName,
			}
			gotRowsAffected, err := sdb.Insert(tt.args.student)
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

func TestStudentDatabase_Update(t *testing.T) {
	type fields struct {
		dataSourceName string
	}
	type args struct {
		sid     string
		student models.Student
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantRowsAffected int64
		wantErr          bool
	}{
		{
			name: "update Foo",
			fields:fields{dataSourceName:"c:000123@/test?charset=utf8"},
			args:args{sid: "201810000999", student: *models.NewStudent("201810000999", "updated", "hahaha")},
			wantRowsAffected: 1,
			wantErr:false,
		},
		{		// 更新不存在的
			name: "update not existed",
			fields:fields{dataSourceName:"c:000123@/test?charset=utf8"},
			args:args{sid: "3e234234", student: *models.NewStudent("3e234234", "23333", "sdfsdfdf")},
			wantRowsAffected: 0,
			wantErr:false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &StudentDatabase{
				dataSourceName: tt.fields.dataSourceName,
			}
			gotRowsAffected, err := sdb.Update(tt.args.sid, tt.args.student)
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
