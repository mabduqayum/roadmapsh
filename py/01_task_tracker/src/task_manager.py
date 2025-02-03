import json
import os
from datetime import datetime

from .task import Task, TaskStatus


class TaskManager:
    def __init__(self, file_path: str = "data/tasks.json"):
        self.file_path = file_path
        self._ensure_file_exists()

    def _ensure_file_exists(self) -> None:
        """Ensure the tasks file exists and is properly initialized."""
        os.makedirs(os.path.dirname(self.file_path), exist_ok=True)
        if not os.path.exists(self.file_path):
            with open(self.file_path, 'w') as f:
                json.dump({"tasks": []}, f)

    def _load_tasks(self) -> [Task]:
        """Load tasks from the JSON file."""
        with open(self.file_path, 'r') as f:
            data = json.load(f)
            return [Task.from_dict(task) for task in data["tasks"]]

    def _save_tasks(self, tasks: [Task]) -> None:
        """Save tasks to the JSON file."""
        with open(self.file_path, 'w') as f:
            json.dump({"tasks": [task.to_dict() for task in tasks]}, f, indent=2)

    def add_task(self, description: str) -> Task:
        """Add a new task."""
        tasks = self._load_tasks()
        new_id = max((task.id for task in tasks), default=0) + 1
        now = datetime.now()

        new_task = Task(
            id=new_id,
            description=description,
            status=TaskStatus.TODO,
            created_at=now,
            updated_at=now
        )

        tasks.append(new_task)
        self._save_tasks(tasks)
        return new_task

    def update_task(self, task_id: int, description: str) -> Task | None:
        """Update a task's description."""
        tasks = self._load_tasks()
        for task in tasks:
            if task.id == task_id:
                task.description = description
                task.updated_at = datetime.now()
                self._save_tasks(tasks)
                return task
        return None

    def delete_task(self, task_id: int) -> bool:
        """Delete a task."""
        tasks = self._load_tasks()
        initial_length = len(tasks)
        tasks = [task for task in tasks if task.id != task_id]
        if len(tasks) != initial_length:
            self._save_tasks(tasks)
            return True
        return False

    def mark_task_status(self, task_id: int, status: TaskStatus) -> list[Task] | None:
        """Mark a task with a new status."""
        tasks = self._load_tasks()
        for task in tasks:
            if task.id == task_id:
                task.status = status
                task.updated_at = datetime.now()
                self._save_tasks(tasks)
                return task
        return None

    def list_tasks(self, status: list[TaskStatus] | None = None) -> [Task]:
        """List tasks, optionally filtered by status."""
        tasks = self._load_tasks()
        if status:
            return [task for task in tasks if task.status == status]
        return tasks
