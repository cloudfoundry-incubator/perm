package rpc

import (
	"context"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/perm-go"
	"code.cloudfoundry.org/perm/messages"
	"code.cloudfoundry.org/perm/models"
)

type RoleServiceServer struct {
	logger lager.Logger

	roleService           models.RoleService
	roleAssignmentService models.RoleAssignmentService
}

func NewRoleServiceServer(
	logger lager.Logger,
	roleService models.RoleService,
	roleAssignmentService models.RoleAssignmentService,
) *RoleServiceServer {
	return &RoleServiceServer{
		logger:                logger,
		roleService:           roleService,
		roleAssignmentService: roleAssignmentService,
	}
}

func (s *RoleServiceServer) CreateRole(
	ctx context.Context,
	req *protos.CreateRoleRequest,
) (*protos.CreateRoleResponse, error) {
	name := models.RoleName(req.GetName())
	var permissions []*models.Permission
	for _, p := range req.GetPermissions() {
		permissions = append(permissions, &models.Permission{
			Name:            models.PermissionName(p.GetName()),
			ResourcePattern: models.PermissionResourcePattern(p.GetResourcePattern()),
		})
	}

	logger := s.logger.Session("create-role").WithData(lager.Data{"role.name": name, "permissions": permissions})
	logger.Debug(messages.Starting)

	role, err := s.roleService.CreateRole(ctx, logger, name, permissions...)

	if err != nil {
		return nil, togRPCError(err)
	}

	logger.Debug(messages.Success)
	return &protos.CreateRoleResponse{
		Role: role.ToProto(),
	}, nil
}

func (s *RoleServiceServer) GetRole(
	ctx context.Context,
	req *protos.GetRoleRequest,
) (*protos.GetRoleResponse, error) {
	name := models.RoleName(req.GetName())
	logger := s.logger.Session("get-role").WithData(lager.Data{"role.name": name})
	logger.Debug(messages.Starting)

	role, err := s.roleService.FindRole(ctx, logger, models.RoleQuery{Name: name})
	if err != nil {
		return nil, togRPCError(err)
	}

	logger.Debug(messages.Success)
	return &protos.GetRoleResponse{
		Role: role.ToProto(),
	}, nil
}

func (s *RoleServiceServer) DeleteRole(
	ctx context.Context,
	req *protos.DeleteRoleRequest,
) (*protos.DeleteRoleResponse, error) {
	name := models.RoleName(req.GetName())
	logger := s.logger.Session("delete-role").WithData(lager.Data{
		"role.name": name,
	})
	logger.Debug(messages.Starting)

	err := s.roleService.DeleteRole(ctx, logger, models.RoleQuery{Name: name})
	if err != nil {
		return nil, togRPCError(err)
	}

	logger.Debug(messages.Success)
	return &protos.DeleteRoleResponse{}, nil
}

func (s *RoleServiceServer) AssignRole(
	ctx context.Context,
	req *protos.AssignRoleRequest,
) (*protos.AssignRoleResponse, error) {
	roleName := models.RoleName(req.GetRoleName())
	pActor := req.GetActor()

	domainID := models.ActorDomainID(pActor.GetID())
	issuer := models.ActorIssuer(pActor.GetIssuer())
	logger := s.logger.Session("assign-role").WithData(lager.Data{
		"actor.id":     domainID,
		"actor.issuer": issuer,
		"role.name":    roleName,
	})
	logger.Debug(messages.Starting)

	err := s.roleAssignmentService.AssignRole(ctx, logger, roleName, domainID, issuer)
	if err != nil {
		return nil, togRPCError(err)
	}

	logger.Debug(messages.Success)
	return &protos.AssignRoleResponse{}, nil
}

