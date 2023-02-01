package session

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/sts"
)

// Session provides a costexplorer specific api session
func Session(
	roleArn string,
	region string,
) (*costexplorer.CostExplorer, error) {
	sess, err := assumedSession(roleArn, region)
	if err != nil {
		return &costexplorer.CostExplorer{}, err
	}
	return costexplorer.New(sess), nil
}

func assumedSession(roleArn string, region string) (*session.Session, error) {
	baseSess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	stsSvc := sts.New(baseSess)
	sessionName := "org-infra-costs"
	assumedRole, err := stsSvc.AssumeRole(&sts.AssumeRoleInput{
		RoleArn:         aws.String(roleArn),
		RoleSessionName: aws.String(sessionName),
	})
	if err != nil {
		return nil, err
	}

	return session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			*assumedRole.Credentials.AccessKeyId,
			*assumedRole.Credentials.SecretAccessKey,
			*assumedRole.Credentials.SessionToken),
		Region: aws.String(region),
	})
}
