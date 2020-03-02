package qzclient

import (
	"example.com/CoursesNotifier/models"
	"example.com/CoursesNotifier/qz/qzapi"
	"testing"
)

func TestClient_AuthUser(t *testing.T) {
	type fields struct {
		Student       models.Student
		token         string
		CurrentXnxqId string
		CurrentWeek   string
		Courses       map[string]models.Course
	}
	tests := []struct {
		name                 string
		fields               fields
		wantAuthUserRespBody *qzapi.AuthUserRespBody
		wantErr              bool
	}{
		{
			name: "test",
			fields: fields{
				Student:       *models.NewStudent("201810000431", "hd270516", "sf"),
				token:         "",
				CurrentXnxqId: "",
				CurrentWeek:   "",
				Courses:       nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Student:       tt.fields.Student,
				token:         tt.fields.token,
				CurrentXnxqId: tt.fields.CurrentXnxqId,
				CurrentWeek:   tt.fields.CurrentWeek,
				Courses:       tt.fields.Courses,
			}
			gotAuthUserRespBody, err := c.AuthUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(gotAuthUserRespBody)
			//if !reflect.DeepEqual(gotAuthUserRespBody, tt.wantAuthUserRespBody) {
			//	t.Errorf("AuthUser() gotAuthUserRespBody = %v, want %v", gotAuthUserRespBody, tt.wantAuthUserRespBody)
			//}
		})
	}
}

func TestClient_FetchAllTermCourses(t *testing.T) {
	type fields struct {
		Student       models.Student
		token         string
		CurrentXnxqId string
		CurrentWeek   string
		Courses       map[string]models.Course
	}
	type args struct {
		ch chan []models.Course
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				Student:       *models.NewStudent("201810000431", "hd270516", "sf"),
				token:         "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE1ODMxNTQwNTksImF1ZCI6IjIwMTgxMDAwMDQzMSJ9.2W8rqPpRGfB95-pFODhCzIi9Z1iCgygC2x1Palk6EE8",
				CurrentXnxqId: "2019-2020-2",
				CurrentWeek:   "2",
				Courses:       nil,
			},
			args: args{
				ch: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Student:       tt.fields.Student,
				token:         tt.fields.token,
				CurrentXnxqId: tt.fields.CurrentXnxqId,
				CurrentWeek:   tt.fields.CurrentWeek,
				Courses:       tt.fields.Courses,
			}

			ch := make(chan []models.Course)
			go c.FetchAllTermCourses(ch)
			t.Log("<-ch: ", <-ch)
			t.Log("c.Course: ", len(c.Courses), " :", c.Courses)
		})
	}
}

func TestClient_FetchCurrentTime(t *testing.T) {
	type fields struct {
		Student       models.Student
		token         string
		CurrentXnxqId string
		CurrentWeek   string
		Courses       map[string]models.Course
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				Student:       *models.NewStudent("201810000431", "hd270516", "sf"),
				token:         "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE1ODI3ODg3NTYsImF1ZCI6IjIwMTgxMDAwMDQzMSJ9.VJPPqTSTZpBbHVEhjy__uWYwzIfF-sQwv7r0vQJ5ndk",
				CurrentXnxqId: "",
				CurrentWeek:   "",
				Courses:       nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Student:       tt.fields.Student,
				token:         tt.fields.token,
				CurrentXnxqId: tt.fields.CurrentXnxqId,
				CurrentWeek:   tt.fields.CurrentWeek,
				Courses:       tt.fields.Courses,
			}
			if err := c.FetchCurrentTime(); (err != nil) != tt.wantErr {
				t.Errorf("FetchCurrentTime() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Log(c)
		})
	}
}

func TestClient_FetchWeekCoursesSlowly(t *testing.T) {
	type fields struct {
		Student       models.Student
		token         string
		CurrentXnxqId string
		CurrentWeek   string
		Courses       map[string]models.Course
	}
	type args struct {
		week int
		ch   chan []models.Course
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test",
			fields: fields{
				Student:       *models.NewStudent("201810000431", "hd270516", "sf"),
				token:         "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE1ODI4NTg0ODIsImF1ZCI6IjIwMTgxMDAwMDQzMSJ9.Ry1xwipdj23ryJYeM1utwU6E4GsIszONe0iYsWqDX4Y",
				CurrentXnxqId: "2019-2020-2",
				CurrentWeek:   "2",
				Courses:       nil,
			},
			args: args{
				week: 2,
				ch:   nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Student:       tt.fields.Student,
				token:         tt.fields.token,
				CurrentXnxqId: tt.fields.CurrentXnxqId,
				CurrentWeek:   tt.fields.CurrentWeek,
				Courses:       tt.fields.Courses,
			}
			ch := make(chan []models.Course)
			go c.FetchWeekCoursesSlowly(2, ch)
			t.Log(<-ch)
		})
	}
}

func TestClient_appendCourse(t *testing.T) {
	type fields struct {
		Student       models.Student
		token         string
		CurrentXnxqId string
		CurrentWeek   string
		Courses       map[string]models.Course
	}
	type args struct {
		courses []models.Course
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//c := &Client{
			//	Student:       tt.fields.Student,
			//	token:         tt.fields.token,
			//	CurrentXnxqId: tt.fields.CurrentXnxqId,
			//	CurrentWeek:   tt.fields.CurrentWeek,
			//	Courses:       tt.fields.Courses,
			//}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		student models.Student
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "new client",
			args: args{student: *models.NewStudent("20100000000", "1231", "sdfds")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.student)
			t.Log(got)
			//if got := New(tt.args.student); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("New() = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestClient_Save(t *testing.T) {
	type fields struct {
		Student       models.Student
		token         string
		CurrentXnxqId string
		CurrentWeek   string
		Courses       map[string]models.Course
	}
	type args struct {
		databaseSource string
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantRowsAffected int64
	}{
		{
			name: "test",
			fields: fields{
				Student:       *models.NewStudent("201810000431", "hd270516", "sf"),
				token:         "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE1ODI4NzE5MjYsImF1ZCI6IjIwMTgxMDAwMDQzMSJ9.QiQvz1Gu-iq2FohEbWxobjsJKGVqgFCwHZ5Bz53lPZM",
				CurrentXnxqId: "2019-2020-2",
				CurrentWeek:   "2",
				Courses:       nil,
			},
			args:args{databaseSource:"c:000123@/test?charset=utf8"},
			wantRowsAffected: int64(15),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				Student:       tt.fields.Student,
				token:         tt.fields.token,
				CurrentXnxqId: tt.fields.CurrentXnxqId,
				CurrentWeek:   tt.fields.CurrentWeek,
				Courses:       tt.fields.Courses,
			}
			ch := make(chan []models.Course)
			go c.FetchAllTermCourses(ch)
			t.Log("len(<-ch)", len(<-ch))
			if gotRowsAffected := c.Save(tt.args.databaseSource); gotRowsAffected == 0 {
				t.Errorf("Save() = %v, want %v", gotRowsAffected, tt.wantRowsAffected)
			} else {
				t.Log(gotRowsAffected)
			}
		})
	}
}