package service

import (
	"context"
	pb "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) FollowUser(context.Context, *pb.FollowUserRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}

func (s *RealWorldService) UnFollowUser(context.Context, *pb.UnFollowUserRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}

func (s *RealWorldService) GetCurrentUser(context.Context, *pb.GetCurrentUserReply) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}

func (s *RealWorldService) UpdateUser(context.Context, *pb.UpdateUserRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}

func (s *RealWorldService) GetProfile(context.Context, *pb.GetProfileRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
