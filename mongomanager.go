package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
	"strings"
	"time"
)

func main() {

	ses := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String("ap-southeast-2")},
		SharedConfigState: session.SharedConfigEnable,
	}))
	cred := credentials.NewStaticCredentials("AKIARWXENY46J4XV2HOS", "9zDEqfJw+F6lGDda5uV0dvk9/MSr+x5mmVKnszhd", "")
	_, err := cred.Get()
	if err != nil {
		fmt.Println(err)
	}
	svc := ec2.New(ses, &aws.Config{Credentials: cred})

	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String(strings.Join([]string{"*", "PdfParser", "*"}, "")),
				},
			},
		},
	}
	resp, err := svc.DescribeInstances(params)
	if err != nil {
		fmt.Println(err)
	}

	dd := "Image_Mongo" + time.Now().Format("2006_01_02_15_04_05")
	fmt.Println(dd)
	for _, item := range resp.Reservations {

		for _, instance := range item.Instances {
			input := &ec2.CreateImageInput{
				Description: aws.String("Image creation test"),
				InstanceId:  instance.InstanceId,
				Name:        aws.String("Image_Mongo" + time.Now().Format("2006_01_02_15_04_05")),
				NoReboot:    aws.Bool(true),
			}
			result, err := svc.CreateImage(input)
			if err != nil {
				log.Fatal("error in creating image")
			}
			fmt.Println(result)
		}
	}

	fmt.Println(resp.Reservations)

}
