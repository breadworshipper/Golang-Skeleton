package adapter

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

type minioStorage struct {
	adapter *Adapter
}

func MinioStorage() Option {
	return &minioStorage{}
}

func (m *minioStorage) Start(a *Adapter) {
	endpoint := "play.min.io"
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Error while connecting to minio storage")
	}

	a.Storage = minioClient
	m.adapter = a
}

func (m *minioStorage) Close() error {
	return nil
}

// func WithDigihubStorage() Option {
// 	return func(a *Adapter) {
// 		env := config.Envs.Storage

// 		a.Storage = s3.New(s3.Options{
// 			BaseEndpoint: aws.String(env.Endpoint),
// 			Region:       env.Region,
// 			Credentials:  aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(env.Key, env.Secret, "")),
// 		})

// 		_, err := a.Storage.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
// 		if err != nil {
// 			log.Fatal().Err(err).Msgf("Error while connecting to crowners storage: %s", env.Endpoint)
// 		}

// 		log.Info().Msg("Digihub storage connected")
// 	}
// }
