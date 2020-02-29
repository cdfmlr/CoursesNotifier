package data

import (
	"example.com/CoursesNotifier/models"
	"reflect"
	"testing"
)

func TestCourseDatabase_Delete(t *testing.T) {
	type fields struct {
		dataSourceName string
	}
	type args struct {
		cid string
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
			sdb := &CourseDatabase{
				dataSourceName: tt.fields.dataSourceName,
			}
			gotRowsAffected, err := sdb.Delete(tt.args.cid)
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

func TestCourseDatabase_GetCourse(t *testing.T) {
	type fields struct {
		dataSourceName string
	}
	type args struct {
		cid string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Course
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &CourseDatabase{
				dataSourceName: tt.fields.dataSourceName,
			}
			got, err := sdb.GetCourse(tt.args.cid)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCourse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCourse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCourseDatabase_GetCourses(t *testing.T) {
	type fields struct {
		dataSourceName string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []models.Course
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &CourseDatabase{
				dataSourceName: tt.fields.dataSourceName,
			}
			got, err := sdb.GetCourses()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCourses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCourses() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCourseDatabase_Insert(t *testing.T) {
	type fields struct {
		dataSourceName string
	}
	type args struct {
		course models.Course
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
			args:args{course:*models.NewCourse("豆腐豆腐", "市人大", "法国", "12:00", "14:00", "1-18", "10102")},
			wantRowsAffected: 1,
			wantErr: false,
		},
		{
			name: "insert Foo Again",
			fields:fields{dataSourceName:"c:000123@/test?charset=utf8"},
			args:args{course:*models.NewCourse("豆腐豆腐", "市人大", "法国", "12:00", "14:00", "1-18", "10102")},
			wantRowsAffected: 0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &CourseDatabase{
				dataSourceName: tt.fields.dataSourceName,
			}
			gotRowsAffected, err := sdb.Insert(tt.args.course)
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

func TestCourseDatabase_Update(t *testing.T) {
	type fields struct {
		dataSourceName string
	}
	type args struct {
		cid    string
		course models.Course
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
			sdb := &CourseDatabase{
				dataSourceName: tt.fields.dataSourceName,
			}
			gotRowsAffected, err := sdb.Update(tt.args.cid, tt.args.course)
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

func TestCourseDatabase_GetCourseOnTime(t *testing.T) {
	type fields struct {
		dataSourceName string
	}
	type args struct {
		week  int
		day   int
		begin string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []models.Course
		wantErr bool
	}{
		{
			name:    "TestCourseDatabase_GetCourseOnTime",
			fields:  fields{dataSourceName: "c:000123@/test?charset=utf8"},
			args:    args{
				week:  2,
				day:   4,
				begin: "08:00",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &CourseDatabase{
				dataSourceName: tt.fields.dataSourceName,
			}
			got, err := sdb.GetCoursesOnTime(tt.args.day, tt.args.begin)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCoursesOnTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}

func TestCourseDatabase_GetCoursesBeginTime(t *testing.T) {
	type fields struct {
		dataSourceName string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []string
		wantErr bool
	}{
		{
			name:    "GetCoursesBeginTime",
			fields:  fields{"c:000123@/test?charset=utf8"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdb := &CourseDatabase{
				dataSourceName: tt.fields.dataSourceName,
			}
			got, err := sdb.GetCoursesBeginTime()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCoursesBeginTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
		})
	}
}