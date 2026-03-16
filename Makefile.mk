# project service name
NAME := apisrv

# local database connection settings
TEST_PGDATABASE ?= test-apisrv

PGDATABASE ?= task_planner
PGHOST ?= localhost
PGPORT ?= 5432
PGUSER ?= mikhail
PGPASSWORD ?= postgres

# add -race to GOFLAGS if RACE=1
RACE=0
