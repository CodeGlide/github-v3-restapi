package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GeneratedType_Timeline_comment_event represents the GeneratedType_Timeline_comment_event schema from the OpenAPI specification
type GeneratedType_Timeline_comment_event struct {
	Node_id string `json:"node_id"`
	Url string `json:"url"` // URL for the issue comment
	User GeneratedType_Simple_user `json:"user"` // A GitHub user.
	Id int `json:"id"` // Unique identifier of the issue comment
	Updated_at string `json:"updated_at"`
	Body_html string `json:"body_html,omitempty"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Body_text string `json:"body_text,omitempty"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Created_at string `json:"created_at"`
	Issue_url string `json:"issue_url"`
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Html_url string `json:"html_url"`
	Body string `json:"body,omitempty"` // Contents of the issue comment
	Event string `json:"event"`
}

// GeneratedType_Pull_request represents the GeneratedType_Pull_request schema from the OpenAPI specification
type GeneratedType_Pull_request struct {
	Id int64 `json:"id"`
	Assignee GeneratedType_Nullable_simple_user `json:"assignee"` // A GitHub user.
	Statuses_url string `json:"statuses_url"`
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Rebaseable bool `json:"rebaseable,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	State string `json:"state"` // State of this Pull Request. Either `open` or `closed`.
	Review_comments_url string `json:"review_comments_url"`
	Base map[string]interface{} `json:"base"`
	Review_comments int `json:"review_comments"`
	Commits_url string `json:"commits_url"`
	Assignees []GeneratedType_Simple_user `json:"assignees,omitempty"`
	Comments_url string `json:"comments_url"`
	Milestone GeneratedType_Nullable_milestone `json:"milestone"` // A collection of related issues and pull requests.
	Node_id string `json:"node_id"`
	Issue_url string `json:"issue_url"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Url string `json:"url"`
	Links map[string]interface{} `json:"_links"`
	Merged bool `json:"merged"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Commits int `json:"commits"`
	Updated_at string `json:"updated_at"`
	User GeneratedType_Simple_user `json:"user"` // A GitHub user.
	Mergeable bool `json:"mergeable"`
	Title string `json:"title"` // The title of the pull request.
	Changed_files int `json:"changed_files"`
	Deletions int `json:"deletions"`
	Mergeable_state string `json:"mergeable_state"`
	Number int `json:"number"` // Number uniquely identifying the pull request within its repository.
	Diff_url string `json:"diff_url"`
	Additions int `json:"additions"`
	Locked bool `json:"locked"`
	Requested_reviewers []GeneratedType_Simple_user `json:"requested_reviewers,omitempty"`
	Head map[string]interface{} `json:"head"`
	Maintainer_can_modify bool `json:"maintainer_can_modify"` // Indicates whether maintainers can modify the pull request.
	Body string `json:"body"`
	Requested_teams []GeneratedType_Team_simple `json:"requested_teams,omitempty"`
	Labels []map[string]interface{} `json:"labels"`
	Auto_merge GeneratedType_Auto_merge `json:"auto_merge"` // The status of auto merging a pull request.
	Closed_at string `json:"closed_at"`
	Patch_url string `json:"patch_url"`
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	Merged_by GeneratedType_Nullable_simple_user `json:"merged_by"` // A GitHub user.
	Review_comment_url string `json:"review_comment_url"`
	Comments int `json:"comments"`
	Merged_at string `json:"merged_at"`
}

// GeneratedType_Copilot_ide_chat represents the GeneratedType_Copilot_ide_chat schema from the OpenAPI specification
type GeneratedType_Copilot_ide_chat struct {
	Editors []map[string]interface{} `json:"editors,omitempty"`
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // Total number of users who prompted Copilot Chat in the IDE.
}

// GeneratedType_Codespaces_public_key represents the GeneratedType_Codespaces_public_key schema from the OpenAPI specification
type GeneratedType_Codespaces_public_key struct {
	Key_id string `json:"key_id"` // The identifier for the key.
	Title string `json:"title,omitempty"`
	Url string `json:"url,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Id int `json:"id,omitempty"`
	Key string `json:"key"` // The Base64 encoded public key.
}

// GeneratedType_Secret_scanning_location_pull_request_review_comment represents the GeneratedType_Secret_scanning_location_pull_request_review_comment schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_pull_request_review_comment struct {
	Pull_request_review_comment_url string `json:"pull_request_review_comment_url"` // The API URL to get the pull request review comment where the secret was detected.
}

// GeneratedType_Repository_rule_detailed represents the GeneratedType_Repository_rule_detailed schema from the OpenAPI specification
type GeneratedType_Repository_rule_detailed struct {
}

// Webhooksworkflowjobrun represents the Webhooksworkflowjobrun schema from the OpenAPI specification
type Webhooksworkflowjobrun struct {
	Name interface{} `json:"name"`
	Status string `json:"status"`
	Updated_at string `json:"updated_at"`
	Conclusion interface{} `json:"conclusion"`
	Created_at string `json:"created_at"`
	Environment string `json:"environment"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
}

// GeneratedType_Webhook_project_column_created represents the GeneratedType_Webhook_project_column_created schema from the OpenAPI specification
type GeneratedType_Webhook_project_column_created struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_column Webhooksprojectcolumn `json:"project_column"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// Reaction represents the Reaction schema from the OpenAPI specification
type Reaction struct {
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Content string `json:"content"` // The reaction to use
	Created_at string `json:"created_at"`
}

// GeneratedType_Branch_restriction_policy represents the GeneratedType_Branch_restriction_policy schema from the OpenAPI specification
type GeneratedType_Branch_restriction_policy struct {
	Url string `json:"url"`
	Users []map[string]interface{} `json:"users"`
	Users_url string `json:"users_url"`
	Apps []map[string]interface{} `json:"apps"`
	Apps_url string `json:"apps_url"`
	Teams []map[string]interface{} `json:"teams"`
	Teams_url string `json:"teams_url"`
}

// GeneratedType_Team_discussion represents the GeneratedType_Team_discussion schema from the OpenAPI specification
type GeneratedType_Team_discussion struct {
	Author GeneratedType_Nullable_simple_user `json:"author"` // A GitHub user.
	Private bool `json:"private"` // Whether or not this discussion should be restricted to team members and organization owners.
	Title string `json:"title"` // The title of the discussion.
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Comments_count int `json:"comments_count"`
	Body_version string `json:"body_version"` // The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
	Number int `json:"number"` // The unique sequence number of a team discussion.
	Last_edited_at string `json:"last_edited_at"`
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Url string `json:"url"`
	Body_html string `json:"body_html"`
	Html_url string `json:"html_url"`
	Team_url string `json:"team_url"`
	Body string `json:"body"` // The main text of the discussion.
	Comments_url string `json:"comments_url"`
	Node_id string `json:"node_id"`
	Pinned bool `json:"pinned"` // Whether or not this discussion should be pinned for easy retrieval.
}

// Activity represents the Activity schema from the OpenAPI specification
type Activity struct {
	Timestamp string `json:"timestamp"` // The time when the activity occurred.
	Activity_type string `json:"activity_type"` // The type of the activity that was performed.
	Actor GeneratedType_Nullable_simple_user `json:"actor"` // A GitHub user.
	After string `json:"after"` // The SHA of the commit after the activity.
	Before string `json:"before"` // The SHA of the commit before the activity.
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Ref string `json:"ref"` // The full Git reference, formatted as `refs/heads/<branch name>`.
}

// Webhooksteam represents the Webhooksteam schema from the OpenAPI specification
type Webhooksteam struct {
	Deleted bool `json:"deleted,omitempty"`
	Members_url string `json:"members_url,omitempty"`
	Parent map[string]interface{} `json:"parent,omitempty"`
	Permission string `json:"permission,omitempty"` // Permission that the team will have for its repositories
	Repositories_url string `json:"repositories_url,omitempty"`
	Privacy string `json:"privacy,omitempty"`
	Slug string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"` // Description of the team
	Id int `json:"id"` // Unique identifier of the team
	Notification_setting string `json:"notification_setting,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Name string `json:"name"` // Name of the team
	Node_id string `json:"node_id,omitempty"`
	Url string `json:"url,omitempty"` // URL for the team
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	Description string `json:"description"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Node_id string `json:"node_id"`
	State string `json:"state"`
	Target_url string `json:"target_url"`
	Context string `json:"context"`
	Creator GeneratedType_Nullable_simple_user `json:"creator"` // A GitHub user.
	Avatar_url string `json:"avatar_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
}

// Team represents the Team schema from the OpenAPI specification
type Team struct {
	Description string `json:"description"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Privacy string `json:"privacy,omitempty"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Notification_setting string `json:"notification_setting,omitempty"`
	Slug string `json:"slug"`
	Members_url string `json:"members_url"`
	Parent GeneratedType_Nullable_team_simple `json:"parent"` // Groups of organization members that gives permissions on specified repositories.
	Permission string `json:"permission"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Repositories_url string `json:"repositories_url"`
}

// GeneratedType_Repository_rule_file_extension_restriction represents the GeneratedType_Repository_rule_file_extension_restriction schema from the OpenAPI specification
type GeneratedType_Repository_rule_file_extension_restriction struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType_Webhook_gollum represents the GeneratedType_Webhook_gollum schema from the OpenAPI specification
type GeneratedType_Webhook_gollum struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pages []map[string]interface{} `json:"pages"` // The pages that were updated.
}

// GeneratedType_Codespaces_permissions_check_for_devcontainer represents the GeneratedType_Codespaces_permissions_check_for_devcontainer schema from the OpenAPI specification
type GeneratedType_Codespaces_permissions_check_for_devcontainer struct {
	Accepted bool `json:"accepted"` // Whether the user has accepted the permissions defined by the devcontainer config
}

// Webhooksprojectchanges represents the Webhooksprojectchanges schema from the OpenAPI specification
type Webhooksprojectchanges struct {
	Archived_at map[string]interface{} `json:"archived_at,omitempty"`
}

// GeneratedType_Webhook_repository_renamed represents the GeneratedType_Webhook_repository_renamed schema from the OpenAPI specification
type GeneratedType_Webhook_repository_renamed struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_dependabot_alert_auto_dismissed represents the GeneratedType_Webhook_dependabot_alert_auto_dismissed schema from the OpenAPI specification
type GeneratedType_Webhook_dependabot_alert_auto_dismissed struct {
	Alert GeneratedType_Dependabot_alert `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Webhooksmembership represents the Webhooksmembership schema from the OpenAPI specification
type Webhooksmembership struct {
	State string `json:"state"`
	Url string `json:"url"`
	User map[string]interface{} `json:"user"`
	Organization_url string `json:"organization_url"`
	Role string `json:"role"`
}

// Webhooksmilestone3 represents the Webhooksmilestone3 schema from the OpenAPI specification
type Webhooksmilestone3 struct {
	Updated_at string `json:"updated_at"`
	Creator map[string]interface{} `json:"creator"`
	Closed_at string `json:"closed_at"`
	Closed_issues int `json:"closed_issues"`
	Url string `json:"url"`
	Number int `json:"number"` // The number of the milestone.
	Id int `json:"id"`
	Open_issues int `json:"open_issues"`
	State string `json:"state"` // The state of the milestone.
	Description string `json:"description"`
	Due_on string `json:"due_on"`
	Node_id string `json:"node_id"`
	Title string `json:"title"` // The title of the milestone.
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Labels_url string `json:"labels_url"`
}

// GeneratedType_Secret_scanning_location_issue_body represents the GeneratedType_Secret_scanning_location_issue_body schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_issue_body struct {
	Issue_body_url string `json:"issue_body_url"` // The API URL to get the issue where the secret was detected.
}

// GeneratedType_Webhook_release_prereleased represents the GeneratedType_Webhook_release_prereleased schema from the OpenAPI specification
type GeneratedType_Webhook_release_prereleased struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release map[string]interface{} `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_discussion_reopened represents the GeneratedType_Webhook_discussion_reopened schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_reopened struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Validation_error represents the GeneratedType_Validation_error schema from the OpenAPI specification
type GeneratedType_Validation_error struct {
	Documentation_url string `json:"documentation_url"`
	Errors []map[string]interface{} `json:"errors,omitempty"`
	Message string `json:"message"`
}

// GeneratedType_Webhook_check_run_rerequested represents the GeneratedType_Webhook_check_run_rerequested schema from the OpenAPI specification
type GeneratedType_Webhook_check_run_rerequested struct {
	Action string `json:"action,omitempty"`
	Check_run GeneratedType_Check_run_with_simple_check_suite `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// Discussion represents the Discussion schema from the OpenAPI specification
type Discussion struct {
	Comments int `json:"comments"`
	Id int `json:"id"`
	State string `json:"state"` // The current state of the discussion. `converting` means that the discussion is being converted from an issue. `transferring` means that the discussion is being transferred from another repository.
	Active_lock_reason string `json:"active_lock_reason"`
	Answer_chosen_by map[string]interface{} `json:"answer_chosen_by"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Html_url string `json:"html_url"`
	Answer_chosen_at string `json:"answer_chosen_at"`
	Answer_html_url string `json:"answer_html_url"`
	State_reason string `json:"state_reason"` // The reason for the current state
	Timeline_url string `json:"timeline_url,omitempty"`
	User map[string]interface{} `json:"user"`
	Body string `json:"body"`
	Labels []Label `json:"labels,omitempty"`
	Locked bool `json:"locked"`
	Category map[string]interface{} `json:"category"`
	Number int `json:"number"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Repository_url string `json:"repository_url"`
	Title string `json:"title"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
}

// GeneratedType_Branch_short represents the GeneratedType_Branch_short schema from the OpenAPI specification
type GeneratedType_Branch_short struct {
	Protected bool `json:"protected"`
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
}

// GeneratedType_Actions_hosted_runner_machine_spec represents the GeneratedType_Actions_hosted_runner_machine_spec schema from the OpenAPI specification
type GeneratedType_Actions_hosted_runner_machine_spec struct {
	Id string `json:"id"` // The ID used for the `size` parameter when creating a new runner.
	Memory_gb int `json:"memory_gb"` // The available RAM for the machine spec.
	Storage_gb int `json:"storage_gb"` // The available SSD storage for the machine spec.
	Cpu_cores int `json:"cpu_cores"` // The number of cores.
}

// GeneratedType_Webhook_dependabot_alert_reopened represents the GeneratedType_Webhook_dependabot_alert_reopened schema from the OpenAPI specification
type GeneratedType_Webhook_dependabot_alert_reopened struct {
	Action string `json:"action"`
	Alert GeneratedType_Dependabot_alert `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_ping represents the GeneratedType_Webhook_ping schema from the OpenAPI specification
type GeneratedType_Webhook_ping struct {
	Hook map[string]interface{} `json:"hook,omitempty"` // The webhook that is being pinged
	Hook_id int `json:"hook_id,omitempty"` // The ID of the webhook that triggered the ping.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Zen string `json:"zen,omitempty"` // Random string of GitHub zen.
}

// Webhooksreview represents the Webhooksreview schema from the OpenAPI specification
type Webhooksreview struct {
	Id int `json:"id"` // Unique identifier of the review
	State string `json:"state"`
	User map[string]interface{} `json:"user"`
	Body string `json:"body"` // The text of the review.
	Commit_id string `json:"commit_id"` // A commit SHA for the review.
	Node_id string `json:"node_id"`
	Html_url string `json:"html_url"`
	Pull_request_url string `json:"pull_request_url"`
	Submitted_at string `json:"submitted_at"`
	Links map[string]interface{} `json:"_links"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
}

// GeneratedType_Organization_actions_secret represents the GeneratedType_Organization_actions_secret schema from the OpenAPI specification
type GeneratedType_Organization_actions_secret struct {
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"`
	Visibility string `json:"visibility"` // Visibility of a secret
}

// GeneratedType_Webhook_member_removed represents the GeneratedType_Webhook_member_removed schema from the OpenAPI specification
type GeneratedType_Webhook_member_removed struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Member Webhooksuser `json:"member"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Sub_issues_summary represents the GeneratedType_Sub_issues_summary schema from the OpenAPI specification
type GeneratedType_Sub_issues_summary struct {
	Percent_completed int `json:"percent_completed"`
	Total int `json:"total"`
	Completed int `json:"completed"`
}

// GeneratedType_Webhook_workflow_job_completed represents the GeneratedType_Webhook_workflow_job_completed schema from the OpenAPI specification
type GeneratedType_Webhook_workflow_job_completed struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Workflow_job interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Starred_repository represents the GeneratedType_Starred_repository schema from the OpenAPI specification
type GeneratedType_Starred_repository struct {
	Starred_at string `json:"starred_at"`
	Repo Repository `json:"repo"` // A repository on GitHub.
}

// Runner represents the Runner schema from the OpenAPI specification
type Runner struct {
	Id int `json:"id"` // The ID of the runner.
	Labels []GeneratedType_Runner_label `json:"labels"`
	Name string `json:"name"` // The name of the runner.
	Os string `json:"os"` // The Operating System of the runner.
	Runner_group_id int `json:"runner_group_id,omitempty"` // The ID of the runner group.
	Status string `json:"status"` // The status of the runner.
	Busy bool `json:"busy"`
	Ephemeral bool `json:"ephemeral,omitempty"`
}

// GeneratedType_Issue_event_project_card represents the GeneratedType_Issue_event_project_card schema from the OpenAPI specification
type GeneratedType_Issue_event_project_card struct {
	Url string `json:"url"`
	Column_name string `json:"column_name"`
	Id int `json:"id"`
	Previous_column_name string `json:"previous_column_name,omitempty"`
	Project_id int `json:"project_id"`
	Project_url string `json:"project_url"`
}

// GeneratedType_Code_scanning_alert_instance represents the GeneratedType_Code_scanning_alert_instance schema from the OpenAPI specification
type GeneratedType_Code_scanning_alert_instance struct {
	Environment string `json:"environment,omitempty"` // Identifies the variable values associated with the environment in which the analysis that generated this alert instance was performed, such as the language that was analyzed.
	State string `json:"state,omitempty"` // State of a code scanning alert.
	Analysis_key string `json:"analysis_key,omitempty"` // Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
	Category string `json:"category,omitempty"` // Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
	Classifications []string `json:"classifications,omitempty"` // Classifications that have been applied to the file that triggered the alert. For example identifying it as documentation, or a generated file.
	Location GeneratedType_Code_scanning_alert_location `json:"location,omitempty"` // Describe a region within a file for the alert.
	Ref string `json:"ref,omitempty"` // The Git reference, formatted as `refs/pull/<number>/merge`, `refs/pull/<number>/head`, `refs/heads/<branch name>` or simply `<branch name>`.
	Commit_sha string `json:"commit_sha,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Message map[string]interface{} `json:"message,omitempty"`
}

// GeneratedType_Webhook_member_added represents the GeneratedType_Webhook_member_added schema from the OpenAPI specification
type GeneratedType_Webhook_member_added struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Member Webhooksuser `json:"member"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_branch_protection_rule_deleted represents the GeneratedType_Webhook_branch_protection_rule_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_branch_protection_rule_deleted struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Rule Webhooksrule `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Projects_v2_status_update represents the GeneratedType_Projects_v2_status_update schema from the OpenAPI specification
type GeneratedType_Projects_v2_status_update struct {
	Updated_at string `json:"updated_at"`
	Node_id string `json:"node_id"`
	Status string `json:"status,omitempty"`
	Target_date string `json:"target_date,omitempty"`
	Project_node_id string `json:"project_node_id,omitempty"`
	Start_date string `json:"start_date,omitempty"`
	Body string `json:"body,omitempty"` // Body of the status update
	Created_at string `json:"created_at"`
	Creator GeneratedType_Simple_user `json:"creator,omitempty"` // A GitHub user.
	Id float64 `json:"id"`
}

// GeneratedType_Simple_installation represents the GeneratedType_Simple_installation schema from the OpenAPI specification
type GeneratedType_Simple_installation struct {
	Id int `json:"id"` // The ID of the installation.
	Node_id string `json:"node_id"` // The global node ID of the installation.
}

// GeneratedType_Basic_error represents the GeneratedType_Basic_error schema from the OpenAPI specification
type GeneratedType_Basic_error struct {
	Documentation_url string `json:"documentation_url,omitempty"`
	Message string `json:"message,omitempty"`
	Status string `json:"status,omitempty"`
	Url string `json:"url,omitempty"`
}

// GeneratedType_Webhook_installation_repositories_removed represents the GeneratedType_Webhook_installation_repositories_removed schema from the OpenAPI specification
type GeneratedType_Webhook_installation_repositories_removed struct {
	Repositories_added []map[string]interface{} `json:"repositories_added"` // An array of repository objects, which were added to the installation.
	Repositories_removed []map[string]interface{} `json:"repositories_removed"` // An array of repository objects, which were removed from the installation.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Requester Webhooksuser `json:"requester"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Issue_event_dismissed_review represents the GeneratedType_Issue_event_dismissed_review schema from the OpenAPI specification
type GeneratedType_Issue_event_dismissed_review struct {
	Review_id int `json:"review_id"`
	State string `json:"state"`
	Dismissal_commit_id string `json:"dismissal_commit_id,omitempty"`
	Dismissal_message string `json:"dismissal_message"`
}

// Webhooksworkflow represents the Webhooksworkflow schema from the OpenAPI specification
type Webhooksworkflow struct {
	Created_at string `json:"created_at"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Path string `json:"path"`
	State string `json:"state"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Badge_url string `json:"badge_url"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
}

// GeneratedType_Issue_event represents the GeneratedType_Issue_event schema from the OpenAPI specification
type GeneratedType_Issue_event struct {
	Actor GeneratedType_Nullable_simple_user `json:"actor"` // A GitHub user.
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Assignee GeneratedType_Nullable_simple_user `json:"assignee,omitempty"` // A GitHub user.
	Author_association string `json:"author_association,omitempty"` // How the author is associated with the repository.
	Created_at string `json:"created_at"`
	Milestone GeneratedType_Issue_event_milestone `json:"milestone,omitempty"` // Issue Event Milestone
	Commit_id string `json:"commit_id"`
	Url string `json:"url"`
	Event string `json:"event"`
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Dismissed_review GeneratedType_Issue_event_dismissed_review `json:"dismissed_review,omitempty"`
	Issue GeneratedType_Nullable_issue `json:"issue,omitempty"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Project_card GeneratedType_Issue_event_project_card `json:"project_card,omitempty"` // Issue Event Project Card
	Id int64 `json:"id"`
	Assigner GeneratedType_Nullable_simple_user `json:"assigner,omitempty"` // A GitHub user.
	Node_id string `json:"node_id"`
	Review_requester GeneratedType_Nullable_simple_user `json:"review_requester,omitempty"` // A GitHub user.
	Label GeneratedType_Issue_event_label `json:"label,omitempty"` // Issue Event Label
	Lock_reason string `json:"lock_reason,omitempty"`
	Requested_reviewer GeneratedType_Nullable_simple_user `json:"requested_reviewer,omitempty"` // A GitHub user.
	Rename GeneratedType_Issue_event_rename `json:"rename,omitempty"` // Issue Event Rename
	Commit_url string `json:"commit_url"`
}

// GeneratedType_Code_scanning_alert represents the GeneratedType_Code_scanning_alert schema from the OpenAPI specification
type GeneratedType_Code_scanning_alert struct {
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Rule GeneratedType_Code_scanning_alert_rule `json:"rule"`
	Dismissed_by GeneratedType_Nullable_simple_user `json:"dismissed_by"` // A GitHub user.
	Number int `json:"number"` // The security alert number.
	Dismissal_approved_by GeneratedType_Nullable_simple_user `json:"dismissal_approved_by,omitempty"` // A GitHub user.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Tool GeneratedType_Code_scanning_analysis_tool `json:"tool"`
	State string `json:"state"` // State of a code scanning alert.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Most_recent_instance GeneratedType_Code_scanning_alert_instance `json:"most_recent_instance"`
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
}

// GeneratedType_Secret_scanning_scan_history represents the GeneratedType_Secret_scanning_scan_history schema from the OpenAPI specification
type GeneratedType_Secret_scanning_scan_history struct {
	Backfill_scans []GeneratedType_Secret_scanning_scan `json:"backfill_scans,omitempty"`
	Custom_pattern_backfill_scans []interface{} `json:"custom_pattern_backfill_scans,omitempty"`
	Incremental_scans []GeneratedType_Secret_scanning_scan `json:"incremental_scans,omitempty"`
	Pattern_update_scans []GeneratedType_Secret_scanning_scan `json:"pattern_update_scans,omitempty"`
}

// Integration represents the Integration schema from the OpenAPI specification
type Integration struct {
	Description string `json:"description"`
	Events []string `json:"events"` // The list of events for the GitHub app. Note that the `installation_target`, `security_advisory`, and `meta` events are not included because they are global events and not specific to an installation.
	Html_url string `json:"html_url"`
	Installations_count int `json:"installations_count,omitempty"` // The number of installations associated with the GitHub app. Only returned when the integration is requesting details about itself.
	Owner interface{} `json:"owner"`
	Client_id string `json:"client_id,omitempty"`
	Id int `json:"id"` // Unique identifier of the GitHub app
	External_url string `json:"external_url"`
	Name string `json:"name"` // The name of the GitHub app
	Node_id string `json:"node_id"`
	Permissions map[string]interface{} `json:"permissions"` // The set of permissions for the GitHub app
	Slug string `json:"slug,omitempty"` // The slug name of the GitHub app
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
}

// Feed represents the Feed schema from the OpenAPI specification
type Feed struct {
	Current_user_organization_urls []string `json:"current_user_organization_urls,omitempty"`
	Repository_discussions_category_url string `json:"repository_discussions_category_url,omitempty"` // A feed of discussions for a given repository and category.
	User_url string `json:"user_url"`
	Current_user_organization_url string `json:"current_user_organization_url,omitempty"`
	Current_user_public_url string `json:"current_user_public_url,omitempty"`
	Current_user_url string `json:"current_user_url,omitempty"`
	Links map[string]interface{} `json:"_links"`
	Repository_discussions_url string `json:"repository_discussions_url,omitempty"` // A feed of discussions for a given repository.
	Security_advisories_url string `json:"security_advisories_url,omitempty"`
	Timeline_url string `json:"timeline_url"`
	Current_user_actor_url string `json:"current_user_actor_url,omitempty"`
}

// GeneratedType_Code_security_configuration_for_repository represents the GeneratedType_Code_security_configuration_for_repository schema from the OpenAPI specification
type GeneratedType_Code_security_configuration_for_repository struct {
	Configuration GeneratedType_Code_security_configuration `json:"configuration,omitempty"` // A code security configuration
	Status string `json:"status,omitempty"` // The attachment status of the code security configuration on the repository.
}

// Root represents the Root schema from the OpenAPI specification
type Root struct {
	Events_url string `json:"events_url"`
	Organization_url string `json:"organization_url"`
	Starred_gists_url string `json:"starred_gists_url"`
	Current_user_url string `json:"current_user_url"`
	Emails_url string `json:"emails_url"`
	Commit_search_url string `json:"commit_search_url"`
	Current_user_repositories_url string `json:"current_user_repositories_url"`
	Keys_url string `json:"keys_url"`
	Gists_url string `json:"gists_url"`
	Issues_url string `json:"issues_url"`
	Feeds_url string `json:"feeds_url"`
	Hub_url string `json:"hub_url,omitempty"`
	Repository_search_url string `json:"repository_search_url"`
	Emojis_url string `json:"emojis_url"`
	Organization_teams_url string `json:"organization_teams_url"`
	Public_gists_url string `json:"public_gists_url"`
	Topic_search_url string `json:"topic_search_url,omitempty"`
	Rate_limit_url string `json:"rate_limit_url"`
	Organization_repositories_url string `json:"organization_repositories_url"`
	Followers_url string `json:"followers_url"`
	User_organizations_url string `json:"user_organizations_url"`
	Notifications_url string `json:"notifications_url"`
	Starred_url string `json:"starred_url"`
	Authorizations_url string `json:"authorizations_url"`
	Code_search_url string `json:"code_search_url"`
	Current_user_authorizations_html_url string `json:"current_user_authorizations_html_url"`
	User_search_url string `json:"user_search_url"`
	User_repositories_url string `json:"user_repositories_url"`
	Repository_url string `json:"repository_url"`
	User_url string `json:"user_url"`
	Following_url string `json:"following_url"`
	Issue_search_url string `json:"issue_search_url"`
	Label_search_url string `json:"label_search_url"`
}

// GeneratedType_Label_search_result_item represents the GeneratedType_Label_search_result_item schema from the OpenAPI specification
type GeneratedType_Label_search_result_item struct {
	DefaultField bool `json:"default"`
	Id int `json:"id"`
	Name string `json:"name"`
	Score float64 `json:"score"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Color string `json:"color"`
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
}

// GeneratedType_Check_automated_security_fixes represents the GeneratedType_Check_automated_security_fixes schema from the OpenAPI specification
type GeneratedType_Check_automated_security_fixes struct {
	Enabled bool `json:"enabled"` // Whether Dependabot security updates are enabled for the repository.
	Paused bool `json:"paused"` // Whether Dependabot security updates are paused for the repository.
}

// GeneratedType_Webhook_organization_deleted represents the GeneratedType_Webhook_organization_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_organization_deleted struct {
	Membership Webhooksmembership `json:"membership,omitempty"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Billing_usage_report represents the GeneratedType_Billing_usage_report schema from the OpenAPI specification
type GeneratedType_Billing_usage_report struct {
	Usageitems []map[string]interface{} `json:"usageItems,omitempty"`
}

// GeneratedType_Nullable_simple_user represents the GeneratedType_Nullable_simple_user schema from the OpenAPI specification
type GeneratedType_Nullable_simple_user struct {
	Avatar_url string `json:"avatar_url"`
	Email string `json:"email,omitempty"`
	Gists_url string `json:"gists_url"`
	Login string `json:"login"`
	Gravatar_id string `json:"gravatar_id"`
	Site_admin bool `json:"site_admin"`
	Starred_at string `json:"starred_at,omitempty"`
	Subscriptions_url string `json:"subscriptions_url"`
	Followers_url string `json:"followers_url"`
	Url string `json:"url"`
	User_view_type string `json:"user_view_type,omitempty"`
	Events_url string `json:"events_url"`
	Id int64 `json:"id"`
	Node_id string `json:"node_id"`
	Organizations_url string `json:"organizations_url"`
	Starred_url string `json:"starred_url"`
	Following_url string `json:"following_url"`
	Html_url string `json:"html_url"`
	Received_events_url string `json:"received_events_url"`
	Name string `json:"name,omitempty"`
	Repos_url string `json:"repos_url"`
	TypeField string `json:"type"`
}

// GeneratedType_Organization_programmatic_access_grant represents the GeneratedType_Organization_programmatic_access_grant schema from the OpenAPI specification
type GeneratedType_Organization_programmatic_access_grant struct {
	Access_granted_at string `json:"access_granted_at"` // Date and time when the fine-grained personal access token was approved to access the organization.
	Token_last_used_at string `json:"token_last_used_at"` // Date and time when the associated fine-grained personal access token was last used for authentication.
	Repository_selection string `json:"repository_selection"` // Type of repository selection requested.
	Token_expired bool `json:"token_expired"` // Whether the associated fine-grained personal access token has expired.
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Repositories_url string `json:"repositories_url"` // URL to the list of repositories the fine-grained personal access token can access. Only follow when `repository_selection` is `subset`.
	Token_id int `json:"token_id"` // Unique identifier of the user's token. This field can also be found in audit log events and the organization's settings for their PAT grants.
	Id int `json:"id"` // Unique identifier of the fine-grained personal access token grant. The `pat_id` used to get details about an approved fine-grained personal access token.
	Permissions map[string]interface{} `json:"permissions"` // Permissions requested, categorized by type of permission.
	Token_expires_at string `json:"token_expires_at"` // Date and time when the associated fine-grained personal access token expires.
	Token_name string `json:"token_name"` // The name given to the user's token. This field can also be found in an organization's settings page for Active Tokens.
}

// GeneratedType_Repository_rule_violation_error represents the GeneratedType_Repository_rule_violation_error schema from the OpenAPI specification
type GeneratedType_Repository_rule_violation_error struct {
	Documentation_url string `json:"documentation_url,omitempty"`
	Message string `json:"message,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Status string `json:"status,omitempty"`
}

// GeneratedType_Code_security_configuration_repositories represents the GeneratedType_Code_security_configuration_repositories schema from the OpenAPI specification
type GeneratedType_Code_security_configuration_repositories struct {
	Repository GeneratedType_Simple_repository `json:"repository,omitempty"` // A GitHub repository.
	Status string `json:"status,omitempty"` // The attachment status of the code security configuration on the repository.
}

// GeneratedType_Repository_rule_params_restricted_commits represents the GeneratedType_Repository_rule_params_restricted_commits schema from the OpenAPI specification
type GeneratedType_Repository_rule_params_restricted_commits struct {
	Oid string `json:"oid"` // Full or abbreviated commit hash to reject
	Reason string `json:"reason,omitempty"` // Reason for restriction
}

// GeneratedType_Repository_ruleset_conditions_repository_name_target represents the GeneratedType_Repository_ruleset_conditions_repository_name_target schema from the OpenAPI specification
type GeneratedType_Repository_ruleset_conditions_repository_name_target struct {
	Repository_name map[string]interface{} `json:"repository_name"`
}

// GeneratedType_Review_custom_gates_state_required represents the GeneratedType_Review_custom_gates_state_required schema from the OpenAPI specification
type GeneratedType_Review_custom_gates_state_required struct {
	Environment_name string `json:"environment_name"` // The name of the environment to approve or reject.
	State string `json:"state"` // Whether to approve or reject deployment to the specified environments.
	Comment string `json:"comment,omitempty"` // Optional comment to include with the review.
}

// GeneratedType_Issue_event_rename represents the GeneratedType_Issue_event_rename schema from the OpenAPI specification
type GeneratedType_Issue_event_rename struct {
	To string `json:"to"`
	From string `json:"from"`
}

// Vulnerability represents the Vulnerability schema from the OpenAPI specification
type Vulnerability struct {
	PackageField map[string]interface{} `json:"package"` // The name of the package affected by the vulnerability.
	Vulnerable_functions []string `json:"vulnerable_functions"` // The functions in the package that are affected by the vulnerability.
	Vulnerable_version_range string `json:"vulnerable_version_range"` // The range of the package versions affected by the vulnerability.
	First_patched_version string `json:"first_patched_version"` // The package version that resolves the vulnerability.
}

// GeneratedType_Webhook_installation_target_renamed represents the GeneratedType_Webhook_installation_target_renamed schema from the OpenAPI specification
type GeneratedType_Webhook_installation_target_renamed struct {
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Target_type string `json:"target_type"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Account map[string]interface{} `json:"account"`
}

// GeneratedType_Webhook_check_run_requested_action represents the GeneratedType_Webhook_check_run_requested_action schema from the OpenAPI specification
type GeneratedType_Webhook_check_run_requested_action struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requested_action map[string]interface{} `json:"requested_action,omitempty"` // The action requested by the user.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Check_run GeneratedType_Check_run_with_simple_check_suite `json:"check_run"` // A check performed on the code of a given code change
}

// GeneratedType_Repository_rule_update represents the GeneratedType_Repository_rule_update schema from the OpenAPI specification
type GeneratedType_Repository_rule_update struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType_Webhook_project_card_moved represents the GeneratedType_Webhook_project_card_moved schema from the OpenAPI specification
type GeneratedType_Webhook_project_card_moved struct {
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card interface{} `json:"project_card"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Repository_rule_required_deployments represents the GeneratedType_Repository_rule_required_deployments schema from the OpenAPI specification
type GeneratedType_Repository_rule_required_deployments struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType_Repository_rule_file_path_restriction represents the GeneratedType_Repository_rule_file_path_restriction schema from the OpenAPI specification
type GeneratedType_Repository_rule_file_path_restriction struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType_Code_scanning_default_setup represents the GeneratedType_Code_scanning_default_setup schema from the OpenAPI specification
type GeneratedType_Code_scanning_default_setup struct {
	Schedule string `json:"schedule,omitempty"` // The frequency of the periodic analysis.
	State string `json:"state,omitempty"` // Code scanning default setup has been configured or not.
	Threat_model string `json:"threat_model,omitempty"` // Threat model to be used for code scanning analysis. Use `remote` to analyze only network sources and `remote_and_local` to include local sources like filesystem access, command-line arguments, database reads, environment variable and standard input.
	Updated_at string `json:"updated_at,omitempty"` // Timestamp of latest configuration update.
	Languages []string `json:"languages,omitempty"` // Languages to be analyzed.
	Query_suite string `json:"query_suite,omitempty"` // CodeQL query suite to be used.
	Runner_label string `json:"runner_label,omitempty"` // Runner label to be used if the runner type is labeled.
	Runner_type string `json:"runner_type,omitempty"` // Runner type to be used.
}

// GeneratedType_Webhook_deployment_review_rejected represents the GeneratedType_Webhook_deployment_review_rejected schema from the OpenAPI specification
type GeneratedType_Webhook_deployment_review_rejected struct {
	Workflow_job_runs []map[string]interface{} `json:"workflow_job_runs,omitempty"`
	Comment string `json:"comment,omitempty"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Workflow_job_run Webhooksworkflowjobrun `json:"workflow_job_run,omitempty"`
	Approver Webhooksapprover `json:"approver,omitempty"`
	Reviewers []map[string]interface{} `json:"reviewers,omitempty"`
	Since string `json:"since"`
}

// GeneratedType_Project_column represents the GeneratedType_Project_column schema from the OpenAPI specification
type GeneratedType_Project_column struct {
	Project_url string `json:"project_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Cards_url string `json:"cards_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // The unique identifier of the project column
	Name string `json:"name"` // Name of the project column
	Node_id string `json:"node_id"`
}

// GeneratedType_Deployment_status represents the GeneratedType_Deployment_status schema from the OpenAPI specification
type GeneratedType_Deployment_status struct {
	Url string `json:"url"`
	Creator GeneratedType_Nullable_simple_user `json:"creator"` // A GitHub user.
	Repository_url string `json:"repository_url"`
	Id int64 `json:"id"`
	Updated_at string `json:"updated_at"`
	Deployment_url string `json:"deployment_url"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Target_url string `json:"target_url"` // Closing down notice: the URL to associate with this status.
	Created_at string `json:"created_at"`
	Description string `json:"description"` // A short description of the status.
	Environment string `json:"environment,omitempty"` // The environment of the deployment that the status is for.
	Log_url string `json:"log_url,omitempty"` // The URL to associate with this status.
	State string `json:"state"` // The state of the status.
	Environment_url string `json:"environment_url,omitempty"` // The URL for accessing your environment.
}

// GeneratedType_Webhook_repository_advisory_reported represents the GeneratedType_Webhook_repository_advisory_reported schema from the OpenAPI specification
type GeneratedType_Webhook_repository_advisory_reported struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_advisory GeneratedType_Repository_advisory `json:"repository_advisory"` // A repository security advisory.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_User_role_assignment represents the GeneratedType_User_role_assignment schema from the OpenAPI specification
type GeneratedType_User_role_assignment struct {
	TypeField string `json:"type"`
	Avatar_url string `json:"avatar_url"`
	Inherited_from []GeneratedType_Team_simple `json:"inherited_from,omitempty"` // Team the user has gotten the role through
	Name string `json:"name,omitempty"`
	Starred_at string `json:"starred_at,omitempty"`
	Repos_url string `json:"repos_url"`
	Node_id string `json:"node_id"`
	Gravatar_id string `json:"gravatar_id"`
	Login string `json:"login"`
	Following_url string `json:"following_url"`
	Site_admin bool `json:"site_admin"`
	Subscriptions_url string `json:"subscriptions_url"`
	Assignment string `json:"assignment,omitempty"` // Determines if the user has a direct, indirect, or mixed relationship to a role
	Html_url string `json:"html_url"`
	Followers_url string `json:"followers_url"`
	Id int `json:"id"`
	Gists_url string `json:"gists_url"`
	Organizations_url string `json:"organizations_url"`
	Received_events_url string `json:"received_events_url"`
	Starred_url string `json:"starred_url"`
	Url string `json:"url"`
	User_view_type string `json:"user_view_type,omitempty"`
	Events_url string `json:"events_url"`
	Email string `json:"email,omitempty"`
}

// GeneratedType_Issue_event_milestone represents the GeneratedType_Issue_event_milestone schema from the OpenAPI specification
type GeneratedType_Issue_event_milestone struct {
	Title string `json:"title"`
}

// GeneratedType_Webhook_pull_request_review_thread_unresolved represents the GeneratedType_Webhook_pull_request_review_thread_unresolved schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_review_thread_unresolved struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Thread map[string]interface{} `json:"thread"`
	Action string `json:"action"`
}

// GeneratedType_Webhook_code_scanning_alert_created represents the GeneratedType_Webhook_code_scanning_alert_created schema from the OpenAPI specification
type GeneratedType_Webhook_code_scanning_alert_created struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_pull_request_review_thread_resolved represents the GeneratedType_Webhook_pull_request_review_thread_resolved schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_review_thread_resolved struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Thread map[string]interface{} `json:"thread"`
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
}

// Verification represents the Verification schema from the OpenAPI specification
type Verification struct {
	Payload string `json:"payload"`
	Reason string `json:"reason"`
	Signature string `json:"signature"`
	Verified bool `json:"verified"`
	Verified_at string `json:"verified_at"`
}

// GeneratedType_Secret_scanning_alert represents the GeneratedType_Secret_scanning_alert schema from the OpenAPI specification
type GeneratedType_Secret_scanning_alert struct {
	First_location_detected interface{} `json:"first_location_detected,omitempty"` // Details on the location where the token was initially detected. This can be a commit, wiki commit, issue, discussion, pull request.
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Supported secret scanning patterns](https://docs.github.com/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
	Multi_repo bool `json:"multi_repo,omitempty"` // Whether the detected secret was found in multiple repositories under the same organization or enterprise.
	Publicly_leaked bool `json:"publicly_leaked,omitempty"` // Whether the detected secret was publicly leaked.
	Secret string `json:"secret,omitempty"` // The secret that was detected.
	Resolution_comment string `json:"resolution_comment,omitempty"` // An optional comment to resolve an alert.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Has_more_locations bool `json:"has_more_locations,omitempty"` // A boolean value representing whether or not the token in the alert was detected in more than one location.
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
	Push_protection_bypass_request_comment string `json:"push_protection_bypass_request_comment,omitempty"` // An optional comment when requesting a push protection bypass.
	Push_protection_bypassed_by GeneratedType_Nullable_simple_user `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Resolved_by GeneratedType_Nullable_simple_user `json:"resolved_by,omitempty"` // A GitHub user.
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	State string `json:"state,omitempty"` // Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
	Validity string `json:"validity,omitempty"` // The token status as of the latest validity check.
	Number int `json:"number,omitempty"` // The security alert number.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Push_protection_bypass_request_html_url string `json:"push_protection_bypass_request_html_url,omitempty"` // The URL to a push protection bypass request.
	Push_protection_bypass_request_reviewer_comment string `json:"push_protection_bypass_request_reviewer_comment,omitempty"` // An optional comment when reviewing a push protection bypass.
	Is_base64_encoded bool `json:"is_base64_encoded,omitempty"` // A boolean value representing whether or not alert is base64 encoded
	Push_protection_bypass_request_reviewer GeneratedType_Nullable_simple_user `json:"push_protection_bypass_request_reviewer,omitempty"` // A GitHub user.
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Resolution string `json:"resolution,omitempty"` // **Required when the `state` is `resolved`.** The reason for resolving the alert.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType_Webhook_issues_closed represents the GeneratedType_Webhook_issues_closed schema from the OpenAPI specification
type GeneratedType_Webhook_issues_closed struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action that was performed.
}

// GeneratedType_Nullable_community_health_file represents the GeneratedType_Nullable_community_health_file schema from the OpenAPI specification
type GeneratedType_Nullable_community_health_file struct {
	Url string `json:"url"`
	Html_url string `json:"html_url"`
}

// GeneratedType_Check_run_with_simple_check_suite represents the GeneratedType_Check_run_with_simple_check_suite schema from the OpenAPI specification
type GeneratedType_Check_run_with_simple_check_suite struct {
	Check_suite GeneratedType_Simple_check_suite `json:"check_suite"` // A suite of checks performed on the code of a given code change
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being checked.
	Node_id string `json:"node_id"`
	Pull_requests []GeneratedType_Pull_request_minimal `json:"pull_requests"`
	Started_at string `json:"started_at"`
	Id int `json:"id"` // The id of the check.
	Status string `json:"status"` // The phase of the lifecycle that the check is currently in.
	Url string `json:"url"`
	Name string `json:"name"` // The name of the check.
	Conclusion string `json:"conclusion"`
	Completed_at string `json:"completed_at"`
	External_id string `json:"external_id"`
	Deployment GeneratedType_Deployment_simple `json:"deployment,omitempty"` // A deployment created as the result of an Actions check run from a workflow that references an environment
	Details_url string `json:"details_url"`
	Output map[string]interface{} `json:"output"`
	Html_url string `json:"html_url"`
	App Integration `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
}

// GeneratedType_Combined_billing_usage represents the GeneratedType_Combined_billing_usage schema from the OpenAPI specification
type GeneratedType_Combined_billing_usage struct {
	Days_left_in_billing_cycle int `json:"days_left_in_billing_cycle"` // Numbers of days left in billing cycle.
	Estimated_paid_storage_for_month int `json:"estimated_paid_storage_for_month"` // Estimated storage space (GB) used in billing cycle.
	Estimated_storage_for_month int `json:"estimated_storage_for_month"` // Estimated sum of free and paid storage space (GB) used in billing cycle.
}

// GeneratedType_Public_ip represents the GeneratedType_Public_ip schema from the OpenAPI specification
type GeneratedType_Public_ip struct {
	Enabled bool `json:"enabled,omitempty"` // Whether public IP is enabled.
	Length int `json:"length,omitempty"` // The length of the IP prefix.
	Prefix string `json:"prefix,omitempty"` // The prefix for the public IP.
}

// GeneratedType_Dependabot_alert_with_repository represents the GeneratedType_Dependabot_alert_with_repository schema from the OpenAPI specification
type GeneratedType_Dependabot_alert_with_repository struct {
	Repository GeneratedType_Simple_repository `json:"repository"` // A GitHub repository.
	State string `json:"state"` // The state of the Dependabot alert.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Fixed_at string `json:"fixed_at"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Auto_dismissed_at string `json:"auto_dismissed_at,omitempty"` // The time that the alert was auto-dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Number int `json:"number"` // The security alert number.
	Dismissed_comment string `json:"dismissed_comment"` // An optional comment associated with the alert's dismissal.
	Security_advisory GeneratedType_Dependabot_alert_security_advisory `json:"security_advisory"` // Details for the GitHub Security Advisory.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_by GeneratedType_Nullable_simple_user `json:"dismissed_by"` // A GitHub user.
	Security_vulnerability GeneratedType_Dependabot_alert_security_vulnerability `json:"security_vulnerability"` // Details pertaining to one vulnerable version range for the advisory.
	Updated_at string `json:"updated_at"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dependency map[string]interface{} `json:"dependency"` // Details for the vulnerable dependency.
	Dismissed_reason string `json:"dismissed_reason"` // The reason that the alert was dismissed.
}

// GeneratedType_Unassigned_issue_event represents the GeneratedType_Unassigned_issue_event schema from the OpenAPI specification
type GeneratedType_Unassigned_issue_event struct {
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Event string `json:"event"`
	Id int `json:"id"`
	Assigner GeneratedType_Simple_user `json:"assigner"` // A GitHub user.
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Assignee GeneratedType_Simple_user `json:"assignee"` // A GitHub user.
	Commit_id string `json:"commit_id"`
}

// GeneratedType_Authentication_token represents the GeneratedType_Authentication_token schema from the OpenAPI specification
type GeneratedType_Authentication_token struct {
	Token string `json:"token"` // The token used for authentication
	Expires_at string `json:"expires_at"` // The time this token expires
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Repositories []Repository `json:"repositories,omitempty"` // The repositories this token has access to
	Repository_selection string `json:"repository_selection,omitempty"` // Describe whether all repositories have been selected or there's a selection involved
	Single_file string `json:"single_file,omitempty"`
}

// GeneratedType_Webhook_pull_request_milestoned represents the GeneratedType_Webhook_pull_request_milestoned schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_milestoned struct {
	Milestone Milestone `json:"milestone,omitempty"` // A collection of related issues and pull requests.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request Webhookspullrequest5 `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Content_tree represents the GeneratedType_Content_tree schema from the OpenAPI specification
type GeneratedType_Content_tree struct {
	Encoding string `json:"encoding,omitempty"`
	Git_url string `json:"git_url"`
	Name string `json:"name"`
	Sha string `json:"sha"`
	Links map[string]interface{} `json:"_links"`
	Size int `json:"size"`
	TypeField string `json:"type"`
	Content string `json:"content,omitempty"`
	Entries []map[string]interface{} `json:"entries,omitempty"`
	Html_url string `json:"html_url"`
	Path string `json:"path"`
	Download_url string `json:"download_url"`
	Url string `json:"url"`
}

// GeneratedType_Webhook_pull_request_ready_for_review represents the GeneratedType_Webhook_pull_request_ready_for_review schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_ready_for_review struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType_Pull_request_webhook `json:"pull_request"`
}

// GeneratedType_Webhook_sponsorship_pending_cancellation represents the GeneratedType_Webhook_sponsorship_pending_cancellation schema from the OpenAPI specification
type GeneratedType_Webhook_sponsorship_pending_cancellation struct {
	Effective_date string `json:"effective_date,omitempty"` // The `pending_cancellation` and `pending_tier_change` event types will include the date the cancellation or tier change will take effect.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
}

// GeneratedType_Repository_ruleset_conditions_repository_property_spec represents the GeneratedType_Repository_ruleset_conditions_repository_property_spec schema from the OpenAPI specification
type GeneratedType_Repository_ruleset_conditions_repository_property_spec struct {
	Name string `json:"name"` // The name of the repository property to target
	Property_values []string `json:"property_values"` // The values to match for the repository property
	Source string `json:"source,omitempty"` // The source of the repository property. Defaults to 'custom' if not specified.
}

// Email represents the Email schema from the OpenAPI specification
type Email struct {
	Email string `json:"email"`
	Primary bool `json:"primary"`
	Verified bool `json:"verified"`
	Visibility string `json:"visibility"`
}

// GeneratedType_Webhook_pull_request_reopened represents the GeneratedType_Webhook_pull_request_reopened schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_reopened struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType_Pull_request_webhook `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Protected_branch_admin_enforced represents the GeneratedType_Protected_branch_admin_enforced schema from the OpenAPI specification
type GeneratedType_Protected_branch_admin_enforced struct {
	Enabled bool `json:"enabled"`
	Url string `json:"url"`
}

// GeneratedType_Porter_large_file represents the GeneratedType_Porter_large_file schema from the OpenAPI specification
type GeneratedType_Porter_large_file struct {
	Path string `json:"path"`
	Ref_name string `json:"ref_name"`
	Size int `json:"size"`
	Oid string `json:"oid"`
}

// GeneratedType_Code_scanning_alert_rule represents the GeneratedType_Code_scanning_alert_rule schema from the OpenAPI specification
type GeneratedType_Code_scanning_alert_rule struct {
	Security_severity_level string `json:"security_severity_level,omitempty"` // The security severity of the alert.
	Tags []string `json:"tags,omitempty"` // A set of tags applicable for the rule.
	Help string `json:"help,omitempty"` // Detailed documentation for the rule as GitHub Flavored Markdown.
	Severity string `json:"severity,omitempty"` // The severity of the alert.
	Full_description string `json:"full_description,omitempty"` // A description of the rule used to detect the alert.
	Name string `json:"name,omitempty"` // The name of the rule used to detect the alert.
	Description string `json:"description,omitempty"` // A short description of the rule used to detect the alert.
	Help_uri string `json:"help_uri,omitempty"` // A link to the documentation for the rule used to detect the alert.
	Id string `json:"id,omitempty"` // A unique identifier for the rule used to detect the alert.
}

// GeneratedType_Secret_scanning_location_pull_request_comment represents the GeneratedType_Secret_scanning_location_pull_request_comment schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_pull_request_comment struct {
	Pull_request_comment_url string `json:"pull_request_comment_url"` // The API URL to get the pull request comment where the secret was detected.
}

// GeneratedType_Webhook_discussion_deleted represents the GeneratedType_Webhook_discussion_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_deleted struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_issues_typed represents the GeneratedType_Webhook_issues_typed schema from the OpenAPI specification
type GeneratedType_Webhook_issues_typed struct {
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	TypeField GeneratedType_Issue_type `json:"type"` // The type of issue.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_merge_group_destroyed represents the GeneratedType_Webhook_merge_group_destroyed schema from the OpenAPI specification
type GeneratedType_Webhook_merge_group_destroyed struct {
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Merge_group GeneratedType_Merge_group `json:"merge_group"` // A group of pull requests that the merge queue has grouped together to be merged.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Reason string `json:"reason,omitempty"` // Explains why the merge group is being destroyed. The group could have been merged, removed from the queue (dequeued), or invalidated by an earlier queue entry being dequeued (invalidated).
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType_Check_suite represents the GeneratedType_Check_suite schema from the OpenAPI specification
type GeneratedType_Check_suite struct {
	Check_runs_url string `json:"check_runs_url"`
	Conclusion string `json:"conclusion"`
	Runs_rerequestable bool `json:"runs_rerequestable,omitempty"`
	Before string `json:"before"`
	Head_sha string `json:"head_sha"` // The SHA of the head commit that is being checked.
	Node_id string `json:"node_id"`
	Pull_requests []GeneratedType_Pull_request_minimal `json:"pull_requests"`
	Repository GeneratedType_Minimal_repository `json:"repository"` // Minimal Repository
	App GeneratedType_Nullable_integration `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Id int64 `json:"id"`
	Head_commit GeneratedType_Simple_commit `json:"head_commit"` // A commit.
	Head_branch string `json:"head_branch"`
	Latest_check_runs_count int `json:"latest_check_runs_count"`
	Status string `json:"status"` // The phase of the lifecycle that the check suite is currently in. Statuses of waiting, requested, and pending are reserved for GitHub Actions check suites.
	After string `json:"after"`
	Created_at string `json:"created_at"`
	Rerequestable bool `json:"rerequestable,omitempty"`
	Url string `json:"url"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType_Webhook_discussion_locked represents the GeneratedType_Webhook_discussion_locked schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_locked struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
}

// GeneratedType_Webhook_deployment_review_requested represents the GeneratedType_Webhook_deployment_review_requested schema from the OpenAPI specification
type GeneratedType_Webhook_deployment_review_requested struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Since string `json:"since"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Reviewers []map[string]interface{} `json:"reviewers"`
	Workflow_job_run map[string]interface{} `json:"workflow_job_run"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Requestor Webhooksuser `json:"requestor"`
	Action string `json:"action"`
	Environment string `json:"environment"`
}

// GeneratedType_Nullable_codespace_machine represents the GeneratedType_Nullable_codespace_machine schema from the OpenAPI specification
type GeneratedType_Nullable_codespace_machine struct {
	Operating_system string `json:"operating_system"` // The operating system of the machine.
	Prebuild_availability string `json:"prebuild_availability"` // Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be "null" if prebuilds are not supported or prebuild availability could not be determined. Value will be "none" if no prebuild is available. Latest values "ready" and "in_progress" indicate the prebuild availability status.
	Storage_in_bytes int `json:"storage_in_bytes"` // How much storage is available to the codespace.
	Cpus int `json:"cpus"` // How many cores are available to the codespace.
	Display_name string `json:"display_name"` // The display name of the machine includes cores, memory, and storage.
	Memory_in_bytes int `json:"memory_in_bytes"` // How much memory is available to the codespace.
	Name string `json:"name"` // The name of the machine.
}

// GeneratedType_Webhook_membership_added represents the GeneratedType_Webhook_membership_added schema from the OpenAPI specification
type GeneratedType_Webhook_membership_added struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Member Webhooksuser `json:"member"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Sender map[string]interface{} `json:"sender"`
	Scope string `json:"scope"` // The scope of the membership. Currently, can only be `team`.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Team Webhooksteam `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
}

// GeneratedType_Webhook_secret_scanning_alert_resolved represents the GeneratedType_Webhook_secret_scanning_alert_resolved schema from the OpenAPI specification
type GeneratedType_Webhook_secret_scanning_alert_resolved struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType_Secret_scanning_alert_webhook `json:"alert"`
}

// GeneratedType_Copilot_ide_code_completions represents the GeneratedType_Copilot_ide_code_completions schema from the OpenAPI specification
type GeneratedType_Copilot_ide_code_completions struct {
	Editors []map[string]interface{} `json:"editors,omitempty"`
	Languages []map[string]interface{} `json:"languages,omitempty"` // Code completion metrics for active languages.
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // Number of users who accepted at least one Copilot code suggestion, across all active editors. Includes both full and partial acceptances.
}

// GeneratedType_Secret_scanning_location_pull_request_review represents the GeneratedType_Secret_scanning_location_pull_request_review schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_pull_request_review struct {
	Pull_request_review_url string `json:"pull_request_review_url"` // The API URL to get the pull request review where the secret was detected.
}

// Topic represents the Topic schema from the OpenAPI specification
type Topic struct {
	Names []string `json:"names"`
}

// GeneratedType_Webhook_marketplace_purchase_changed represents the GeneratedType_Webhook_marketplace_purchase_changed schema from the OpenAPI specification
type GeneratedType_Webhook_marketplace_purchase_changed struct {
	Effective_date string `json:"effective_date"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Action string `json:"action"`
}

// GeneratedType_Webhook_installation_unsuspend represents the GeneratedType_Webhook_installation_unsuspend schema from the OpenAPI specification
type GeneratedType_Webhook_installation_unsuspend struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
}

// GeneratedType_Code_scanning_analysis represents the GeneratedType_Code_scanning_analysis schema from the OpenAPI specification
type GeneratedType_Code_scanning_analysis struct {
	Rules_count int `json:"rules_count"` // The total number of rules used in the analysis.
	Url string `json:"url"` // The REST API URL of the analysis resource.
	Category string `json:"category,omitempty"` // Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
	Id int `json:"id"` // Unique identifier for this analysis.
	Environment string `json:"environment"` // Identifies the variable values associated with the environment in which this analysis was performed.
	Ref string `json:"ref"` // The Git reference, formatted as `refs/pull/<number>/merge`, `refs/pull/<number>/head`, `refs/heads/<branch name>` or simply `<branch name>`.
	Results_count int `json:"results_count"` // The total number of results in the analysis.
	ErrorField string `json:"error"`
	Deletable bool `json:"deletable"`
	Sarif_id string `json:"sarif_id"` // An identifier for the upload.
	Commit_sha string `json:"commit_sha"` // The SHA of the commit to which the analysis you are uploading relates.
	Created_at string `json:"created_at"` // The time that the analysis was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Tool GeneratedType_Code_scanning_analysis_tool `json:"tool"`
	Analysis_key string `json:"analysis_key"` // Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
	Warning string `json:"warning"` // Warning generated when processing the analysis
}

// Artifact represents the Artifact schema from the OpenAPI specification
type Artifact struct {
	Digest string `json:"digest,omitempty"` // The SHA256 digest of the artifact. This field will only be populated on artifacts uploaded with upload-artifact v4 or newer. For older versions, this field will be null.
	Expired bool `json:"expired"` // Whether or not the artifact has expired.
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Expires_at string `json:"expires_at"`
	Name string `json:"name"` // The name of the artifact.
	Updated_at string `json:"updated_at"`
	Archive_download_url string `json:"archive_download_url"`
	Size_in_bytes int `json:"size_in_bytes"` // The size in bytes of the artifact.
	Workflow_run map[string]interface{} `json:"workflow_run,omitempty"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
}

// GeneratedType_Webhook_repository_deleted represents the GeneratedType_Webhook_repository_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_repository_deleted struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_projects_v2_project_reopened represents the GeneratedType_Webhook_projects_v2_project_reopened schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_project_reopened struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2 GeneratedType_Projects_v2 `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_project_column_edited represents the GeneratedType_Webhook_project_column_edited schema from the OpenAPI specification
type GeneratedType_Webhook_project_column_edited struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_column Webhooksprojectcolumn `json:"project_column"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType_Repository_rule_merge_queue represents the GeneratedType_Repository_rule_merge_queue schema from the OpenAPI specification
type GeneratedType_Repository_rule_merge_queue struct {
	TypeField string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType_Webhook_secret_scanning_alert_validated represents the GeneratedType_Webhook_secret_scanning_alert_validated schema from the OpenAPI specification
type GeneratedType_Webhook_secret_scanning_alert_validated struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType_Secret_scanning_alert_webhook `json:"alert"`
}

// GeneratedType_Timeline_cross_referenced_event represents the GeneratedType_Timeline_cross_referenced_event schema from the OpenAPI specification
type GeneratedType_Timeline_cross_referenced_event struct {
	Source map[string]interface{} `json:"source"`
	Updated_at string `json:"updated_at"`
	Actor GeneratedType_Simple_user `json:"actor,omitempty"` // A GitHub user.
	Created_at string `json:"created_at"`
	Event string `json:"event"`
}

// GeneratedType_Webhook_deploy_key_deleted represents the GeneratedType_Webhook_deploy_key_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_deploy_key_deleted struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Key Webhooksdeploykey `json:"key"` // The [`deploy key`](https://docs.github.com/rest/deploy-keys/deploy-keys#get-a-deploy-key) resource.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_pull_request_review_edited represents the GeneratedType_Webhook_pull_request_review_edited schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_review_edited struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Review Webhooksreview `json:"review"` // The review that was affected.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Changes map[string]interface{} `json:"changes"`
}

// GeneratedType_Git_ref represents the GeneratedType_Git_ref schema from the OpenAPI specification
type GeneratedType_Git_ref struct {
	Node_id string `json:"node_id"`
	Object map[string]interface{} `json:"object"`
	Ref string `json:"ref"`
	Url string `json:"url"`
}

// GeneratedType_Webhook_milestone_opened represents the GeneratedType_Webhook_milestone_opened schema from the OpenAPI specification
type GeneratedType_Webhook_milestone_opened struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone3 `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Deployment_branch_policy_name_pattern represents the GeneratedType_Deployment_branch_policy_name_pattern schema from the OpenAPI specification
type GeneratedType_Deployment_branch_policy_name_pattern struct {
	Name string `json:"name"` // The name pattern that branches must match in order to deploy to the environment. Wildcard characters will not match `/`. For example, to match branches that begin with `release/` and contain an additional single slash, use `release/*/*`. For more information about pattern matching syntax, see the [Ruby File.fnmatch documentation](https://ruby-doc.org/core-2.5.1/File.html#method-c-fnmatch).
}

// GeneratedType_Global_advisory represents the GeneratedType_Global_advisory schema from the OpenAPI specification
type GeneratedType_Global_advisory struct {
	Nvd_published_at string `json:"nvd_published_at"` // The date and time when the advisory was published in the National Vulnerability Database, in ISO 8601 format. This field is only populated when the advisory is imported from the National Vulnerability Database.
	References []string `json:"references"`
	Url string `json:"url"` // The API URL for the advisory.
	Html_url string `json:"html_url"` // The URL for the advisory.
	Updated_at string `json:"updated_at"` // The date and time of when the advisory was last updated, in ISO 8601 format.
	Cve_id string `json:"cve_id"` // The Common Vulnerabilities and Exposures (CVE) ID.
	Identifiers []map[string]interface{} `json:"identifiers"`
	Source_code_location string `json:"source_code_location"` // The URL of the advisory's source code.
	Epss GeneratedType_Security_advisory_epss `json:"epss,omitempty"` // The EPSS scores as calculated by the [Exploit Prediction Scoring System](https://www.first.org/epss).
	Ghsa_id string `json:"ghsa_id"` // The GitHub Security Advisory ID.
	Cvss map[string]interface{} `json:"cvss"`
	Repository_advisory_url string `json:"repository_advisory_url"` // The API URL for the repository advisory.
	Summary string `json:"summary"` // A short summary of the advisory.
	Cwes []map[string]interface{} `json:"cwes"`
	Description string `json:"description"` // A detailed description of what the advisory entails.
	Published_at string `json:"published_at"` // The date and time of when the advisory was published, in ISO 8601 format.
	TypeField string `json:"type"` // The type of advisory.
	Credits []map[string]interface{} `json:"credits"` // The users who contributed to the advisory.
	Cvss_severities GeneratedType_Cvss_severities `json:"cvss_severities,omitempty"`
	Github_reviewed_at string `json:"github_reviewed_at"` // The date and time of when the advisory was reviewed by GitHub, in ISO 8601 format.
	Severity string `json:"severity"` // The severity of the advisory.
	Vulnerabilities []Vulnerability `json:"vulnerabilities"` // The products and respective version ranges affected by the advisory.
	Withdrawn_at string `json:"withdrawn_at"` // The date and time of when the advisory was withdrawn, in ISO 8601 format.
}

// Language represents the Language schema from the OpenAPI specification
type Language struct {
}

// GeneratedType_Issue_event_label represents the GeneratedType_Issue_event_label schema from the OpenAPI specification
type GeneratedType_Issue_event_label struct {
	Color string `json:"color"`
	Name string `json:"name"`
}

// GeneratedType_Webhook_workflow_job_in_progress represents the GeneratedType_Webhook_workflow_job_in_progress schema from the OpenAPI specification
type GeneratedType_Webhook_workflow_job_in_progress struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Workflow_job interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
}

// Webhooksmarketplacepurchase represents the Webhooksmarketplacepurchase schema from the OpenAPI specification
type Webhooksmarketplacepurchase struct {
	Next_billing_date string `json:"next_billing_date"`
	On_free_trial bool `json:"on_free_trial"`
	Plan map[string]interface{} `json:"plan"`
	Unit_count int `json:"unit_count"`
	Account map[string]interface{} `json:"account"`
	Billing_cycle string `json:"billing_cycle"`
	Free_trial_ends_on string `json:"free_trial_ends_on"`
}

// GeneratedType_Webhook_organization_member_invited represents the GeneratedType_Webhook_organization_member_invited schema from the OpenAPI specification
type GeneratedType_Webhook_organization_member_invited struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	User Webhooksuser `json:"user,omitempty"`
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Invitation map[string]interface{} `json:"invitation"` // The invitation for the user or email if the action is `member_invited`.
}

// GeneratedType_Webhook_sponsorship_edited represents the GeneratedType_Webhook_sponsorship_edited schema from the OpenAPI specification
type GeneratedType_Webhook_sponsorship_edited struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
}

// GeneratedType_Enterprise_team represents the GeneratedType_Enterprise_team schema from the OpenAPI specification
type GeneratedType_Enterprise_team struct {
	Group_id string `json:"group_id,omitempty"`
	Name string `json:"name"`
	Sync_to_organizations string `json:"sync_to_organizations,omitempty"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Id int64 `json:"id"`
	Members_url string `json:"members_url"`
	Organization_selection_type string `json:"organization_selection_type,omitempty"`
	Slug string `json:"slug"`
	Group_name string `json:"group_name,omitempty"`
	Html_url string `json:"html_url"`
	Description string `json:"description,omitempty"`
	Url string `json:"url"`
}

// GeneratedType_Classroom_assignment represents the GeneratedType_Classroom_assignment schema from the OpenAPI specification
type GeneratedType_Classroom_assignment struct {
	Invite_link string `json:"invite_link"` // The link that a student can use to accept the assignment.
	Max_members int `json:"max_members"` // The maximum allowable members per team.
	Max_teams int `json:"max_teams"` // The maximum allowable teams for the assignment.
	Deadline string `json:"deadline"` // The time at which the assignment is due.
	Public_repo bool `json:"public_repo"` // Whether an accepted assignment creates a public repository.
	TypeField string `json:"type"` // Whether it's a group assignment or individual assignment.
	Editor string `json:"editor"` // The selected editor for the assignment.
	Students_are_repo_admins bool `json:"students_are_repo_admins"` // Whether students are admins on created repository when a student accepts the assignment.
	Feedback_pull_requests_enabled bool `json:"feedback_pull_requests_enabled"` // Whether feedback pull request will be created when a student accepts the assignment.
	Invitations_enabled bool `json:"invitations_enabled"` // Whether the invitation link is enabled. Visiting an enabled invitation link will accept the assignment.
	Submitted int `json:"submitted"` // The number of students that have submitted the assignment.
	Title string `json:"title"` // Assignment title.
	Starter_code_repository GeneratedType_Simple_classroom_repository `json:"starter_code_repository"` // A GitHub repository view for Classroom
	Accepted int `json:"accepted"` // The number of students that have accepted the assignment.
	Classroom Classroom `json:"classroom"` // A GitHub Classroom classroom
	Language string `json:"language"` // The programming language used in the assignment.
	Id int `json:"id"` // Unique identifier of the repository.
	Slug string `json:"slug"` // Sluggified name of the assignment.
	Passing int `json:"passing"` // The number of students that have passed the assignment.
}

// GeneratedType_Webhook_issues_unpinned represents the GeneratedType_Webhook_issues_unpinned schema from the OpenAPI specification
type GeneratedType_Webhook_issues_unpinned struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue2 `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Webhook_repository_edited represents the GeneratedType_Webhook_repository_edited schema from the OpenAPI specification
type GeneratedType_Webhook_repository_edited struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Packages_billing_usage represents the GeneratedType_Packages_billing_usage schema from the OpenAPI specification
type GeneratedType_Packages_billing_usage struct {
	Total_gigabytes_bandwidth_used int `json:"total_gigabytes_bandwidth_used"` // Sum of the free and paid storage space (GB) for GitHuub Packages.
	Total_paid_gigabytes_bandwidth_used int `json:"total_paid_gigabytes_bandwidth_used"` // Total paid storage space (GB) for GitHuub Packages.
	Included_gigabytes_bandwidth int `json:"included_gigabytes_bandwidth"` // Free storage space (GB) for GitHub Packages.
}

// GeneratedType_Webhook_workflow_run_requested represents the GeneratedType_Webhook_workflow_run_requested schema from the OpenAPI specification
type GeneratedType_Webhook_workflow_run_requested struct {
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Workflow Webhooksworkflow `json:"workflow"`
}

// GeneratedType_Simple_commit represents the GeneratedType_Simple_commit schema from the OpenAPI specification
type GeneratedType_Simple_commit struct {
	Tree_id string `json:"tree_id"` // SHA for the commit's tree
	Author map[string]interface{} `json:"author"` // Information about the Git author
	Committer map[string]interface{} `json:"committer"` // Information about the Git committer
	Id string `json:"id"` // SHA for the commit
	Message string `json:"message"` // Message describing the purpose of the commit
	Timestamp string `json:"timestamp"` // Timestamp of the commit
}

// Webhookspullrequest5 represents the Webhookspullrequest5 schema from the OpenAPI specification
type Webhookspullrequest5 struct {
	Review_comments_url string `json:"review_comments_url"`
	Mergeable_state string `json:"mergeable_state,omitempty"`
	Locked bool `json:"locked"`
	Updated_at string `json:"updated_at"`
	Diff_url string `json:"diff_url"`
	Additions int `json:"additions,omitempty"`
	Statuses_url string `json:"statuses_url"`
	Changed_files int `json:"changed_files,omitempty"`
	Merged_by map[string]interface{} `json:"merged_by,omitempty"`
	Patch_url string `json:"patch_url"`
	Assignee map[string]interface{} `json:"assignee"`
	Review_comments int `json:"review_comments,omitempty"`
	Url string `json:"url"`
	Requested_teams []map[string]interface{} `json:"requested_teams"`
	Base map[string]interface{} `json:"base"`
	Closed_at string `json:"closed_at"`
	Auto_merge map[string]interface{} `json:"auto_merge"` // The status of auto merging a pull request.
	Created_at string `json:"created_at"`
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Body string `json:"body"`
	Links map[string]interface{} `json:"_links"`
	Comments_url string `json:"comments_url"`
	Review_comment_url string `json:"review_comment_url"`
	Commits int `json:"commits,omitempty"`
	Requested_reviewers []interface{} `json:"requested_reviewers"`
	Commits_url string `json:"commits_url"`
	Mergeable bool `json:"mergeable,omitempty"`
	Issue_url string `json:"issue_url"`
	Rebaseable bool `json:"rebaseable,omitempty"`
	Merged_at string `json:"merged_at"`
	Html_url string `json:"html_url"`
	State string `json:"state"` // State of this Pull Request. Either `open` or `closed`.
	Title string `json:"title"` // The title of the pull request.
	Head map[string]interface{} `json:"head"`
	Id int `json:"id"`
	Number int `json:"number"` // Number uniquely identifying the pull request within its repository.
	Deletions int `json:"deletions,omitempty"`
	Labels []map[string]interface{} `json:"labels"`
	Node_id string `json:"node_id"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Draft bool `json:"draft"` // Indicates whether or not the pull request is a draft.
	Maintainer_can_modify bool `json:"maintainer_can_modify,omitempty"` // Indicates whether maintainers can modify the pull request.
	User map[string]interface{} `json:"user"`
	Merged bool `json:"merged,omitempty"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Comments int `json:"comments,omitempty"`
	Active_lock_reason string `json:"active_lock_reason"`
	Assignees []map[string]interface{} `json:"assignees"`
}

// GeneratedType_Dependabot_alert_package represents the GeneratedType_Dependabot_alert_package schema from the OpenAPI specification
type GeneratedType_Dependabot_alert_package struct {
	Ecosystem string `json:"ecosystem"` // The package's language or package management ecosystem.
	Name string `json:"name"` // The unique package name within its ecosystem.
}

// GeneratedType_Secret_scanning_location_commit represents the GeneratedType_Secret_scanning_location_commit schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_commit struct {
	Commit_url string `json:"commit_url"` // The API URL to get the associated commit resource
	End_line float64 `json:"end_line"` // Line number at which the secret ends in the file
	Start_column float64 `json:"start_column"` // The column at which the secret starts within the start line when the file is interpreted as 8BIT ASCII
	Path string `json:"path"` // The file path in the repository
	End_column float64 `json:"end_column"` // The column at which the secret ends within the end line when the file is interpreted as 8BIT ASCII
	Start_line float64 `json:"start_line"` // Line number at which the secret starts in the file
	Blob_sha string `json:"blob_sha"` // SHA-1 hash ID of the associated blob
	Blob_url string `json:"blob_url"` // The API URL to get the associated blob resource
	Commit_sha string `json:"commit_sha"` // SHA-1 hash ID of the associated commit
}

// GeneratedType_Webhook_sub_issues_sub_issue_removed represents the GeneratedType_Webhook_sub_issues_sub_issue_removed schema from the OpenAPI specification
type GeneratedType_Webhook_sub_issues_sub_issue_removed struct {
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Sub_issue_repo Repository `json:"sub_issue_repo"` // A repository on GitHub.
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType_Network_settings represents the GeneratedType_Network_settings schema from the OpenAPI specification
type GeneratedType_Network_settings struct {
	Id string `json:"id"` // The unique identifier of the network settings resource.
	Name string `json:"name"` // The name of the network settings resource.
	Network_configuration_id string `json:"network_configuration_id,omitempty"` // The identifier of the network configuration that is using this settings resource.
	Region string `json:"region"` // The location of the subnet this network settings resource is configured for.
	Subnet_id string `json:"subnet_id"` // The subnet this network settings resource is configured for.
}

// GeneratedType_Webhook_milestone_edited represents the GeneratedType_Webhook_milestone_edited schema from the OpenAPI specification
type GeneratedType_Webhook_milestone_edited struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the milestone if the action was `edited`.
}

// GeneratedType_Webhook_fork represents the GeneratedType_Webhook_fork schema from the OpenAPI specification
type GeneratedType_Webhook_fork struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Forkee interface{} `json:"forkee"` // The created [`repository`](https://docs.github.com/rest/repos/repos#get-a-repository) resource.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Code_scanning_alert_rule_summary represents the GeneratedType_Code_scanning_alert_rule_summary schema from the OpenAPI specification
type GeneratedType_Code_scanning_alert_rule_summary struct {
	Full_description string `json:"full_description,omitempty"` // A description of the rule used to detect the alert.
	Help string `json:"help,omitempty"` // Detailed documentation for the rule as GitHub Flavored Markdown.
	Help_uri string `json:"help_uri,omitempty"` // A link to the documentation for the rule used to detect the alert.
	Id string `json:"id,omitempty"` // A unique identifier for the rule used to detect the alert.
	Name string `json:"name,omitempty"` // The name of the rule used to detect the alert.
	Severity string `json:"severity,omitempty"` // The severity of the alert.
	Tags []string `json:"tags,omitempty"` // A set of tags applicable for the rule.
	Description string `json:"description,omitempty"` // A short description of the rule used to detect the alert.
	Security_severity_level string `json:"security_severity_level,omitempty"` // The security severity of the alert.
}

// GeneratedType_Webhook_secret_scanning_scan_completed represents the GeneratedType_Webhook_secret_scanning_scan_completed schema from the OpenAPI specification
type GeneratedType_Webhook_secret_scanning_scan_completed struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Secret_types []string `json:"secret_types,omitempty"` // List of patterns that were updated. This will be empty for normal backfill scans or custom pattern updates
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Source string `json:"source"` // What type of content was scanned
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	TypeField string `json:"type"` // What type of scan was completed
	Action string `json:"action"`
	Completed_at string `json:"completed_at"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Custom_pattern_name string `json:"custom_pattern_name,omitempty"` // If the scan was triggered by a custom pattern update, this will be the name of the pattern that was updated
	Custom_pattern_scope string `json:"custom_pattern_scope,omitempty"` // If the scan was triggered by a custom pattern update, this will be the scope of the pattern that was updated
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Started_at string `json:"started_at"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType_Social_account represents the GeneratedType_Social_account schema from the OpenAPI specification
type GeneratedType_Social_account struct {
	Provider string `json:"provider"`
	Url string `json:"url"`
}

// GeneratedType_Repository_ruleset_bypass_actor represents the GeneratedType_Repository_ruleset_bypass_actor schema from the OpenAPI specification
type GeneratedType_Repository_ruleset_bypass_actor struct {
	Actor_id int `json:"actor_id,omitempty"` // The ID of the actor that can bypass a ruleset. If `actor_type` is `OrganizationAdmin`, this should be `1`. If `actor_type` is `DeployKey`, this should be null. `OrganizationAdmin` is not applicable for personal repositories.
	Actor_type string `json:"actor_type"` // The type of actor that can bypass a ruleset.
	Bypass_mode string `json:"bypass_mode,omitempty"` // When the specified actor can bypass the ruleset. `pull_request` means that an actor can only bypass rules on pull requests. `pull_request` is not applicable for the `DeployKey` actor type. Also, `pull_request` is only applicable to branch rulesets.
}

// GeneratedType_Webhook_repository_ruleset_created represents the GeneratedType_Webhook_repository_ruleset_created schema from the OpenAPI specification
type GeneratedType_Webhook_repository_ruleset_created struct {
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_ruleset GeneratedType_Repository_ruleset `json:"repository_ruleset"` // A set of rules to apply when specified conditions are met.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_pull_request_review_comment_created represents the GeneratedType_Webhook_pull_request_review_comment_created schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_review_comment_created struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request) itself.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
}

// GeneratedType_Repository_advisory_credit represents the GeneratedType_Repository_advisory_credit schema from the OpenAPI specification
type GeneratedType_Repository_advisory_credit struct {
	State string `json:"state"` // The state of the user's acceptance of the credit.
	TypeField string `json:"type"` // The type of credit the user is receiving.
	User GeneratedType_Simple_user `json:"user"` // A GitHub user.
}

// GeneratedType_Webhook_discussion_answered represents the GeneratedType_Webhook_discussion_answered schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_answered struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Answer Webhooksanswer `json:"answer"`
}

// GeneratedType_Webhook_public represents the GeneratedType_Webhook_public schema from the OpenAPI specification
type GeneratedType_Webhook_public struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_milestone_closed represents the GeneratedType_Webhook_milestone_closed schema from the OpenAPI specification
type GeneratedType_Webhook_milestone_closed struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone `json:"milestone"` // A collection of related issues and pull requests.
}

// GeneratedType_Dependabot_alert represents the GeneratedType_Dependabot_alert schema from the OpenAPI specification
type GeneratedType_Dependabot_alert struct {
	Dismissed_by GeneratedType_Nullable_simple_user `json:"dismissed_by"` // A GitHub user.
	Fixed_at string `json:"fixed_at"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_comment string `json:"dismissed_comment"` // An optional comment associated with the alert's dismissal.
	Updated_at string `json:"updated_at"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_reason string `json:"dismissed_reason"` // The reason that the alert was dismissed.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Auto_dismissed_at string `json:"auto_dismissed_at,omitempty"` // The time that the alert was auto-dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dependency map[string]interface{} `json:"dependency"` // Details for the vulnerable dependency.
	Security_advisory GeneratedType_Dependabot_alert_security_advisory `json:"security_advisory"` // Details for the GitHub Security Advisory.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	State string `json:"state"` // The state of the Dependabot alert.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Security_vulnerability GeneratedType_Dependabot_alert_security_vulnerability `json:"security_vulnerability"` // Details pertaining to one vulnerable version range for the advisory.
	Number int `json:"number"` // The security alert number.
}

// GeneratedType_Porter_author represents the GeneratedType_Porter_author schema from the OpenAPI specification
type GeneratedType_Porter_author struct {
	Name string `json:"name"`
	Remote_id string `json:"remote_id"`
	Remote_name string `json:"remote_name"`
	Url string `json:"url"`
	Email string `json:"email"`
	Id int `json:"id"`
	Import_url string `json:"import_url"`
}

// GeneratedType_Repository_subscription represents the GeneratedType_Repository_subscription schema from the OpenAPI specification
type GeneratedType_Repository_subscription struct {
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Ignored bool `json:"ignored"` // Determines if all notifications should be blocked from this repository.
	Reason string `json:"reason"`
	Repository_url string `json:"repository_url"`
	Subscribed bool `json:"subscribed"` // Determines if notifications should be received from this repository.
}

// Webhooksusermannequin represents the Webhooksusermannequin schema from the OpenAPI specification
type Webhooksusermannequin struct {
	Organizations_url string `json:"organizations_url,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Email string `json:"email,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Name string `json:"name,omitempty"`
	Gists_url string `json:"gists_url,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
	Url string `json:"url,omitempty"`
	Login string `json:"login"`
	Id int `json:"id"`
	User_view_type string `json:"user_view_type,omitempty"`
	TypeField string `json:"type,omitempty"`
}

// GeneratedType_Custom_property_value represents the GeneratedType_Custom_property_value schema from the OpenAPI specification
type GeneratedType_Custom_property_value struct {
	Value string `json:"value"` // The value assigned to the property
	Property_name string `json:"property_name"` // The name of the property
}

// GeneratedType_Webhook_project_column_deleted represents the GeneratedType_Webhook_project_column_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_project_column_deleted struct {
	Repository GeneratedType_Nullable_repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_column Webhooksprojectcolumn `json:"project_column"`
}

// GeneratedType_Organization_dependabot_secret represents the GeneratedType_Organization_dependabot_secret schema from the OpenAPI specification
type GeneratedType_Organization_dependabot_secret struct {
	Visibility string `json:"visibility"` // Visibility of a secret
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType_Webhook_branch_protection_rule_created represents the GeneratedType_Webhook_branch_protection_rule_created schema from the OpenAPI specification
type GeneratedType_Webhook_branch_protection_rule_created struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Rule Webhooksrule `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
}

// Environment represents the Environment schema from the OpenAPI specification
type Environment struct {
	Name string `json:"name"` // The name of the environment.
	Protection_rules []interface{} `json:"protection_rules,omitempty"` // Built-in deployment protection rules for the environment.
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Id int64 `json:"id"` // The id of the environment.
	Updated_at string `json:"updated_at"` // The time that the environment was last updated, in ISO 8601 format.
	Created_at string `json:"created_at"` // The time that the environment was created, in ISO 8601 format.
	Deployment_branch_policy GeneratedType_Deployment_branch_policy_settings `json:"deployment_branch_policy,omitempty"` // The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
	Node_id string `json:"node_id"`
}

// GeneratedType_Repository_invitation represents the GeneratedType_Repository_invitation schema from the OpenAPI specification
type GeneratedType_Repository_invitation struct {
	Id int64 `json:"id"` // Unique identifier of the repository invitation.
	Invitee GeneratedType_Nullable_simple_user `json:"invitee"` // A GitHub user.
	Node_id string `json:"node_id"`
	Permissions string `json:"permissions"` // The permission associated with the invitation.
	Repository GeneratedType_Minimal_repository `json:"repository"` // Minimal Repository
	Html_url string `json:"html_url"`
	Url string `json:"url"` // URL for the repository invitation
	Created_at string `json:"created_at"`
	Inviter GeneratedType_Nullable_simple_user `json:"inviter"` // A GitHub user.
	Expired bool `json:"expired,omitempty"` // Whether or not the invitation has expired
}

// GeneratedType_Webhook_sub_issues_sub_issue_added represents the GeneratedType_Webhook_sub_issues_sub_issue_added schema from the OpenAPI specification
type GeneratedType_Webhook_sub_issues_sub_issue_added struct {
	Sub_issue_repo Repository `json:"sub_issue_repo"` // A repository on GitHub.
	Action string `json:"action"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
}

// GeneratedType_Timeline_commit_commented_event represents the GeneratedType_Timeline_commit_commented_event schema from the OpenAPI specification
type GeneratedType_Timeline_commit_commented_event struct {
	Comments []GeneratedType_Commit_comment `json:"comments,omitempty"`
	Commit_id string `json:"commit_id,omitempty"`
	Event string `json:"event,omitempty"`
	Node_id string `json:"node_id,omitempty"`
}

// GeneratedType_Integration_installation_request represents the GeneratedType_Integration_installation_request schema from the OpenAPI specification
type GeneratedType_Integration_installation_request struct {
	Account interface{} `json:"account"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // Unique identifier of the request installation.
	Node_id string `json:"node_id,omitempty"`
	Requester GeneratedType_Simple_user `json:"requester"` // A GitHub user.
}

// GeneratedType_Commit_search_result_item represents the GeneratedType_Commit_search_result_item schema from the OpenAPI specification
type GeneratedType_Commit_search_result_item struct {
	Author GeneratedType_Nullable_simple_user `json:"author"` // A GitHub user.
	Committer GeneratedType_Nullable_git_user `json:"committer"` // Metaproperties for Git author/committer information.
	Html_url string `json:"html_url"`
	Score float64 `json:"score"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Commit map[string]interface{} `json:"commit"`
	Parents []map[string]interface{} `json:"parents"`
	Url string `json:"url"`
	Comments_url string `json:"comments_url"`
	Sha string `json:"sha"`
	Node_id string `json:"node_id"`
	Repository GeneratedType_Minimal_repository `json:"repository"` // Minimal Repository
}

// GeneratedType_Webhook_security_advisory_published represents the GeneratedType_Webhook_security_advisory_published schema from the OpenAPI specification
type GeneratedType_Webhook_security_advisory_published struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Security_advisory Webhookssecurityadvisory `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_release_created represents the GeneratedType_Webhook_release_created schema from the OpenAPI specification
type GeneratedType_Webhook_release_created struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Dependabot_alert_security_vulnerability represents the GeneratedType_Dependabot_alert_security_vulnerability schema from the OpenAPI specification
type GeneratedType_Dependabot_alert_security_vulnerability struct {
	Vulnerable_version_range string `json:"vulnerable_version_range"` // Conditions that identify vulnerable versions of this vulnerability's package.
	First_patched_version map[string]interface{} `json:"first_patched_version"` // Details pertaining to the package version that patches this vulnerability.
	PackageField GeneratedType_Dependabot_alert_package `json:"package"` // Details for the vulnerable package.
	Severity string `json:"severity"` // The severity of the vulnerability.
}

// GeneratedType_Webhook_star_created represents the GeneratedType_Webhook_star_created schema from the OpenAPI specification
type GeneratedType_Webhook_star_created struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Starred_at string `json:"starred_at"` // The time the star was created. This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Will be `null` for the `deleted` action.
}

// GeneratedType_Webhook_package_updated represents the GeneratedType_Webhook_package_updated schema from the OpenAPI specification
type GeneratedType_Webhook_package_updated struct {
	PackageField map[string]interface{} `json:"package"` // Information about the package.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_personal_access_token_request_created represents the GeneratedType_Webhook_personal_access_token_request_created schema from the OpenAPI specification
type GeneratedType_Webhook_personal_access_token_request_created struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Personal_access_token_request GeneratedType_Personal_access_token_request `json:"personal_access_token_request"` // Details of a Personal Access Token Request.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Pull_request_merge_result represents the GeneratedType_Pull_request_merge_result schema from the OpenAPI specification
type GeneratedType_Pull_request_merge_result struct {
	Merged bool `json:"merged"`
	Message string `json:"message"`
	Sha string `json:"sha"`
}

// Installation represents the Installation schema from the OpenAPI specification
type Installation struct {
	Target_type string `json:"target_type"`
	Account interface{} `json:"account"`
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Target_id int `json:"target_id"` // The ID of the user or organization this token is being scoped to.
	Access_tokens_url string `json:"access_tokens_url"`
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Id int `json:"id"` // The ID of the installation.
	Single_file_name string `json:"single_file_name"`
	Suspended_at string `json:"suspended_at"`
	Html_url string `json:"html_url"`
	Suspended_by GeneratedType_Nullable_simple_user `json:"suspended_by"` // A GitHub user.
	Updated_at string `json:"updated_at"`
	App_slug string `json:"app_slug"`
	Contact_email string `json:"contact_email,omitempty"`
	Repositories_url string `json:"repositories_url"`
	Created_at string `json:"created_at"`
	App_id int `json:"app_id"`
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Permissions GeneratedType_App_permissions `json:"permissions"` // The permissions granted to the user access token.
	Events []string `json:"events"`
}

// GeneratedType_Webhook_discussion_unpinned represents the GeneratedType_Webhook_discussion_unpinned schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_unpinned struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// Autolink represents the Autolink schema from the OpenAPI specification
type Autolink struct {
	Id int `json:"id"`
	Is_alphanumeric bool `json:"is_alphanumeric"` // Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters.
	Key_prefix string `json:"key_prefix"` // The prefix of a key that is linkified.
	Url_template string `json:"url_template"` // A template for the target URL that is generated if a key was found.
}

// GeneratedType_Workflow_usage represents the GeneratedType_Workflow_usage schema from the OpenAPI specification
type GeneratedType_Workflow_usage struct {
	Billable map[string]interface{} `json:"billable"`
}

// GeneratedType_Branch_protection represents the GeneratedType_Branch_protection schema from the OpenAPI specification
type GeneratedType_Branch_protection struct {
	Url string `json:"url,omitempty"`
	Allow_deletions map[string]interface{} `json:"allow_deletions,omitempty"`
	Required_linear_history map[string]interface{} `json:"required_linear_history,omitempty"`
	Enforce_admins GeneratedType_Protected_branch_admin_enforced `json:"enforce_admins,omitempty"` // Protected Branch Admin Enforced
	Restrictions GeneratedType_Branch_restriction_policy `json:"restrictions,omitempty"` // Branch Restriction Policy
	Required_signatures map[string]interface{} `json:"required_signatures,omitempty"`
	Allow_fork_syncing map[string]interface{} `json:"allow_fork_syncing,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
	Block_creations map[string]interface{} `json:"block_creations,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Name string `json:"name,omitempty"`
	Required_conversation_resolution map[string]interface{} `json:"required_conversation_resolution,omitempty"`
	Required_pull_request_reviews GeneratedType_Protected_branch_pull_request_review `json:"required_pull_request_reviews,omitempty"` // Protected Branch Pull Request Review
	Required_status_checks GeneratedType_Protected_branch_required_status_check `json:"required_status_checks,omitempty"` // Protected Branch Required Status Check
	Lock_branch map[string]interface{} `json:"lock_branch,omitempty"` // Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
	Allow_force_pushes map[string]interface{} `json:"allow_force_pushes,omitempty"`
	Protection_url string `json:"protection_url,omitempty"`
}

// GeneratedType_Dependabot_alert_security_advisory represents the GeneratedType_Dependabot_alert_security_advisory schema from the OpenAPI specification
type GeneratedType_Dependabot_alert_security_advisory struct {
	Cvss map[string]interface{} `json:"cvss"` // Details for the advisory pertaining to the Common Vulnerability Scoring System.
	Updated_at string `json:"updated_at"` // The time that the advisory was last modified in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Epss GeneratedType_Security_advisory_epss `json:"epss,omitempty"` // The EPSS scores as calculated by the [Exploit Prediction Scoring System](https://www.first.org/epss).
	References []map[string]interface{} `json:"references"` // Links to additional advisory information.
	Vulnerabilities []GeneratedType_Dependabot_alert_security_vulnerability `json:"vulnerabilities"` // Vulnerable version range information for the advisory.
	Cve_id string `json:"cve_id"` // The unique CVE ID assigned to the advisory.
	Identifiers []map[string]interface{} `json:"identifiers"` // Values that identify this advisory among security information sources.
	Cvss_severities GeneratedType_Cvss_severities `json:"cvss_severities,omitempty"`
	Cwes []map[string]interface{} `json:"cwes"` // Details for the advisory pertaining to Common Weakness Enumeration.
	Description string `json:"description"` // A long-form Markdown-supported description of the advisory.
	Ghsa_id string `json:"ghsa_id"` // The unique GitHub Security Advisory ID assigned to the advisory.
	Published_at string `json:"published_at"` // The time that the advisory was published in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Severity string `json:"severity"` // The severity of the advisory.
	Summary string `json:"summary"` // A short, plain text summary of the advisory.
	Withdrawn_at string `json:"withdrawn_at"` // The time that the advisory was withdrawn in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType_Interaction_limit_response represents the GeneratedType_Interaction_limit_response schema from the OpenAPI specification
type GeneratedType_Interaction_limit_response struct {
	Limit string `json:"limit"` // The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
	Origin string `json:"origin"`
	Expires_at string `json:"expires_at"`
}

// GeneratedType_Webhook_pull_request_opened represents the GeneratedType_Webhook_pull_request_opened schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_opened struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType_Pull_request_webhook `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
}

// GeneratedType_Repository_rule_params_workflow_file_reference represents the GeneratedType_Repository_rule_params_workflow_file_reference schema from the OpenAPI specification
type GeneratedType_Repository_rule_params_workflow_file_reference struct {
	Ref string `json:"ref,omitempty"` // The ref (branch or tag) of the workflow file to use
	Repository_id int `json:"repository_id"` // The ID of the repository where the workflow is defined
	Sha string `json:"sha,omitempty"` // The commit SHA of the workflow file to use
	Path string `json:"path"` // The path to the workflow file
}

// GeneratedType_Webhook_organization_member_removed represents the GeneratedType_Webhook_organization_member_removed schema from the OpenAPI specification
type GeneratedType_Webhook_organization_member_removed struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Membership Webhooksmembership `json:"membership"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_discussion_comment_edited represents the GeneratedType_Webhook_discussion_comment_edited schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_comment_edited struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Changes map[string]interface{} `json:"changes"`
	Comment Webhookscomment `json:"comment"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_dependabot_alert_auto_reopened represents the GeneratedType_Webhook_dependabot_alert_auto_reopened schema from the OpenAPI specification
type GeneratedType_Webhook_dependabot_alert_auto_reopened struct {
	Action string `json:"action"`
	Alert GeneratedType_Dependabot_alert `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Milestoned_issue_event represents the GeneratedType_Milestoned_issue_event schema from the OpenAPI specification
type GeneratedType_Milestoned_issue_event struct {
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_id string `json:"commit_id"`
	Id int `json:"id"`
	Milestone map[string]interface{} `json:"milestone"`
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
}

// GeneratedType_Nullable_integration represents the GeneratedType_Nullable_integration schema from the OpenAPI specification
type GeneratedType_Nullable_integration struct {
	Permissions map[string]interface{} `json:"permissions"` // The set of permissions for the GitHub app
	Description string `json:"description"`
	Events []string `json:"events"` // The list of events for the GitHub app. Note that the `installation_target`, `security_advisory`, and `meta` events are not included because they are global events and not specific to an installation.
	Installations_count int `json:"installations_count,omitempty"` // The number of installations associated with the GitHub app. Only returned when the integration is requesting details about itself.
	Owner interface{} `json:"owner"`
	External_url string `json:"external_url"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // Unique identifier of the GitHub app
	Name string `json:"name"` // The name of the GitHub app
	Slug string `json:"slug,omitempty"` // The slug name of the GitHub app
	Client_id string `json:"client_id,omitempty"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
}

// GeneratedType_Webhook_discussion_pinned represents the GeneratedType_Webhook_discussion_pinned schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_pinned struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Code_of_conduct_simple represents the GeneratedType_Code_of_conduct_simple schema from the OpenAPI specification
type GeneratedType_Code_of_conduct_simple struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Key string `json:"key"`
}

// Webhooksrule represents the Webhooksrule schema from the OpenAPI specification
type Webhooksrule struct {
	Required_status_checks_enforcement_level string `json:"required_status_checks_enforcement_level"`
	Id int `json:"id"`
	Pull_request_reviews_enforcement_level string `json:"pull_request_reviews_enforcement_level"`
	Require_last_push_approval bool `json:"require_last_push_approval,omitempty"` // Whether the most recent push must be approved by someone other than the person who pushed it
	Dismiss_stale_reviews_on_push bool `json:"dismiss_stale_reviews_on_push"`
	Ignore_approvals_from_contributors bool `json:"ignore_approvals_from_contributors"`
	Authorized_actor_names []string `json:"authorized_actor_names"`
	Linear_history_requirement_enforcement_level string `json:"linear_history_requirement_enforcement_level"`
	Allow_deletions_enforcement_level string `json:"allow_deletions_enforcement_level"`
	Updated_at string `json:"updated_at"`
	Merge_queue_enforcement_level string `json:"merge_queue_enforcement_level"`
	Required_approving_review_count int `json:"required_approving_review_count"`
	Required_deployments_enforcement_level string `json:"required_deployments_enforcement_level"`
	Authorized_actors_only bool `json:"authorized_actors_only"`
	Create_protected bool `json:"create_protected,omitempty"`
	Repository_id int `json:"repository_id"`
	Name string `json:"name"`
	Allow_force_pushes_enforcement_level string `json:"allow_force_pushes_enforcement_level"`
	Admin_enforced bool `json:"admin_enforced"`
	Signature_requirement_enforcement_level string `json:"signature_requirement_enforcement_level"`
	Authorized_dismissal_actors_only bool `json:"authorized_dismissal_actors_only"`
	Created_at string `json:"created_at"`
	Lock_branch_enforcement_level string `json:"lock_branch_enforcement_level"` // The enforcement level of the branch lock setting. `off` means the branch is not locked, `non_admins` means the branch is read-only for non_admins, and `everyone` means the branch is read-only for everyone.
	Require_code_owner_review bool `json:"require_code_owner_review"`
	Required_status_checks []string `json:"required_status_checks"`
	Strict_required_status_checks_policy bool `json:"strict_required_status_checks_policy"`
	Lock_allows_fork_sync bool `json:"lock_allows_fork_sync,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow users to pull changes from upstream when the branch is locked. This setting is only applicable for forks.
	Required_conversation_resolution_level string `json:"required_conversation_resolution_level"`
}

// GeneratedType_Webhook_sponsorship_created represents the GeneratedType_Webhook_sponsorship_created schema from the OpenAPI specification
type GeneratedType_Webhook_sponsorship_created struct {
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// Authorization represents the Authorization schema from the OpenAPI specification
type Authorization struct {
	Scopes []string `json:"scopes"` // A list of scopes that this authorization is in.
	Note string `json:"note"`
	App map[string]interface{} `json:"app"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Id int64 `json:"id"`
	Fingerprint string `json:"fingerprint"`
	Token string `json:"token"`
	Expires_at string `json:"expires_at"`
	Updated_at string `json:"updated_at"`
	Installation GeneratedType_Nullable_scoped_installation `json:"installation,omitempty"`
	Note_url string `json:"note_url"`
	User GeneratedType_Nullable_simple_user `json:"user,omitempty"` // A GitHub user.
	Hashed_token string `json:"hashed_token"`
	Token_last_eight string `json:"token_last_eight"`
}

// GeneratedType_Organization_update_issue_type represents the GeneratedType_Organization_update_issue_type schema from the OpenAPI specification
type GeneratedType_Organization_update_issue_type struct {
	Name string `json:"name"` // Name of the issue type.
	Color string `json:"color,omitempty"` // Color for the issue type.
	Description string `json:"description,omitempty"` // Description of the issue type.
	Is_enabled bool `json:"is_enabled"` // Whether or not the issue type is enabled at the organization level.
}

// GeneratedType_Simple_classroom_user represents the GeneratedType_Simple_classroom_user schema from the OpenAPI specification
type GeneratedType_Simple_classroom_user struct {
	Avatar_url string `json:"avatar_url"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Login string `json:"login"`
}

// GeneratedType_Repository_rule_required_linear_history represents the GeneratedType_Repository_rule_required_linear_history schema from the OpenAPI specification
type GeneratedType_Repository_rule_required_linear_history struct {
	TypeField string `json:"type"`
}

// GeneratedType_Org_membership represents the GeneratedType_Org_membership schema from the OpenAPI specification
type GeneratedType_Org_membership struct {
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Organization GeneratedType_Organization_simple `json:"organization"` // A GitHub organization.
	Organization_url string `json:"organization_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Role string `json:"role"` // The user's membership type in the organization.
	State string `json:"state"` // The state of the member in the organization. The `pending` state indicates the user has not yet accepted an invitation.
	Url string `json:"url"`
}

// GeneratedType_Enterprise_webhooks represents the GeneratedType_Enterprise_webhooks schema from the OpenAPI specification
type GeneratedType_Enterprise_webhooks struct {
	Id int `json:"id"` // Unique identifier of the enterprise
	Name string `json:"name"` // The name of the enterprise.
	Node_id string `json:"node_id"`
	Html_url string `json:"html_url"`
	Website_url string `json:"website_url,omitempty"` // The enterprise's website URL.
	Description string `json:"description,omitempty"` // A short description of the enterprise.
	Slug string `json:"slug"` // The slug url identifier for the enterprise.
	Updated_at string `json:"updated_at"`
	Avatar_url string `json:"avatar_url"`
	Created_at string `json:"created_at"`
}

// GeneratedType_Public_user represents the GeneratedType_Public_user schema from the OpenAPI specification
type GeneratedType_Public_user struct {
	Followers_url string `json:"followers_url"`
	Avatar_url string `json:"avatar_url"`
	Hireable bool `json:"hireable"`
	Created_at string `json:"created_at"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Updated_at string `json:"updated_at"`
	Name string `json:"name"`
	Html_url string `json:"html_url"`
	Location string `json:"location"`
	Events_url string `json:"events_url"`
	Id int64 `json:"id"`
	Private_gists int `json:"private_gists,omitempty"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Gravatar_id string `json:"gravatar_id"`
	Starred_url string `json:"starred_url"`
	Followers int `json:"followers"`
	User_view_type string `json:"user_view_type,omitempty"`
	Twitter_username string `json:"twitter_username,omitempty"`
	TypeField string `json:"type"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Collaborators int `json:"collaborators,omitempty"`
	Bio string `json:"bio"`
	Company string `json:"company"`
	Site_admin bool `json:"site_admin"`
	Node_id string `json:"node_id"`
	Email string `json:"email"`
	Repos_url string `json:"repos_url"`
	Following int `json:"following"`
	Received_events_url string `json:"received_events_url"`
	Following_url string `json:"following_url"`
	Organizations_url string `json:"organizations_url"`
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Blog string `json:"blog"`
	Public_gists int `json:"public_gists"`
	Gists_url string `json:"gists_url"`
	Login string `json:"login"`
	Subscriptions_url string `json:"subscriptions_url"`
	Public_repos int `json:"public_repos"`
	Notification_email string `json:"notification_email,omitempty"`
	Url string `json:"url"`
}

// GeneratedType_Deployment_branch_policy represents the GeneratedType_Deployment_branch_policy schema from the OpenAPI specification
type GeneratedType_Deployment_branch_policy struct {
	Id int `json:"id,omitempty"` // The unique identifier of the branch or tag policy.
	Name string `json:"name,omitempty"` // The name pattern that branches or tags must match in order to deploy to the environment.
	Node_id string `json:"node_id,omitempty"`
	TypeField string `json:"type,omitempty"` // Whether this rule targets a branch or tag.
}

// GeneratedType_Pages_https_certificate represents the GeneratedType_Pages_https_certificate schema from the OpenAPI specification
type GeneratedType_Pages_https_certificate struct {
	Description string `json:"description"`
	Domains []string `json:"domains"` // Array of the domain set and its alternate name (if it is configured)
	Expires_at string `json:"expires_at,omitempty"`
	State string `json:"state"`
}

// GeneratedType_Code_scanning_autofix_commits_response represents the GeneratedType_Code_scanning_autofix_commits_response schema from the OpenAPI specification
type GeneratedType_Code_scanning_autofix_commits_response struct {
	Sha string `json:"sha,omitempty"` // SHA of commit with autofix.
	Target_ref string `json:"target_ref,omitempty"` // The Git reference of target branch for the commit. For more information, see "[Git References](https://git-scm.com/book/en/v2/Git-Internals-Git-References)" in the Git documentation.
}

// GeneratedType_Removed_from_project_issue_event represents the GeneratedType_Removed_from_project_issue_event schema from the OpenAPI specification
type GeneratedType_Removed_from_project_issue_event struct {
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Event string `json:"event"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_url string `json:"commit_url"`
}

// GeneratedType_Issue_comment represents the GeneratedType_Issue_comment schema from the OpenAPI specification
type GeneratedType_Issue_comment struct {
	Body_text string `json:"body_text,omitempty"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"` // URL for the issue comment
	Created_at string `json:"created_at"`
	Id int64 `json:"id"` // Unique identifier of the issue comment
	Issue_url string `json:"issue_url"`
	Node_id string `json:"node_id"`
	Body_html string `json:"body_html,omitempty"`
	Html_url string `json:"html_url"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body,omitempty"` // Contents of the issue comment
}

// GeneratedType_Issue_search_result_item represents the GeneratedType_Issue_search_result_item schema from the OpenAPI specification
type GeneratedType_Issue_search_result_item struct {
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	State string `json:"state"`
	Body_html string `json:"body_html,omitempty"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Events_url string `json:"events_url"`
	Id int64 `json:"id"`
	Draft bool `json:"draft,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	Closed_at string `json:"closed_at"`
	Node_id string `json:"node_id"`
	Number int `json:"number"`
	Body string `json:"body,omitempty"`
	Url string `json:"url"`
	State_reason string `json:"state_reason,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Score float64 `json:"score"`
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Comments_url string `json:"comments_url"`
	Updated_at string `json:"updated_at"`
	Assignees []GeneratedType_Simple_user `json:"assignees,omitempty"`
	Sub_issues_summary map[string]interface{} `json:"sub_issues_summary,omitempty"`
	Locked bool `json:"locked"`
	TypeField GeneratedType_Issue_type `json:"type,omitempty"` // The type of issue.
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Labels []map[string]interface{} `json:"labels"`
	Labels_url string `json:"labels_url"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Comments int `json:"comments"`
	Assignee GeneratedType_Nullable_simple_user `json:"assignee"` // A GitHub user.
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Milestone GeneratedType_Nullable_milestone `json:"milestone"` // A collection of related issues and pull requests.
	Title string `json:"title"`
	Repository_url string `json:"repository_url"`
}

// GeneratedType_Webhook_code_scanning_alert_closed_by_user represents the GeneratedType_Webhook_code_scanning_alert_closed_by_user schema from the OpenAPI specification
type GeneratedType_Webhook_code_scanning_alert_closed_by_user struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Hook_delivery_item represents the GeneratedType_Hook_delivery_item schema from the OpenAPI specification
type GeneratedType_Hook_delivery_item struct {
	Id int64 `json:"id"` // Unique identifier of the webhook delivery.
	Installation_id int64 `json:"installation_id"` // The id of the GitHub App installation associated with this event.
	Status_code int `json:"status_code"` // Status code received when delivery was made.
	Delivered_at string `json:"delivered_at"` // Time when the webhook delivery occurred.
	Guid string `json:"guid"` // Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
	Redelivery bool `json:"redelivery"` // Whether the webhook delivery is a redelivery.
	Repository_id int64 `json:"repository_id"` // The id of the repository associated with this event.
	Event string `json:"event"` // The event that triggered the delivery.
	Action string `json:"action"` // The type of activity for the event that triggered the delivery.
	Status string `json:"status"` // Describes the response returned after attempting the delivery.
	Throttled_at string `json:"throttled_at,omitempty"` // Time when the webhook delivery was throttled.
	Duration float64 `json:"duration"` // Time spent delivering.
}

// GeneratedType_Actions_secret represents the GeneratedType_Actions_secret schema from the OpenAPI specification
type GeneratedType_Actions_secret struct {
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
}

// GeneratedType_Webhook_marketplace_purchase_pending_change_cancelled represents the GeneratedType_Webhook_marketplace_purchase_pending_change_cancelled schema from the OpenAPI specification
type GeneratedType_Webhook_marketplace_purchase_pending_change_cancelled struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Previous_marketplace_purchase Webhookspreviousmarketplacepurchase `json:"previous_marketplace_purchase,omitempty"`
	Marketplace_purchase map[string]interface{} `json:"marketplace_purchase"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Effective_date string `json:"effective_date"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Simple_user represents the GeneratedType_Simple_user schema from the OpenAPI specification
type GeneratedType_Simple_user struct {
	Repos_url string `json:"repos_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Id int64 `json:"id"`
	Login string `json:"login"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Events_url string `json:"events_url"`
	Following_url string `json:"following_url"`
	Html_url string `json:"html_url"`
	Subscriptions_url string `json:"subscriptions_url"`
	Name string `json:"name,omitempty"`
	Starred_url string `json:"starred_url"`
	User_view_type string `json:"user_view_type,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Gists_url string `json:"gists_url"`
	Email string `json:"email,omitempty"`
	Site_admin bool `json:"site_admin"`
	Gravatar_id string `json:"gravatar_id"`
	TypeField string `json:"type"`
	Organizations_url string `json:"organizations_url"`
	Followers_url string `json:"followers_url"`
	Received_events_url string `json:"received_events_url"`
}

// GeneratedType_Webhook_code_scanning_alert_fixed represents the GeneratedType_Webhook_code_scanning_alert_fixed schema from the OpenAPI specification
type GeneratedType_Webhook_code_scanning_alert_fixed struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Repository_rule_pull_request represents the GeneratedType_Repository_rule_pull_request schema from the OpenAPI specification
type GeneratedType_Repository_rule_pull_request struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType_Repository_rule_deletion represents the GeneratedType_Repository_rule_deletion schema from the OpenAPI specification
type GeneratedType_Repository_rule_deletion struct {
	TypeField string `json:"type"`
}

// GeneratedType_Repository_rule_commit_author_email_pattern represents the GeneratedType_Repository_rule_commit_author_email_pattern schema from the OpenAPI specification
type GeneratedType_Repository_rule_commit_author_email_pattern struct {
	TypeField string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType_Protected_branch_required_status_check represents the GeneratedType_Protected_branch_required_status_check schema from the OpenAPI specification
type GeneratedType_Protected_branch_required_status_check struct {
	Contexts []string `json:"contexts"`
	Contexts_url string `json:"contexts_url,omitempty"`
	Enforcement_level string `json:"enforcement_level,omitempty"`
	Strict bool `json:"strict,omitempty"`
	Url string `json:"url,omitempty"`
	Checks []map[string]interface{} `json:"checks"`
}

// Webhookssecurityadvisory represents the Webhookssecurityadvisory schema from the OpenAPI specification
type Webhookssecurityadvisory struct {
	References []map[string]interface{} `json:"references"`
	Summary string `json:"summary"`
	Description string `json:"description"`
	Ghsa_id string `json:"ghsa_id"`
	Withdrawn_at string `json:"withdrawn_at"`
	Identifiers []map[string]interface{} `json:"identifiers"`
	Severity string `json:"severity"`
	Updated_at string `json:"updated_at"`
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities"`
	Cvss map[string]interface{} `json:"cvss"`
	Cvss_severities GeneratedType_Cvss_severities `json:"cvss_severities,omitempty"`
	Cwes []map[string]interface{} `json:"cwes"`
	Published_at string `json:"published_at"`
}

// GeneratedType_Repository_rule_creation represents the GeneratedType_Repository_rule_creation schema from the OpenAPI specification
type GeneratedType_Repository_rule_creation struct {
	TypeField string `json:"type"`
}

// Key represents the Key schema from the OpenAPI specification
type Key struct {
	Read_only bool `json:"read_only"`
	Title string `json:"title"`
	Url string `json:"url"`
	Verified bool `json:"verified"`
	Created_at string `json:"created_at"`
	Id int64 `json:"id"`
	Key string `json:"key"`
}

// GeneratedType_Webhook_project_edited represents the GeneratedType_Webhook_project_edited schema from the OpenAPI specification
type GeneratedType_Webhook_project_edited struct {
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // The changes to the project if the action was `edited`.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Codespaces_user_public_key represents the GeneratedType_Codespaces_user_public_key schema from the OpenAPI specification
type GeneratedType_Codespaces_user_public_key struct {
	Key_id string `json:"key_id"` // The identifier for the key.
	Key string `json:"key"` // The Base64 encoded public key.
}

// GeneratedType_Custom_property_set_payload represents the GeneratedType_Custom_property_set_payload schema from the OpenAPI specification
type GeneratedType_Custom_property_set_payload struct {
	Values_editable_by string `json:"values_editable_by,omitempty"` // Who can edit the values of the property
	Allowed_values []string `json:"allowed_values,omitempty"` // An ordered list of the allowed values of the property. The property can have up to 200 allowed values.
	Default_value string `json:"default_value,omitempty"` // Default value of the property
	Description string `json:"description,omitempty"` // Short description of the property
	Required bool `json:"required,omitempty"` // Whether the property is required.
	Value_type string `json:"value_type"` // The type of the value for the property
}

// GeneratedType_Review_comment represents the GeneratedType_Review_comment schema from the OpenAPI specification
type GeneratedType_Review_comment struct {
	Original_commit_id string `json:"original_commit_id"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Body_text string `json:"body_text,omitempty"`
	Body_html string `json:"body_html,omitempty"`
	Diff_hunk string `json:"diff_hunk"`
	Original_position int `json:"original_position"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Links map[string]interface{} `json:"_links"`
	Subject_type string `json:"subject_type,omitempty"` // The level at which the comment is targeted, can be a diff line or a file.
	Original_line int `json:"original_line,omitempty"` // The original line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Original_start_line int `json:"original_start_line,omitempty"` // The original first line of the range for a multi-line comment.
	Node_id string `json:"node_id"`
	Start_side string `json:"start_side,omitempty"` // The side of the first line of the range for a multi-line comment.
	In_reply_to_id int `json:"in_reply_to_id,omitempty"`
	Url string `json:"url"`
	Line int `json:"line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Commit_id string `json:"commit_id"`
	Side string `json:"side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Body string `json:"body"`
	Pull_request_review_id int64 `json:"pull_request_review_id"`
	Html_url string `json:"html_url"`
	Pull_request_url string `json:"pull_request_url"`
	Id int64 `json:"id"`
	Start_line int `json:"start_line,omitempty"` // The first line of the range for a multi-line comment.
	Path string `json:"path"`
	Position int `json:"position"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
}

// GeneratedType_Webhook_project_column_moved represents the GeneratedType_Webhook_project_column_moved schema from the OpenAPI specification
type GeneratedType_Webhook_project_column_moved struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_column Webhooksprojectcolumn `json:"project_column"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Organization_simple_webhooks represents the GeneratedType_Organization_simple_webhooks schema from the OpenAPI specification
type GeneratedType_Organization_simple_webhooks struct {
	Login string `json:"login"`
	Members_url string `json:"members_url"`
	Node_id string `json:"node_id"`
	Repos_url string `json:"repos_url"`
	Description string `json:"description"`
	Public_members_url string `json:"public_members_url"`
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Issues_url string `json:"issues_url"`
	Events_url string `json:"events_url"`
	Hooks_url string `json:"hooks_url"`
	Id int `json:"id"`
}

// GeneratedType_Nullable_simple_commit represents the GeneratedType_Nullable_simple_commit schema from the OpenAPI specification
type GeneratedType_Nullable_simple_commit struct {
	Committer map[string]interface{} `json:"committer"` // Information about the Git committer
	Id string `json:"id"` // SHA for the commit
	Message string `json:"message"` // Message describing the purpose of the commit
	Timestamp string `json:"timestamp"` // Timestamp of the commit
	Tree_id string `json:"tree_id"` // SHA for the commit's tree
	Author map[string]interface{} `json:"author"` // Information about the Git author
}

// GeneratedType_Webhook_check_run_created_form_encoded represents the GeneratedType_Webhook_check_run_created_form_encoded schema from the OpenAPI specification
type GeneratedType_Webhook_check_run_created_form_encoded struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.created JSON payload. The decoded payload is a JSON object.
}

// GeneratedType_Private_vulnerability_report_create represents the GeneratedType_Private_vulnerability_report_create schema from the OpenAPI specification
type GeneratedType_Private_vulnerability_report_create struct {
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities,omitempty"` // An array of products affected by the vulnerability detailed in a repository security advisory.
	Cvss_vector_string string `json:"cvss_vector_string,omitempty"` // The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
	Cwe_ids []string `json:"cwe_ids,omitempty"` // A list of Common Weakness Enumeration (CWE) IDs.
	Description string `json:"description"` // A detailed description of what the advisory impacts.
	Severity string `json:"severity,omitempty"` // The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
	Start_private_fork bool `json:"start_private_fork,omitempty"` // Whether to create a temporary private fork of the repository to collaborate on a fix.
	Summary string `json:"summary"` // A short summary of the advisory.
}

// GeneratedType_Nullable_code_of_conduct_simple represents the GeneratedType_Nullable_code_of_conduct_simple schema from the OpenAPI specification
type GeneratedType_Nullable_code_of_conduct_simple struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Key string `json:"key"`
}

// GeneratedType_License_simple represents the GeneratedType_License_simple schema from the OpenAPI specification
type GeneratedType_License_simple struct {
	Url string `json:"url"`
	Html_url string `json:"html_url,omitempty"`
	Key string `json:"key"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Spdx_id string `json:"spdx_id"`
}

// GeneratedType_Webhook_discussion_created represents the GeneratedType_Webhook_discussion_created schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_created struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_label_deleted represents the GeneratedType_Webhook_label_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_label_deleted struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Nullable_repository represents the GeneratedType_Nullable_repository schema from the OpenAPI specification
type GeneratedType_Nullable_repository struct {
	Collaborators_url string `json:"collaborators_url"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Forks_count int `json:"forks_count"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Issues_url string `json:"issues_url"`
	Languages_url string `json:"languages_url"`
	Pulls_url string `json:"pulls_url"`
	Svn_url string `json:"svn_url"`
	Id int64 `json:"id"` // Unique identifier of the repository
	Topics []string `json:"topics,omitempty"`
	Subscribers_url string `json:"subscribers_url"`
	Issue_events_url string `json:"issue_events_url"`
	Notifications_url string `json:"notifications_url"`
	Open_issues int `json:"open_issues"`
	Created_at string `json:"created_at"`
	Milestones_url string `json:"milestones_url"`
	Watchers int `json:"watchers"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Git_url string `json:"git_url"`
	Merges_url string `json:"merges_url"`
	Commits_url string `json:"commits_url"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Pushed_at string `json:"pushed_at"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Contents_url string `json:"contents_url"`
	Mirror_url string `json:"mirror_url"`
	Node_id string `json:"node_id"`
	Stargazers_count int `json:"stargazers_count"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Downloads_url string `json:"downloads_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Git_refs_url string `json:"git_refs_url"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Statuses_url string `json:"statuses_url"`
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Releases_url string `json:"releases_url"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Git_commits_url string `json:"git_commits_url"`
	Tags_url string `json:"tags_url"`
	Has_pages bool `json:"has_pages"`
	Contributors_url string `json:"contributors_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Master_branch string `json:"master_branch,omitempty"`
	Updated_at string `json:"updated_at"`
	Ssh_url string `json:"ssh_url"`
	Compare_url string `json:"compare_url"`
	Comments_url string `json:"comments_url"`
	Deployments_url string `json:"deployments_url"`
	Keys_url string `json:"keys_url"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Open_issues_count int `json:"open_issues_count"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Assignees_url string `json:"assignees_url"`
	Code_search_index_status map[string]interface{} `json:"code_search_index_status,omitempty"` // The status of the code search index for this repository
	Blobs_url string `json:"blobs_url"`
	Trees_url string `json:"trees_url"`
	License GeneratedType_Nullable_license_simple `json:"license"` // License Simple
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Fork bool `json:"fork"`
	Url string `json:"url"`
	Forks int `json:"forks"`
	Watchers_count int `json:"watchers_count"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Description string `json:"description"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Issue_comment_url string `json:"issue_comment_url"`
	Full_name string `json:"full_name"`
	Stargazers_url string `json:"stargazers_url"`
	Git_tags_url string `json:"git_tags_url"`
	Hooks_url string `json:"hooks_url"`
	Teams_url string `json:"teams_url"`
	Events_url string `json:"events_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Name string `json:"name"` // The name of the repository.
	Labels_url string `json:"labels_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Clone_url string `json:"clone_url"`
	Homepage string `json:"homepage"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Branches_url string `json:"branches_url"`
	Archive_url string `json:"archive_url"`
	Forks_url string `json:"forks_url"`
	Language string `json:"language"`
	Subscription_url string `json:"subscription_url"`
	Html_url string `json:"html_url"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
}

// GeneratedType_Webhook_sponsorship_tier_changed represents the GeneratedType_Webhook_sponsorship_tier_changed schema from the OpenAPI specification
type GeneratedType_Webhook_sponsorship_tier_changed struct {
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
	Changes Webhookschanges8 `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// Blob represents the Blob schema from the OpenAPI specification
type Blob struct {
	Size int `json:"size"`
	Url string `json:"url"`
	Content string `json:"content"`
	Encoding string `json:"encoding"`
	Highlighted_content string `json:"highlighted_content,omitempty"`
	Node_id string `json:"node_id"`
	Sha string `json:"sha"`
}

// GeneratedType_Webhook_projects_v2_project_edited represents the GeneratedType_Webhook_projects_v2_project_edited schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_project_edited struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2 GeneratedType_Projects_v2 `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Actions_billing_usage represents the GeneratedType_Actions_billing_usage schema from the OpenAPI specification
type GeneratedType_Actions_billing_usage struct {
	Included_minutes int `json:"included_minutes"` // The amount of free GitHub Actions minutes available.
	Minutes_used_breakdown map[string]interface{} `json:"minutes_used_breakdown"`
	Total_minutes_used int `json:"total_minutes_used"` // The sum of the free and paid GitHub Actions minutes used.
	Total_paid_minutes_used int `json:"total_paid_minutes_used"` // The total paid GitHub Actions minutes used.
}

// GeneratedType_Webhook_personal_access_token_request_approved represents the GeneratedType_Webhook_personal_access_token_request_approved schema from the OpenAPI specification
type GeneratedType_Webhook_personal_access_token_request_approved struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Personal_access_token_request GeneratedType_Personal_access_token_request `json:"personal_access_token_request"` // Details of a Personal Access Token Request.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Runner_label represents the GeneratedType_Runner_label schema from the OpenAPI specification
type GeneratedType_Runner_label struct {
	Id int `json:"id,omitempty"` // Unique identifier of the label.
	Name string `json:"name"` // Name of the label.
	TypeField string `json:"type,omitempty"` // The type of label. Read-only labels are applied automatically when the runner is configured.
}

// GeneratedType_Team_membership represents the GeneratedType_Team_membership schema from the OpenAPI specification
type GeneratedType_Team_membership struct {
	Url string `json:"url"`
	Role string `json:"role"` // The role of the user in the team.
	State string `json:"state"` // The state of the user's membership in the team.
}

// GeneratedType_Code_scanning_variant_analysis_repository represents the GeneratedType_Code_scanning_variant_analysis_repository schema from the OpenAPI specification
type GeneratedType_Code_scanning_variant_analysis_repository struct {
	Id int `json:"id"` // A unique identifier of the repository.
	Name string `json:"name"` // The name of the repository.
	Private bool `json:"private"` // Whether the repository is private.
	Stargazers_count int `json:"stargazers_count"`
	Updated_at string `json:"updated_at"`
	Full_name string `json:"full_name"` // The full, globally unique, name of the repository.
}

// Hovercard represents the Hovercard schema from the OpenAPI specification
type Hovercard struct {
	Contexts []map[string]interface{} `json:"contexts"`
}

// GeneratedType_Webhook_milestone_created represents the GeneratedType_Webhook_milestone_created schema from the OpenAPI specification
type GeneratedType_Webhook_milestone_created struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone3 `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_ping_form_encoded represents the GeneratedType_Webhook_ping_form_encoded schema from the OpenAPI specification
type GeneratedType_Webhook_ping_form_encoded struct {
	Payload string `json:"payload"` // A URL-encoded string of the ping JSON payload. The decoded payload is a JSON object.
}

// GeneratedType_Webhook_workflow_job_queued represents the GeneratedType_Webhook_workflow_job_queued schema from the OpenAPI specification
type GeneratedType_Webhook_workflow_job_queued struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Workflow_job map[string]interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_code_scanning_alert_reopened represents the GeneratedType_Webhook_code_scanning_alert_reopened schema from the OpenAPI specification
type GeneratedType_Webhook_code_scanning_alert_reopened struct {
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Action string `json:"action"`
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
}

// GeneratedType_Actions_set_default_workflow_permissions represents the GeneratedType_Actions_set_default_workflow_permissions schema from the OpenAPI specification
type GeneratedType_Actions_set_default_workflow_permissions struct {
	Can_approve_pull_request_reviews bool `json:"can_approve_pull_request_reviews,omitempty"` // Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
	Default_workflow_permissions string `json:"default_workflow_permissions,omitempty"` // The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
}

// GeneratedType_Release_asset represents the GeneratedType_Release_asset schema from the OpenAPI specification
type GeneratedType_Release_asset struct {
	Download_count int `json:"download_count"`
	Id int `json:"id"`
	State string `json:"state"` // State of the release asset.
	Updated_at string `json:"updated_at"`
	Label string `json:"label"`
	Node_id string `json:"node_id"`
	Digest string `json:"digest"`
	Name string `json:"name"` // The file name of the asset.
	Size int `json:"size"`
	Uploader GeneratedType_Nullable_simple_user `json:"uploader"` // A GitHub user.
	Url string `json:"url"`
	Browser_download_url string `json:"browser_download_url"`
	Content_type string `json:"content_type"`
	Created_at string `json:"created_at"`
}

// GeneratedType_Repository_rule_committer_email_pattern represents the GeneratedType_Repository_rule_committer_email_pattern schema from the OpenAPI specification
type GeneratedType_Repository_rule_committer_email_pattern struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType_Webhook_discussion_comment_created represents the GeneratedType_Webhook_discussion_comment_created schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_comment_created struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhookscomment `json:"comment"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_repository_advisory_published represents the GeneratedType_Webhook_repository_advisory_published schema from the OpenAPI specification
type GeneratedType_Webhook_repository_advisory_published struct {
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_advisory GeneratedType_Repository_advisory `json:"repository_advisory"` // A repository security advisory.
}

// Webhookschanges8 represents the Webhookschanges8 schema from the OpenAPI specification
type Webhookschanges8 struct {
	Tier map[string]interface{} `json:"tier"`
}

// GeneratedType_Ssh_signing_key represents the GeneratedType_Ssh_signing_key schema from the OpenAPI specification
type GeneratedType_Ssh_signing_key struct {
	Id int `json:"id"`
	Key string `json:"key"`
	Title string `json:"title"`
	Created_at string `json:"created_at"`
}

// Collaborator represents the Collaborator schema from the OpenAPI specification
type Collaborator struct {
	Html_url string `json:"html_url"`
	Following_url string `json:"following_url"`
	Gravatar_id string `json:"gravatar_id"`
	Repos_url string `json:"repos_url"`
	TypeField string `json:"type"`
	Id int64 `json:"id"`
	Url string `json:"url"`
	User_view_type string `json:"user_view_type,omitempty"`
	Events_url string `json:"events_url"`
	Organizations_url string `json:"organizations_url"`
	Email string `json:"email,omitempty"`
	Starred_url string `json:"starred_url"`
	Received_events_url string `json:"received_events_url"`
	Subscriptions_url string `json:"subscriptions_url"`
	Login string `json:"login"`
	Site_admin bool `json:"site_admin"`
	Gists_url string `json:"gists_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Role_name string `json:"role_name"`
	Name string `json:"name,omitempty"`
	Followers_url string `json:"followers_url"`
	Node_id string `json:"node_id"`
	Avatar_url string `json:"avatar_url"`
}

// Traffic represents the Traffic schema from the OpenAPI specification
type Traffic struct {
	Count int `json:"count"`
	Timestamp string `json:"timestamp"`
	Uniques int `json:"uniques"`
}

// GeneratedType_Webhook_repository_created represents the GeneratedType_Webhook_repository_created schema from the OpenAPI specification
type GeneratedType_Webhook_repository_created struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_repository_unarchived represents the GeneratedType_Webhook_repository_unarchived schema from the OpenAPI specification
type GeneratedType_Webhook_repository_unarchived struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Repository_rule_non_fast_forward represents the GeneratedType_Repository_rule_non_fast_forward schema from the OpenAPI specification
type GeneratedType_Repository_rule_non_fast_forward struct {
	TypeField string `json:"type"`
}

// GeneratedType_Repository_rule_params_status_check_configuration represents the GeneratedType_Repository_rule_params_status_check_configuration schema from the OpenAPI specification
type GeneratedType_Repository_rule_params_status_check_configuration struct {
	Context string `json:"context"` // The status check context name that must be present on the commit.
	Integration_id int `json:"integration_id,omitempty"` // The optional integration ID that this status check must originate from.
}

// GeneratedType_Repository_rule_params_reviewer represents the GeneratedType_Repository_rule_params_reviewer schema from the OpenAPI specification
type GeneratedType_Repository_rule_params_reviewer struct {
	TypeField string `json:"type"` // The type of the reviewer
	Id int `json:"id"` // ID of the reviewer which must review changes to matching files.
}

// GeneratedType_Webhook_issues_demilestoned represents the GeneratedType_Webhook_issues_demilestoned schema from the OpenAPI specification
type GeneratedType_Webhook_issues_demilestoned struct {
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Milestone Webhooksmilestone `json:"milestone,omitempty"` // A collection of related issues and pull requests.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Simple_commit_status represents the GeneratedType_Simple_commit_status schema from the OpenAPI specification
type GeneratedType_Simple_commit_status struct {
	Required bool `json:"required,omitempty"`
	Target_url string `json:"target_url"`
	Avatar_url string `json:"avatar_url"`
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	State string `json:"state"`
	Url string `json:"url"`
	Context string `json:"context"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType_Actions_organization_permissions represents the GeneratedType_Actions_organization_permissions schema from the OpenAPI specification
type GeneratedType_Actions_organization_permissions struct {
	Enabled_repositories string `json:"enabled_repositories"` // The policy that controls the repositories in the organization that are allowed to run GitHub Actions.
	Selected_actions_url string `json:"selected_actions_url,omitempty"` // The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // The API URL to use to get or set the selected repositories that are allowed to run GitHub Actions, when `enabled_repositories` is set to `selected`.
	Allowed_actions string `json:"allowed_actions,omitempty"` // The permissions policy that controls the actions and reusable workflows that are allowed to run.
}

// GeneratedType_Gitignore_template represents the GeneratedType_Gitignore_template schema from the OpenAPI specification
type GeneratedType_Gitignore_template struct {
	Source string `json:"source"`
	Name string `json:"name"`
}

// GeneratedType_Webhook_project_reopened represents the GeneratedType_Webhook_project_reopened schema from the OpenAPI specification
type GeneratedType_Webhook_project_reopened struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Rate_limit_overview represents the GeneratedType_Rate_limit_overview schema from the OpenAPI specification
type GeneratedType_Rate_limit_overview struct {
	Resources map[string]interface{} `json:"resources"`
	Rate GeneratedType_Rate_limit `json:"rate"`
}

// GeneratedType_Status_check_policy represents the GeneratedType_Status_check_policy schema from the OpenAPI specification
type GeneratedType_Status_check_policy struct {
	Url string `json:"url"`
	Checks []map[string]interface{} `json:"checks"`
	Contexts []string `json:"contexts"`
	Contexts_url string `json:"contexts_url"`
	Strict bool `json:"strict"`
}

// GeneratedType_Webhook_watch_started represents the GeneratedType_Webhook_watch_started schema from the OpenAPI specification
type GeneratedType_Webhook_watch_started struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Repository_rule_branch_name_pattern represents the GeneratedType_Repository_rule_branch_name_pattern schema from the OpenAPI specification
type GeneratedType_Repository_rule_branch_name_pattern struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType_Webhook_projects_v2_item_converted represents the GeneratedType_Webhook_projects_v2_item_converted schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_item_converted struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType_Projects_v2_item `json:"projects_v2_item"` // An item belonging to a project
}

// GeneratedType_Full_repository represents the GeneratedType_Full_repository schema from the OpenAPI specification
type GeneratedType_Full_repository struct {
	Commits_url string `json:"commits_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Tags_url string `json:"tags_url"`
	Updated_at string `json:"updated_at"`
	Fork bool `json:"fork"`
	Forks_count int `json:"forks_count"`
	Parent Repository `json:"parent,omitempty"` // A repository on GitHub.
	Has_discussions bool `json:"has_discussions"`
	Svn_url string `json:"svn_url"`
	Name string `json:"name"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"`
	Mirror_url string `json:"mirror_url"`
	Watchers_count int `json:"watchers_count"`
	Watchers int `json:"watchers"`
	Trees_url string `json:"trees_url"`
	Git_commits_url string `json:"git_commits_url"`
	Open_issues_count int `json:"open_issues_count"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Created_at string `json:"created_at"`
	Subscription_url string `json:"subscription_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Subscribers_count int `json:"subscribers_count"`
	Organization GeneratedType_Nullable_simple_user `json:"organization,omitempty"` // A GitHub user.
	Has_projects bool `json:"has_projects"`
	Language string `json:"language"`
	Has_downloads bool `json:"has_downloads,omitempty"`
	Stargazers_count int `json:"stargazers_count"`
	Teams_url string `json:"teams_url"`
	Default_branch string `json:"default_branch"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Full_name string `json:"full_name"`
	Assignees_url string `json:"assignees_url"`
	Branches_url string `json:"branches_url"`
	Comments_url string `json:"comments_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Clone_url string `json:"clone_url"`
	Source Repository `json:"source,omitempty"` // A repository on GitHub.
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Forks int `json:"forks"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Releases_url string `json:"releases_url"`
	Git_tags_url string `json:"git_tags_url"`
	Node_id string `json:"node_id"`
	Html_url string `json:"html_url"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Events_url string `json:"events_url"`
	Description string `json:"description"`
	Forks_url string `json:"forks_url"`
	Topics []string `json:"topics,omitempty"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"`
	Deployments_url string `json:"deployments_url"`
	Archive_url string `json:"archive_url"`
	Labels_url string `json:"labels_url"`
	Languages_url string `json:"languages_url"`
	Downloads_url string `json:"downloads_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Ssh_url string `json:"ssh_url"`
	Git_refs_url string `json:"git_refs_url"`
	Network_count int `json:"network_count"`
	Id int64 `json:"id"`
	Homepage string `json:"homepage"`
	Template_repository GeneratedType_Nullable_repository `json:"template_repository,omitempty"` // A repository on GitHub.
	Issues_url string `json:"issues_url"`
	Code_of_conduct GeneratedType_Code_of_conduct_simple `json:"code_of_conduct,omitempty"` // Code of Conduct Simple
	Statuses_url string `json:"statuses_url"`
	Compare_url string `json:"compare_url"`
	Git_url string `json:"git_url"`
	Notifications_url string `json:"notifications_url"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Open_issues int `json:"open_issues"`
	Collaborators_url string `json:"collaborators_url"`
	Blobs_url string `json:"blobs_url"`
	Security_and_analysis GeneratedType_Security_and_analysis `json:"security_and_analysis,omitempty"`
	Archived bool `json:"archived"`
	Issue_events_url string `json:"issue_events_url"`
	Has_wiki bool `json:"has_wiki"`
	Is_template bool `json:"is_template,omitempty"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"`
	Contributors_url string `json:"contributors_url"`
	Private bool `json:"private"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Pushed_at string `json:"pushed_at"`
	Pulls_url string `json:"pulls_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Milestones_url string `json:"milestones_url"`
	Hooks_url string `json:"hooks_url"`
	Has_issues bool `json:"has_issues"`
	License GeneratedType_Nullable_license_simple `json:"license"` // License Simple
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Contents_url string `json:"contents_url"`
	Subscribers_url string `json:"subscribers_url"`
	Has_pages bool `json:"has_pages"`
	Keys_url string `json:"keys_url"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"`
	Stargazers_url string `json:"stargazers_url"`
	Url string `json:"url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is allowed.
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Merges_url string `json:"merges_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"`
}

// GeneratedType_Repository_rule_required_status_checks represents the GeneratedType_Repository_rule_required_status_checks schema from the OpenAPI specification
type GeneratedType_Repository_rule_required_status_checks struct {
	TypeField string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType_Webhook_code_scanning_alert_appeared_in_branch represents the GeneratedType_Webhook_code_scanning_alert_appeared_in_branch schema from the OpenAPI specification
type GeneratedType_Webhook_code_scanning_alert_appeared_in_branch struct {
	Action string `json:"action"`
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_github_app_authorization_revoked represents the GeneratedType_Webhook_github_app_authorization_revoked schema from the OpenAPI specification
type GeneratedType_Webhook_github_app_authorization_revoked struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Community_profile represents the GeneratedType_Community_profile schema from the OpenAPI specification
type GeneratedType_Community_profile struct {
	Health_percentage int `json:"health_percentage"`
	Updated_at string `json:"updated_at"`
	Content_reports_enabled bool `json:"content_reports_enabled,omitempty"`
	Description string `json:"description"`
	Documentation string `json:"documentation"`
	Files map[string]interface{} `json:"files"`
}

// GeneratedType_Participation_stats represents the GeneratedType_Participation_stats schema from the OpenAPI specification
type GeneratedType_Participation_stats struct {
	All []int `json:"all"`
	Owner []int `json:"owner"`
}

// GeneratedType_Webhook_discussion_comment_deleted represents the GeneratedType_Webhook_discussion_comment_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_comment_deleted struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhookscomment `json:"comment"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_secret_scanning_alert_created represents the GeneratedType_Webhook_secret_scanning_alert_created schema from the OpenAPI specification
type GeneratedType_Webhook_secret_scanning_alert_created struct {
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType_Secret_scanning_alert_webhook `json:"alert"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_project_card_deleted represents the GeneratedType_Webhook_project_card_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_project_card_deleted struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card map[string]interface{} `json:"project_card"`
	Repository GeneratedType_Nullable_repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Webhook_pull_request_assigned represents the GeneratedType_Webhook_pull_request_assigned schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_assigned struct {
	Pull_request map[string]interface{} `json:"pull_request"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Assignee Webhooksuser `json:"assignee"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Number int `json:"number"` // The pull request number.
}

// GeneratedType_Projects_v2_single_select_option represents the GeneratedType_Projects_v2_single_select_option schema from the OpenAPI specification
type GeneratedType_Projects_v2_single_select_option struct {
	Color string `json:"color,omitempty"`
	Description string `json:"description,omitempty"`
	Id string `json:"id"`
	Name string `json:"name"`
}

// GeneratedType_Code_scanning_codeql_database represents the GeneratedType_Code_scanning_codeql_database schema from the OpenAPI specification
type GeneratedType_Code_scanning_codeql_database struct {
	Name string `json:"name"` // The name of the CodeQL database.
	Url string `json:"url"` // The URL at which to download the CodeQL database. The `Accept` header must be set to the value of the `content_type` property.
	Commit_oid string `json:"commit_oid,omitempty"` // The commit SHA of the repository at the time the CodeQL database was created.
	Language string `json:"language"` // The language of the CodeQL database.
	Content_type string `json:"content_type"` // The MIME type of the CodeQL database file.
	Created_at string `json:"created_at"` // The date and time at which the CodeQL database was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Id int `json:"id"` // The ID of the CodeQL database.
	Size int `json:"size"` // The size of the CodeQL database file in bytes.
	Updated_at string `json:"updated_at"` // The date and time at which the CodeQL database was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Uploader GeneratedType_Simple_user `json:"uploader"` // A GitHub user.
}

// GeneratedType_Webhook_rubygems_metadata represents the GeneratedType_Webhook_rubygems_metadata schema from the OpenAPI specification
type GeneratedType_Webhook_rubygems_metadata struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Repo string `json:"repo,omitempty"`
	Version_info map[string]interface{} `json:"version_info,omitempty"`
	Readme string `json:"readme,omitempty"`
	Description string `json:"description,omitempty"`
	Name string `json:"name,omitempty"`
	Platform string `json:"platform,omitempty"`
	Commit_oid string `json:"commit_oid,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Dependencies []map[string]interface{} `json:"dependencies,omitempty"`
}

// GeneratedType_Pages_health_check represents the GeneratedType_Pages_health_check schema from the OpenAPI specification
type GeneratedType_Pages_health_check struct {
	Alt_domain map[string]interface{} `json:"alt_domain,omitempty"`
	Domain map[string]interface{} `json:"domain,omitempty"`
}

// GeneratedType_Referenced_workflow represents the GeneratedType_Referenced_workflow schema from the OpenAPI specification
type GeneratedType_Referenced_workflow struct {
	Ref string `json:"ref,omitempty"`
	Sha string `json:"sha"`
	Path string `json:"path"`
}

// Migration represents the Migration schema from the OpenAPI specification
type Migration struct {
	State string `json:"state"`
	Exclude_metadata bool `json:"exclude_metadata"`
	Archive_url string `json:"archive_url,omitempty"`
	Lock_repositories bool `json:"lock_repositories"`
	Url string `json:"url"`
	Exclude_git_data bool `json:"exclude_git_data"`
	Guid string `json:"guid"`
	Id int64 `json:"id"`
	Exclude_owner_projects bool `json:"exclude_owner_projects"`
	Owner GeneratedType_Nullable_simple_user `json:"owner"` // A GitHub user.
	Org_metadata_only bool `json:"org_metadata_only"`
	Repositories []Repository `json:"repositories"` // The repositories included in the migration. Only returned for export migrations.
	Exclude_attachments bool `json:"exclude_attachments"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Exclude []string `json:"exclude,omitempty"` // Exclude related items from being returned in the response in order to improve performance of the request. The array can include any of: `"repositories"`.
	Exclude_releases bool `json:"exclude_releases"`
	Node_id string `json:"node_id"`
}

// GeneratedType_Dependabot_public_key represents the GeneratedType_Dependabot_public_key schema from the OpenAPI specification
type GeneratedType_Dependabot_public_key struct {
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
}

// GeneratedType_Webhook_check_run_rerequested_form_encoded represents the GeneratedType_Webhook_check_run_rerequested_form_encoded schema from the OpenAPI specification
type GeneratedType_Webhook_check_run_rerequested_form_encoded struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.rerequested JSON payload. The decoded payload is a JSON object.
}

// GeneratedType_Review_request_removed_issue_event represents the GeneratedType_Review_request_removed_issue_event schema from the OpenAPI specification
type GeneratedType_Review_request_removed_issue_event struct {
	Requested_reviewer GeneratedType_Simple_user `json:"requested_reviewer,omitempty"` // A GitHub user.
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Review_requester GeneratedType_Simple_user `json:"review_requester"` // A GitHub user.
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
}

// GeneratedType_Nullable_simple_repository represents the GeneratedType_Nullable_simple_repository schema from the OpenAPI specification
type GeneratedType_Nullable_simple_repository struct {
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Teams_url string `json:"teams_url"` // The API URL to list the teams on the repository.
	Keys_url string `json:"keys_url"` // A template for the API URL to get information about deploy keys on the repository.
	Id int64 `json:"id"` // A unique identifier of the repository.
	Languages_url string `json:"languages_url"` // The API URL to get information about the languages of the repository.
	Html_url string `json:"html_url"` // The URL to view the repository on GitHub.com.
	Compare_url string `json:"compare_url"` // A template for the API URL to compare two commits or refs.
	Releases_url string `json:"releases_url"` // A template for the API URL to get information about releases on the repository.
	Fork bool `json:"fork"` // Whether the repository is a fork.
	Private bool `json:"private"` // Whether the repository is private.
	Full_name string `json:"full_name"` // The full, globally unique, name of the repository.
	Description string `json:"description"` // The repository description.
	Hooks_url string `json:"hooks_url"` // The API URL to list the hooks on the repository.
	Merges_url string `json:"merges_url"` // The API URL to merge branches in the repository.
	Issues_url string `json:"issues_url"` // A template for the API URL to get information about issues on the repository.
	Events_url string `json:"events_url"` // The API URL to list the events of the repository.
	Git_tags_url string `json:"git_tags_url"` // A template for the API URL to get information about Git tags of the repository.
	Pulls_url string `json:"pulls_url"` // A template for the API URL to get information about pull requests on the repository.
	Commits_url string `json:"commits_url"` // A template for the API URL to get information about commits on the repository.
	Contributors_url string `json:"contributors_url"` // A template for the API URL to list the contributors to the repository.
	Forks_url string `json:"forks_url"` // The API URL to list the forks of the repository.
	Subscription_url string `json:"subscription_url"` // The API URL to subscribe to notifications for this repository.
	Url string `json:"url"` // The URL to get more information about the repository from the GitHub API.
	Archive_url string `json:"archive_url"` // A template for the API URL to download the repository as an archive.
	Subscribers_url string `json:"subscribers_url"` // The API URL to list the subscribers on the repository.
	Labels_url string `json:"labels_url"` // A template for the API URL to get information about labels of the repository.
	Contents_url string `json:"contents_url"` // A template for the API URL to get the contents of the repository.
	Blobs_url string `json:"blobs_url"` // A template for the API URL to create or retrieve a raw Git blob in the repository.
	Collaborators_url string `json:"collaborators_url"` // A template for the API URL to get information about collaborators of the repository.
	Notifications_url string `json:"notifications_url"` // A template for the API URL to get information about notifications on the repository.
	Branches_url string `json:"branches_url"` // A template for the API URL to get information about branches in the repository.
	Assignees_url string `json:"assignees_url"` // A template for the API URL to list the available assignees for issues in the repository.
	Statuses_url string `json:"statuses_url"` // A template for the API URL to get information about statuses of a commit.
	Git_commits_url string `json:"git_commits_url"` // A template for the API URL to get information about Git commits of the repository.
	Deployments_url string `json:"deployments_url"` // The API URL to list the deployments of the repository.
	Trees_url string `json:"trees_url"` // A template for the API URL to create or retrieve a raw Git tree of the repository.
	Comments_url string `json:"comments_url"` // A template for the API URL to get information about comments on the repository.
	Stargazers_url string `json:"stargazers_url"` // The API URL to list the stargazers on the repository.
	Name string `json:"name"` // The name of the repository.
	Node_id string `json:"node_id"` // The GraphQL identifier of the repository.
	Git_refs_url string `json:"git_refs_url"` // A template for the API URL to get information about Git refs of the repository.
	Issue_comment_url string `json:"issue_comment_url"` // A template for the API URL to get information about issue comments on the repository.
	Tags_url string `json:"tags_url"` // The API URL to get information about tags on the repository.
	Milestones_url string `json:"milestones_url"` // A template for the API URL to get information about milestones of the repository.
	Downloads_url string `json:"downloads_url"` // The API URL to list the downloads on the repository.
	Issue_events_url string `json:"issue_events_url"` // A template for the API URL to get information about issue events on the repository.
}

// GeneratedType_Repository_advisory_create represents the GeneratedType_Repository_advisory_create schema from the OpenAPI specification
type GeneratedType_Repository_advisory_create struct {
	Cve_id string `json:"cve_id,omitempty"` // The Common Vulnerabilities and Exposures (CVE) ID.
	Description string `json:"description"` // A detailed description of what the advisory impacts.
	Cvss_vector_string string `json:"cvss_vector_string,omitempty"` // The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
	Cwe_ids []string `json:"cwe_ids,omitempty"` // A list of Common Weakness Enumeration (CWE) IDs.
	Severity string `json:"severity,omitempty"` // The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
	Start_private_fork bool `json:"start_private_fork,omitempty"` // Whether to create a temporary private fork of the repository to collaborate on a fix.
	Summary string `json:"summary"` // A short summary of the advisory.
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities"` // A product affected by the vulnerability detailed in a repository security advisory.
	Credits []map[string]interface{} `json:"credits,omitempty"` // A list of users receiving credit for their participation in the security advisory.
}

// GeneratedType_Code_scanning_variant_analysis_repo_task represents the GeneratedType_Code_scanning_variant_analysis_repo_task schema from the OpenAPI specification
type GeneratedType_Code_scanning_variant_analysis_repo_task struct {
	Failure_message string `json:"failure_message,omitempty"` // The reason of the failure of this repo task. This is only available if the repository task has failed.
	Repository GeneratedType_Simple_repository `json:"repository"` // A GitHub repository.
	Result_count int `json:"result_count,omitempty"` // The number of results in the case of a successful analysis. This is only available for successful analyses.
	Source_location_prefix string `json:"source_location_prefix,omitempty"` // The source location prefix to use. This is only available for successful analyses.
	Analysis_status string `json:"analysis_status"` // The new status of the CodeQL variant analysis repository task.
	Artifact_size_in_bytes int `json:"artifact_size_in_bytes,omitempty"` // The size of the artifact. This is only available for successful analyses.
	Artifact_url string `json:"artifact_url,omitempty"` // The URL of the artifact. This is only available for successful analyses.
	Database_commit_sha string `json:"database_commit_sha,omitempty"` // The SHA of the commit the CodeQL database was built against. This is only available for successful analyses.
}

// GeneratedType_Content_submodule represents the GeneratedType_Content_submodule schema from the OpenAPI specification
type GeneratedType_Content_submodule struct {
	Download_url string `json:"download_url"`
	Html_url string `json:"html_url"`
	Name string `json:"name"`
	Sha string `json:"sha"`
	Submodule_git_url string `json:"submodule_git_url"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Git_url string `json:"git_url"`
	Path string `json:"path"`
	Size int `json:"size"`
	Links map[string]interface{} `json:"_links"`
}

// GeneratedType_Oidc_custom_sub_repo represents the GeneratedType_Oidc_custom_sub_repo schema from the OpenAPI specification
type GeneratedType_Oidc_custom_sub_repo struct {
	Include_claim_keys []string `json:"include_claim_keys,omitempty"` // Array of unique strings. Each claim key can only contain alphanumeric characters and underscores.
	Use_default bool `json:"use_default"` // Whether to use the default template or not. If `true`, the `include_claim_keys` field is ignored.
}

// GeneratedType_Webhook_pull_request_demilestoned represents the GeneratedType_Webhook_pull_request_demilestoned schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_demilestoned struct {
	Pull_request Webhookspullrequest5 `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Milestone Milestone `json:"milestone,omitempty"` // A collection of related issues and pull requests.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Repository_rule_ruleset_info represents the GeneratedType_Repository_rule_ruleset_info schema from the OpenAPI specification
type GeneratedType_Repository_rule_ruleset_info struct {
	Ruleset_source string `json:"ruleset_source,omitempty"` // The name of the source of the ruleset that includes this rule.
	Ruleset_source_type string `json:"ruleset_source_type,omitempty"` // The type of source for the ruleset that includes this rule.
	Ruleset_id int `json:"ruleset_id,omitempty"` // The ID of the ruleset that includes this rule.
}

// GeneratedType_Webhook_repository_ruleset_edited represents the GeneratedType_Webhook_repository_ruleset_edited schema from the OpenAPI specification
type GeneratedType_Webhook_repository_ruleset_edited struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_ruleset GeneratedType_Repository_ruleset `json:"repository_ruleset"` // A set of rules to apply when specified conditions are met.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Review_dismissed_issue_event represents the GeneratedType_Review_dismissed_issue_event schema from the OpenAPI specification
type GeneratedType_Review_dismissed_issue_event struct {
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Event string `json:"event"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Dismissed_review map[string]interface{} `json:"dismissed_review"`
	Created_at string `json:"created_at"`
}

// GeneratedType_Webhook_deployment_created represents the GeneratedType_Webhook_deployment_created schema from the OpenAPI specification
type GeneratedType_Webhook_deployment_created struct {
	Action string `json:"action"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Deployment map[string]interface{} `json:"deployment"` // The [deployment](https://docs.github.com/rest/deployments/deployments#list-deployments).
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Workflow Webhooksworkflow `json:"workflow"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_projects_v2_item_archived represents the GeneratedType_Webhook_projects_v2_item_archived schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_item_archived struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType_Projects_v2_item `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes Webhooksprojectchanges `json:"changes"`
}

// GeneratedType_Repository_rule_params_code_scanning_tool represents the GeneratedType_Repository_rule_params_code_scanning_tool schema from the OpenAPI specification
type GeneratedType_Repository_rule_params_code_scanning_tool struct {
	Tool string `json:"tool"` // The name of a code scanning tool
	Alerts_threshold string `json:"alerts_threshold"` // The severity level at which code scanning results that raise alerts block a reference update. For more information on alert severity levels, see "[About code scanning alerts](https://docs.github.com/code-security/code-scanning/managing-code-scanning-alerts/about-code-scanning-alerts#about-alert-severity-and-security-severity-levels)."
	Security_alerts_threshold string `json:"security_alerts_threshold"` // The severity level at which code scanning results that raise security alerts block a reference update. For more information on security severity levels, see "[About code scanning alerts](https://docs.github.com/code-security/code-scanning/managing-code-scanning-alerts/about-code-scanning-alerts#about-alert-severity-and-security-severity-levels)."
}

// GeneratedType_Projects_v2_item represents the GeneratedType_Projects_v2_item schema from the OpenAPI specification
type GeneratedType_Projects_v2_item struct {
	Updated_at string `json:"updated_at"`
	Content_type string `json:"content_type"` // The type of content tracked in a project item
	Project_node_id string `json:"project_node_id,omitempty"`
	Archived_at string `json:"archived_at"`
	Content_node_id string `json:"content_node_id"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id,omitempty"`
	Creator GeneratedType_Simple_user `json:"creator,omitempty"` // A GitHub user.
	Id float64 `json:"id"`
}

// GeneratedType_Secret_scanning_location_wiki_commit represents the GeneratedType_Secret_scanning_location_wiki_commit schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_wiki_commit struct {
	Blob_sha string `json:"blob_sha"` // SHA-1 hash ID of the associated blob
	Commit_sha string `json:"commit_sha"` // SHA-1 hash ID of the associated commit
	End_column float64 `json:"end_column"` // The column at which the secret ends within the end line when the file is interpreted as 8-bit ASCII.
	Page_url string `json:"page_url"` // The GitHub URL to get the associated wiki page
	Start_line float64 `json:"start_line"` // Line number at which the secret starts in the file
	Commit_url string `json:"commit_url"` // The GitHub URL to get the associated wiki commit
	End_line float64 `json:"end_line"` // Line number at which the secret ends in the file
	Path string `json:"path"` // The file path of the wiki page
	Start_column float64 `json:"start_column"` // The column at which the secret starts within the start line when the file is interpreted as 8-bit ASCII.
}

// GeneratedType_Nullable_milestone represents the GeneratedType_Nullable_milestone schema from the OpenAPI specification
type GeneratedType_Nullable_milestone struct {
	Created_at string `json:"created_at"`
	Creator GeneratedType_Nullable_simple_user `json:"creator"` // A GitHub user.
	Labels_url string `json:"labels_url"`
	Updated_at string `json:"updated_at"`
	Id int `json:"id"`
	State string `json:"state"` // The state of the milestone.
	Closed_issues int `json:"closed_issues"`
	Due_on string `json:"due_on"`
	Description string `json:"description"`
	Title string `json:"title"` // The title of the milestone.
	Open_issues int `json:"open_issues"`
	Number int `json:"number"` // The number of the milestone.
	Url string `json:"url"`
	Node_id string `json:"node_id"`
	Closed_at string `json:"closed_at"`
	Html_url string `json:"html_url"`
}

// GeneratedType_Webhook_team_created represents the GeneratedType_Webhook_team_created schema from the OpenAPI specification
type GeneratedType_Webhook_team_created struct {
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Hook_response represents the GeneratedType_Hook_response schema from the OpenAPI specification
type GeneratedType_Hook_response struct {
	Status string `json:"status"`
	Code int `json:"code"`
	Message string `json:"message"`
}

// Manifest represents the Manifest schema from the OpenAPI specification
type Manifest struct {
	File map[string]interface{} `json:"file,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Name string `json:"name"` // The name of the manifest.
	Resolved map[string]interface{} `json:"resolved,omitempty"` // A collection of resolved package dependencies.
}

// GeneratedType_Link_with_type represents the GeneratedType_Link_with_type schema from the OpenAPI specification
type GeneratedType_Link_with_type struct {
	Href string `json:"href"`
	TypeField string `json:"type"`
}

// Package represents the Package schema from the OpenAPI specification
type Package struct {
	Created_at string `json:"created_at"`
	Package_type string `json:"package_type"`
	Html_url string `json:"html_url"`
	Repository GeneratedType_Nullable_minimal_repository `json:"repository,omitempty"` // Minimal Repository
	Id int `json:"id"` // Unique identifier of the package.
	Name string `json:"name"` // The name of the package.
	Owner GeneratedType_Nullable_simple_user `json:"owner,omitempty"` // A GitHub user.
	Version_count int `json:"version_count"` // The number of versions of the package.
	Visibility string `json:"visibility"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
}

// GeneratedType_Webhook_issues_unlabeled represents the GeneratedType_Webhook_issues_unlabeled schema from the OpenAPI specification
type GeneratedType_Webhook_issues_unlabeled struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Label Webhookslabel `json:"label,omitempty"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Code_scanning_default_setup_options represents the GeneratedType_Code_scanning_default_setup_options schema from the OpenAPI specification
type GeneratedType_Code_scanning_default_setup_options struct {
	Runner_type string `json:"runner_type,omitempty"` // Whether to use labeled runners or standard GitHub runners.
	Runner_label string `json:"runner_label,omitempty"` // The label of the runner to use for code scanning default setup when runner_type is 'labeled'.
}

// GeneratedType_Webhook_release_published represents the GeneratedType_Webhook_release_published schema from the OpenAPI specification
type GeneratedType_Webhook_release_published struct {
	Release Webhooksrelease1 `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Minimal_repository represents the GeneratedType_Minimal_repository schema from the OpenAPI specification
type GeneratedType_Minimal_repository struct {
	Allow_forking bool `json:"allow_forking,omitempty"`
	Forks int `json:"forks,omitempty"`
	Has_pages bool `json:"has_pages,omitempty"`
	Size int `json:"size,omitempty"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Url string `json:"url"`
	Has_downloads bool `json:"has_downloads,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Downloads_url string `json:"downloads_url"`
	Watchers_count int `json:"watchers_count,omitempty"`
	Compare_url string `json:"compare_url"`
	Milestones_url string `json:"milestones_url"`
	Has_issues bool `json:"has_issues,omitempty"`
	Git_url string `json:"git_url,omitempty"`
	Labels_url string `json:"labels_url"`
	Notifications_url string `json:"notifications_url"`
	Hooks_url string `json:"hooks_url"`
	Contents_url string `json:"contents_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Network_count int `json:"network_count,omitempty"`
	Languages_url string `json:"languages_url"`
	Pulls_url string `json:"pulls_url"`
	Commits_url string `json:"commits_url"`
	Has_wiki bool `json:"has_wiki,omitempty"`
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Role_name string `json:"role_name,omitempty"`
	Contributors_url string `json:"contributors_url"`
	Mirror_url string `json:"mirror_url,omitempty"`
	Archive_url string `json:"archive_url"`
	Keys_url string `json:"keys_url"`
	Trees_url string `json:"trees_url"`
	Pushed_at string `json:"pushed_at,omitempty"`
	Merges_url string `json:"merges_url"`
	Blobs_url string `json:"blobs_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Default_branch string `json:"default_branch,omitempty"`
	Git_commits_url string `json:"git_commits_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Forks_count int `json:"forks_count,omitempty"`
	Has_projects bool `json:"has_projects,omitempty"`
	Issues_url string `json:"issues_url"`
	Node_id string `json:"node_id"`
	Subscribers_url string `json:"subscribers_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Assignees_url string `json:"assignees_url"`
	Language string `json:"language,omitempty"`
	Id int64 `json:"id"`
	Stargazers_url string `json:"stargazers_url"`
	Statuses_url string `json:"statuses_url"`
	Private bool `json:"private"`
	Deployments_url string `json:"deployments_url"`
	Stargazers_count int `json:"stargazers_count,omitempty"`
	Forks_url string `json:"forks_url"`
	Html_url string `json:"html_url"`
	Name string `json:"name"`
	Tags_url string `json:"tags_url"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Comments_url string `json:"comments_url"`
	Fork bool `json:"fork"`
	Teams_url string `json:"teams_url"`
	Security_and_analysis GeneratedType_Security_and_analysis `json:"security_and_analysis,omitempty"`
	Svn_url string `json:"svn_url,omitempty"`
	Git_tags_url string `json:"git_tags_url"`
	Topics []string `json:"topics,omitempty"`
	Releases_url string `json:"releases_url"`
	Open_issues_count int `json:"open_issues_count,omitempty"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	License map[string]interface{} `json:"license,omitempty"`
	Branches_url string `json:"branches_url"`
	Code_of_conduct GeneratedType_Code_of_conduct `json:"code_of_conduct,omitempty"` // Code Of Conduct
	Ssh_url string `json:"ssh_url,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	Issue_comment_url string `json:"issue_comment_url"`
	Collaborators_url string `json:"collaborators_url"`
	Events_url string `json:"events_url"`
	Issue_events_url string `json:"issue_events_url"`
	Is_template bool `json:"is_template,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Watchers int `json:"watchers,omitempty"`
	Description string `json:"description"`
	Git_refs_url string `json:"git_refs_url"`
	Open_issues int `json:"open_issues,omitempty"`
	Archived bool `json:"archived,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Clone_url string `json:"clone_url,omitempty"`
	Full_name string `json:"full_name"`
	Subscription_url string `json:"subscription_url"`
	Visibility string `json:"visibility,omitempty"`
}

// GeneratedType_Simple_classroom_repository represents the GeneratedType_Simple_classroom_repository schema from the OpenAPI specification
type GeneratedType_Simple_classroom_repository struct {
	Private bool `json:"private"` // Whether the repository is private.
	Default_branch string `json:"default_branch"` // The default branch for the repository.
	Full_name string `json:"full_name"` // The full, globally unique name of the repository.
	Html_url string `json:"html_url"` // The URL to view the repository on GitHub.com.
	Id int `json:"id"` // A unique identifier of the repository.
	Node_id string `json:"node_id"` // The GraphQL identifier of the repository.
}

// GeneratedType_Webhook_personal_access_token_request_denied represents the GeneratedType_Webhook_personal_access_token_request_denied schema from the OpenAPI specification
type GeneratedType_Webhook_personal_access_token_request_denied struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Personal_access_token_request GeneratedType_Personal_access_token_request `json:"personal_access_token_request"` // Details of a Personal Access Token Request.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Environment_approvals represents the GeneratedType_Environment_approvals schema from the OpenAPI specification
type GeneratedType_Environment_approvals struct {
	Environments []map[string]interface{} `json:"environments"` // The list of environments that were approved or rejected
	State string `json:"state"` // Whether deployment to the environment(s) was approved or rejected or pending (with comments)
	User GeneratedType_Simple_user `json:"user"` // A GitHub user.
	Comment string `json:"comment"` // The comment submitted with the deployment review
}

// GeneratedType_Scim_error represents the GeneratedType_Scim_error schema from the OpenAPI specification
type GeneratedType_Scim_error struct {
	Documentation_url string `json:"documentation_url,omitempty"`
	Message string `json:"message,omitempty"`
	Schemas []string `json:"schemas,omitempty"`
	Scimtype string `json:"scimType,omitempty"`
	Status int `json:"status,omitempty"`
	Detail string `json:"detail,omitempty"`
}

// GeneratedType_Copilot_seat_details represents the GeneratedType_Copilot_seat_details schema from the OpenAPI specification
type GeneratedType_Copilot_seat_details struct {
	Assigning_team interface{} `json:"assigning_team,omitempty"` // The team through which the assignee is granted access to GitHub Copilot, if applicable.
	Last_activity_at string `json:"last_activity_at,omitempty"` // Timestamp of user's last GitHub Copilot activity, in ISO 8601 format.
	Last_activity_editor string `json:"last_activity_editor,omitempty"` // Last editor that was used by the user for a GitHub Copilot completion.
	Pending_cancellation_date string `json:"pending_cancellation_date,omitempty"` // The pending cancellation date for the seat, in `YYYY-MM-DD` format. This will be null unless the assignee's Copilot access has been canceled during the current billing cycle. If the seat has been cancelled, this corresponds to the start of the organization's next billing cycle.
	Updated_at string `json:"updated_at,omitempty"` // **Closing down notice:** This field is no longer relevant and is closing down. Use the `created_at` field to determine when the assignee was last granted access to GitHub Copilot. Timestamp of when the assignee's GitHub Copilot access was last updated, in ISO 8601 format.
	Assignee GeneratedType_Nullable_simple_user `json:"assignee,omitempty"` // A GitHub user.
	Created_at string `json:"created_at"` // Timestamp of when the assignee was last granted access to GitHub Copilot, in ISO 8601 format.
	Organization GeneratedType_Nullable_organization_simple `json:"organization,omitempty"` // A GitHub organization.
	Plan_type string `json:"plan_type,omitempty"` // The Copilot plan of the organization, or the parent enterprise, when applicable.
}

// GeneratedType_Webhook_discussion_unlabeled represents the GeneratedType_Webhook_discussion_unlabeled schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_unlabeled struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
}

// GeneratedType_Converted_note_to_issue_issue_event represents the GeneratedType_Converted_note_to_issue_issue_event schema from the OpenAPI specification
type GeneratedType_Converted_note_to_issue_issue_event struct {
	Node_id string `json:"node_id"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Event string `json:"event"`
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_url string `json:"commit_url"`
}

// GeneratedType_Actions_hosted_runner represents the GeneratedType_Actions_hosted_runner schema from the OpenAPI specification
type GeneratedType_Actions_hosted_runner struct {
	Image_details GeneratedType_Nullable_actions_hosted_runner_pool_image `json:"image_details"` // Provides details of a hosted runner image
	Platform string `json:"platform"` // The operating system of the image.
	Id int `json:"id"` // The unique identifier of the hosted runner.
	Last_active_on string `json:"last_active_on,omitempty"` // The time at which the runner was last used, in ISO 8601 format.
	Machine_size_details GeneratedType_Actions_hosted_runner_machine_spec `json:"machine_size_details"` // Provides details of a particular machine spec.
	Public_ips []GeneratedType_Public_ip `json:"public_ips,omitempty"` // The public IP ranges when public IP is enabled for the hosted runners.
	Runner_group_id int `json:"runner_group_id,omitempty"` // The unique identifier of the group that the hosted runner belongs to.
	Status string `json:"status"` // The status of the runner.
	Public_ip_enabled bool `json:"public_ip_enabled"` // Whether public IP is enabled for the hosted runners.
	Maximum_runners int `json:"maximum_runners,omitempty"` // The maximum amount of hosted runners. Runners will not scale automatically above this number. Use this setting to limit your cost.
	Name string `json:"name"` // The name of the hosted runner.
}

// GeneratedType_Webhook_pull_request_closed represents the GeneratedType_Webhook_pull_request_closed schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_closed struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType_Pull_request_webhook `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
}

// GeneratedType_Webhook_project_card_edited represents the GeneratedType_Webhook_project_card_edited schema from the OpenAPI specification
type GeneratedType_Webhook_project_card_edited struct {
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card Webhooksprojectcard `json:"project_card"`
}

// GeneratedType_Webhook_repository_privatized represents the GeneratedType_Webhook_repository_privatized schema from the OpenAPI specification
type GeneratedType_Webhook_repository_privatized struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_repository_archived represents the GeneratedType_Webhook_repository_archived schema from the OpenAPI specification
type GeneratedType_Webhook_repository_archived struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_workflow_run_in_progress represents the GeneratedType_Webhook_workflow_run_in_progress schema from the OpenAPI specification
type GeneratedType_Webhook_workflow_run_in_progress struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Workflow Webhooksworkflow `json:"workflow"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Timeline_assigned_issue_event represents the GeneratedType_Timeline_assigned_issue_event schema from the OpenAPI specification
type GeneratedType_Timeline_assigned_issue_event struct {
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
	Id int `json:"id"`
	Assignee GeneratedType_Simple_user `json:"assignee"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
}

// GeneratedType_Code_scanning_sarifs_receipt represents the GeneratedType_Code_scanning_sarifs_receipt schema from the OpenAPI specification
type GeneratedType_Code_scanning_sarifs_receipt struct {
	Id string `json:"id,omitempty"` // An identifier for the upload.
	Url string `json:"url,omitempty"` // The REST API URL for checking the status of the upload.
}

// GeneratedType_User_marketplace_purchase represents the GeneratedType_User_marketplace_purchase schema from the OpenAPI specification
type GeneratedType_User_marketplace_purchase struct {
	Next_billing_date string `json:"next_billing_date"`
	On_free_trial bool `json:"on_free_trial"`
	Plan GeneratedType_Marketplace_listing_plan `json:"plan"` // Marketplace Listing Plan
	Unit_count int `json:"unit_count"`
	Updated_at string `json:"updated_at"`
	Account GeneratedType_Marketplace_account `json:"account"`
	Billing_cycle string `json:"billing_cycle"`
	Free_trial_ends_on string `json:"free_trial_ends_on"`
}

// GeneratedType_Commit_comment represents the GeneratedType_Commit_comment schema from the OpenAPI specification
type GeneratedType_Commit_comment struct {
	Id int `json:"id"`
	Position int `json:"position"`
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
	Line int `json:"line"`
	Updated_at string `json:"updated_at"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Html_url string `json:"html_url"`
	Path string `json:"path"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"`
}

// Webhooksproject represents the Webhooksproject schema from the OpenAPI specification
type Webhooksproject struct {
	Created_at string `json:"created_at"`
	Body string `json:"body"` // Body of the project
	Id int `json:"id"`
	Updated_at string `json:"updated_at"`
	Number int `json:"number"`
	State string `json:"state"` // State of the project; either 'open' or 'closed'
	Url string `json:"url"`
	Creator map[string]interface{} `json:"creator"`
	Html_url string `json:"html_url"`
	Name string `json:"name"` // Name of the project
	Node_id string `json:"node_id"`
	Owner_url string `json:"owner_url"`
	Columns_url string `json:"columns_url"`
}

// Repository represents the Repository schema from the OpenAPI specification
type Repository struct {
	Open_issues_count int `json:"open_issues_count"`
	Tags_url string `json:"tags_url"`
	Git_commits_url string `json:"git_commits_url"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Stargazers_count int `json:"stargazers_count"`
	Notifications_url string `json:"notifications_url"`
	Id int64 `json:"id"` // Unique identifier of the repository
	Mirror_url string `json:"mirror_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Branches_url string `json:"branches_url"`
	Pulls_url string `json:"pulls_url"`
	Watchers_count int `json:"watchers_count"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Svn_url string `json:"svn_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Comments_url string `json:"comments_url"`
	Assignees_url string `json:"assignees_url"`
	Language string `json:"language"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Deployments_url string `json:"deployments_url"`
	Issue_events_url string `json:"issue_events_url"`
	Name string `json:"name"` // The name of the repository.
	Releases_url string `json:"releases_url"`
	Downloads_url string `json:"downloads_url"`
	Url string `json:"url"`
	Topics []string `json:"topics,omitempty"`
	Archive_url string `json:"archive_url"`
	Pushed_at string `json:"pushed_at"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Clone_url string `json:"clone_url"`
	Full_name string `json:"full_name"`
	Issue_comment_url string `json:"issue_comment_url"`
	Milestones_url string `json:"milestones_url"`
	Git_tags_url string `json:"git_tags_url"`
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Blobs_url string `json:"blobs_url"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Teams_url string `json:"teams_url"`
	Git_url string `json:"git_url"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Has_pages bool `json:"has_pages"`
	Ssh_url string `json:"ssh_url"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Git_refs_url string `json:"git_refs_url"`
	Languages_url string `json:"languages_url"`
	Forks_url string `json:"forks_url"`
	Statuses_url string `json:"statuses_url"`
	Labels_url string `json:"labels_url"`
	Merges_url string `json:"merges_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Forks int `json:"forks"`
	Subscription_url string `json:"subscription_url"`
	Events_url string `json:"events_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	License GeneratedType_Nullable_license_simple `json:"license"` // License Simple
	Subscribers_url string `json:"subscribers_url"`
	Collaborators_url string `json:"collaborators_url"`
	Stargazers_url string `json:"stargazers_url"`
	Contents_url string `json:"contents_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Issues_url string `json:"issues_url"`
	Contributors_url string `json:"contributors_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Node_id string `json:"node_id"`
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Html_url string `json:"html_url"`
	Fork bool `json:"fork"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Trees_url string `json:"trees_url"`
	Created_at string `json:"created_at"`
	Homepage string `json:"homepage"`
	Updated_at string `json:"updated_at"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Keys_url string `json:"keys_url"`
	Forks_count int `json:"forks_count"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Code_search_index_status map[string]interface{} `json:"code_search_index_status,omitempty"` // The status of the code search index for this repository
	Commits_url string `json:"commits_url"`
	Open_issues int `json:"open_issues"`
	Hooks_url string `json:"hooks_url"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Compare_url string `json:"compare_url"`
	Description string `json:"description"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Watchers int `json:"watchers"`
}

// GeneratedType_Moved_column_in_project_issue_event represents the GeneratedType_Moved_column_in_project_issue_event schema from the OpenAPI specification
type GeneratedType_Moved_column_in_project_issue_event struct {
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Url string `json:"url"`
	Id int `json:"id"`
	Commit_id string `json:"commit_id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
}

// GeneratedType_Content_file represents the GeneratedType_Content_file schema from the OpenAPI specification
type GeneratedType_Content_file struct {
	Encoding string `json:"encoding"`
	Sha string `json:"sha"`
	Submodule_git_url string `json:"submodule_git_url,omitempty"`
	Name string `json:"name"`
	Content string `json:"content"`
	Git_url string `json:"git_url"`
	Html_url string `json:"html_url"`
	Path string `json:"path"`
	Download_url string `json:"download_url"`
	Size int `json:"size"`
	Target string `json:"target,omitempty"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Links map[string]interface{} `json:"_links"`
}

// GeneratedType_Review_custom_gates_comment_required represents the GeneratedType_Review_custom_gates_comment_required schema from the OpenAPI specification
type GeneratedType_Review_custom_gates_comment_required struct {
	Comment string `json:"comment"` // Comment associated with the pending deployment protection rule. **Required when state is not provided.**
	Environment_name string `json:"environment_name"` // The name of the environment to approve or reject.
}

// GeneratedType_Webhook_workflow_run_completed represents the GeneratedType_Webhook_workflow_run_completed schema from the OpenAPI specification
type GeneratedType_Webhook_workflow_run_completed struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Workflow Webhooksworkflow `json:"workflow"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
}

// GeneratedType_Private_user represents the GeneratedType_Private_user schema from the OpenAPI specification
type GeneratedType_Private_user struct {
	Total_private_repos int `json:"total_private_repos"`
	Notification_email string `json:"notification_email,omitempty"`
	Url string `json:"url"`
	Gravatar_id string `json:"gravatar_id"`
	Node_id string `json:"node_id"`
	Html_url string `json:"html_url"`
	Site_admin bool `json:"site_admin"`
	Subscriptions_url string `json:"subscriptions_url"`
	Repos_url string `json:"repos_url"`
	Location string `json:"location"`
	Events_url string `json:"events_url"`
	Following_url string `json:"following_url"`
	TypeField string `json:"type"`
	Avatar_url string `json:"avatar_url"`
	Created_at string `json:"created_at"`
	Owned_private_repos int `json:"owned_private_repos"`
	Ldap_dn string `json:"ldap_dn,omitempty"`
	Starred_url string `json:"starred_url"`
	Gists_url string `json:"gists_url"`
	Private_gists int `json:"private_gists"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Company string `json:"company"`
	Hireable bool `json:"hireable"`
	Updated_at string `json:"updated_at"`
	Email string `json:"email"`
	Disk_usage int `json:"disk_usage"`
	Login string `json:"login"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Bio string `json:"bio"`
	User_view_type string `json:"user_view_type,omitempty"`
	Public_gists int `json:"public_gists"`
	Collaborators int `json:"collaborators"`
	Followers_url string `json:"followers_url"`
	Two_factor_authentication bool `json:"two_factor_authentication"`
	Id int64 `json:"id"`
	Blog string `json:"blog"`
	Following int `json:"following"`
	Name string `json:"name"`
	Public_repos int `json:"public_repos"`
	Organizations_url string `json:"organizations_url"`
	Followers int `json:"followers"`
	Received_events_url string `json:"received_events_url"`
	Business_plus bool `json:"business_plus,omitempty"`
}

// GeneratedType_Gist_simple represents the GeneratedType_Gist_simple schema from the OpenAPI specification
type GeneratedType_Gist_simple struct {
	Owner GeneratedType_Simple_user `json:"owner,omitempty"` // A GitHub user.
	Git_pull_url string `json:"git_pull_url,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Forks []map[string]interface{} `json:"forks,omitempty"`
	History []GeneratedType_Gist_history `json:"history,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Public bool `json:"public,omitempty"`
	Truncated bool `json:"truncated,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	User string `json:"user,omitempty"`
	Comments int `json:"comments,omitempty"`
	Git_push_url string `json:"git_push_url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Comments_url string `json:"comments_url,omitempty"`
	Forks_url string `json:"forks_url,omitempty"`
	Id string `json:"id,omitempty"`
	Comments_enabled bool `json:"comments_enabled,omitempty"`
	Files map[string]interface{} `json:"files,omitempty"`
	Fork_of map[string]interface{} `json:"fork_of,omitempty"` // Gist
	Url string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	Commits_url string `json:"commits_url,omitempty"`
}

// GeneratedType_Installation_token represents the GeneratedType_Installation_token schema from the OpenAPI specification
type GeneratedType_Installation_token struct {
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Permissions GeneratedType_App_permissions `json:"permissions,omitempty"` // The permissions granted to the user access token.
	Repositories []Repository `json:"repositories,omitempty"`
	Repository_selection string `json:"repository_selection,omitempty"`
	Single_file string `json:"single_file,omitempty"`
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Token string `json:"token"`
	Expires_at string `json:"expires_at"`
}

// GeneratedType_Runner_groups_org represents the GeneratedType_Runner_groups_org schema from the OpenAPI specification
type GeneratedType_Runner_groups_org struct {
	Selected_workflows []string `json:"selected_workflows,omitempty"` // List of workflows the runner group should be allowed to run. This setting will be ignored unless `restricted_to_workflows` is set to `true`.
	Visibility string `json:"visibility"`
	Workflow_restrictions_read_only bool `json:"workflow_restrictions_read_only,omitempty"` // If `true`, the `restricted_to_workflows` and `selected_workflows` fields cannot be modified.
	Hosted_runners_url string `json:"hosted_runners_url,omitempty"`
	Inherited bool `json:"inherited"`
	Name string `json:"name"`
	Allows_public_repositories bool `json:"allows_public_repositories"`
	Id float64 `json:"id"`
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // Link to the selected repositories resource for this runner group. Not present unless visibility was set to `selected`
	DefaultField bool `json:"default"`
	Inherited_allows_public_repositories bool `json:"inherited_allows_public_repositories,omitempty"`
	Network_configuration_id string `json:"network_configuration_id,omitempty"` // The identifier of a hosted compute network configuration.
	Restricted_to_workflows bool `json:"restricted_to_workflows,omitempty"` // If `true`, the runner group will be restricted to running only the workflows specified in the `selected_workflows` array.
	Runners_url string `json:"runners_url"`
}

// GeneratedType_Secret_scanning_alert_webhook represents the GeneratedType_Secret_scanning_alert_webhook schema from the OpenAPI specification
type GeneratedType_Secret_scanning_alert_webhook struct {
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypass_request_html_url string `json:"push_protection_bypass_request_html_url,omitempty"` // The URL to a push protection bypass request.
	Push_protection_bypassed_by GeneratedType_Nullable_simple_user `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Number int `json:"number,omitempty"` // The security alert number.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Supported secret scanning patterns](https://docs.github.com/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
	Publicly_leaked bool `json:"publicly_leaked,omitempty"` // Whether the detected secret was publicly leaked.
	Resolved_by GeneratedType_Nullable_simple_user `json:"resolved_by,omitempty"` // A GitHub user.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
	Push_protection_bypass_request_comment string `json:"push_protection_bypass_request_comment,omitempty"` // An optional comment when requesting a push protection bypass.
	Resolution_comment string `json:"resolution_comment,omitempty"` // An optional comment to resolve an alert.
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
	Resolution string `json:"resolution,omitempty"` // The reason for resolving the alert.
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	Validity string `json:"validity,omitempty"` // The token status as of the latest validity check.
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
	Multi_repo bool `json:"multi_repo,omitempty"` // Whether the detected secret was found in multiple repositories in the same organization or business.
	Push_protection_bypass_request_reviewer GeneratedType_Nullable_simple_user `json:"push_protection_bypass_request_reviewer,omitempty"` // A GitHub user.
	Push_protection_bypass_request_reviewer_comment string `json:"push_protection_bypass_request_reviewer_comment,omitempty"` // An optional comment when reviewing a push protection bypass.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType_Deploy_key represents the GeneratedType_Deploy_key schema from the OpenAPI specification
type GeneratedType_Deploy_key struct {
	Added_by string `json:"added_by,omitempty"`
	Created_at string `json:"created_at"`
	Read_only bool `json:"read_only"`
	Title string `json:"title"`
	Verified bool `json:"verified"`
	Url string `json:"url"`
	Enabled bool `json:"enabled,omitempty"`
	Id int `json:"id"`
	Last_used string `json:"last_used,omitempty"`
	Key string `json:"key"`
}

// GeneratedType_Repo_search_result_item represents the GeneratedType_Repo_search_result_item schema from the OpenAPI specification
type GeneratedType_Repo_search_result_item struct {
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Trees_url string `json:"trees_url"`
	Language string `json:"language"`
	Has_projects bool `json:"has_projects"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Blobs_url string `json:"blobs_url"`
	Events_url string `json:"events_url"`
	Description string `json:"description"`
	Owner GeneratedType_Nullable_simple_user `json:"owner"` // A GitHub user.
	Issues_url string `json:"issues_url"`
	License GeneratedType_Nullable_license_simple `json:"license"` // License Simple
	Has_issues bool `json:"has_issues"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"`
	Has_wiki bool `json:"has_wiki"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"`
	Tags_url string `json:"tags_url"`
	Compare_url string `json:"compare_url"`
	Git_refs_url string `json:"git_refs_url"`
	Statuses_url string `json:"statuses_url"`
	Downloads_url string `json:"downloads_url"`
	Mirror_url string `json:"mirror_url"`
	Archived bool `json:"archived"`
	Branches_url string `json:"branches_url"`
	Topics []string `json:"topics,omitempty"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Clone_url string `json:"clone_url"`
	Languages_url string `json:"languages_url"`
	Size int `json:"size"`
	Comments_url string `json:"comments_url"`
	Updated_at string `json:"updated_at"`
	Fork bool `json:"fork"`
	Git_commits_url string `json:"git_commits_url"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Commits_url string `json:"commits_url"`
	Milestones_url string `json:"milestones_url"`
	Node_id string `json:"node_id"`
	Notifications_url string `json:"notifications_url"`
	Stargazers_count int `json:"stargazers_count"`
	Private bool `json:"private"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"`
	Issue_comment_url string `json:"issue_comment_url"`
	Labels_url string `json:"labels_url"`
	Subscribers_url string `json:"subscribers_url"`
	Id int `json:"id"`
	Contributors_url string `json:"contributors_url"`
	Score float64 `json:"score"`
	Svn_url string `json:"svn_url"`
	Stargazers_url string `json:"stargazers_url"`
	Subscription_url string `json:"subscription_url"`
	Hooks_url string `json:"hooks_url"`
	Homepage string `json:"homepage"`
	Issue_events_url string `json:"issue_events_url"`
	Open_issues int `json:"open_issues"`
	Created_at string `json:"created_at"`
	Forks_count int `json:"forks_count"`
	Html_url string `json:"html_url"`
	Contents_url string `json:"contents_url"`
	Is_template bool `json:"is_template,omitempty"`
	Keys_url string `json:"keys_url"`
	Pulls_url string `json:"pulls_url"`
	Has_downloads bool `json:"has_downloads"`
	Git_tags_url string `json:"git_tags_url"`
	Has_pages bool `json:"has_pages"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Forks_url string `json:"forks_url"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Merges_url string `json:"merges_url"`
	Teams_url string `json:"teams_url"`
	Watchers_count int `json:"watchers_count"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Git_url string `json:"git_url"`
	Archive_url string `json:"archive_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Deployments_url string `json:"deployments_url"`
	Full_name string `json:"full_name"`
	Watchers int `json:"watchers"`
	Collaborators_url string `json:"collaborators_url"`
	Default_branch string `json:"default_branch"`
	Releases_url string `json:"releases_url"`
	Url string `json:"url"`
	Ssh_url string `json:"ssh_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Open_issues_count int `json:"open_issues_count"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"`
	Pushed_at string `json:"pushed_at"`
	Assignees_url string `json:"assignees_url"`
	Name string `json:"name"`
	Forks int `json:"forks"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
}

// Webhooksreviewcomment represents the Webhooksreviewcomment schema from the OpenAPI specification
type Webhooksreviewcomment struct {
	Subject_type string `json:"subject_type,omitempty"` // The level at which the comment is targeted, can be a diff line or a file.
	Body string `json:"body"` // The text of the comment.
	Original_commit_id string `json:"original_commit_id"` // The SHA of the original commit to which the comment applies.
	Original_line int `json:"original_line"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Links map[string]interface{} `json:"_links"`
	Id int `json:"id"` // The ID of the pull request review comment.
	Pull_request_url string `json:"pull_request_url"` // URL for the pull request that the review comment belongs to.
	Start_side string `json:"start_side"` // The side of the first line of the range for a multi-line comment.
	Updated_at string `json:"updated_at"`
	Diff_hunk string `json:"diff_hunk"` // The diff of the line that the comment refers to.
	Original_position int `json:"original_position"` // The index of the original line in the diff to which the comment applies.
	Position int `json:"position"` // The line index in the diff to which the comment applies.
	Start_line int `json:"start_line"` // The first line of the range for a multi-line comment.
	Url string `json:"url"` // URL for the pull request review comment
	Commit_id string `json:"commit_id"` // The SHA of the commit to which the comment applies.
	Side string `json:"side"` // The side of the first line of the range for a multi-line comment.
	User map[string]interface{} `json:"user"`
	Html_url string `json:"html_url"` // HTML URL for the pull request review comment.
	In_reply_to_id int `json:"in_reply_to_id,omitempty"` // The comment ID to reply to.
	Line int `json:"line"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Path string `json:"path"` // The relative path of the file to which the comment applies.
	Pull_request_review_id int `json:"pull_request_review_id"` // The ID of the pull request review to which the comment belongs.
	Original_start_line int `json:"original_start_line"` // The first line of the range for a multi-line comment.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"` // The node ID of the pull request review comment.
	Reactions map[string]interface{} `json:"reactions"`
}

// GeneratedType_Webhook_projects_v2_item_restored represents the GeneratedType_Webhook_projects_v2_item_restored schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_item_restored struct {
	Changes Webhooksprojectchanges `json:"changes"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType_Projects_v2_item `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Page represents the Page schema from the OpenAPI specification
type Page struct {
	Protected_domain_state string `json:"protected_domain_state,omitempty"` // The state if the domain is verified
	Cname string `json:"cname"` // The Pages site's custom domain
	Https_certificate GeneratedType_Pages_https_certificate `json:"https_certificate,omitempty"`
	Https_enforced bool `json:"https_enforced,omitempty"` // Whether https is enabled on the domain
	Status string `json:"status"` // The status of the most recent build of the Page.
	Html_url string `json:"html_url,omitempty"` // The web address the Page can be accessed from.
	Url string `json:"url"` // The API address for accessing this Page resource.
	Build_type string `json:"build_type,omitempty"` // The process in which the Page will be built.
	Pending_domain_unverified_at string `json:"pending_domain_unverified_at,omitempty"` // The timestamp when a pending domain becomes unverified.
	Public bool `json:"public"` // Whether the GitHub Pages site is publicly visible. If set to `true`, the site is accessible to anyone on the internet. If set to `false`, the site will only be accessible to users who have at least `read` access to the repository that published the site.
	Source GeneratedType_Pages_source_hash `json:"source,omitempty"`
	Custom_404 bool `json:"custom_404"` // Whether the Page has a custom 404 page.
}

// GeneratedType_Rate_limit represents the GeneratedType_Rate_limit schema from the OpenAPI specification
type GeneratedType_Rate_limit struct {
	Reset int `json:"reset"`
	Used int `json:"used"`
	Limit int `json:"limit"`
	Remaining int `json:"remaining"`
}

// Webhooksrelease represents the Webhooksrelease schema from the OpenAPI specification
type Webhooksrelease struct {
	Url string `json:"url"`
	Assets []map[string]interface{} `json:"assets"`
	Assets_url string `json:"assets_url"`
	Id int `json:"id"`
	Target_commitish string `json:"target_commitish"` // Specifies the commitish value that determines where the Git tag is created from.
	Draft bool `json:"draft"` // Whether the release is a draft or published
	Created_at string `json:"created_at"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Upload_url string `json:"upload_url"`
	Body string `json:"body"`
	Prerelease bool `json:"prerelease"` // Whether the release is identified as a prerelease or a full release.
	Tarball_url string `json:"tarball_url"`
	Author map[string]interface{} `json:"author"`
	Tag_name string `json:"tag_name"` // The name of the tag.
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Published_at string `json:"published_at"`
	Discussion_url string `json:"discussion_url,omitempty"`
	Zipball_url string `json:"zipball_url"`
	Html_url string `json:"html_url"`
}

// GeneratedType_Team_discussion_comment represents the GeneratedType_Team_discussion_comment schema from the OpenAPI specification
type GeneratedType_Team_discussion_comment struct {
	Body_version string `json:"body_version"` // The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
	Discussion_url string `json:"discussion_url"`
	Last_edited_at string `json:"last_edited_at"`
	Number int `json:"number"` // The unique sequence number of a team discussion comment.
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Author GeneratedType_Nullable_simple_user `json:"author"` // A GitHub user.
	Node_id string `json:"node_id"`
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Body string `json:"body"` // The main text of the comment.
	Body_html string `json:"body_html"`
}

// GeneratedType_Webhook_deploy_key_created represents the GeneratedType_Webhook_deploy_key_created schema from the OpenAPI specification
type GeneratedType_Webhook_deploy_key_created struct {
	Key Webhooksdeploykey `json:"key"` // The [`deploy key`](https://docs.github.com/rest/deploy-keys/deploy-keys#get-a-deploy-key) resource.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_installation_created represents the GeneratedType_Webhook_installation_created schema from the OpenAPI specification
type GeneratedType_Webhook_installation_created struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester Webhooksuser `json:"requester,omitempty"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_issues_reopened represents the GeneratedType_Webhook_issues_reopened schema from the OpenAPI specification
type GeneratedType_Webhook_issues_reopened struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Pull_request_webhook represents the GeneratedType_Pull_request_webhook schema from the OpenAPI specification
type GeneratedType_Pull_request_webhook struct {
	Links map[string]interface{} `json:"_links"`
	Merged bool `json:"merged"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Commits int `json:"commits"`
	Updated_at string `json:"updated_at"`
	User GeneratedType_Simple_user `json:"user"` // A GitHub user.
	Mergeable bool `json:"mergeable"`
	Title string `json:"title"` // The title of the pull request.
	Changed_files int `json:"changed_files"`
	Deletions int `json:"deletions"`
	Mergeable_state string `json:"mergeable_state"`
	Number int `json:"number"` // Number uniquely identifying the pull request within its repository.
	Diff_url string `json:"diff_url"`
	Additions int `json:"additions"`
	Locked bool `json:"locked"`
	Requested_reviewers []GeneratedType_Simple_user `json:"requested_reviewers,omitempty"`
	Head map[string]interface{} `json:"head"`
	Maintainer_can_modify bool `json:"maintainer_can_modify"` // Indicates whether maintainers can modify the pull request.
	Body string `json:"body"`
	Requested_teams []GeneratedType_Team_simple `json:"requested_teams,omitempty"`
	Labels []map[string]interface{} `json:"labels"`
	Auto_merge GeneratedType_Auto_merge `json:"auto_merge"` // The status of auto merging a pull request.
	Closed_at string `json:"closed_at"`
	Patch_url string `json:"patch_url"`
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	Merged_by GeneratedType_Nullable_simple_user `json:"merged_by"` // A GitHub user.
	Review_comment_url string `json:"review_comment_url"`
	Comments int `json:"comments"`
	Merged_at string `json:"merged_at"`
	Id int64 `json:"id"`
	Assignee GeneratedType_Nullable_simple_user `json:"assignee"` // A GitHub user.
	Statuses_url string `json:"statuses_url"`
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Rebaseable bool `json:"rebaseable,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	State string `json:"state"` // State of this Pull Request. Either `open` or `closed`.
	Review_comments_url string `json:"review_comments_url"`
	Base map[string]interface{} `json:"base"`
	Review_comments int `json:"review_comments"`
	Commits_url string `json:"commits_url"`
	Assignees []GeneratedType_Simple_user `json:"assignees,omitempty"`
	Comments_url string `json:"comments_url"`
	Milestone GeneratedType_Nullable_milestone `json:"milestone"` // A collection of related issues and pull requests.
	Node_id string `json:"node_id"`
	Issue_url string `json:"issue_url"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Url string `json:"url"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether to allow updating the pull request's branch.
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged.
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., "Merge pull request #123 from branch-name").
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.**
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow auto-merge for pull requests.
}

// Webhookslabel represents the Webhookslabel schema from the OpenAPI specification
type Webhookslabel struct {
	Name string `json:"name"` // The name of the label.
	Node_id string `json:"node_id"`
	Url string `json:"url"` // URL for the label
	Color string `json:"color"` // 6-character hex code, without the leading #, identifying the color
	DefaultField bool `json:"default"`
	Description string `json:"description"`
	Id int `json:"id"`
}

// GeneratedType_Dependabot_secret represents the GeneratedType_Dependabot_secret schema from the OpenAPI specification
type GeneratedType_Dependabot_secret struct {
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Updated_at string `json:"updated_at"`
}

// GeneratedType_Workflow_run represents the GeneratedType_Workflow_run schema from the OpenAPI specification
type GeneratedType_Workflow_run struct {
	Check_suite_id int `json:"check_suite_id,omitempty"` // The ID of the associated check suite.
	Artifacts_url string `json:"artifacts_url"` // The URL to the artifacts for the workflow run.
	Run_started_at string `json:"run_started_at,omitempty"` // The start time of the latest run. Resets on re-run.
	Url string `json:"url"` // The URL to the workflow run.
	Triggering_actor GeneratedType_Simple_user `json:"triggering_actor,omitempty"` // A GitHub user.
	Repository GeneratedType_Minimal_repository `json:"repository"` // Minimal Repository
	Display_title string `json:"display_title"` // The event-specific title associated with the run or the run-name if set, or the value of `run-name` if it is set in the workflow.
	Logs_url string `json:"logs_url"` // The URL to download the logs for the workflow run.
	Head_repository GeneratedType_Minimal_repository `json:"head_repository"` // Minimal Repository
	Jobs_url string `json:"jobs_url"` // The URL to the jobs for the workflow run.
	Path string `json:"path"` // The full path of the workflow
	Workflow_url string `json:"workflow_url"` // The URL to the workflow.
	Event string `json:"event"`
	Rerun_url string `json:"rerun_url"` // The URL to rerun the workflow run.
	Previous_attempt_url string `json:"previous_attempt_url,omitempty"` // The URL to the previous attempted run of this workflow, if one exists.
	Head_repository_id int `json:"head_repository_id,omitempty"`
	Node_id string `json:"node_id"`
	Cancel_url string `json:"cancel_url"` // The URL to cancel the workflow run.
	Status string `json:"status"`
	Pull_requests []GeneratedType_Pull_request_minimal `json:"pull_requests"` // Pull requests that are open with a `head_sha` or `head_branch` that matches the workflow run. The returned pull requests do not necessarily indicate pull requests that triggered the run.
	Referenced_workflows []GeneratedType_Referenced_workflow `json:"referenced_workflows,omitempty"`
	Run_number int `json:"run_number"` // The auto incrementing run number for the workflow run.
	Conclusion string `json:"conclusion"`
	Head_commit GeneratedType_Nullable_simple_commit `json:"head_commit"` // A commit.
	Head_sha string `json:"head_sha"` // The SHA of the head commit that points to the version of the workflow being run.
	Created_at string `json:"created_at"`
	Check_suite_node_id string `json:"check_suite_node_id,omitempty"` // The node ID of the associated check suite.
	Id int `json:"id"` // The ID of the workflow run.
	Updated_at string `json:"updated_at"`
	Workflow_id int `json:"workflow_id"` // The ID of the parent workflow.
	Actor GeneratedType_Simple_user `json:"actor,omitempty"` // A GitHub user.
	Head_branch string `json:"head_branch"`
	Html_url string `json:"html_url"`
	Check_suite_url string `json:"check_suite_url"` // The URL to the associated check suite.
	Name string `json:"name,omitempty"` // The name of the workflow run.
	Run_attempt int `json:"run_attempt,omitempty"` // Attempt number of the run, 1 for first attempt and higher if the workflow was re-run.
}

// GeneratedType_Webhook_pull_request_labeled represents the GeneratedType_Webhook_pull_request_labeled schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_labeled struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Number int `json:"number"` // The pull request number.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Label Webhookslabel `json:"label,omitempty"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
}

// GeneratedType_Deployment_protection_rule represents the GeneratedType_Deployment_protection_rule schema from the OpenAPI specification
type GeneratedType_Deployment_protection_rule struct {
	Node_id string `json:"node_id"` // The node ID for the deployment protection rule.
	App GeneratedType_Custom_deployment_rule_app `json:"app"` // A GitHub App that is providing a custom deployment protection rule.
	Enabled bool `json:"enabled"` // Whether the deployment protection rule is enabled for the environment.
	Id int `json:"id"` // The unique identifier for the deployment protection rule.
}

// GeneratedType_Tag_protection represents the GeneratedType_Tag_protection schema from the OpenAPI specification
type GeneratedType_Tag_protection struct {
	Id int `json:"id,omitempty"`
	Pattern string `json:"pattern"`
	Updated_at string `json:"updated_at,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
}

// GeneratedType_Codespaces_secret represents the GeneratedType_Codespaces_secret schema from the OpenAPI specification
type GeneratedType_Codespaces_secret struct {
	Created_at string `json:"created_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the secret
	Selected_repositories_url string `json:"selected_repositories_url"` // The API URL at which the list of repositories this secret is visible to can be retrieved
	Updated_at string `json:"updated_at"` // The date and time at which the secret was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Visibility string `json:"visibility"` // The type of repositories in the organization that the secret is visible to
}

// GeneratedType_Webhook_issues_opened represents the GeneratedType_Webhook_issues_opened schema from the OpenAPI specification
type GeneratedType_Webhook_issues_opened struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_label_edited represents the GeneratedType_Webhook_label_edited schema from the OpenAPI specification
type GeneratedType_Webhook_label_edited struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // The changes to the label if the action was `edited`.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Page_build_status represents the GeneratedType_Page_build_status schema from the OpenAPI specification
type GeneratedType_Page_build_status struct {
	Status string `json:"status"`
	Url string `json:"url"`
}

// GeneratedType_Webhook_organization_renamed represents the GeneratedType_Webhook_organization_renamed schema from the OpenAPI specification
type GeneratedType_Webhook_organization_renamed struct {
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Membership Webhooksmembership `json:"membership,omitempty"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_personal_access_token_request_cancelled represents the GeneratedType_Webhook_personal_access_token_request_cancelled schema from the OpenAPI specification
type GeneratedType_Webhook_personal_access_token_request_cancelled struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Personal_access_token_request GeneratedType_Personal_access_token_request `json:"personal_access_token_request"` // Details of a Personal Access Token Request.
}

// GeneratedType_Nullable_actions_hosted_runner_pool_image represents the GeneratedType_Nullable_actions_hosted_runner_pool_image schema from the OpenAPI specification
type GeneratedType_Nullable_actions_hosted_runner_pool_image struct {
	Size_gb int `json:"size_gb"` // Image size in GB.
	Source string `json:"source"` // The image provider.
	Display_name string `json:"display_name"` // Display name for this image.
	Id string `json:"id"` // The ID of the image. Use this ID for the `image` parameter when creating a new larger runner.
}

// GeneratedType_Webhook_check_run_completed represents the GeneratedType_Webhook_check_run_completed schema from the OpenAPI specification
type GeneratedType_Webhook_check_run_completed struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
	Check_run GeneratedType_Check_run_with_simple_check_suite `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Branch_with_protection represents the GeneratedType_Branch_with_protection schema from the OpenAPI specification
type GeneratedType_Branch_with_protection struct {
	Pattern string `json:"pattern,omitempty"`
	Protected bool `json:"protected"`
	Protection GeneratedType_Branch_protection `json:"protection"` // Branch Protection
	Protection_url string `json:"protection_url"`
	Required_approving_review_count int `json:"required_approving_review_count,omitempty"`
	Links map[string]interface{} `json:"_links"`
	Commit Commit `json:"commit"` // Commit
	Name string `json:"name"`
}

// Webhookspreviousmarketplacepurchase represents the Webhookspreviousmarketplacepurchase schema from the OpenAPI specification
type Webhookspreviousmarketplacepurchase struct {
	Plan map[string]interface{} `json:"plan"`
	Unit_count int `json:"unit_count"`
	Account map[string]interface{} `json:"account"`
	Billing_cycle string `json:"billing_cycle"`
	Free_trial_ends_on interface{} `json:"free_trial_ends_on"`
	Next_billing_date string `json:"next_billing_date,omitempty"`
	On_free_trial bool `json:"on_free_trial"`
}

// Enterprise represents the Enterprise schema from the OpenAPI specification
type Enterprise struct {
	Avatar_url string `json:"avatar_url"`
	Created_at string `json:"created_at"`
	Description string `json:"description,omitempty"` // A short description of the enterprise.
	Name string `json:"name"` // The name of the enterprise.
	Updated_at string `json:"updated_at"`
	Html_url string `json:"html_url"`
	Slug string `json:"slug"` // The slug url identifier for the enterprise.
	Website_url string `json:"website_url,omitempty"` // The enterprise's website URL.
	Id int `json:"id"` // Unique identifier of the enterprise
	Node_id string `json:"node_id"`
}

// GeneratedType_Webhook_security_advisory_updated represents the GeneratedType_Webhook_security_advisory_updated schema from the OpenAPI specification
type GeneratedType_Webhook_security_advisory_updated struct {
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Security_advisory Webhookssecurityadvisory `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
}

// Classroom represents the Classroom schema from the OpenAPI specification
type Classroom struct {
	Archived bool `json:"archived"` // Whether classroom is archived.
	Id int `json:"id"` // Unique identifier of the classroom.
	Name string `json:"name"` // The name of the classroom.
	Organization GeneratedType_Simple_classroom_organization `json:"organization"` // A GitHub organization.
	Url string `json:"url"` // The URL of the classroom on GitHub Classroom.
}

// GeneratedType_Repository_advisory_update represents the GeneratedType_Repository_advisory_update schema from the OpenAPI specification
type GeneratedType_Repository_advisory_update struct {
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities,omitempty"` // A product affected by the vulnerability detailed in a repository security advisory.
	Cwe_ids []string `json:"cwe_ids,omitempty"` // A list of Common Weakness Enumeration (CWE) IDs.
	Collaborating_users []string `json:"collaborating_users,omitempty"` // A list of usernames who have been granted write access to the advisory.
	Cvss_vector_string string `json:"cvss_vector_string,omitempty"` // The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
	Severity string `json:"severity,omitempty"` // The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
	Summary string `json:"summary,omitempty"` // A short summary of the advisory.
	Collaborating_teams []string `json:"collaborating_teams,omitempty"` // A list of team slugs which have been granted write access to the advisory.
	Cve_id string `json:"cve_id,omitempty"` // The Common Vulnerabilities and Exposures (CVE) ID.
	Description string `json:"description,omitempty"` // A detailed description of what the advisory impacts.
	State string `json:"state,omitempty"` // The state of the advisory.
	Credits []map[string]interface{} `json:"credits,omitempty"` // A list of users receiving credit for their participation in the security advisory.
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Tarball_url string `json:"tarball_url"`
	Zipball_url string `json:"zipball_url"`
	Commit map[string]interface{} `json:"commit"`
}

// GeneratedType_Webhook_commit_comment_created represents the GeneratedType_Webhook_commit_comment_created schema from the OpenAPI specification
type GeneratedType_Webhook_commit_comment_created struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action performed. Can be `created`.
	Comment map[string]interface{} `json:"comment"` // The [commit comment](${externalDocsUpapp/api/description/components/schemas/webhooks/issue-comment-created.yamlrl}/rest/commits/comments#get-a-commit-comment) resource.
}

// GeneratedType_Short_blob represents the GeneratedType_Short_blob schema from the OpenAPI specification
type GeneratedType_Short_blob struct {
	Url string `json:"url"`
	Sha string `json:"sha"`
}

// GeneratedType_Webhook_discussion_closed represents the GeneratedType_Webhook_discussion_closed schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_closed struct {
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_repository_transferred represents the GeneratedType_Webhook_repository_transferred schema from the OpenAPI specification
type GeneratedType_Webhook_repository_transferred struct {
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Selected_actions represents the GeneratedType_Selected_actions schema from the OpenAPI specification
type GeneratedType_Selected_actions struct {
	Github_owned_allowed bool `json:"github_owned_allowed,omitempty"` // Whether GitHub-owned actions are allowed. For example, this includes the actions in the `actions` organization.
	Patterns_allowed []string `json:"patterns_allowed,omitempty"` // Specifies a list of string-matching patterns to allow specific action(s) and reusable workflow(s). Wildcards, tags, and SHAs are allowed. For example, `monalisa/octocat@*`, `monalisa/octocat@v2`, `monalisa/*`. > [!NOTE] > The `patterns_allowed` setting only applies to public repositories.
	Verified_allowed bool `json:"verified_allowed,omitempty"` // Whether actions from GitHub Marketplace verified creators are allowed. Set to `true` to allow all actions by GitHub Marketplace verified creators.
}

// GeneratedType_Locked_issue_event represents the GeneratedType_Locked_issue_event schema from the OpenAPI specification
type GeneratedType_Locked_issue_event struct {
	Node_id string `json:"node_id"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_url string `json:"commit_url"`
	Lock_reason string `json:"lock_reason"`
}

// GeneratedType_Webhook_projects_v2_item_created represents the GeneratedType_Webhook_projects_v2_item_created schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_item_created struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType_Projects_v2_item `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Workflow represents the Workflow schema from the OpenAPI specification
type Workflow struct {
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Name string `json:"name"`
	State string `json:"state"`
	Deleted_at string `json:"deleted_at,omitempty"`
	Html_url string `json:"html_url"`
	Path string `json:"path"`
	Badge_url string `json:"badge_url"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType_Code_scanning_default_setup_update represents the GeneratedType_Code_scanning_default_setup_update schema from the OpenAPI specification
type GeneratedType_Code_scanning_default_setup_update struct {
	State string `json:"state,omitempty"` // The desired state of code scanning default setup.
	Threat_model string `json:"threat_model,omitempty"` // Threat model to be used for code scanning analysis. Use `remote` to analyze only network sources and `remote_and_local` to include local sources like filesystem access, command-line arguments, database reads, environment variable and standard input.
	Languages []string `json:"languages,omitempty"` // CodeQL languages to be analyzed.
	Query_suite string `json:"query_suite,omitempty"` // CodeQL query suite to be used.
	Runner_label string `json:"runner_label,omitempty"` // Runner label to be used if the runner type is labeled.
	Runner_type string `json:"runner_type,omitempty"` // Runner type to be used.
}

// GeneratedType_Webhook_project_card_created represents the GeneratedType_Webhook_project_card_created schema from the OpenAPI specification
type GeneratedType_Webhook_project_card_created struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card Webhooksprojectcard `json:"project_card"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Code_scanning_alert_location represents the GeneratedType_Code_scanning_alert_location schema from the OpenAPI specification
type GeneratedType_Code_scanning_alert_location struct {
	End_line int `json:"end_line,omitempty"`
	Path string `json:"path,omitempty"`
	Start_column int `json:"start_column,omitempty"`
	Start_line int `json:"start_line,omitempty"`
	End_column int `json:"end_column,omitempty"`
}

// GeneratedType_Labeled_issue_event represents the GeneratedType_Labeled_issue_event schema from the OpenAPI specification
type GeneratedType_Labeled_issue_event struct {
	Node_id string `json:"node_id"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Label map[string]interface{} `json:"label"`
	Commit_id string `json:"commit_id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Event string `json:"event"`
}

// GeneratedType_Page_build represents the GeneratedType_Page_build schema from the OpenAPI specification
type GeneratedType_Page_build struct {
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Commit string `json:"commit"`
	Created_at string `json:"created_at"`
	Duration int `json:"duration"`
	ErrorField map[string]interface{} `json:"error"`
	Pusher GeneratedType_Nullable_simple_user `json:"pusher"` // A GitHub user.
	Status string `json:"status"`
}

// GeneratedType_Webhook_secret_scanning_alert_location_created represents the GeneratedType_Webhook_secret_scanning_alert_location_created schema from the OpenAPI specification
type GeneratedType_Webhook_secret_scanning_alert_location_created struct {
	Action string `json:"action,omitempty"`
	Alert GeneratedType_Secret_scanning_alert_webhook `json:"alert"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Location GeneratedType_Secret_scanning_location `json:"location"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_discussion_transferred represents the GeneratedType_Webhook_discussion_transferred schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_transferred struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Personal_access_token_request represents the GeneratedType_Personal_access_token_request schema from the OpenAPI specification
type GeneratedType_Personal_access_token_request struct {
	Permissions_added map[string]interface{} `json:"permissions_added"` // New requested permissions, categorized by type of permission.
	Repository_count int `json:"repository_count"` // The number of repositories the token is requesting access to. This field is only populated when `repository_selection` is `subset`.
	Token_expired bool `json:"token_expired"` // Whether the associated fine-grained personal access token has expired.
	Created_at string `json:"created_at"` // Date and time when the request for access was created.
	Repositories []map[string]interface{} `json:"repositories"` // An array of repository objects the token is requesting access to. This field is only populated when `repository_selection` is `subset`.
	Token_expires_at string `json:"token_expires_at"` // Date and time when the associated fine-grained personal access token expires.
	Token_last_used_at string `json:"token_last_used_at"` // Date and time when the associated fine-grained personal access token was last used for authentication.
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Token_name string `json:"token_name"` // The name given to the user's token. This field can also be found in an organization's settings page for Active Tokens.
	Permissions_result map[string]interface{} `json:"permissions_result"` // Permissions requested, categorized by type of permission. This field incorporates `permissions_added` and `permissions_upgraded`.
	Permissions_upgraded map[string]interface{} `json:"permissions_upgraded"` // Requested permissions that elevate access for a previously approved request for access, categorized by type of permission.
	Repository_selection string `json:"repository_selection"` // Type of repository selection requested.
	Token_id int `json:"token_id"` // Unique identifier of the user's token. This field can also be found in audit log events and the organization's settings for their PAT grants.
	Id int `json:"id"` // Unique identifier of the request for access via fine-grained personal access token. Used as the `pat_request_id` parameter in the list and review API calls.
}

// GeneratedType_Validation_error_simple represents the GeneratedType_Validation_error_simple schema from the OpenAPI specification
type GeneratedType_Validation_error_simple struct {
	Documentation_url string `json:"documentation_url"`
	Errors []string `json:"errors,omitempty"`
	Message string `json:"message"`
}

// GeneratedType_Webhook_repository_import represents the GeneratedType_Webhook_repository_import schema from the OpenAPI specification
type GeneratedType_Webhook_repository_import struct {
	Status string `json:"status"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// Actor represents the Actor schema from the OpenAPI specification
type Actor struct {
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Display_login string `json:"display_login,omitempty"`
	Gravatar_id string `json:"gravatar_id"`
	Id int `json:"id"`
	Login string `json:"login"`
}

// GeneratedType_Repository_rule_commit_message_pattern represents the GeneratedType_Repository_rule_commit_message_pattern schema from the OpenAPI specification
type GeneratedType_Repository_rule_commit_message_pattern struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType_Webhook_team_edited represents the GeneratedType_Webhook_team_edited schema from the OpenAPI specification
type GeneratedType_Webhook_team_edited struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the team if the action was `edited`.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_custom_property_values_updated represents the GeneratedType_Webhook_custom_property_values_updated schema from the OpenAPI specification
type GeneratedType_Webhook_custom_property_values_updated struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	New_property_values []GeneratedType_Custom_property_value `json:"new_property_values"` // The new custom property values for the repository.
	Old_property_values []GeneratedType_Custom_property_value `json:"old_property_values"` // The old custom property values for the repository.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_release_unpublished represents the GeneratedType_Webhook_release_unpublished schema from the OpenAPI specification
type GeneratedType_Webhook_release_unpublished struct {
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease1 `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Gist_comment represents the GeneratedType_Gist_comment schema from the OpenAPI specification
type GeneratedType_Gist_comment struct {
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"` // The comment text.
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
}

// GeneratedType_Copilot_dotcom_pull_requests represents the GeneratedType_Copilot_dotcom_pull_requests schema from the OpenAPI specification
type GeneratedType_Copilot_dotcom_pull_requests struct {
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // The number of users who used Copilot for Pull Requests on github.com to generate a pull request summary at least once.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // Repositories in which users used Copilot for Pull Requests to generate pull request summaries
}

// GeneratedType_Webhook_dependabot_alert_reintroduced represents the GeneratedType_Webhook_dependabot_alert_reintroduced schema from the OpenAPI specification
type GeneratedType_Webhook_dependabot_alert_reintroduced struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType_Dependabot_alert `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_License_content represents the GeneratedType_License_content schema from the OpenAPI specification
type GeneratedType_License_content struct {
	Name string `json:"name"`
	Size int `json:"size"`
	Path string `json:"path"`
	Url string `json:"url"`
	Content string `json:"content"`
	Git_url string `json:"git_url"`
	Html_url string `json:"html_url"`
	License GeneratedType_Nullable_license_simple `json:"license"` // License Simple
	Sha string `json:"sha"`
	TypeField string `json:"type"`
	Links map[string]interface{} `json:"_links"`
	Download_url string `json:"download_url"`
	Encoding string `json:"encoding"`
}

// GeneratedType_Timeline_issue_events represents the GeneratedType_Timeline_issue_events schema from the OpenAPI specification
type GeneratedType_Timeline_issue_events struct {
}

// GeneratedType_Projects_v2_iteration_setting represents the GeneratedType_Projects_v2_iteration_setting schema from the OpenAPI specification
type GeneratedType_Projects_v2_iteration_setting struct {
	Duration float64 `json:"duration,omitempty"`
	Id string `json:"id"`
	Start_date string `json:"start_date,omitempty"`
	Title string `json:"title"`
}

// GeneratedType_Secret_scanning_location_issue_title represents the GeneratedType_Secret_scanning_location_issue_title schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_issue_title struct {
	Issue_title_url string `json:"issue_title_url"` // The API URL to get the issue where the secret was detected.
}

// GeneratedType_Webhook_check_run_completed_form_encoded represents the GeneratedType_Webhook_check_run_completed_form_encoded schema from the OpenAPI specification
type GeneratedType_Webhook_check_run_completed_form_encoded struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.completed JSON payload. The decoded payload is a JSON object.
}

// GeneratedType_Webhook_security_advisory_withdrawn represents the GeneratedType_Webhook_security_advisory_withdrawn schema from the OpenAPI specification
type GeneratedType_Webhook_security_advisory_withdrawn struct {
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Security_advisory map[string]interface{} `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
}

// GeneratedType_Webhook_issues_untyped represents the GeneratedType_Webhook_issues_untyped schema from the OpenAPI specification
type GeneratedType_Webhook_issues_untyped struct {
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	TypeField GeneratedType_Issue_type `json:"type"` // The type of issue.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_projects_v2_item_deleted represents the GeneratedType_Webhook_projects_v2_item_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_item_deleted struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType_Projects_v2_item `json:"projects_v2_item"` // An item belonging to a project
}

// GeneratedType_Nullable_git_user represents the GeneratedType_Nullable_git_user schema from the OpenAPI specification
type GeneratedType_Nullable_git_user struct {
	Date string `json:"date,omitempty"`
	Email string `json:"email,omitempty"`
	Name string `json:"name,omitempty"`
}

// GeneratedType_Repository_rule_code_scanning represents the GeneratedType_Repository_rule_code_scanning schema from the OpenAPI specification
type GeneratedType_Repository_rule_code_scanning struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType_Actions_workflow_access_to_repository represents the GeneratedType_Actions_workflow_access_to_repository schema from the OpenAPI specification
type GeneratedType_Actions_workflow_access_to_repository struct {
	Access_level string `json:"access_level"` // Defines the level of access that workflows outside of the repository have to actions and reusable workflows within the repository. `none` means the access is only possible from workflows in this repository. `user` level access allows sharing across user owned private repositories only. `organization` level access allows sharing across the organization.
}

// GeneratedType_Org_private_registry_configuration_with_selected_repositories represents the GeneratedType_Org_private_registry_configuration_with_selected_repositories schema from the OpenAPI specification
type GeneratedType_Org_private_registry_configuration_with_selected_repositories struct {
	Selected_repository_ids []int `json:"selected_repository_ids,omitempty"` // An array of repository IDs that can access the organization private registry when `visibility` is set to `selected`.
	Updated_at string `json:"updated_at"`
	Username string `json:"username,omitempty"` // The username to use when authenticating with the private registry.
	Visibility string `json:"visibility"` // Which type of organization repositories have access to the private registry. `selected` means only the repositories specified by `selected_repository_ids` can access the private registry.
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the private registry configuration.
	Registry_type string `json:"registry_type"` // The registry type.
}

// GeneratedType_Cvss_severities represents the GeneratedType_Cvss_severities schema from the OpenAPI specification
type GeneratedType_Cvss_severities struct {
	Cvss_v4 map[string]interface{} `json:"cvss_v4,omitempty"`
	Cvss_v3 map[string]interface{} `json:"cvss_v3,omitempty"`
}

// GeneratedType_Webhook_merge_group_checks_requested represents the GeneratedType_Webhook_merge_group_checks_requested schema from the OpenAPI specification
type GeneratedType_Webhook_merge_group_checks_requested struct {
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Merge_group GeneratedType_Merge_group `json:"merge_group"` // A group of pull requests that the merge queue has grouped together to be merged.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Combined_commit_status represents the GeneratedType_Combined_commit_status schema from the OpenAPI specification
type GeneratedType_Combined_commit_status struct {
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
	Repository GeneratedType_Minimal_repository `json:"repository"` // Minimal Repository
	Sha string `json:"sha"`
	State string `json:"state"`
	Statuses []GeneratedType_Simple_commit_status `json:"statuses"`
	Total_count int `json:"total_count"`
}

// GeneratedType_Webhook_repository_ruleset_deleted represents the GeneratedType_Webhook_repository_ruleset_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_repository_ruleset_deleted struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_ruleset GeneratedType_Repository_ruleset `json:"repository_ruleset"` // A set of rules to apply when specified conditions are met.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Code_scanning_alert_items represents the GeneratedType_Code_scanning_alert_items schema from the OpenAPI specification
type GeneratedType_Code_scanning_alert_items struct {
	State string `json:"state"` // State of a code scanning alert.
	Dismissed_by GeneratedType_Nullable_simple_user `json:"dismissed_by"` // A GitHub user.
	Number int `json:"number"` // The security alert number.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Rule GeneratedType_Code_scanning_alert_rule_summary `json:"rule"`
	Most_recent_instance GeneratedType_Code_scanning_alert_instance `json:"most_recent_instance"`
	Tool GeneratedType_Code_scanning_analysis_tool `json:"tool"`
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Dismissal_approved_by GeneratedType_Nullable_simple_user `json:"dismissal_approved_by,omitempty"` // A GitHub user.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType_Page_deployment represents the GeneratedType_Page_deployment schema from the OpenAPI specification
type GeneratedType_Page_deployment struct {
	Id string `json:"id"` // The ID of the GitHub Pages deployment. This is the Git SHA of the deployed commit.
	Page_url string `json:"page_url"` // The URI to the deployed GitHub Pages.
	Preview_url string `json:"preview_url,omitempty"` // The URI to the deployed GitHub Pages preview.
	Status_url string `json:"status_url"` // The URI to monitor GitHub Pages deployment status.
}

// GeneratedType_Code_scanning_autofix represents the GeneratedType_Code_scanning_autofix schema from the OpenAPI specification
type GeneratedType_Code_scanning_autofix struct {
	Description string `json:"description"` // The description of an autofix.
	Started_at string `json:"started_at"` // The start time of an autofix in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Status string `json:"status"` // The status of an autofix.
}

// Commit represents the Commit schema from the OpenAPI specification
type Commit struct {
	Committer interface{} `json:"committer"`
	Node_id string `json:"node_id"`
	Sha string `json:"sha"`
	Stats map[string]interface{} `json:"stats,omitempty"`
	Files []GeneratedType_Diff_entry `json:"files,omitempty"`
	Url string `json:"url"`
	Author interface{} `json:"author"`
	Comments_url string `json:"comments_url"`
	Commit map[string]interface{} `json:"commit"`
	Parents []map[string]interface{} `json:"parents"`
	Html_url string `json:"html_url"`
}

// GeneratedType_Org_ruleset_conditions represents the GeneratedType_Org_ruleset_conditions schema from the OpenAPI specification
type GeneratedType_Org_ruleset_conditions struct {
}

// GeneratedType_Secret_scanning_location_pull_request_body represents the GeneratedType_Secret_scanning_location_pull_request_body schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_pull_request_body struct {
	Pull_request_body_url string `json:"pull_request_body_url"` // The API URL to get the pull request where the secret was detected.
}

// GeneratedType_Webhook_team_removed_from_repository represents the GeneratedType_Webhook_team_removed_from_repository schema from the OpenAPI specification
type GeneratedType_Webhook_team_removed_from_repository struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_issues_pinned represents the GeneratedType_Webhook_issues_pinned schema from the OpenAPI specification
type GeneratedType_Webhook_issues_pinned struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue2 `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// Webhooksissue2 represents the Webhooksissue2 schema from the OpenAPI specification
type Webhooksissue2 struct {
	Assignees []map[string]interface{} `json:"assignees"`
	User map[string]interface{} `json:"user"`
	Html_url string `json:"html_url"`
	Updated_at string `json:"updated_at"`
	Sub_issues_summary map[string]interface{} `json:"sub_issues_summary,omitempty"`
	Id int64 `json:"id"`
	Labels []map[string]interface{} `json:"labels,omitempty"`
	Reactions map[string]interface{} `json:"reactions"`
	Created_at string `json:"created_at"`
	Timeline_url string `json:"timeline_url,omitempty"`
	State string `json:"state,omitempty"` // State of the issue; either 'open' or 'closed'
	Events_url string `json:"events_url"`
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Url string `json:"url"` // URL for the issue
	Number int `json:"number"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Performed_via_github_app map[string]interface{} `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Draft bool `json:"draft,omitempty"`
	Node_id string `json:"node_id"`
	Closed_at string `json:"closed_at"`
	Assignee map[string]interface{} `json:"assignee,omitempty"`
	Comments int `json:"comments"`
	Title string `json:"title"` // Title of the issue
	Locked bool `json:"locked,omitempty"`
	TypeField GeneratedType_Issue_type `json:"type,omitempty"` // The type of issue.
	State_reason string `json:"state_reason,omitempty"`
	Comments_url string `json:"comments_url"`
	Active_lock_reason string `json:"active_lock_reason"`
	Labels_url string `json:"labels_url"`
	Body string `json:"body"` // Contents of the issue
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Repository_url string `json:"repository_url"`
}

// GeneratedType_Security_advisory_epss represents the GeneratedType_Security_advisory_epss schema from the OpenAPI specification
type GeneratedType_Security_advisory_epss struct {
	Percentage float64 `json:"percentage,omitempty"`
	Percentile float64 `json:"percentile,omitempty"`
}

// GeneratedType_Key_simple represents the GeneratedType_Key_simple schema from the OpenAPI specification
type GeneratedType_Key_simple struct {
	Created_at string `json:"created_at,omitempty"`
	Id int `json:"id"`
	Key string `json:"key"`
}

// Deployment represents the Deployment schema from the OpenAPI specification
type Deployment struct {
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Id int64 `json:"id"` // Unique identifier of the deployment
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Repository_url string `json:"repository_url"`
	Url string `json:"url"`
	Original_environment string `json:"original_environment,omitempty"`
	Statuses_url string `json:"statuses_url"`
	Payload interface{} `json:"payload"`
	Task string `json:"task"` // Parameter to specify a task to execute
	Transient_environment bool `json:"transient_environment,omitempty"` // Specifies if the given environment is will no longer exist at some point in the future. Default: false.
	Environment string `json:"environment"` // Name for the target deployment environment.
	Production_environment bool `json:"production_environment,omitempty"` // Specifies if the given environment is one that end-users directly interact with. Default: false.
	Ref string `json:"ref"` // The ref to deploy. This can be a branch, tag, or sha.
	Creator GeneratedType_Nullable_simple_user `json:"creator"` // A GitHub user.
	Description string `json:"description"`
	Sha string `json:"sha"`
	Created_at string `json:"created_at"`
}

// GeneratedType_Webhook_star_deleted represents the GeneratedType_Webhook_star_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_star_deleted struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Starred_at interface{} `json:"starred_at"` // The time the star was created. This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Will be `null` for the `deleted` action.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_issues_labeled represents the GeneratedType_Webhook_issues_labeled schema from the OpenAPI specification
type GeneratedType_Webhook_issues_labeled struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Label Webhookslabel `json:"label,omitempty"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Organization_programmatic_access_grant_request represents the GeneratedType_Organization_programmatic_access_grant_request schema from the OpenAPI specification
type GeneratedType_Organization_programmatic_access_grant_request struct {
	Token_id int `json:"token_id"` // Unique identifier of the user's token. This field can also be found in audit log events and the organization's settings for their PAT grants.
	Token_name string `json:"token_name"` // The name given to the user's token. This field can also be found in an organization's settings page for Active Tokens.
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Permissions map[string]interface{} `json:"permissions"` // Permissions requested, categorized by type of permission.
	Id int `json:"id"` // Unique identifier of the request for access via fine-grained personal access token. The `pat_request_id` used to review PAT requests.
	Repositories_url string `json:"repositories_url"` // URL to the list of repositories requested to be accessed via fine-grained personal access token. Should only be followed when `repository_selection` is `subset`.
	Repository_selection string `json:"repository_selection"` // Type of repository selection requested.
	Token_expired bool `json:"token_expired"` // Whether the associated fine-grained personal access token has expired.
	Reason string `json:"reason"` // Reason for requesting access.
	Token_last_used_at string `json:"token_last_used_at"` // Date and time when the associated fine-grained personal access token was last used for authentication.
	Created_at string `json:"created_at"` // Date and time when the request for access was created.
	Token_expires_at string `json:"token_expires_at"` // Date and time when the associated fine-grained personal access token expires.
}

// GeneratedType_Nullable_team_simple represents the GeneratedType_Nullable_team_simple schema from the OpenAPI specification
type GeneratedType_Nullable_team_simple struct {
	Repositories_url string `json:"repositories_url"`
	Url string `json:"url"` // URL for the team
	Name string `json:"name"` // Name of the team
	Node_id string `json:"node_id"`
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Description string `json:"description"` // Description of the team
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the team
	Members_url string `json:"members_url"`
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Notification_setting string `json:"notification_setting,omitempty"` // The notification setting the team has set
	Slug string `json:"slug"`
}

// GeneratedType_Secret_scanning_location_discussion_comment represents the GeneratedType_Secret_scanning_location_discussion_comment schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_discussion_comment struct {
	Discussion_comment_url string `json:"discussion_comment_url"` // The API URL to get the discussion comment where the secret was detected.
}

// GeneratedType_Webhook_issue_comment_deleted represents the GeneratedType_Webhook_issue_comment_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_issue_comment_deleted struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhooksissuecomment `json:"comment"` // The [comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment) itself.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) the comment belongs to.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// Hook represents the Hook schema from the OpenAPI specification
type Hook struct {
	Updated_at string `json:"updated_at"`
	Last_response GeneratedType_Hook_response `json:"last_response"`
	Url string `json:"url"`
	Events []string `json:"events"` // Determines what events the hook is triggered for. Default: ['push'].
	Config GeneratedType_Webhook_config `json:"config"` // Configuration object of the webhook
	Id int `json:"id"` // Unique identifier of the webhook.
	Ping_url string `json:"ping_url"`
	Active bool `json:"active"` // Determines whether the hook is actually triggered on pushes.
	Created_at string `json:"created_at"`
	Deliveries_url string `json:"deliveries_url,omitempty"`
	Name string `json:"name"` // The name of a valid service, use 'web' for a webhook.
	Test_url string `json:"test_url"`
	TypeField string `json:"type"`
}

// GeneratedType_Network_configuration represents the GeneratedType_Network_configuration schema from the OpenAPI specification
type GeneratedType_Network_configuration struct {
	Compute_service string `json:"compute_service,omitempty"` // The hosted compute service the network configuration supports.
	Created_on string `json:"created_on"` // The time at which the network configuration was created, in ISO 8601 format.
	Id string `json:"id"` // The unique identifier of the network configuration.
	Name string `json:"name"` // The name of the network configuration.
	Network_settings_ids []string `json:"network_settings_ids,omitempty"` // The unique identifier of each network settings in the configuration.
}

// Webhooksissue represents the Webhooksissue schema from the OpenAPI specification
type Webhooksissue struct {
	Locked bool `json:"locked,omitempty"`
	Node_id string `json:"node_id"`
	Comments_url string `json:"comments_url"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Updated_at string `json:"updated_at"`
	Reactions map[string]interface{} `json:"reactions"`
	TypeField GeneratedType_Issue_type `json:"type,omitempty"` // The type of issue.
	Draft bool `json:"draft,omitempty"`
	Performed_via_github_app map[string]interface{} `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Number int `json:"number"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Created_at string `json:"created_at"`
	State_reason string `json:"state_reason,omitempty"`
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	User map[string]interface{} `json:"user"`
	Repository_url string `json:"repository_url"`
	Labels []map[string]interface{} `json:"labels,omitempty"`
	State string `json:"state,omitempty"` // State of the issue; either 'open' or 'closed'
	Comments int `json:"comments"`
	Url string `json:"url"` // URL for the issue
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Labels_url string `json:"labels_url"`
	Html_url string `json:"html_url"`
	Closed_at string `json:"closed_at"`
	Events_url string `json:"events_url"`
	Active_lock_reason string `json:"active_lock_reason"`
	Id int64 `json:"id"`
	Body string `json:"body"` // Contents of the issue
	Title string `json:"title"` // Title of the issue
	Assignees []map[string]interface{} `json:"assignees"`
	Sub_issues_summary map[string]interface{} `json:"sub_issues_summary,omitempty"`
	Assignee map[string]interface{} `json:"assignee,omitempty"`
}

// GeneratedType_Webhook_release_released represents the GeneratedType_Webhook_release_released schema from the OpenAPI specification
type GeneratedType_Webhook_release_released struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Organization_full represents the GeneratedType_Organization_full schema from the OpenAPI specification
type GeneratedType_Organization_full struct {
	Members_can_delete_issues bool `json:"members_can_delete_issues,omitempty"`
	Members_can_create_internal_repositories bool `json:"members_can_create_internal_repositories,omitempty"`
	Public_members_url string `json:"public_members_url"`
	Members_can_fork_private_repositories bool `json:"members_can_fork_private_repositories,omitempty"`
	Location string `json:"location,omitempty"`
	Created_at string `json:"created_at"`
	Secret_scanning_enabled_for_new_repositories bool `json:"secret_scanning_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether secret scanning is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Repos_url string `json:"repos_url"`
	Two_factor_requirement_enabled bool `json:"two_factor_requirement_enabled,omitempty"`
	Default_repository_branch string `json:"default_repository_branch,omitempty"` // The default branch for repositories created in this organization.
	Billing_email string `json:"billing_email,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Members_can_create_public_repositories bool `json:"members_can_create_public_repositories,omitempty"`
	Html_url string `json:"html_url"`
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Events_url string `json:"events_url"`
	Members_can_create_private_repositories bool `json:"members_can_create_private_repositories,omitempty"`
	Members_allowed_repository_creation_type string `json:"members_allowed_repository_creation_type,omitempty"`
	Secret_scanning_push_protection_custom_link_enabled bool `json:"secret_scanning_push_protection_custom_link_enabled,omitempty"` // Whether a custom link is shown to contributors who are blocked from pushing a secret by push protection.
	Blog string `json:"blog,omitempty"`
	Public_repos int `json:"public_repos"`
	Public_gists int `json:"public_gists"`
	TypeField string `json:"type"`
	Deploy_keys_enabled_for_repositories bool `json:"deploy_keys_enabled_for_repositories,omitempty"` // Controls whether or not deploy keys may be added and used for repositories in the organization.
	Id int `json:"id"`
	Updated_at string `json:"updated_at"`
	Archived_at string `json:"archived_at"`
	Dependabot_security_updates_enabled_for_new_repositories bool `json:"dependabot_security_updates_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether Dependabot security updates are automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Default_repository_permission string `json:"default_repository_permission,omitempty"`
	Node_id string `json:"node_id"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Hooks_url string `json:"hooks_url"`
	Is_verified bool `json:"is_verified,omitempty"`
	Has_organization_projects bool `json:"has_organization_projects"`
	Private_gists int `json:"private_gists,omitempty"`
	Members_can_create_teams bool `json:"members_can_create_teams,omitempty"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Secret_scanning_push_protection_enabled_for_new_repositories bool `json:"secret_scanning_push_protection_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether secret scanning push protection is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Issues_url string `json:"issues_url"`
	Followers int `json:"followers"`
	Display_commenter_full_name_setting_enabled bool `json:"display_commenter_full_name_setting_enabled,omitempty"`
	Company string `json:"company,omitempty"`
	Url string `json:"url"`
	Dependabot_alerts_enabled_for_new_repositories bool `json:"dependabot_alerts_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether Dependabot alerts are automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Members_can_change_repo_visibility bool `json:"members_can_change_repo_visibility,omitempty"`
	Name string `json:"name,omitempty"`
	Dependency_graph_enabled_for_new_repositories bool `json:"dependency_graph_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether dependency graph is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Following int `json:"following"`
	Members_can_view_dependency_insights bool `json:"members_can_view_dependency_insights,omitempty"`
	Members_can_create_private_pages bool `json:"members_can_create_private_pages,omitempty"`
	Readers_can_create_discussions bool `json:"readers_can_create_discussions,omitempty"`
	Email string `json:"email,omitempty"`
	Advanced_security_enabled_for_new_repositories bool `json:"advanced_security_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether GitHub Advanced Security is enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Plan map[string]interface{} `json:"plan,omitempty"`
	Members_can_invite_outside_collaborators bool `json:"members_can_invite_outside_collaborators,omitempty"`
	Members_url string `json:"members_url"`
	Collaborators int `json:"collaborators,omitempty"` // The number of collaborators on private repositories. This field may be null if the number of private repositories is over 50,000.
	Description string `json:"description"`
	Login string `json:"login"`
	Members_can_create_repositories bool `json:"members_can_create_repositories,omitempty"`
	Members_can_create_public_pages bool `json:"members_can_create_public_pages,omitempty"`
	Has_repository_projects bool `json:"has_repository_projects"`
	Members_can_delete_repositories bool `json:"members_can_delete_repositories,omitempty"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Members_can_create_pages bool `json:"members_can_create_pages,omitempty"`
	Secret_scanning_push_protection_custom_link string `json:"secret_scanning_push_protection_custom_link,omitempty"` // An optional URL string to display to contributors who are blocked from pushing a secret.
}

// GeneratedType_Demilestoned_issue_event represents the GeneratedType_Demilestoned_issue_event schema from the OpenAPI specification
type GeneratedType_Demilestoned_issue_event struct {
	Milestone map[string]interface{} `json:"milestone"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Commit_id string `json:"commit_id"`
	Id int `json:"id"`
	Url string `json:"url"`
	Event string `json:"event"`
}

// Codespace represents the Codespace schema from the OpenAPI specification
type Codespace struct {
	Web_url string `json:"web_url"` // URL to access this codespace on the web.
	Idle_timeout_notice string `json:"idle_timeout_notice,omitempty"` // Text to show user when codespace idle timeout minutes has been overriden by an organization policy
	Devcontainer_path string `json:"devcontainer_path,omitempty"` // Path to devcontainer.json from repo root used to create Codespace.
	Start_url string `json:"start_url"` // API URL to start this codespace.
	Last_known_stop_notice string `json:"last_known_stop_notice,omitempty"` // The text to display to a user when a codespace has been stopped for a potentially actionable reason.
	Location string `json:"location"` // The initally assigned location of a new codespace.
	Display_name string `json:"display_name,omitempty"` // Display name for this codespace.
	Id int64 `json:"id"`
	Billable_owner GeneratedType_Simple_user `json:"billable_owner"` // A GitHub user.
	Url string `json:"url"` // API URL for this codespace.
	Idle_timeout_minutes int `json:"idle_timeout_minutes"` // The number of minutes of inactivity after which this codespace will be automatically stopped.
	Machines_url string `json:"machines_url"` // API URL to access available alternate machine types for this codespace.
	Pending_operation_disabled_reason string `json:"pending_operation_disabled_reason,omitempty"` // Text to show user when codespace is disabled by a pending operation
	Prebuild bool `json:"prebuild"` // Whether the codespace was created from a prebuild.
	Publish_url string `json:"publish_url,omitempty"` // API URL to publish this codespace to a new repository.
	Pulls_url string `json:"pulls_url"` // API URL for the Pull Request associated with this codespace, if any.
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Name string `json:"name"` // Automatically generated name of this codespace.
	Runtime_constraints map[string]interface{} `json:"runtime_constraints,omitempty"`
	Retention_expires_at string `json:"retention_expires_at,omitempty"` // When a codespace will be auto-deleted based on the "retention_period_minutes" and "last_used_at"
	Environment_id string `json:"environment_id"` // UUID identifying this codespace's environment.
	Updated_at string `json:"updated_at"`
	Stop_url string `json:"stop_url"` // API URL to stop this codespace.
	Repository GeneratedType_Minimal_repository `json:"repository"` // Minimal Repository
	Git_status map[string]interface{} `json:"git_status"` // Details about the codespace's git repository.
	Created_at string `json:"created_at"`
	Machine GeneratedType_Nullable_codespace_machine `json:"machine"` // A description of the machine powering a codespace.
	Retention_period_minutes int `json:"retention_period_minutes,omitempty"` // Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
	Recent_folders []string `json:"recent_folders"`
	State string `json:"state"` // State of this codespace.
	Last_used_at string `json:"last_used_at"` // Last known time this codespace was started.
	Pending_operation bool `json:"pending_operation,omitempty"` // Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
}

// GeneratedType_Classroom_assignment_grade represents the GeneratedType_Classroom_assignment_grade schema from the OpenAPI specification
type GeneratedType_Classroom_assignment_grade struct {
	Assignment_name string `json:"assignment_name"` // Name of the assignment
	Group_name string `json:"group_name,omitempty"` // If a group assignment, name of the group the student is in
	Points_available int `json:"points_available"` // Number of points available for the assignment
	Roster_identifier string `json:"roster_identifier"` // Roster identifier of the student
	Starter_code_url string `json:"starter_code_url"` // URL of the starter code for the assignment
	Student_repository_url string `json:"student_repository_url"` // URL of the student's assignment repository
	Submission_timestamp string `json:"submission_timestamp"` // Timestamp of the student's assignment submission
	Points_awarded int `json:"points_awarded"` // Number of points awarded to the student
	Assignment_url string `json:"assignment_url"` // URL of the assignment
	Github_username string `json:"github_username"` // GitHub username of the student
	Student_repository_name string `json:"student_repository_name"` // Name of the student's assignment repository
}

// GeneratedType_Repository_rule_max_file_size represents the GeneratedType_Repository_rule_max_file_size schema from the OpenAPI specification
type GeneratedType_Repository_rule_max_file_size struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType_Issue_type represents the GeneratedType_Issue_type schema from the OpenAPI specification
type GeneratedType_Issue_type struct {
	Is_enabled bool `json:"is_enabled,omitempty"` // The enabled state of the issue type.
	Name string `json:"name"` // The name of the issue type.
	Node_id string `json:"node_id"` // The node identifier of the issue type.
	Updated_at string `json:"updated_at,omitempty"` // The time the issue type last updated.
	Color string `json:"color,omitempty"` // The color of the issue type.
	Created_at string `json:"created_at,omitempty"` // The time the issue type created.
	Description string `json:"description"` // The description of the issue type.
	Id int `json:"id"` // The unique identifier of the issue type.
}

// GeneratedType_Webhook_pull_request_review_submitted represents the GeneratedType_Webhook_pull_request_review_submitted schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_review_submitted struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Review Webhooksreview `json:"review"` // The review that was affected.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_repository_vulnerability_alert_create represents the GeneratedType_Webhook_repository_vulnerability_alert_create schema from the OpenAPI specification
type GeneratedType_Webhook_repository_vulnerability_alert_create struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert Webhooksalert `json:"alert"` // The security alert of the vulnerable dependency.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Code_scanning_sarifs_status represents the GeneratedType_Code_scanning_sarifs_status schema from the OpenAPI specification
type GeneratedType_Code_scanning_sarifs_status struct {
	Errors []string `json:"errors,omitempty"` // Any errors that ocurred during processing of the delivery.
	Processing_status string `json:"processing_status,omitempty"` // `pending` files have not yet been processed, while `complete` means results from the SARIF have been stored. `failed` files have either not been processed at all, or could only be partially processed.
	Analyses_url string `json:"analyses_url,omitempty"` // The REST API URL for getting the analyses associated with the upload.
}

// Webhooksmilestone represents the Webhooksmilestone schema from the OpenAPI specification
type Webhooksmilestone struct {
	Labels_url string `json:"labels_url"`
	Open_issues int `json:"open_issues"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Closed_at string `json:"closed_at"`
	Id int `json:"id"`
	Creator map[string]interface{} `json:"creator"`
	Closed_issues int `json:"closed_issues"`
	State string `json:"state"` // The state of the milestone.
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Title string `json:"title"` // The title of the milestone.
	Html_url string `json:"html_url"`
	Number int `json:"number"` // The number of the milestone.
	Due_on string `json:"due_on"`
	Description string `json:"description"`
}

// GeneratedType_Simple_classroom represents the GeneratedType_Simple_classroom schema from the OpenAPI specification
type GeneratedType_Simple_classroom struct {
	Archived bool `json:"archived"` // Returns whether classroom is archived or not.
	Id int `json:"id"` // Unique identifier of the classroom.
	Name string `json:"name"` // The name of the classroom.
	Url string `json:"url"` // The url of the classroom on GitHub Classroom.
}

// GeneratedType_Marketplace_listing_plan represents the GeneratedType_Marketplace_listing_plan schema from the OpenAPI specification
type GeneratedType_Marketplace_listing_plan struct {
	Yearly_price_in_cents int `json:"yearly_price_in_cents"`
	Accounts_url string `json:"accounts_url"`
	Number int `json:"number"`
	Url string `json:"url"`
	Description string `json:"description"`
	Id int `json:"id"`
	Name string `json:"name"`
	Price_model string `json:"price_model"`
	Bullets []string `json:"bullets"`
	Has_free_trial bool `json:"has_free_trial"`
	Monthly_price_in_cents int `json:"monthly_price_in_cents"`
	Unit_name string `json:"unit_name"`
	State string `json:"state"`
}

// GeneratedType_Timeline_line_commented_event represents the GeneratedType_Timeline_line_commented_event schema from the OpenAPI specification
type GeneratedType_Timeline_line_commented_event struct {
	Comments []GeneratedType_Pull_request_review_comment `json:"comments,omitempty"`
	Event string `json:"event,omitempty"`
	Node_id string `json:"node_id,omitempty"`
}

// GeneratedType_Empty_object represents the GeneratedType_Empty_object schema from the OpenAPI specification
type GeneratedType_Empty_object struct {
}

// GeneratedType_Codespace_machine represents the GeneratedType_Codespace_machine schema from the OpenAPI specification
type GeneratedType_Codespace_machine struct {
	Operating_system string `json:"operating_system"` // The operating system of the machine.
	Prebuild_availability string `json:"prebuild_availability"` // Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be "null" if prebuilds are not supported or prebuild availability could not be determined. Value will be "none" if no prebuild is available. Latest values "ready" and "in_progress" indicate the prebuild availability status.
	Storage_in_bytes int `json:"storage_in_bytes"` // How much storage is available to the codespace.
	Cpus int `json:"cpus"` // How many cores are available to the codespace.
	Display_name string `json:"display_name"` // The display name of the machine includes cores, memory, and storage.
	Memory_in_bytes int `json:"memory_in_bytes"` // How much memory is available to the codespace.
	Name string `json:"name"` // The name of the machine.
}

// GeneratedType_Webhook_installation_repositories_added represents the GeneratedType_Webhook_installation_repositories_added schema from the OpenAPI specification
type GeneratedType_Webhook_installation_repositories_added struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories_removed []map[string]interface{} `json:"repositories_removed"` // An array of repository objects, which were removed from the installation.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation Installation `json:"installation"` // Installation
	Repositories_added []map[string]interface{} `json:"repositories_added"` // An array of repository objects, which were added to the installation.
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester Webhooksuser `json:"requester"`
}

// GeneratedType_Webhook_pull_request_review_comment_deleted represents the GeneratedType_Webhook_pull_request_review_comment_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_review_comment_deleted struct {
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhooksreviewcomment `json:"comment"` // The [comment](https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request) itself.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// Event represents the Event schema from the OpenAPI specification
type Event struct {
	TypeField string `json:"type"`
	Actor Actor `json:"actor"` // Actor
	Created_at string `json:"created_at"`
	Id string `json:"id"`
	Org Actor `json:"org,omitempty"` // Actor
	Payload map[string]interface{} `json:"payload"`
	Public bool `json:"public"`
	Repo map[string]interface{} `json:"repo"`
}

// GeneratedType_Webhook_workflow_job_waiting represents the GeneratedType_Webhook_workflow_job_waiting schema from the OpenAPI specification
type GeneratedType_Webhook_workflow_job_waiting struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Workflow_job map[string]interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_installation_new_permissions_accepted represents the GeneratedType_Webhook_installation_new_permissions_accepted schema from the OpenAPI specification
type GeneratedType_Webhook_installation_new_permissions_accepted struct {
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
}

// GeneratedType_Webhook_check_run_requested_action_form_encoded represents the GeneratedType_Webhook_check_run_requested_action_form_encoded schema from the OpenAPI specification
type GeneratedType_Webhook_check_run_requested_action_form_encoded struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.requested_action JSON payload. The decoded payload is a JSON object.
}

// GeneratedType_Webhook_repository_vulnerability_alert_dismiss represents the GeneratedType_Webhook_repository_vulnerability_alert_dismiss schema from the OpenAPI specification
type GeneratedType_Webhook_repository_vulnerability_alert_dismiss struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert map[string]interface{} `json:"alert"` // The security alert of the vulnerable dependency.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_label_created represents the GeneratedType_Webhook_label_created schema from the OpenAPI specification
type GeneratedType_Webhook_label_created struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Repository_advisory represents the GeneratedType_Repository_advisory schema from the OpenAPI specification
type GeneratedType_Repository_advisory struct {
	Summary string `json:"summary"` // A short summary of the advisory.
	Cwes []map[string]interface{} `json:"cwes"`
	Description string `json:"description"` // A detailed description of what the advisory entails.
	Submission map[string]interface{} `json:"submission"`
	Collaborating_users []GeneratedType_Simple_user `json:"collaborating_users"` // A list of users that collaborate on the advisory.
	Cvss_severities GeneratedType_Cvss_severities `json:"cvss_severities,omitempty"`
	Cvss map[string]interface{} `json:"cvss"`
	Identifiers []map[string]interface{} `json:"identifiers"`
	Vulnerabilities []GeneratedType_Repository_advisory_vulnerability `json:"vulnerabilities"`
	Severity string `json:"severity"` // The severity of the advisory.
	State string `json:"state"` // The state of the advisory.
	Url string `json:"url"` // The API URL for the advisory.
	Author interface{} `json:"author"` // The author of the advisory.
	Collaborating_teams []Team `json:"collaborating_teams"` // A list of teams that collaborate on the advisory.
	Created_at string `json:"created_at"` // The date and time of when the advisory was created, in ISO 8601 format.
	Credits_detailed []GeneratedType_Repository_advisory_credit `json:"credits_detailed"`
	Ghsa_id string `json:"ghsa_id"` // The GitHub Security Advisory ID.
	Private_fork interface{} `json:"private_fork"` // A temporary private fork of the advisory's repository for collaborating on a fix.
	Credits []map[string]interface{} `json:"credits"`
	Updated_at string `json:"updated_at"` // The date and time of when the advisory was last updated, in ISO 8601 format.
	Withdrawn_at string `json:"withdrawn_at"` // The date and time of when the advisory was withdrawn, in ISO 8601 format.
	Cve_id string `json:"cve_id"` // The Common Vulnerabilities and Exposures (CVE) ID.
	Published_at string `json:"published_at"` // The date and time of when the advisory was published, in ISO 8601 format.
	Publisher interface{} `json:"publisher"` // The publisher of the advisory.
	Closed_at string `json:"closed_at"` // The date and time of when the advisory was closed, in ISO 8601 format.
	Html_url string `json:"html_url"` // The URL for the advisory.
	Cwe_ids []string `json:"cwe_ids"` // A list of only the CWE IDs.
}

// GeneratedType_Gpg_key represents the GeneratedType_Gpg_key schema from the OpenAPI specification
type GeneratedType_Gpg_key struct {
	Created_at string `json:"created_at"`
	Id int64 `json:"id"`
	Can_encrypt_comms bool `json:"can_encrypt_comms"`
	Revoked bool `json:"revoked"`
	Can_sign bool `json:"can_sign"`
	Public_key string `json:"public_key"`
	Emails []map[string]interface{} `json:"emails"`
	Key_id string `json:"key_id"`
	Expires_at string `json:"expires_at"`
	Raw_key string `json:"raw_key"`
	Subkeys []map[string]interface{} `json:"subkeys"`
	Can_certify bool `json:"can_certify"`
	Primary_key_id int `json:"primary_key_id"`
	Can_encrypt_storage bool `json:"can_encrypt_storage"`
	Name string `json:"name,omitempty"`
}

// GeneratedType_Copilot_organization_details represents the GeneratedType_Copilot_organization_details schema from the OpenAPI specification
type GeneratedType_Copilot_organization_details struct {
	Seat_breakdown GeneratedType_Copilot_organization_seat_breakdown `json:"seat_breakdown"` // The breakdown of Copilot Business seats for the organization.
	Seat_management_setting string `json:"seat_management_setting"` // The mode of assigning new seats.
	Cli string `json:"cli,omitempty"` // The organization policy for allowing or disallowing Copilot in the CLI.
	Ide_chat string `json:"ide_chat,omitempty"` // The organization policy for allowing or disallowing Copilot Chat in the IDE.
	Plan_type string `json:"plan_type,omitempty"` // The Copilot plan of the organization, or the parent enterprise, when applicable.
	Platform_chat string `json:"platform_chat,omitempty"` // The organization policy for allowing or disallowing Copilot features on GitHub.com.
	Public_code_suggestions string `json:"public_code_suggestions"` // The organization policy for allowing or blocking suggestions matching public code (duplication detection filter).
}

// GeneratedType_Assigned_issue_event represents the GeneratedType_Assigned_issue_event schema from the OpenAPI specification
type GeneratedType_Assigned_issue_event struct {
	Created_at string `json:"created_at"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Id int `json:"id"`
	Assignee GeneratedType_Simple_user `json:"assignee"` // A GitHub user.
	Node_id string `json:"node_id"`
	Event string `json:"event"`
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Assigner GeneratedType_Simple_user `json:"assigner"` // A GitHub user.
	Commit_url string `json:"commit_url"`
}

// GeneratedType_Webhook_custom_property_promoted_to_enterprise represents the GeneratedType_Webhook_custom_property_promoted_to_enterprise schema from the OpenAPI specification
type GeneratedType_Webhook_custom_property_promoted_to_enterprise struct {
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Definition GeneratedType_Custom_property `json:"definition"` // Custom property defined on an organization
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_pull_request_converted_to_draft represents the GeneratedType_Webhook_pull_request_converted_to_draft schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_converted_to_draft struct {
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType_Pull_request_webhook `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_delete represents the GeneratedType_Webhook_delete schema from the OpenAPI specification
type GeneratedType_Webhook_delete struct {
	Pusher_type string `json:"pusher_type"` // The pusher type for the event. Can be either `user` or a deploy key.
	Ref string `json:"ref"` // The [`git ref`](https://docs.github.com/rest/git/refs#get-a-reference) resource.
	Ref_type string `json:"ref_type"` // The type of Git ref object deleted in the repository.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Code_scanning_autofix_commits represents the GeneratedType_Code_scanning_autofix_commits schema from the OpenAPI specification
type GeneratedType_Code_scanning_autofix_commits struct {
	Message string `json:"message,omitempty"` // Commit message to be used.
	Target_ref string `json:"target_ref,omitempty"` // The Git reference of target branch for the commit. Branch needs to already exist. For more information, see "[Git References](https://git-scm.com/book/en/v2/Git-Internals-Git-References)" in the Git documentation.
}

// GeneratedType_Codespace_export_details represents the GeneratedType_Codespace_export_details schema from the OpenAPI specification
type GeneratedType_Codespace_export_details struct {
	Completed_at string `json:"completed_at,omitempty"` // Completion time of the last export operation
	Export_url string `json:"export_url,omitempty"` // Url for fetching export details
	Html_url string `json:"html_url,omitempty"` // Web url for the exported branch
	Id string `json:"id,omitempty"` // Id for the export details
	Sha string `json:"sha,omitempty"` // Git commit SHA of the exported branch
	State string `json:"state,omitempty"` // State of the latest export
	Branch string `json:"branch,omitempty"` // Name of the exported branch
}

// GeneratedType_Code_scanning_variant_analysis represents the GeneratedType_Code_scanning_variant_analysis schema from the OpenAPI specification
type GeneratedType_Code_scanning_variant_analysis struct {
	Updated_at string `json:"updated_at,omitempty"` // The date and time at which the variant analysis was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Skipped_repositories map[string]interface{} `json:"skipped_repositories,omitempty"` // Information about repositories that were skipped from processing. This information is only available to the user that initiated the variant analysis.
	Completed_at string `json:"completed_at,omitempty"` // The date and time at which the variant analysis was completed, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ. Will be null if the variant analysis has not yet completed or this information is not available.
	Id int `json:"id"` // The ID of the variant analysis.
	Query_pack_url string `json:"query_pack_url"` // The download url for the query pack.
	Scanned_repositories []map[string]interface{} `json:"scanned_repositories,omitempty"`
	Actions_workflow_run_id int `json:"actions_workflow_run_id,omitempty"` // The GitHub Actions workflow run used to execute this variant analysis. This is only available if the workflow run has started.
	Controller_repo GeneratedType_Simple_repository `json:"controller_repo"` // A GitHub repository.
	Created_at string `json:"created_at,omitempty"` // The date and time at which the variant analysis was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Failure_reason string `json:"failure_reason,omitempty"` // The reason for a failure of the variant analysis. This is only available if the variant analysis has failed.
	Query_language string `json:"query_language"` // The language targeted by the CodeQL query
	Status string `json:"status"`
}

// GeneratedType_Codeowners_errors represents the GeneratedType_Codeowners_errors schema from the OpenAPI specification
type GeneratedType_Codeowners_errors struct {
	Errors []map[string]interface{} `json:"errors"`
}

// GeneratedType_Code_scanning_analysis_deletion represents the GeneratedType_Code_scanning_analysis_deletion schema from the OpenAPI specification
type GeneratedType_Code_scanning_analysis_deletion struct {
	Confirm_delete_url string `json:"confirm_delete_url"` // Next deletable analysis in chain, with last analysis deletion confirmation
	Next_analysis_url string `json:"next_analysis_url"` // Next deletable analysis in chain, without last analysis deletion confirmation
}

// GeneratedType_Webhook_registry_package_published represents the GeneratedType_Webhook_registry_package_published schema from the OpenAPI specification
type GeneratedType_Webhook_registry_package_published struct {
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Registry_package map[string]interface{} `json:"registry_package"`
}

// GeneratedType_Git_tree represents the GeneratedType_Git_tree schema from the OpenAPI specification
type GeneratedType_Git_tree struct {
	Url string `json:"url,omitempty"`
	Sha string `json:"sha"`
	Tree []map[string]interface{} `json:"tree"` // Objects specifying a tree structure
	Truncated bool `json:"truncated"`
}

// GeneratedType_Webhook_marketplace_purchase_purchased represents the GeneratedType_Webhook_marketplace_purchase_purchased schema from the OpenAPI specification
type GeneratedType_Webhook_marketplace_purchase_purchased struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Effective_date string `json:"effective_date"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Previous_marketplace_purchase Webhookspreviousmarketplacepurchase `json:"previous_marketplace_purchase,omitempty"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_issues_deleted represents the GeneratedType_Webhook_issues_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_issues_deleted struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Release represents the Release schema from the OpenAPI specification
type Release struct {
	Html_url string `json:"html_url"`
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Upload_url string `json:"upload_url"`
	Author GeneratedType_Simple_user `json:"author"` // A GitHub user.
	Published_at string `json:"published_at"`
	Tag_name string `json:"tag_name"` // The name of the tag.
	Mentions_count int `json:"mentions_count,omitempty"`
	Target_commitish string `json:"target_commitish"` // Specifies the commitish value that determines where the Git tag is created from.
	Body string `json:"body,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	Created_at string `json:"created_at"`
	Discussion_url string `json:"discussion_url,omitempty"` // The URL of the release discussion.
	Name string `json:"name"`
	Prerelease bool `json:"prerelease"` // Whether to identify the release as a prerelease or a full release.
	Assets_url string `json:"assets_url"`
	Body_html string `json:"body_html,omitempty"`
	Tarball_url string `json:"tarball_url"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Zipball_url string `json:"zipball_url"`
	Draft bool `json:"draft"` // true to create a draft (unpublished) release, false to create a published one.
	Assets []GeneratedType_Release_asset `json:"assets"`
}

// GeneratedType_Webhook_issue_comment_created represents the GeneratedType_Webhook_issue_comment_created schema from the OpenAPI specification
type GeneratedType_Webhook_issue_comment_created struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment) itself.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) the comment belongs to.
}

// GeneratedType_Org_repo_custom_property_values represents the GeneratedType_Org_repo_custom_property_values schema from the OpenAPI specification
type GeneratedType_Org_repo_custom_property_values struct {
	Properties []GeneratedType_Custom_property_value `json:"properties"` // List of custom property names and associated values
	Repository_full_name string `json:"repository_full_name"`
	Repository_id int `json:"repository_id"`
	Repository_name string `json:"repository_name"`
}

// GeneratedType_Marketplace_purchase represents the GeneratedType_Marketplace_purchase schema from the OpenAPI specification
type GeneratedType_Marketplace_purchase struct {
	Marketplace_pending_change map[string]interface{} `json:"marketplace_pending_change,omitempty"`
	Marketplace_purchase map[string]interface{} `json:"marketplace_purchase"`
	Organization_billing_email string `json:"organization_billing_email,omitempty"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Email string `json:"email,omitempty"`
	Id int `json:"id"`
	Login string `json:"login"`
}

// Import represents the Import schema from the OpenAPI specification
type Import struct {
	Large_files_count int `json:"large_files_count,omitempty"`
	Large_files_size int `json:"large_files_size,omitempty"`
	Status string `json:"status"`
	Import_percent int `json:"import_percent,omitempty"`
	Use_lfs bool `json:"use_lfs,omitempty"`
	Commit_count int `json:"commit_count,omitempty"`
	Push_percent int `json:"push_percent,omitempty"`
	Project_choices []map[string]interface{} `json:"project_choices,omitempty"`
	Repository_url string `json:"repository_url"`
	Svc_root string `json:"svc_root,omitempty"`
	Vcs_url string `json:"vcs_url"` // The URL of the originating repository.
	Failed_step string `json:"failed_step,omitempty"`
	Tfvc_project string `json:"tfvc_project,omitempty"`
	Vcs string `json:"vcs"`
	Authors_count int `json:"authors_count,omitempty"`
	Message string `json:"message,omitempty"`
	Has_large_files bool `json:"has_large_files,omitempty"`
	Error_message string `json:"error_message,omitempty"`
	Html_url string `json:"html_url"`
	Status_text string `json:"status_text,omitempty"`
	Authors_url string `json:"authors_url"`
	Svn_root string `json:"svn_root,omitempty"`
	Url string `json:"url"`
}

// Job represents the Job schema from the OpenAPI specification
type Job struct {
	Name string `json:"name"` // The name of the job.
	Run_id int `json:"run_id"` // The id of the associated workflow run.
	Html_url string `json:"html_url"`
	Created_at string `json:"created_at"` // The time that the job created, in ISO 8601 format.
	Node_id string `json:"node_id"`
	Started_at string `json:"started_at"` // The time that the job started, in ISO 8601 format.
	Steps []map[string]interface{} `json:"steps,omitempty"` // Steps in this job.
	Workflow_name string `json:"workflow_name"` // The name of the workflow.
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being run.
	Check_run_url string `json:"check_run_url"`
	Completed_at string `json:"completed_at"` // The time that the job finished, in ISO 8601 format.
	Runner_group_id int `json:"runner_group_id"` // The ID of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Head_branch string `json:"head_branch"` // The name of the current branch.
	Url string `json:"url"`
	Conclusion string `json:"conclusion"` // The outcome of the job.
	Run_attempt int `json:"run_attempt,omitempty"` // Attempt number of the associated workflow run, 1 for first attempt and higher if the workflow was re-run.
	Labels []string `json:"labels"` // Labels for the workflow job. Specified by the "runs_on" attribute in the action's workflow file.
	Status string `json:"status"` // The phase of the lifecycle that the job is currently in.
	Id int `json:"id"` // The id of the job.
	Run_url string `json:"run_url"`
	Runner_group_name string `json:"runner_group_name"` // The name of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Runner_id int `json:"runner_id"` // The ID of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Runner_name string `json:"runner_name"` // The name of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
}

// GeneratedType_Webhook_pull_request_dequeued represents the GeneratedType_Webhook_pull_request_dequeued schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_dequeued struct {
	Reason string `json:"reason"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Pull_request_minimal represents the GeneratedType_Pull_request_minimal schema from the OpenAPI specification
type GeneratedType_Pull_request_minimal struct {
	Base map[string]interface{} `json:"base"`
	Head map[string]interface{} `json:"head"`
	Id int64 `json:"id"`
	Number int `json:"number"`
	Url string `json:"url"`
}

// GeneratedType_Base_gist represents the GeneratedType_Base_gist schema from the OpenAPI specification
type GeneratedType_Base_gist struct {
	Files map[string]interface{} `json:"files"`
	Html_url string `json:"html_url"`
	Updated_at string `json:"updated_at"`
	History []interface{} `json:"history,omitempty"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Comments_enabled bool `json:"comments_enabled,omitempty"`
	Comments_url string `json:"comments_url"`
	Public bool `json:"public"`
	Git_push_url string `json:"git_push_url"`
	Owner GeneratedType_Simple_user `json:"owner,omitempty"` // A GitHub user.
	Description string `json:"description"`
	Forks_url string `json:"forks_url"`
	Truncated bool `json:"truncated,omitempty"`
	Commits_url string `json:"commits_url"`
	Id string `json:"id"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Created_at string `json:"created_at"`
	Comments int `json:"comments"`
	Forks []interface{} `json:"forks,omitempty"`
	Git_pull_url string `json:"git_pull_url"`
}

// GeneratedType_Webhook_projects_v2_project_deleted represents the GeneratedType_Webhook_projects_v2_project_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_project_deleted struct {
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2 GeneratedType_Projects_v2 `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_discussion_edited represents the GeneratedType_Webhook_discussion_edited schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_edited struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_discussion_category_changed represents the GeneratedType_Webhook_discussion_category_changed schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_category_changed struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// License represents the License schema from the OpenAPI specification
type License struct {
	Body string `json:"body"`
	Permissions []string `json:"permissions"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Conditions []string `json:"conditions"`
	Featured bool `json:"featured"`
	Implementation string `json:"implementation"`
	Key string `json:"key"`
	Url string `json:"url"`
	Description string `json:"description"`
	Html_url string `json:"html_url"`
	Limitations []string `json:"limitations"`
	Spdx_id string `json:"spdx_id"`
}

// GeneratedType_Team_full represents the GeneratedType_Team_full schema from the OpenAPI specification
type GeneratedType_Team_full struct {
	Repositories_url string `json:"repositories_url"`
	Node_id string `json:"node_id"`
	Members_count int `json:"members_count"`
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Description string `json:"description"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // Unique identifier of the team
	Members_url string `json:"members_url"`
	Html_url string `json:"html_url"`
	Url string `json:"url"` // URL for the team
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Name string `json:"name"` // Name of the team
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Repos_count int `json:"repos_count"`
	Parent GeneratedType_Nullable_team_simple `json:"parent,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Organization GeneratedType_Team_organization `json:"organization"` // Team Organization
	Updated_at string `json:"updated_at"`
	Notification_setting string `json:"notification_setting,omitempty"` // The notification setting the team has set
	Slug string `json:"slug"`
}

// GeneratedType_Repository_ruleset_conditions_repository_id_target represents the GeneratedType_Repository_ruleset_conditions_repository_id_target schema from the OpenAPI specification
type GeneratedType_Repository_ruleset_conditions_repository_id_target struct {
	Repository_id map[string]interface{} `json:"repository_id"`
}

// GeneratedType_Repository_ruleset_conditions_repository_property_target represents the GeneratedType_Repository_ruleset_conditions_repository_property_target schema from the OpenAPI specification
type GeneratedType_Repository_ruleset_conditions_repository_property_target struct {
	Repository_property map[string]interface{} `json:"repository_property"`
}

// Webhookscomment represents the Webhookscomment schema from the OpenAPI specification
type Webhookscomment struct {
	Parent_id int `json:"parent_id"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Reactions map[string]interface{} `json:"reactions"`
	Updated_at string `json:"updated_at"`
	User map[string]interface{} `json:"user"`
	Created_at string `json:"created_at"`
	Discussion_id int `json:"discussion_id"`
	Repository_url string `json:"repository_url"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"`
	Child_comment_count int `json:"child_comment_count"`
	Id int `json:"id"`
}

// GeneratedType_Webhook_repository_dispatch_sample represents the GeneratedType_Webhook_repository_dispatch_sample schema from the OpenAPI specification
type GeneratedType_Webhook_repository_dispatch_sample struct {
	Branch string `json:"branch"`
	Client_payload map[string]interface{} `json:"client_payload"` // The `client_payload` that was specified in the `POST /repos/{owner}/{repo}/dispatches` request body.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The `event_type` that was specified in the `POST /repos/{owner}/{repo}/dispatches` request body.
}

// GeneratedType_Repository_rule represents the GeneratedType_Repository_rule schema from the OpenAPI specification
type GeneratedType_Repository_rule struct {
}

// GeneratedType_Pull_request_simple represents the GeneratedType_Pull_request_simple schema from the OpenAPI specification
type GeneratedType_Pull_request_simple struct {
	Commits_url string `json:"commits_url"`
	Issue_url string `json:"issue_url"`
	State string `json:"state"`
	Title string `json:"title"`
	Head map[string]interface{} `json:"head"`
	Requested_teams []Team `json:"requested_teams,omitempty"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Body string `json:"body"`
	Html_url string `json:"html_url"`
	Comments_url string `json:"comments_url"`
	Review_comment_url string `json:"review_comment_url"`
	Updated_at string `json:"updated_at"`
	Id int64 `json:"id"`
	Auto_merge GeneratedType_Auto_merge `json:"auto_merge"` // The status of auto merging a pull request.
	Assignee GeneratedType_Nullable_simple_user `json:"assignee"` // A GitHub user.
	Requested_reviewers []GeneratedType_Simple_user `json:"requested_reviewers,omitempty"`
	Node_id string `json:"node_id"`
	Statuses_url string `json:"statuses_url"`
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	Milestone GeneratedType_Nullable_milestone `json:"milestone"` // A collection of related issues and pull requests.
	Assignees []GeneratedType_Simple_user `json:"assignees,omitempty"`
	Review_comments_url string `json:"review_comments_url"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Merge_commit_sha string `json:"merge_commit_sha"`
	Links map[string]interface{} `json:"_links"`
	Labels []map[string]interface{} `json:"labels"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Url string `json:"url"`
	Number int `json:"number"`
	Created_at string `json:"created_at"`
	Locked bool `json:"locked"`
	Closed_at string `json:"closed_at"`
	Merged_at string `json:"merged_at"`
	Base map[string]interface{} `json:"base"`
	Patch_url string `json:"patch_url"`
	Diff_url string `json:"diff_url"`
}

// GeneratedType_Webhook_sub_issues_parent_issue_added represents the GeneratedType_Webhook_sub_issues_parent_issue_added schema from the OpenAPI specification
type GeneratedType_Webhook_sub_issues_parent_issue_added struct {
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
	Action string `json:"action"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Parent_issue_repo Repository `json:"parent_issue_repo"` // A repository on GitHub.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_milestone_deleted represents the GeneratedType_Webhook_milestone_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_milestone_deleted struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// Webhooksdeploykey represents the Webhooksdeploykey schema from the OpenAPI specification
type Webhooksdeploykey struct {
	Id int `json:"id"`
	Key string `json:"key"`
	Read_only bool `json:"read_only"`
	Title string `json:"title"`
	Verified bool `json:"verified"`
	Enabled bool `json:"enabled,omitempty"`
	Last_used string `json:"last_used,omitempty"`
	Created_at string `json:"created_at"`
	Added_by string `json:"added_by,omitempty"`
	Url string `json:"url"`
}

// Contributor represents the Contributor schema from the OpenAPI specification
type Contributor struct {
	Html_url string `json:"html_url,omitempty"`
	TypeField string `json:"type"`
	User_view_type string `json:"user_view_type,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Url string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Contributions int `json:"contributions"`
	Id int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Gists_url string `json:"gists_url,omitempty"`
	Login string `json:"login,omitempty"`
	Organizations_url string `json:"organizations_url,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
}

// GeneratedType_Interaction_limit represents the GeneratedType_Interaction_limit schema from the OpenAPI specification
type GeneratedType_Interaction_limit struct {
	Expiry string `json:"expiry,omitempty"` // The duration of the interaction restriction. Default: `one_day`.
	Limit string `json:"limit"` // The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
}

// GeneratedType_Webhook_project_created represents the GeneratedType_Webhook_project_created schema from the OpenAPI specification
type GeneratedType_Webhook_project_created struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Git_tag represents the GeneratedType_Git_tag schema from the OpenAPI specification
type GeneratedType_Git_tag struct {
	Node_id string `json:"node_id"`
	Object map[string]interface{} `json:"object"`
	Sha string `json:"sha"`
	Tag string `json:"tag"` // Name of the tag
	Tagger map[string]interface{} `json:"tagger"`
	Url string `json:"url"` // URL for the tag
	Verification Verification `json:"verification,omitempty"`
	Message string `json:"message"` // Message describing the purpose of the tag
}

// Webhooksprojectcard represents the Webhooksprojectcard schema from the OpenAPI specification
type Webhooksprojectcard struct {
	After_id int `json:"after_id,omitempty"`
	Column_id int `json:"column_id"`
	Column_url string `json:"column_url"`
	Content_url string `json:"content_url,omitempty"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // The project card's ID
	Node_id string `json:"node_id"`
	Note string `json:"note"`
	Archived bool `json:"archived"` // Whether or not the card is archived
	Creator map[string]interface{} `json:"creator"`
	Project_url string `json:"project_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
}

// GeneratedType_Webhook_member_edited represents the GeneratedType_Webhook_member_edited schema from the OpenAPI specification
type GeneratedType_Webhook_member_edited struct {
	Member Webhooksuser `json:"member"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the collaborator permissions
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Code_security_configuration represents the GeneratedType_Code_security_configuration schema from the OpenAPI specification
type GeneratedType_Code_security_configuration struct {
	Secret_scanning_delegated_bypass string `json:"secret_scanning_delegated_bypass,omitempty"` // The enablement status of secret scanning delegated bypass
	Updated_at string `json:"updated_at,omitempty"`
	Dependency_graph_autosubmit_action string `json:"dependency_graph_autosubmit_action,omitempty"` // The enablement status of Automatic dependency submission
	Enforcement string `json:"enforcement,omitempty"` // The enforcement status for a security configuration
	Secret_scanning_non_provider_patterns string `json:"secret_scanning_non_provider_patterns,omitempty"` // The enablement status of secret scanning non-provider patterns
	Code_scanning_delegated_alert_dismissal string `json:"code_scanning_delegated_alert_dismissal,omitempty"` // The enablement status of code scanning delegated alert dismissal
	Secret_scanning_push_protection string `json:"secret_scanning_push_protection,omitempty"` // The enablement status of secret scanning push protection
	Advanced_security string `json:"advanced_security,omitempty"` // The enablement status of GitHub Advanced Security
	Code_scanning_default_setup string `json:"code_scanning_default_setup,omitempty"` // The enablement status of code scanning default setup
	Code_scanning_options map[string]interface{} `json:"code_scanning_options,omitempty"` // Feature options for code scanning
	Dependency_graph_autosubmit_action_options map[string]interface{} `json:"dependency_graph_autosubmit_action_options,omitempty"` // Feature options for Automatic dependency submission
	Description string `json:"description,omitempty"` // A description of the code security configuration
	Secret_scanning_delegated_alert_dismissal string `json:"secret_scanning_delegated_alert_dismissal,omitempty"` // The enablement status of secret scanning delegated alert dismissal
	Dependabot_alerts string `json:"dependabot_alerts,omitempty"` // The enablement status of Dependabot alerts
	Secret_scanning string `json:"secret_scanning,omitempty"` // The enablement status of secret scanning
	Secret_scanning_generic_secrets string `json:"secret_scanning_generic_secrets,omitempty"` // The enablement status of Copilot secret scanning
	Dependabot_security_updates string `json:"dependabot_security_updates,omitempty"` // The enablement status of Dependabot security updates
	Html_url string `json:"html_url,omitempty"` // The URL of the configuration
	Name string `json:"name,omitempty"` // The name of the code security configuration. Must be unique within the organization.
	Private_vulnerability_reporting string `json:"private_vulnerability_reporting,omitempty"` // The enablement status of private vulnerability reporting
	Secret_scanning_delegated_bypass_options map[string]interface{} `json:"secret_scanning_delegated_bypass_options,omitempty"` // Feature options for secret scanning delegated bypass
	Code_scanning_default_setup_options map[string]interface{} `json:"code_scanning_default_setup_options,omitempty"` // Feature options for code scanning default setup
	Id int `json:"id,omitempty"` // The ID of the code security configuration
	Target_type string `json:"target_type,omitempty"` // The type of the code security configuration.
	Secret_scanning_validity_checks string `json:"secret_scanning_validity_checks,omitempty"` // The enablement status of secret scanning validity checks
	Url string `json:"url,omitempty"` // The URL of the configuration
	Created_at string `json:"created_at,omitempty"`
	Dependency_graph string `json:"dependency_graph,omitempty"` // The enablement status of Dependency Graph
}

// GeneratedType_Webhook_issues_assigned represents the GeneratedType_Webhook_issues_assigned schema from the OpenAPI specification
type GeneratedType_Webhook_issues_assigned struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action that was performed.
	Assignee Webhooksuser `json:"assignee,omitempty"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Campaign_summary represents the GeneratedType_Campaign_summary schema from the OpenAPI specification
type GeneratedType_Campaign_summary struct {
	Description string `json:"description"` // The campaign description
	State string `json:"state"` // Indicates whether a campaign is open or closed
	Closed_at string `json:"closed_at,omitempty"` // The date and time the campaign was closed, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ. Will be null if the campaign is still open.
	Published_at string `json:"published_at,omitempty"` // The date and time the campaign was published, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Team_managers []Team `json:"team_managers,omitempty"` // The campaign team managers
	Managers []GeneratedType_Simple_user `json:"managers"` // The campaign managers
	Name string `json:"name,omitempty"` // The campaign name
	Updated_at string `json:"updated_at"` // The date and time the campaign was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Ends_at string `json:"ends_at"` // The date and time the campaign has ended, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Number int `json:"number"` // The number of the newly created campaign
	Alert_stats map[string]interface{} `json:"alert_stats,omitempty"`
	Contact_link string `json:"contact_link"` // The contact link of the campaign.
	Created_at string `json:"created_at"` // The date and time the campaign was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
}

// GeneratedType_Deployment_branch_policy_settings represents the GeneratedType_Deployment_branch_policy_settings schema from the OpenAPI specification
type GeneratedType_Deployment_branch_policy_settings struct {
	Custom_branch_policies bool `json:"custom_branch_policies"` // Whether only branches that match the specified name patterns can deploy to this environment. If `custom_branch_policies` is `true`, `protected_branches` must be `false`; if `custom_branch_policies` is `false`, `protected_branches` must be `true`.
	Protected_branches bool `json:"protected_branches"` // Whether only branches with branch protection rules can deploy to this environment. If `protected_branches` is `true`, `custom_branch_policies` must be `false`; if `protected_branches` is `false`, `custom_branch_policies` must be `true`.
}

// GeneratedType_Actions_public_key represents the GeneratedType_Actions_public_key schema from the OpenAPI specification
type GeneratedType_Actions_public_key struct {
	Key_id string `json:"key_id"` // The identifier for the key.
	Title string `json:"title,omitempty"`
	Url string `json:"url,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Id int `json:"id,omitempty"`
	Key string `json:"key"` // The Base64 encoded public key.
}

// GeneratedType_Codespace_with_full_repository represents the GeneratedType_Codespace_with_full_repository schema from the OpenAPI specification
type GeneratedType_Codespace_with_full_repository struct {
	Idle_timeout_minutes int `json:"idle_timeout_minutes"` // The number of minutes of inactivity after which this codespace will be automatically stopped.
	Updated_at string `json:"updated_at"`
	Pending_operation bool `json:"pending_operation,omitempty"` // Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
	Retention_period_minutes int `json:"retention_period_minutes,omitempty"` // Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
	Id int64 `json:"id"`
	Pending_operation_disabled_reason string `json:"pending_operation_disabled_reason,omitempty"` // Text to show user when codespace is disabled by a pending operation
	Location string `json:"location"` // The initally assigned location of a new codespace.
	Created_at string `json:"created_at"`
	Git_status map[string]interface{} `json:"git_status"` // Details about the codespace's git repository.
	Url string `json:"url"` // API URL for this codespace.
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Web_url string `json:"web_url"` // URL to access this codespace on the web.
	Pulls_url string `json:"pulls_url"` // API URL for the Pull Request associated with this codespace, if any.
	Name string `json:"name"` // Automatically generated name of this codespace.
	Publish_url string `json:"publish_url,omitempty"` // API URL to publish this codespace to a new repository.
	Machines_url string `json:"machines_url"` // API URL to access available alternate machine types for this codespace.
	Machine GeneratedType_Nullable_codespace_machine `json:"machine"` // A description of the machine powering a codespace.
	Display_name string `json:"display_name,omitempty"` // Display name for this codespace.
	Prebuild bool `json:"prebuild"` // Whether the codespace was created from a prebuild.
	Last_used_at string `json:"last_used_at"` // Last known time this codespace was started.
	Retention_expires_at string `json:"retention_expires_at,omitempty"` // When a codespace will be auto-deleted based on the "retention_period_minutes" and "last_used_at"
	Start_url string `json:"start_url"` // API URL to start this codespace.
	Environment_id string `json:"environment_id"` // UUID identifying this codespace's environment.
	State string `json:"state"` // State of this codespace.
	Devcontainer_path string `json:"devcontainer_path,omitempty"` // Path to devcontainer.json from repo root used to create Codespace.
	Runtime_constraints map[string]interface{} `json:"runtime_constraints,omitempty"`
	Stop_url string `json:"stop_url"` // API URL to stop this codespace.
	Billable_owner GeneratedType_Simple_user `json:"billable_owner"` // A GitHub user.
	Idle_timeout_notice string `json:"idle_timeout_notice,omitempty"` // Text to show user when codespace idle timeout minutes has been overriden by an organization policy
	Repository GeneratedType_Full_repository `json:"repository"` // Full Repository
	Recent_folders []string `json:"recent_folders"`
}

// GeneratedType_Webhook_marketplace_purchase_pending_change represents the GeneratedType_Webhook_marketplace_purchase_pending_change schema from the OpenAPI specification
type GeneratedType_Webhook_marketplace_purchase_pending_change struct {
	Effective_date string `json:"effective_date"`
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// Webhooksanswer represents the Webhooksanswer schema from the OpenAPI specification
type Webhooksanswer struct {
	Parent_id interface{} `json:"parent_id"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Body string `json:"body"`
	Node_id string `json:"node_id"`
	Repository_url string `json:"repository_url"`
	Updated_at string `json:"updated_at"`
	User map[string]interface{} `json:"user"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Child_comment_count int `json:"child_comment_count"`
	Created_at string `json:"created_at"`
	Discussion_id int `json:"discussion_id"`
}

// GeneratedType_Repository_rule_max_file_path_length represents the GeneratedType_Repository_rule_max_file_path_length schema from the OpenAPI specification
type GeneratedType_Repository_rule_max_file_path_length struct {
	TypeField string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType_Topic_search_result_item represents the GeneratedType_Topic_search_result_item schema from the OpenAPI specification
type GeneratedType_Topic_search_result_item struct {
	Related []map[string]interface{} `json:"related,omitempty"`
	Description string `json:"description"`
	Logo_url string `json:"logo_url,omitempty"`
	Repository_count int `json:"repository_count,omitempty"`
	Featured bool `json:"featured"`
	Score float64 `json:"score"`
	Created_by string `json:"created_by"`
	Curated bool `json:"curated"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Updated_at string `json:"updated_at"`
	Name string `json:"name"`
	Released string `json:"released"`
	Aliases []map[string]interface{} `json:"aliases,omitempty"`
	Short_description string `json:"short_description"`
	Created_at string `json:"created_at"`
	Display_name string `json:"display_name"`
}

// GeneratedType_Webhook_pull_request_review_comment_edited represents the GeneratedType_Webhook_pull_request_review_comment_edited schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_review_comment_edited struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Comment Webhooksreviewcomment `json:"comment"` // The [comment](https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request) itself.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Changes Webhookschanges `json:"changes"` // The changes to the comment.
}

// GeneratedType_Api_insights_summary_stats represents the GeneratedType_Api_insights_summary_stats schema from the OpenAPI specification
type GeneratedType_Api_insights_summary_stats struct {
	Rate_limited_request_count int64 `json:"rate_limited_request_count,omitempty"` // The total number of requests that were rate limited within the queried time period
	Total_request_count int64 `json:"total_request_count,omitempty"` // The total number of requests within the queried time period
}

// GeneratedType_Webhook_project_deleted represents the GeneratedType_Webhook_project_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_project_deleted struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType_Nullable_repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType_Workflow_run_usage represents the GeneratedType_Workflow_run_usage schema from the OpenAPI specification
type GeneratedType_Workflow_run_usage struct {
	Billable map[string]interface{} `json:"billable"`
	Run_duration_ms int `json:"run_duration_ms,omitempty"`
}

// Webhooksprojectcolumn represents the Webhooksprojectcolumn schema from the OpenAPI specification
type Webhooksprojectcolumn struct {
	Cards_url string `json:"cards_url"`
	Created_at string `json:"created_at"`
	Project_url string `json:"project_url"`
	Url string `json:"url"`
	Updated_at string `json:"updated_at"`
	After_id int `json:"after_id,omitempty"`
	Name string `json:"name"` // Name of the project column
	Id int `json:"id"` // The unique identifier of the project column
	Node_id string `json:"node_id"`
}

// GeneratedType_Added_to_project_issue_event represents the GeneratedType_Added_to_project_issue_event schema from the OpenAPI specification
type GeneratedType_Added_to_project_issue_event struct {
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Commit_id string `json:"commit_id"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_url string `json:"commit_url"`
}

// GeneratedType_Webhook_pull_request_unassigned represents the GeneratedType_Webhook_pull_request_unassigned schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_unassigned struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Number int `json:"number"` // The pull request number.
	Assignee Webhooksusermannequin `json:"assignee,omitempty"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_projects_v2_project_created represents the GeneratedType_Webhook_projects_v2_project_created schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_project_created struct {
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2 GeneratedType_Projects_v2 `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Repository_rule_tag_name_pattern represents the GeneratedType_Repository_rule_tag_name_pattern schema from the OpenAPI specification
type GeneratedType_Repository_rule_tag_name_pattern struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType_Actions_variable represents the GeneratedType_Actions_variable schema from the OpenAPI specification
type GeneratedType_Actions_variable struct {
	Name string `json:"name"` // The name of the variable.
	Updated_at string `json:"updated_at"` // The date and time at which the variable was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Value string `json:"value"` // The value of the variable.
	Created_at string `json:"created_at"` // The date and time at which the variable was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
}

// Snapshot represents the Snapshot schema from the OpenAPI specification
type Snapshot struct {
	Detector map[string]interface{} `json:"detector"` // A description of the detector used.
	Job map[string]interface{} `json:"job"`
	Manifests map[string]interface{} `json:"manifests,omitempty"` // A collection of package manifests, which are a collection of related dependencies declared in a file or representing a logical group of dependencies.
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Ref string `json:"ref"` // The repository branch that triggered this snapshot.
	Scanned string `json:"scanned"` // The time at which the snapshot was scanned.
	Sha string `json:"sha"` // The commit SHA associated with this dependency snapshot. Maximum length: 40 characters.
	Version int `json:"version"` // The version of the repository snapshot submission.
}

// GeneratedType_Pending_deployment represents the GeneratedType_Pending_deployment schema from the OpenAPI specification
type GeneratedType_Pending_deployment struct {
	Wait_timer int `json:"wait_timer"` // The set duration of the wait timer
	Wait_timer_started_at string `json:"wait_timer_started_at"` // The time that the wait timer began.
	Current_user_can_approve bool `json:"current_user_can_approve"` // Whether the currently authenticated user can approve the deployment
	Environment map[string]interface{} `json:"environment"`
	Reviewers []map[string]interface{} `json:"reviewers"` // The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
}

// GeneratedType_Check_suite_preference represents the GeneratedType_Check_suite_preference schema from the OpenAPI specification
type GeneratedType_Check_suite_preference struct {
	Preferences map[string]interface{} `json:"preferences"`
	Repository GeneratedType_Minimal_repository `json:"repository"` // Minimal Repository
}

// Link represents the Link schema from the OpenAPI specification
type Link struct {
	Href string `json:"href"`
}

// GeneratedType_Protected_branch_pull_request_review represents the GeneratedType_Protected_branch_pull_request_review schema from the OpenAPI specification
type GeneratedType_Protected_branch_pull_request_review struct {
	Dismissal_restrictions map[string]interface{} `json:"dismissal_restrictions,omitempty"`
	Require_code_owner_reviews bool `json:"require_code_owner_reviews"`
	Require_last_push_approval bool `json:"require_last_push_approval,omitempty"` // Whether the most recent push must be approved by someone other than the person who pushed it.
	Required_approving_review_count int `json:"required_approving_review_count,omitempty"`
	Url string `json:"url,omitempty"`
	Bypass_pull_request_allowances map[string]interface{} `json:"bypass_pull_request_allowances,omitempty"` // Allow specific users, teams, or apps to bypass pull request requirements.
	Dismiss_stale_reviews bool `json:"dismiss_stale_reviews"`
}

// GeneratedType_Organization_secret_scanning_alert represents the GeneratedType_Organization_secret_scanning_alert schema from the OpenAPI specification
type GeneratedType_Organization_secret_scanning_alert struct {
	State string `json:"state,omitempty"` // Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
	Publicly_leaked bool `json:"publicly_leaked,omitempty"` // Whether the secret was publicly leaked.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypass_request_reviewer_comment string `json:"push_protection_bypass_request_reviewer_comment,omitempty"` // An optional comment when reviewing a push protection bypass.
	Repository GeneratedType_Simple_repository `json:"repository,omitempty"` // A GitHub repository.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Resolved_by GeneratedType_Nullable_simple_user `json:"resolved_by,omitempty"` // A GitHub user.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Supported secret scanning patterns](https://docs.github.com/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
	Number int `json:"number,omitempty"` // The security alert number.
	Push_protection_bypass_request_comment string `json:"push_protection_bypass_request_comment,omitempty"` // An optional comment when requesting a push protection bypass.
	Push_protection_bypass_request_html_url string `json:"push_protection_bypass_request_html_url,omitempty"` // The URL to a push protection bypass request.
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
	Multi_repo bool `json:"multi_repo,omitempty"` // Whether the detected secret was found in multiple repositories in the same organization or enterprise.
	Is_base64_encoded bool `json:"is_base64_encoded,omitempty"` // A boolean value representing whether or not alert is base64 encoded
	Push_protection_bypass_request_reviewer GeneratedType_Nullable_simple_user `json:"push_protection_bypass_request_reviewer,omitempty"` // A GitHub user.
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	Validity string `json:"validity,omitempty"` // The token status as of the latest validity check.
	Push_protection_bypassed_by GeneratedType_Nullable_simple_user `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Resolution_comment string `json:"resolution_comment,omitempty"` // The comment that was optionally added when this alert was closed
	Has_more_locations bool `json:"has_more_locations,omitempty"` // A boolean value representing whether or not the token in the alert was detected in more than one location.
	Secret string `json:"secret,omitempty"` // The secret that was detected.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
	Resolution string `json:"resolution,omitempty"` // **Required when the `state` is `resolved`.** The reason for resolving the alert.
	First_location_detected interface{} `json:"first_location_detected,omitempty"` // Details on the location where the token was initially detected. This can be a commit, wiki commit, issue, discussion, pull request.
}

// GeneratedType_Webhook_project_card_converted represents the GeneratedType_Webhook_project_card_converted schema from the OpenAPI specification
type GeneratedType_Webhook_project_card_converted struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card Webhooksprojectcard `json:"project_card"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Secret_scanning_push_protection_bypass represents the GeneratedType_Secret_scanning_push_protection_bypass schema from the OpenAPI specification
type GeneratedType_Secret_scanning_push_protection_bypass struct {
	Expire_at string `json:"expire_at,omitempty"` // The time that the bypass will expire in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Reason string `json:"reason,omitempty"` // The reason for bypassing push protection.
	Token_type string `json:"token_type,omitempty"` // The token type this bypass is for.
}

// GeneratedType_Webhook_org_block_unblocked represents the GeneratedType_Webhook_org_block_unblocked schema from the OpenAPI specification
type GeneratedType_Webhook_org_block_unblocked struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Blocked_user Webhooksuser `json:"blocked_user"`
}

// GeneratedType_Secret_scanning_location_discussion_body represents the GeneratedType_Secret_scanning_location_discussion_body schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_discussion_body struct {
	Discussion_body_url string `json:"discussion_body_url"` // The URL to the discussion where the secret was detected.
}

// GeneratedType_Webhook_team_deleted represents the GeneratedType_Webhook_team_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_team_deleted struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
}

// GeneratedType_Marketplace_account represents the GeneratedType_Marketplace_account schema from the OpenAPI specification
type GeneratedType_Marketplace_account struct {
	Id int `json:"id"`
	Login string `json:"login"`
	Node_id string `json:"node_id,omitempty"`
	Organization_billing_email string `json:"organization_billing_email,omitempty"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Email string `json:"email,omitempty"`
}

// GeneratedType_Protected_branch represents the GeneratedType_Protected_branch schema from the OpenAPI specification
type GeneratedType_Protected_branch struct {
	Required_pull_request_reviews map[string]interface{} `json:"required_pull_request_reviews,omitempty"`
	Required_signatures map[string]interface{} `json:"required_signatures,omitempty"`
	Restrictions GeneratedType_Branch_restriction_policy `json:"restrictions,omitempty"` // Branch Restriction Policy
	Allow_deletions map[string]interface{} `json:"allow_deletions,omitempty"`
	Enforce_admins map[string]interface{} `json:"enforce_admins,omitempty"`
	Required_linear_history map[string]interface{} `json:"required_linear_history,omitempty"`
	Required_status_checks GeneratedType_Status_check_policy `json:"required_status_checks,omitempty"` // Status Check Policy
	Url string `json:"url"`
	Allow_force_pushes map[string]interface{} `json:"allow_force_pushes,omitempty"`
	Allow_fork_syncing map[string]interface{} `json:"allow_fork_syncing,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
	Block_creations map[string]interface{} `json:"block_creations,omitempty"`
	Required_conversation_resolution map[string]interface{} `json:"required_conversation_resolution,omitempty"`
	Lock_branch map[string]interface{} `json:"lock_branch,omitempty"` // Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
}

// GeneratedType_Deployment_branch_policy_name_pattern_with_type represents the GeneratedType_Deployment_branch_policy_name_pattern_with_type schema from the OpenAPI specification
type GeneratedType_Deployment_branch_policy_name_pattern_with_type struct {
	TypeField string `json:"type,omitempty"` // Whether this rule targets a branch or tag
	Name string `json:"name"` // The name pattern that branches or tags must match in order to deploy to the environment. Wildcard characters will not match `/`. For example, to match branches that begin with `release/` and contain an additional single slash, use `release/*/*`. For more information about pattern matching syntax, see the [Ruby File.fnmatch documentation](https://ruby-doc.org/core-2.5.1/File.html#method-c-fnmatch).
}

// GeneratedType_Security_and_analysis represents the GeneratedType_Security_and_analysis schema from the OpenAPI specification
type GeneratedType_Security_and_analysis struct {
	Dependabot_security_updates map[string]interface{} `json:"dependabot_security_updates,omitempty"` // Enable or disable Dependabot security updates for the repository.
	Secret_scanning map[string]interface{} `json:"secret_scanning,omitempty"`
	Secret_scanning_ai_detection map[string]interface{} `json:"secret_scanning_ai_detection,omitempty"`
	Secret_scanning_non_provider_patterns map[string]interface{} `json:"secret_scanning_non_provider_patterns,omitempty"`
	Secret_scanning_push_protection map[string]interface{} `json:"secret_scanning_push_protection,omitempty"`
	Advanced_security map[string]interface{} `json:"advanced_security,omitempty"`
	Code_security map[string]interface{} `json:"code_security,omitempty"`
}

// Milestone represents the Milestone schema from the OpenAPI specification
type Milestone struct {
	Closed_issues int `json:"closed_issues"`
	State string `json:"state"` // The state of the milestone.
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Description string `json:"description"`
	Title string `json:"title"` // The title of the milestone.
	Url string `json:"url"`
	Due_on string `json:"due_on"`
	Open_issues int `json:"open_issues"`
	Updated_at string `json:"updated_at"`
	Creator GeneratedType_Nullable_simple_user `json:"creator"` // A GitHub user.
	Html_url string `json:"html_url"`
	Labels_url string `json:"labels_url"`
	Node_id string `json:"node_id"`
	Number int `json:"number"` // The number of the milestone.
	Closed_at string `json:"closed_at"`
}

// GeneratedType_Simple_classroom_organization represents the GeneratedType_Simple_classroom_organization schema from the OpenAPI specification
type GeneratedType_Simple_classroom_organization struct {
	Node_id string `json:"node_id"`
	Avatar_url string `json:"avatar_url"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Login string `json:"login"`
	Name string `json:"name"`
}

// GeneratedType_Webhook_sponsorship_cancelled represents the GeneratedType_Webhook_sponsorship_cancelled schema from the OpenAPI specification
type GeneratedType_Webhook_sponsorship_cancelled struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Gist_commit represents the GeneratedType_Gist_commit schema from the OpenAPI specification
type GeneratedType_Gist_commit struct {
	Change_status map[string]interface{} `json:"change_status"`
	Committed_at string `json:"committed_at"`
	Url string `json:"url"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Version string `json:"version"`
}

// GeneratedType_Copilot_usage_metrics_day represents the GeneratedType_Copilot_usage_metrics_day schema from the OpenAPI specification
type GeneratedType_Copilot_usage_metrics_day struct {
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // The total number of Copilot users who engaged with any Copilot feature, for the given day. Examples include but are not limited to accepting a code suggestion, prompting Copilot chat, or triggering a PR Summary. Does not include authentication events. Is not limited to the individual features detailed on the endpoint.
	Copilot_dotcom_chat GeneratedType_Copilot_dotcom_chat `json:"copilot_dotcom_chat,omitempty"` // Usage metrics for Copilot Chat in GitHub.com
	Copilot_dotcom_pull_requests GeneratedType_Copilot_dotcom_pull_requests `json:"copilot_dotcom_pull_requests,omitempty"` // Usage metrics for Copilot for pull requests.
	Copilot_ide_chat GeneratedType_Copilot_ide_chat `json:"copilot_ide_chat,omitempty"` // Usage metrics for Copilot Chat in the IDE.
	Copilot_ide_code_completions GeneratedType_Copilot_ide_code_completions `json:"copilot_ide_code_completions,omitempty"` // Usage metrics for Copilot editor code completions in the IDE.
	Date string `json:"date"` // The date for which the usage metrics are aggregated, in `YYYY-MM-DD` format.
	Total_active_users int `json:"total_active_users,omitempty"` // The total number of Copilot users with activity belonging to any Copilot feature, globally, for the given day. Includes passive activity such as receiving a code suggestion, as well as engagement activity such as accepting a code suggestion or prompting chat. Does not include authentication events. Is not limited to the individual features detailed on the endpoint.
}

// GeneratedType_Repository_advisory_vulnerability represents the GeneratedType_Repository_advisory_vulnerability schema from the OpenAPI specification
type GeneratedType_Repository_advisory_vulnerability struct {
	PackageField map[string]interface{} `json:"package"` // The name of the package affected by the vulnerability.
	Patched_versions string `json:"patched_versions"` // The package version(s) that resolve the vulnerability.
	Vulnerable_functions []string `json:"vulnerable_functions"` // The functions in the package that are affected.
	Vulnerable_version_range string `json:"vulnerable_version_range"` // The range of the package versions affected by the vulnerability.
}

// GeneratedType_Repository_ruleset represents the GeneratedType_Repository_ruleset schema from the OpenAPI specification
type GeneratedType_Repository_ruleset struct {
	Rules []GeneratedType_Repository_rule `json:"rules,omitempty"`
	Source_type string `json:"source_type,omitempty"` // The type of the source of the ruleset
	Target string `json:"target,omitempty"` // The target of the ruleset
	Links map[string]interface{} `json:"_links,omitempty"`
	Current_user_can_bypass string `json:"current_user_can_bypass,omitempty"` // The bypass type of the user making the API request for this ruleset. This field is only returned when querying the repository-level endpoint.
	Enforcement string `json:"enforcement"` // The enforcement level of the ruleset. `evaluate` allows admins to test rules before enforcing them. Admins can view insights on the Rule Insights page (`evaluate` is only available with GitHub Enterprise).
	Id int `json:"id"` // The ID of the ruleset
	Source string `json:"source"` // The name of the source
	Updated_at string `json:"updated_at,omitempty"`
	Bypass_actors []GeneratedType_Repository_ruleset_bypass_actor `json:"bypass_actors,omitempty"` // The actors that can bypass the rules in this ruleset
	Conditions interface{} `json:"conditions,omitempty"`
	Name string `json:"name"` // The name of the ruleset
	Node_id string `json:"node_id,omitempty"`
	Created_at string `json:"created_at,omitempty"`
}

// GeneratedType_Webhook_package_published represents the GeneratedType_Webhook_package_published schema from the OpenAPI specification
type GeneratedType_Webhook_package_published struct {
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	PackageField map[string]interface{} `json:"package"` // Information about the package.
}

// GeneratedType_Webhook_issues_locked represents the GeneratedType_Webhook_issues_locked schema from the OpenAPI specification
type GeneratedType_Webhook_issues_locked struct {
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Actions_hosted_runner_limits represents the GeneratedType_Actions_hosted_runner_limits schema from the OpenAPI specification
type GeneratedType_Actions_hosted_runner_limits struct {
	Public_ips map[string]interface{} `json:"public_ips"` // Provides details of static public IP limits for GitHub-hosted Hosted Runners
}

// GeneratedType_Actions_cache_list represents the GeneratedType_Actions_cache_list schema from the OpenAPI specification
type GeneratedType_Actions_cache_list struct {
	Total_count int `json:"total_count"` // Total number of caches
	Actions_caches []map[string]interface{} `json:"actions_caches"` // Array of caches
}

// GeneratedType_Deployment_simple represents the GeneratedType_Deployment_simple schema from the OpenAPI specification
type GeneratedType_Deployment_simple struct {
	Original_environment string `json:"original_environment,omitempty"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Repository_url string `json:"repository_url"`
	Url string `json:"url"`
	Node_id string `json:"node_id"`
	Statuses_url string `json:"statuses_url"`
	Task string `json:"task"` // Parameter to specify a task to execute
	Transient_environment bool `json:"transient_environment,omitempty"` // Specifies if the given environment is will no longer exist at some point in the future. Default: false.
	Id int `json:"id"` // Unique identifier of the deployment
	Production_environment bool `json:"production_environment,omitempty"` // Specifies if the given environment is one that end-users directly interact with. Default: false.
	Description string `json:"description"`
	Environment string `json:"environment"` // Name for the target deployment environment.
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
}

// GeneratedType_Org_hook represents the GeneratedType_Org_hook schema from the OpenAPI specification
type GeneratedType_Org_hook struct {
	TypeField string `json:"type"`
	Config map[string]interface{} `json:"config"`
	Deliveries_url string `json:"deliveries_url,omitempty"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Active bool `json:"active"`
	Events []string `json:"events"`
	Name string `json:"name"`
	Ping_url string `json:"ping_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
}

// GeneratedType_Webhook_pull_request_locked represents the GeneratedType_Webhook_pull_request_locked schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_locked struct {
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_deployment_review_approved represents the GeneratedType_Webhook_deployment_review_approved schema from the OpenAPI specification
type GeneratedType_Webhook_deployment_review_approved struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Since string `json:"since"`
	Workflow_job_run Webhooksworkflowjobrun `json:"workflow_job_run,omitempty"`
	Reviewers []map[string]interface{} `json:"reviewers,omitempty"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Approver Webhooksapprover `json:"approver,omitempty"`
	Comment string `json:"comment,omitempty"`
	Workflow_job_runs []map[string]interface{} `json:"workflow_job_runs,omitempty"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// Issue represents the Issue schema from the OpenAPI specification
type Issue struct {
	Events_url string `json:"events_url"`
	Created_at string `json:"created_at"`
	Repository_url string `json:"repository_url"`
	Url string `json:"url"` // URL for the issue
	Locked bool `json:"locked"`
	Updated_at string `json:"updated_at"`
	Assignees []GeneratedType_Simple_user `json:"assignees,omitempty"`
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Title string `json:"title"` // Title of the issue
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Body_html string `json:"body_html,omitempty"`
	Draft bool `json:"draft,omitempty"`
	Html_url string `json:"html_url"`
	TypeField GeneratedType_Issue_type `json:"type,omitempty"` // The type of issue.
	Body_text string `json:"body_text,omitempty"`
	Closed_at string `json:"closed_at"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Id int64 `json:"id"`
	Comments int `json:"comments"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Comments_url string `json:"comments_url"`
	State string `json:"state"` // State of the issue; either 'open' or 'closed'
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Number int `json:"number"` // Number uniquely identifying the issue within its repository
	Body string `json:"body,omitempty"` // Contents of the issue
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Closed_by GeneratedType_Nullable_simple_user `json:"closed_by,omitempty"` // A GitHub user.
	Labels []interface{} `json:"labels"` // Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository
	State_reason string `json:"state_reason,omitempty"` // The reason for the current state
	Node_id string `json:"node_id"`
	Labels_url string `json:"labels_url"`
	Milestone GeneratedType_Nullable_milestone `json:"milestone"` // A collection of related issues and pull requests.
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Sub_issues_summary GeneratedType_Sub_issues_summary `json:"sub_issues_summary,omitempty"`
	Assignee GeneratedType_Nullable_simple_user `json:"assignee"` // A GitHub user.
}

// GeneratedType_Webhook_pull_request_enqueued represents the GeneratedType_Webhook_pull_request_enqueued schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_enqueued struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Repository_ruleset_conditions represents the GeneratedType_Repository_ruleset_conditions schema from the OpenAPI specification
type GeneratedType_Repository_ruleset_conditions struct {
	Ref_name map[string]interface{} `json:"ref_name,omitempty"`
}

// Thread represents the Thread schema from the OpenAPI specification
type Thread struct {
	Updated_at string `json:"updated_at"`
	Subject map[string]interface{} `json:"subject"`
	Subscription_url string `json:"subscription_url"`
	Id string `json:"id"`
	Url string `json:"url"`
	Last_read_at string `json:"last_read_at"`
	Reason string `json:"reason"`
	Repository GeneratedType_Minimal_repository `json:"repository"` // Minimal Repository
	Unread bool `json:"unread"`
}

// GeneratedType_Webhook_custom_property_updated represents the GeneratedType_Webhook_custom_property_updated schema from the OpenAPI specification
type GeneratedType_Webhook_custom_property_updated struct {
	Definition GeneratedType_Custom_property `json:"definition"` // Custom property defined on an organization
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Unlabeled_issue_event represents the GeneratedType_Unlabeled_issue_event schema from the OpenAPI specification
type GeneratedType_Unlabeled_issue_event struct {
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Node_id string `json:"node_id"`
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Id int `json:"id"`
	Label map[string]interface{} `json:"label"`
}

// GeneratedType_Webhook_projects_v2_status_update_deleted represents the GeneratedType_Webhook_projects_v2_status_update_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_status_update_deleted struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_status_update GeneratedType_Projects_v2_status_update `json:"projects_v2_status_update"` // An status update belonging to a project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Secret_scanning_scan represents the GeneratedType_Secret_scanning_scan schema from the OpenAPI specification
type GeneratedType_Secret_scanning_scan struct {
	Status string `json:"status,omitempty"` // The state of the scan. Either "completed", "running", or "pending"
	TypeField string `json:"type,omitempty"` // The type of scan
	Completed_at string `json:"completed_at,omitempty"` // The time that the scan was completed. Empty if the scan is running
	Started_at string `json:"started_at,omitempty"` // The time that the scan was started. Empty if the scan is pending
}

// GeneratedType_Webhook_pull_request_unlocked represents the GeneratedType_Webhook_pull_request_unlocked schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_unlocked struct {
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Nullable_issue represents the GeneratedType_Nullable_issue schema from the OpenAPI specification
type GeneratedType_Nullable_issue struct {
	Url string `json:"url"` // URL for the issue
	Events_url string `json:"events_url"`
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	TypeField GeneratedType_Issue_type `json:"type,omitempty"` // The type of issue.
	Assignees []GeneratedType_Simple_user `json:"assignees,omitempty"`
	Body_html string `json:"body_html,omitempty"`
	Html_url string `json:"html_url"`
	Id int64 `json:"id"`
	Body string `json:"body,omitempty"` // Contents of the issue
	Sub_issues_summary GeneratedType_Sub_issues_summary `json:"sub_issues_summary,omitempty"`
	Draft bool `json:"draft,omitempty"`
	State string `json:"state"` // State of the issue; either 'open' or 'closed'
	Timeline_url string `json:"timeline_url,omitempty"`
	Labels_url string `json:"labels_url"`
	Locked bool `json:"locked"`
	Title string `json:"title"` // Title of the issue
	Repository_url string `json:"repository_url"`
	Created_at string `json:"created_at"`
	Body_text string `json:"body_text,omitempty"`
	Comments int `json:"comments"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Assignee GeneratedType_Nullable_simple_user `json:"assignee"` // A GitHub user.
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Number int `json:"number"` // Number uniquely identifying the issue within its repository
	Closed_at string `json:"closed_at"`
	Milestone GeneratedType_Nullable_milestone `json:"milestone"` // A collection of related issues and pull requests.
	Labels []interface{} `json:"labels"` // Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository
	Comments_url string `json:"comments_url"`
	Node_id string `json:"node_id"`
	State_reason string `json:"state_reason,omitempty"` // The reason for the current state
	Closed_by GeneratedType_Nullable_simple_user `json:"closed_by,omitempty"` // A GitHub user.
	Updated_at string `json:"updated_at"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
}

// GeneratedType_Commit_activity represents the GeneratedType_Commit_activity schema from the OpenAPI specification
type GeneratedType_Commit_activity struct {
	Total int `json:"total"`
	Week int `json:"week"`
	Days []int `json:"days"`
}

// GeneratedType_Clone_traffic represents the GeneratedType_Clone_traffic schema from the OpenAPI specification
type GeneratedType_Clone_traffic struct {
	Uniques int `json:"uniques"`
	Clones []Traffic `json:"clones"`
	Count int `json:"count"`
}

// GeneratedType_Review_requested_issue_event represents the GeneratedType_Review_requested_issue_event schema from the OpenAPI specification
type GeneratedType_Review_requested_issue_event struct {
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Id int `json:"id"`
	Requested_reviewer GeneratedType_Simple_user `json:"requested_reviewer,omitempty"` // A GitHub user.
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Event string `json:"event"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Review_requester GeneratedType_Simple_user `json:"review_requester"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
	Node_id string `json:"node_id"`
}

// GeneratedType_Git_commit represents the GeneratedType_Git_commit schema from the OpenAPI specification
type GeneratedType_Git_commit struct {
	Author map[string]interface{} `json:"author"` // Identifying information for the git-user
	Parents []map[string]interface{} `json:"parents"`
	Tree map[string]interface{} `json:"tree"`
	Message string `json:"message"` // Message describing the purpose of the commit
	Sha string `json:"sha"` // SHA for the commit
	Url string `json:"url"`
	Committer map[string]interface{} `json:"committer"` // Identifying information for the git-user
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Verification map[string]interface{} `json:"verification"`
}

// GeneratedType_Webhook_installation_deleted represents the GeneratedType_Webhook_installation_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_installation_deleted struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
}

// GeneratedType_Organization_role represents the GeneratedType_Organization_role schema from the OpenAPI specification
type GeneratedType_Organization_role struct {
	Created_at string `json:"created_at"` // The date and time the role was created.
	Base_role string `json:"base_role,omitempty"` // The system role from which this role inherits permissions.
	Description string `json:"description,omitempty"` // A short description about who this role is for or what permissions it grants.
	Name string `json:"name"` // The name of the role.
	Permissions []string `json:"permissions"` // A list of permissions included in this role.
	Updated_at string `json:"updated_at"` // The date and time the role was last updated.
	Id int64 `json:"id"` // The unique identifier of the role.
	Organization GeneratedType_Nullable_simple_user `json:"organization"` // A GitHub user.
	Source string `json:"source,omitempty"` // Source answers the question, "where did this role come from?"
}

// GeneratedType_Reaction_rollup represents the GeneratedType_Reaction_rollup schema from the OpenAPI specification
type GeneratedType_Reaction_rollup struct {
	Hooray int `json:"hooray"`
	Field1 int `json:"-1"`
	Eyes int `json:"eyes"`
	Heart int `json:"heart"`
	Rocket int `json:"rocket"`
	Total_count int `json:"total_count"`
	Url string `json:"url"`
	Laugh int `json:"laugh"`
	Field1_1 int `json:"+1"`
	Confused int `json:"confused"`
}

// GeneratedType_Webhook_sponsorship_pending_tier_change represents the GeneratedType_Webhook_sponsorship_pending_tier_change schema from the OpenAPI specification
type GeneratedType_Webhook_sponsorship_pending_tier_change struct {
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Effective_date string `json:"effective_date,omitempty"` // The `pending_cancellation` and `pending_tier_change` event types will include the date the cancellation or tier change will take effect.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Changes Webhookschanges8 `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Custom_property represents the GeneratedType_Custom_property schema from the OpenAPI specification
type GeneratedType_Custom_property struct {
	Allowed_values []string `json:"allowed_values,omitempty"` // An ordered list of the allowed values of the property. The property can have up to 200 allowed values.
	Default_value string `json:"default_value,omitempty"` // Default value of the property
	Required bool `json:"required,omitempty"` // Whether the property is required.
	Source_type string `json:"source_type,omitempty"` // The source type of the property
	Property_name string `json:"property_name"` // The name of the property
	Url string `json:"url,omitempty"` // The URL that can be used to fetch, update, or delete info about this property via the API.
	Value_type string `json:"value_type"` // The type of the value for the property
	Values_editable_by string `json:"values_editable_by,omitempty"` // Who can edit the values of the property
	Description string `json:"description,omitempty"` // Short description of the property
}

// Webhookschanges represents the Webhookschanges schema from the OpenAPI specification
type Webhookschanges struct {
	Body map[string]interface{} `json:"body,omitempty"`
}

// GeneratedType_Check_run represents the GeneratedType_Check_run schema from the OpenAPI specification
type GeneratedType_Check_run struct {
	Conclusion string `json:"conclusion"`
	Url string `json:"url"`
	Check_suite map[string]interface{} `json:"check_suite"`
	Status string `json:"status"` // The phase of the lifecycle that the check is currently in. Statuses of waiting, requested, and pending are reserved for GitHub Actions check runs.
	App GeneratedType_Nullable_integration `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Details_url string `json:"details_url"`
	External_id string `json:"external_id"`
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being checked.
	Id int64 `json:"id"` // The id of the check.
	Output map[string]interface{} `json:"output"`
	Html_url string `json:"html_url"`
	Name string `json:"name"` // The name of the check.
	Node_id string `json:"node_id"`
	Started_at string `json:"started_at"`
	Completed_at string `json:"completed_at"`
	Pull_requests []GeneratedType_Pull_request_minimal `json:"pull_requests"` // Pull requests that are open with a `head_sha` or `head_branch` that matches the check. The returned pull requests do not necessarily indicate pull requests that triggered the check.
	Deployment GeneratedType_Deployment_simple `json:"deployment,omitempty"` // A deployment created as the result of an Actions check run from a workflow that references an environment
}

// GeneratedType_Merge_group represents the GeneratedType_Merge_group schema from the OpenAPI specification
type GeneratedType_Merge_group struct {
	Base_sha string `json:"base_sha"` // The SHA of the merge group's parent commit.
	Head_commit GeneratedType_Simple_commit `json:"head_commit"` // A commit.
	Head_ref string `json:"head_ref"` // The full ref of the merge group.
	Head_sha string `json:"head_sha"` // The SHA of the merge group.
	Base_ref string `json:"base_ref"` // The full ref of the branch the merge group will be merged into.
}

// GeneratedType_Projects_v2 represents the GeneratedType_Projects_v2 schema from the OpenAPI specification
type GeneratedType_Projects_v2 struct {
	Creator GeneratedType_Simple_user `json:"creator"` // A GitHub user.
	Id float64 `json:"id"`
	Short_description string `json:"short_description"`
	Title string `json:"title"`
	Deleted_at string `json:"deleted_at"`
	Deleted_by GeneratedType_Nullable_simple_user `json:"deleted_by"` // A GitHub user.
	Node_id string `json:"node_id"`
	Public bool `json:"public"`
	Updated_at string `json:"updated_at"`
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Description string `json:"description"`
	Number int `json:"number"`
	Closed_at string `json:"closed_at"`
	Created_at string `json:"created_at"`
}

// GeneratedType_Repository_collaborator_permission represents the GeneratedType_Repository_collaborator_permission schema from the OpenAPI specification
type GeneratedType_Repository_collaborator_permission struct {
	Permission string `json:"permission"`
	Role_name string `json:"role_name"`
	User GeneratedType_Nullable_collaborator `json:"user"` // Collaborator
}

// GeneratedType_Package_version represents the GeneratedType_Package_version schema from the OpenAPI specification
type GeneratedType_Package_version struct {
	Url string `json:"url"`
	Deleted_at string `json:"deleted_at,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // Unique identifier of the package version.
	License string `json:"license,omitempty"`
	Name string `json:"name"` // The name of the package version.
	Package_html_url string `json:"package_html_url"`
	Description string `json:"description,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType_Timeline_committed_event represents the GeneratedType_Timeline_committed_event schema from the OpenAPI specification
type GeneratedType_Timeline_committed_event struct {
	Verification map[string]interface{} `json:"verification"`
	Author map[string]interface{} `json:"author"` // Identifying information for the git-user
	Message string `json:"message"` // Message describing the purpose of the commit
	Committer map[string]interface{} `json:"committer"` // Identifying information for the git-user
	Html_url string `json:"html_url"`
	Tree map[string]interface{} `json:"tree"`
	Event string `json:"event,omitempty"`
	Node_id string `json:"node_id"`
	Parents []map[string]interface{} `json:"parents"`
	Sha string `json:"sha"` // SHA for the commit
	Url string `json:"url"`
}

// GeneratedType_Actions_cache_usage_by_repository represents the GeneratedType_Actions_cache_usage_by_repository schema from the OpenAPI specification
type GeneratedType_Actions_cache_usage_by_repository struct {
	Active_caches_count int `json:"active_caches_count"` // The number of active caches in the repository.
	Active_caches_size_in_bytes int `json:"active_caches_size_in_bytes"` // The sum of the size in bytes of all the active cache items in the repository.
	Full_name string `json:"full_name"` // The repository owner and name for the cache usage being shown.
}

// GeneratedType_Webhook_issues_edited represents the GeneratedType_Webhook_issues_edited schema from the OpenAPI specification
type GeneratedType_Webhook_issues_edited struct {
	Changes map[string]interface{} `json:"changes"` // The changes to the issue.
	Label Webhookslabel `json:"label,omitempty"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Gist_history represents the GeneratedType_Gist_history schema from the OpenAPI specification
type GeneratedType_Gist_history struct {
	Url string `json:"url,omitempty"`
	User GeneratedType_Nullable_simple_user `json:"user,omitempty"` // A GitHub user.
	Version string `json:"version,omitempty"`
	Change_status map[string]interface{} `json:"change_status,omitempty"`
	Committed_at string `json:"committed_at,omitempty"`
}

// GeneratedType_Merged_upstream represents the GeneratedType_Merged_upstream schema from the OpenAPI specification
type GeneratedType_Merged_upstream struct {
	Merge_type string `json:"merge_type,omitempty"`
	Message string `json:"message,omitempty"`
	Base_branch string `json:"base_branch,omitempty"`
}

// GeneratedType_Ruleset_version represents the GeneratedType_Ruleset_version schema from the OpenAPI specification
type GeneratedType_Ruleset_version struct {
	Actor map[string]interface{} `json:"actor"` // The actor who updated the ruleset
	Updated_at string `json:"updated_at"`
	Version_id int `json:"version_id"` // The ID of the previous version of the ruleset
}

// GeneratedType_Timeline_reviewed_event represents the GeneratedType_Timeline_reviewed_event schema from the OpenAPI specification
type GeneratedType_Timeline_reviewed_event struct {
	Body_html string `json:"body_html,omitempty"`
	Event string `json:"event"`
	Id int `json:"id"` // Unique identifier of the review
	Body string `json:"body"` // The text of the review.
	Body_text string `json:"body_text,omitempty"`
	Node_id string `json:"node_id"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Commit_id string `json:"commit_id"` // A commit SHA for the review.
	Html_url string `json:"html_url"`
	Pull_request_url string `json:"pull_request_url"`
	State string `json:"state"`
	Submitted_at string `json:"submitted_at,omitempty"`
	User GeneratedType_Simple_user `json:"user"` // A GitHub user.
	Links map[string]interface{} `json:"_links"`
}

// GeneratedType_Webhook_discussion_unlocked represents the GeneratedType_Webhook_discussion_unlocked schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_unlocked struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Hook_delivery represents the GeneratedType_Hook_delivery schema from the OpenAPI specification
type GeneratedType_Hook_delivery struct {
	Event string `json:"event"` // The event that triggered the delivery.
	Guid string `json:"guid"` // Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
	Status string `json:"status"` // Description of the status of the attempted delivery
	Status_code int `json:"status_code"` // Status code received when delivery was made.
	Url string `json:"url,omitempty"` // The URL target of the delivery.
	Action string `json:"action"` // The type of activity for the event that triggered the delivery.
	Id int `json:"id"` // Unique identifier of the delivery.
	Installation_id int `json:"installation_id"` // The id of the GitHub App installation associated with this event.
	Throttled_at string `json:"throttled_at,omitempty"` // Time when the webhook delivery was throttled.
	Response map[string]interface{} `json:"response"`
	Duration float64 `json:"duration"` // Time spent delivering.
	Request map[string]interface{} `json:"request"`
	Delivered_at string `json:"delivered_at"` // Time when the delivery was delivered.
	Redelivery bool `json:"redelivery"` // Whether the delivery is a redelivery.
	Repository_id int `json:"repository_id"` // The id of the repository associated with this event.
}

// GeneratedType_Webhook_projects_v2_item_reordered represents the GeneratedType_Webhook_projects_v2_item_reordered schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_item_reordered struct {
	Projects_v2_item GeneratedType_Projects_v2_item `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Repository_rule_workflows represents the GeneratedType_Repository_rule_workflows schema from the OpenAPI specification
type GeneratedType_Repository_rule_workflows struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// Webhooksteam1 represents the Webhooksteam1 schema from the OpenAPI specification
type Webhooksteam1 struct {
	Slug string `json:"slug,omitempty"`
	Url string `json:"url,omitempty"` // URL for the team
	Deleted bool `json:"deleted,omitempty"`
	Id int `json:"id"` // Unique identifier of the team
	Node_id string `json:"node_id,omitempty"`
	Notification_setting string `json:"notification_setting,omitempty"` // Whether team members will receive notifications when their team is @mentioned
	Privacy string `json:"privacy,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Members_url string `json:"members_url,omitempty"`
	Parent map[string]interface{} `json:"parent,omitempty"`
	Permission string `json:"permission,omitempty"` // Permission that the team will have for its repositories
	Description string `json:"description,omitempty"` // Description of the team
	Name string `json:"name"` // Name of the team
	Repositories_url string `json:"repositories_url,omitempty"`
}

// GeneratedType_Webhook_dependabot_alert_fixed represents the GeneratedType_Webhook_dependabot_alert_fixed schema from the OpenAPI specification
type GeneratedType_Webhook_dependabot_alert_fixed struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType_Dependabot_alert `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Pull_request_review_comment represents the GeneratedType_Pull_request_review_comment schema from the OpenAPI specification
type GeneratedType_Pull_request_review_comment struct {
	Path string `json:"path"` // The relative path of the file to which the comment applies.
	Reactions GeneratedType_Reaction_rollup `json:"reactions,omitempty"`
	Side string `json:"side,omitempty"` // The side of the diff to which the comment applies. The side of the last line of the range for a multi-line comment
	Body_text string `json:"body_text,omitempty"`
	Id int64 `json:"id"` // The ID of the pull request review comment.
	User GeneratedType_Simple_user `json:"user"` // A GitHub user.
	Start_line int `json:"start_line,omitempty"` // The first line of the range for a multi-line comment.
	Updated_at string `json:"updated_at"`
	Original_line int `json:"original_line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Commit_id string `json:"commit_id"` // The SHA of the commit to which the comment applies.
	Html_url string `json:"html_url"` // HTML URL for the pull request review comment.
	Pull_request_review_id int64 `json:"pull_request_review_id"` // The ID of the pull request review to which the comment belongs.
	Url string `json:"url"` // URL for the pull request review comment
	Pull_request_url string `json:"pull_request_url"` // URL for the pull request that the review comment belongs to.
	Created_at string `json:"created_at"`
	Links map[string]interface{} `json:"_links"`
	Position int `json:"position,omitempty"` // The line index in the diff to which the comment applies. This field is closing down; use `line` instead.
	Line int `json:"line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Original_commit_id string `json:"original_commit_id"` // The SHA of the original commit to which the comment applies.
	Subject_type string `json:"subject_type,omitempty"` // The level at which the comment is targeted, can be a diff line or a file.
	Start_side string `json:"start_side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Diff_hunk string `json:"diff_hunk"` // The diff of the line that the comment refers to.
	Body string `json:"body"` // The text of the comment.
	Body_html string `json:"body_html,omitempty"`
	Original_start_line int `json:"original_start_line,omitempty"` // The first line of the range for a multi-line comment.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Original_position int `json:"original_position,omitempty"` // The index of the original line in the diff to which the comment applies. This field is closing down; use `original_line` instead.
	In_reply_to_id int `json:"in_reply_to_id,omitempty"` // The comment ID to reply to.
	Node_id string `json:"node_id"` // The node ID of the pull request review comment.
}

// GeneratedType_Simple_repository represents the GeneratedType_Simple_repository schema from the OpenAPI specification
type GeneratedType_Simple_repository struct {
	Blobs_url string `json:"blobs_url"` // A template for the API URL to create or retrieve a raw Git blob in the repository.
	Issues_url string `json:"issues_url"` // A template for the API URL to get information about issues on the repository.
	Milestones_url string `json:"milestones_url"` // A template for the API URL to get information about milestones of the repository.
	Assignees_url string `json:"assignees_url"` // A template for the API URL to list the available assignees for issues in the repository.
	Deployments_url string `json:"deployments_url"` // The API URL to list the deployments of the repository.
	Tags_url string `json:"tags_url"` // The API URL to get information about tags on the repository.
	Html_url string `json:"html_url"` // The URL to view the repository on GitHub.com.
	Git_commits_url string `json:"git_commits_url"` // A template for the API URL to get information about Git commits of the repository.
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Teams_url string `json:"teams_url"` // The API URL to list the teams on the repository.
	Forks_url string `json:"forks_url"` // The API URL to list the forks of the repository.
	Contributors_url string `json:"contributors_url"` // A template for the API URL to list the contributors to the repository.
	Commits_url string `json:"commits_url"` // A template for the API URL to get information about commits on the repository.
	Statuses_url string `json:"statuses_url"` // A template for the API URL to get information about statuses of a commit.
	Notifications_url string `json:"notifications_url"` // A template for the API URL to get information about notifications on the repository.
	Subscribers_url string `json:"subscribers_url"` // The API URL to list the subscribers on the repository.
	Subscription_url string `json:"subscription_url"` // The API URL to subscribe to notifications for this repository.
	Compare_url string `json:"compare_url"` // A template for the API URL to compare two commits or refs.
	Issue_comment_url string `json:"issue_comment_url"` // A template for the API URL to get information about issue comments on the repository.
	Branches_url string `json:"branches_url"` // A template for the API URL to get information about branches in the repository.
	Keys_url string `json:"keys_url"` // A template for the API URL to get information about deploy keys on the repository.
	Node_id string `json:"node_id"` // The GraphQL identifier of the repository.
	Full_name string `json:"full_name"` // The full, globally unique, name of the repository.
	Id int64 `json:"id"` // A unique identifier of the repository.
	Fork bool `json:"fork"` // Whether the repository is a fork.
	Releases_url string `json:"releases_url"` // A template for the API URL to get information about releases on the repository.
	Labels_url string `json:"labels_url"` // A template for the API URL to get information about labels of the repository.
	Private bool `json:"private"` // Whether the repository is private.
	Archive_url string `json:"archive_url"` // A template for the API URL to download the repository as an archive.
	Hooks_url string `json:"hooks_url"` // The API URL to list the hooks on the repository.
	Events_url string `json:"events_url"` // The API URL to list the events of the repository.
	Git_refs_url string `json:"git_refs_url"` // A template for the API URL to get information about Git refs of the repository.
	Issue_events_url string `json:"issue_events_url"` // A template for the API URL to get information about issue events on the repository.
	Languages_url string `json:"languages_url"` // The API URL to get information about the languages of the repository.
	Name string `json:"name"` // The name of the repository.
	Git_tags_url string `json:"git_tags_url"` // A template for the API URL to get information about Git tags of the repository.
	Contents_url string `json:"contents_url"` // A template for the API URL to get the contents of the repository.
	Collaborators_url string `json:"collaborators_url"` // A template for the API URL to get information about collaborators of the repository.
	Merges_url string `json:"merges_url"` // The API URL to merge branches in the repository.
	Downloads_url string `json:"downloads_url"` // The API URL to list the downloads on the repository.
	Stargazers_url string `json:"stargazers_url"` // The API URL to list the stargazers on the repository.
	Trees_url string `json:"trees_url"` // A template for the API URL to create or retrieve a raw Git tree of the repository.
	Url string `json:"url"` // The URL to get more information about the repository from the GitHub API.
	Pulls_url string `json:"pulls_url"` // A template for the API URL to get information about pull requests on the repository.
	Comments_url string `json:"comments_url"` // A template for the API URL to get information about comments on the repository.
	Description string `json:"description"` // The repository description.
}

// GeneratedType_Team_role_assignment represents the GeneratedType_Team_role_assignment schema from the OpenAPI specification
type GeneratedType_Team_role_assignment struct {
	Repositories_url string `json:"repositories_url"`
	Parent GeneratedType_Nullable_team_simple `json:"parent"` // Groups of organization members that gives permissions on specified repositories.
	Description string `json:"description"`
	Members_url string `json:"members_url"`
	Name string `json:"name"`
	Url string `json:"url"`
	Id int `json:"id"`
	Notification_setting string `json:"notification_setting,omitempty"`
	Privacy string `json:"privacy,omitempty"`
	Node_id string `json:"node_id"`
	Permission string `json:"permission"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Slug string `json:"slug"`
	Assignment string `json:"assignment,omitempty"` // Determines if the team has a direct, indirect, or mixed relationship to a role
	Html_url string `json:"html_url"`
}

// GeneratedType_Project_card represents the GeneratedType_Project_card schema from the OpenAPI specification
type GeneratedType_Project_card struct {
	Archived bool `json:"archived,omitempty"` // Whether or not the card is archived
	Column_name string `json:"column_name,omitempty"`
	Column_url string `json:"column_url"`
	Content_url string `json:"content_url,omitempty"`
	Creator GeneratedType_Nullable_simple_user `json:"creator"` // A GitHub user.
	Note string `json:"note"`
	Project_url string `json:"project_url"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Id int64 `json:"id"` // The project card's ID
	Node_id string `json:"node_id"`
	Project_id string `json:"project_id,omitempty"`
	Url string `json:"url"`
}

// GeneratedType_Check_annotation represents the GeneratedType_Check_annotation schema from the OpenAPI specification
type GeneratedType_Check_annotation struct {
	Blob_href string `json:"blob_href"`
	End_column int `json:"end_column"`
	End_line int `json:"end_line"`
	Path string `json:"path"`
	Message string `json:"message"`
	Raw_details string `json:"raw_details"`
	Start_line int `json:"start_line"`
	Title string `json:"title"`
	Annotation_level string `json:"annotation_level"`
	Start_column int `json:"start_column"`
}

// GeneratedType_Webhook_projects_v2_project_closed represents the GeneratedType_Webhook_projects_v2_project_closed schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_project_closed struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2 GeneratedType_Projects_v2 `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Project_collaborator_permission represents the GeneratedType_Project_collaborator_permission schema from the OpenAPI specification
type GeneratedType_Project_collaborator_permission struct {
	Permission string `json:"permission"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
}

// GeneratedType_Team_simple represents the GeneratedType_Team_simple schema from the OpenAPI specification
type GeneratedType_Team_simple struct {
	Repositories_url string `json:"repositories_url"`
	Url string `json:"url"` // URL for the team
	Members_url string `json:"members_url"`
	Name string `json:"name"` // Name of the team
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Notification_setting string `json:"notification_setting,omitempty"` // The notification setting the team has set
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Slug string `json:"slug"`
	Description string `json:"description"` // Description of the team
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the team
	Node_id string `json:"node_id"`
}

// GeneratedType_Webhook_create represents the GeneratedType_Webhook_create schema from the OpenAPI specification
type GeneratedType_Webhook_create struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Pusher_type string `json:"pusher_type"` // The pusher type for the event. Can be either `user` or a deploy key.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Ref string `json:"ref"` // The [`git ref`](https://docs.github.com/rest/git/refs#get-a-reference) resource.
	Ref_type string `json:"ref_type"` // The type of Git ref object created in the repository.
	Description string `json:"description"` // The repository's current description.
	Master_branch string `json:"master_branch"` // The name of the repository's default branch (usually `main`).
}

// GeneratedType_Webhook_issues_transferred represents the GeneratedType_Webhook_issues_transferred schema from the OpenAPI specification
type GeneratedType_Webhook_issues_transferred struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue2 `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Actions_hosted_runner_image represents the GeneratedType_Actions_hosted_runner_image schema from the OpenAPI specification
type GeneratedType_Actions_hosted_runner_image struct {
	Source string `json:"source"` // The image provider.
	Display_name string `json:"display_name"` // Display name for this image.
	Id string `json:"id"` // The ID of the image. Use this ID for the `image` parameter when creating a new larger runner.
	Platform string `json:"platform"` // The operating system of the image.
	Size_gb int `json:"size_gb"` // Image size in GB.
}

// GeneratedType_Ruleset_version_with_state represents the GeneratedType_Ruleset_version_with_state schema from the OpenAPI specification
type GeneratedType_Ruleset_version_with_state struct {
	Actor map[string]interface{} `json:"actor"` // The actor who updated the ruleset
	Updated_at string `json:"updated_at"`
	Version_id int `json:"version_id"` // The ID of the previous version of the ruleset
	State map[string]interface{} `json:"state"` // The state of the ruleset version
}

// GeneratedType_Webhook_workflow_dispatch represents the GeneratedType_Webhook_workflow_dispatch schema from the OpenAPI specification
type GeneratedType_Webhook_workflow_dispatch struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Ref string `json:"ref"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Workflow string `json:"workflow"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Inputs map[string]interface{} `json:"inputs"`
}

// GeneratedType_Webhook_projects_v2_item_edited represents the GeneratedType_Webhook_projects_v2_item_edited schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_item_edited struct {
	Action string `json:"action"`
	Changes interface{} `json:"changes,omitempty"` // The changes made to the item may involve modifications in the item's fields and draft issue body. It includes altered values for text, number, date, single select, and iteration fields, along with the GraphQL node ID of the changed field.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType_Projects_v2_item `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Org_private_registry_configuration represents the GeneratedType_Org_private_registry_configuration schema from the OpenAPI specification
type GeneratedType_Org_private_registry_configuration struct {
	Registry_type string `json:"registry_type"` // The registry type.
	Updated_at string `json:"updated_at"`
	Username string `json:"username,omitempty"` // The username to use when authenticating with the private registry.
	Visibility string `json:"visibility"` // Which type of organization repositories have access to the private registry.
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the private registry configuration.
}

// GeneratedType_Code_search_result_item represents the GeneratedType_Code_search_result_item schema from the OpenAPI specification
type GeneratedType_Code_search_result_item struct {
	Url string `json:"url"`
	File_size int `json:"file_size,omitempty"`
	Git_url string `json:"git_url"`
	Html_url string `json:"html_url"`
	Name string `json:"name"`
	Repository GeneratedType_Minimal_repository `json:"repository"` // Minimal Repository
	Line_numbers []string `json:"line_numbers,omitempty"`
	Sha string `json:"sha"`
	Path string `json:"path"`
	Language string `json:"language,omitempty"`
	Last_modified_at string `json:"last_modified_at,omitempty"`
	Score float64 `json:"score"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
}

// GeneratedType_User_search_result_item represents the GeneratedType_User_search_result_item schema from the OpenAPI specification
type GeneratedType_User_search_result_item struct {
	Events_url string `json:"events_url"`
	Location string `json:"location,omitempty"`
	Followers int `json:"followers,omitempty"`
	Id int64 `json:"id"`
	Public_gists int `json:"public_gists,omitempty"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Node_id string `json:"node_id"`
	Html_url string `json:"html_url"`
	Suspended_at string `json:"suspended_at,omitempty"`
	Following_url string `json:"following_url"`
	Email string `json:"email,omitempty"`
	Hireable bool `json:"hireable,omitempty"`
	Url string `json:"url"`
	Subscriptions_url string `json:"subscriptions_url"`
	Blog string `json:"blog,omitempty"`
	Organizations_url string `json:"organizations_url"`
	Public_repos int `json:"public_repos,omitempty"`
	Followers_url string `json:"followers_url"`
	Gists_url string `json:"gists_url"`
	TypeField string `json:"type"`
	Updated_at string `json:"updated_at,omitempty"`
	Score float64 `json:"score"`
	Created_at string `json:"created_at,omitempty"`
	Site_admin bool `json:"site_admin"`
	Gravatar_id string `json:"gravatar_id"`
	Name string `json:"name,omitempty"`
	Repos_url string `json:"repos_url"`
	Company string `json:"company,omitempty"`
	Login string `json:"login"`
	Starred_url string `json:"starred_url"`
	User_view_type string `json:"user_view_type,omitempty"`
	Following int `json:"following,omitempty"`
	Bio string `json:"bio,omitempty"`
	Received_events_url string `json:"received_events_url"`
	Avatar_url string `json:"avatar_url"`
}

// Stargazer represents the Stargazer schema from the OpenAPI specification
type Stargazer struct {
	Starred_at string `json:"starred_at"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
}

// GeneratedType_Runner_application represents the GeneratedType_Runner_application schema from the OpenAPI specification
type GeneratedType_Runner_application struct {
	Download_url string `json:"download_url"`
	Filename string `json:"filename"`
	Os string `json:"os"`
	Sha256_checksum string `json:"sha256_checksum,omitempty"`
	Temp_download_token string `json:"temp_download_token,omitempty"` // A short lived bearer token used to download the runner, if needed.
	Architecture string `json:"architecture"`
}

// GeneratedType_Repository_webhooks represents the GeneratedType_Repository_webhooks schema from the OpenAPI specification
type GeneratedType_Repository_webhooks struct {
	Pushed_at string `json:"pushed_at"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Labels_url string `json:"labels_url"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Description string `json:"description"`
	Watchers_count int `json:"watchers_count"`
	Issues_url string `json:"issues_url"`
	Teams_url string `json:"teams_url"`
	Created_at string `json:"created_at"`
	Subscribers_url string `json:"subscribers_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Contents_url string `json:"contents_url"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Keys_url string `json:"keys_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Releases_url string `json:"releases_url"`
	Deployments_url string `json:"deployments_url"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Id int64 `json:"id"` // Unique identifier of the repository
	Url string `json:"url"`
	Organization GeneratedType_Nullable_simple_user `json:"organization,omitempty"` // A GitHub user.
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Archive_url string `json:"archive_url"`
	Template_repository map[string]interface{} `json:"template_repository,omitempty"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Issue_events_url string `json:"issue_events_url"`
	Topics []string `json:"topics,omitempty"`
	Events_url string `json:"events_url"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Svn_url string `json:"svn_url"`
	Assignees_url string `json:"assignees_url"`
	Forks_url string `json:"forks_url"`
	Statuses_url string `json:"statuses_url"`
	License GeneratedType_Nullable_license_simple `json:"license"` // License Simple
	Tags_url string `json:"tags_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Open_issues int `json:"open_issues"`
	Open_issues_count int `json:"open_issues_count"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Blobs_url string `json:"blobs_url"`
	Branches_url string `json:"branches_url"`
	Compare_url string `json:"compare_url"`
	Pulls_url string `json:"pulls_url"`
	Git_refs_url string `json:"git_refs_url"`
	Html_url string `json:"html_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Issue_comment_url string `json:"issue_comment_url"`
	Subscription_url string `json:"subscription_url"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Archived bool `json:"archived"` // Whether the repository is archived.
	Language string `json:"language"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Fork bool `json:"fork"`
	Mirror_url string `json:"mirror_url"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Has_pages bool `json:"has_pages"`
	Notifications_url string `json:"notifications_url"`
	Collaborators_url string `json:"collaborators_url"`
	Trees_url string `json:"trees_url"`
	Ssh_url string `json:"ssh_url"`
	Updated_at string `json:"updated_at"`
	Languages_url string `json:"languages_url"`
	Milestones_url string `json:"milestones_url"`
	Clone_url string `json:"clone_url"`
	Git_commits_url string `json:"git_commits_url"`
	Stargazers_count int `json:"stargazers_count"`
	Forks_count int `json:"forks_count"`
	Commits_url string `json:"commits_url"`
	Git_tags_url string `json:"git_tags_url"`
	Contributors_url string `json:"contributors_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Hooks_url string `json:"hooks_url"`
	Forks int `json:"forks"`
	Downloads_url string `json:"downloads_url"`
	Full_name string `json:"full_name"`
	Name string `json:"name"` // The name of the repository.
	Merges_url string `json:"merges_url"`
	Node_id string `json:"node_id"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Homepage string `json:"homepage"`
	Stargazers_url string `json:"stargazers_url"`
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Comments_url string `json:"comments_url"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Git_url string `json:"git_url"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Network_count int `json:"network_count,omitempty"`
	Watchers int `json:"watchers"`
}

// GeneratedType_Nullable_organization_simple represents the GeneratedType_Nullable_organization_simple schema from the OpenAPI specification
type GeneratedType_Nullable_organization_simple struct {
	Avatar_url string `json:"avatar_url"`
	Description string `json:"description"`
	Events_url string `json:"events_url"`
	Issues_url string `json:"issues_url"`
	Members_url string `json:"members_url"`
	Url string `json:"url"`
	Hooks_url string `json:"hooks_url"`
	Id int `json:"id"`
	Login string `json:"login"`
	Node_id string `json:"node_id"`
	Public_members_url string `json:"public_members_url"`
	Repos_url string `json:"repos_url"`
}

// GeneratedType_Actions_repository_permissions represents the GeneratedType_Actions_repository_permissions schema from the OpenAPI specification
type GeneratedType_Actions_repository_permissions struct {
	Enabled bool `json:"enabled"` // Whether GitHub Actions is enabled on the repository.
	Selected_actions_url string `json:"selected_actions_url,omitempty"` // The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
	Allowed_actions string `json:"allowed_actions,omitempty"` // The permissions policy that controls the actions and reusable workflows that are allowed to run.
}

// GeneratedType_Referrer_traffic represents the GeneratedType_Referrer_traffic schema from the OpenAPI specification
type GeneratedType_Referrer_traffic struct {
	Referrer string `json:"referrer"`
	Uniques int `json:"uniques"`
	Count int `json:"count"`
}

// GeneratedType_Nullable_repository_webhooks represents the GeneratedType_Nullable_repository_webhooks schema from the OpenAPI specification
type GeneratedType_Nullable_repository_webhooks struct {
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Issue_comment_url string `json:"issue_comment_url"`
	Milestones_url string `json:"milestones_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Collaborators_url string `json:"collaborators_url"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Network_count int `json:"network_count,omitempty"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Clone_url string `json:"clone_url"`
	Archive_url string `json:"archive_url"`
	Git_commits_url string `json:"git_commits_url"`
	Comments_url string `json:"comments_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	License GeneratedType_Nullable_license_simple `json:"license"` // License Simple
	Created_at string `json:"created_at"`
	Deployments_url string `json:"deployments_url"`
	Subscription_url string `json:"subscription_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Pushed_at string `json:"pushed_at"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Commits_url string `json:"commits_url"`
	Assignees_url string `json:"assignees_url"`
	Template_repository map[string]interface{} `json:"template_repository,omitempty"`
	Forks_count int `json:"forks_count"`
	Hooks_url string `json:"hooks_url"`
	Contributors_url string `json:"contributors_url"`
	Statuses_url string `json:"statuses_url"`
	Stargazers_count int `json:"stargazers_count"`
	Watchers int `json:"watchers"`
	Labels_url string `json:"labels_url"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Languages_url string `json:"languages_url"`
	Mirror_url string `json:"mirror_url"`
	Notifications_url string `json:"notifications_url"`
	Trees_url string `json:"trees_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Git_tags_url string `json:"git_tags_url"`
	Name string `json:"name"` // The name of the repository.
	Full_name string `json:"full_name"`
	Issues_url string `json:"issues_url"`
	Ssh_url string `json:"ssh_url"`
	Id int64 `json:"id"` // Unique identifier of the repository
	Keys_url string `json:"keys_url"`
	Forks int `json:"forks"`
	Url string `json:"url"`
	Stargazers_url string `json:"stargazers_url"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Issue_events_url string `json:"issue_events_url"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Pulls_url string `json:"pulls_url"`
	Open_issues int `json:"open_issues"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Homepage string `json:"homepage"`
	Blobs_url string `json:"blobs_url"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Updated_at string `json:"updated_at"`
	Releases_url string `json:"releases_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Node_id string `json:"node_id"`
	Description string `json:"description"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Svn_url string `json:"svn_url"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Compare_url string `json:"compare_url"`
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Private bool `json:"private"` // Whether the repository is private or public.
	Branches_url string `json:"branches_url"`
	Git_refs_url string `json:"git_refs_url"`
	Open_issues_count int `json:"open_issues_count"`
	Html_url string `json:"html_url"`
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Teams_url string `json:"teams_url"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Language string `json:"language"`
	Organization GeneratedType_Nullable_simple_user `json:"organization,omitempty"` // A GitHub user.
	Git_url string `json:"git_url"`
	Forks_url string `json:"forks_url"`
	Fork bool `json:"fork"`
	Has_pages bool `json:"has_pages"`
	Merges_url string `json:"merges_url"`
	Contents_url string `json:"contents_url"`
	Tags_url string `json:"tags_url"`
	Subscribers_url string `json:"subscribers_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Downloads_url string `json:"downloads_url"`
	Watchers_count int `json:"watchers_count"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Topics []string `json:"topics,omitempty"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Events_url string `json:"events_url"`
}

// GeneratedType_Copilot_organization_seat_breakdown represents the GeneratedType_Copilot_organization_seat_breakdown schema from the OpenAPI specification
type GeneratedType_Copilot_organization_seat_breakdown struct {
	Active_this_cycle int `json:"active_this_cycle,omitempty"` // The number of seats that have used Copilot during the current billing cycle.
	Added_this_cycle int `json:"added_this_cycle,omitempty"` // Seats added during the current billing cycle.
	Inactive_this_cycle int `json:"inactive_this_cycle,omitempty"` // The number of seats that have not used Copilot during the current billing cycle.
	Pending_cancellation int `json:"pending_cancellation,omitempty"` // The number of seats that are pending cancellation at the end of the current billing cycle.
	Pending_invitation int `json:"pending_invitation,omitempty"` // The number of users who have been invited to receive a Copilot seat through this organization.
	Total int `json:"total,omitempty"` // The total number of seats being billed for the organization as of the current billing cycle.
}

// GeneratedType_Webhook_status represents the GeneratedType_Webhook_status schema from the OpenAPI specification
type GeneratedType_Webhook_status struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sha string `json:"sha"` // The Commit SHA.
	Updated_at string `json:"updated_at"`
	Commit map[string]interface{} `json:"commit"`
	Description string `json:"description"` // The optional human-readable description added to the status.
	State string `json:"state"` // The new state. Can be `pending`, `success`, `failure`, or `error`.
	Id int `json:"id"` // The unique identifier of the status.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Branches []map[string]interface{} `json:"branches"` // An array of branch objects containing the status' SHA. Each branch contains the given SHA, but the SHA may or may not be the head of the branch. The array includes a maximum of 10 branches.
	Created_at string `json:"created_at"`
	Target_url string `json:"target_url"` // The optional link added to the status.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Avatar_url string `json:"avatar_url,omitempty"`
	Context string `json:"context"`
	Name string `json:"name"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_custom_property_deleted represents the GeneratedType_Webhook_custom_property_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_custom_property_deleted struct {
	Definition map[string]interface{} `json:"definition"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Webhook_pull_request_auto_merge_enabled represents the GeneratedType_Webhook_pull_request_auto_merge_enabled schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_auto_merge_enabled struct {
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Pull_request map[string]interface{} `json:"pull_request"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"`
	Reason string `json:"reason,omitempty"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
}

// GeneratedType_Oidc_custom_sub represents the GeneratedType_Oidc_custom_sub schema from the OpenAPI specification
type GeneratedType_Oidc_custom_sub struct {
	Include_claim_keys []string `json:"include_claim_keys"` // Array of unique strings. Each claim key can only contain alphanumeric characters and underscores.
}

// GeneratedType_Webhook_discussion_unanswered represents the GeneratedType_Webhook_discussion_unanswered schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_unanswered struct {
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Old_answer Webhooksanswer `json:"old_answer"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_custom_property_created represents the GeneratedType_Webhook_custom_property_created schema from the OpenAPI specification
type GeneratedType_Webhook_custom_property_created struct {
	Action string `json:"action"`
	Definition GeneratedType_Custom_property `json:"definition"` // Custom property defined on an organization
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType_Webhook_registry_package_updated represents the GeneratedType_Webhook_registry_package_updated schema from the OpenAPI specification
type GeneratedType_Webhook_registry_package_updated struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Registry_package map[string]interface{} `json:"registry_package"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Api_overview represents the GeneratedType_Api_overview schema from the OpenAPI specification
type GeneratedType_Api_overview struct {
	Github_enterprise_importer []string `json:"github_enterprise_importer,omitempty"`
	Importer []string `json:"importer,omitempty"`
	Web []string `json:"web,omitempty"`
	Verifiable_password_authentication bool `json:"verifiable_password_authentication"`
	Actions []string `json:"actions,omitempty"`
	Codespaces []string `json:"codespaces,omitempty"`
	Packages []string `json:"packages,omitempty"`
	Ssh_keys []string `json:"ssh_keys,omitempty"`
	Hooks []string `json:"hooks,omitempty"`
	Domains map[string]interface{} `json:"domains,omitempty"`
	Ssh_key_fingerprints map[string]interface{} `json:"ssh_key_fingerprints,omitempty"`
	Dependabot []string `json:"dependabot,omitempty"`
	Api []string `json:"api,omitempty"`
	Copilot []string `json:"copilot,omitempty"`
	Actions_macos []string `json:"actions_macos,omitempty"`
	Pages []string `json:"pages,omitempty"`
	Git []string `json:"git,omitempty"`
}

// GeneratedType_Webhook_pull_request_unlabeled represents the GeneratedType_Webhook_pull_request_unlabeled schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_unlabeled struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label,omitempty"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Number int `json:"number"` // The pull request number.
	Pull_request map[string]interface{} `json:"pull_request"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Webhook_check_suite_completed represents the GeneratedType_Webhook_check_suite_completed schema from the OpenAPI specification
type GeneratedType_Webhook_check_suite_completed struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/checks/suites#get-a-check-suite).
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_organization_member_added represents the GeneratedType_Webhook_organization_member_added schema from the OpenAPI specification
type GeneratedType_Webhook_organization_member_added struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Membership Webhooksmembership `json:"membership"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Webhook_deployment_status_created represents the GeneratedType_Webhook_deployment_status_created schema from the OpenAPI specification
type GeneratedType_Webhook_deployment_status_created struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Workflow Webhooksworkflow `json:"workflow,omitempty"`
	Workflow_run map[string]interface{} `json:"workflow_run,omitempty"`
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Deployment map[string]interface{} `json:"deployment"` // The [deployment](https://docs.github.com/rest/deployments/deployments#list-deployments).
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Deployment_status map[string]interface{} `json:"deployment_status"` // The [deployment status](https://docs.github.com/rest/deployments/statuses#list-deployment-statuses).
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Check_run map[string]interface{} `json:"check_run,omitempty"`
}

// GeneratedType_Repository_rule_required_signatures represents the GeneratedType_Repository_rule_required_signatures schema from the OpenAPI specification
type GeneratedType_Repository_rule_required_signatures struct {
	TypeField string `json:"type"`
}

// GeneratedType_Dependabot_repository_access_details represents the GeneratedType_Dependabot_repository_access_details schema from the OpenAPI specification
type GeneratedType_Dependabot_repository_access_details struct {
	Accessible_repositories []GeneratedType_Nullable_simple_repository `json:"accessible_repositories,omitempty"`
	Default_level string `json:"default_level,omitempty"` // The default repository access level for Dependabot updates.
}

// GeneratedType_Webhook_release_edited represents the GeneratedType_Webhook_release_edited schema from the OpenAPI specification
type GeneratedType_Webhook_release_edited struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
}

// GeneratedType_Webhook_issues_milestoned represents the GeneratedType_Webhook_issues_milestoned schema from the OpenAPI specification
type GeneratedType_Webhook_issues_milestoned struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Milestone Webhooksmilestone `json:"milestone"` // A collection of related issues and pull requests.
}

// GeneratedType_Renamed_issue_event represents the GeneratedType_Renamed_issue_event schema from the OpenAPI specification
type GeneratedType_Renamed_issue_event struct {
	Rename map[string]interface{} `json:"rename"`
	Url string `json:"url"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
}

// GeneratedType_Webhook_security_and_analysis represents the GeneratedType_Webhook_security_and_analysis schema from the OpenAPI specification
type GeneratedType_Webhook_security_and_analysis struct {
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Full_repository `json:"repository"` // Full Repository
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType_Classroom_accepted_assignment represents the GeneratedType_Classroom_accepted_assignment schema from the OpenAPI specification
type GeneratedType_Classroom_accepted_assignment struct {
	Repository GeneratedType_Simple_classroom_repository `json:"repository"` // A GitHub repository view for Classroom
	Students []GeneratedType_Simple_classroom_user `json:"students"`
	Submitted bool `json:"submitted"` // Whether an accepted assignment has been submitted.
	Assignment GeneratedType_Simple_classroom_assignment `json:"assignment"` // A GitHub Classroom assignment
	Commit_count int `json:"commit_count"` // Count of student commits.
	Grade string `json:"grade"` // Most recent grade.
	Id int `json:"id"` // Unique identifier of the repository.
	Passing bool `json:"passing"` // Whether a submission passed.
}

// GeneratedType_Organization_simple represents the GeneratedType_Organization_simple schema from the OpenAPI specification
type GeneratedType_Organization_simple struct {
	Login string `json:"login"`
	Hooks_url string `json:"hooks_url"`
	Node_id string `json:"node_id"`
	Avatar_url string `json:"avatar_url"`
	Id int `json:"id"`
	Members_url string `json:"members_url"`
	Public_members_url string `json:"public_members_url"`
	Url string `json:"url"`
	Issues_url string `json:"issues_url"`
	Repos_url string `json:"repos_url"`
	Description string `json:"description"`
	Events_url string `json:"events_url"`
}

// GeneratedType_Secret_scanning_location_pull_request_title represents the GeneratedType_Secret_scanning_location_pull_request_title schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_pull_request_title struct {
	Pull_request_title_url string `json:"pull_request_title_url"` // The API URL to get the pull request where the secret was detected.
}

// GeneratedType_Webhook_pull_request_synchronize represents the GeneratedType_Webhook_pull_request_synchronize schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_synchronize struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	After string `json:"after"`
	Pull_request map[string]interface{} `json:"pull_request"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Before string `json:"before"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
}

// GeneratedType_View_traffic represents the GeneratedType_View_traffic schema from the OpenAPI specification
type GeneratedType_View_traffic struct {
	Count int `json:"count"`
	Uniques int `json:"uniques"`
	Views []Traffic `json:"views"`
}

// GeneratedType_Webhook_marketplace_purchase_cancelled represents the GeneratedType_Webhook_marketplace_purchase_cancelled schema from the OpenAPI specification
type GeneratedType_Webhook_marketplace_purchase_cancelled struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Previous_marketplace_purchase Webhookspreviousmarketplacepurchase `json:"previous_marketplace_purchase,omitempty"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Effective_date string `json:"effective_date"`
}

// GeneratedType_Webhook_secret_scanning_alert_publicly_leaked represents the GeneratedType_Webhook_secret_scanning_alert_publicly_leaked schema from the OpenAPI specification
type GeneratedType_Webhook_secret_scanning_alert_publicly_leaked struct {
	Action string `json:"action"`
	Alert GeneratedType_Secret_scanning_alert_webhook `json:"alert"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType_Pull_request_review_request represents the GeneratedType_Pull_request_review_request schema from the OpenAPI specification
type GeneratedType_Pull_request_review_request struct {
	Teams []Team `json:"teams"`
	Users []GeneratedType_Simple_user `json:"users"`
}

// GeneratedType_Webhook_check_suite_requested represents the GeneratedType_Webhook_check_suite_requested schema from the OpenAPI specification
type GeneratedType_Webhook_check_suite_requested struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/checks/suites#get-a-check-suite).
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Secret_scanning_location_issue_comment represents the GeneratedType_Secret_scanning_location_issue_comment schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_issue_comment struct {
	Issue_comment_url string `json:"issue_comment_url"` // The API URL to get the issue comment where the secret was detected.
}

// GeneratedType_Diff_entry represents the GeneratedType_Diff_entry schema from the OpenAPI specification
type GeneratedType_Diff_entry struct {
	Patch string `json:"patch,omitempty"`
	Blob_url string `json:"blob_url"`
	Deletions int `json:"deletions"`
	Filename string `json:"filename"`
	Raw_url string `json:"raw_url"`
	Sha string `json:"sha"`
	Additions int `json:"additions"`
	Contents_url string `json:"contents_url"`
	Previous_filename string `json:"previous_filename,omitempty"`
	Status string `json:"status"`
	Changes int `json:"changes"`
}

// GeneratedType_Code_of_conduct represents the GeneratedType_Code_of_conduct schema from the OpenAPI specification
type GeneratedType_Code_of_conduct struct {
	Url string `json:"url"`
	Body string `json:"body,omitempty"`
	Html_url string `json:"html_url"`
	Key string `json:"key"`
	Name string `json:"name"`
}

// Label represents the Label schema from the OpenAPI specification
type Label struct {
	Color string `json:"color"` // 6-character hex code, without the leading #, identifying the color
	DefaultField bool `json:"default"` // Whether this label comes by default in a new repository.
	Description string `json:"description"` // Optional description of the label, such as its purpose.
	Id int64 `json:"id"` // Unique identifier for the label.
	Name string `json:"name"` // The name of the label.
	Node_id string `json:"node_id"`
	Url string `json:"url"` // URL for the label
}

// GeneratedType_Secret_scanning_location_discussion_title represents the GeneratedType_Secret_scanning_location_discussion_title schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location_discussion_title struct {
	Discussion_title_url string `json:"discussion_title_url"` // The URL to the discussion where the secret was detected.
}

// GeneratedType_Dependency_graph_spdx_sbom represents the GeneratedType_Dependency_graph_spdx_sbom schema from the OpenAPI specification
type GeneratedType_Dependency_graph_spdx_sbom struct {
	Sbom map[string]interface{} `json:"sbom"`
}

// GeneratedType_Commit_comparison represents the GeneratedType_Commit_comparison schema from the OpenAPI specification
type GeneratedType_Commit_comparison struct {
	Patch_url string `json:"patch_url"`
	Total_commits int `json:"total_commits"`
	Base_commit Commit `json:"base_commit"` // Commit
	Behind_by int `json:"behind_by"`
	Commits []Commit `json:"commits"`
	Files []GeneratedType_Diff_entry `json:"files,omitempty"`
	Permalink_url string `json:"permalink_url"`
	Status string `json:"status"`
	Html_url string `json:"html_url"`
	Url string `json:"url"`
	Ahead_by int `json:"ahead_by"`
	Diff_url string `json:"diff_url"`
	Merge_base_commit Commit `json:"merge_base_commit"` // Commit
}

// GeneratedType_Thread_subscription represents the GeneratedType_Thread_subscription schema from the OpenAPI specification
type GeneratedType_Thread_subscription struct {
	Reason string `json:"reason"`
	Repository_url string `json:"repository_url,omitempty"`
	Subscribed bool `json:"subscribed"`
	Thread_url string `json:"thread_url,omitempty"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Ignored bool `json:"ignored"`
}

// GeneratedType_Custom_deployment_rule_app represents the GeneratedType_Custom_deployment_rule_app schema from the OpenAPI specification
type GeneratedType_Custom_deployment_rule_app struct {
	Node_id string `json:"node_id"` // The node ID for the deployment protection rule integration.
	Slug string `json:"slug"` // The slugified name of the deployment protection rule integration.
	Id int `json:"id"` // The unique identifier of the deployment protection rule integration.
	Integration_url string `json:"integration_url"` // The URL for the endpoint to get details about the app.
}

// GeneratedType_Simple_check_suite represents the GeneratedType_Simple_check_suite schema from the OpenAPI specification
type GeneratedType_Simple_check_suite struct {
	Pull_requests []GeneratedType_Pull_request_minimal `json:"pull_requests,omitempty"`
	App Integration `json:"app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Head_sha string `json:"head_sha,omitempty"` // The SHA of the head commit that is being checked.
	Repository GeneratedType_Minimal_repository `json:"repository,omitempty"` // Minimal Repository
	Before string `json:"before,omitempty"`
	Conclusion string `json:"conclusion,omitempty"`
	Head_branch string `json:"head_branch,omitempty"`
	Id int `json:"id,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Status string `json:"status,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Url string `json:"url,omitempty"`
	After string `json:"after,omitempty"`
	Created_at string `json:"created_at,omitempty"`
}

// GeneratedType_File_commit represents the GeneratedType_File_commit schema from the OpenAPI specification
type GeneratedType_File_commit struct {
	Content map[string]interface{} `json:"content"`
	Commit map[string]interface{} `json:"commit"`
}

// GeneratedType_Webhook_discussion_labeled represents the GeneratedType_Webhook_discussion_labeled schema from the OpenAPI specification
type GeneratedType_Webhook_discussion_labeled struct {
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_team_add represents the GeneratedType_Webhook_team_add schema from the OpenAPI specification
type GeneratedType_Webhook_team_add struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Content_symlink represents the GeneratedType_Content_symlink schema from the OpenAPI specification
type GeneratedType_Content_symlink struct {
	Html_url string `json:"html_url"`
	Sha string `json:"sha"`
	Target string `json:"target"`
	Download_url string `json:"download_url"`
	Path string `json:"path"`
	Size int `json:"size"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Git_url string `json:"git_url"`
	Name string `json:"name"`
	Links map[string]interface{} `json:"_links"`
}

// GeneratedType_Nullable_scoped_installation represents the GeneratedType_Nullable_scoped_installation schema from the OpenAPI specification
type GeneratedType_Nullable_scoped_installation struct {
	Permissions GeneratedType_App_permissions `json:"permissions"` // The permissions granted to the user access token.
	Repositories_url string `json:"repositories_url"`
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Single_file_name string `json:"single_file_name"`
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Account GeneratedType_Simple_user `json:"account"` // A GitHub user.
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
}

// GeneratedType_Code_scanning_analysis_tool represents the GeneratedType_Code_scanning_analysis_tool schema from the OpenAPI specification
type GeneratedType_Code_scanning_analysis_tool struct {
	Name string `json:"name,omitempty"` // The name of the tool used to generate the code scanning analysis.
	Version string `json:"version,omitempty"` // The version of the tool used to generate the code scanning analysis.
	Guid string `json:"guid,omitempty"` // The GUID of the tool used to generate the code scanning analysis, if provided in the uploaded SARIF data.
}

// GeneratedType_Short_branch represents the GeneratedType_Short_branch schema from the OpenAPI specification
type GeneratedType_Short_branch struct {
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
	Protected bool `json:"protected"`
	Protection GeneratedType_Branch_protection `json:"protection,omitempty"` // Branch Protection
	Protection_url string `json:"protection_url,omitempty"`
}

// GeneratedType_Webhook_membership_removed represents the GeneratedType_Webhook_membership_removed schema from the OpenAPI specification
type GeneratedType_Webhook_membership_removed struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Member Webhooksuser `json:"member"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Team Webhooksteam `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Scope string `json:"scope"` // The scope of the membership. Currently, can only be `team`.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Sender map[string]interface{} `json:"sender"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType_Webhook_push represents the GeneratedType_Webhook_push schema from the OpenAPI specification
type GeneratedType_Webhook_push struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Ref string `json:"ref"` // The full git ref that was pushed. Example: `refs/heads/main` or `refs/tags/v3.14.1`.
	After string `json:"after"` // The SHA of the most recent commit on `ref` after the push.
	Head_commit map[string]interface{} `json:"head_commit"`
	Compare string `json:"compare"` // URL that shows the changes in this `ref` update, from the `before` commit to the `after` commit. For a newly created `ref` that is directly based on the default branch, this is the comparison between the head of the default branch and the `after` commit. Otherwise, this shows all commits until the `after` commit.
	Before string `json:"before"` // The SHA of the most recent commit on `ref` before the push.
	Pusher map[string]interface{} `json:"pusher"` // Metaproperties for Git author/committer information.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Commits []map[string]interface{} `json:"commits"` // An array of commit objects describing the pushed commits. (Pushed commits are all commits that are included in the `compare` between the `before` commit and the `after` commit.) The array includes a maximum of 2048 commits. If necessary, you can use the [Commits API](https://docs.github.com/rest/commits) to fetch additional commits.
	Deleted bool `json:"deleted"` // Whether this push deleted the `ref`.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository map[string]interface{} `json:"repository"` // A git repository
	Base_ref string `json:"base_ref"`
	Created bool `json:"created"` // Whether this push created the `ref`.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Forced bool `json:"forced"` // Whether this push was a force push of the `ref`.
}

// GeneratedType_Pull_request_review represents the GeneratedType_Pull_request_review schema from the OpenAPI specification
type GeneratedType_Pull_request_review struct {
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Commit_id string `json:"commit_id"` // A commit SHA for the review. If the commit object was garbage collected or forcibly deleted, then it no longer exists in Git and this value will be `null`.
	Body string `json:"body"` // The text of the review.
	Body_html string `json:"body_html,omitempty"`
	Id int64 `json:"id"` // Unique identifier of the review
	Pull_request_url string `json:"pull_request_url"`
	Submitted_at string `json:"submitted_at,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	User GeneratedType_Nullable_simple_user `json:"user"` // A GitHub user.
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	State string `json:"state"`
	Links map[string]interface{} `json:"_links"`
}

// GeneratedType_Nullable_minimal_repository represents the GeneratedType_Nullable_minimal_repository schema from the OpenAPI specification
type GeneratedType_Nullable_minimal_repository struct {
	Private bool `json:"private"`
	Language string `json:"language,omitempty"`
	Deployments_url string `json:"deployments_url"`
	Languages_url string `json:"languages_url"`
	Svn_url string `json:"svn_url,omitempty"`
	Mirror_url string `json:"mirror_url,omitempty"`
	Notifications_url string `json:"notifications_url"`
	Fork bool `json:"fork"`
	Full_name string `json:"full_name"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Open_issues int `json:"open_issues,omitempty"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Merges_url string `json:"merges_url"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Assignees_url string `json:"assignees_url"`
	Hooks_url string `json:"hooks_url"`
	Updated_at string `json:"updated_at,omitempty"`
	Url string `json:"url"`
	Git_commits_url string `json:"git_commits_url"`
	Trees_url string `json:"trees_url"`
	Comments_url string `json:"comments_url"`
	Open_issues_count int `json:"open_issues_count,omitempty"`
	Events_url string `json:"events_url"`
	Forks_url string `json:"forks_url"`
	Git_url string `json:"git_url,omitempty"`
	Network_count int `json:"network_count,omitempty"`
	Contents_url string `json:"contents_url"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Visibility string `json:"visibility,omitempty"`
	Stargazers_count int `json:"stargazers_count,omitempty"`
	Contributors_url string `json:"contributors_url"`
	Watchers int `json:"watchers,omitempty"`
	Default_branch string `json:"default_branch,omitempty"`
	Issue_events_url string `json:"issue_events_url"`
	Size int `json:"size,omitempty"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Watchers_count int `json:"watchers_count,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Clone_url string `json:"clone_url,omitempty"`
	Has_downloads bool `json:"has_downloads,omitempty"`
	Archive_url string `json:"archive_url"`
	Archived bool `json:"archived,omitempty"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Has_issues bool `json:"has_issues,omitempty"`
	Is_template bool `json:"is_template,omitempty"`
	Subscribers_url string `json:"subscribers_url"`
	Pushed_at string `json:"pushed_at,omitempty"`
	Blobs_url string `json:"blobs_url"`
	Forks int `json:"forks,omitempty"`
	Description string `json:"description"`
	Releases_url string `json:"releases_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Issues_url string `json:"issues_url"`
	Has_pages bool `json:"has_pages,omitempty"`
	License map[string]interface{} `json:"license,omitempty"`
	Statuses_url string `json:"statuses_url"`
	Labels_url string `json:"labels_url"`
	Commits_url string `json:"commits_url"`
	Html_url string `json:"html_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Git_tags_url string `json:"git_tags_url"`
	Tags_url string `json:"tags_url"`
	Subscription_url string `json:"subscription_url"`
	Teams_url string `json:"teams_url"`
	Disabled bool `json:"disabled,omitempty"`
	Pulls_url string `json:"pulls_url"`
	Git_refs_url string `json:"git_refs_url"`
	Topics []string `json:"topics,omitempty"`
	Has_wiki bool `json:"has_wiki,omitempty"`
	Downloads_url string `json:"downloads_url"`
	Name string `json:"name"`
	Milestones_url string `json:"milestones_url"`
	Has_projects bool `json:"has_projects,omitempty"`
	Ssh_url string `json:"ssh_url,omitempty"`
	Branches_url string `json:"branches_url"`
	Role_name string `json:"role_name,omitempty"`
	Node_id string `json:"node_id"`
	Owner GeneratedType_Simple_user `json:"owner"` // A GitHub user.
	Issue_comment_url string `json:"issue_comment_url"`
	Compare_url string `json:"compare_url"`
	Collaborators_url string `json:"collaborators_url"`
	Security_and_analysis GeneratedType_Security_and_analysis `json:"security_and_analysis,omitempty"`
	Id int64 `json:"id"`
	Forks_count int `json:"forks_count,omitempty"`
	Keys_url string `json:"keys_url"`
	Stargazers_url string `json:"stargazers_url"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Code_of_conduct GeneratedType_Code_of_conduct `json:"code_of_conduct,omitempty"` // Code Of Conduct
}

// Webhooksapprover represents the Webhooksapprover schema from the OpenAPI specification
type Webhooksapprover struct {
	Starred_url string `json:"starred_url,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Url string `json:"url,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Id int `json:"id,omitempty"`
	TypeField string `json:"type,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Login string `json:"login,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Gists_url string `json:"gists_url,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Organizations_url string `json:"organizations_url,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
}

// GeneratedType_Organization_create_issue_type represents the GeneratedType_Organization_create_issue_type schema from the OpenAPI specification
type GeneratedType_Organization_create_issue_type struct {
	Description string `json:"description,omitempty"` // Description of the issue type.
	Is_enabled bool `json:"is_enabled"` // Whether or not the issue type is enabled at the organization level.
	Name string `json:"name"` // Name of the issue type.
	Color string `json:"color,omitempty"` // Color for the issue type.
}

// GeneratedType_Webhook_team_added_to_repository represents the GeneratedType_Webhook_team_added_to_repository schema from the OpenAPI specification
type GeneratedType_Webhook_team_added_to_repository struct {
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
}

// GeneratedType_Simple_classroom_assignment represents the GeneratedType_Simple_classroom_assignment schema from the OpenAPI specification
type GeneratedType_Simple_classroom_assignment struct {
	Public_repo bool `json:"public_repo"` // Whether an accepted assignment creates a public repository.
	Students_are_repo_admins bool `json:"students_are_repo_admins"` // Whether students are admins on created repository on accepted assignment.
	Feedback_pull_requests_enabled bool `json:"feedback_pull_requests_enabled"` // Whether feedback pull request will be created on assignment acceptance.
	Title string `json:"title"` // Assignment title.
	TypeField string `json:"type"` // Whether it's a Group Assignment or Individual Assignment.
	Language string `json:"language"` // The programming language used in the assignment.
	Invitations_enabled bool `json:"invitations_enabled"` // Whether the invitation link is enabled. Visiting an enabled invitation link will accept the assignment.
	Classroom GeneratedType_Simple_classroom `json:"classroom"` // A GitHub Classroom classroom
	Slug string `json:"slug"` // Sluggified name of the assignment.
	Max_teams int `json:"max_teams,omitempty"` // The maximum allowable teams for the assignment.
	Editor string `json:"editor"` // The selected editor for the assignment.
	Passing int `json:"passing"` // The number of students that have passed the assignment.
	Submitted int `json:"submitted"` // The number of students that have submitted the assignment.
	Invite_link string `json:"invite_link"` // The link that a student can use to accept the assignment.
	Deadline string `json:"deadline"` // The time at which the assignment is due.
	Id int `json:"id"` // Unique identifier of the repository.
	Accepted int `json:"accepted"` // The number of students that have accepted the assignment.
	Max_members int `json:"max_members,omitempty"` // The maximum allowable members per team.
}

// GeneratedType_Webhook_deployment_protection_rule_requested represents the GeneratedType_Webhook_deployment_protection_rule_requested schema from the OpenAPI specification
type GeneratedType_Webhook_deployment_protection_rule_requested struct {
	Deployment_callback_url string `json:"deployment_callback_url,omitempty"` // The URL to review the deployment protection rule.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_requests []GeneratedType_Pull_request `json:"pull_requests,omitempty"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Environment string `json:"environment,omitempty"` // The name of the environment that has the deployment protection rule.
	Event string `json:"event,omitempty"` // The event that triggered the deployment protection rule.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action,omitempty"`
}

// GeneratedType_Webhook_projects_v2_status_update_created represents the GeneratedType_Webhook_projects_v2_status_update_created schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_status_update_created struct {
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_status_update GeneratedType_Projects_v2_status_update `json:"projects_v2_status_update"` // An status update belonging to a project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Code_scanning_variant_analysis_skipped_repo_group represents the GeneratedType_Code_scanning_variant_analysis_skipped_repo_group schema from the OpenAPI specification
type GeneratedType_Code_scanning_variant_analysis_skipped_repo_group struct {
	Repository_count int `json:"repository_count"` // The total number of repositories that were skipped for this reason.
	Repositories []GeneratedType_Code_scanning_variant_analysis_repository `json:"repositories"` // A list of repositories that were skipped. This list may not include all repositories that were skipped. This is only available when the repository was found and the user has access to it.
}

// Webhooksrelease1 represents the Webhooksrelease1 schema from the OpenAPI specification
type Webhooksrelease1 struct {
	Author map[string]interface{} `json:"author"`
	Target_commitish string `json:"target_commitish"` // Specifies the commitish value that determines where the Git tag is created from.
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Assets_url string `json:"assets_url"`
	Tag_name string `json:"tag_name"` // The name of the tag.
	Tarball_url string `json:"tarball_url"`
	Node_id string `json:"node_id"`
	Body string `json:"body"`
	Zipball_url string `json:"zipball_url"`
	Draft bool `json:"draft"` // Whether the release is a draft or published
	Published_at string `json:"published_at"`
	Assets []map[string]interface{} `json:"assets"`
	Discussion_url string `json:"discussion_url,omitempty"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Upload_url string `json:"upload_url"`
	Name string `json:"name"`
	Prerelease bool `json:"prerelease"` // Whether the release is identified as a prerelease or a full release.
}

// GeneratedType_Webhook_issue_comment_edited represents the GeneratedType_Webhook_issue_comment_edited schema from the OpenAPI specification
type GeneratedType_Webhook_issue_comment_edited struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Action string `json:"action"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Changes Webhookschanges `json:"changes"` // The changes to the comment.
	Comment Webhooksissuecomment `json:"comment"` // The [comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment) itself.
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) the comment belongs to.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_release_deleted represents the GeneratedType_Webhook_release_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_release_deleted struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_secret_scanning_alert_reopened represents the GeneratedType_Webhook_secret_scanning_alert_reopened schema from the OpenAPI specification
type GeneratedType_Webhook_secret_scanning_alert_reopened struct {
	Action string `json:"action"`
	Alert GeneratedType_Secret_scanning_alert_webhook `json:"alert"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType_Organization_invitation represents the GeneratedType_Organization_invitation schema from the OpenAPI specification
type GeneratedType_Organization_invitation struct {
	Team_count int `json:"team_count"`
	Failed_at string `json:"failed_at,omitempty"`
	Failed_reason string `json:"failed_reason,omitempty"`
	Id int64 `json:"id"`
	Role string `json:"role"`
	Invitation_teams_url string `json:"invitation_teams_url"`
	Node_id string `json:"node_id"`
	Login string `json:"login"`
	Created_at string `json:"created_at"`
	Email string `json:"email"`
	Invitation_source string `json:"invitation_source,omitempty"`
	Inviter GeneratedType_Simple_user `json:"inviter"` // A GitHub user.
}

// Dependency represents the Dependency schema from the OpenAPI specification
type Dependency struct {
	Dependencies []string `json:"dependencies,omitempty"` // Array of package-url (PURLs) of direct child dependencies.
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Package_url string `json:"package_url,omitempty"` // Package-url (PURL) of dependency. See https://github.com/package-url/purl-spec for more details.
	Relationship string `json:"relationship,omitempty"` // A notation of whether a dependency is requested directly by this manifest or is a dependency of another dependency.
	Scope string `json:"scope,omitempty"` // A notation of whether the dependency is required for the primary build artifact (runtime) or is only used for development. Future versions of this specification may allow for more granular scopes.
}

// GeneratedType_Team_project represents the GeneratedType_Team_project schema from the OpenAPI specification
type GeneratedType_Team_project struct {
	Updated_at string `json:"updated_at"`
	Private bool `json:"private,omitempty"` // Whether the project is private or not. Only present when owner is an organization.
	Node_id string `json:"node_id"`
	Organization_permission string `json:"organization_permission,omitempty"` // The organization permission for this project. Only present when owner is an organization.
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Number int `json:"number"`
	Columns_url string `json:"columns_url"`
	Permissions map[string]interface{} `json:"permissions"`
	Name string `json:"name"`
	Body string `json:"body"`
	Creator GeneratedType_Simple_user `json:"creator"` // A GitHub user.
	Owner_url string `json:"owner_url"`
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	State string `json:"state"`
}

// GeneratedType_Webhook_repository_publicized represents the GeneratedType_Webhook_repository_publicized schema from the OpenAPI specification
type GeneratedType_Webhook_repository_publicized struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_installation_suspend represents the GeneratedType_Webhook_installation_suspend schema from the OpenAPI specification
type GeneratedType_Webhook_installation_suspend struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_check_suite_rerequested represents the GeneratedType_Webhook_check_suite_rerequested schema from the OpenAPI specification
type GeneratedType_Webhook_check_suite_rerequested struct {
	Action string `json:"action"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/checks/suites#get-a-check-suite).
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_State_change_issue_event represents the GeneratedType_State_change_issue_event schema from the OpenAPI specification
type GeneratedType_State_change_issue_event struct {
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Id int `json:"id"`
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Url string `json:"url"`
	State_reason string `json:"state_reason,omitempty"`
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Created_at string `json:"created_at"`
	Event string `json:"event"`
}

// GeneratedType_Webhook_code_scanning_alert_reopened_by_user represents the GeneratedType_Webhook_code_scanning_alert_reopened_by_user schema from the OpenAPI specification
type GeneratedType_Webhook_code_scanning_alert_reopened_by_user struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Repository_rule_params_required_reviewer_configuration represents the GeneratedType_Repository_rule_params_required_reviewer_configuration schema from the OpenAPI specification
type GeneratedType_Repository_rule_params_required_reviewer_configuration struct {
	File_patterns []string `json:"file_patterns"` // Array of file patterns. Pull requests which change matching files must be approved by the specified team. File patterns use the same syntax as `.gitignore` files.
	Minimum_approvals int `json:"minimum_approvals"` // Minimum number of approvals required from the specified team. If set to zero, the team will be added to the pull request but approval is optional.
	Reviewer GeneratedType_Repository_rule_params_reviewer `json:"reviewer"` // A required reviewing team
}

// GeneratedType_Webhook_issues_unlocked represents the GeneratedType_Webhook_issues_unlocked schema from the OpenAPI specification
type GeneratedType_Webhook_issues_unlocked struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_Webhook_repository_vulnerability_alert_reopen represents the GeneratedType_Webhook_repository_vulnerability_alert_reopen schema from the OpenAPI specification
type GeneratedType_Webhook_repository_vulnerability_alert_reopen struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert Webhooksalert `json:"alert"` // The security alert of the vulnerable dependency.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType_App_permissions represents the GeneratedType_App_permissions schema from the OpenAPI specification
type GeneratedType_App_permissions struct {
	Administration string `json:"administration,omitempty"` // The level of permission to grant the access token for repository creation, deletion, settings, teams, and collaborators creation.
	Environments string `json:"environments,omitempty"` // The level of permission to grant the access token for managing repository environments.
	Organization_copilot_seat_management string `json:"organization_copilot_seat_management,omitempty"` // The level of permission to grant the access token for managing access to GitHub Copilot for members of an organization with a Copilot Business subscription. This property is in public preview and is subject to change.
	Organization_self_hosted_runners string `json:"organization_self_hosted_runners,omitempty"` // The level of permission to grant the access token to view and manage GitHub Actions self-hosted runners available to an organization.
	Pull_requests string `json:"pull_requests,omitempty"` // The level of permission to grant the access token for pull requests and related comments, assignees, labels, milestones, and merges.
	Organization_events string `json:"organization_events,omitempty"` // The level of permission to grant the access token to view events triggered by an activity in an organization.
	Security_events string `json:"security_events,omitempty"` // The level of permission to grant the access token to view and manage security events like code scanning alerts.
	Repository_hooks string `json:"repository_hooks,omitempty"` // The level of permission to grant the access token to manage the post-receive hooks for a repository.
	Organization_custom_properties string `json:"organization_custom_properties,omitempty"` // The level of permission to grant the access token for custom property management.
	Statuses string `json:"statuses,omitempty"` // The level of permission to grant the access token for commit statuses.
	Followers string `json:"followers,omitempty"` // The level of permission to grant the access token to manage the followers belonging to a user.
	Email_addresses string `json:"email_addresses,omitempty"` // The level of permission to grant the access token to manage the email addresses belonging to a user.
	Repository_projects string `json:"repository_projects,omitempty"` // The level of permission to grant the access token to manage repository projects, columns, and cards.
	Organization_custom_roles string `json:"organization_custom_roles,omitempty"` // The level of permission to grant the access token for custom repository roles management.
	Secrets string `json:"secrets,omitempty"` // The level of permission to grant the access token to manage repository secrets.
	Organization_custom_org_roles string `json:"organization_custom_org_roles,omitempty"` // The level of permission to grant the access token for custom organization roles management.
	Packages string `json:"packages,omitempty"` // The level of permission to grant the access token for packages published to GitHub Packages.
	Organization_announcement_banners string `json:"organization_announcement_banners,omitempty"` // The level of permission to grant the access token to view and manage announcement banners for an organization.
	Metadata string `json:"metadata,omitempty"` // The level of permission to grant the access token to search repositories, list collaborators, and access repository metadata.
	Secret_scanning_alerts string `json:"secret_scanning_alerts,omitempty"` // The level of permission to grant the access token to view and manage secret scanning alerts.
	Workflows string `json:"workflows,omitempty"` // The level of permission to grant the access token to update GitHub Actions workflow files.
	Team_discussions string `json:"team_discussions,omitempty"` // The level of permission to grant the access token to manage team discussions and related comments.
	Interaction_limits string `json:"interaction_limits,omitempty"` // The level of permission to grant the access token to view and manage interaction limits on a repository.
	Members string `json:"members,omitempty"` // The level of permission to grant the access token for organization teams and members.
	Vulnerability_alerts string `json:"vulnerability_alerts,omitempty"` // The level of permission to grant the access token to manage Dependabot alerts.
	Git_ssh_keys string `json:"git_ssh_keys,omitempty"` // The level of permission to grant the access token to manage git SSH keys.
	Organization_secrets string `json:"organization_secrets,omitempty"` // The level of permission to grant the access token to manage organization secrets.
	Organization_personal_access_token_requests string `json:"organization_personal_access_token_requests,omitempty"` // The level of permission to grant the access token for viewing and managing fine-grained personal access tokens that have been approved by an organization.
	Starring string `json:"starring,omitempty"` // The level of permission to grant the access token to list and manage repositories a user is starring.
	Dependabot_secrets string `json:"dependabot_secrets,omitempty"` // The level of permission to grant the access token to manage Dependabot secrets.
	Organization_plan string `json:"organization_plan,omitempty"` // The level of permission to grant the access token for viewing an organization's plan.
	Checks string `json:"checks,omitempty"` // The level of permission to grant the access token for checks on code.
	Deployments string `json:"deployments,omitempty"` // The level of permission to grant the access token for deployments and deployment statuses.
	Organization_user_blocking string `json:"organization_user_blocking,omitempty"` // The level of permission to grant the access token to view and manage users blocked by the organization.
	Pages string `json:"pages,omitempty"` // The level of permission to grant the access token to retrieve Pages statuses, configuration, and builds, as well as create new builds.
	Organization_administration string `json:"organization_administration,omitempty"` // The level of permission to grant the access token to manage access to an organization.
	Codespaces string `json:"codespaces,omitempty"` // The level of permission to grant the access token to create, edit, delete, and list Codespaces.
	Repository_custom_properties string `json:"repository_custom_properties,omitempty"` // The level of permission to grant the access token to view and edit custom properties for a repository, when allowed by the property.
	Actions string `json:"actions,omitempty"` // The level of permission to grant the access token for GitHub Actions workflows, workflow runs, and artifacts.
	Gpg_keys string `json:"gpg_keys,omitempty"` // The level of permission to grant the access token to view and manage GPG keys belonging to a user.
	Organization_personal_access_tokens string `json:"organization_personal_access_tokens,omitempty"` // The level of permission to grant the access token for viewing and managing fine-grained personal access token requests to an organization.
	Organization_packages string `json:"organization_packages,omitempty"` // The level of permission to grant the access token for organization packages published to GitHub Packages.
	Issues string `json:"issues,omitempty"` // The level of permission to grant the access token for issues and related comments, assignees, labels, and milestones.
	Contents string `json:"contents,omitempty"` // The level of permission to grant the access token for repository contents, commits, branches, downloads, releases, and merges.
	Organization_projects string `json:"organization_projects,omitempty"` // The level of permission to grant the access token to manage organization projects and projects public preview (where available).
	Profile string `json:"profile,omitempty"` // The level of permission to grant the access token to manage the profile settings belonging to a user.
	Single_file string `json:"single_file,omitempty"` // The level of permission to grant the access token to manage just a single file.
	Organization_hooks string `json:"organization_hooks,omitempty"` // The level of permission to grant the access token to manage the post-receive hooks for an organization.
}

// GeneratedType_Release_notes_content represents the GeneratedType_Release_notes_content schema from the OpenAPI specification
type GeneratedType_Release_notes_content struct {
	Body string `json:"body"` // The generated body describing the contents of the release supporting markdown formatting
	Name string `json:"name"` // The generated name of the release
}

// GeneratedType_Webhook_dependabot_alert_dismissed represents the GeneratedType_Webhook_dependabot_alert_dismissed schema from the OpenAPI specification
type GeneratedType_Webhook_dependabot_alert_dismissed struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType_Dependabot_alert `json:"alert"` // A Dependabot alert.
}

// GeneratedType_Webhook_branch_protection_configuration_disabled represents the GeneratedType_Webhook_branch_protection_configuration_disabled schema from the OpenAPI specification
type GeneratedType_Webhook_branch_protection_configuration_disabled struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Actions_cache_usage_org_enterprise represents the GeneratedType_Actions_cache_usage_org_enterprise schema from the OpenAPI specification
type GeneratedType_Actions_cache_usage_org_enterprise struct {
	Total_active_caches_count int `json:"total_active_caches_count"` // The count of active caches across all repositories of an enterprise or an organization.
	Total_active_caches_size_in_bytes int `json:"total_active_caches_size_in_bytes"` // The total size in bytes of all active cache items across all repositories of an enterprise or an organization.
}

// GeneratedType_Auto_merge represents the GeneratedType_Auto_merge schema from the OpenAPI specification
type GeneratedType_Auto_merge struct {
	Enabled_by GeneratedType_Simple_user `json:"enabled_by"` // A GitHub user.
	Merge_method string `json:"merge_method"` // The merge method to use.
	Commit_message string `json:"commit_message"` // Commit message for the merge commit.
	Commit_title string `json:"commit_title"` // Title for the merge commit message.
}

// GeneratedType_Webhook_sub_issues_parent_issue_removed represents the GeneratedType_Webhook_sub_issues_parent_issue_removed schema from the OpenAPI specification
type GeneratedType_Webhook_sub_issues_parent_issue_removed struct {
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
	Parent_issue_repo Repository `json:"parent_issue_repo"` // A repository on GitHub.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
	Action string `json:"action"`
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
}

// Webhookssponsorship represents the Webhookssponsorship schema from the OpenAPI specification
type Webhookssponsorship struct {
	Sponsorable map[string]interface{} `json:"sponsorable"`
	Tier map[string]interface{} `json:"tier"` // The `tier_changed` and `pending_tier_change` will include the original tier before the change or pending change. For more information, see the pending tier change payload.
	Created_at string `json:"created_at"`
	Maintainer map[string]interface{} `json:"maintainer,omitempty"`
	Node_id string `json:"node_id"`
	Privacy_level string `json:"privacy_level"`
	Sponsor map[string]interface{} `json:"sponsor"`
}

// GeneratedType_Webhook_pull_request_review_dismissed represents the GeneratedType_Webhook_pull_request_review_dismissed schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_review_dismissed struct {
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Review map[string]interface{} `json:"review"` // The review that was affected.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Webhook_meta_deleted represents the GeneratedType_Webhook_meta_deleted schema from the OpenAPI specification
type GeneratedType_Webhook_meta_deleted struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Nullable_repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Hook map[string]interface{} `json:"hook"` // The deleted webhook. This will contain different keys based on the type of webhook it is: repository, organization, business, app, or GitHub Marketplace.
	Hook_id int `json:"hook_id"` // The id of the modified webhook.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Secret_scanning_location represents the GeneratedType_Secret_scanning_location schema from the OpenAPI specification
type GeneratedType_Secret_scanning_location struct {
	Details interface{} `json:"details,omitempty"`
	TypeField string `json:"type,omitempty"` // The location type. Because secrets may be found in different types of resources (ie. code, comments, issues, pull requests, discussions), this field identifies the type of resource where the secret was found.
}

// GeneratedType_Repo_codespaces_secret represents the GeneratedType_Repo_codespaces_secret schema from the OpenAPI specification
type GeneratedType_Repo_codespaces_secret struct {
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Updated_at string `json:"updated_at"`
}

// GeneratedType_Nullable_license_simple represents the GeneratedType_Nullable_license_simple schema from the OpenAPI specification
type GeneratedType_Nullable_license_simple struct {
	Html_url string `json:"html_url,omitempty"`
	Key string `json:"key"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Spdx_id string `json:"spdx_id"`
	Url string `json:"url"`
}

// GeneratedType_Webhook_pull_request_edited represents the GeneratedType_Webhook_pull_request_edited schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_edited struct {
	Pull_request GeneratedType_Pull_request_webhook `json:"pull_request"`
	Changes map[string]interface{} `json:"changes"` // The changes to the comment if the action was `edited`.
	Number int `json:"number"` // The pull request number.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Pages_deployment_status represents the GeneratedType_Pages_deployment_status schema from the OpenAPI specification
type GeneratedType_Pages_deployment_status struct {
	Status string `json:"status,omitempty"` // The current status of the deployment.
}

// Webhooksuser represents the Webhooksuser schema from the OpenAPI specification
type Webhooksuser struct {
	Login string `json:"login"`
	Email string `json:"email,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
	Gists_url string `json:"gists_url,omitempty"`
	Url string `json:"url,omitempty"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Name string `json:"name,omitempty"`
	Organizations_url string `json:"organizations_url,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	TypeField string `json:"type,omitempty"`
	Id int64 `json:"id"`
	Node_id string `json:"node_id,omitempty"`
	Events_url string `json:"events_url,omitempty"`
}

// GeneratedType_Contributor_activity represents the GeneratedType_Contributor_activity schema from the OpenAPI specification
type GeneratedType_Contributor_activity struct {
	Weeks []map[string]interface{} `json:"weeks"`
	Author GeneratedType_Nullable_simple_user `json:"author"` // A GitHub user.
	Total int `json:"total"`
}

// GeneratedType_Code_scanning_organization_alert_items represents the GeneratedType_Code_scanning_organization_alert_items schema from the OpenAPI specification
type GeneratedType_Code_scanning_organization_alert_items struct {
	Most_recent_instance GeneratedType_Code_scanning_alert_instance `json:"most_recent_instance"`
	Number int `json:"number"` // The security alert number.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Rule GeneratedType_Code_scanning_alert_rule_summary `json:"rule"`
	Url string `json:"url"` // The REST API URL of the alert resource.
	State string `json:"state"` // State of a code scanning alert.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Dismissal_approved_by GeneratedType_Nullable_simple_user `json:"dismissal_approved_by,omitempty"` // A GitHub user.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Tool GeneratedType_Code_scanning_analysis_tool `json:"tool"`
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Repository GeneratedType_Simple_repository `json:"repository"` // A GitHub repository.
	Dismissed_by GeneratedType_Nullable_simple_user `json:"dismissed_by"` // A GitHub user.
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
}

// GeneratedType_Actions_get_default_workflow_permissions represents the GeneratedType_Actions_get_default_workflow_permissions schema from the OpenAPI specification
type GeneratedType_Actions_get_default_workflow_permissions struct {
	Can_approve_pull_request_reviews bool `json:"can_approve_pull_request_reviews"` // Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
	Default_workflow_permissions string `json:"default_workflow_permissions"` // The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
}

// GeneratedType_Codespaces_org_secret represents the GeneratedType_Codespaces_org_secret schema from the OpenAPI specification
type GeneratedType_Codespaces_org_secret struct {
	Created_at string `json:"created_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the secret
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // The API URL at which the list of repositories this secret is visible to can be retrieved
	Updated_at string `json:"updated_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Visibility string `json:"visibility"` // The type of repositories in the organization that the secret is visible to
}

// Metadata represents the Metadata schema from the OpenAPI specification
type Metadata struct {
}

// GeneratedType_Code_scanning_default_setup_update_response represents the GeneratedType_Code_scanning_default_setup_update_response schema from the OpenAPI specification
type GeneratedType_Code_scanning_default_setup_update_response struct {
	Run_id int `json:"run_id,omitempty"` // ID of the corresponding run.
	Run_url string `json:"run_url,omitempty"` // URL of the corresponding run.
}

// GeneratedType_Pages_source_hash represents the GeneratedType_Pages_source_hash schema from the OpenAPI specification
type GeneratedType_Pages_source_hash struct {
	Branch string `json:"branch"`
	Path string `json:"path"`
}

// GeneratedType_Nullable_collaborator represents the GeneratedType_Nullable_collaborator schema from the OpenAPI specification
type GeneratedType_Nullable_collaborator struct {
	Html_url string `json:"html_url"`
	Avatar_url string `json:"avatar_url"`
	Events_url string `json:"events_url"`
	Subscriptions_url string `json:"subscriptions_url"`
	Following_url string `json:"following_url"`
	Received_events_url string `json:"received_events_url"`
	Url string `json:"url"`
	TypeField string `json:"type"`
	Email string `json:"email,omitempty"`
	Gists_url string `json:"gists_url"`
	Gravatar_id string `json:"gravatar_id"`
	Login string `json:"login"`
	Name string `json:"name,omitempty"`
	Site_admin bool `json:"site_admin"`
	Starred_url string `json:"starred_url"`
	Repos_url string `json:"repos_url"`
	Followers_url string `json:"followers_url"`
	Id int64 `json:"id"`
	Node_id string `json:"node_id"`
	Organizations_url string `json:"organizations_url"`
	Role_name string `json:"role_name"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
}

// GeneratedType_Webhook_branch_protection_rule_edited represents the GeneratedType_Webhook_branch_protection_rule_edited schema from the OpenAPI specification
type GeneratedType_Webhook_branch_protection_rule_edited struct {
	Changes map[string]interface{} `json:"changes,omitempty"` // If the action was `edited`, the changes to the rule.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Rule Webhooksrule `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Team_repository represents the GeneratedType_Team_repository schema from the OpenAPI specification
type GeneratedType_Team_repository struct {
	Ssh_url string `json:"ssh_url"`
	Issue_events_url string `json:"issue_events_url"`
	Node_id string `json:"node_id"`
	Contents_url string `json:"contents_url"`
	Svn_url string `json:"svn_url"`
	Blobs_url string `json:"blobs_url"`
	Forks_url string `json:"forks_url"`
	Name string `json:"name"` // The name of the repository.
	Language string `json:"language"`
	Assignees_url string `json:"assignees_url"`
	Git_commits_url string `json:"git_commits_url"`
	Milestones_url string `json:"milestones_url"`
	Updated_at string `json:"updated_at"`
	Fork bool `json:"fork"`
	Hooks_url string `json:"hooks_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Languages_url string `json:"languages_url"`
	Teams_url string `json:"teams_url"`
	Network_count int `json:"network_count,omitempty"`
	Owner GeneratedType_Nullable_simple_user `json:"owner"` // A GitHub user.
	Private bool `json:"private"` // Whether the repository is private or public.
	Git_refs_url string `json:"git_refs_url"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Deployments_url string `json:"deployments_url"`
	Statuses_url string `json:"statuses_url"`
	Notifications_url string `json:"notifications_url"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Size int `json:"size"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Clone_url string `json:"clone_url"`
	Contributors_url string `json:"contributors_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Subscribers_url string `json:"subscribers_url"`
	Events_url string `json:"events_url"`
	Role_name string `json:"role_name,omitempty"`
	Topics []string `json:"topics,omitempty"`
	Branches_url string `json:"branches_url"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Stargazers_url string `json:"stargazers_url"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Git_tags_url string `json:"git_tags_url"`
	Trees_url string `json:"trees_url"`
	Has_pages bool `json:"has_pages"`
	Issues_url string `json:"issues_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Description string `json:"description"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Merges_url string `json:"merges_url"`
	Releases_url string `json:"releases_url"`
	Open_issues int `json:"open_issues"`
	Html_url string `json:"html_url"`
	Forks_count int `json:"forks_count"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Archive_url string `json:"archive_url"`
	Labels_url string `json:"labels_url"`
	Pushed_at string `json:"pushed_at"`
	Pulls_url string `json:"pulls_url"`
	Homepage string `json:"homepage"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Tags_url string `json:"tags_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Mirror_url string `json:"mirror_url"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Id int `json:"id"` // Unique identifier of the repository
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Downloads_url string `json:"downloads_url"`
	Watchers int `json:"watchers"`
	Open_issues_count int `json:"open_issues_count"`
	Watchers_count int `json:"watchers_count"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Commits_url string `json:"commits_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Compare_url string `json:"compare_url"`
	License GeneratedType_Nullable_license_simple `json:"license"` // License Simple
	Git_url string `json:"git_url"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Full_name string `json:"full_name"`
	Forks int `json:"forks"`
	Keys_url string `json:"keys_url"`
	Stargazers_count int `json:"stargazers_count"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Collaborators_url string `json:"collaborators_url"`
	Comments_url string `json:"comments_url"`
	Subscription_url string `json:"subscription_url"`
}

// GeneratedType_Copilot_dotcom_chat represents the GeneratedType_Copilot_dotcom_chat schema from the OpenAPI specification
type GeneratedType_Copilot_dotcom_chat struct {
	Models []map[string]interface{} `json:"models,omitempty"` // List of model metrics for a custom models and the default model.
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // Total number of users who prompted Copilot Chat on github.com at least once.
}

// GeneratedType_Webhook_project_closed represents the GeneratedType_Webhook_project_closed schema from the OpenAPI specification
type GeneratedType_Webhook_project_closed struct {
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_check_run_created represents the GeneratedType_Webhook_check_run_created schema from the OpenAPI specification
type GeneratedType_Webhook_check_run_created struct {
	Action string `json:"action,omitempty"`
	Check_run GeneratedType_Check_run_with_simple_check_suite `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Billing_usage_report_user represents the GeneratedType_Billing_usage_report_user schema from the OpenAPI specification
type GeneratedType_Billing_usage_report_user struct {
	Usageitems []map[string]interface{} `json:"usageItems,omitempty"`
}

// Webhooksalert represents the Webhooksalert schema from the OpenAPI specification
type Webhooksalert struct {
	Number int `json:"number"`
	State string `json:"state"`
	Affected_range string `json:"affected_range"`
	Created_at string `json:"created_at"`
	Dismissed_at string `json:"dismissed_at,omitempty"`
	Dismiss_reason string `json:"dismiss_reason,omitempty"`
	Node_id string `json:"node_id"`
	Affected_package_name string `json:"affected_package_name"`
	Ghsa_id string `json:"ghsa_id"`
	Fixed_at string `json:"fixed_at,omitempty"`
	Id int `json:"id"`
	External_identifier string `json:"external_identifier"`
	Dismisser map[string]interface{} `json:"dismisser,omitempty"`
	External_reference string `json:"external_reference"`
	Fix_reason string `json:"fix_reason,omitempty"`
	Severity string `json:"severity"`
	Fixed_in string `json:"fixed_in,omitempty"`
}

// GeneratedType_Webhook_projects_v2_status_update_edited represents the GeneratedType_Webhook_projects_v2_status_update_edited schema from the OpenAPI specification
type GeneratedType_Webhook_projects_v2_status_update_edited struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_status_update GeneratedType_Projects_v2_status_update `json:"projects_v2_status_update"` // An status update belonging to a project
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Webhook_branch_protection_configuration_enabled represents the GeneratedType_Webhook_branch_protection_configuration_enabled schema from the OpenAPI specification
type GeneratedType_Webhook_branch_protection_configuration_enabled struct {
	Action string `json:"action"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}

// GeneratedType_Organization_actions_variable represents the GeneratedType_Organization_actions_variable schema from the OpenAPI specification
type GeneratedType_Organization_actions_variable struct {
	Created_at string `json:"created_at"` // The date and time at which the variable was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the variable.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"` // The date and time at which the variable was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Value string `json:"value"` // The value of the variable.
	Visibility string `json:"visibility"` // Visibility of a variable
}

// GeneratedType_Timeline_unassigned_issue_event represents the GeneratedType_Timeline_unassigned_issue_event schema from the OpenAPI specification
type GeneratedType_Timeline_unassigned_issue_event struct {
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Url string `json:"url"`
	Assignee GeneratedType_Simple_user `json:"assignee"` // A GitHub user.
	Id int `json:"id"`
	Performed_via_github_app GeneratedType_Nullable_integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType_Simple_user `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
}

// GeneratedType_Rule_suite represents the GeneratedType_Rule_suite schema from the OpenAPI specification
type GeneratedType_Rule_suite struct {
	After_sha string `json:"after_sha,omitempty"` // The last commit sha in the push evaluation.
	Evaluation_result string `json:"evaluation_result,omitempty"` // The result of the rule evaluations for rules with the `active` and `evaluate` enforcement statuses, demonstrating whether rules would pass or fail if all rules in the rule suite were `active`. Null if no rules with `evaluate` enforcement status were run.
	Repository_id int `json:"repository_id,omitempty"` // The ID of the repository associated with the rule evaluation.
	Ref string `json:"ref,omitempty"` // The ref name that the evaluation ran on.
	Result string `json:"result,omitempty"` // The result of the rule evaluations for rules with the `active` enforcement status.
	Rule_evaluations []map[string]interface{} `json:"rule_evaluations,omitempty"` // Details on the evaluated rules.
	Actor_id int `json:"actor_id,omitempty"` // The number that identifies the user.
	Before_sha string `json:"before_sha,omitempty"` // The first commit sha before the push evaluation.
	Id int `json:"id,omitempty"` // The unique identifier of the rule insight.
	Pushed_at string `json:"pushed_at,omitempty"`
	Repository_name string `json:"repository_name,omitempty"` // The name of the repository without the `.git` extension.
	Actor_name string `json:"actor_name,omitempty"` // The handle for the GitHub user account.
}

// GeneratedType_Team_organization represents the GeneratedType_Team_organization schema from the OpenAPI specification
type GeneratedType_Team_organization struct {
	TypeField string `json:"type"`
	Events_url string `json:"events_url"`
	Node_id string `json:"node_id"`
	Members_can_create_public_pages bool `json:"members_can_create_public_pages,omitempty"`
	Following int `json:"following"`
	Members_can_create_pages bool `json:"members_can_create_pages,omitempty"`
	Email string `json:"email,omitempty"`
	Public_members_url string `json:"public_members_url"`
	Is_verified bool `json:"is_verified,omitempty"`
	Members_can_create_private_repositories bool `json:"members_can_create_private_repositories,omitempty"`
	Members_can_create_private_pages bool `json:"members_can_create_private_pages,omitempty"`
	Updated_at string `json:"updated_at"`
	Members_can_create_public_repositories bool `json:"members_can_create_public_repositories,omitempty"`
	Members_can_create_repositories bool `json:"members_can_create_repositories,omitempty"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Members_url string `json:"members_url"`
	Has_organization_projects bool `json:"has_organization_projects"`
	Members_can_fork_private_repositories bool `json:"members_can_fork_private_repositories,omitempty"`
	Id int `json:"id"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Members_can_create_internal_repositories bool `json:"members_can_create_internal_repositories,omitempty"`
	Has_repository_projects bool `json:"has_repository_projects"`
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Issues_url string `json:"issues_url"`
	Billing_email string `json:"billing_email,omitempty"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Followers int `json:"followers"`
	Private_gists int `json:"private_gists,omitempty"`
	Url string `json:"url"`
	Public_gists int `json:"public_gists"`
	Blog string `json:"blog,omitempty"`
	Html_url string `json:"html_url"`
	Name string `json:"name,omitempty"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Default_repository_permission string `json:"default_repository_permission,omitempty"`
	Login string `json:"login"`
	Hooks_url string `json:"hooks_url"`
	Members_allowed_repository_creation_type string `json:"members_allowed_repository_creation_type,omitempty"`
	Repos_url string `json:"repos_url"`
	Archived_at string `json:"archived_at"`
	Description string `json:"description"`
	Location string `json:"location,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Created_at string `json:"created_at"`
	Collaborators int `json:"collaborators,omitempty"`
	Two_factor_requirement_enabled bool `json:"two_factor_requirement_enabled,omitempty"`
	Public_repos int `json:"public_repos"`
	Company string `json:"company,omitempty"`
}

// GeneratedType_Webhook_page_build represents the GeneratedType_Webhook_page_build schema from the OpenAPI specification
type GeneratedType_Webhook_page_build struct {
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Build map[string]interface{} `json:"build"` // The [List GitHub Pages builds](https://docs.github.com/rest/pages/pages#list-github-pages-builds) itself.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Id int `json:"id"`
}

// GeneratedType_Webhook_dependabot_alert_created represents the GeneratedType_Webhook_dependabot_alert_created schema from the OpenAPI specification
type GeneratedType_Webhook_dependabot_alert_created struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType_Dependabot_alert `json:"alert"` // A Dependabot alert.
}

// GeneratedType_Webhook_secret_scanning_alert_location_created_form_encoded represents the GeneratedType_Webhook_secret_scanning_alert_location_created_form_encoded schema from the OpenAPI specification
type GeneratedType_Webhook_secret_scanning_alert_location_created_form_encoded struct {
	Payload string `json:"payload"` // A URL-encoded string of the secret_scanning_alert_location.created JSON payload. The decoded payload is a JSON object.
}

// GeneratedType_Webhook_issues_unassigned represents the GeneratedType_Webhook_issues_unassigned schema from the OpenAPI specification
type GeneratedType_Webhook_issues_unassigned struct {
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action that was performed.
	Assignee Webhooksusermannequin `json:"assignee,omitempty"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType_Content_traffic represents the GeneratedType_Content_traffic schema from the OpenAPI specification
type GeneratedType_Content_traffic struct {
	Count int `json:"count"`
	Path string `json:"path"`
	Title string `json:"title"`
	Uniques int `json:"uniques"`
}

// Project represents the Project schema from the OpenAPI specification
type Project struct {
	State string `json:"state"` // State of the project; either 'open' or 'closed'
	Id int `json:"id"`
	Name string `json:"name"` // Name of the project
	Updated_at string `json:"updated_at"`
	Body string `json:"body"` // Body of the project
	Html_url string `json:"html_url"`
	Url string `json:"url"`
	Node_id string `json:"node_id"`
	Organization_permission string `json:"organization_permission,omitempty"` // The baseline permission that all organization members have on this project. Only present if owner is an organization.
	Creator GeneratedType_Nullable_simple_user `json:"creator"` // A GitHub user.
	Owner_url string `json:"owner_url"`
	Private bool `json:"private,omitempty"` // Whether or not this project can be seen by everyone. Only present if owner is an organization.
	Columns_url string `json:"columns_url"`
	Created_at string `json:"created_at"`
	Number int `json:"number"`
}

// GeneratedType_Webhook_config represents the GeneratedType_Webhook_config schema from the OpenAPI specification
type GeneratedType_Webhook_config struct {
	Content_type string `json:"content_type,omitempty"` // The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
	Insecure_ssl string `json:"insecure_ssl,omitempty"`
	Secret string `json:"secret,omitempty"` // If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
	Url string `json:"url,omitempty"` // The URL to which the payloads will be delivered.
}

// Webhooksissuecomment represents the Webhooksissuecomment schema from the OpenAPI specification
type Webhooksissuecomment struct {
	Id int64 `json:"id"` // Unique identifier of the issue comment
	Node_id string `json:"node_id"`
	Body string `json:"body"` // Contents of the issue comment
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Issue_url string `json:"issue_url"`
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Reactions map[string]interface{} `json:"reactions"`
	Url string `json:"url"` // URL for the issue comment
	User map[string]interface{} `json:"user"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Updated_at string `json:"updated_at"`
}

// GeneratedType_Webhook_pull_request_auto_merge_disabled represents the GeneratedType_Webhook_pull_request_auto_merge_disabled schema from the OpenAPI specification
type GeneratedType_Webhook_pull_request_auto_merge_disabled struct {
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Reason string `json:"reason"`
	Number int `json:"number"`
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType_Webhook_org_block_blocked represents the GeneratedType_Webhook_org_block_blocked schema from the OpenAPI specification
type GeneratedType_Webhook_org_block_blocked struct {
	Blocked_user Webhooksuser `json:"blocked_user"`
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType_Webhook_repository_vulnerability_alert_resolve represents the GeneratedType_Webhook_repository_vulnerability_alert_resolve schema from the OpenAPI specification
type GeneratedType_Webhook_repository_vulnerability_alert_resolve struct {
	Action string `json:"action"`
	Alert map[string]interface{} `json:"alert"` // The security alert of the vulnerable dependency.
	Enterprise GeneratedType_Enterprise_webhooks `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType_Simple_installation `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType_Organization_simple_webhooks `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType_Repository_webhooks `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType_Simple_user `json:"sender"` // A GitHub user.
}
