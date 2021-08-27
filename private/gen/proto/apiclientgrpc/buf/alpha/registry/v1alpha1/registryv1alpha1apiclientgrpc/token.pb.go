// Copyright 2020-2021 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-apiclientgrpc. DO NOT EDIT.

package registryv1alpha1apiclientgrpc

import (
	context "context"
	v1alpha1 "github.com/bufbuild/buf/private/gen/proto/go/buf/alpha/registry/v1alpha1"
	zap "go.uber.org/zap"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type tokenService struct {
	logger          *zap.Logger
	client          v1alpha1.TokenServiceClient
	contextModifier func(context.Context) context.Context
}

// CreateToken creates a new token suitable for machine-to-machine authentication.
func (s *tokenService) CreateToken(
	ctx context.Context,
	note string,
	expireTime *timestamppb.Timestamp,
) (token string, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.CreateToken(
		ctx,
		&v1alpha1.CreateTokenRequest{
			Note:       note,
			ExpireTime: expireTime,
		},
	)
	if err != nil {
		return "", err
	}
	return response.Token, nil
}

// GetToken gets the specific token for the user
//
// This method requires authentication.
func (s *tokenService) GetToken(ctx context.Context, tokenId string) (token *v1alpha1.Token, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.GetToken(
		ctx,
		&v1alpha1.GetTokenRequest{
			TokenId: tokenId,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Token, nil
}

// ListTokens lists the users active tokens
//
// This method requires authentication.
func (s *tokenService) ListTokens(
	ctx context.Context,
	pageSize uint32,
	pageToken string,
	reverse bool,
) (tokens []*v1alpha1.Token, nextPageToken string, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.ListTokens(
		ctx,
		&v1alpha1.ListTokensRequest{
			PageSize:  pageSize,
			PageToken: pageToken,
			Reverse:   reverse,
		},
	)
	if err != nil {
		return nil, "", err
	}
	return response.Tokens, response.NextPageToken, nil
}

// DeleteToken deletes an existing token.
//
// This method requires authentication.
func (s *tokenService) DeleteToken(ctx context.Context, tokenId string) (_ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	_, err := s.client.DeleteToken(
		ctx,
		&v1alpha1.DeleteTokenRequest{
			TokenId: tokenId,
		},
	)
	if err != nil {
		return err
	}
	return nil
}