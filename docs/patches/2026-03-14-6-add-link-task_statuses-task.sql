ALTER TABLE tasks
ADD CONSTRAINT "Ref_tasks_to_task_statuses"
FOREIGN KEY (statusId)
REFERENCES "task_statuses"("statusId");


