# todo
A command line ToDo list application

## Compile
go build todo.go

## Commands

Add a new todo
```
todo add "Remove all comments."
  Item added.
```

List all todos
```
todo list
  1: Remove all comments.
```

Complete a todo item
```
todo done 1
  Marked 1 item as done.
```

## Edit to list manually
This todo app simple uses a text file to manage the todo.  This make managing the todo list easy, but doesn't break your flow of working right from the command line or easly inside VIM.

## Useful Tips
Add the 'togo list' command to your bootable shell script to see all your todo's when you start terminal.

todo list

