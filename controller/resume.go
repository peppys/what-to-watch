package controller

import (
	"context"

	"github.com/PeppyS/personal-site-api/model"
	pb "github.com/PeppyS/personal-site-api/proto"
	google_proto_empty "github.com/golang/protobuf/ptypes/empty"
	google_proto_timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

type resumeService interface {
	Get() (model.Resume, error)
}

// ResumeController defines controller structure
type ResumeController struct {
	service resumeService
}

// NewResume instantiates ResumeController, handling resume CRUD operations
func NewResume(rs resumeService) *ResumeController {
	return &ResumeController{
		service: rs,
	}
}

// Get retrieves active resume
func (c *ResumeController) Get(ctx context.Context, empty *google_proto_empty.Empty) (*pb.ResumeResponse, error) {
	resume, err := c.service.Get()
	if err != nil {
		return nil, err
	}

	return &pb.ResumeResponse{
		FirstName: resume.FirstName,
		LastName:  resume.LastName,
		Headline:  resume.Headline,
		Summary:   resume.Summary,
		Experience: []*pb.ResumeResponse_Experience{
			// TODO - loop through and set these entries
			&pb.ResumeResponse_Experience{
				Title:       resume.Experience[0].Title,
				Company:     resume.Experience[0].Company,
				Location:    resume.Experience[0].Location,
				Description: resume.Experience[0].Description,
				DateRange: &pb.ResumeResponse_DateRange{
					Start: &google_proto_timestamp.Timestamp{
						Seconds: int64(resume.Experience[0].DateRange.Start.Second()),
					},
					End: &google_proto_timestamp.Timestamp{
						Seconds: int64(resume.Experience[0].DateRange.End.Second()),
					},
					Display: resume.Experience[0].DateRange.Display,
				},
			},
		},
	}, nil
}
