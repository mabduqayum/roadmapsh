import argparse

from .task_manager import TaskManager
from .task import TaskStatus

class CLI:
    def __init__(self):
        self.task_manager = TaskManager()

    @staticmethod
    def setup_parser() -> argparse.ArgumentParser:
        parser = argparse.ArgumentParser(description="Task Tracker CLI")
        subparsers = parser.add_subparsers(dest="command", help="Available commands")

        # Add task
        add_parser = subparsers.add_parser("add", help="Add a new task")
        add_parser.add_argument("description", help="Task description")

        # Update task
        update_parser = subparsers.add_parser("update", help="Update a task")
        update_parser.add_argument("id", type=int, help="Task ID")
        update_parser.add_argument("description", help="New task description")

        # Delete task
        delete_parser = subparsers.add_parser("delete", help="Delete a task")
        delete_parser.add_argument("id", type=int, help="Task ID")

        # Mark task status
        mark_in_progress = subparsers.add_parser("mark-in-progress", help="Mark task as in progress")
        mark_in_progress.add_argument("id", type=int, help="Task ID")

        mark_done = subparsers.add_parser("mark-done", help="Mark task as done")
        mark_done.add_argument("id", type=int, help="Task ID")

        # List tasks
        list_parser = subparsers.add_parser("list", help="List tasks")
        list_parser.add_argument("status", nargs="?", choices=["todo", "in-progress", "done"],
                               help="Filter tasks by status")

        return parser

    def run(self) -> None:
        parser = self.setup_parser()
        args = parser.parse_args()

        if not args.command:
            parser.print_help()
            return

        try:
            if args.command == "add":
                task = self.task_manager.add_task(args.description)
                print(f"Task added successfully (ID: {task.id})")

            elif args.command == "update":
                task = self.task_manager.update_task(args.id, args.description)
                if task:
                    print(f"Task {task.id} updated successfully")
                else:
                    print(f"Task {args.id} not found")

            elif args.command == "delete":
                if self.task_manager.delete_task(args.id):
                    print(f"Task {args.id} deleted successfully")
                else:
                    print(f"Task {args.id} not found")

            elif args.command == "mark-in-progress":
                task = self.task_manager.mark_task_status(args.id, TaskStatus.IN_PROGRESS)
                if task:
                    print(f"Task {task.id} marked as in progress")
                else:
                    print(f"Task {args.id} not found")

            elif args.command == "mark-done":
                task = self.task_manager.mark_task_status(args.id, TaskStatus.DONE)
                if task:
                    print(f"Task {task.id} marked as done")
                else:
                    print(f"Task {args.id} not found")

            elif args.command == "list":
                status = TaskStatus(args.status) if args.status else None
                tasks = self.task_manager.list_tasks(status)
                if tasks:
                    for task in tasks:
                        print(f"[{task.id}] {task.description} ({task.status.value})")
                else:
                    print("No tasks found")

        except Exception as e:
            print(f"Error: {str(e)}")