func (s *RoleServiceServer) UnassignRole(
	ctx context.Context,
	req *protos.UnassignRoleRequest,
) (*protos.UnassignRoleResponse, error) {
	roleName := models.RoleName(req.GetRoleName())
	pActor := req.GetActor()

	domainID := models.ActorDomainID(pActor.GetID())
	issuer := models.ActorIssuer(pActor.GetIssuer())
	actor := models.Actor{
		DomainID: domainID,
		Issuer:   issuer,
	}
	logger := s.logger.Session("unassign-role").WithData(lager.Data{
		"actor.id":     actor.DomainID,
		"actor.issuer": actor.Issuer,
		"role.name":    roleName,
	})
	logger.Debug(messages.Starting)

	err := s.roleAssignmentService.UnassignRole(ctx, logger, roleName, domainID, issuer)
	if err != nil {
		return nil, togRPCError(err)
	}

	logger.Debug(messages.Success)
	return &protos.UnassignRoleResponse{}, nil
}

func (s *RoleServiceServer) HasRole(
	ctx context.Context,
	req *protos.HasRoleRequest,
) (*protos.HasRoleResponse, error) {
	roleName := models.RoleName(req.GetRoleName())
	pActor := req.GetActor()
	domainID := models.ActorDomainID(pActor.GetID())
	issuer := models.ActorIssuer(pActor.GetIssuer())

	logger := s.logger.Session("has-role").WithData(lager.Data{
		"actor.id":     domainID,
		"actor.issuer": issuer,
		"role.name":    roleName,
	})
	logger.Debug(messages.Starting)

	query := models.RoleAssignmentQuery{
		ActorQuery: models.ActorQuery{
			DomainID: domainID,
			Issuer:   issuer,
		},
		RoleQuery: models.RoleQuery{
			Name: roleName,
		},
	}

	found, err := s.roleAssignmentService.HasRole(ctx, logger, query)
	if err != nil {
		if err == models.ErrRoleNotFound || err == models.ErrActorNotFound {
			return &protos.HasRoleResponse{HasRole: false}, nil
		}

		return nil, togRPCError(err)
	}

	logger.Debug(messages.Success)
	return &protos.HasRoleResponse{HasRole: found}, nil
}

func (s *RoleServiceServer) ListActorRoles(
	ctx context.Context,
	req *protos.ListActorRolesRequest,
) (*protos.ListActorRolesResponse, error) {
	pActor := req.GetActor()
	domainID := models.ActorDomainID(pActor.GetID())
	issuer := models.ActorIssuer(pActor.GetIssuer())
	logger := s.logger.Session("list-actor-roles").WithData(lager.Data{
		"actor.id":     domainID,
		"actor.issuer": issuer,
	})
	logger.Debug(messages.Starting)

	actorQuery := models.ActorQuery{DomainID: domainID, Issuer: issuer}
	roles, err := s.roleAssignmentService.ListActorRoles(ctx, logger, actorQuery)
	if err != nil {
		if err == models.ErrActorNotFound {
			roles = []*models.Role{}
		} else {
			return nil, togRPCError(err)
		}
	}

	var pRoles []*protos.Role

	for _, r := range roles {
		pRoles = append(pRoles, r.ToProto())
	}

	logger.Debug(messages.Success)
	return &protos.ListActorRolesResponse{
		Roles: pRoles,
	}, nil
}

func (s *RoleServiceServer) ListRolePermissions(
	ctx context.Context,
	req *protos.ListRolePermissionsRequest,
) (*protos.ListRolePermissionsResponse, error) {
	roleName := models.RoleName(req.GetRoleName())
	logger := s.logger.Session("list-role-permissions").WithData(lager.Data{
		"role.name": roleName,
	})
	logger.Debug(messages.Starting)

	permissions, err := s.roleService.ListRolePermissions(ctx, logger, models.RoleQuery{Name: roleName})
	if err != nil {
		if err == models.ErrRoleNotFound {
			permissions = []*models.Permission{}
		} else {
			return nil, togRPCError(err)
		}
	}

	var pPermissions []*protos.Permission

	for _, p := range permissions {
		pPermissions = append(pPermissions, p.ToProto())
	}

	logger.Debug(messages.Success)
	return &protos.ListRolePermissionsResponse{
		Permissions: pPermissions,
	}, nil
}
