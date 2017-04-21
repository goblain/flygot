# FlyGoT

FlyGot is a tiny wrapper over concourses `fly` binary that takes file for pipeline configuration and passes it via Go langs templating engine before involing fly with the rendered contents.

## Usecases

### Avoiding boilerplate code

	{{ define "blockfoo" }}
	- name: {{.}}-git
          type: git
          source:
             ...
	{{ end }}
	{{ template "blockfoo" "value1" }}
	{{ template "blockfoo" "value2" }}
	{{ template "blockfoo" "value3" }}

