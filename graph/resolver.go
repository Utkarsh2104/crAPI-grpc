package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import pb "grpc/pb/community_grpc/proto"

type Resolver struct {
	pb.CommunityServiceClient
}
