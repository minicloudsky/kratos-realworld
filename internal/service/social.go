package service

import (
	"context"
	pb "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserReply, error) {
	_, _ = s.uc.Login(ctx, req.User.Email, req.User.Password)
	return &pb.UserReply{
		User: &pb.UserReply_User{
			Email:    "1397111131@qq.com",
			Token:    "xxx",
			Username: "minicloudsky",
			Bio:      "xxx",
			Image:    "xxx.png",
		},
	}, nil
}

func (s *RealWorldService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserReply, error) {
	user, err := s.uc.Register(ctx, req.User.Username, req.User.Email, req.User.Password)
	if err != nil {
		panic(err)
	}
	return &pb.UserReply{
		User: &pb.UserReply_User{
			Email:    user.Email,
			Token:    user.Token,
			Username: user.Email,
			Bio:      user.Bio,
			Image:    user.Image,
		},
	}, nil
}

func (s *RealWorldService) GetArticle(context.Context, *pb.GetArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}

func (s *RealWorldService) ListArticles(context.Context, *pb.ListArticlesRequest) (*pb.MultipleArticlesReply, error) {
	return &pb.MultipleArticlesReply{}, nil
}

func (s *RealWorldService) FeedArticles(context.Context, *pb.FeedArticlesRequest) (*pb.MultipleArticlesReply, error) {
	return &pb.MultipleArticlesReply{}, nil
}

func (s *RealWorldService) CreateArticle(context.Context, *pb.CreateArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}

func (s *RealWorldService) UpdateArticle(context.Context, *pb.UpdateArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}

func (s *RealWorldService) DeleteArticle(context.Context, *pb.DeleteArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}

func (s *RealWorldService) AddComment(context.Context, *pb.AddCommentRequest) (*pb.SingleCommentReply, error) {
	return &pb.SingleCommentReply{}, nil
}

func (s *RealWorldService) GetComments(context.Context, *pb.GetCommentRequest) (*pb.MultipleCommentsReply, error) {
	return &pb.MultipleCommentsReply{}, nil
}

func (s *RealWorldService) DeleteComment(context.Context, *pb.DeleteCommentRequest) (*pb.SingleCommentReply, error) {
	return &pb.SingleCommentReply{}, nil
}

func (s *RealWorldService) FavoriteArticle(context.Context, *pb.FavoriteArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}

func (s *RealWorldService) UnFavoriteArticle(context.Context, *pb.UnFavoriteArticleRequest) (*pb.SingleArticleReply, error) {
	return &pb.SingleArticleReply{}, nil
}

func (s *RealWorldService) GetTags(context.Context, *pb.GetTagsRequest) (*pb.TagListReply, error) {
	return &pb.TagListReply{}, nil
}
