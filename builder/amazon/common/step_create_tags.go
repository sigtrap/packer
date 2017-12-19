package common

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	retry "github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/packer"
	"github.com/hashicorp/packer/template/interpolate"
	"github.com/mitchellh/multistep"
)

type StepCreateTags struct {
	AccessConfig *AccessConfig
	Tags         map[string]string
	SnapshotTags map[string]string
	Ctx          interpolate.Context
}

func (s *StepCreateTags) Run(state multistep.StateBag) multistep.StepAction {
	ec2conn := state.Get("ec2").(*ec2.EC2)
	ui := state.Get("ui").(packer.Ui)
	amis := state.Get("amis").(map[string]string)

	var sourceAMI string
	if rawSourceAMI, hasSourceAMI := state.GetOk("source_image"); hasSourceAMI {
		sourceAMI = *rawSourceAMI.(*ec2.Image).ImageId
	} else {
		sourceAMI = ""
	}

	if len(s.Tags) == 0 && len(s.SnapshotTags) == 0 {
		return multistep.ActionContinue
	}

	// Adds tags to AMIs and snapshots
	for ami := range amis {
		ui.Say(fmt.Sprintf("Adding tags to AMI (%s)...", ami))

		// Declare list of resources to tag
		session, err := s.AccessConfig.Session()
		if err != nil {
			err := fmt.Errorf("Error creating AWS session: %s", err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
		regionconn := ec2.New(session)

		// Retrieve image list for given AMI
		resourceIds := []*string{&ami}
		imageResp, err := regionconn.DescribeImages(&ec2.DescribeImagesInput{
			ImageIds: resourceIds,
		})

		if err != nil {
			err := fmt.Errorf("Error retrieving details for AMI (%s): %s", ami, err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}

		if len(imageResp.Images) == 0 {
			err := fmt.Errorf("Error retrieving details for AMI (%s), no images found", ami)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}

		image := imageResp.Images[0]
		snapshotIds := []*string{}

		// Add only those with a Snapshot ID, i.e. not Ephemeral
		for _, device := range image.BlockDeviceMappings {
			if device.Ebs != nil && device.Ebs.SnapshotId != nil {
				ui.Say(fmt.Sprintf("Tagging snapshot: %s", *device.Ebs.SnapshotId))
				resourceIds = append(resourceIds, device.Ebs.SnapshotId)
				snapshotIds = append(snapshotIds, device.Ebs.SnapshotId)
			}
		}

		// Convert tags to ec2.Tag format
		ui.Say("Creating AMI tags")
		amiTags, err := ConvertToEC2Tags(s.Tags, *ec2conn.Config.Region, sourceAMI, s.Ctx)
		if err != nil {
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
		ReportTags(ui, amiTags)

		ui.Say("Creating snapshot tags")
		snapshotTags, err := ConvertToEC2Tags(s.SnapshotTags, *ec2conn.Config.Region, sourceAMI, s.Ctx)
		if err != nil {
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
		ReportTags(ui, snapshotTags)

		// Retry creating tags for about 2.5 minutes
		err = retry.Retry(0.2, 30, 11, func(_ uint) (bool, error) {
			// Tag images and snapshots
			_, err := regionconn.CreateTags(&ec2.CreateTagsInput{
				Resources: resourceIds,
				Tags:      amiTags,
			})
			if awsErr, ok := err.(awserr.Error); ok {
				if awsErr.Code() == "InvalidAMIID.NotFound" ||
					awsErr.Code() == "InvalidSnapshot.NotFound" {
					return false, nil
				}
			}

			// Override tags on snapshots
			if len(snapshotTags) > 0 {
				_, err = regionconn.CreateTags(&ec2.CreateTagsInput{
					Resources: snapshotIds,
					Tags:      snapshotTags,
				})
			}
			if err == nil {
				return true, nil
			}
			if awsErr, ok := err.(awserr.Error); ok {
				if awsErr.Code() == "InvalidSnapshot.NotFound" {
					return false, nil
				}
			}
			return true, err
		})

		if err != nil {
			err := fmt.Errorf("Error adding tags to Resources (%#v): %s", resourceIds, err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
	}

	return multistep.ActionContinue
}

func (s *StepCreateTags) Cleanup(state multistep.StateBag) {
	// No cleanup...
}

func ReportTags(ui packer.Ui, tags []*ec2.Tag) {
	for _, tag := range tags {
		ui.Message(fmt.Sprintf("Adding tag: \"%s\": \"%s\"",
			aws.StringValue(tag.Key), aws.StringValue(tag.Value)))
	}
}

func ConvertToEC2Tags(tags map[string]string, region, sourceAmiId string, ctx interpolate.Context) ([]*ec2.Tag, error) {
	var ec2Tags []*ec2.Tag
	for key, value := range tags {

		ctx.Data = &BuildInfoTemplate{
			SourceAMI:   sourceAmiId,
			BuildRegion: region,
		}
		interpolatedKey, err := interpolate.Render(key, &ctx)
		if err != nil {
			return ec2Tags, fmt.Errorf("Error processing tag: %s:%s - %s", key, value, err)
		}
		interpolatedValue, err := interpolate.Render(value, &ctx)
		if err != nil {
			return ec2Tags, fmt.Errorf("Error processing tag: %s:%s - %s", key, value, err)
		}

		ec2Tags = append(ec2Tags, &ec2.Tag{
			Key:   aws.String(interpolatedKey),
			Value: aws.String(interpolatedValue),
		})
	}

	return ec2Tags, nil
}
