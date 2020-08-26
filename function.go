package p

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/Neu-Robotics/Sanigift-CF/shared"
	"github.com/pkg/errors"
)

const (
	// BucketName SaniGift uploaded image bucket
	BucketName = "sanigift-uploaded-images"
	// ProjectID GCP project_id
	ProjectID = "nextep-279317"
)

type res struct {
	Message string `json:"message"`
}

// UploadImage func dumps file into Storage bucket
func UploadImage(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(res{Message: "Upload Unsuccessful!"})
			fmt.Println(r)
			log.Fatalf("%s", debug.Stack())
		}

		json.NewEncoder(w).Encode(res{Message: "File Uploaded"})
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	f, h, err := r.FormFile("image")
	if err != nil {
		panic(errors.Wrap(err, "Failed to process FormFile"))
	}

	opts := shared.StorageOpts{
		BucketName: BucketName,
		ProjectID:  ProjectID,
	}

	s := shared.NewStorage(ctx, opts)

	s.AddFile(h.Filename, f)
}
