// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"bytes"
	"cas/pkg/ent"
	"cas/pkg/graphql/model"
	"context"
	"embed"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Entity() EntityResolver
	Mutation() MutationResolver
	Query() QueryResolver
	Role() RoleResolver
	User() UserResolver
	UserRole() UserRoleResolver
	CreateRoleInput() CreateRoleInputResolver
	CreateUserInput() CreateUserInputResolver
	RoleWhereInput() RoleWhereInputResolver
	UpdateRoleInput() UpdateRoleInputResolver
	UpdateUserInput() UpdateUserInputResolver
	UserRoleWhereInput() UserRoleWhereInputResolver
	UserWhereInput() UserWhereInputResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Entity struct {
		FindUserByID func(childComplexity int, id string) int
	}

	Mutation struct {
		CreateRole func(childComplexity int, input ent.CreateRoleInput) int
		CreateUser func(childComplexity int, input ent.CreateUserInput) int
		DeleteRole func(childComplexity int, id string) int
		DeleteUser func(childComplexity int, id string) int
		UpdateRole func(childComplexity int, id string, input ent.UpdateRoleInput) int
		UpdateUser func(childComplexity int, id string, input ent.UpdateUserInput) int
	}

	PageInfo struct {
		EndCursor       func(childComplexity int) int
		HasNextPage     func(childComplexity int) int
		HasPreviousPage func(childComplexity int) int
		StartCursor     func(childComplexity int) int
	}

	Query struct {
		Login              func(childComplexity int, req model.LoginReq) int
		Node               func(childComplexity int, id string) int
		Nodes              func(childComplexity int, ids []string) int
		Roles              func(childComplexity int) int
		Users              func(childComplexity int) int
		__resolve__service func(childComplexity int) int
		__resolve_entities func(childComplexity int, representations []map[string]interface{}) int
	}

	Role struct {
		CreatedAt func(childComplexity int) int
		CreatedBy func(childComplexity int) int
		DeletedAt func(childComplexity int) int
		ID        func(childComplexity int) int
		Name      func(childComplexity int) int
		UpdatedAt func(childComplexity int) int
		UpdatedBy func(childComplexity int) int
		UserRoles func(childComplexity int) int
		Users     func(childComplexity int) int
	}

	User struct {
		CreatedAt func(childComplexity int) int
		CreatedBy func(childComplexity int) int
		DeletedAt func(childComplexity int) int
		ID        func(childComplexity int) int
		Name      func(childComplexity int) int
		Phone     func(childComplexity int) int
		Roles     func(childComplexity int) int
		UpdatedAt func(childComplexity int) int
		UpdatedBy func(childComplexity int) int
		UserRoles func(childComplexity int) int
	}

	UserRole struct {
		CreatedAt func(childComplexity int) int
		CreatedBy func(childComplexity int) int
		DeletedAt func(childComplexity int) int
		ID        func(childComplexity int) int
		Role      func(childComplexity int) int
		RoleID    func(childComplexity int) int
		UpdatedAt func(childComplexity int) int
		UpdatedBy func(childComplexity int) int
		User      func(childComplexity int) int
		UserID    func(childComplexity int) int
	}

	_Service struct {
		SDL func(childComplexity int) int
	}
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "Entity.findUserByID":
		if e.complexity.Entity.FindUserByID == nil {
			break
		}

		args, err := ec.field_Entity_findUserByID_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Entity.FindUserByID(childComplexity, args["id"].(string)), true

	case "Mutation.createRole":
		if e.complexity.Mutation.CreateRole == nil {
			break
		}

		args, err := ec.field_Mutation_createRole_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateRole(childComplexity, args["input"].(ent.CreateRoleInput)), true

	case "Mutation.createUser":
		if e.complexity.Mutation.CreateUser == nil {
			break
		}

		args, err := ec.field_Mutation_createUser_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateUser(childComplexity, args["input"].(ent.CreateUserInput)), true

	case "Mutation.deleteRole":
		if e.complexity.Mutation.DeleteRole == nil {
			break
		}

		args, err := ec.field_Mutation_deleteRole_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeleteRole(childComplexity, args["id"].(string)), true

	case "Mutation.deleteUser":
		if e.complexity.Mutation.DeleteUser == nil {
			break
		}

		args, err := ec.field_Mutation_deleteUser_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeleteUser(childComplexity, args["id"].(string)), true

	case "Mutation.updateRole":
		if e.complexity.Mutation.UpdateRole == nil {
			break
		}

		args, err := ec.field_Mutation_updateRole_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.UpdateRole(childComplexity, args["id"].(string), args["input"].(ent.UpdateRoleInput)), true

	case "Mutation.updateUser":
		if e.complexity.Mutation.UpdateUser == nil {
			break
		}

		args, err := ec.field_Mutation_updateUser_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.UpdateUser(childComplexity, args["id"].(string), args["input"].(ent.UpdateUserInput)), true

	case "PageInfo.endCursor":
		if e.complexity.PageInfo.EndCursor == nil {
			break
		}

		return e.complexity.PageInfo.EndCursor(childComplexity), true

	case "PageInfo.hasNextPage":
		if e.complexity.PageInfo.HasNextPage == nil {
			break
		}

		return e.complexity.PageInfo.HasNextPage(childComplexity), true

	case "PageInfo.hasPreviousPage":
		if e.complexity.PageInfo.HasPreviousPage == nil {
			break
		}

		return e.complexity.PageInfo.HasPreviousPage(childComplexity), true

	case "PageInfo.startCursor":
		if e.complexity.PageInfo.StartCursor == nil {
			break
		}

		return e.complexity.PageInfo.StartCursor(childComplexity), true

	case "Query.login":
		if e.complexity.Query.Login == nil {
			break
		}

		args, err := ec.field_Query_login_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Login(childComplexity, args["req"].(model.LoginReq)), true

	case "Query.node":
		if e.complexity.Query.Node == nil {
			break
		}

		args, err := ec.field_Query_node_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Node(childComplexity, args["id"].(string)), true

	case "Query.nodes":
		if e.complexity.Query.Nodes == nil {
			break
		}

		args, err := ec.field_Query_nodes_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Nodes(childComplexity, args["ids"].([]string)), true

	case "Query.roles":
		if e.complexity.Query.Roles == nil {
			break
		}

		return e.complexity.Query.Roles(childComplexity), true

	case "Query.users":
		if e.complexity.Query.Users == nil {
			break
		}

		return e.complexity.Query.Users(childComplexity), true

	case "Query._service":
		if e.complexity.Query.__resolve__service == nil {
			break
		}

		return e.complexity.Query.__resolve__service(childComplexity), true

	case "Query._entities":
		if e.complexity.Query.__resolve_entities == nil {
			break
		}

		args, err := ec.field_Query__entities_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.__resolve_entities(childComplexity, args["representations"].([]map[string]interface{})), true

	case "Role.createdAt":
		if e.complexity.Role.CreatedAt == nil {
			break
		}

		return e.complexity.Role.CreatedAt(childComplexity), true

	case "Role.createdBy":
		if e.complexity.Role.CreatedBy == nil {
			break
		}

		return e.complexity.Role.CreatedBy(childComplexity), true

	case "Role.deletedAt":
		if e.complexity.Role.DeletedAt == nil {
			break
		}

		return e.complexity.Role.DeletedAt(childComplexity), true

	case "Role.id":
		if e.complexity.Role.ID == nil {
			break
		}

		return e.complexity.Role.ID(childComplexity), true

	case "Role.name":
		if e.complexity.Role.Name == nil {
			break
		}

		return e.complexity.Role.Name(childComplexity), true

	case "Role.updatedAt":
		if e.complexity.Role.UpdatedAt == nil {
			break
		}

		return e.complexity.Role.UpdatedAt(childComplexity), true

	case "Role.updatedBy":
		if e.complexity.Role.UpdatedBy == nil {
			break
		}

		return e.complexity.Role.UpdatedBy(childComplexity), true

	case "Role.userRoles":
		if e.complexity.Role.UserRoles == nil {
			break
		}

		return e.complexity.Role.UserRoles(childComplexity), true

	case "Role.users":
		if e.complexity.Role.Users == nil {
			break
		}

		return e.complexity.Role.Users(childComplexity), true

	case "User.createdAt":
		if e.complexity.User.CreatedAt == nil {
			break
		}

		return e.complexity.User.CreatedAt(childComplexity), true

	case "User.createdBy":
		if e.complexity.User.CreatedBy == nil {
			break
		}

		return e.complexity.User.CreatedBy(childComplexity), true

	case "User.deletedAt":
		if e.complexity.User.DeletedAt == nil {
			break
		}

		return e.complexity.User.DeletedAt(childComplexity), true

	case "User.id":
		if e.complexity.User.ID == nil {
			break
		}

		return e.complexity.User.ID(childComplexity), true

	case "User.name":
		if e.complexity.User.Name == nil {
			break
		}

		return e.complexity.User.Name(childComplexity), true

	case "User.phone":
		if e.complexity.User.Phone == nil {
			break
		}

		return e.complexity.User.Phone(childComplexity), true

	case "User.roles":
		if e.complexity.User.Roles == nil {
			break
		}

		return e.complexity.User.Roles(childComplexity), true

	case "User.updatedAt":
		if e.complexity.User.UpdatedAt == nil {
			break
		}

		return e.complexity.User.UpdatedAt(childComplexity), true

	case "User.updatedBy":
		if e.complexity.User.UpdatedBy == nil {
			break
		}

		return e.complexity.User.UpdatedBy(childComplexity), true

	case "User.userRoles":
		if e.complexity.User.UserRoles == nil {
			break
		}

		return e.complexity.User.UserRoles(childComplexity), true

	case "UserRole.createdAt":
		if e.complexity.UserRole.CreatedAt == nil {
			break
		}

		return e.complexity.UserRole.CreatedAt(childComplexity), true

	case "UserRole.createdBy":
		if e.complexity.UserRole.CreatedBy == nil {
			break
		}

		return e.complexity.UserRole.CreatedBy(childComplexity), true

	case "UserRole.deletedAt":
		if e.complexity.UserRole.DeletedAt == nil {
			break
		}

		return e.complexity.UserRole.DeletedAt(childComplexity), true

	case "UserRole.id":
		if e.complexity.UserRole.ID == nil {
			break
		}

		return e.complexity.UserRole.ID(childComplexity), true

	case "UserRole.role":
		if e.complexity.UserRole.Role == nil {
			break
		}

		return e.complexity.UserRole.Role(childComplexity), true

	case "UserRole.roleID":
		if e.complexity.UserRole.RoleID == nil {
			break
		}

		return e.complexity.UserRole.RoleID(childComplexity), true

	case "UserRole.updatedAt":
		if e.complexity.UserRole.UpdatedAt == nil {
			break
		}

		return e.complexity.UserRole.UpdatedAt(childComplexity), true

	case "UserRole.updatedBy":
		if e.complexity.UserRole.UpdatedBy == nil {
			break
		}

		return e.complexity.UserRole.UpdatedBy(childComplexity), true

	case "UserRole.user":
		if e.complexity.UserRole.User == nil {
			break
		}

		return e.complexity.UserRole.User(childComplexity), true

	case "UserRole.userID":
		if e.complexity.UserRole.UserID == nil {
			break
		}

		return e.complexity.UserRole.UserID(childComplexity), true

	case "_Service.sdl":
		if e.complexity._Service.SDL == nil {
			break
		}

		return e.complexity._Service.SDL(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputCreateRoleInput,
		ec.unmarshalInputCreateUserInput,
		ec.unmarshalInputRoleOrder,
		ec.unmarshalInputRoleWhereInput,
		ec.unmarshalInputUpdateRoleInput,
		ec.unmarshalInputUpdateUserInput,
		ec.unmarshalInputUserOrder,
		ec.unmarshalInputUserRoleOrder,
		ec.unmarshalInputUserRoleWhereInput,
		ec.unmarshalInputUserWhereInput,
		ec.unmarshalInputloginReq,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Query(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

//go:embed "cas.graphql"
var sourcesFS embed.FS

func sourceData(filename string) string {
	data, err := sourcesFS.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("codegen problem: %s not available", filename))
	}
	return string(data)
}

