package courseTicker

import (
	"example.com/CoursesNotifier/models"
	"testing"
	"time"
)

func Test_isCourseInWeek(t *testing.T) {
	type args struct {
		course *models.Course
		week   int
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "a",
			args: args{
				course: models.NewCourse("", "", "", "", "", "2", ""),
				week:   2,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "^a",
			args: args{
				course: models.NewCourse("", "", "", "", "", "3", ""),
				week:   2,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "a-b",
			args: args{
				course: models.NewCourse("", "", "", "", "", "1-3", ""),
				week:   2,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "^a-b",
			args: args{
				course: models.NewCourse("", "", "", "", "", "9-12", ""),
				week:   2,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "a,b",
			args: args{
				course: models.NewCourse("", "", "", "", "", "2,9", ""),
				week:   2,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "^a,b",
			args: args{
				course: models.NewCourse("", "", "", "", "", "5,9", ""),
				week:   2,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "a-b,c,d-e,f",
			args: args{
				course: models.NewCourse("", "", "", "", "", "1-3,4,5-9,18", ""),
				week:   2,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "^a-b,c,d-e,f",
			args: args{
				course: models.NewCourse("", "", "", "", "", "1-3,5,7-9,18", ""),
				week:   6,
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isCourseInWeek(tt.args.course, tt.args.week)
			if (err != nil) != tt.wantErr {
				t.Errorf("isCourseInWeek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isCourseInWeek() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNearestBeginTime(t *testing.T) {
	type args struct {
		databaseSource string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "test",
			args: args{"c:000123@/test?charset=utf8"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getNearestBeginTime(tt.args.databaseSource)
			t.Log(got)
		})
	}
}

func Test_getCurrentWeek(t *testing.T) {
	type args struct {
		databaseSource string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "getCurrentWeek",
			args: args{"c:000123@/test?charset=utf8"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCurrentWeek(tt.args.databaseSource); got != tt.want {
				t.Errorf("getCurrentWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_durationToWeek(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2020-02-29",
			args: args{"2020-02-29"},
			want: 0,
		},
		{
			name: "2020-03-01",
			args: args{"2020-03-01"},
			want: 0,
		},
		{
			name: "2020-02-28",
			args: args{"2020-02-28"},
			want: 0,
		},
		{
			name: "2020-02-23",
			args: args{"2020-02-23"},
			want: 0,
		},
		{
			name: "2020-02-22",
			args: args{"2020-02-22"},
			want: 1,
		},
		{
			name: "2020-02-16",
			args: args{"2020-02-16"},
			want: 1,
		},
		{
			name: "2020-01-29",
			args: args{"2020-01-29"},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, _ := time.Parse("2006-01-02", tt.args.s)
			if got := _durationToWeek(time.Since(d)); got != tt.want {
				t.Errorf("durationToWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotifyApproachingCourses(t *testing.T) {
	type args struct {
		databaseSource     string
		minuteBeforeCourse float64
		notifiers          []Notifier
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "a",
			args: args{
				databaseSource:     "c:000123@/test?charset=utf8",
				minuteBeforeCourse: 10,
				notifiers:          []Notifier{LogNotifier("hahaha")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NotifyApproachingCourses(tt.args.databaseSource, tt.args.minuteBeforeCourse, tt.args.notifiers...)
		})
	}
}
