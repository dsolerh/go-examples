# API description

## Routes

### GET /chores

Retrieves the currently active chores on the system. Can be paginated via params.

### POST /chores

Creates a new chore. Requires an admin role for the user who is performing the task.

### GET /chores/:id

Gets a specific chore details

### PATCH /chores/:id

### PATCH /chores/:id/deactivate

### PATCH /chores/:id/reschedule

Updates the specific details about a chore. Requires an admin role for the user who is performing the task.

### GET /members

### POST /members

### GET /members/:id

### PATCH /members/:id