var sources = []*ast.Source{
	{Name: "cas.graphql", Input: sourceData("cas.graphql"), BuiltIn: false},
	{Name: "federation/directives.graphql", Input: `
	scalar _Any
	scalar _FieldSet

	directive @external on FIELD_DEFINITION
	directive @requires(fields: _FieldSet!) on FIELD_DEFINITION
	directive @provides(fields: _FieldSet!) on FIELD_DEFINITION
	directive @extends on OBJECT | INTERFACE

	directive @key(fields: _FieldSet!, resolvable: Boolean = true) repeatable on OBJECT | INTERFACE
	directive @link(import: [String!], url: String!) repeatable on SCHEMA
	directive @shareable on OBJECT | FIELD_DEFINITION
	directive @tag(name: String!) repeatable on FIELD_DEFINITION | INTERFACE | OBJECT | UNION | ARGUMENT_DEFINITION | SCALAR | ENUM | ENUM_VALUE | INPUT_OBJECT | INPUT_FIELD_DEFINITION
	directive @override(from: String!) on FIELD_DEFINITION
	directive @inaccessible on SCALAR | OBJECT | FIELD_DEFINITION | ARGUMENT_DEFINITION | INTERFACE | UNION | ENUM | ENUM_VALUE | INPUT_OBJECT | INPUT_FIELD_DEFINITION
`, BuiltIn: true},
	{Name: "federation/entity.graphql", Input: `
# a union of all types that use the @key directive
union _Entity = User

# fake type to build resolver interfaces for users to implement
type Entity {
		findUserByID(id: ID!,): User!

}

type _Service {
  sdl: String
}

extend type Query {
  _entities(representations: [_Any!]!): [_Entity]!
  _service: _Service!
}
`, BuiltIn: true},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
