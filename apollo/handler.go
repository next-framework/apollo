package apollo

import (
	"net/http"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

const (
	GET     string = "get"
	HEAD    string = "head"
	POST    string = "post"
	PUT     string = "put"
	PATCH   string = "patch"
	DELETE  string = "delete"
	OPTIONS string = "options"
	TRACE   string = "trace"
)

type Handler interface {
	Handle(c *Context)
}

type HandlerRouterMapping interface {
	Add(router *Router, handler Handler)
	Resolve(r *http.Request) Handler
}

type DefaultHandlerRouterMapping struct {
	Handlers map[string]Handler
	Routers  map[string]*Router
	pMap     map[string]void
	patterns []string
}

type void struct{}

var member void

func (hrm *DefaultHandlerRouterMapping) Add(router *Router, handler Handler) {
	methods := router.Methods
	if len(methods) == 0 {
		methods = getAllMethods()
	}

	if hrm.Routers == nil {
		hrm.Routers = make(map[string]*Router)
	}
	if hrm.pMap == nil {
		hrm.pMap = make(map[string]void)
	}
	for _, v := range methods {
		path := filepath.Clean(router.Path)

		// todo 检查配置的路径格式是否符合要求

		key := buildRouterKey(path, strings.ToLower(v))
		hrm.Routers[key] = router

		_, existed := hrm.pMap[path]
		if !existed {
			hrm.pMap[path] = member
			hrm.patterns = append(hrm.patterns, path)
			sort.Sort(sort.Reverse(sort.StringSlice(hrm.patterns)))
		}
	}

	if hrm.Handlers == nil {
		hrm.Handlers = make(map[string]Handler)
	}
	hrm.Handlers[router.Handler] = handler
}

func (hrm *DefaultHandlerRouterMapping) Resolve(r *http.Request) Handler {
	if len(hrm.patterns) == 0 {
		return nil
	}

	path := r.URL.Path
	method := strings.ToLower(r.Method)
	for _, pattern := range hrm.patterns {
		if match(pattern, path) {
			key := buildRouterKey(pattern, method)
			router, existed := hrm.Routers[key]
			if !existed {
				return nil
			}
			handler, existed := hrm.Handlers[router.Handler]
			if !existed {
				return nil
			}
			return handler
		}
	}

	return nil
}

func getAllMethods() []string {
	methods := []string{GET, HEAD, POST, PUT, PATCH, DELETE, OPTIONS, TRACE}
	return methods
}

func buildRouterKey(path, method string) string {
	key := "[" + path + "]" + ":" + "[" + method + "]"
	return key
}

func match(pattern, path string) bool {
	if pattern == "/*" {
		return true
	}

	patterns := strings.Split(pattern, "/")
	paths := strings.Split(path, "/")

	if len(patterns) > len(paths) {
		return false
	}

	patternsIndex, pathsIndex := 0, 0
	for {
		p1 := patterns[patternsIndex]
		p2 := paths[pathsIndex]

		if p1 == "**" {
			if patternsIndex+1 == len(patterns) || pathsIndex+1 == len(paths) {
				return true
			}
			if touch(patterns[patternsIndex+1], paths[pathsIndex+1]) {
				patternsIndex++
			}
			pathsIndex++
			continue
		}

		if touch(p1, p2) {
			if patternsIndex+1 == len(patterns) && pathsIndex+1 != len(paths) {
				return false
			}
			if patternsIndex+1 == len(patterns) && pathsIndex+1 == len(paths) {
				return true
			}
			patternsIndex++
			pathsIndex++
			continue
		} else {
			return false
		}
	}
}

func touch(pattern, name string) bool {
	if pattern == "*" || pattern == "**" || (strings.HasPrefix(pattern, "{") && strings.HasSuffix(pattern, "}")) {
		return true
	}

	if strings.HasPrefix(pattern, "*") {
		pattern = "." + pattern
	}
	matched, _ := regexp.MatchString(pattern, name)
	return matched
}
