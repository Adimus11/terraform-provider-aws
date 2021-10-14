// Code generated by internal/generate/tags/main.go; DO NOT EDIT.
package acmpca

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/acmpca"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

// ListTags lists acmpca service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func ListTags(conn *acmpca.ACMPCA, identifier string) (tftags.KeyValueTags, error) {
	input := &acmpca.ListTagsInput{
		CertificateAuthorityArn: aws.String(identifier),
	}

	output, err := conn.ListTags(input)

	if err != nil {
		return tftags.New(nil), err
	}

	return KeyValueTags(output.Tags), nil
}

// []*SERVICE.Tag handling

// Tags returns acmpca service tags.
func Tags(tags tftags.KeyValueTags) []*acmpca.Tag {
	result := make([]*acmpca.Tag, 0, len(tags))

	for k, v := range tags.Map() {
		tag := &acmpca.Tag{
			Key:   aws.String(k),
			Value: aws.String(v),
		}

		result = append(result, tag)
	}

	return result
}

// KeyValueTags creates tftags.KeyValueTags from acmpca service tags.
func KeyValueTags(tags []*acmpca.Tag) tftags.KeyValueTags {
	m := make(map[string]*string, len(tags))

	for _, tag := range tags {
		m[aws.StringValue(tag.Key)] = tag.Value
	}

	return tftags.New(m)
}

// UpdateTags updates acmpca service tags.
// The identifier is typically the Amazon Resource Name (ARN), although
// it may also be a different identifier depending on the service.
func UpdateTags(conn *acmpca.ACMPCA, identifier string, oldTagsMap interface{}, newTagsMap interface{}) error {
	oldTags := tftags.New(oldTagsMap)
	newTags := tftags.New(newTagsMap)

	if removedTags := oldTags.Removed(newTags); len(removedTags) > 0 {
		input := &acmpca.UntagCertificateAuthorityInput{
			CertificateAuthorityArn: aws.String(identifier),
			Tags:                    Tags(removedTags.IgnoreAWS()),
		}

		_, err := conn.UntagCertificateAuthority(input)

		if err != nil {
			return fmt.Errorf("error untagging resource (%s): %w", identifier, err)
		}
	}

	if updatedTags := oldTags.Updated(newTags); len(updatedTags) > 0 {
		input := &acmpca.TagCertificateAuthorityInput{
			CertificateAuthorityArn: aws.String(identifier),
			Tags:                    Tags(updatedTags.IgnoreAWS()),
		}

		_, err := conn.TagCertificateAuthority(input)

		if err != nil {
			return fmt.Errorf("error tagging resource (%s): %w", identifier, err)
		}
	}

	return nil
}
