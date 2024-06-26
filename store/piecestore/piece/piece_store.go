package piece

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/zkMeLabs/mechain-storage-provider/pkg/log"
	"github.com/zkMeLabs/mechain-storage-provider/store/piecestore/storage"
)

// NewPieceStore returns an instance of PieceStore
func NewPieceStore(pieceConfig *storage.PieceStoreConfig) (*PieceStore, error) {
	if err := checkConfig(pieceConfig); err != nil {
		log.Errorw("failed to check piece store config", "error", err)
		return nil, err
	}
	blob, err := createStorage(*pieceConfig)
	if err != nil {
		log.Errorw("failed to create storage", "error", err)
		return nil, err
	}
	log.Infow("piece store is running", "storage type", pieceConfig.Store.Storage,
		"shards", pieceConfig.Shards)

	return &PieceStore{blob}, nil
}

// checkConfig checks config if right
func checkConfig(cfg *storage.PieceStoreConfig) error {
	overrideConfigFromEnv(cfg)
	if cfg.Shards > 256 {
		return fmt.Errorf("too many shards: %d", cfg.Shards)
	}
	if cfg.Store.IAMType != storage.AKSKIAMType && cfg.Store.IAMType != storage.SAIAMType {
		return fmt.Errorf("invalid iam type: %s", cfg.Store.IAMType)
	}
	if cfg.Store.MaxRetries < 0 {
		return fmt.Errorf("MaxRetries should be equal or greater than zero")
	}
	if cfg.Store.MinRetryDelay < 0 {
		return fmt.Errorf("MinRetryDelay should be equal or greater than zero")
	}
	if cfg.Store.Storage == storage.DiskFileStore {
		if cfg.Store.BucketURL == "" {
			cfg.Store.BucketURL = setDefaultFileStorePath()
		}
		p, err := filepath.Abs(cfg.Store.BucketURL)
		if err != nil {
			log.Errorw("failed to get absolute path", "bucket", cfg.Store.BucketURL, "error", err)
			return err
		}
		cfg.Store.BucketURL = p
		cfg.Store.BucketURL += "/"
	}
	return nil
}

func overrideConfigFromEnv(cfg *storage.PieceStoreConfig) {
	if val, ok := os.LookupEnv(storage.BucketURL); ok {
		cfg.Store.BucketURL = val
	}
}

func createStorage(cfg storage.PieceStoreConfig) (storage.ObjectStorage, error) {
	var (
		object storage.ObjectStorage
		err    error
	)
	if cfg.Shards > 1 {
		object, err = storage.NewSharded(cfg)
	} else {
		object, err = storage.NewObjectStorage(cfg.Store)
	}
	if err != nil {
		log.Errorw("failed to create storage", "error", err, "object", object)
		return nil, err
	}

	if err = checkBucket(context.Background(), object); err != nil {
		log.Errorw("failed to check bucket due to storage is not configured rightly ", "error", err,
			"object", object)
		return nil, err
	}

	return object, nil
}

// checkBucket checks bucket if exists
func checkBucket(ctx context.Context, store storage.ObjectStorage) error {
	if err := store.HeadBucket(ctx); err != nil {
		log.Errorw("failed to head bucket", "error", err)
		if errors.Is(err, storage.ErrNoSuchBucket) {
			if err2 := store.CreateBucket(ctx); err2 != nil {
				return fmt.Errorf("failed to create bucket in %s: %s, previous err: %s", store, err2, err)
			}
			log.Info("create bucket successfully!")
			return nil
		}
		return storage.ErrNoPermissionAccessBucket
	}
	log.Debugf("succeed to head bucket in %s", store)
	return nil
}

func setDefaultFileStorePath() string {
	defaultBucket := "/var/piecestore"
	switch runtime.GOOS {
	case "linux":
		if os.Getuid() == 0 {
			break
		}
		fallthrough
	case "darwin":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Panicw("failed to get current user's home directory", "error", err)
		}
		defaultBucket = path.Join(homeDir, ".piecestore", "local")
	case "windows":
		defaultBucket = path.Join("C:/piecestore/local")
	default:
		log.Panic("Unknown operating system!")
	}
	return defaultBucket
}
