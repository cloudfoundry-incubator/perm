package rpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"

	"code.cloudfoundry.org/perm/protos"
	"github.com/satori/go.uuid"
)

type RoleServiceServer struct {
	roles        map[string]*protos.Role
	roleBindings map[protos.Actor][]string
}

func NewRoleServiceServer() *RoleServiceServer {
	return &RoleServiceServer{
		roles:        make(map[string]*protos.Role),
		roleBindings: make(map[protos.Actor][]string),
	}
}

func (s *RoleServiceServer) CreateRole(ctx context.Context, req *protos.CreateRoleRequest) (*protos.CreateRoleResponse, error) {
	name := req.GetName()

	if _, exists := s.roles[name]; exists {
		return nil, togRPCError(codes.AlreadyExists, errors.New("role already exists"))
	}

	role := &protos.Role{
		Name: name,
		ID:   uuid.NewV4().String(),
	}
	s.roles[name] = role

	return &protos.CreateRoleResponse{
		Role: role,
	}, nil
}

func (s *RoleServiceServer) GetRole(ctx context.Context, req *protos.GetRoleRequest) (*protos.GetRoleResponse, error) {
	name := req.GetName()

	for _, role := range s.roles {
		if role.GetName() == name {
			return &protos.GetRoleResponse{
				Role: role,
			}, nil
		}
	}

	return nil, togRPCError(codes.NotFound, errors.New("could not find role"))
}

func (s *RoleServiceServer) DeleteRole(ctx context.Context, req *protos.DeleteRoleRequest) (*protos.DeleteRoleResponse, error) {
	name := req.GetName()

	_, ok := s.roles[name]

	if !ok {
		return nil, togRPCError(codes.NotFound, errors.New("could not find role"))
	}

	delete(s.roles, req.GetName())

	// "Cascade"
	// Remove role bindings for role
	for actor, roles := range s.roleBindings {
		for i, role := range roles {
			if role == name {
				s.roleBindings[actor] = append(roles[:i], roles[i+1:]...)
				break
			}
		}
	}

	return &protos.DeleteRoleResponse{}, nil
}

func (s *RoleServiceServer) AssignRole(ctx context.Context, req *protos.AssignRoleRequest) (*protos.AssignRoleResponse, error) {
	roleName := req.GetRoleName()
	actor := req.GetActor()

	if _, exists := s.roles[roleName]; !exists {
		return nil, togRPCError(codes.NotFound, errors.New("could not find role"))
	}

	roleBindings, ok := s.roleBindings[*actor]
	if !ok {
		roleBindings = []string{}
	}

	for _, role := range roleBindings {
		if role == roleName {
			return nil, togRPCError(codes.AlreadyExists, errors.New("actor is already assigned to role"))
		}
	}

	roleBindings = append(roleBindings, roleName)

	s.roleBindings[*actor] = roleBindings

	return &protos.AssignRoleResponse{}, nil
}

func (s *RoleServiceServer) UnassignRole(ctx context.Context, req *protos.UnassignRoleRequest) (*protos.UnassignRoleResponse, error) {
	roleName := req.GetRoleName()
	actor := req.GetActor()

	if _, exists := s.roles[roleName]; !exists {
		return nil, togRPCError(codes.NotFound, errors.New("could not find role"))
	}

	roleBindings, ok := s.roleBindings[*actor]
	if !ok {
		roleBindings = []string{}
	}

	for i, role := range roleBindings {
		if role == roleName {
			s.roleBindings[*actor] = append(roleBindings[:i], roleBindings[i+1:]...)
			return &protos.UnassignRoleResponse{}, nil
		}
	}

	return nil, togRPCError(codes.NotFound, errors.New("actor is not assigned to role"))
}

func (s *RoleServiceServer) HasRole(ctx context.Context, req *protos.HasRoleRequest) (*protos.HasRoleResponse, error) {
	actor := req.GetActor()
	role := req.GetRoleName()
	roleBindings, ok := s.roleBindings[*actor]
	if !ok {
		return &protos.HasRoleResponse{HasRole: false}, nil
	}

	var found bool

	for _, name := range roleBindings {
		if name == role {
			found = true
			break
		}
	}

	return &protos.HasRoleResponse{HasRole: found}, nil
}

func (s *RoleServiceServer) ListActorRoles(ctx context.Context, req *protos.ListActorRolesRequest) (*protos.ListActorRolesResponse, error) {
	actor := req.GetActor()
	roleBindings, ok := s.roleBindings[*actor]
	if !ok {
		return &protos.ListActorRolesResponse{
			Roles: []*protos.Role{},
		}, nil
	}

	var roles []*protos.Role

	for _, id := range roleBindings {
		role, found := s.roles[id]
		if !found {
			return nil, togRPCError(codes.Unknown, errors.New("found a role-binding for non-existent role"))
		}

		roles = append(roles, role)
	}

	return &protos.ListActorRolesResponse{
		Roles: roles,
	}, nil
}