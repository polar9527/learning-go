# an ordinary rule
# target grpc/test has no recipe
# target grpc/grpc_demo5_client.go also has no recipe
grpc/test: grpc/grpc_demo5_client.go

# a pattern rule
# Static Pattern Rules versus Implicit Rules

# Rules without Recipes or Prerequisites
# If a rule has no prerequisites or recipe,
# and the target of the rule is a nonexistent file,
# then make imagines this target to have been updated whenever its rule is run.
# This implies that all targets depending on this one will always have their recipe run.

# non-match-anything rule
# a pattern rule without Recipes or Prerequisites
# this is an Implicit Rules below

# An implicit rule can apply to any target that matches its pattern,
# but it does apply only when the target has no recipe otherwise specified,
# and only when the prerequisites can be found.
# so grpc/test and grpc/grpc_demo5_client.go are all in a rule with the prerequisites can be found,
# but grpc/grpc_demo5_server.go is not in a rule with prerequisites can be found,
# If more than one implicit rule appears applicable, only one applies; the choice depends on the order of rules.
grpc/%:
	echo $@ " - " $^

#src
#├── Makefile
#├── grpc
#│   ├── grpc_demo5_client.go
#└── └── grpc_demo5_server.go
