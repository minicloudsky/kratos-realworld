// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: realworld/v1/realworld.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RealworldClient is the client API for Realworld service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RealworldClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserReply, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserReply, error)
	GetCurrentUser(ctx context.Context, in *GetCurrentUserReply, opts ...grpc.CallOption) (*UserReply, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserReply, error)
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	UnFollowUser(ctx context.Context, in *UnFollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error)
	ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error)
	FeedArticles(ctx context.Context, in *FeedArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error)
	GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*SingleCommentReply, error)
	GetComments(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*MultipleCommentsReply, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*SingleCommentReply, error)
	FavoriteArticle(ctx context.Context, in *FavoriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	UnFavoriteArticle(ctx context.Context, in *UnFavoriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error)
	GetTags(ctx context.Context, in *GetTagsRequest, opts ...grpc.CallOption) (*TagListReply, error)
}

type realworldClient struct {
	cc grpc.ClientConnInterface
}

func NewRealworldClient(cc grpc.ClientConnInterface) RealworldClient {
	return &realworldClient{cc}
}

func (c *realworldClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realworldClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realworldClient) GetCurrentUser(ctx context.Context, in *GetCurrentUserReply, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/GetCurrentUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realworldClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realworldClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/GetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realworldClient) FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/FollowUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realworldClient) UnFollowUser(ctx context.Context, in *UnFollowUserRequest, opts ...grpc.CallOption) (*ProfileReply, error) {
	out := new(ProfileReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/UnFollowUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realworldClient) ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error) {
	out := new(MultipleArticlesReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/ListArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realworldClient) FeedArticles(ctx context.Context, in *FeedArticlesRequest, opts ...grpc.CallOption) (*MultipleArticlesReply, error) {
	out := new(MultipleArticlesReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/FeedArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realworldClient) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/GetArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realworldClient) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/CreateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realworldClient) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/UpdateArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realworldClient) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/DeleteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realworldClient) AddComment(ctx context.Context, in *AddCommentRequest, opts ...grpc.CallOption) (*SingleCommentReply, error) {
	out := new(SingleCommentReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/AddComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realworldClient) GetComments(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*MultipleCommentsReply, error) {
	out := new(MultipleCommentsReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/GetComments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realworldClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*SingleCommentReply, error) {
	out := new(SingleCommentReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realworldClient) FavoriteArticle(ctx context.Context, in *FavoriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/FavoriteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realworldClient) UnFavoriteArticle(ctx context.Context, in *UnFavoriteArticleRequest, opts ...grpc.CallOption) (*SingleArticleReply, error) {
	out := new(SingleArticleReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/UnFavoriteArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *realworldClient) GetTags(ctx context.Context, in *GetTagsRequest, opts ...grpc.CallOption) (*TagListReply, error) {
	out := new(TagListReply)
	err := c.cc.Invoke(ctx, "/realworld.v1.Realworld/GetTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RealworldServer is the server API for Realworld service.
// All implementations must embed UnimplementedRealworldServer
// for forward compatibility
type RealworldServer interface {
	Login(context.Context, *LoginRequest) (*UserReply, error)
	Register(context.Context, *RegisterRequest) (*UserReply, error)
	GetCurrentUser(context.Context, *GetCurrentUserReply) (*UserReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UserReply, error)
	GetProfile(context.Context, *GetProfileRequest) (*ProfileReply, error)
	FollowUser(context.Context, *FollowUserRequest) (*ProfileReply, error)
	UnFollowUser(context.Context, *UnFollowUserRequest) (*ProfileReply, error)
	ListArticles(context.Context, *ListArticlesRequest) (*MultipleArticlesReply, error)
	FeedArticles(context.Context, *FeedArticlesRequest) (*MultipleArticlesReply, error)
	GetArticle(context.Context, *GetArticleRequest) (*SingleArticleReply, error)
	CreateArticle(context.Context, *CreateArticleRequest) (*SingleArticleReply, error)
	UpdateArticle(context.Context, *UpdateArticleRequest) (*SingleArticleReply, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*SingleArticleReply, error)
	AddComment(context.Context, *AddCommentRequest) (*SingleCommentReply, error)
	GetComments(context.Context, *GetCommentRequest) (*MultipleCommentsReply, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*SingleCommentReply, error)
	FavoriteArticle(context.Context, *FavoriteArticleRequest) (*SingleArticleReply, error)
	UnFavoriteArticle(context.Context, *UnFavoriteArticleRequest) (*SingleArticleReply, error)
	GetTags(context.Context, *GetTagsRequest) (*TagListReply, error)
	mustEmbedUnimplementedRealworldServer()
}

// UnimplementedRealworldServer must be embedded to have forward compatible implementations.
type UnimplementedRealworldServer struct {
}

func (UnimplementedRealworldServer) Login(context.Context, *LoginRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedRealworldServer) Register(context.Context, *RegisterRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRealworldServer) GetCurrentUser(context.Context, *GetCurrentUserReply) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentUser not implemented")
}
func (UnimplementedRealworldServer) UpdateUser(context.Context, *UpdateUserRequest) (*UserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedRealworldServer) GetProfile(context.Context, *GetProfileRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (UnimplementedRealworldServer) FollowUser(context.Context, *FollowUserRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")
}
func (UnimplementedRealworldServer) UnFollowUser(context.Context, *UnFollowUserRequest) (*ProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFollowUser not implemented")
}
func (UnimplementedRealworldServer) ListArticles(context.Context, *ListArticlesRequest) (*MultipleArticlesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticles not implemented")
}
func (UnimplementedRealworldServer) FeedArticles(context.Context, *FeedArticlesRequest) (*MultipleArticlesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeedArticles not implemented")
}
func (UnimplementedRealworldServer) GetArticle(context.Context, *GetArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedRealworldServer) CreateArticle(context.Context, *CreateArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedRealworldServer) UpdateArticle(context.Context, *UpdateArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArticle not implemented")
}
func (UnimplementedRealworldServer) DeleteArticle(context.Context, *DeleteArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArticle not implemented")
}
func (UnimplementedRealworldServer) AddComment(context.Context, *AddCommentRequest) (*SingleCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComment not implemented")
}
func (UnimplementedRealworldServer) GetComments(context.Context, *GetCommentRequest) (*MultipleCommentsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComments not implemented")
}
func (UnimplementedRealworldServer) DeleteComment(context.Context, *DeleteCommentRequest) (*SingleCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedRealworldServer) FavoriteArticle(context.Context, *FavoriteArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteArticle not implemented")
}
func (UnimplementedRealworldServer) UnFavoriteArticle(context.Context, *UnFavoriteArticleRequest) (*SingleArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFavoriteArticle not implemented")
}
func (UnimplementedRealworldServer) GetTags(context.Context, *GetTagsRequest) (*TagListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTags not implemented")
}
func (UnimplementedRealworldServer) mustEmbedUnimplementedRealworldServer() {}

// UnsafeRealworldServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RealworldServer will
// result in compilation errorhandler.
type UnsafeRealworldServer interface {
	mustEmbedUnimplementedRealworldServer()
}

func RegisterRealworldServer(s grpc.ServiceRegistrar, srv RealworldServer) {
	s.RegisterService(&Realworld_ServiceDesc, srv)
}

func _Realworld_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realworld_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realworld_GetCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentUserReply)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).GetCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/GetCurrentUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).GetCurrentUser(ctx, req.(*GetCurrentUserReply))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realworld_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realworld_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).GetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/GetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).GetProfile(ctx, req.(*GetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realworld_FollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).FollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/FollowUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).FollowUser(ctx, req.(*FollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realworld_UnFollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnFollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).UnFollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/UnFollowUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).UnFollowUser(ctx, req.(*UnFollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realworld_ListArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).ListArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/ListArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).ListArticles(ctx, req.(*ListArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realworld_FeedArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).FeedArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/FeedArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).FeedArticles(ctx, req.(*FeedArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realworld_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/GetArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).GetArticle(ctx, req.(*GetArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realworld_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/CreateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).CreateArticle(ctx, req.(*CreateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realworld_UpdateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).UpdateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/UpdateArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).UpdateArticle(ctx, req.(*UpdateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realworld_DeleteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).DeleteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/DeleteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).DeleteArticle(ctx, req.(*DeleteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realworld_AddComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).AddComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/AddComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).AddComment(ctx, req.(*AddCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realworld_GetComments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).GetComments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/GetComments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).GetComments(ctx, req.(*GetCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realworld_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realworld_FavoriteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).FavoriteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/FavoriteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).FavoriteArticle(ctx, req.(*FavoriteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realworld_UnFavoriteArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnFavoriteArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).UnFavoriteArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/UnFavoriteArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).UnFavoriteArticle(ctx, req.(*UnFavoriteArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Realworld_GetTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RealworldServer).GetTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/realworld.v1.Realworld/GetTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RealworldServer).GetTags(ctx, req.(*GetTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Realworld_ServiceDesc is the grpc.ServiceDesc for Realworld service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Realworld_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "realworld.v1.Realworld",
	HandlerType: (*RealworldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Realworld_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Realworld_Register_Handler,
		},
		{
			MethodName: "GetCurrentUser",
			Handler:    _Realworld_GetCurrentUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _Realworld_UpdateUser_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _Realworld_GetProfile_Handler,
		},
		{
			MethodName: "FollowUser",
			Handler:    _Realworld_FollowUser_Handler,
		},
		{
			MethodName: "UnFollowUser",
			Handler:    _Realworld_UnFollowUser_Handler,
		},
		{
			MethodName: "ListArticles",
			Handler:    _Realworld_ListArticles_Handler,
		},
		{
			MethodName: "FeedArticles",
			Handler:    _Realworld_FeedArticles_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _Realworld_GetArticle_Handler,
		},
		{
			MethodName: "CreateArticle",
			Handler:    _Realworld_CreateArticle_Handler,
		},
		{
			MethodName: "UpdateArticle",
			Handler:    _Realworld_UpdateArticle_Handler,
		},
		{
			MethodName: "DeleteArticle",
			Handler:    _Realworld_DeleteArticle_Handler,
		},
		{
			MethodName: "AddComment",
			Handler:    _Realworld_AddComment_Handler,
		},
		{
			MethodName: "GetComments",
			Handler:    _Realworld_GetComments_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _Realworld_DeleteComment_Handler,
		},
		{
			MethodName: "FavoriteArticle",
			Handler:    _Realworld_FavoriteArticle_Handler,
		},
		{
			MethodName: "UnFavoriteArticle",
			Handler:    _Realworld_UnFavoriteArticle_Handler,
		},
		{
			MethodName: "GetTags",
			Handler:    _Realworld_GetTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "realworld/v1/realworld.proto",
}
