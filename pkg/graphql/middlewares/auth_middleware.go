package middlewares

import (
	"cas/tools"
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	"github.com/vektah/gqlparser/v2/ast"
	"net/http"
)

const ResponseWriter = "RESPONSE_WRITER"

// InjectableResponseWriter 将 writer 载入到 ctx 中
type InjectableResponseWriter struct {
	http.ResponseWriter
	Cookie *http.Cookie
}

func (i *InjectableResponseWriter) Write(data []byte) (int, error) {
	if i.Cookie != nil {
		http.SetCookie(i.ResponseWriter, i.Cookie)
	}
	return i.ResponseWriter.Write(data)
}

/*
WriterMiddleware 配置 Cookie 输出流到 ctx
*/
func WriterMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		injectableResponseWriter := InjectableResponseWriter{
			ResponseWriter: c.Writer,
			Cookie:         nil,
		}
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, ResponseWriter, &injectableResponseWriter)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// NewAuthenticationMiddleware 自定义 Graphql HandlerExtension
func NewAuthenticationMiddleware(skipAuthForDirectives ...string) DirectiveDrivenAuthenticator {
	return DirectiveDrivenAuthenticator{SkipAuthFor: skipAuthForDirectives}
}

type DirectiveDrivenAuthenticator struct {
	SkipAuthFor []string
}

func (DirectiveDrivenAuthenticator) ExtensionName() string {
	return "DirectiveDrivenAuthenticator"
}

func (DirectiveDrivenAuthenticator) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

/*
InterceptField 检查是否需要校验 Cookie 以及校验后的值设置
*/
func (d DirectiveDrivenAuthenticator) InterceptField(ctx context.Context, next graphql.Resolver) (interface{}, error) {
	rc := graphql.GetOperationContext(ctx)
	fc := graphql.GetFieldContext(ctx)
	// only apply auth checks to these
	if rc.OperationName != "IntrospectionQuery" && tools.IsOneOf(fc.Object, "Query", "Mutation", "Subscription") {
		// skip auth check
		if containsAnyOfDirectives(fc.Field.Definition.Directives, d.SkipAuthFor...) {
			return next(ctx)
		}
		// get cookie from context
		rawCookie := ctx.Value(tools.CookieName)
		if rawCookie == nil {
			return nil, http.ErrNoCookie
		}
		// assert type as *http.Cookie
		cookie, ok := rawCookie.(*http.Cookie)
		if !ok {
			return nil, errors.New("not valid cookie")
		}
		// parse and validate JWT token using cookie
		token, err := tools.ParseToken(cookie.Value)
		if err != nil {
			return nil, err
		}
		// 保存当前用户 ID
		ctx = context.WithValue(ctx, "UserID", token.UserID)
		// 将 token 保存到上下文中，便于发送 GRPC 请求，由于 GRPC 请求 METADATA key 全小写，所以 Authorization 换成 token
		ctx = context.WithValue(ctx, "token", cookie)
	}
	return next(ctx)
}

func containsAnyOfDirectives(directives ast.DirectiveList, skipAuths ...string) bool {
	for _, directive := range directives {
		for _, skipAuth := range skipAuths {
			if directive.Name == skipAuth {
				return true
			}
		}
	}
	return false
}
