package ec2_instance

type ListAllRegionsJob struct {
	processor *Processor
}

func NewListAllRegionsJob(processor *Processor) *ListAllRegionsJob {
	return &ListAllRegionsJob{
		processor: processor,
	}
}

func (j *ListAllRegionsJob) Id() string {
	return "list_all_regions_for_ec2_instance"
}
func (j *ListAllRegionsJob) Description() string {
	return "Listing all available regions (EC2 Instance)"
}
func (j *ListAllRegionsJob) Run() error {
	regions, err := j.processor.provider.ListAllRegions()
	if err != nil {
		return err
	}
	for _, region := range regions {
		j.processor.jobQueue.Push(NewListEC2InstancesInRegionJob(j.processor, region))
	}
	return nil
}
