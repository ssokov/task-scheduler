CREATE TABLE "task_statuses" (
	"statusId" SERIAL NOT NULL,
	"title" varchar(255) NOT NULL,
	"alias" varchar(64) NOT NULL,
	CONSTRAINT "task_statuses_pkey" PRIMARY KEY("statusId"),
	CONSTRAINT "task_statuses_alias_key" UNIQUE("alias")
);