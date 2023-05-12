package main

import (
	"reflect"
	"testing"
)

func Test_extractImagesForLine(t *testing.T) {
	type args struct {
		line        string
		imagePrefix string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "single image",
			args: args{
				imagePrefix: "gcr.io/tekton-releases/github.com/tektoncd/pipeline/cmd",
				line:        "image: gcr.io/tekton-releases/github.com/tektoncd/pipeline/cmd/controller:v0.47.0@sha256:f2d03e5b00345da4bf91044daff32795f6f54edb23f8a36742abd729929c7943",
			},
			want: []string{
				"gcr.io/tekton-releases/github.com/tektoncd/pipeline/cmd/controller:v0.47.0",
			},
		},
		{
			name: "multi image",
			args: args{
				imagePrefix: "gcr.io/tekton-releases/github.com/tektoncd/pipeline/cmd",
				line:        `"-entrypoint-image", "gcr.io/tekton-releases/github.com/tektoncd/pipeline/cmd/entrypoint:v0.47.0@sha256:5282e057d67e18632b4158994d5a4af50799568d67fcc6db4da53004ae5f4dd5", "-nop-image", "gcr.io/tekton-releases/github.com/tektoncd/pipeline/cmd/nop:v0.47.0@sha256:3bd15d5ea0f19f439c02bc629d04b5759ec0f4e01e84f1963f3533b7e96643c4", "-sidecarlogresults-image", "gcr.io/tekton-releases/github.com/tektoncd/pipeline/cmd/sidecarlogresults:v0.47.0@sha256:f443ac86d9a453c35344c662f34636dc767b31ace68250b8980177917ee9951d", "-workingdirinit-image", "gcr.io/tekton-releases/github.com/tektoncd/pipeline/cmd/workingdirinit:v0.47.0@sha256:e7b8fe95065252123112e43d9a445dc7957adb344a30cb01c38f1b4268d90d6e",`,
			},
			want: []string{
				"gcr.io/tekton-releases/github.com/tektoncd/pipeline/cmd/entrypoint:v0.47.0",
				"gcr.io/tekton-releases/github.com/tektoncd/pipeline/cmd/nop:v0.47.0",
				"gcr.io/tekton-releases/github.com/tektoncd/pipeline/cmd/sidecarlogresults:v0.47.0",
				"gcr.io/tekton-releases/github.com/tektoncd/pipeline/cmd/workingdirinit:v0.47.0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractImagesForLine(tt.args.line, tt.args.imagePrefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractImagesForLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
