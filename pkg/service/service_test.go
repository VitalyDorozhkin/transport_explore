package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/VitalyDorozhkin/transport_explore/pkg/models"
)

func Test_service_GetUser(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *models.GetUserRequest
	}
	tests := []struct {
		name    string
		args    args
		want    models.Response
		wantErr bool
	}{
		{
			name: "good",
			args: args{
				ctx: nil,
				request: &models.GetUserRequest{Id: 1},
			},
			want: models.Response{
				Error: false,
				Data: &models.Data{Res: true},
			},
			wantErr: false,
		},
		{
			name: "bad",
			args: args{
				ctx: nil,
				request:  &models.GetUserRequest{Id: -1},
			},
			want: models.Response{
				Error: true,
				ErrorText: "bad id",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{}
			got, err := s.GetUser(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_GetUserCount(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *models.GetUserCountRequest
	}
	tests := []struct {
		name    string
		args    args
		want    models.Response
		wantErr bool
	}{
		{
			name: "good",
			args: args{
				ctx: nil,
				request: &models.GetUserCountRequest{Id: 1},
			},
			want: models.Response{
				Error: false,
				Data: &models.Data{Res: true},
			},
			wantErr: false,
		},
		{
			name: "bad",
			args: args{
				ctx: nil,
				request:  &models.GetUserCountRequest{Id: -1},
			},
			want: models.Response{
				Error: true,
				ErrorText: "bad id",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{}
			got, err := s.GetUserCount(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserCount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_PostOrder(t *testing.T) {
	type args struct {
		ctx     context.Context
		request *models.PostOrderRequest
	}
	tests := []struct {
		name    string
		args    args
		want    models.Response
		wantErr bool
	}{
		{
			name: "good",
			args: args{
				ctx: nil,
				request: &models.PostOrderRequest{Id: 1},
			},
			want: models.Response{
				Error: false,
				Data: &models.Data{Res: true},
			},
			wantErr: false,
		},
		{
			name: "bad",
			args: args{
				ctx: nil,
				request:  &models.PostOrderRequest{Id: -1},
			},
			want: models.Response{
				Error: true,
				ErrorText: "bad id",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{}
			got, err := s.PostOrder(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostOrder() got = %v, want %v", got, tt.want)
			}
		})
	}
}
