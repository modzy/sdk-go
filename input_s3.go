package modzy

type S3KeyDefinition struct {
	Bucket string
	Key    string
}

type S3Inputable func() (*S3KeyDefinition, error)

func S3Key(bucket string, key string) S3Inputable {
	return func() (*S3KeyDefinition, error) {
		return &S3KeyDefinition{
			Bucket: bucket,
			Key:    key,
		}, nil
	}
}
