# CodeLoom

## Purpose
This is supposed to be a sample project to behave as a ticketing system that also allows for repo actions


### Ticketing System
1. CRUD on `tickets`
2. Buckets for `development`, `staging`, `testing`, `production` as defaults with support for +5 custom stages in between.



### PR Management System 

All GitHub API references are here[https://docs.github.com/en/rest/pulls?apiVersion=2022-11-28]

#### Basic PR interactions
1. List all open PRs.
2. List all commits on a PR.
3. Merge a PR.
4. Check if a PR has been merged. 

#### Tier 1 PR interactions
1. List all review comments in a PR
2. CRUD a particular comment for a PR
3. Reply to a review comment on a PR

#### Tier 2 PR interactions
1. List reviews for a pull request
2. Create a review for a pull request
3. Get a review for a pull request
4. Update a review for a pull request
5. Delete a pending review for a pull request
6. List comments for a pull request review


#### User-related PR interactions
1. List all PRs requested to be reviewed
2. Request review on a PR to an user
3. Remove all requested reviewers from PR
4. Dismiss a review for a pull request
5. Submit a review for a pull request


### Running the project
1. Navigate to cmd/codeloom.
2. Run the server in live-reload mode with `air -c .air.toml` command.


### Tables (MongoDB)

1. Users
    - first_name
    - last_name
    - title (Optional)
    - email
    - is_admin
    - is_reviewer
    - is_active
    - can_merge
    - token (Encrypted)
    - token_end_date
    - password
    - group [List of Groups]
    - date_joined

2. Groups
    - group_name
    - group_description
    - group_permissions [List of Permission objects]

3. Permission
    - permission_name
    - permission_description
    - is_active

4. TicketBuckets
    - id
    - bucket_name
    - bucket_description (Optional)
    - start_date
    - end_date
    - tickets (List of Ticket objects)

5. Tickets
    - ticket_id
    - ticket_name
    - ticket_description
    - attachments (Optional)
    - started_by (User object)
    - start_date
    - end_date
    - assigned_to (List of User objects)
    - notify_on_activity (Boolean | False)
    - notify_before_expiry (Boolean | False)
    - notify_on_expiry (Boolean | False)
    - sent_notification_for_expiry (Boolean | False)


