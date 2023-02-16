package graph

import pb "grpc/pb/community_grpc/proto"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	pb.CommunityServiceClient
}
