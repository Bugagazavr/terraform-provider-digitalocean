package digitalocean

import (
	"github.com/digitalocean/terraform-provider-digitalocean/internal/datalist"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDigitalOceanSpacesBuckets() *schema.Resource {
	dataListConfig := &datalist.ResourceConfig{
		RecordSchema:        spacesBucketSchema(),
		ResultAttributeName: "buckets",
		FlattenRecord:       flattenSpacesBucket,
		GetRecords:          getDigitalOceanBuckets,
	}

	return datalist.NewResource(dataListConfig)
}
