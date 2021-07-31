package modzy

type S3KeyDefinition struct {
	Bucket string
	Key    string
}

// S3Inputable is a data input that can be provided when using Jobs().SubmitJobS3(...).
//
// Provided implementations of this function are:
//	S3Input
type S3Inputable func() (*S3KeyDefinition, error)

func S3Input(bucket string, key string) S3Inputable {
	return func() (*S3KeyDefinition, error) {
		return &S3KeyDefinition{
			Bucket: bucket,
			Key:    key,
		}, nil
	}
}
