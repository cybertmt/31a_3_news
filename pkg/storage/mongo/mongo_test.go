package mongo

import (
	"GoNews/pkg/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestInstance_AddPost(t *testing.T) {
	type fields struct {
		Client *mongo.Client
		DB     *mongo.Database
	}
	type args struct {
		p storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Instance{
				Client: tt.fields.Client,
				DB:     tt.fields.DB,
			}
			if err := i.AddPost(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("AddPost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInstance_DeletePost(t *testing.T) {
	type fields struct {
		Client *mongo.Client
		DB     *mongo.Database
	}
	type args struct {
		p storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Instance{
				Client: tt.fields.Client,
				DB:     tt.fields.DB,
			}
			if err := i.DeletePost(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInstance_Posts(t *testing.T) {
	type fields struct {
		Client *mongo.Client
		DB     *mongo.Database
	}
	tests := []struct {
		name    string
		fields  fields
		want    []storage.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Instance{
				Client: tt.fields.Client,
				DB:     tt.fields.DB,
			}
			got, err := i.Posts()
			if (err != nil) != tt.wantErr {
				t.Errorf("Posts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Posts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstance_UpdatePost(t *testing.T) {
	type fields struct {
		Client *mongo.Client
		DB     *mongo.Database
	}
	type args struct {
		p storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Instance{
				Client: tt.fields.Client,
				DB:     tt.fields.DB,
			}
			if err := i.UpdatePost(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		constr string
	}
	tests := []struct {
		name    string
		args    args
		want    *Instance
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.constr)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}