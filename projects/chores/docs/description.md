# Chores App

## Idea

Being able to create chores for house keeping with a partner, ensuring fearness on the task split.

Offer the possibility to create chores (cleaning fridge, whashing dishes)

Offer the posibility to create members (the people who is going to participate in the chores)

Create a record of completion for each chore to be reviewed by the participants (TBD the method for the review)

Example:

```
Members: Alice and Bob
Chores: Clean Living Room
Records:
    - Alice at `some date` reviewed by `list of reviewers` approved `true|false`
    - Bob at `some date` reviewed by `list of reviewers` approved `true|false`
```

Being able to extract information about the members performance based on their record

Report to a member his status regarding the active chores

Set rules for the chores split, being able to assign chores to specific members only, (by default all members are obliged to participate in non specific chores)

### Chore Definition

- The chore will be defined dinamically by any member with permision to do so.
- Contain a **Schedule** for the performance of the task.
- Can contain a list of members who are assigned to this task (by default all members are assigned to non specific tasks)
- Has a name and description (optional).
- Can contain a priority (optional).
- Can contain an expiration for reviews.
- Can contain an expiration for reviews of extra tasks.
- Can specify rewards upon completion (For the future).

A **Schedule** for the chores is:

- Type (Daily, Monthly, Exact Date).
- Frequency:
    - Days of the week that the chore is to be performed (if daily)
    - Days of the month.
    - Exact date to be performed
- Times per day the chore needs to be completed. Can be an amount or a list of times (Can be used or daily or specific dates)

Note: A chore cannot be deleted, but can be deactivated

The operations tha can be done over the chores are:

- Create a chore.
- Edit the chore.
- Deactivate the chore.
- Complete the chore. (Everyone can complete a chore)

### Members Definition

- All members can by default complete active chores
- Members have a Name, and a role.
- The member roles can be:
  - Admin: can create and edit roles + Reviewer
  - Reviewer: can review the completion of a chore (exept it's own chores)
  - Basic: can only complete chores

Operations that can be done with the members

- Add a member.
- Edit member details (name and other miscs).
- Promote a member (Update the role of a member, Admin only)
- Deactivate a member.

### Chore completion record

**Scheduled Chore**: it's a chore that is scheduled for the day. Upon completion can be reviewed and it will be marqued as _"fully completed"_ if accepted. It not reviewed in the time window specified by the Chore Definition it will be marked as _"partially completed"_.

**Extra Chore**: it's a chore that is not scheduled for the day. Upon completion needs to be reviewed to be marked as _completed_. If not reviewed will stay as pending for a longer time period.

- A chore can be completed based on it's schedule.
- A member can complete all the chores scheduled for the day.
- A member can complete extra chores not scheduled for the day, which are going to valid only if accepted by a reviewer.
- A completed chore needs to be reviewed to be completley accepted in case of a scheduled chore and accepted in an extra chore.

A **Chore Record** has:

- The chore name (and identifier).
- The time of completion.
- The Member who complete it.
- The completion status, which can be:
  - completed
  - partialy completed
  - pending review
  - rejected
- The completion review
- If the chore was completed as an extra (not scheduled)

### Metrics (for the future)
