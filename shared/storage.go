package shared

import (
	"context"
	"io"
	"io/ioutil"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
	"github.com/pkg/errors"
)

// Storage wraps gcloud storage bucket
type Storage struct {
	c  *storage.Client
	at *storage.BucketAttrs
}

// StorageOpts required New Storage params
type StorageOpts struct {
	BucketName string
	ProjectID  string
}

// NewStorage creats new storage
func NewStorage(ctx context.Context, opts StorageOpts) Storage {
	client, err := storage.NewClient(ctx)
	if err != nil {
		panic(errors.Wrap(err, "Storage client failed to initialize"))
	}

	attrs, err := client.Bucket(opts.BucketName).Attrs(ctx)
	if err == storage.ErrBucketNotExist {
		if err := client.Bucket(opts.BucketName).Create(ctx, opts.ProjectID, nil); err != nil {
			panic(errors.Wrap(err, "Failed to create New Bucket"))
		}
	}
	if err != nil {
		panic(errors.Wrap(err, "Failed to retrieve Existing Bucket attr"))
	}

	return Storage{c: client, at: attrs}
}

// AddFile add new Object to cloud storage bucket
func (s *Storage) AddFile(name string, file multipart.File) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*2)
	defer cancel()

	w := s.c.Bucket(s.at.Name).Object(name).NewWriter(ctx)
	defer w.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(errors.Wrap(err, "Failed to read uploaded file"))
	}

	if _, err := w.Write(b); err != io.EOF {
		panic(errors.Wrap(w.CloseWithError(err), "Failed to write to Object to bucket"))
	}
}
