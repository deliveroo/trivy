package aws

import (
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/accessanalyzer"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/apigateway"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/athena"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/cloudfront"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/cloudtrail"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/cloudwatch"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/codebuild"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/config"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/documentdb"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/dynamodb"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/ec2"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/ecr"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/ecs"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/efs"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/eks"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/elasticache"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/elasticsearch"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/elb"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/emr"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/iam"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/kinesis"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/kms"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/lambda"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/mq"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/msk"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/neptune"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/rds"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/redshift"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/s3"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/sam"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/sns"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/sqs"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/ssm"
	"github.com/deliveroo/trivy/pkg/iac/providers/aws/workspaces"
)

type AWS struct {
	Meta           Meta
	AccessAnalyzer accessanalyzer.AccessAnalyzer
	APIGateway     apigateway.APIGateway
	Athena         athena.Athena
	Cloudfront     cloudfront.Cloudfront
	CloudTrail     cloudtrail.CloudTrail
	CloudWatch     cloudwatch.CloudWatch
	CodeBuild      codebuild.CodeBuild
	Config         config.Config
	DocumentDB     documentdb.DocumentDB
	DynamoDB       dynamodb.DynamoDB
	EC2            ec2.EC2
	ECR            ecr.ECR
	ECS            ecs.ECS
	EFS            efs.EFS
	EKS            eks.EKS
	ElastiCache    elasticache.ElastiCache
	Elasticsearch  elasticsearch.Elasticsearch
	ELB            elb.ELB
	EMR            emr.EMR
	IAM            iam.IAM
	Kinesis        kinesis.Kinesis
	KMS            kms.KMS
	Lambda         lambda.Lambda
	MQ             mq.MQ
	MSK            msk.MSK
	Neptune        neptune.Neptune
	RDS            rds.RDS
	Redshift       redshift.Redshift
	SAM            sam.SAM
	S3             s3.S3
	SNS            sns.SNS
	SQS            sqs.SQS
	SSM            ssm.SSM
	WorkSpaces     workspaces.WorkSpaces
}

type Meta struct {
	TFProviders []TerraformProvider
}
