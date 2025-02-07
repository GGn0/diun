package grpc

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/crazy-max/diun/v4/internal/model"
	"github.com/crazy-max/diun/v4/pb"
	"github.com/crazy-max/diun/v4/pkg/registry"
)

func (c *Client) NotifTest(ctx context.Context, request *pb.NotifTestRequest) (*pb.NotifTestResponse, error) {
	createdAt, _ := time.Parse("2006-01-02T15:04:05Z", "2020-03-26T12:23:56Z")
	image, _ := registry.ParseImage(registry.ParseImageOptions{
		Name: "diun/testnotif:latest",
	})
	image.HubLink = ""

	entry := model.NotifEntry{
		Status:   "new",
		Provider: "file",
		Image:    image,
		Manifest: registry.Manifest{
			Name:          "diun/testnotif",
			Tag:           "latest",
			MIMEType:      "application/vnd.docker.distribution.manifest.list.v2+json",
			Digest:        "sha256:216e3ae7de4ca8b553eb11ef7abda00651e79e537e85c46108284e5e91673e01",
			Created:       &createdAt,
			DockerVersion: "",
			Labels: map[string]string{
				"maintainer":                      "CrazyMax",
				"org.label-schema.build-date":     "2020-03-26T12:23:56Z",
				"org.label-schema.description":    "Docker image update notifier",
				"org.label-schema.name":           "Diun",
				"org.label-schema.schema-version": "1.0",
				"org.label-schema.url":            "https://github.com/crazy-max/diun",
				"org.label-schema.vcs-ref":        "e13f097c",
				"org.label-schema.vcs-url":        "https://github.com/crazy-max/diun",
				"org.label-schema.vendor":         "CrazyMax",
				"org.label-schema.version":        "x.x.x",
			},
			Layers: []string{
				"sha256:aad63a9339440e7c3e1fff2b988991b9bfb81280042fa7f39a5e327023056819",
				"sha256:166c6f165b73185ede72415d780538a55c0c8e854bd177925bc007193e5b0d1b",
				"sha256:e05682efa9cc9d6239b2b9252fe0dc1e58d6e1585679733bb94a6549d49e9b10",
				"sha256:c6a5bfed445b3ed7e85523cd73c6532ac9f9b72bb588ca728fd5b33987ca6538",
				"sha256:df2140efb8abeb727ef0b27ff158b7010a7941eb1cfdade505f510a6e1eaf016",
			},
			Platform: "linux/amd64",
		},
	}

	if len(c.notif.List()) == 0 {
		return &pb.NotifTestResponse{
			Message: "No notifier available",
		}, nil
	}

	var sent []string
	for _, n := range c.notif.List() {
		if err := n.Send(entry); err != nil {
			return nil, err
		}
		sent = append(sent, n.Name())
	}

	return &pb.NotifTestResponse{
		Message: fmt.Sprintf("Notification sent for %s notifier(s)", strings.Join(sent, ", ")),
	}, nil
}
