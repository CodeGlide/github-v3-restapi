package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Webhooksproject represents the Webhooksproject schema from the OpenAPI specification
type Webhooksproject struct {
	Body string `json:"body"` // Body of the project
	Updated_at string `json:"updated_at"`
	Columns_url string `json:"columns_url"`
	Creator map[string]interface{} `json:"creator"`
	Number int `json:"number"`
	State string `json:"state"` // State of the project; either 'open' or 'closed'
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Name string `json:"name"` // Name of the project
	Node_id string `json:"node_id"`
	Owner_url string `json:"owner_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Label Webhookslabel `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // The changes to the label if the action was `edited`.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_ruleset GeneratedType `json:"repository_ruleset"` // A set of rules to apply when specified conditions are met.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Color string `json:"color"`
	DefaultField bool `json:"default"`
	Description string `json:"description"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Score float64 `json:"score"`
	Url string `json:"url"`
	Name string `json:"name"`
}

// Team represents the Team schema from the OpenAPI specification
type Team struct {
	Parent GeneratedType `json:"parent"` // Groups of organization members that gives permissions on specified repositories.
	Permission string `json:"permission"`
	Repositories_url string `json:"repositories_url"`
	Name string `json:"name"`
	Privacy string `json:"privacy,omitempty"`
	Slug string `json:"slug"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Description string `json:"description"`
	Members_url string `json:"members_url"`
	Notification_setting string `json:"notification_setting,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Forks_url string `json:"forks_url"`
	Has_pages bool `json:"has_pages,omitempty"`
	Forks_count int `json:"forks_count,omitempty"`
	Labels_url string `json:"labels_url"`
	Events_url string `json:"events_url"`
	Issue_events_url string `json:"issue_events_url"`
	Issues_url string `json:"issues_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Has_downloads bool `json:"has_downloads,omitempty"`
	Pushed_at string `json:"pushed_at,omitempty"`
	Stargazers_url string `json:"stargazers_url"`
	Role_name string `json:"role_name,omitempty"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Open_issues int `json:"open_issues,omitempty"`
	Keys_url string `json:"keys_url"`
	Releases_url string `json:"releases_url"`
	Clone_url string `json:"clone_url,omitempty"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Description string `json:"description"`
	Created_at string `json:"created_at,omitempty"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Size int `json:"size,omitempty"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Git_refs_url string `json:"git_refs_url"`
	Ssh_url string `json:"ssh_url,omitempty"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Code_of_conduct GeneratedType `json:"code_of_conduct,omitempty"` // Code Of Conduct
	Downloads_url string `json:"downloads_url"`
	Contributors_url string `json:"contributors_url"`
	Stargazers_count int `json:"stargazers_count,omitempty"`
	Network_count int `json:"network_count,omitempty"`
	Assignees_url string `json:"assignees_url"`
	Has_issues bool `json:"has_issues,omitempty"`
	Trees_url string `json:"trees_url"`
	Disabled bool `json:"disabled,omitempty"`
	Fork bool `json:"fork"`
	Has_wiki bool `json:"has_wiki,omitempty"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Issue_comment_url string `json:"issue_comment_url"`
	Subscribers_url string `json:"subscribers_url"`
	Compare_url string `json:"compare_url"`
	Git_url string `json:"git_url,omitempty"`
	Has_projects bool `json:"has_projects,omitempty"`
	Archived bool `json:"archived,omitempty"`
	Branches_url string `json:"branches_url"`
	Contents_url string `json:"contents_url"`
	Private bool `json:"private"`
	Url string `json:"url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Name string `json:"name"`
	Homepage string `json:"homepage,omitempty"`
	Statuses_url string `json:"statuses_url"`
	Security_and_analysis GeneratedType `json:"security_and_analysis,omitempty"`
	Languages_url string `json:"languages_url"`
	Default_branch string `json:"default_branch,omitempty"`
	Subscription_url string `json:"subscription_url"`
	Tags_url string `json:"tags_url"`
	Comments_url string `json:"comments_url"`
	Deployments_url string `json:"deployments_url"`
	Html_url string `json:"html_url"`
	Git_commits_url string `json:"git_commits_url"`
	Full_name string `json:"full_name"`
	Merges_url string `json:"merges_url"`
	Hooks_url string `json:"hooks_url"`
	Is_template bool `json:"is_template,omitempty"`
	Blobs_url string `json:"blobs_url"`
	Notifications_url string `json:"notifications_url"`
	Language string `json:"language,omitempty"`
	Watchers int `json:"watchers,omitempty"`
	License map[string]interface{} `json:"license,omitempty"`
	Visibility string `json:"visibility,omitempty"`
	Topics []string `json:"topics,omitempty"`
	Commits_url string `json:"commits_url"`
	Pulls_url string `json:"pulls_url"`
	Mirror_url string `json:"mirror_url,omitempty"`
	Collaborators_url string `json:"collaborators_url"`
	Milestones_url string `json:"milestones_url"`
	Git_tags_url string `json:"git_tags_url"`
	Archive_url string `json:"archive_url"`
	Svn_url string `json:"svn_url,omitempty"`
	Id int64 `json:"id"`
	Watchers_count int `json:"watchers_count,omitempty"`
	Open_issues_count int `json:"open_issues_count,omitempty"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Forks int `json:"forks,omitempty"`
	Teams_url string `json:"teams_url"`
	Updated_at string `json:"updated_at,omitempty"`
	Node_id string `json:"node_id"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
}

// Label represents the Label schema from the OpenAPI specification
type Label struct {
	Id int64 `json:"id"` // Unique identifier for the label.
	Name string `json:"name"` // The name of the label.
	Node_id string `json:"node_id"`
	Url string `json:"url"` // URL for the label
	Color string `json:"color"` // 6-character hex code, without the leading #, identifying the color
	DefaultField bool `json:"default"` // Whether this label comes by default in a new repository.
	Description string `json:"description"` // Optional description of the label, such as its purpose.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Code_security map[string]interface{} `json:"code_security,omitempty"`
	Dependabot_security_updates map[string]interface{} `json:"dependabot_security_updates,omitempty"` // Enable or disable Dependabot security updates for the repository.
	Secret_scanning map[string]interface{} `json:"secret_scanning,omitempty"`
	Secret_scanning_ai_detection map[string]interface{} `json:"secret_scanning_ai_detection,omitempty"`
	Secret_scanning_non_provider_patterns map[string]interface{} `json:"secret_scanning_non_provider_patterns,omitempty"`
	Secret_scanning_push_protection map[string]interface{} `json:"secret_scanning_push_protection,omitempty"`
	Advanced_security map[string]interface{} `json:"advanced_security,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Url string `json:"url"`
	Added_by string `json:"added_by,omitempty"`
	Key string `json:"key"`
	Read_only bool `json:"read_only"`
	Verified bool `json:"verified"`
	Created_at string `json:"created_at"`
	Enabled bool `json:"enabled,omitempty"`
	Last_used string `json:"last_used,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Event string `json:"event"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Assigner GeneratedType `json:"assigner"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Environments []map[string]interface{} `json:"environments"` // The list of environments that were approved or rejected
	State string `json:"state"` // Whether deployment to the environment(s) was approved or rejected or pending (with comments)
	User GeneratedType `json:"user"` // A GitHub user.
	Comment string `json:"comment"` // The comment submitted with the deployment review
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/checks/suites#get-a-check-suite).
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Old_answer Webhooksanswer `json:"old_answer"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// Commit represents the Commit schema from the OpenAPI specification
type Commit struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
	Author interface{} `json:"author"`
	Node_id string `json:"node_id"`
	Stats map[string]interface{} `json:"stats,omitempty"`
	Comments_url string `json:"comments_url"`
	Html_url string `json:"html_url"`
	Commit map[string]interface{} `json:"commit"`
	Committer interface{} `json:"committer"`
	Files []GeneratedType `json:"files,omitempty"`
	Parents []map[string]interface{} `json:"parents"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Label Webhookslabel `json:"label,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Action string `json:"action"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Assignee Webhooksuser `json:"assignee"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"` // The ID of the installation.
	Node_id string `json:"node_id"` // The global node ID of the installation.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Version_id int `json:"version_id"` // The ID of the previous version of the ruleset
	Actor map[string]interface{} `json:"actor"` // The actor who updated the ruleset
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Columns_url string `json:"columns_url"`
	Updated_at string `json:"updated_at"`
	Id int `json:"id"`
	Number int `json:"number"`
	Owner_url string `json:"owner_url"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Private bool `json:"private,omitempty"` // Whether the project is private or not. Only present when owner is an organization.
	Url string `json:"url"`
	Body string `json:"body"`
	Permissions map[string]interface{} `json:"permissions"`
	Html_url string `json:"html_url"`
	State string `json:"state"`
	Node_id string `json:"node_id"`
	Name string `json:"name"`
	Organization_permission string `json:"organization_permission,omitempty"` // The organization permission for this project. Only present when owner is an organization.
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Default_level string `json:"default_level,omitempty"` // The default repository access level for Dependabot updates.
	Accessible_repositories []GeneratedType `json:"accessible_repositories,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Permissions GeneratedType `json:"permissions"` // The permissions granted to the user access token.
	Repositories_url string `json:"repositories_url"`
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Single_file_name string `json:"single_file_name"`
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Account GeneratedType `json:"account"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ref_name map[string]interface{} `json:"ref_name,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	TypeField GeneratedType `json:"type"` // The type of issue.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Number int `json:"number"`
	Pull_request map[string]interface{} `json:"pull_request"`
	Reason string `json:"reason"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_column Webhooksprojectcolumn `json:"project_column"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int64 `json:"id"`
	Node_id string `json:"node_id"`
	Repos_url string `json:"repos_url"`
	Following_url string `json:"following_url"`
	Gravatar_id string `json:"gravatar_id"`
	Html_url string `json:"html_url"`
	Url string `json:"url"`
	Site_admin bool `json:"site_admin"`
	Organizations_url string `json:"organizations_url"`
	TypeField string `json:"type"`
	Received_events_url string `json:"received_events_url"`
	Events_url string `json:"events_url"`
	Email string `json:"email,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Login string `json:"login"`
	Starred_url string `json:"starred_url"`
	Subscriptions_url string `json:"subscriptions_url"`
	User_view_type string `json:"user_view_type,omitempty"`
	Followers_url string `json:"followers_url"`
	Name string `json:"name,omitempty"`
	Starred_at string `json:"starred_at,omitempty"`
	Gists_url string `json:"gists_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Public_ips map[string]interface{} `json:"public_ips"` // Provides details of static public IP limits for GitHub-hosted Hosted Runners
}

// Workflow represents the Workflow schema from the OpenAPI specification
type Workflow struct {
	Url string `json:"url"`
	Name string `json:"name"`
	State string `json:"state"`
	Deleted_at string `json:"deleted_at,omitempty"`
	Path string `json:"path"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Badge_url string `json:"badge_url"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pattern_update_scans []GeneratedType `json:"pattern_update_scans,omitempty"`
	Backfill_scans []GeneratedType `json:"backfill_scans,omitempty"`
	Custom_pattern_backfill_scans []interface{} `json:"custom_pattern_backfill_scans,omitempty"`
	Incremental_scans []GeneratedType `json:"incremental_scans,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_column Webhooksprojectcolumn `json:"project_column"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Deleted_at string `json:"deleted_at"`
	Description string `json:"description"`
	Public bool `json:"public"`
	Short_description string `json:"short_description"`
	Closed_at string `json:"closed_at"`
	Created_at string `json:"created_at"`
	Deleted_by GeneratedType `json:"deleted_by"` // A GitHub user.
	Id float64 `json:"id"`
	Updated_at string `json:"updated_at"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Title string `json:"title"`
	Node_id string `json:"node_id"`
	Number int `json:"number"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request_body_url string `json:"pull_request_body_url"` // The API URL to get the pull request where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	State_reason string `json:"state_reason,omitempty"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Body string `json:"body,omitempty"`
	Html_url string `json:"html_url"`
	Key string `json:"key"`
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// Job represents the Job schema from the OpenAPI specification
type Job struct {
	Created_at string `json:"created_at"` // The time that the job created, in ISO 8601 format.
	Runner_name string `json:"runner_name"` // The name of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Started_at string `json:"started_at"` // The time that the job started, in ISO 8601 format.
	Workflow_name string `json:"workflow_name"` // The name of the workflow.
	Run_url string `json:"run_url"`
	Runner_group_name string `json:"runner_group_name"` // The name of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Url string `json:"url"`
	Completed_at string `json:"completed_at"` // The time that the job finished, in ISO 8601 format.
	Name string `json:"name"` // The name of the job.
	Check_run_url string `json:"check_run_url"`
	Head_branch string `json:"head_branch"` // The name of the current branch.
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being run.
	Labels []string `json:"labels"` // Labels for the workflow job. Specified by the "runs_on" attribute in the action's workflow file.
	Run_id int `json:"run_id"` // The id of the associated workflow run.
	Runner_id int `json:"runner_id"` // The ID of the runner to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Steps []map[string]interface{} `json:"steps,omitempty"` // Steps in this job.
	Conclusion string `json:"conclusion"` // The outcome of the job.
	Run_attempt int `json:"run_attempt,omitempty"` // Attempt number of the associated workflow run, 1 for first attempt and higher if the workflow was re-run.
	Node_id string `json:"node_id"`
	Html_url string `json:"html_url"`
	Runner_group_id int `json:"runner_group_id"` // The ID of the runner group to which this job has been assigned. (If a runner hasn't yet been assigned, this will be null.)
	Status string `json:"status"` // The phase of the lifecycle that the job is currently in.
	Id int `json:"id"` // The id of the job.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Review map[string]interface{} `json:"review"` // The review that was affected.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// Webhooksprojectcard represents the Webhooksprojectcard schema from the OpenAPI specification
type Webhooksprojectcard struct {
	Project_url string `json:"project_url"`
	After_id int `json:"after_id,omitempty"`
	Column_id int `json:"column_id"`
	Note string `json:"note"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Archived bool `json:"archived"` // Whether or not the card is archived
	Column_url string `json:"column_url"`
	Content_url string `json:"content_url,omitempty"`
	Created_at string `json:"created_at"`
	Creator map[string]interface{} `json:"creator"`
	Id int `json:"id"` // The project card's ID
	Node_id string `json:"node_id"`
}

// Thread represents the Thread schema from the OpenAPI specification
type Thread struct {
	Last_read_at string `json:"last_read_at"`
	Reason string `json:"reason"`
	Subscription_url string `json:"subscription_url"`
	Unread bool `json:"unread"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Id string `json:"id"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Subject map[string]interface{} `json:"subject"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Deployment_status map[string]interface{} `json:"deployment_status"` // The [deployment status](https://docs.github.com/rest/deployments/statuses#list-deployment-statuses).
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Deployment map[string]interface{} `json:"deployment"` // The [deployment](https://docs.github.com/rest/deployments/deployments#list-deployments).
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Check_run map[string]interface{} `json:"check_run,omitempty"`
	Workflow Webhooksworkflow `json:"workflow,omitempty"`
	Workflow_run map[string]interface{} `json:"workflow_run,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Status string `json:"status"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Status string `json:"status,omitempty"` // The attachment status of the code security configuration on the repository.
	Repository GeneratedType `json:"repository,omitempty"` // A GitHub repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Location string `json:"location,omitempty"`
	Members_can_change_repo_visibility bool `json:"members_can_change_repo_visibility,omitempty"`
	Members_can_create_private_repositories bool `json:"members_can_create_private_repositories,omitempty"`
	Two_factor_requirement_enabled bool `json:"two_factor_requirement_enabled,omitempty"`
	Secret_scanning_enabled_for_new_repositories bool `json:"secret_scanning_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether secret scanning is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Description string `json:"description"`
	Archived_at string `json:"archived_at"`
	Members_can_delete_issues bool `json:"members_can_delete_issues,omitempty"`
	Members_can_invite_outside_collaborators bool `json:"members_can_invite_outside_collaborators,omitempty"`
	Followers int `json:"followers"`
	Created_at string `json:"created_at"`
	Members_can_create_public_pages bool `json:"members_can_create_public_pages,omitempty"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Email string `json:"email,omitempty"`
	Dependency_graph_enabled_for_new_repositories bool `json:"dependency_graph_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether dependency graph is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Secret_scanning_push_protection_custom_link_enabled bool `json:"secret_scanning_push_protection_custom_link_enabled,omitempty"` // Whether a custom link is shown to contributors who are blocked from pushing a secret by push protection.
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Login string `json:"login"`
	Updated_at string `json:"updated_at"`
	Members_can_create_public_repositories bool `json:"members_can_create_public_repositories,omitempty"`
	Default_repository_permission string `json:"default_repository_permission,omitempty"`
	Dependabot_security_updates_enabled_for_new_repositories bool `json:"dependabot_security_updates_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether Dependabot security updates are automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Html_url string `json:"html_url"`
	Public_members_url string `json:"public_members_url"`
	Display_commenter_full_name_setting_enabled bool `json:"display_commenter_full_name_setting_enabled,omitempty"`
	Members_can_create_private_pages bool `json:"members_can_create_private_pages,omitempty"`
	Members_can_create_internal_repositories bool `json:"members_can_create_internal_repositories,omitempty"`
	Id int `json:"id"`
	Members_can_create_repositories bool `json:"members_can_create_repositories,omitempty"`
	Billing_email string `json:"billing_email,omitempty"`
	Blog string `json:"blog,omitempty"`
	Company string `json:"company,omitempty"`
	Deploy_keys_enabled_for_repositories bool `json:"deploy_keys_enabled_for_repositories,omitempty"` // Controls whether or not deploy keys may be added and used for repositories in the organization.
	Disk_usage int `json:"disk_usage,omitempty"`
	Advanced_security_enabled_for_new_repositories bool `json:"advanced_security_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether GitHub Advanced Security is enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Twitter_username string `json:"twitter_username,omitempty"`
	Default_repository_branch string `json:"default_repository_branch,omitempty"` // The default branch for repositories created in this organization.
	Members_url string `json:"members_url"`
	Secret_scanning_push_protection_enabled_for_new_repositories bool `json:"secret_scanning_push_protection_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether secret scanning push protection is automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Name string `json:"name,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Secret_scanning_push_protection_custom_link string `json:"secret_scanning_push_protection_custom_link,omitempty"` // An optional URL string to display to contributors who are blocked from pushing a secret.
	Is_verified bool `json:"is_verified,omitempty"`
	Members_can_create_teams bool `json:"members_can_create_teams,omitempty"`
	Members_can_view_dependency_insights bool `json:"members_can_view_dependency_insights,omitempty"`
	Members_can_fork_private_repositories bool `json:"members_can_fork_private_repositories,omitempty"`
	Issues_url string `json:"issues_url"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Members_allowed_repository_creation_type string `json:"members_allowed_repository_creation_type,omitempty"`
	Has_repository_projects bool `json:"has_repository_projects"`
	Members_can_create_pages bool `json:"members_can_create_pages,omitempty"`
	Collaborators int `json:"collaborators,omitempty"` // The number of collaborators on private repositories. This field may be null if the number of private repositories is over 50,000.
	Events_url string `json:"events_url"`
	TypeField string `json:"type"`
	Public_gists int `json:"public_gists"`
	Members_can_delete_repositories bool `json:"members_can_delete_repositories,omitempty"`
	Has_organization_projects bool `json:"has_organization_projects"`
	Dependabot_alerts_enabled_for_new_repositories bool `json:"dependabot_alerts_enabled_for_new_repositories,omitempty"` // **Endpoint closing down notice.** Please use [code security configurations](https://docs.github.com/rest/code-security/configurations) instead. Whether Dependabot alerts are automatically enabled for new repositories and repositories transferred to this organization. This field is only visible to organization owners or members of a team with the security manager role.
	Following int `json:"following"`
	Public_repos int `json:"public_repos"`
	Hooks_url string `json:"hooks_url"`
	Private_gists int `json:"private_gists,omitempty"`
	Repos_url string `json:"repos_url"`
	Readers_can_create_discussions bool `json:"readers_can_create_discussions,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Runner_label string `json:"runner_label,omitempty"` // The label of the runner to use for code scanning default setup when runner_type is 'labeled'.
	Runner_type string `json:"runner_type,omitempty"` // Whether to use labeled runners or standard GitHub runners.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Name string `json:"name"` // Automatically generated name of this codespace.
	Billable_owner GeneratedType `json:"billable_owner"` // A GitHub user.
	State string `json:"state"` // State of this codespace.
	Stop_url string `json:"stop_url"` // API URL to stop this codespace.
	Id int64 `json:"id"`
	Runtime_constraints map[string]interface{} `json:"runtime_constraints,omitempty"`
	Idle_timeout_minutes int `json:"idle_timeout_minutes"` // The number of minutes of inactivity after which this codespace will be automatically stopped.
	Idle_timeout_notice string `json:"idle_timeout_notice,omitempty"` // Text to show user when codespace idle timeout minutes has been overriden by an organization policy
	Url string `json:"url"` // API URL for this codespace.
	Start_url string `json:"start_url"` // API URL to start this codespace.
	Display_name string `json:"display_name,omitempty"` // Display name for this codespace.
	Last_used_at string `json:"last_used_at"` // Last known time this codespace was started.
	Repository GeneratedType `json:"repository"` // Full Repository
	Retention_expires_at string `json:"retention_expires_at,omitempty"` // When a codespace will be auto-deleted based on the "retention_period_minutes" and "last_used_at"
	Machines_url string `json:"machines_url"` // API URL to access available alternate machine types for this codespace.
	Git_status map[string]interface{} `json:"git_status"` // Details about the codespace's git repository.
	Location string `json:"location"` // The initally assigned location of a new codespace.
	Devcontainer_path string `json:"devcontainer_path,omitempty"` // Path to devcontainer.json from repo root used to create Codespace.
	Pending_operation_disabled_reason string `json:"pending_operation_disabled_reason,omitempty"` // Text to show user when codespace is disabled by a pending operation
	Machine GeneratedType `json:"machine"` // A description of the machine powering a codespace.
	Prebuild bool `json:"prebuild"` // Whether the codespace was created from a prebuild.
	Environment_id string `json:"environment_id"` // UUID identifying this codespace's environment.
	Pending_operation bool `json:"pending_operation,omitempty"` // Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
	Retention_period_minutes int `json:"retention_period_minutes,omitempty"` // Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
	Web_url string `json:"web_url"` // URL to access this codespace on the web.
	Recent_folders []string `json:"recent_folders"`
	Created_at string `json:"created_at"`
	Pulls_url string `json:"pulls_url"` // API URL for the Pull Request associated with this codespace, if any.
	Updated_at string `json:"updated_at"`
	Publish_url string `json:"publish_url,omitempty"` // API URL to publish this codespace to a new repository.
}

// Hovercard represents the Hovercard schema from the OpenAPI specification
type Hovercard struct {
	Contexts []map[string]interface{} `json:"contexts"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Owner []int `json:"owner"`
	All []int `json:"all"`
}

// Webhooksmilestone represents the Webhooksmilestone schema from the OpenAPI specification
type Webhooksmilestone struct {
	Html_url string `json:"html_url"`
	Url string `json:"url"`
	Creator map[string]interface{} `json:"creator"`
	Updated_at string `json:"updated_at"`
	Id int `json:"id"`
	Labels_url string `json:"labels_url"`
	Node_id string `json:"node_id"`
	Closed_at string `json:"closed_at"`
	Description string `json:"description"`
	Open_issues int `json:"open_issues"`
	Created_at string `json:"created_at"`
	Number int `json:"number"` // The number of the milestone.
	State string `json:"state"` // The state of the milestone.
	Closed_issues int `json:"closed_issues"`
	Due_on string `json:"due_on"`
	Title string `json:"title"` // The title of the milestone.
}

// Link represents the Link schema from the OpenAPI specification
type Link struct {
	Href string `json:"href"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/checks/suites#get-a-check-suite).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Source_type string `json:"source_type,omitempty"` // The source type of the property
	Value_type string `json:"value_type"` // The type of the value for the property
	Url string `json:"url,omitempty"` // The URL that can be used to fetch, update, or delete info about this property via the API.
	Values_editable_by string `json:"values_editable_by,omitempty"` // Who can edit the values of the property
	Property_name string `json:"property_name"` // The name of the property
	Description string `json:"description,omitempty"` // Short description of the property
	Allowed_values []string `json:"allowed_values,omitempty"` // An ordered list of the allowed values of the property. The property can have up to 200 allowed values.
	Default_value string `json:"default_value,omitempty"` // Default value of the property
	Required bool `json:"required,omitempty"` // Whether the property is required.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description,omitempty"`
	Id string `json:"id"`
	Name string `json:"name"`
	Color string `json:"color,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Run_url string `json:"run_url,omitempty"` // URL of the corresponding run.
	Run_id int `json:"run_id,omitempty"` // ID of the corresponding run.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender map[string]interface{} `json:"sender"`
	Team Webhooksteam `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Scope string `json:"scope"` // The scope of the membership. Currently, can only be `team`.
	Action string `json:"action"`
	Member Webhooksuser `json:"member"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// Webhooksrelease1 represents the Webhooksrelease1 schema from the OpenAPI specification
type Webhooksrelease1 struct {
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Body string `json:"body"`
	Discussion_url string `json:"discussion_url,omitempty"`
	Node_id string `json:"node_id"`
	Assets_url string `json:"assets_url"`
	Prerelease bool `json:"prerelease"` // Whether the release is identified as a prerelease or a full release.
	Target_commitish string `json:"target_commitish"` // Specifies the commitish value that determines where the Git tag is created from.
	Published_at string `json:"published_at"`
	Url string `json:"url"`
	Tarball_url string `json:"tarball_url"`
	Created_at string `json:"created_at"`
	Draft bool `json:"draft"` // Whether the release is a draft or published
	Name string `json:"name"`
	Zipball_url string `json:"zipball_url"`
	Id int `json:"id"`
	Author map[string]interface{} `json:"author"`
	Tag_name string `json:"tag_name"` // The name of the tag.
	Assets []map[string]interface{} `json:"assets"`
	Upload_url string `json:"upload_url"`
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Webhooksdeploykey represents the Webhooksdeploykey schema from the OpenAPI specification
type Webhooksdeploykey struct {
	Added_by string `json:"added_by,omitempty"`
	Last_used string `json:"last_used,omitempty"`
	Verified bool `json:"verified"`
	Enabled bool `json:"enabled,omitempty"`
	Id int `json:"id"`
	Read_only bool `json:"read_only"`
	Created_at string `json:"created_at"`
	Title string `json:"title"`
	Url string `json:"url"`
	Key string `json:"key"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Id int64 `json:"id"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Updated_at string `json:"updated_at"`
	After string `json:"after"`
	Head_branch string `json:"head_branch"`
	Head_commit GeneratedType `json:"head_commit"` // A commit.
	Status string `json:"status"` // The phase of the lifecycle that the check suite is currently in. Statuses of waiting, requested, and pending are reserved for GitHub Actions check suites.
	Before string `json:"before"`
	Conclusion string `json:"conclusion"`
	App GeneratedType `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Check_runs_url string `json:"check_runs_url"`
	Rerequestable bool `json:"rerequestable,omitempty"`
	Latest_check_runs_count int `json:"latest_check_runs_count"`
	Runs_rerequestable bool `json:"runs_rerequestable,omitempty"`
	Url string `json:"url"`
	Pull_requests []GeneratedType `json:"pull_requests"`
	Node_id string `json:"node_id"`
	Head_sha string `json:"head_sha"` // The SHA of the head commit that is being checked.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Cvss map[string]interface{} `json:"cvss"` // Details for the advisory pertaining to the Common Vulnerability Scoring System.
	Cwes []map[string]interface{} `json:"cwes"` // Details for the advisory pertaining to Common Weakness Enumeration.
	Published_at string `json:"published_at"` // The time that the advisory was published in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Summary string `json:"summary"` // A short, plain text summary of the advisory.
	References []map[string]interface{} `json:"references"` // Links to additional advisory information.
	Severity string `json:"severity"` // The severity of the advisory.
	Updated_at string `json:"updated_at"` // The time that the advisory was last modified in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Description string `json:"description"` // A long-form Markdown-supported description of the advisory.
	Withdrawn_at string `json:"withdrawn_at"` // The time that the advisory was withdrawn in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Cvss_severities GeneratedType `json:"cvss_severities,omitempty"`
	Identifiers []map[string]interface{} `json:"identifiers"` // Values that identify this advisory among security information sources.
	Vulnerabilities []GeneratedType `json:"vulnerabilities"` // Vulnerable version range information for the advisory.
	Ghsa_id string `json:"ghsa_id"` // The unique GitHub Security Advisory ID assigned to the advisory.
	Cve_id string `json:"cve_id"` // The unique CVE ID assigned to the advisory.
	Epss GeneratedType `json:"epss,omitempty"` // The EPSS scores as calculated by the [Exploit Prediction Scoring System](https://www.first.org/epss).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Include_claim_keys []string `json:"include_claim_keys,omitempty"` // Array of unique strings. Each claim key can only contain alphanumeric characters and underscores.
	Use_default bool `json:"use_default"` // Whether to use the default template or not. If `true`, the `include_claim_keys` field is ignored.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Expire_at string `json:"expire_at,omitempty"` // The time that the bypass will expire in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Reason string `json:"reason,omitempty"` // The reason for bypassing push protection.
	Token_type string `json:"token_type,omitempty"` // The token type this bypass is for.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Editors []map[string]interface{} `json:"editors,omitempty"`
	Languages []map[string]interface{} `json:"languages,omitempty"` // Code completion metrics for active languages.
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // Number of users who accepted at least one Copilot code suggestion, across all active editors. Includes both full and partial acceptances.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Author GeneratedType `json:"author"` // A GitHub user.
	Total int `json:"total"`
	Weeks []map[string]interface{} `json:"weeks"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Version_id int `json:"version_id"` // The ID of the previous version of the ruleset
	Actor map[string]interface{} `json:"actor"` // The actor who updated the ruleset
	State map[string]interface{} `json:"state"` // The state of the ruleset version
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue2 `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Permission string `json:"permission"`
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
	Title string `json:"title,omitempty"`
	Url string `json:"url,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Id int `json:"id,omitempty"`
}

// Runner represents the Runner schema from the OpenAPI specification
type Runner struct {
	Id int `json:"id"` // The ID of the runner.
	Labels []GeneratedType `json:"labels"`
	Name string `json:"name"` // The name of the runner.
	Os string `json:"os"` // The Operating System of the runner.
	Runner_group_id int `json:"runner_group_id,omitempty"` // The ID of the runner group.
	Status string `json:"status"` // The status of the runner.
	Busy bool `json:"busy"`
	Ephemeral bool `json:"ephemeral,omitempty"`
}

// Webhookssecurityadvisory represents the Webhookssecurityadvisory schema from the OpenAPI specification
type Webhookssecurityadvisory struct {
	Published_at string `json:"published_at"`
	Severity string `json:"severity"`
	Summary string `json:"summary"`
	Description string `json:"description"`
	Updated_at string `json:"updated_at"`
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities"`
	Cvss map[string]interface{} `json:"cvss"`
	Cvss_severities GeneratedType `json:"cvss_severities,omitempty"`
	References []map[string]interface{} `json:"references"`
	Withdrawn_at string `json:"withdrawn_at"`
	Cwes []map[string]interface{} `json:"cwes"`
	Ghsa_id string `json:"ghsa_id"`
	Identifiers []map[string]interface{} `json:"identifiers"`
}

// Webhookslabel represents the Webhookslabel schema from the OpenAPI specification
type Webhookslabel struct {
	DefaultField bool `json:"default"`
	Description string `json:"description"`
	Id int `json:"id"`
	Name string `json:"name"` // The name of the label.
	Node_id string `json:"node_id"`
	Url string `json:"url"` // URL for the label
	Color string `json:"color"` // 6-character hex code, without the leading #, identifying the color
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enabled bool `json:"enabled"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Status string `json:"status,omitempty"`
	Url string `json:"url,omitempty"`
	Documentation_url string `json:"documentation_url,omitempty"`
	Message string `json:"message,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Site_admin bool `json:"site_admin"`
	User_view_type string `json:"user_view_type,omitempty"`
	Subscriptions_url string `json:"subscriptions_url"`
	Hireable bool `json:"hireable,omitempty"`
	Gists_url string `json:"gists_url"`
	Login string `json:"login"`
	Public_gists int `json:"public_gists,omitempty"`
	Followers_url string `json:"followers_url"`
	Avatar_url string `json:"avatar_url"`
	Events_url string `json:"events_url"`
	Starred_url string `json:"starred_url"`
	Bio string `json:"bio,omitempty"`
	Received_events_url string `json:"received_events_url"`
	Gravatar_id string `json:"gravatar_id"`
	Repos_url string `json:"repos_url"`
	Suspended_at string `json:"suspended_at,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Email string `json:"email,omitempty"`
	TypeField string `json:"type"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Score float64 `json:"score"`
	Url string `json:"url"`
	Followers int `json:"followers,omitempty"`
	Blog string `json:"blog,omitempty"`
	Following int `json:"following,omitempty"`
	Company string `json:"company,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Name string `json:"name,omitempty"`
	Public_repos int `json:"public_repos,omitempty"`
	Following_url string `json:"following_url"`
	Html_url string `json:"html_url"`
	Organizations_url string `json:"organizations_url"`
	Node_id string `json:"node_id"`
	Id int64 `json:"id"`
	Location string `json:"location,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_ruleset GeneratedType `json:"repository_ruleset"` // A set of rules to apply when specified conditions are met.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Visibility string `json:"visibility"` // Visibility of a secret
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_status_update GeneratedType `json:"projects_v2_status_update"` // An status update belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repositories_removed []map[string]interface{} `json:"repositories_removed"` // An array of repository objects, which were removed from the installation.
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Installation Installation `json:"installation"` // Installation
	Requester Webhooksuser `json:"requester"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories_added []map[string]interface{} `json:"repositories_added"` // An array of repository objects, which were added to the installation.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Environment_name string `json:"environment_name"` // The name of the environment to approve or reject.
	State string `json:"state"` // Whether to approve or reject deployment to the specified environments.
	Comment string `json:"comment,omitempty"` // Optional comment to include with the review.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_name map[string]interface{} `json:"repository_name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Definition GeneratedType `json:"definition"` // Custom property defined on an organization
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Html_url string `json:"html_url"`
}

// Webhooksteam represents the Webhooksteam schema from the OpenAPI specification
type Webhooksteam struct {
	Privacy string `json:"privacy,omitempty"`
	Slug string `json:"slug,omitempty"`
	Name string `json:"name"` // Name of the team
	Parent map[string]interface{} `json:"parent,omitempty"`
	Permission string `json:"permission,omitempty"` // Permission that the team will have for its repositories
	Repositories_url string `json:"repositories_url,omitempty"`
	Url string `json:"url,omitempty"` // URL for the team
	Html_url string `json:"html_url,omitempty"`
	Notification_setting string `json:"notification_setting,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Description string `json:"description,omitempty"` // Description of the team
	Id int `json:"id"` // Unique identifier of the team
	Members_url string `json:"members_url,omitempty"`
	Node_id string `json:"node_id,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert map[string]interface{} `json:"alert"` // The security alert of the vulnerable dependency.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Expiry string `json:"expiry,omitempty"` // The duration of the interaction restriction. Default: `one_day`.
	Limit string `json:"limit"` // The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enabled bool `json:"enabled"` // Whether Dependabot security updates are enabled for the repository.
	Paused bool `json:"paused"` // Whether Dependabot security updates are paused for the repository.
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

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Security_advisory Webhookssecurityadvisory `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue2 `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
	Label map[string]interface{} `json:"label"`
	Node_id string `json:"node_id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body string `json:"body"` // The generated body describing the contents of the release supporting markdown formatting
	Name string `json:"name"` // The generated name of the release
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Base_ref string `json:"base_ref"`
	Compare string `json:"compare"` // URL that shows the changes in this `ref` update, from the `before` commit to the `after` commit. For a newly created `ref` that is directly based on the default branch, this is the comparison between the head of the default branch and the `after` commit. Otherwise, this shows all commits until the `after` commit.
	Created bool `json:"created"` // Whether this push created the `ref`.
	Ref string `json:"ref"` // The full git ref that was pushed. Example: `refs/heads/main` or `refs/tags/v3.14.1`.
	Before string `json:"before"` // The SHA of the most recent commit on `ref` before the push.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Head_commit map[string]interface{} `json:"head_commit"`
	Forced bool `json:"forced"` // Whether this push was a force push of the `ref`.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Deleted bool `json:"deleted"` // Whether this push deleted the `ref`.
	Pusher map[string]interface{} `json:"pusher"` // Metaproperties for Git author/committer information.
	Commits []map[string]interface{} `json:"commits"` // An array of commit objects describing the pushed commits. (Pushed commits are all commits that are included in the `compare` between the `before` commit and the `after` commit.) The array includes a maximum of 2048 commits. If necessary, you can use the [Commits API](https://docs.github.com/rest/commits) to fetch additional commits.
	After string `json:"after"` // The SHA of the most recent commit on `ref` after the push.
	Repository map[string]interface{} `json:"repository"` // A git repository
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Href string `json:"href"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Thread_url string `json:"thread_url,omitempty"`
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Ignored bool `json:"ignored"`
	Reason string `json:"reason"`
	Repository_url string `json:"repository_url,omitempty"`
	Subscribed bool `json:"subscribed"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"` // The action that was performed.
	Assignee Webhooksuser `json:"assignee,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Actor represents the Actor schema from the OpenAPI specification
type Actor struct {
	Gravatar_id string `json:"gravatar_id"`
	Id int `json:"id"`
	Login string `json:"login"`
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Display_login string `json:"display_login,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enabled bool `json:"enabled"` // Whether GitHub Actions is enabled on the repository.
	Selected_actions_url string `json:"selected_actions_url,omitempty"` // The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
	Allowed_actions string `json:"allowed_actions,omitempty"` // The permissions policy that controls the actions and reusable workflows that are allowed to run.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Role_name string `json:"role_name,omitempty"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Mirror_url string `json:"mirror_url"`
	Issues_url string `json:"issues_url"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Git_tags_url string `json:"git_tags_url"`
	Git_commits_url string `json:"git_commits_url"`
	Statuses_url string `json:"statuses_url"`
	Issue_events_url string `json:"issue_events_url"`
	Network_count int `json:"network_count,omitempty"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Releases_url string `json:"releases_url"`
	Languages_url string `json:"languages_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Topics []string `json:"topics,omitempty"`
	Pushed_at string `json:"pushed_at"`
	Subscribers_url string `json:"subscribers_url"`
	Commits_url string `json:"commits_url"`
	Labels_url string `json:"labels_url"`
	Merges_url string `json:"merges_url"`
	Html_url string `json:"html_url"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Private bool `json:"private"` // Whether the repository is private or public.
	Homepage string `json:"homepage"`
	Downloads_url string `json:"downloads_url"`
	Ssh_url string `json:"ssh_url"`
	License GeneratedType `json:"license"` // License Simple
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Contributors_url string `json:"contributors_url"`
	Stargazers_count int `json:"stargazers_count"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Stargazers_url string `json:"stargazers_url"`
	Notifications_url string `json:"notifications_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Git_refs_url string `json:"git_refs_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Trees_url string `json:"trees_url"`
	Keys_url string `json:"keys_url"`
	Clone_url string `json:"clone_url"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Deployments_url string `json:"deployments_url"`
	Forks int `json:"forks"`
	Forks_url string `json:"forks_url"`
	Url string `json:"url"`
	Forks_count int `json:"forks_count"`
	Blobs_url string `json:"blobs_url"`
	Collaborators_url string `json:"collaborators_url"`
	Assignees_url string `json:"assignees_url"`
	Description string `json:"description"`
	Open_issues int `json:"open_issues"`
	Teams_url string `json:"teams_url"`
	Fork bool `json:"fork"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Created_at string `json:"created_at"`
	Svn_url string `json:"svn_url"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Has_pages bool `json:"has_pages"`
	Comments_url string `json:"comments_url"`
	Open_issues_count int `json:"open_issues_count"`
	Archive_url string `json:"archive_url"`
	Full_name string `json:"full_name"`
	Pulls_url string `json:"pulls_url"`
	Id int `json:"id"` // Unique identifier of the repository
	Branches_url string `json:"branches_url"`
	Events_url string `json:"events_url"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Watchers int `json:"watchers"`
	Tags_url string `json:"tags_url"`
	Language string `json:"language"`
	Subscription_url string `json:"subscription_url"`
	Size int `json:"size"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Updated_at string `json:"updated_at"`
	Git_url string `json:"git_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Watchers_count int `json:"watchers_count"`
	Compare_url string `json:"compare_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Milestones_url string `json:"milestones_url"`
	Node_id string `json:"node_id"`
	Contents_url string `json:"contents_url"`
	Name string `json:"name"` // The name of the repository.
	Issue_comment_url string `json:"issue_comment_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Hooks_url string `json:"hooks_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Member Webhooksuser `json:"member"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Previous_column_name string `json:"previous_column_name,omitempty"`
	Project_id int `json:"project_id"`
	Project_url string `json:"project_url"`
	Url string `json:"url"`
	Column_name string `json:"column_name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request_comment_url string `json:"pull_request_comment_url"` // The API URL to get the pull request comment where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id string `json:"id"` // The unique identifier of the network settings resource.
	Name string `json:"name"` // The name of the network settings resource.
	Network_configuration_id string `json:"network_configuration_id,omitempty"` // The identifier of the network configuration that is using this settings resource.
	Region string `json:"region"` // The location of the subnet this network settings resource is configured for.
	Subnet_id string `json:"subnet_id"` // The subnet this network settings resource is configured for.
}

// Hook represents the Hook schema from the OpenAPI specification
type Hook struct {
	Test_url string `json:"test_url"`
	Config GeneratedType `json:"config"` // Configuration object of the webhook
	Ping_url string `json:"ping_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Events []string `json:"events"` // Determines what events the hook is triggered for. Default: ['push'].
	Id int `json:"id"` // Unique identifier of the webhook.
	TypeField string `json:"type"`
	Created_at string `json:"created_at"`
	Deliveries_url string `json:"deliveries_url,omitempty"`
	Last_response GeneratedType `json:"last_response"`
	Name string `json:"name"` // The name of a valid service, use 'web' for a webhook.
	Active bool `json:"active"` // Determines whether the hook is actually triggered on pushes.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Starred_at string `json:"starred_at"`
	Repo Repository `json:"repo"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Role string `json:"role"` // The role of the user in the team.
	State string `json:"state"` // The state of the user's membership in the team.
}

// Webhooksworkflow represents the Webhooksworkflow schema from the OpenAPI specification
type Webhooksworkflow struct {
	Url string `json:"url"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Path string `json:"path"`
	State string `json:"state"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Badge_url string `json:"badge_url"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Member Webhooksuser `json:"member"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Scope string `json:"scope"` // The scope of the membership. Currently, can only be `team`.
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Team Webhooksteam `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Sender map[string]interface{} `json:"sender"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Topic represents the Topic schema from the OpenAPI specification
type Topic struct {
	Names []string `json:"names"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Personal_access_token_request GeneratedType `json:"personal_access_token_request"` // Details of a Personal Access Token Request.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// Migration represents the Migration schema from the OpenAPI specification
type Migration struct {
	Guid string `json:"guid"`
	Exclude_releases bool `json:"exclude_releases"`
	Archive_url string `json:"archive_url,omitempty"`
	Exclude_attachments bool `json:"exclude_attachments"`
	Lock_repositories bool `json:"lock_repositories"`
	Updated_at string `json:"updated_at"`
	Exclude_git_data bool `json:"exclude_git_data"`
	Exclude []string `json:"exclude,omitempty"` // Exclude related items from being returned in the response in order to improve performance of the request. The array can include any of: `"repositories"`.
	Exclude_metadata bool `json:"exclude_metadata"`
	Id int64 `json:"id"`
	Node_id string `json:"node_id"`
	Org_metadata_only bool `json:"org_metadata_only"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Created_at string `json:"created_at"`
	State string `json:"state"`
	Url string `json:"url"`
	Exclude_owner_projects bool `json:"exclude_owner_projects"`
	Repositories []Repository `json:"repositories"` // The repositories included in the migration. Only returned for export migrations.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id string `json:"id"` // The ID of the image. Use this ID for the `image` parameter when creating a new larger runner.
	Platform string `json:"platform"` // The operating system of the image.
	Size_gb int `json:"size_gb"` // Image size in GB.
	Source string `json:"source"` // The image provider.
	Display_name string `json:"display_name"` // Display name for this image.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Reason string `json:"reason,omitempty"` // Explains why the merge group is being destroyed. The group could have been merged, removed from the queue (dequeued), or invalidated by an earlier queue entry being dequeued (invalidated).
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Merge_group GeneratedType `json:"merge_group"` // A group of pull requests that the merge queue has grouped together to be merged.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Label Webhookslabel `json:"label,omitempty"`
	Number int `json:"number"` // The pull request number.
	Pull_request map[string]interface{} `json:"pull_request"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alert GeneratedType `json:"alert"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Location GeneratedType `json:"location"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Releases_url string `json:"releases_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Code_search_index_status map[string]interface{} `json:"code_search_index_status,omitempty"` // The status of the code search index for this repository
	Git_tags_url string `json:"git_tags_url"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Trees_url string `json:"trees_url"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Name string `json:"name"` // The name of the repository.
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Svn_url string `json:"svn_url"`
	Watchers_count int `json:"watchers_count"`
	Ssh_url string `json:"ssh_url"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Pushed_at string `json:"pushed_at"`
	Has_pages bool `json:"has_pages"`
	Milestones_url string `json:"milestones_url"`
	Notifications_url string `json:"notifications_url"`
	Watchers int `json:"watchers"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Mirror_url string `json:"mirror_url"`
	License GeneratedType `json:"license"` // License Simple
	Branches_url string `json:"branches_url"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Node_id string `json:"node_id"`
	Stargazers_count int `json:"stargazers_count"`
	Keys_url string `json:"keys_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Topics []string `json:"topics,omitempty"`
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Git_url string `json:"git_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Commits_url string `json:"commits_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Stargazers_url string `json:"stargazers_url"`
	Clone_url string `json:"clone_url"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Languages_url string `json:"languages_url"`
	Labels_url string `json:"labels_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Collaborators_url string `json:"collaborators_url"`
	Full_name string `json:"full_name"`
	Html_url string `json:"html_url"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Assignees_url string `json:"assignees_url"`
	Forks_count int `json:"forks_count"`
	Issues_url string `json:"issues_url"`
	Tags_url string `json:"tags_url"`
	Compare_url string `json:"compare_url"`
	Git_commits_url string `json:"git_commits_url"`
	Blobs_url string `json:"blobs_url"`
	Open_issues int `json:"open_issues"`
	Fork bool `json:"fork"`
	Id int64 `json:"id"` // Unique identifier of the repository
	Language string `json:"language"`
	Teams_url string `json:"teams_url"`
	Downloads_url string `json:"downloads_url"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Subscription_url string `json:"subscription_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Statuses_url string `json:"statuses_url"`
	Merges_url string `json:"merges_url"`
	Homepage string `json:"homepage"`
	Url string `json:"url"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Private bool `json:"private"` // Whether the repository is private or public.
	Contributors_url string `json:"contributors_url"`
	Issue_events_url string `json:"issue_events_url"`
	Comments_url string `json:"comments_url"`
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Open_issues_count int `json:"open_issues_count"`
	Hooks_url string `json:"hooks_url"`
	Pulls_url string `json:"pulls_url"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Forks_url string `json:"forks_url"`
	Events_url string `json:"events_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Archive_url string `json:"archive_url"`
	Description string `json:"description"`
	Contents_url string `json:"contents_url"`
	Forks int `json:"forks"`
	Git_refs_url string `json:"git_refs_url"`
	Subscribers_url string `json:"subscribers_url"`
	Deployments_url string `json:"deployments_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Custom_pattern_name string `json:"custom_pattern_name,omitempty"` // If the scan was triggered by a custom pattern update, this will be the name of the pattern that was updated
	Custom_pattern_scope string `json:"custom_pattern_scope,omitempty"` // If the scan was triggered by a custom pattern update, this will be the scope of the pattern that was updated
	Secret_types []string `json:"secret_types,omitempty"` // List of patterns that were updated. This will be empty for normal backfill scans or custom pattern updates
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Source string `json:"source"` // What type of content was scanned
	Started_at string `json:"started_at"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	TypeField string `json:"type"` // What type of scan was completed
	Action string `json:"action"`
	Completed_at string `json:"completed_at"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name of the classroom.
	Url string `json:"url"` // The url of the classroom on GitHub Classroom.
	Archived bool `json:"archived"` // Returns whether classroom is archived or not.
	Id int `json:"id"` // Unique identifier of the classroom.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Completed_at string `json:"completed_at,omitempty"` // The time that the scan was completed. Empty if the scan is running
	Started_at string `json:"started_at,omitempty"` // The time that the scan was started. Empty if the scan is pending
	Status string `json:"status,omitempty"` // The state of the scan. Either "completed", "running", or "pending"
	TypeField string `json:"type,omitempty"` // The type of scan
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone3 `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id string `json:"id"` // The unique identifier of the network configuration.
	Name string `json:"name"` // The name of the network configuration.
	Network_settings_ids []string `json:"network_settings_ids,omitempty"` // The unique identifier of each network settings in the configuration.
	Compute_service string `json:"compute_service,omitempty"` // The hosted compute service the network configuration supports.
	Created_on string `json:"created_on"` // The time at which the network configuration was created, in ISO 8601 format.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Severity string `json:"severity"` // The severity of the vulnerability.
	Vulnerable_version_range string `json:"vulnerable_version_range"` // Conditions that identify vulnerable versions of this vulnerability's package.
	First_patched_version map[string]interface{} `json:"first_patched_version"` // Details pertaining to the package version that patches this vulnerability.
	PackageField GeneratedType `json:"package"` // Details for the vulnerable package.
}

// Webhooksanswer represents the Webhooksanswer schema from the OpenAPI specification
type Webhooksanswer struct {
	Parent_id interface{} `json:"parent_id"`
	Child_comment_count int `json:"child_comment_count"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Repository_url string `json:"repository_url"`
	Body string `json:"body"`
	Created_at string `json:"created_at"`
	Discussion_id int `json:"discussion_id"`
	Id int `json:"id"`
	User map[string]interface{} `json:"user"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Context string `json:"context"`
	Commit map[string]interface{} `json:"commit"`
	Description string `json:"description"` // The optional human-readable description added to the status.
	State string `json:"state"` // The new state. Can be `pending`, `success`, `failure`, or `error`.
	Sha string `json:"sha"` // The Commit SHA.
	Id int `json:"id"` // The unique identifier of the status.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Branches []map[string]interface{} `json:"branches"` // An array of branch objects containing the status' SHA. Each branch contains the given SHA, but the SHA may or may not be the head of the branch. The array includes a maximum of 10 branches.
	Name string `json:"name"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Target_url string `json:"target_url"` // The optional link added to the status.
	Avatar_url string `json:"avatar_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType `json:"pull_request"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Comment string `json:"comment,omitempty"`
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Since string `json:"since"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Workflow_job_run Webhooksworkflowjobrun `json:"workflow_job_run,omitempty"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Approver Webhooksapprover `json:"approver,omitempty"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Reviewers []map[string]interface{} `json:"reviewers,omitempty"`
	Workflow_job_runs []map[string]interface{} `json:"workflow_job_runs,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Reason string `json:"reason,omitempty"` // Reason for restriction
	Oid string `json:"oid"` // Full or abbreviated commit hash to reject
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Can_approve_pull_request_reviews bool `json:"can_approve_pull_request_reviews,omitempty"` // Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
	Default_workflow_permissions string `json:"default_workflow_permissions,omitempty"` // The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Personal_access_token_request GeneratedType `json:"personal_access_token_request"` // Details of a Personal Access Token Request.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Color string `json:"color"`
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Membership Webhooksmembership `json:"membership"` // The membership between the user and the organization. Not present when the action is `member_invited`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Registry_package map[string]interface{} `json:"registry_package"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Annotation_level string `json:"annotation_level"`
	Message string `json:"message"`
	Raw_details string `json:"raw_details"`
	Start_column int `json:"start_column"`
	Title string `json:"title"`
	Blob_href string `json:"blob_href"`
	End_line int `json:"end_line"`
	End_column int `json:"end_column"`
	Path string `json:"path"`
	Start_line int `json:"start_line"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// Autolink represents the Autolink schema from the OpenAPI specification
type Autolink struct {
	Id int `json:"id"`
	Is_alphanumeric bool `json:"is_alphanumeric"` // Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters.
	Key_prefix string `json:"key_prefix"` // The prefix of a key that is linkified.
	Url_template string `json:"url_template"` // A template for the target URL that is generated if a key was found.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Key string `json:"key"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"`
	Name string `json:"name"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
}

// Snapshot represents the Snapshot schema from the OpenAPI specification
type Snapshot struct {
	Sha string `json:"sha"` // The commit SHA associated with this dependency snapshot. Maximum length: 40 characters.
	Version int `json:"version"` // The version of the repository snapshot submission.
	Detector map[string]interface{} `json:"detector"` // A description of the detector used.
	Job map[string]interface{} `json:"job"`
	Manifests map[string]interface{} `json:"manifests,omitempty"` // A collection of package manifests, which are a collection of related dependencies declared in a file or representing a logical group of dependencies.
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Ref string `json:"ref"` // The repository branch that triggered this snapshot.
	Scanned string `json:"scanned"` // The time at which the snapshot was scanned.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Object map[string]interface{} `json:"object"`
	Ref string `json:"ref"`
	Url string `json:"url"`
}

// Feed represents the Feed schema from the OpenAPI specification
type Feed struct {
	Current_user_organization_url string `json:"current_user_organization_url,omitempty"`
	Security_advisories_url string `json:"security_advisories_url,omitempty"`
	Current_user_actor_url string `json:"current_user_actor_url,omitempty"`
	Current_user_organization_urls []string `json:"current_user_organization_urls,omitempty"`
	Current_user_url string `json:"current_user_url,omitempty"`
	Repository_discussions_category_url string `json:"repository_discussions_category_url,omitempty"` // A feed of discussions for a given repository and category.
	User_url string `json:"user_url"`
	Current_user_public_url string `json:"current_user_public_url,omitempty"`
	Repository_discussions_url string `json:"repository_discussions_url,omitempty"` // A feed of discussions for a given repository.
	Links map[string]interface{} `json:"_links"`
	Timeline_url string `json:"timeline_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Count int `json:"count"`
	Path string `json:"path"`
	Title string `json:"title"`
	Uniques int `json:"uniques"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Role string `json:"role"` // The user's membership type in the organization.
	State string `json:"state"` // The state of the member in the organization. The `pending` state indicates the user has not yet accepted an invitation.
	Url string `json:"url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Organization_url string `json:"organization_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Events_url string `json:"events_url"`
	Id int `json:"id"`
	Name string `json:"name,omitempty"`
	Members_allowed_repository_creation_type string `json:"members_allowed_repository_creation_type,omitempty"`
	Email string `json:"email,omitempty"`
	Members_url string `json:"members_url"`
	Description string `json:"description"`
	Members_can_fork_private_repositories bool `json:"members_can_fork_private_repositories,omitempty"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Html_url string `json:"html_url"`
	Updated_at string `json:"updated_at"`
	Archived_at string `json:"archived_at"`
	Public_members_url string `json:"public_members_url"`
	Blog string `json:"blog,omitempty"`
	Members_can_create_private_pages bool `json:"members_can_create_private_pages,omitempty"`
	Public_repos int `json:"public_repos"`
	Members_can_create_internal_repositories bool `json:"members_can_create_internal_repositories,omitempty"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Has_organization_projects bool `json:"has_organization_projects"`
	Following int `json:"following"`
	Default_repository_permission string `json:"default_repository_permission,omitempty"`
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Followers int `json:"followers"`
	Public_gists int `json:"public_gists"`
	Members_can_create_pages bool `json:"members_can_create_pages,omitempty"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Collaborators int `json:"collaborators,omitempty"`
	Has_repository_projects bool `json:"has_repository_projects"`
	Members_can_create_private_repositories bool `json:"members_can_create_private_repositories,omitempty"`
	Repos_url string `json:"repos_url"`
	TypeField string `json:"type"`
	Company string `json:"company,omitempty"`
	Location string `json:"location,omitempty"`
	Private_gists int `json:"private_gists,omitempty"`
	Node_id string `json:"node_id"`
	Issues_url string `json:"issues_url"`
	Is_verified bool `json:"is_verified,omitempty"`
	Two_factor_requirement_enabled bool `json:"two_factor_requirement_enabled,omitempty"`
	Hooks_url string `json:"hooks_url"`
	Members_can_create_public_repositories bool `json:"members_can_create_public_repositories,omitempty"`
	Url string `json:"url"`
	Login string `json:"login"`
	Avatar_url string `json:"avatar_url"`
	Billing_email string `json:"billing_email,omitempty"`
	Members_can_create_repositories bool `json:"members_can_create_repositories,omitempty"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Members_can_create_public_pages bool `json:"members_can_create_public_pages,omitempty"`
	Created_at string `json:"created_at"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Display_name string `json:"display_name"` // Display name for this image.
	Id string `json:"id"` // The ID of the image. Use this ID for the `image` parameter when creating a new larger runner.
	Size_gb int `json:"size_gb"` // Image size in GB.
	Source string `json:"source"` // The image provider.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Teams []Team `json:"teams"`
	Users []GeneratedType `json:"users"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Event string `json:"event"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Head_commit GeneratedType `json:"head_commit"` // A commit.
	Head_ref string `json:"head_ref"` // The full ref of the merge group.
	Head_sha string `json:"head_sha"` // The SHA of the merge group.
	Base_ref string `json:"base_ref"` // The full ref of the branch the merge group will be merged into.
	Base_sha string `json:"base_sha"` // The SHA of the merge group's parent commit.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Created_at string `json:"created_at"`
	Id int `json:"id"` // Unique identifier of the issue comment
	Updated_at string `json:"updated_at"`
	Html_url string `json:"html_url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Body_html string `json:"body_html,omitempty"`
	Event string `json:"event"`
	Issue_url string `json:"issue_url"`
	Body_text string `json:"body_text,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
	Url string `json:"url"` // URL for the issue comment
	Body string `json:"body,omitempty"` // Contents of the issue comment
	Author_association string `json:"author_association"` // How the author is associated with the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow string `json:"workflow"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Inputs map[string]interface{} `json:"inputs"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Ref string `json:"ref"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Starred_url string `json:"starred_url"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Public_gists int `json:"public_gists"`
	Blog string `json:"blog"`
	Node_id string `json:"node_id"`
	Name string `json:"name"`
	Updated_at string `json:"updated_at"`
	Bio string `json:"bio"`
	Total_private_repos int `json:"total_private_repos"`
	Avatar_url string `json:"avatar_url"`
	Private_gists int `json:"private_gists"`
	Location string `json:"location"`
	Site_admin bool `json:"site_admin"`
	Ldap_dn string `json:"ldap_dn,omitempty"`
	Followers_url string `json:"followers_url"`
	Business_plus bool `json:"business_plus,omitempty"`
	Created_at string `json:"created_at"`
	Company string `json:"company"`
	Url string `json:"url"`
	Two_factor_authentication bool `json:"two_factor_authentication"`
	Following int `json:"following"`
	Hireable bool `json:"hireable"`
	Html_url string `json:"html_url"`
	Followers int `json:"followers"`
	Disk_usage int `json:"disk_usage"`
	Events_url string `json:"events_url"`
	TypeField string `json:"type"`
	Public_repos int `json:"public_repos"`
	Subscriptions_url string `json:"subscriptions_url"`
	Gists_url string `json:"gists_url"`
	Received_events_url string `json:"received_events_url"`
	User_view_type string `json:"user_view_type,omitempty"`
	Following_url string `json:"following_url"`
	Gravatar_id string `json:"gravatar_id"`
	Id int64 `json:"id"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Notification_email string `json:"notification_email,omitempty"`
	Login string `json:"login"`
	Repos_url string `json:"repos_url"`
	Collaborators int `json:"collaborators"`
	Email string `json:"email"`
	Organizations_url string `json:"organizations_url"`
	Owned_private_repos int `json:"owned_private_repos"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Event string `json:"event,omitempty"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Sha string `json:"sha"` // SHA for the commit
	Committer map[string]interface{} `json:"committer"` // Identifying information for the git-user
	Url string `json:"url"`
	Verification map[string]interface{} `json:"verification"`
	Message string `json:"message"` // Message describing the purpose of the commit
	Tree map[string]interface{} `json:"tree"`
	Author map[string]interface{} `json:"author"` // Identifying information for the git-user
	Parents []map[string]interface{} `json:"parents"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	File_patterns []string `json:"file_patterns"` // Array of file patterns. Pull requests which change matching files must be approved by the specified team. File patterns use the same syntax as `.gitignore` files.
	Minimum_approvals int `json:"minimum_approvals"` // Minimum number of approvals required from the specified team. If set to zero, the team will be added to the pull request but approval is optional.
	Reviewer GeneratedType `json:"reviewer"` // A required reviewing team
}

// Enterprise represents the Enterprise schema from the OpenAPI specification
type Enterprise struct {
	Description string `json:"description,omitempty"` // A short description of the enterprise.
	Id int `json:"id"` // Unique identifier of the enterprise
	Node_id string `json:"node_id"`
	Slug string `json:"slug"` // The slug url identifier for the enterprise.
	Updated_at string `json:"updated_at"`
	Html_url string `json:"html_url"`
	Name string `json:"name"` // The name of the enterprise.
	Website_url string `json:"website_url,omitempty"` // The enterprise's website URL.
	Avatar_url string `json:"avatar_url"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.created JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Accepted bool `json:"accepted"` // Whether the user has accepted the permissions defined by the devcontainer config
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Identifiers []map[string]interface{} `json:"identifiers"`
	Cvss map[string]interface{} `json:"cvss"`
	References []string `json:"references"`
	Repository_advisory_url string `json:"repository_advisory_url"` // The API URL for the repository advisory.
	Html_url string `json:"html_url"` // The URL for the advisory.
	Source_code_location string `json:"source_code_location"` // The URL of the advisory's source code.
	Ghsa_id string `json:"ghsa_id"` // The GitHub Security Advisory ID.
	Cwes []map[string]interface{} `json:"cwes"`
	Withdrawn_at string `json:"withdrawn_at"` // The date and time of when the advisory was withdrawn, in ISO 8601 format.
	Nvd_published_at string `json:"nvd_published_at"` // The date and time when the advisory was published in the National Vulnerability Database, in ISO 8601 format. This field is only populated when the advisory is imported from the National Vulnerability Database.
	Updated_at string `json:"updated_at"` // The date and time of when the advisory was last updated, in ISO 8601 format.
	Credits []map[string]interface{} `json:"credits"` // The users who contributed to the advisory.
	Github_reviewed_at string `json:"github_reviewed_at"` // The date and time of when the advisory was reviewed by GitHub, in ISO 8601 format.
	Published_at string `json:"published_at"` // The date and time of when the advisory was published, in ISO 8601 format.
	Vulnerabilities []Vulnerability `json:"vulnerabilities"` // The products and respective version ranges affected by the advisory.
	Cvss_severities GeneratedType `json:"cvss_severities,omitempty"`
	Epss GeneratedType `json:"epss,omitempty"` // The EPSS scores as calculated by the [Exploit Prediction Scoring System](https://www.first.org/epss).
	TypeField string `json:"type"` // The type of advisory.
	Severity string `json:"severity"` // The severity of the advisory.
	Url string `json:"url"` // The API URL for the advisory.
	Cve_id string `json:"cve_id"` // The Common Vulnerabilities and Exposures (CVE) ID.
	Summary string `json:"summary"` // A short summary of the advisory.
	Description string `json:"description"` // A detailed description of what the advisory entails.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Merge_group GeneratedType `json:"merge_group"` // A group of pull requests that the merge queue has grouped together to be merged.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	End_column int `json:"end_column,omitempty"`
	End_line int `json:"end_line,omitempty"`
	Path string `json:"path,omitempty"`
	Start_column int `json:"start_column,omitempty"`
	Start_line int `json:"start_line,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Delivered_at string `json:"delivered_at"` // Time when the delivery was delivered.
	Action string `json:"action"` // The type of activity for the event that triggered the delivery.
	Id int `json:"id"` // Unique identifier of the delivery.
	Url string `json:"url,omitempty"` // The URL target of the delivery.
	Installation_id int `json:"installation_id"` // The id of the GitHub App installation associated with this event.
	Request map[string]interface{} `json:"request"`
	Response map[string]interface{} `json:"response"`
	Status string `json:"status"` // Description of the status of the attempted delivery
	Guid string `json:"guid"` // Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
	Duration float64 `json:"duration"` // Time spent delivering.
	Redelivery bool `json:"redelivery"` // Whether the delivery is a redelivery.
	Repository_id int `json:"repository_id"` // The id of the repository associated with this event.
	Throttled_at string `json:"throttled_at,omitempty"` // Time when the webhook delivery was throttled.
	Status_code int `json:"status_code"` // Status code received when delivery was made.
	Event string `json:"event"` // The event that triggered the delivery.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the milestone if the action was `edited`.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at,omitempty"`
	Id int `json:"id"`
	Key string `json:"key"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Cwe_ids []string `json:"cwe_ids,omitempty"` // A list of Common Weakness Enumeration (CWE) IDs.
	Severity string `json:"severity,omitempty"` // The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities,omitempty"` // A product affected by the vulnerability detailed in a repository security advisory.
	Collaborating_teams []string `json:"collaborating_teams,omitempty"` // A list of team slugs which have been granted write access to the advisory.
	State string `json:"state,omitempty"` // The state of the advisory.
	Cve_id string `json:"cve_id,omitempty"` // The Common Vulnerabilities and Exposures (CVE) ID.
	Cvss_vector_string string `json:"cvss_vector_string,omitempty"` // The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
	Summary string `json:"summary,omitempty"` // A short summary of the advisory.
	Collaborating_users []string `json:"collaborating_users,omitempty"` // A list of usernames who have been granted write access to the advisory.
	Credits []map[string]interface{} `json:"credits,omitempty"` // A list of users receiving credit for their participation in the security advisory.
	Description string `json:"description,omitempty"` // A detailed description of what the advisory impacts.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Usageitems []map[string]interface{} `json:"usageItems,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Action string `json:"action"`
	Effective_date string `json:"effective_date,omitempty"` // The `pending_cancellation` and `pending_tier_change` event types will include the date the cancellation or tier change will take effect.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Pull_request_url string `json:"pull_request_url"`
	Submitted_at string `json:"submitted_at,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	State string `json:"state"`
	Body string `json:"body"` // The text of the review.
	Html_url string `json:"html_url"`
	Id int64 `json:"id"` // Unique identifier of the review
	User GeneratedType `json:"user"` // A GitHub user.
	Links map[string]interface{} `json:"_links"`
	Body_html string `json:"body_html,omitempty"`
	Commit_id string `json:"commit_id"` // A commit SHA for the review. If the commit object was garbage collected or forcibly deleted, then it no longer exists in Git and this value will be `null`.
	Node_id string `json:"node_id"`
}

// Webhooksissue represents the Webhooksissue schema from the OpenAPI specification
type Webhooksissue struct {
	Repository_url string `json:"repository_url"`
	Title string `json:"title"` // Title of the issue
	Created_at string `json:"created_at"`
	Url string `json:"url"` // URL for the issue
	Active_lock_reason string `json:"active_lock_reason"`
	Draft bool `json:"draft,omitempty"`
	Labels_url string `json:"labels_url"`
	TypeField GeneratedType `json:"type,omitempty"` // The type of issue.
	Sub_issues_summary map[string]interface{} `json:"sub_issues_summary,omitempty"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Node_id string `json:"node_id"`
	Number int `json:"number"`
	Reactions map[string]interface{} `json:"reactions"`
	Events_url string `json:"events_url"`
	Closed_at string `json:"closed_at"`
	Updated_at string `json:"updated_at"`
	Body string `json:"body"` // Contents of the issue
	Assignee map[string]interface{} `json:"assignee,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	State_reason string `json:"state_reason,omitempty"`
	Comments_url string `json:"comments_url"`
	Assignees []map[string]interface{} `json:"assignees"`
	Locked bool `json:"locked,omitempty"`
	Id int64 `json:"id"`
	Performed_via_github_app map[string]interface{} `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Html_url string `json:"html_url"`
	User map[string]interface{} `json:"user"`
	State string `json:"state,omitempty"` // State of the issue; either 'open' or 'closed'
	Labels []map[string]interface{} `json:"labels,omitempty"`
	Comments int `json:"comments"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Access_level string `json:"access_level"` // Defines the level of access that workflows outside of the repository have to actions and reusable workflows within the repository. `none` means the access is only possible from workflows in this repository. `user` level access allows sharing across user owned private repositories only. `organization` level access allows sharing across the organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request GeneratedType `json:"pull_request"`
	Number int `json:"number"` // The pull request number.
	Changes map[string]interface{} `json:"changes"` // The changes to the comment if the action was `edited`.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Help_uri string `json:"help_uri,omitempty"` // A link to the documentation for the rule used to detect the alert.
	Name string `json:"name,omitempty"` // The name of the rule used to detect the alert.
	Severity string `json:"severity,omitempty"` // The severity of the alert.
	Full_description string `json:"full_description,omitempty"` // A description of the rule used to detect the alert.
	Id string `json:"id,omitempty"` // A unique identifier for the rule used to detect the alert.
	Security_severity_level string `json:"security_severity_level,omitempty"` // The security severity of the alert.
	Tags []string `json:"tags,omitempty"` // A set of tags applicable for the rule.
	Description string `json:"description,omitempty"` // A short description of the rule used to detect the alert.
	Help string `json:"help,omitempty"` // Detailed documentation for the rule as GitHub Flavored Markdown.
}

// Installation represents the Installation schema from the OpenAPI specification
type Installation struct {
	Repositories_url string `json:"repositories_url"`
	Created_at string `json:"created_at"`
	Suspended_by GeneratedType `json:"suspended_by"` // A GitHub user.
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Suspended_at string `json:"suspended_at"`
	Access_tokens_url string `json:"access_tokens_url"`
	Html_url string `json:"html_url"`
	Single_file_name string `json:"single_file_name"`
	App_slug string `json:"app_slug"`
	Updated_at string `json:"updated_at"`
	Events []string `json:"events"`
	Id int `json:"id"` // The ID of the installation.
	Account interface{} `json:"account"`
	Permissions GeneratedType `json:"permissions"` // The permissions granted to the user access token.
	App_id int `json:"app_id"`
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Target_type string `json:"target_type"`
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Target_id int `json:"target_id"` // The ID of the user or organization this token is being scoped to.
	Contact_email string `json:"contact_email,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue_body_url string `json:"issue_body_url"` // The API URL to get the issue where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Allow_force_pushes map[string]interface{} `json:"allow_force_pushes,omitempty"`
	Enforce_admins GeneratedType `json:"enforce_admins,omitempty"` // Protected Branch Admin Enforced
	Url string `json:"url,omitempty"`
	Block_creations map[string]interface{} `json:"block_creations,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Protection_url string `json:"protection_url,omitempty"`
	Restrictions GeneratedType `json:"restrictions,omitempty"` // Branch Restriction Policy
	Required_linear_history map[string]interface{} `json:"required_linear_history,omitempty"`
	Required_pull_request_reviews GeneratedType `json:"required_pull_request_reviews,omitempty"` // Protected Branch Pull Request Review
	Required_signatures map[string]interface{} `json:"required_signatures,omitempty"`
	Allow_deletions map[string]interface{} `json:"allow_deletions,omitempty"`
	Required_status_checks GeneratedType `json:"required_status_checks,omitempty"` // Protected Branch Required Status Check
	Name string `json:"name,omitempty"`
	Required_conversation_resolution map[string]interface{} `json:"required_conversation_resolution,omitempty"`
	Allow_fork_syncing map[string]interface{} `json:"allow_fork_syncing,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
	Lock_branch map[string]interface{} `json:"lock_branch,omitempty"` // Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id,omitempty"` // The unique identifier of the branch or tag policy.
	Name string `json:"name,omitempty"` // The name pattern that branches or tags must match in order to deploy to the environment.
	Node_id string `json:"node_id,omitempty"`
	TypeField string `json:"type,omitempty"` // Whether this rule targets a branch or tag.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Reason string `json:"reason,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Pull_request map[string]interface{} `json:"pull_request"`
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"`
}

// Activity represents the Activity schema from the OpenAPI specification
type Activity struct {
	Activity_type string `json:"activity_type"` // The type of the activity that was performed.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	After string `json:"after"` // The SHA of the commit after the activity.
	Before string `json:"before"` // The SHA of the commit before the activity.
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Ref string `json:"ref"` // The full Git reference, formatted as `refs/heads/<branch name>`.
	Timestamp string `json:"timestamp"` // The time when the activity occurred.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request GeneratedType `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request_review_comment_url string `json:"pull_request_review_comment_url"` // The API URL to get the pull request review comment where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"` // The type of credit the user is receiving.
	User GeneratedType `json:"user"` // A GitHub user.
	State string `json:"state"` // The state of the user's acceptance of the credit.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parent GeneratedType `json:"parent"` // Groups of organization members that gives permissions on specified repositories.
	Privacy string `json:"privacy,omitempty"`
	Permission string `json:"permission"`
	Id int `json:"id"`
	Url string `json:"url"`
	Assignment string `json:"assignment,omitempty"` // Determines if the team has a direct, indirect, or mixed relationship to a role
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Slug string `json:"slug"`
	Html_url string `json:"html_url"`
	Members_url string `json:"members_url"`
	Repositories_url string `json:"repositories_url"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Notification_setting string `json:"notification_setting,omitempty"`
	Description string `json:"description"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_advisory GeneratedType `json:"repository_advisory"` // A repository security advisory.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card Webhooksprojectcard `json:"project_card"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	After string `json:"after"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Pull_request map[string]interface{} `json:"pull_request"`
	Before string `json:"before"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Include_claim_keys []string `json:"include_claim_keys"` // Array of unique strings. Each claim key can only contain alphanumeric characters and underscores.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enabled_by GeneratedType `json:"enabled_by"` // A GitHub user.
	Merge_method string `json:"merge_method"` // The merge method to use.
	Commit_message string `json:"commit_message"` // Commit message for the merge commit.
	Commit_title string `json:"commit_title"` // Title for the merge commit message.
}

// Webhookspreviousmarketplacepurchase represents the Webhookspreviousmarketplacepurchase schema from the OpenAPI specification
type Webhookspreviousmarketplacepurchase struct {
	Account map[string]interface{} `json:"account"`
	Billing_cycle string `json:"billing_cycle"`
	Free_trial_ends_on interface{} `json:"free_trial_ends_on"`
	Next_billing_date string `json:"next_billing_date,omitempty"`
	On_free_trial bool `json:"on_free_trial"`
	Plan map[string]interface{} `json:"plan"`
	Unit_count int `json:"unit_count"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // Total number of users who prompted Copilot Chat on github.com at least once.
	Models []map[string]interface{} `json:"models,omitempty"` // List of model metrics for a custom models and the default model.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Id int `json:"id"`
	Commit_url string `json:"commit_url"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Label map[string]interface{} `json:"label"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alt_domain map[string]interface{} `json:"alt_domain,omitempty"`
	Domain map[string]interface{} `json:"domain,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Single_file string `json:"single_file,omitempty"`
	Token string `json:"token"` // The token used for authentication
	Expires_at string `json:"expires_at"` // The time this token expires
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Repositories []Repository `json:"repositories,omitempty"` // The repositories this token has access to
	Repository_selection string `json:"repository_selection,omitempty"` // Describe whether all repositories have been selected or there's a selection involved
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow Webhooksworkflow `json:"workflow"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Details interface{} `json:"details,omitempty"`
	TypeField string `json:"type,omitempty"` // The location type. Because secrets may be found in different types of resources (ie. code, comments, issues, pull requests, discussions), this field identifies the type of resource where the secret was found.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_count int `json:"total_count"`
	Confused int `json:"confused"`
	Hooray int `json:"hooray"`
	Rocket int `json:"rocket"`
	Eyes int `json:"eyes"`
	Url string `json:"url"`
	Field1 int `json:"-1"`
	Field1 int `json:"+1"`
	Heart int `json:"heart"`
	Laugh int `json:"laugh"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Status string `json:"status"`
}

// Webhookspullrequest5 represents the Webhookspullrequest5 schema from the OpenAPI specification
type Webhookspullrequest5 struct {
	Mergeable bool `json:"mergeable,omitempty"`
	Review_comments_url string `json:"review_comments_url"`
	Merged_at string `json:"merged_at"`
	Rebaseable bool `json:"rebaseable,omitempty"`
	Closed_at string `json:"closed_at"`
	Draft bool `json:"draft"` // Indicates whether or not the pull request is a draft.
	Comments int `json:"comments,omitempty"`
	Html_url string `json:"html_url"`
	Number int `json:"number"` // Number uniquely identifying the pull request within its repository.
	Commits int `json:"commits,omitempty"`
	Locked bool `json:"locked"`
	Auto_merge map[string]interface{} `json:"auto_merge"` // The status of auto merging a pull request.
	User map[string]interface{} `json:"user"`
	Head map[string]interface{} `json:"head"`
	Merged bool `json:"merged,omitempty"`
	Changed_files int `json:"changed_files,omitempty"`
	Links map[string]interface{} `json:"_links"`
	Title string `json:"title"` // The title of the pull request.
	State string `json:"state"` // State of this Pull Request. Either `open` or `closed`.
	Body string `json:"body"`
	Created_at string `json:"created_at"`
	Active_lock_reason string `json:"active_lock_reason"`
	Assignee map[string]interface{} `json:"assignee"`
	Mergeable_state string `json:"mergeable_state,omitempty"`
	Review_comment_url string `json:"review_comment_url"`
	Maintainer_can_modify bool `json:"maintainer_can_modify,omitempty"` // Indicates whether maintainers can modify the pull request.
	Statuses_url string `json:"statuses_url"`
	Issue_url string `json:"issue_url"`
	Labels []map[string]interface{} `json:"labels"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Comments_url string `json:"comments_url"`
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Deletions int `json:"deletions,omitempty"`
	Commits_url string `json:"commits_url"`
	Id int `json:"id"`
	Diff_url string `json:"diff_url"`
	Requested_reviewers []interface{} `json:"requested_reviewers"`
	Url string `json:"url"`
	Patch_url string `json:"patch_url"`
	Assignees []map[string]interface{} `json:"assignees"`
	Merged_by map[string]interface{} `json:"merged_by,omitempty"`
	Review_comments int `json:"review_comments,omitempty"`
	Base map[string]interface{} `json:"base"`
	Updated_at string `json:"updated_at"`
	Additions int `json:"additions,omitempty"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Node_id string `json:"node_id"`
	Requested_teams []map[string]interface{} `json:"requested_teams"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Deployment GeneratedType `json:"deployment,omitempty"` // A deployment created as the result of an Actions check run from a workflow that references an environment
	Details_url string `json:"details_url"`
	Status string `json:"status"` // The phase of the lifecycle that the check is currently in. Statuses of waiting, requested, and pending are reserved for GitHub Actions check runs.
	Url string `json:"url"`
	App GeneratedType `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Name string `json:"name"` // The name of the check.
	Output map[string]interface{} `json:"output"`
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being checked.
	External_id string `json:"external_id"`
	Started_at string `json:"started_at"`
	Check_suite map[string]interface{} `json:"check_suite"`
	Conclusion string `json:"conclusion"`
	Completed_at string `json:"completed_at"`
	Id int64 `json:"id"` // The id of the check.
	Html_url string `json:"html_url"`
	Pull_requests []GeneratedType `json:"pull_requests"` // Pull requests that are open with a `head_sha` or `head_branch` that matches the check. The returned pull requests do not necessarily indicate pull requests that triggered the check.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Events []string `json:"events"` // The list of events for the GitHub app. Note that the `installation_target`, `security_advisory`, and `meta` events are not included because they are global events and not specific to an installation.
	Installations_count int `json:"installations_count,omitempty"` // The number of installations associated with the GitHub app. Only returned when the integration is requesting details about itself.
	Client_id string `json:"client_id,omitempty"`
	Id int `json:"id"` // Unique identifier of the GitHub app
	Permissions map[string]interface{} `json:"permissions"` // The set of permissions for the GitHub app
	Description string `json:"description"`
	Owner interface{} `json:"owner"`
	Updated_at string `json:"updated_at"`
	External_url string `json:"external_url"`
	Html_url string `json:"html_url"`
	Name string `json:"name"` // The name of the GitHub app
	Node_id string `json:"node_id"`
	Slug string `json:"slug,omitempty"` // The slug name of the GitHub app
	Created_at string `json:"created_at"`
}

// Webhooksreview represents the Webhooksreview schema from the OpenAPI specification
type Webhooksreview struct {
	Id int `json:"id"` // Unique identifier of the review
	Body string `json:"body"` // The text of the review.
	Commit_id string `json:"commit_id"` // A commit SHA for the review.
	State string `json:"state"`
	Links map[string]interface{} `json:"_links"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Submitted_at string `json:"submitted_at"`
	Node_id string `json:"node_id"`
	Pull_request_url string `json:"pull_request_url"`
	User map[string]interface{} `json:"user"`
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Project_column Webhooksprojectcolumn `json:"project_column"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Security_advisory Webhookssecurityadvisory `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Date string `json:"date,omitempty"`
	Email string `json:"email,omitempty"`
	Name string `json:"name,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Registry_package map[string]interface{} `json:"registry_package"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	From string `json:"from"`
	To string `json:"to"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Changes Webhookschanges `json:"changes"` // The changes to the comment.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Action string `json:"action"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Comment Webhooksreviewcomment `json:"comment"` // The [comment](https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request) itself.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow Webhooksworkflow `json:"workflow"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Deployment map[string]interface{} `json:"deployment"` // The [deployment](https://docs.github.com/rest/deployments/deployments#list-deployments).
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_ruleset GeneratedType `json:"repository_ruleset"` // A set of rules to apply when specified conditions are met.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// Webhooksissue2 represents the Webhooksissue2 schema from the OpenAPI specification
type Webhooksissue2 struct {
	Updated_at string `json:"updated_at"`
	Node_id string `json:"node_id"`
	Assignees []map[string]interface{} `json:"assignees"`
	Url string `json:"url"` // URL for the issue
	Title string `json:"title"` // Title of the issue
	Reactions map[string]interface{} `json:"reactions"`
	Milestone map[string]interface{} `json:"milestone"` // A collection of related issues and pull requests.
	Id int64 `json:"id"`
	Labels_url string `json:"labels_url"`
	Repository_url string `json:"repository_url"`
	Sub_issues_summary map[string]interface{} `json:"sub_issues_summary,omitempty"`
	Created_at string `json:"created_at"`
	TypeField GeneratedType `json:"type,omitempty"` // The type of issue.
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Comments int `json:"comments"`
	Comments_url string `json:"comments_url"`
	Draft bool `json:"draft,omitempty"`
	Labels []map[string]interface{} `json:"labels,omitempty"`
	State string `json:"state,omitempty"` // State of the issue; either 'open' or 'closed'
	Body string `json:"body"` // Contents of the issue
	Html_url string `json:"html_url"`
	Performed_via_github_app map[string]interface{} `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Assignee map[string]interface{} `json:"assignee,omitempty"`
	Closed_at string `json:"closed_at"`
	State_reason string `json:"state_reason,omitempty"`
	User map[string]interface{} `json:"user"`
	Locked bool `json:"locked,omitempty"`
	Active_lock_reason string `json:"active_lock_reason"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Events_url string `json:"events_url"`
	Number int `json:"number"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Classifications []string `json:"classifications,omitempty"` // Classifications that have been applied to the file that triggered the alert. For example identifying it as documentation, or a generated file.
	Html_url string `json:"html_url,omitempty"`
	Location GeneratedType `json:"location,omitempty"` // Describe a region within a file for the alert.
	Ref string `json:"ref,omitempty"` // The Git reference, formatted as `refs/pull/<number>/merge`, `refs/pull/<number>/head`, `refs/heads/<branch name>` or simply `<branch name>`.
	Environment string `json:"environment,omitempty"` // Identifies the variable values associated with the environment in which the analysis that generated this alert instance was performed, such as the language that was analyzed.
	Message map[string]interface{} `json:"message,omitempty"`
	Analysis_key string `json:"analysis_key,omitempty"` // Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
	Category string `json:"category,omitempty"` // Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
	Commit_sha string `json:"commit_sha,omitempty"`
	State string `json:"state,omitempty"` // State of a code scanning alert.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int64 `json:"id"`
	Number int `json:"number"`
	Url string `json:"url"`
	Base map[string]interface{} `json:"base"`
	Head map[string]interface{} `json:"head"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
}

// Webhookschanges8 represents the Webhookschanges8 schema from the OpenAPI specification
type Webhookschanges8 struct {
	Tier map[string]interface{} `json:"tier"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // Repositories in which users used Copilot for Pull Requests to generate pull request summaries
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // The number of users who used Copilot for Pull Requests on github.com to generate a pull request summary at least once.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField GeneratedType `json:"type"` // The type of issue.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
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

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
	Protected bool `json:"protected"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parents []map[string]interface{} `json:"parents"`
	Verification map[string]interface{} `json:"verification"`
	Committer map[string]interface{} `json:"committer"` // Identifying information for the git-user
	Sha string `json:"sha"` // SHA for the commit
	Tree map[string]interface{} `json:"tree"`
	Message string `json:"message"` // Message describing the purpose of the commit
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Author map[string]interface{} `json:"author"` // Identifying information for the git-user
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ref string `json:"ref,omitempty"` // The ref name that the evaluation ran on.
	Repository_name string `json:"repository_name,omitempty"` // The name of the repository without the `.git` extension.
	Rule_evaluations []map[string]interface{} `json:"rule_evaluations,omitempty"` // Details on the evaluated rules.
	Actor_id int `json:"actor_id,omitempty"` // The number that identifies the user.
	Evaluation_result string `json:"evaluation_result,omitempty"` // The result of the rule evaluations for rules with the `active` and `evaluate` enforcement statuses, demonstrating whether rules would pass or fail if all rules in the rule suite were `active`. Null if no rules with `evaluate` enforcement status were run.
	Id int `json:"id,omitempty"` // The unique identifier of the rule insight.
	Pushed_at string `json:"pushed_at,omitempty"`
	Repository_id int `json:"repository_id,omitempty"` // The ID of the repository associated with the rule evaluation.
	Actor_name string `json:"actor_name,omitempty"` // The handle for the GitHub user account.
	After_sha string `json:"after_sha,omitempty"` // The last commit sha in the push evaluation.
	Result string `json:"result,omitempty"` // The result of the rule evaluations for rules with the `active` enforcement status.
	Before_sha string `json:"before_sha,omitempty"` // The first commit sha before the push evaluation.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Count int `json:"count"`
	Uniques int `json:"uniques"`
	Clones []Traffic `json:"clones"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Members_url string `json:"members_url"`
	Repositories_url string `json:"repositories_url"`
	Slug string `json:"slug"`
	Notification_setting string `json:"notification_setting,omitempty"` // The notification setting the team has set
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Url string `json:"url"` // URL for the team
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Name string `json:"name"` // Name of the team
	Node_id string `json:"node_id"`
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the team
	Description string `json:"description"` // Description of the team
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name of the variable.
	Updated_at string `json:"updated_at"` // The date and time at which the variable was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Value string `json:"value"` // The value of the variable.
	Created_at string `json:"created_at"` // The date and time at which the variable was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Hooks_url string `json:"hooks_url"`
	Id int `json:"id"`
	Public_members_url string `json:"public_members_url"`
	Repos_url string `json:"repos_url"`
	Description string `json:"description"`
	Login string `json:"login"`
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Members_url string `json:"members_url"`
	Events_url string `json:"events_url"`
	Issues_url string `json:"issues_url"`
	Node_id string `json:"node_id"`
}

// Webhooksmembership represents the Webhooksmembership schema from the OpenAPI specification
type Webhooksmembership struct {
	Organization_url string `json:"organization_url"`
	Role string `json:"role"`
	State string `json:"state"`
	Url string `json:"url"`
	User map[string]interface{} `json:"user"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue_title_url string `json:"issue_title_url"` // The API URL to get the issue where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"` // The date and time the campaign was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Created_at string `json:"created_at"` // The date and time the campaign was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Published_at string `json:"published_at,omitempty"` // The date and time the campaign was published, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Description string `json:"description"` // The campaign description
	Closed_at string `json:"closed_at,omitempty"` // The date and time the campaign was closed, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ. Will be null if the campaign is still open.
	Contact_link string `json:"contact_link"` // The contact link of the campaign.
	Ends_at string `json:"ends_at"` // The date and time the campaign has ended, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name,omitempty"` // The campaign name
	Number int `json:"number"` // The number of the newly created campaign
	Alert_stats map[string]interface{} `json:"alert_stats,omitempty"`
	Managers []GeneratedType `json:"managers"` // The campaign managers
	State string `json:"state"` // Indicates whether a campaign is open or closed
	Team_managers []Team `json:"team_managers,omitempty"` // The campaign team managers
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester Webhooksuser `json:"requester,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // The changes to the project if the action was `edited`.
}

// Metadata represents the Metadata schema from the OpenAPI specification
type Metadata struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Id int `json:"id"`
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Created_at string `json:"created_at"`
	Commit_url string `json:"commit_url"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Requested_reviewer GeneratedType `json:"requested_reviewer,omitempty"` // A GitHub user.
	Event string `json:"event"`
	Review_requester GeneratedType `json:"review_requester"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url,omitempty"` // The URL of the configuration
	Code_scanning_default_setup string `json:"code_scanning_default_setup,omitempty"` // The enablement status of code scanning default setup
	Secret_scanning string `json:"secret_scanning,omitempty"` // The enablement status of secret scanning
	Html_url string `json:"html_url,omitempty"` // The URL of the configuration
	Secret_scanning_delegated_bypass string `json:"secret_scanning_delegated_bypass,omitempty"` // The enablement status of secret scanning delegated bypass
	Secret_scanning_non_provider_patterns string `json:"secret_scanning_non_provider_patterns,omitempty"` // The enablement status of secret scanning non-provider patterns
	Secret_scanning_validity_checks string `json:"secret_scanning_validity_checks,omitempty"` // The enablement status of secret scanning validity checks
	Code_scanning_default_setup_options map[string]interface{} `json:"code_scanning_default_setup_options,omitempty"` // Feature options for code scanning default setup
	Dependabot_alerts string `json:"dependabot_alerts,omitempty"` // The enablement status of Dependabot alerts
	Secret_scanning_delegated_alert_dismissal string `json:"secret_scanning_delegated_alert_dismissal,omitempty"` // The enablement status of secret scanning delegated alert dismissal
	Code_scanning_options map[string]interface{} `json:"code_scanning_options,omitempty"` // Feature options for code scanning
	Private_vulnerability_reporting string `json:"private_vulnerability_reporting,omitempty"` // The enablement status of private vulnerability reporting
	Updated_at string `json:"updated_at,omitempty"`
	Code_scanning_delegated_alert_dismissal string `json:"code_scanning_delegated_alert_dismissal,omitempty"` // The enablement status of code scanning delegated alert dismissal
	Dependabot_security_updates string `json:"dependabot_security_updates,omitempty"` // The enablement status of Dependabot security updates
	Dependency_graph string `json:"dependency_graph,omitempty"` // The enablement status of Dependency Graph
	Created_at string `json:"created_at,omitempty"`
	Enforcement string `json:"enforcement,omitempty"` // The enforcement status for a security configuration
	Secret_scanning_delegated_bypass_options map[string]interface{} `json:"secret_scanning_delegated_bypass_options,omitempty"` // Feature options for secret scanning delegated bypass
	Description string `json:"description,omitempty"` // A description of the code security configuration
	Secret_scanning_generic_secrets string `json:"secret_scanning_generic_secrets,omitempty"` // The enablement status of Copilot secret scanning
	Secret_scanning_push_protection string `json:"secret_scanning_push_protection,omitempty"` // The enablement status of secret scanning push protection
	Advanced_security string `json:"advanced_security,omitempty"` // The enablement status of GitHub Advanced Security
	Dependency_graph_autosubmit_action string `json:"dependency_graph_autosubmit_action,omitempty"` // The enablement status of Automatic dependency submission
	Dependency_graph_autosubmit_action_options map[string]interface{} `json:"dependency_graph_autosubmit_action_options,omitempty"` // Feature options for Automatic dependency submission
	Id int `json:"id,omitempty"` // The ID of the code security configuration
	Name string `json:"name,omitempty"` // The name of the code security configuration. Must be unique within the organization.
	Target_type string `json:"target_type,omitempty"` // The type of the code security configuration.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body_version string `json:"body_version"` // The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
	Comments_url string `json:"comments_url"`
	Body_html string `json:"body_html"`
	Updated_at string `json:"updated_at"`
	Html_url string `json:"html_url"`
	Body string `json:"body"` // The main text of the discussion.
	Reactions GeneratedType `json:"reactions,omitempty"`
	Created_at string `json:"created_at"`
	Number int `json:"number"` // The unique sequence number of a team discussion.
	Pinned bool `json:"pinned"` // Whether or not this discussion should be pinned for easy retrieval.
	Team_url string `json:"team_url"`
	Title string `json:"title"` // The title of the discussion.
	Comments_count int `json:"comments_count"`
	Last_edited_at string `json:"last_edited_at"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Private bool `json:"private"` // Whether or not this discussion should be restricted to team members and organization owners.
	Author GeneratedType `json:"author"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Content_node_id string `json:"content_node_id"`
	Created_at string `json:"created_at"`
	Id float64 `json:"id"`
	Project_node_id string `json:"project_node_id,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Content_type string `json:"content_type"` // The type of content tracked in a project item
	Creator GeneratedType `json:"creator,omitempty"` // A GitHub user.
	Updated_at string `json:"updated_at"`
	Archived_at string `json:"archived_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease1 `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion_title_url string `json:"discussion_title_url"` // The URL to the discussion where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Label Webhookslabel `json:"label,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at,omitempty"` // The time the issue type created.
	Description string `json:"description"` // The description of the issue type.
	Id int `json:"id"` // The unique identifier of the issue type.
	Is_enabled bool `json:"is_enabled,omitempty"` // The enabled state of the issue type.
	Name string `json:"name"` // The name of the issue type.
	Node_id string `json:"node_id"` // The node identifier of the issue type.
	Updated_at string `json:"updated_at,omitempty"` // The time the issue type last updated.
	Color string `json:"color,omitempty"` // The color of the issue type.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repos_url string `json:"repos_url"`
	Description string `json:"description"`
	Members_url string `json:"members_url"`
	Node_id string `json:"node_id"`
	Public_members_url string `json:"public_members_url"`
	Events_url string `json:"events_url"`
	Avatar_url string `json:"avatar_url"`
	Issues_url string `json:"issues_url"`
	Login string `json:"login"`
	Url string `json:"url"`
	Hooks_url string `json:"hooks_url"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// Contributor represents the Contributor schema from the OpenAPI specification
type Contributor struct {
	Gists_url string `json:"gists_url,omitempty"`
	Id int `json:"id,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
	Login string `json:"login,omitempty"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Contributions int `json:"contributions"`
	Repos_url string `json:"repos_url,omitempty"`
	Email string `json:"email,omitempty"`
	Url string `json:"url,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Organizations_url string `json:"organizations_url,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	TypeField string `json:"type"`
	Name string `json:"name,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Remote_name string `json:"remote_name"`
	Url string `json:"url"`
	Email string `json:"email"`
	Id int `json:"id"`
	Import_url string `json:"import_url"`
	Name string `json:"name"`
	Remote_id string `json:"remote_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Thread map[string]interface{} `json:"thread"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Definition map[string]interface{} `json:"definition"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhookscomment `json:"comment"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Expires_at string `json:"expires_at,omitempty"`
	State string `json:"state"`
	Description string `json:"description"`
	Domains []string `json:"domains"` // Array of the domain set and its alternate name (if it is configured)
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Review Webhooksreview `json:"review"` // The review that was affected.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_url string `json:"commit_url"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Dismissed_review map[string]interface{} `json:"dismissed_review"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Active_this_cycle int `json:"active_this_cycle,omitempty"` // The number of seats that have used Copilot during the current billing cycle.
	Added_this_cycle int `json:"added_this_cycle,omitempty"` // Seats added during the current billing cycle.
	Inactive_this_cycle int `json:"inactive_this_cycle,omitempty"` // The number of seats that have not used Copilot during the current billing cycle.
	Pending_cancellation int `json:"pending_cancellation,omitempty"` // The number of seats that are pending cancellation at the end of the current billing cycle.
	Pending_invitation int `json:"pending_invitation,omitempty"` // The number of users who have been invited to receive a Copilot seat through this organization.
	Total int `json:"total,omitempty"` // The total number of seats being billed for the organization as of the current billing cycle.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Assignee GeneratedType `json:"assignee,omitempty"` // A GitHub user.
	Milestone GeneratedType `json:"milestone,omitempty"` // Issue Event Milestone
	Rename GeneratedType `json:"rename,omitempty"` // Issue Event Rename
	Author_association string `json:"author_association,omitempty"` // How the author is associated with the repository.
	Issue GeneratedType `json:"issue,omitempty"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Url string `json:"url"`
	Assigner GeneratedType `json:"assigner,omitempty"` // A GitHub user.
	Node_id string `json:"node_id"`
	Project_card GeneratedType `json:"project_card,omitempty"` // Issue Event Project Card
	Created_at string `json:"created_at"`
	Requested_reviewer GeneratedType `json:"requested_reviewer,omitempty"` // A GitHub user.
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Lock_reason string `json:"lock_reason,omitempty"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Dismissed_review GeneratedType `json:"dismissed_review,omitempty"`
	Label GeneratedType `json:"label,omitempty"` // Issue Event Label
	Commit_url string `json:"commit_url"`
	Review_requester GeneratedType `json:"review_requester,omitempty"` // A GitHub user.
	Id int64 `json:"id"`
	Event string `json:"event"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Milestone Milestone `json:"milestone,omitempty"` // A collection of related issues and pull requests.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request Webhookspullrequest5 `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Actor GeneratedType `json:"actor,omitempty"` // A GitHub user.
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Source map[string]interface{} `json:"source"`
}

// Repository represents the Repository schema from the OpenAPI specification
type Repository struct {
	Trees_url string `json:"trees_url"`
	Commits_url string `json:"commits_url"`
	Updated_at string `json:"updated_at"`
	Languages_url string `json:"languages_url"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Name string `json:"name"` // The name of the repository.
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Open_issues int `json:"open_issues"`
	Subscribers_url string `json:"subscribers_url"`
	Git_commits_url string `json:"git_commits_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Html_url string `json:"html_url"`
	Code_search_index_status map[string]interface{} `json:"code_search_index_status,omitempty"` // The status of the code search index for this repository
	Master_branch string `json:"master_branch,omitempty"`
	Issue_comment_url string `json:"issue_comment_url"`
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Svn_url string `json:"svn_url"`
	Homepage string `json:"homepage"`
	Watchers_count int `json:"watchers_count"`
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Stargazers_url string `json:"stargazers_url"`
	Labels_url string `json:"labels_url"`
	Keys_url string `json:"keys_url"`
	Forks_count int `json:"forks_count"`
	Forks_url string `json:"forks_url"`
	Teams_url string `json:"teams_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Contributors_url string `json:"contributors_url"`
	Deployments_url string `json:"deployments_url"`
	Url string `json:"url"`
	Git_refs_url string `json:"git_refs_url"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Tags_url string `json:"tags_url"`
	Node_id string `json:"node_id"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Private bool `json:"private"` // Whether the repository is private or public.
	Contents_url string `json:"contents_url"`
	Topics []string `json:"topics,omitempty"`
	Notifications_url string `json:"notifications_url"`
	Compare_url string `json:"compare_url"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Subscription_url string `json:"subscription_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Starred_at string `json:"starred_at,omitempty"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Archive_url string `json:"archive_url"`
	Clone_url string `json:"clone_url"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Mirror_url string `json:"mirror_url"`
	Stargazers_count int `json:"stargazers_count"`
	Language string `json:"language"`
	Collaborators_url string `json:"collaborators_url"`
	Ssh_url string `json:"ssh_url"`
	Assignees_url string `json:"assignees_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Milestones_url string `json:"milestones_url"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Statuses_url string `json:"statuses_url"`
	Downloads_url string `json:"downloads_url"`
	Releases_url string `json:"releases_url"`
	Description string `json:"description"`
	Id int64 `json:"id"` // Unique identifier of the repository
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Merges_url string `json:"merges_url"`
	Git_tags_url string `json:"git_tags_url"`
	Events_url string `json:"events_url"`
	Git_url string `json:"git_url"`
	Forks int `json:"forks"`
	Hooks_url string `json:"hooks_url"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Issue_events_url string `json:"issue_events_url"`
	Issues_url string `json:"issues_url"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	License GeneratedType `json:"license"` // License Simple
	Full_name string `json:"full_name"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Blobs_url string `json:"blobs_url"`
	Created_at string `json:"created_at"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Has_pages bool `json:"has_pages"`
	Watchers int `json:"watchers"`
	Branches_url string `json:"branches_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Pushed_at string `json:"pushed_at"`
	Pulls_url string `json:"pulls_url"`
	Fork bool `json:"fork"`
	Comments_url string `json:"comments_url"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Open_issues_count int `json:"open_issues_count"`
}

// Webhookscomment represents the Webhookscomment schema from the OpenAPI specification
type Webhookscomment struct {
	Id int `json:"id"`
	Updated_at string `json:"updated_at"`
	User map[string]interface{} `json:"user"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body"`
	Html_url string `json:"html_url"`
	Parent_id int `json:"parent_id"`
	Repository_url string `json:"repository_url"`
	Child_comment_count int `json:"child_comment_count"`
	Discussion_id int `json:"discussion_id"`
	Node_id string `json:"node_id"`
	Reactions map[string]interface{} `json:"reactions"`
	Created_at string `json:"created_at"`
}

// Webhooksuser represents the Webhooksuser schema from the OpenAPI specification
type Webhooksuser struct {
	Repos_url string `json:"repos_url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Organizations_url string `json:"organizations_url,omitempty"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Url string `json:"url,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	TypeField string `json:"type,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Gists_url string `json:"gists_url,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Name string `json:"name,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
	Id int64 `json:"id"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Login string `json:"login"`
	Deleted bool `json:"deleted,omitempty"`
	Email string `json:"email,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Requestor Webhooksuser `json:"requestor"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Environment string `json:"environment"`
	Workflow_job_run map[string]interface{} `json:"workflow_job_run"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Since string `json:"since"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Reviewers []map[string]interface{} `json:"reviewers"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card map[string]interface{} `json:"project_card"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Git_refs_url string `json:"git_refs_url"`
	Languages_url string `json:"languages_url"`
	Default_branch string `json:"default_branch"`
	Issue_comment_url string `json:"issue_comment_url"`
	Trees_url string `json:"trees_url"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Commits_url string `json:"commits_url"`
	Has_discussions bool `json:"has_discussions"`
	Url string `json:"url"`
	Issue_events_url string `json:"issue_events_url"`
	Has_issues bool `json:"has_issues"`
	Collaborators_url string `json:"collaborators_url"`
	Git_url string `json:"git_url"`
	Teams_url string `json:"teams_url"`
	Is_template bool `json:"is_template,omitempty"`
	Events_url string `json:"events_url"`
	Parent Repository `json:"parent,omitempty"` // A repository on GitHub.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub user.
	Stargazers_url string `json:"stargazers_url"`
	Git_tags_url string `json:"git_tags_url"`
	Compare_url string `json:"compare_url"`
	Subscribers_count int `json:"subscribers_count"`
	Comments_url string `json:"comments_url"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Assignees_url string `json:"assignees_url"`
	Mirror_url string `json:"mirror_url"`
	Watchers int `json:"watchers"`
	Stargazers_count int `json:"stargazers_count"`
	Security_and_analysis GeneratedType `json:"security_and_analysis,omitempty"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Created_at string `json:"created_at"`
	Archived bool `json:"archived"`
	Git_commits_url string `json:"git_commits_url"`
	Branches_url string `json:"branches_url"`
	Description string `json:"description"`
	Statuses_url string `json:"statuses_url"`
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"`
	Contents_url string `json:"contents_url"`
	Has_wiki bool `json:"has_wiki"`
	Fork bool `json:"fork"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Forks int `json:"forks"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Id int64 `json:"id"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Name string `json:"name"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Notifications_url string `json:"notifications_url"`
	Open_issues int `json:"open_issues"`
	Pulls_url string `json:"pulls_url"`
	Clone_url string `json:"clone_url"`
	Milestones_url string `json:"milestones_url"`
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Pushed_at string `json:"pushed_at"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Releases_url string `json:"releases_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Homepage string `json:"homepage"`
	Topics []string `json:"topics,omitempty"`
	Svn_url string `json:"svn_url"`
	Has_pages bool `json:"has_pages"`
	Language string `json:"language"`
	Master_branch string `json:"master_branch,omitempty"`
	Forks_url string `json:"forks_url"`
	License GeneratedType `json:"license"` // License Simple
	Network_count int `json:"network_count"`
	Hooks_url string `json:"hooks_url"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is allowed.
	Subscription_url string `json:"subscription_url"`
	Updated_at string `json:"updated_at"`
	Private bool `json:"private"`
	Downloads_url string `json:"downloads_url"`
	Forks_count int `json:"forks_count"`
	Issues_url string `json:"issues_url"`
	Template_repository GeneratedType `json:"template_repository,omitempty"` // A repository on GitHub.
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"`
	Tags_url string `json:"tags_url"`
	Node_id string `json:"node_id"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"`
	Keys_url string `json:"keys_url"`
	Labels_url string `json:"labels_url"`
	Ssh_url string `json:"ssh_url"`
	Deployments_url string `json:"deployments_url"`
	Full_name string `json:"full_name"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Html_url string `json:"html_url"`
	Subscribers_url string `json:"subscribers_url"`
	Has_downloads bool `json:"has_downloads,omitempty"`
	Watchers_count int `json:"watchers_count"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Source Repository `json:"source,omitempty"` // A repository on GitHub.
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"`
	Merges_url string `json:"merges_url"`
	Archive_url string `json:"archive_url"`
	Contributors_url string `json:"contributors_url"`
	Has_projects bool `json:"has_projects"`
	Open_issues_count int `json:"open_issues_count"`
	Blobs_url string `json:"blobs_url"`
	Code_of_conduct GeneratedType `json:"code_of_conduct,omitempty"` // Code of Conduct Simple
}

// Webhookschanges represents the Webhookschanges schema from the OpenAPI specification
type Webhookschanges struct {
	Body map[string]interface{} `json:"body,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Membership Webhooksmembership `json:"membership,omitempty"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Required_linear_history map[string]interface{} `json:"required_linear_history,omitempty"`
	Required_signatures map[string]interface{} `json:"required_signatures,omitempty"`
	Restrictions GeneratedType `json:"restrictions,omitempty"` // Branch Restriction Policy
	Url string `json:"url"`
	Lock_branch map[string]interface{} `json:"lock_branch,omitempty"` // Whether to set the branch as read-only. If this is true, users will not be able to push to the branch.
	Required_status_checks GeneratedType `json:"required_status_checks,omitempty"` // Status Check Policy
	Allow_force_pushes map[string]interface{} `json:"allow_force_pushes,omitempty"`
	Block_creations map[string]interface{} `json:"block_creations,omitempty"`
	Required_pull_request_reviews map[string]interface{} `json:"required_pull_request_reviews,omitempty"`
	Allow_deletions map[string]interface{} `json:"allow_deletions,omitempty"`
	Allow_fork_syncing map[string]interface{} `json:"allow_fork_syncing,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow fork syncing. Set to `false` to prevent fork syncing.
	Enforce_admins map[string]interface{} `json:"enforce_admins,omitempty"`
	Required_conversation_resolution map[string]interface{} `json:"required_conversation_resolution,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Effective_date string `json:"effective_date"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the secret_scanning_alert_location.created JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description,omitempty"` // Description of the issue type.
	Is_enabled bool `json:"is_enabled"` // Whether or not the issue type is enabled at the organization level.
	Name string `json:"name"` // Name of the issue type.
	Color string `json:"color,omitempty"` // Color for the issue type.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Week int `json:"week"`
	Days []int `json:"days"`
	Total int `json:"total"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Plan GeneratedType `json:"plan"` // Marketplace Listing Plan
	Unit_count int `json:"unit_count"`
	Updated_at string `json:"updated_at"`
	Account GeneratedType `json:"account"`
	Billing_cycle string `json:"billing_cycle"`
	Free_trial_ends_on string `json:"free_trial_ends_on"`
	Next_billing_date string `json:"next_billing_date"`
	On_free_trial bool `json:"on_free_trial"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"` // The date and time of when the advisory was last updated, in ISO 8601 format.
	Html_url string `json:"html_url"` // The URL for the advisory.
	Closed_at string `json:"closed_at"` // The date and time of when the advisory was closed, in ISO 8601 format.
	Cwe_ids []string `json:"cwe_ids"` // A list of only the CWE IDs.
	Ghsa_id string `json:"ghsa_id"` // The GitHub Security Advisory ID.
	Url string `json:"url"` // The API URL for the advisory.
	Withdrawn_at string `json:"withdrawn_at"` // The date and time of when the advisory was withdrawn, in ISO 8601 format.
	Collaborating_users []GeneratedType `json:"collaborating_users"` // A list of users that collaborate on the advisory.
	Identifiers []map[string]interface{} `json:"identifiers"`
	Created_at string `json:"created_at"` // The date and time of when the advisory was created, in ISO 8601 format.
	Credits []map[string]interface{} `json:"credits"`
	State string `json:"state"` // The state of the advisory.
	Summary string `json:"summary"` // A short summary of the advisory.
	Private_fork interface{} `json:"private_fork"` // A temporary private fork of the advisory's repository for collaborating on a fix.
	Cvss_severities GeneratedType `json:"cvss_severities,omitempty"`
	Publisher interface{} `json:"publisher"` // The publisher of the advisory.
	Author interface{} `json:"author"` // The author of the advisory.
	Collaborating_teams []Team `json:"collaborating_teams"` // A list of teams that collaborate on the advisory.
	Cve_id string `json:"cve_id"` // The Common Vulnerabilities and Exposures (CVE) ID.
	Description string `json:"description"` // A detailed description of what the advisory entails.
	Published_at string `json:"published_at"` // The date and time of when the advisory was published, in ISO 8601 format.
	Submission map[string]interface{} `json:"submission"`
	Severity string `json:"severity"` // The severity of the advisory.
	Vulnerabilities []GeneratedType `json:"vulnerabilities"`
	Cvss map[string]interface{} `json:"cvss"`
	Cwes []map[string]interface{} `json:"cwes"`
	Credits_detailed []GeneratedType `json:"credits_detailed"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow Webhooksworkflow `json:"workflow"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Starred_at interface{} `json:"starred_at"` // The time the star was created. This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Will be `null` for the `deleted` action.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Assignee Webhooksusermannequin `json:"assignee,omitempty"`
	Number int `json:"number"` // The pull request number.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_property map[string]interface{} `json:"repository_property"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body string `json:"body"` // The comment text.
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Tree_id string `json:"tree_id"` // SHA for the commit's tree
	Author map[string]interface{} `json:"author"` // Information about the Git author
	Committer map[string]interface{} `json:"committer"` // Information about the Git committer
	Id string `json:"id"` // SHA for the commit
	Message string `json:"message"` // Message describing the purpose of the commit
	Timestamp string `json:"timestamp"` // Timestamp of the commit
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Forkee interface{} `json:"forkee"` // The created [`repository`](https://docs.github.com/rest/repos/repos#get-a-repository) resource.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
	Push_protection_bypass_request_reviewer_comment string `json:"push_protection_bypass_request_reviewer_comment,omitempty"` // An optional comment when reviewing a push protection bypass.
	Repository GeneratedType `json:"repository,omitempty"` // A GitHub repository.
	Resolution_comment string `json:"resolution_comment,omitempty"` // The comment that was optionally added when this alert was closed
	Multi_repo bool `json:"multi_repo,omitempty"` // Whether the detected secret was found in multiple repositories in the same organization or enterprise.
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Has_more_locations bool `json:"has_more_locations,omitempty"` // A boolean value representing whether or not the token in the alert was detected in more than one location.
	Number int `json:"number,omitempty"` // The security alert number.
	State string `json:"state,omitempty"` // Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Resolution string `json:"resolution,omitempty"` // **Required when the `state` is `resolved`.** The reason for resolving the alert.
	Validity string `json:"validity,omitempty"` // The token status as of the latest validity check.
	Push_protection_bypass_request_comment string `json:"push_protection_bypass_request_comment,omitempty"` // An optional comment when requesting a push protection bypass.
	First_location_detected GeneratedType `json:"first_location_detected,omitempty"` // Details on the location where the token was initially detected. This can be a commit, wiki commit, issue, discussion, pull request.
	Resolved_by GeneratedType `json:"resolved_by,omitempty"` // A GitHub user.
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
	Push_protection_bypass_request_reviewer GeneratedType `json:"push_protection_bypass_request_reviewer,omitempty"` // A GitHub user.
	Secret string `json:"secret,omitempty"` // The secret that was detected.
	Publicly_leaked bool `json:"publicly_leaked,omitempty"` // Whether the secret was publicly leaked.
	Push_protection_bypass_request_html_url string `json:"push_protection_bypass_request_html_url,omitempty"` // The URL to a push protection bypass request.
	Push_protection_bypassed_by GeneratedType `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Supported secret scanning patterns](https://docs.github.com/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	Is_base64_encoded bool `json:"is_base64_encoded,omitempty"` // A boolean value representing whether or not alert is base64 encoded
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow Webhooksworkflow `json:"workflow"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Views []Traffic `json:"views"`
	Count int `json:"count"`
	Uniques int `json:"uniques"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"`
	Reason string `json:"reason"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Security_advisory map[string]interface{} `json:"security_advisory"` // The details of the security advisory, including summary, description, and severity.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Login string `json:"login"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
	Avatar_url string `json:"avatar_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at,omitempty"` // **Closing down notice:** This field is no longer relevant and is closing down. Use the `created_at` field to determine when the assignee was last granted access to GitHub Copilot. Timestamp of when the assignee's GitHub Copilot access was last updated, in ISO 8601 format.
	Last_activity_at string `json:"last_activity_at,omitempty"` // Timestamp of user's last GitHub Copilot activity, in ISO 8601 format.
	Last_activity_editor string `json:"last_activity_editor,omitempty"` // Last editor that was used by the user for a GitHub Copilot completion.
	Assigning_team interface{} `json:"assigning_team,omitempty"` // The team through which the assignee is granted access to GitHub Copilot, if applicable.
	Created_at string `json:"created_at"` // Timestamp of when the assignee was last granted access to GitHub Copilot, in ISO 8601 format.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization.
	Plan_type string `json:"plan_type,omitempty"` // The Copilot plan of the organization, or the parent enterprise, when applicable.
	Assignee GeneratedType `json:"assignee,omitempty"` // A GitHub user.
	Pending_cancellation_date string `json:"pending_cancellation_date,omitempty"` // The pending cancellation date for the seat, in `YYYY-MM-DD` format. This will be null unless the assignee's Copilot access has been canceled during the current billing cycle. If the seat has been cancelled, this corresponds to the start of the organization's next billing cycle.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Events []string `json:"events"`
	Updated_at string `json:"updated_at"`
	Id int `json:"id"`
	Ping_url string `json:"ping_url"`
	TypeField string `json:"type"`
	Active bool `json:"active"`
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Deliveries_url string `json:"deliveries_url,omitempty"`
	Name string `json:"name"`
	Config map[string]interface{} `json:"config"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Property_name string `json:"property_name"` // The name of the property
	Value string `json:"value"` // The value assigned to the property
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Status string `json:"status,omitempty"`
	Documentation_url string `json:"documentation_url,omitempty"`
	Message string `json:"message,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Id int `json:"id"`
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Assigner GeneratedType `json:"assigner"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Node_id string `json:"node_id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Value_type string `json:"value_type"` // The type of the value for the property
	Values_editable_by string `json:"values_editable_by,omitempty"` // Who can edit the values of the property
	Allowed_values []string `json:"allowed_values,omitempty"` // An ordered list of the allowed values of the property. The property can have up to 200 allowed values.
	Default_value string `json:"default_value,omitempty"` // Default value of the property
	Description string `json:"description,omitempty"` // Short description of the property
	Required bool `json:"required,omitempty"` // Whether the property is required.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name of the hosted runner.
	Platform string `json:"platform"` // The operating system of the image.
	Public_ip_enabled bool `json:"public_ip_enabled"` // Whether public IP is enabled for the hosted runners.
	Public_ips []GeneratedType `json:"public_ips,omitempty"` // The public IP ranges when public IP is enabled for the hosted runners.
	Image_details GeneratedType `json:"image_details"` // Provides details of a hosted runner image
	Status string `json:"status"` // The status of the runner.
	Last_active_on string `json:"last_active_on,omitempty"` // The time at which the runner was last used, in ISO 8601 format.
	Maximum_runners int `json:"maximum_runners,omitempty"` // The maximum amount of hosted runners. Runners will not scale automatically above this number. Use this setting to limit your cost.
	Id int `json:"id"` // The unique identifier of the hosted runner.
	Runner_group_id int `json:"runner_group_id,omitempty"` // The unique identifier of the group that the hosted runner belongs to.
	Machine_size_details GeneratedType `json:"machine_size_details"` // Provides details of a particular machine spec.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Editors []map[string]interface{} `json:"editors,omitempty"`
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // Total number of users who prompted Copilot Chat in the IDE.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue2 `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.completed JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Custom_branch_policies bool `json:"custom_branch_policies"` // Whether only branches that match the specified name patterns can deploy to this environment. If `custom_branch_policies` is `true`, `protected_branches` must be `false`; if `custom_branch_policies` is `false`, `protected_branches` must be `true`.
	Protected_branches bool `json:"protected_branches"` // Whether only branches with branch protection rules can deploy to this environment. If `protected_branches` is `true`, `custom_branch_policies` must be `false`; if `protected_branches` is `false`, `custom_branch_policies` must be `true`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Target_ref string `json:"target_ref,omitempty"` // The Git reference of target branch for the commit. Branch needs to already exist. For more information, see "[Git References](https://git-scm.com/book/en/v2/Git-Internals-Git-References)" in the Git documentation.
	Message string `json:"message,omitempty"` // Commit message to be used.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card Webhooksprojectcard `json:"project_card"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_column Webhooksprojectcolumn `json:"project_column"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Rule Webhooksrule `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Usageitems []map[string]interface{} `json:"usageItems,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	Size int `json:"size"`
	Content string `json:"content,omitempty"`
	Encoding string `json:"encoding,omitempty"`
	Download_url string `json:"download_url"`
	Entries []map[string]interface{} `json:"entries,omitempty"`
	Html_url string `json:"html_url"`
	Name string `json:"name"`
	Links map[string]interface{} `json:"_links"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Git_url string `json:"git_url"`
	Path string `json:"path"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description"` // A detailed description of what the advisory impacts.
	Severity string `json:"severity,omitempty"` // The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
	Summary string `json:"summary"` // A short summary of the advisory.
	Start_private_fork bool `json:"start_private_fork,omitempty"` // Whether to create a temporary private fork of the repository to collaborate on a fix.
	Credits []map[string]interface{} `json:"credits,omitempty"` // A list of users receiving credit for their participation in the security advisory.
	Cve_id string `json:"cve_id,omitempty"` // The Common Vulnerabilities and Exposures (CVE) ID.
	Cwe_ids []string `json:"cwe_ids,omitempty"` // A list of Common Weakness Enumeration (CWE) IDs.
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities"` // A product affected by the vulnerability detailed in a repository security advisory.
	Cvss_vector_string string `json:"cvss_vector_string,omitempty"` // The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alerts_threshold string `json:"alerts_threshold"` // The severity level at which code scanning results that raise alerts block a reference update. For more information on alert severity levels, see "[About code scanning alerts](https://docs.github.com/code-security/code-scanning/managing-code-scanning-alerts/about-code-scanning-alerts#about-alert-severity-and-security-severity-levels)."
	Security_alerts_threshold string `json:"security_alerts_threshold"` // The severity level at which code scanning results that raise security alerts block a reference update. For more information on security severity levels, see "[About code scanning alerts](https://docs.github.com/code-security/code-scanning/managing-code-scanning-alerts/about-code-scanning-alerts#about-alert-severity-and-security-severity-levels)."
	Tool string `json:"tool"` // The name of a code scanning tool
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Changes map[string]interface{} `json:"changes"` // The changes to the issue.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label,omitempty"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	New_property_values []GeneratedType `json:"new_property_values"` // The new custom property values for the repository.
	Old_property_values []GeneratedType `json:"old_property_values"` // The old custom property values for the repository.
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Milestone map[string]interface{} `json:"milestone"`
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
}

// Codespace represents the Codespace schema from the OpenAPI specification
type Codespace struct {
	Pulls_url string `json:"pulls_url"` // API URL for the Pull Request associated with this codespace, if any.
	Devcontainer_path string `json:"devcontainer_path,omitempty"` // Path to devcontainer.json from repo root used to create Codespace.
	Id int64 `json:"id"`
	Prebuild bool `json:"prebuild"` // Whether the codespace was created from a prebuild.
	Web_url string `json:"web_url"` // URL to access this codespace on the web.
	Stop_url string `json:"stop_url"` // API URL to stop this codespace.
	Runtime_constraints map[string]interface{} `json:"runtime_constraints,omitempty"`
	Name string `json:"name"` // Automatically generated name of this codespace.
	Recent_folders []string `json:"recent_folders"`
	Last_known_stop_notice string `json:"last_known_stop_notice,omitempty"` // The text to display to a user when a codespace has been stopped for a potentially actionable reason.
	Url string `json:"url"` // API URL for this codespace.
	Retention_period_minutes int `json:"retention_period_minutes,omitempty"` // Duration in minutes after codespace has gone idle in which it will be deleted. Must be integer minutes between 0 and 43200 (30 days).
	Machine GeneratedType `json:"machine"` // A description of the machine powering a codespace.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Publish_url string `json:"publish_url,omitempty"` // API URL to publish this codespace to a new repository.
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Retention_expires_at string `json:"retention_expires_at,omitempty"` // When a codespace will be auto-deleted based on the "retention_period_minutes" and "last_used_at"
	Pending_operation bool `json:"pending_operation,omitempty"` // Whether or not a codespace has a pending async operation. This would mean that the codespace is temporarily unavailable. The only thing that you can do with a codespace in this state is delete it.
	Start_url string `json:"start_url"` // API URL to start this codespace.
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Last_used_at string `json:"last_used_at"` // Last known time this codespace was started.
	State string `json:"state"` // State of this codespace.
	Display_name string `json:"display_name,omitempty"` // Display name for this codespace.
	Idle_timeout_minutes int `json:"idle_timeout_minutes"` // The number of minutes of inactivity after which this codespace will be automatically stopped.
	Machines_url string `json:"machines_url"` // API URL to access available alternate machine types for this codespace.
	Idle_timeout_notice string `json:"idle_timeout_notice,omitempty"` // Text to show user when codespace idle timeout minutes has been overriden by an organization policy
	Pending_operation_disabled_reason string `json:"pending_operation_disabled_reason,omitempty"` // Text to show user when codespace is disabled by a pending operation
	Billable_owner GeneratedType `json:"billable_owner"` // A GitHub user.
	Environment_id string `json:"environment_id"` // UUID identifying this codespace's environment.
	Location string `json:"location"` // The initally assigned location of a new codespace.
	Git_status map[string]interface{} `json:"git_status"` // Details about the codespace's git repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Compare_url string `json:"compare_url"` // A template for the API URL to compare two commits or refs.
	Labels_url string `json:"labels_url"` // A template for the API URL to get information about labels of the repository.
	Forks_url string `json:"forks_url"` // The API URL to list the forks of the repository.
	Git_commits_url string `json:"git_commits_url"` // A template for the API URL to get information about Git commits of the repository.
	Stargazers_url string `json:"stargazers_url"` // The API URL to list the stargazers on the repository.
	Subscribers_url string `json:"subscribers_url"` // The API URL to list the subscribers on the repository.
	Issues_url string `json:"issues_url"` // A template for the API URL to get information about issues on the repository.
	Releases_url string `json:"releases_url"` // A template for the API URL to get information about releases on the repository.
	Downloads_url string `json:"downloads_url"` // The API URL to list the downloads on the repository.
	Events_url string `json:"events_url"` // The API URL to list the events of the repository.
	Git_refs_url string `json:"git_refs_url"` // A template for the API URL to get information about Git refs of the repository.
	Notifications_url string `json:"notifications_url"` // A template for the API URL to get information about notifications on the repository.
	Id int64 `json:"id"` // A unique identifier of the repository.
	Milestones_url string `json:"milestones_url"` // A template for the API URL to get information about milestones of the repository.
	Assignees_url string `json:"assignees_url"` // A template for the API URL to list the available assignees for issues in the repository.
	Contributors_url string `json:"contributors_url"` // A template for the API URL to list the contributors to the repository.
	Keys_url string `json:"keys_url"` // A template for the API URL to get information about deploy keys on the repository.
	Subscription_url string `json:"subscription_url"` // The API URL to subscribe to notifications for this repository.
	Languages_url string `json:"languages_url"` // The API URL to get information about the languages of the repository.
	Description string `json:"description"` // The repository description.
	Issue_comment_url string `json:"issue_comment_url"` // A template for the API URL to get information about issue comments on the repository.
	Issue_events_url string `json:"issue_events_url"` // A template for the API URL to get information about issue events on the repository.
	Teams_url string `json:"teams_url"` // The API URL to list the teams on the repository.
	Comments_url string `json:"comments_url"` // A template for the API URL to get information about comments on the repository.
	Commits_url string `json:"commits_url"` // A template for the API URL to get information about commits on the repository.
	Branches_url string `json:"branches_url"` // A template for the API URL to get information about branches in the repository.
	Tags_url string `json:"tags_url"` // The API URL to get information about tags on the repository.
	Deployments_url string `json:"deployments_url"` // The API URL to list the deployments of the repository.
	Fork bool `json:"fork"` // Whether the repository is a fork.
	Html_url string `json:"html_url"` // The URL to view the repository on GitHub.com.
	Collaborators_url string `json:"collaborators_url"` // A template for the API URL to get information about collaborators of the repository.
	Statuses_url string `json:"statuses_url"` // A template for the API URL to get information about statuses of a commit.
	Hooks_url string `json:"hooks_url"` // The API URL to list the hooks on the repository.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Trees_url string `json:"trees_url"` // A template for the API URL to create or retrieve a raw Git tree of the repository.
	Contents_url string `json:"contents_url"` // A template for the API URL to get the contents of the repository.
	Full_name string `json:"full_name"` // The full, globally unique, name of the repository.
	Name string `json:"name"` // The name of the repository.
	Merges_url string `json:"merges_url"` // The API URL to merge branches in the repository.
	Private bool `json:"private"` // Whether the repository is private.
	Pulls_url string `json:"pulls_url"` // A template for the API URL to get information about pull requests on the repository.
	Url string `json:"url"` // The URL to get more information about the repository from the GitHub API.
	Archive_url string `json:"archive_url"` // A template for the API URL to download the repository as an archive.
	Git_tags_url string `json:"git_tags_url"` // A template for the API URL to get information about Git tags of the repository.
	Node_id string `json:"node_id"` // The GraphQL identifier of the repository.
	Blobs_url string `json:"blobs_url"` // A template for the API URL to create or retrieve a raw Git blob in the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	After string `json:"after,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Head_sha string `json:"head_sha,omitempty"` // The SHA of the head commit that is being checked.
	Id int `json:"id,omitempty"`
	Repository GeneratedType `json:"repository,omitempty"` // Minimal Repository
	Status string `json:"status,omitempty"`
	Before string `json:"before,omitempty"`
	Head_branch string `json:"head_branch,omitempty"`
	Url string `json:"url,omitempty"`
	Conclusion string `json:"conclusion,omitempty"`
	App Integration `json:"app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Node_id string `json:"node_id,omitempty"`
	Pull_requests []GeneratedType `json:"pull_requests,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Repositories_url string `json:"repositories_url"`
	Slug string `json:"slug"`
	Description string `json:"description"` // Description of the team
	Node_id string `json:"node_id"`
	Url string `json:"url"` // URL for the team
	Notification_setting string `json:"notification_setting,omitempty"` // The notification setting the team has set
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Html_url string `json:"html_url"`
	Members_url string `json:"members_url"`
	Name string `json:"name"` // Name of the team
	Id int `json:"id"` // Unique identifier of the team
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Storage_gb int `json:"storage_gb"` // The available SSD storage for the machine spec.
	Cpu_cores int `json:"cpu_cores"` // The number of cores.
	Id string `json:"id"` // The ID used for the `size` parameter when creating a new runner.
	Memory_gb int `json:"memory_gb"` // The available RAM for the machine spec.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Visibility string `json:"visibility"` // Which type of organization repositories have access to the private registry.
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the private registry configuration.
	Registry_type string `json:"registry_type"` // The registry type.
	Updated_at string `json:"updated_at"`
	Username string `json:"username,omitempty"` // The username to use when authenticating with the private registry.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Commit_id string `json:"commit_id"`
	Event string `json:"event"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Files map[string]interface{} `json:"files"`
	Health_percentage int `json:"health_percentage"`
	Updated_at string `json:"updated_at"`
	Content_reports_enabled bool `json:"content_reports_enabled,omitempty"`
	Description string `json:"description"`
	Documentation string `json:"documentation"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Help string `json:"help,omitempty"` // Detailed documentation for the rule as GitHub Flavored Markdown.
	Help_uri string `json:"help_uri,omitempty"` // A link to the documentation for the rule used to detect the alert.
	Security_severity_level string `json:"security_severity_level,omitempty"` // The security severity of the alert.
	Name string `json:"name,omitempty"` // The name of the rule used to detect the alert.
	Tags []string `json:"tags,omitempty"` // A set of tags applicable for the rule.
	Id string `json:"id,omitempty"` // A unique identifier for the rule used to detect the alert.
	Severity string `json:"severity,omitempty"` // The severity of the alert.
	Description string `json:"description,omitempty"` // A short description of the rule used to detect the alert.
	Full_description string `json:"full_description,omitempty"` // A description of the rule used to detect the alert.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Dependency map[string]interface{} `json:"dependency"` // Details for the vulnerable dependency.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Dismissed_comment string `json:"dismissed_comment"` // An optional comment associated with the alert's dismissal.
	Number int `json:"number"` // The security alert number.
	Dismissed_reason string `json:"dismissed_reason"` // The reason that the alert was dismissed.
	Updated_at string `json:"updated_at"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Security_vulnerability GeneratedType `json:"security_vulnerability"` // Details pertaining to one vulnerable version range for the advisory.
	Fixed_at string `json:"fixed_at"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Security_advisory GeneratedType `json:"security_advisory"` // Details for the GitHub Security Advisory.
	State string `json:"state"` // The state of the Dependabot alert.
	Auto_dismissed_at string `json:"auto_dismissed_at,omitempty"` // The time that the alert was auto-dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Repository GeneratedType `json:"repository"` // A GitHub repository.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Hosted_runners_url string `json:"hosted_runners_url,omitempty"`
	Inherited bool `json:"inherited"`
	Inherited_allows_public_repositories bool `json:"inherited_allows_public_repositories,omitempty"`
	Network_configuration_id string `json:"network_configuration_id,omitempty"` // The identifier of a hosted compute network configuration.
	Restricted_to_workflows bool `json:"restricted_to_workflows,omitempty"` // If `true`, the runner group will be restricted to running only the workflows specified in the `selected_workflows` array.
	Runners_url string `json:"runners_url"`
	Allows_public_repositories bool `json:"allows_public_repositories"`
	DefaultField bool `json:"default"`
	Id float64 `json:"id"`
	Name string `json:"name"`
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // Link to the selected repositories resource for this runner group. Not present unless visibility was set to `selected`
	Selected_workflows []string `json:"selected_workflows,omitempty"` // List of workflows the runner group should be allowed to run. This setting will be ignored unless `restricted_to_workflows` is set to `true`.
	Visibility string `json:"visibility"`
	Workflow_restrictions_read_only bool `json:"workflow_restrictions_read_only,omitempty"` // If `true`, the `restricted_to_workflows` and `selected_workflows` fields cannot be modified.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Key string `json:"key"`
	Title string `json:"title"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Full_name string `json:"full_name"` // The repository owner and name for the cache usage being shown.
	Active_caches_count int `json:"active_caches_count"` // The number of active caches in the repository.
	Active_caches_size_in_bytes int `json:"active_caches_size_in_bytes"` // The sum of the size in bytes of all the active cache items in the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Build map[string]interface{} `json:"build"` // The [List GitHub Pages builds](https://docs.github.com/rest/pages/pages#list-github-pages-builds) itself.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Id int `json:"id"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Days_left_in_billing_cycle int `json:"days_left_in_billing_cycle"` // Numbers of days left in billing cycle.
	Estimated_paid_storage_for_month int `json:"estimated_paid_storage_for_month"` // Estimated storage space (GB) used in billing cycle.
	Estimated_storage_for_month int `json:"estimated_storage_for_month"` // Estimated sum of free and paid storage space (GB) used in billing cycle.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action performed. Can be `created`.
	Comment map[string]interface{} `json:"comment"` // The [commit comment](${externalDocsUpapp/api/description/components/schemas/webhooks/issue-comment-created.yamlrl}/rest/commits/comments#get-a-commit-comment) resource.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes,omitempty"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_status_update GeneratedType `json:"projects_v2_status_update"` // An status update belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
	Node_id string `json:"node_id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
	Action string `json:"action"`
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
	Parent_issue_repo Repository `json:"parent_issue_repo"` // A repository on GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhooksreviewcomment `json:"comment"` // The [comment](https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request) itself.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Starred_at string `json:"starred_at"` // The time the star was created. This is a timestamp in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`. Will be `null` for the `deleted` action.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id,omitempty"`
	Event string `json:"event,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Comments []GeneratedType `json:"comments,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Protection GeneratedType `json:"protection"` // Branch Protection
	Protection_url string `json:"protection_url"`
	Required_approving_review_count int `json:"required_approving_review_count,omitempty"`
	Links map[string]interface{} `json:"_links"`
	Commit Commit `json:"commit"` // Commit
	Name string `json:"name"`
	Pattern string `json:"pattern,omitempty"`
	Protected bool `json:"protected"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment) itself.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) the comment belongs to.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Answer Webhooksanswer `json:"answer"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Visibility string `json:"visibility"` // The type of repositories in the organization that the secret is visible to
	Created_at string `json:"created_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the secret
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // The API URL at which the list of repositories this secret is visible to can be retrieved
	Updated_at string `json:"updated_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Archived bool `json:"archived,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Network_count int `json:"network_count,omitempty"`
	Has_downloads bool `json:"has_downloads,omitempty"`
	Has_issues bool `json:"has_issues,omitempty"`
	Languages_url string `json:"languages_url"`
	Code_of_conduct GeneratedType `json:"code_of_conduct,omitempty"` // Code Of Conduct
	Svn_url string `json:"svn_url,omitempty"`
	Size int `json:"size,omitempty"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Git_tags_url string `json:"git_tags_url"`
	Topics []string `json:"topics,omitempty"`
	Forks_url string `json:"forks_url"`
	Default_branch string `json:"default_branch,omitempty"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Commits_url string `json:"commits_url"`
	Has_pages bool `json:"has_pages,omitempty"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Id int64 `json:"id"`
	Assignees_url string `json:"assignees_url"`
	Tags_url string `json:"tags_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Events_url string `json:"events_url"`
	Full_name string `json:"full_name"`
	Role_name string `json:"role_name,omitempty"`
	Branches_url string `json:"branches_url"`
	Notifications_url string `json:"notifications_url"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Labels_url string `json:"labels_url"`
	Updated_at string `json:"updated_at,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Clone_url string `json:"clone_url,omitempty"`
	Git_refs_url string `json:"git_refs_url"`
	Security_and_analysis GeneratedType `json:"security_and_analysis,omitempty"`
	Open_issues_count int `json:"open_issues_count,omitempty"`
	Releases_url string `json:"releases_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Git_commits_url string `json:"git_commits_url"`
	Trees_url string `json:"trees_url"`
	Contributors_url string `json:"contributors_url"`
	Watchers int `json:"watchers,omitempty"`
	Ssh_url string `json:"ssh_url,omitempty"`
	Description string `json:"description"`
	Subscribers_url string `json:"subscribers_url"`
	Teams_url string `json:"teams_url"`
	Node_id string `json:"node_id"`
	Private bool `json:"private"`
	Watchers_count int `json:"watchers_count,omitempty"`
	Name string `json:"name"`
	Mirror_url string `json:"mirror_url,omitempty"`
	Open_issues int `json:"open_issues,omitempty"`
	Collaborators_url string `json:"collaborators_url"`
	Subscription_url string `json:"subscription_url"`
	Stargazers_count int `json:"stargazers_count,omitempty"`
	Contents_url string `json:"contents_url"`
	Downloads_url string `json:"downloads_url"`
	Forks_count int `json:"forks_count,omitempty"`
	Has_projects bool `json:"has_projects,omitempty"`
	Keys_url string `json:"keys_url"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Html_url string `json:"html_url"`
	Statuses_url string `json:"statuses_url"`
	Compare_url string `json:"compare_url"`
	Has_wiki bool `json:"has_wiki,omitempty"`
	Deployments_url string `json:"deployments_url"`
	Forks int `json:"forks,omitempty"`
	Url string `json:"url"`
	Pulls_url string `json:"pulls_url"`
	Hooks_url string `json:"hooks_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Issues_url string `json:"issues_url"`
	Fork bool `json:"fork"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Comments_url string `json:"comments_url"`
	Language string `json:"language,omitempty"`
	Git_url string `json:"git_url,omitempty"`
	Pushed_at string `json:"pushed_at,omitempty"`
	Blobs_url string `json:"blobs_url"`
	Stargazers_url string `json:"stargazers_url"`
	Visibility string `json:"visibility,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	Is_template bool `json:"is_template,omitempty"`
	Issue_comment_url string `json:"issue_comment_url"`
	License map[string]interface{} `json:"license,omitempty"`
	Milestones_url string `json:"milestones_url"`
	Merges_url string `json:"merges_url"`
	Archive_url string `json:"archive_url"`
	Issue_events_url string `json:"issue_events_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Dismiss_stale_reviews bool `json:"dismiss_stale_reviews"`
	Dismissal_restrictions map[string]interface{} `json:"dismissal_restrictions,omitempty"`
	Require_code_owner_reviews bool `json:"require_code_owner_reviews"`
	Require_last_push_approval bool `json:"require_last_push_approval,omitempty"` // Whether the most recent push must be approved by someone other than the person who pushed it.
	Required_approving_review_count int `json:"required_approving_review_count,omitempty"`
	Url string `json:"url,omitempty"`
	Bypass_pull_request_allowances map[string]interface{} `json:"bypass_pull_request_allowances,omitempty"` // Allow specific users, teams, or apps to bypass pull request requirements.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pages []map[string]interface{} `json:"pages"` // The pages that were updated.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id string `json:"id"` // The ID of the GitHub Pages deployment. This is the Git SHA of the deployed commit.
	Page_url string `json:"page_url"` // The URI to the deployed GitHub Pages.
	Preview_url string `json:"preview_url,omitempty"` // The URI to the deployed GitHub Pages preview.
	Status_url string `json:"status_url"` // The URI to monitor GitHub Pages deployment status.
}

// Webhooksrelease represents the Webhooksrelease schema from the OpenAPI specification
type Webhooksrelease struct {
	Created_at string `json:"created_at"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	Prerelease bool `json:"prerelease"` // Whether the release is identified as a prerelease or a full release.
	Tarball_url string `json:"tarball_url"`
	Discussion_url string `json:"discussion_url,omitempty"`
	Target_commitish string `json:"target_commitish"` // Specifies the commitish value that determines where the Git tag is created from.
	Upload_url string `json:"upload_url"`
	Node_id string `json:"node_id"`
	Author map[string]interface{} `json:"author"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Name string `json:"name"`
	Tag_name string `json:"tag_name"` // The name of the tag.
	Zipball_url string `json:"zipball_url"`
	Assets []map[string]interface{} `json:"assets"`
	Published_at string `json:"published_at"`
	Assets_url string `json:"assets_url"`
	Body string `json:"body"`
	Draft bool `json:"draft"` // Whether the release is a draft or published
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int64 `json:"id"` // Unique identifier of the repository
	Contents_url string `json:"contents_url"`
	Url string `json:"url"`
	Branches_url string `json:"branches_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Open_issues int `json:"open_issues"`
	Template_repository map[string]interface{} `json:"template_repository,omitempty"`
	Hooks_url string `json:"hooks_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Trees_url string `json:"trees_url"`
	Comments_url string `json:"comments_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Contributors_url string `json:"contributors_url"`
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Updated_at string `json:"updated_at"`
	Mirror_url string `json:"mirror_url"`
	Pulls_url string `json:"pulls_url"`
	Forks_count int `json:"forks_count"`
	Stargazers_url string `json:"stargazers_url"`
	Has_pages bool `json:"has_pages"`
	Fork bool `json:"fork"`
	Commits_url string `json:"commits_url"`
	Topics []string `json:"topics,omitempty"`
	Compare_url string `json:"compare_url"`
	Assignees_url string `json:"assignees_url"`
	Blobs_url string `json:"blobs_url"`
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Deployments_url string `json:"deployments_url"`
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Forks int `json:"forks"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Releases_url string `json:"releases_url"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Tags_url string `json:"tags_url"`
	Labels_url string `json:"labels_url"`
	Archived bool `json:"archived"` // Whether the repository is archived.
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Subscription_url string `json:"subscription_url"`
	License GeneratedType `json:"license"` // License Simple
	Node_id string `json:"node_id"`
	Watchers_count int `json:"watchers_count"`
	Archive_url string `json:"archive_url"`
	Full_name string `json:"full_name"`
	Forks_url string `json:"forks_url"`
	Issues_url string `json:"issues_url"`
	Homepage string `json:"homepage"`
	Master_branch string `json:"master_branch,omitempty"`
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Issue_comment_url string `json:"issue_comment_url"`
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Watchers int `json:"watchers"`
	Clone_url string `json:"clone_url"`
	Network_count int `json:"network_count,omitempty"`
	Stargazers_count int `json:"stargazers_count"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub user.
	Teams_url string `json:"teams_url"`
	Name string `json:"name"` // The name of the repository.
	Collaborators_url string `json:"collaborators_url"`
	Git_commits_url string `json:"git_commits_url"`
	Git_tags_url string `json:"git_tags_url"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Languages_url string `json:"languages_url"`
	Pushed_at string `json:"pushed_at"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Created_at string `json:"created_at"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Merges_url string `json:"merges_url"`
	Keys_url string `json:"keys_url"`
	Statuses_url string `json:"statuses_url"`
	Git_refs_url string `json:"git_refs_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Issue_events_url string `json:"issue_events_url"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Svn_url string `json:"svn_url"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Notifications_url string `json:"notifications_url"`
	Html_url string `json:"html_url"`
	Downloads_url string `json:"downloads_url"`
	Events_url string `json:"events_url"`
	Milestones_url string `json:"milestones_url"`
	Subscribers_url string `json:"subscribers_url"`
	Git_url string `json:"git_url"`
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Description string `json:"description"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Open_issues_count int `json:"open_issues_count"`
	Ssh_url string `json:"ssh_url"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Language string `json:"language"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Dependabot []string `json:"dependabot,omitempty"`
	Importer []string `json:"importer,omitempty"`
	Ssh_keys []string `json:"ssh_keys,omitempty"`
	Github_enterprise_importer []string `json:"github_enterprise_importer,omitempty"`
	Web []string `json:"web,omitempty"`
	Actions_macos []string `json:"actions_macos,omitempty"`
	Domains map[string]interface{} `json:"domains,omitempty"`
	Actions []string `json:"actions,omitempty"`
	Ssh_key_fingerprints map[string]interface{} `json:"ssh_key_fingerprints,omitempty"`
	Api []string `json:"api,omitempty"`
	Codespaces []string `json:"codespaces,omitempty"`
	Copilot []string `json:"copilot,omitempty"`
	Packages []string `json:"packages,omitempty"`
	Pages []string `json:"pages,omitempty"`
	Git []string `json:"git,omitempty"`
	Hooks []string `json:"hooks,omitempty"`
	Verifiable_password_authentication bool `json:"verifiable_password_authentication"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Roster_identifier string `json:"roster_identifier"` // Roster identifier of the student
	Assignment_name string `json:"assignment_name"` // Name of the assignment
	Github_username string `json:"github_username"` // GitHub username of the student
	Student_repository_name string `json:"student_repository_name"` // Name of the student's assignment repository
	Group_name string `json:"group_name,omitempty"` // If a group assignment, name of the group the student is in
	Points_awarded int `json:"points_awarded"` // Number of points awarded to the student
	Starter_code_url string `json:"starter_code_url"` // URL of the starter code for the assignment
	Student_repository_url string `json:"student_repository_url"` // URL of the student's assignment repository
	Submission_timestamp string `json:"submission_timestamp"` // Timestamp of the student's assignment submission
	Assignment_url string `json:"assignment_url"` // URL of the assignment
	Points_available int `json:"points_available"` // Number of points available for the assignment
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Size int `json:"size"`
	Encoding string `json:"encoding"`
	Target string `json:"target,omitempty"`
	Url string `json:"url"`
	Git_url string `json:"git_url"`
	Path string `json:"path"`
	Sha string `json:"sha"`
	Content string `json:"content"`
	Submodule_git_url string `json:"submodule_git_url,omitempty"`
	TypeField string `json:"type"`
	Links map[string]interface{} `json:"_links"`
	Download_url string `json:"download_url"`
	Html_url string `json:"html_url"`
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Statuses_url string `json:"statuses_url"`
	Id int `json:"id"` // Unique identifier of the deployment
	Task string `json:"task"` // Parameter to specify a task to execute
	Transient_environment bool `json:"transient_environment,omitempty"` // Specifies if the given environment is will no longer exist at some point in the future. Default: false.
	Node_id string `json:"node_id"`
	Production_environment bool `json:"production_environment,omitempty"` // Specifies if the given environment is one that end-users directly interact with. Default: false.
	Repository_url string `json:"repository_url"`
	Updated_at string `json:"updated_at"`
	Description string `json:"description"`
	Environment string `json:"environment"` // Name for the target deployment environment.
	Original_environment string `json:"original_environment,omitempty"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Thread map[string]interface{} `json:"thread"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_minutes_used int `json:"total_minutes_used"` // The sum of the free and paid GitHub Actions minutes used.
	Total_paid_minutes_used int `json:"total_paid_minutes_used"` // The total paid GitHub Actions minutes used.
	Included_minutes int `json:"included_minutes"` // The amount of free GitHub Actions minutes available.
	Minutes_used_breakdown map[string]interface{} `json:"minutes_used_breakdown"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"` // The changes to the collaborator permissions
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Member Webhooksuser `json:"member"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ldap_dn string `json:"ldap_dn,omitempty"` // Distinguished Name (DN) that team maps to within LDAP environment
	Node_id string `json:"node_id"`
	Description string `json:"description"`
	Members_url string `json:"members_url"`
	Parent GeneratedType `json:"parent,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Created_at string `json:"created_at"`
	Id int `json:"id"` // Unique identifier of the team
	Name string `json:"name"` // Name of the team
	Permission string `json:"permission"` // Permission that the team will have for its repositories
	Privacy string `json:"privacy,omitempty"` // The level of privacy this team should have
	Organization GeneratedType `json:"organization"` // Team Organization
	Slug string `json:"slug"`
	Html_url string `json:"html_url"`
	Notification_setting string `json:"notification_setting,omitempty"` // The notification setting the team has set
	Repos_count int `json:"repos_count"`
	Repositories_url string `json:"repositories_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"` // URL for the team
	Members_count int `json:"members_count"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes interface{} `json:"changes,omitempty"` // The changes made to the item may involve modifications in the item's fields and draft issue body. It includes altered values for text, number, date, single select, and iteration fields, along with the GraphQL node ID of the changed field.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request_title_url string `json:"pull_request_title_url"` // The API URL to get the pull request where the secret was detected.
}

// Webhooksissuecomment represents the Webhooksissuecomment schema from the OpenAPI specification
type Webhooksissuecomment struct {
	Id int64 `json:"id"` // Unique identifier of the issue comment
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Node_id string `json:"node_id"`
	Reactions map[string]interface{} `json:"reactions"`
	Body string `json:"body"` // Contents of the issue comment
	Created_at string `json:"created_at"`
	Issue_url string `json:"issue_url"`
	Updated_at string `json:"updated_at"`
	User map[string]interface{} `json:"user"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Html_url string `json:"html_url"`
	Url string `json:"url"` // URL for the issue comment
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Previous_marketplace_purchase Webhookspreviousmarketplacepurchase `json:"previous_marketplace_purchase,omitempty"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Action string `json:"action"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Effective_date string `json:"effective_date"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repositories_url string `json:"repositories_url"` // URL to the list of repositories requested to be accessed via fine-grained personal access token. Should only be followed when `repository_selection` is `subset`.
	Token_last_used_at string `json:"token_last_used_at"` // Date and time when the associated fine-grained personal access token was last used for authentication.
	Token_name string `json:"token_name"` // The name given to the user's token. This field can also be found in an organization's settings page for Active Tokens.
	Reason string `json:"reason"` // Reason for requesting access.
	Repository_selection string `json:"repository_selection"` // Type of repository selection requested.
	Token_expired bool `json:"token_expired"` // Whether the associated fine-grained personal access token has expired.
	Token_id int `json:"token_id"` // Unique identifier of the user's token. This field can also be found in audit log events and the organization's settings for their PAT grants.
	Created_at string `json:"created_at"` // Date and time when the request for access was created.
	Id int `json:"id"` // Unique identifier of the request for access via fine-grained personal access token. The `pat_request_id` used to review PAT requests.
	Permissions map[string]interface{} `json:"permissions"` // Permissions requested, categorized by type of permission.
	Token_expires_at string `json:"token_expires_at"` // Date and time when the associated fine-grained personal access token expires.
	Owner GeneratedType `json:"owner"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Comment map[string]interface{} `json:"comment"` // The [comment](https://docs.github.com/rest/pulls/comments#get-a-review-comment-for-a-pull-request) itself.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int64 `json:"id"` // Unique identifier of the webhook delivery.
	Installation_id int64 `json:"installation_id"` // The id of the GitHub App installation associated with this event.
	Status string `json:"status"` // Describes the response returned after attempting the delivery.
	Action string `json:"action"` // The type of activity for the event that triggered the delivery.
	Delivered_at string `json:"delivered_at"` // Time when the webhook delivery occurred.
	Redelivery bool `json:"redelivery"` // Whether the webhook delivery is a redelivery.
	Status_code int `json:"status_code"` // Status code received when delivery was made.
	Event string `json:"event"` // The event that triggered the delivery.
	Guid string `json:"guid"` // Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
	Repository_id int64 `json:"repository_id"` // The id of the repository associated with this event.
	Throttled_at string `json:"throttled_at,omitempty"` // Time when the webhook delivery was throttled.
	Duration float64 `json:"duration"` // Time spent delivering.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Workflow_job_runs []map[string]interface{} `json:"workflow_job_runs,omitempty"`
	Workflow_job_run Webhooksworkflowjobrun `json:"workflow_job_run,omitempty"`
	Comment string `json:"comment,omitempty"`
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Reviewers []map[string]interface{} `json:"reviewers,omitempty"`
	Workflow_run map[string]interface{} `json:"workflow_run"`
	Action string `json:"action"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Since string `json:"since"`
	Approver Webhooksapprover `json:"approver,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Login string `json:"login"`
	Avatar_url string `json:"avatar_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Token_name string `json:"token_name"` // The name given to the user's token. This field can also be found in an organization's settings page for Active Tokens.
	Permissions_upgraded map[string]interface{} `json:"permissions_upgraded"` // Requested permissions that elevate access for a previously approved request for access, categorized by type of permission.
	Token_id int `json:"token_id"` // Unique identifier of the user's token. This field can also be found in audit log events and the organization's settings for their PAT grants.
	Id int `json:"id"` // Unique identifier of the request for access via fine-grained personal access token. Used as the `pat_request_id` parameter in the list and review API calls.
	Token_expires_at string `json:"token_expires_at"` // Date and time when the associated fine-grained personal access token expires.
	Token_last_used_at string `json:"token_last_used_at"` // Date and time when the associated fine-grained personal access token was last used for authentication.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Permissions_result map[string]interface{} `json:"permissions_result"` // Permissions requested, categorized by type of permission. This field incorporates `permissions_added` and `permissions_upgraded`.
	Repository_count int `json:"repository_count"` // The number of repositories the token is requesting access to. This field is only populated when `repository_selection` is `subset`.
	Repository_selection string `json:"repository_selection"` // Type of repository selection requested.
	Created_at string `json:"created_at"` // Date and time when the request for access was created.
	Permissions_added map[string]interface{} `json:"permissions_added"` // New requested permissions, categorized by type of permission.
	Repositories []map[string]interface{} `json:"repositories"` // An array of repository objects the token is requesting access to. This field is only populated when `repository_selection` is `subset`.
	Token_expired bool `json:"token_expired"` // Whether the associated fine-grained personal access token has expired.
}

// Webhooksapprover represents the Webhooksapprover schema from the OpenAPI specification
type Webhooksapprover struct {
	Id int `json:"id,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Login string `json:"login,omitempty"`
	Gists_url string `json:"gists_url,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	TypeField string `json:"type,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Organizations_url string `json:"organizations_url,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Received_events_url string `json:"received_events_url,omitempty"`
	Url string `json:"url,omitempty"`
	Node_id string `json:"node_id,omitempty"`
}

// Classroom represents the Classroom schema from the OpenAPI specification
type Classroom struct {
	Id int `json:"id"` // Unique identifier of the classroom.
	Name string `json:"name"` // The name of the classroom.
	Organization GeneratedType `json:"organization"` // A GitHub organization.
	Url string `json:"url"` // The URL of the classroom on GitHub Classroom.
	Archived bool `json:"archived"` // Whether classroom is archived.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Rename map[string]interface{} `json:"rename"`
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.rerequested JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert Webhooksalert `json:"alert"` // The security alert of the vulnerable dependency.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Starred_url string `json:"starred_url"`
	Node_id string `json:"node_id"`
	Organizations_url string `json:"organizations_url"`
	User_view_type string `json:"user_view_type,omitempty"`
	Email string `json:"email,omitempty"`
	Received_events_url string `json:"received_events_url"`
	TypeField string `json:"type"`
	Role_name string `json:"role_name"`
	Avatar_url string `json:"avatar_url"`
	Following_url string `json:"following_url"`
	Events_url string `json:"events_url"`
	Name string `json:"name,omitempty"`
	Gists_url string `json:"gists_url"`
	Repos_url string `json:"repos_url"`
	Url string `json:"url"`
	Id int64 `json:"id"`
	Login string `json:"login"`
	Subscriptions_url string `json:"subscriptions_url"`
	Gravatar_id string `json:"gravatar_id"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Site_admin bool `json:"site_admin"`
	Followers_url string `json:"followers_url"`
	Html_url string `json:"html_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Strict bool `json:"strict"`
	Url string `json:"url"`
	Checks []map[string]interface{} `json:"checks"`
	Contexts []string `json:"contexts"`
	Contexts_url string `json:"contexts_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Title string `json:"title"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Limit string `json:"limit"` // The type of GitHub user that can comment, open issues, or create pull requests while the interaction limit is in effect.
	Origin string `json:"origin"`
	Expires_at string `json:"expires_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Oid string `json:"oid"`
	Path string `json:"path"`
	Ref_name string `json:"ref_name"`
	Size int `json:"size"`
}

// Webhooksprojectchanges represents the Webhooksprojectchanges schema from the OpenAPI specification
type Webhooksprojectchanges struct {
	Archived_at map[string]interface{} `json:"archived_at,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	End_column float64 `json:"end_column"` // The column at which the secret ends within the end line when the file is interpreted as 8BIT ASCII
	Start_line float64 `json:"start_line"` // Line number at which the secret starts in the file
	Blob_sha string `json:"blob_sha"` // SHA-1 hash ID of the associated blob
	Blob_url string `json:"blob_url"` // The API URL to get the associated blob resource
	End_line float64 `json:"end_line"` // Line number at which the secret ends in the file
	Path string `json:"path"` // The file path in the repository
	Commit_sha string `json:"commit_sha"` // SHA-1 hash ID of the associated commit
	Commit_url string `json:"commit_url"` // The API URL to get the associated commit resource
	Start_column float64 `json:"start_column"` // The column at which the secret starts within the start line when the file is interpreted as 8BIT ASCII
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Original_start_line int `json:"original_start_line,omitempty"` // The first line of the range for a multi-line comment.
	Path string `json:"path"` // The relative path of the file to which the comment applies.
	User GeneratedType `json:"user"` // A GitHub user.
	Position int `json:"position,omitempty"` // The line index in the diff to which the comment applies. This field is closing down; use `line` instead.
	Pull_request_review_id int64 `json:"pull_request_review_id"` // The ID of the pull request review to which the comment belongs.
	Links map[string]interface{} `json:"_links"`
	Id int64 `json:"id"` // The ID of the pull request review comment.
	Created_at string `json:"created_at"`
	Start_line int `json:"start_line,omitempty"` // The first line of the range for a multi-line comment.
	Commit_id string `json:"commit_id"` // The SHA of the commit to which the comment applies.
	Original_position int `json:"original_position,omitempty"` // The index of the original line in the diff to which the comment applies. This field is closing down; use `original_line` instead.
	Start_side string `json:"start_side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Html_url string `json:"html_url"` // HTML URL for the pull request review comment.
	Body_html string `json:"body_html,omitempty"`
	Side string `json:"side,omitempty"` // The side of the diff to which the comment applies. The side of the last line of the range for a multi-line comment
	Body_text string `json:"body_text,omitempty"`
	Updated_at string `json:"updated_at"`
	Diff_hunk string `json:"diff_hunk"` // The diff of the line that the comment refers to.
	Original_line int `json:"original_line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Line int `json:"line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Reactions GeneratedType `json:"reactions,omitempty"`
	Original_commit_id string `json:"original_commit_id"` // The SHA of the original commit to which the comment applies.
	Pull_request_url string `json:"pull_request_url"` // URL for the pull request that the review comment belongs to.
	Node_id string `json:"node_id"` // The node ID of the pull request review comment.
	Body string `json:"body"` // The text of the comment.
	In_reply_to_id int `json:"in_reply_to_id,omitempty"` // The comment ID to reply to.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Subject_type string `json:"subject_type,omitempty"` // The level at which the comment is targeted, can be a diff line or a file.
	Url string `json:"url"` // URL for the pull request review comment
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action that was performed.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issues_url string `json:"issues_url"`
	Node_id string `json:"node_id"`
	Events_url string `json:"events_url"`
	Id int `json:"id"`
	Repos_url string `json:"repos_url"`
	Avatar_url string `json:"avatar_url"`
	Public_members_url string `json:"public_members_url"`
	Url string `json:"url"`
	Hooks_url string `json:"hooks_url"`
	Login string `json:"login"`
	Members_url string `json:"members_url"`
	Description string `json:"description"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"` // The changes to the team if the action was `edited`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Errors []string `json:"errors,omitempty"` // Any errors that ocurred during processing of the delivery.
	Processing_status string `json:"processing_status,omitempty"` // `pending` files have not yet been processed, while `complete` means results from the SARIF have been stored. `failed` files have either not been processed at all, or could only be partially processed.
	Analyses_url string `json:"analyses_url,omitempty"` // The REST API URL for getting the analyses associated with the upload.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Slug string `json:"slug"` // The slugified name of the deployment protection rule integration.
	Id int `json:"id"` // The unique identifier of the deployment protection rule integration.
	Integration_url string `json:"integration_url"` // The URL for the endpoint to get details about the app.
	Node_id string `json:"node_id"` // The node ID for the deployment protection rule integration.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Links map[string]interface{} `json:"_links"`
	Git_url string `json:"git_url"`
	Sha string `json:"sha"`
	Download_url string `json:"download_url"`
	Html_url string `json:"html_url"`
	Submodule_git_url string `json:"submodule_git_url"`
	TypeField string `json:"type"`
	Name string `json:"name"`
	Path string `json:"path"`
	Size int `json:"size"`
}

// Release represents the Release schema from the OpenAPI specification
type Release struct {
	Body_text string `json:"body_text,omitempty"`
	Discussion_url string `json:"discussion_url,omitempty"` // The URL of the release discussion.
	Draft bool `json:"draft"` // true to create a draft (unpublished) release, false to create a published one.
	Mentions_count int `json:"mentions_count,omitempty"`
	Html_url string `json:"html_url"`
	Assets []GeneratedType `json:"assets"`
	Assets_url string `json:"assets_url"`
	Tag_name string `json:"tag_name"` // The name of the tag.
	Prerelease bool `json:"prerelease"` // Whether to identify the release as a prerelease or a full release.
	Tarball_url string `json:"tarball_url"`
	Upload_url string `json:"upload_url"`
	Url string `json:"url"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Created_at string `json:"created_at"`
	Zipball_url string `json:"zipball_url"`
	Id int `json:"id"`
	Body string `json:"body,omitempty"`
	Node_id string `json:"node_id"`
	Name string `json:"name"`
	Published_at string `json:"published_at"`
	Target_commitish string `json:"target_commitish"` // Specifies the commitish value that determines where the Git tag is created from.
	Body_html string `json:"body_html,omitempty"`
	Reactions GeneratedType `json:"reactions,omitempty"`
}

// Package represents the Package schema from the OpenAPI specification
type Package struct {
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Owner GeneratedType `json:"owner,omitempty"` // A GitHub user.
	Package_type string `json:"package_type"`
	Repository GeneratedType `json:"repository,omitempty"` // Minimal Repository
	Version_count int `json:"version_count"` // The number of versions of the package.
	Visibility string `json:"visibility"`
	Id int `json:"id"` // Unique identifier of the package.
	Name string `json:"name"` // The name of the package.
}

// Manifest represents the Manifest schema from the OpenAPI specification
type Manifest struct {
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
	Name string `json:"name"` // The name of the manifest.
	Resolved map[string]interface{} `json:"resolved,omitempty"` // A collection of resolved package dependencies.
	File map[string]interface{} `json:"file,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ref string `json:"ref,omitempty"` // The ref (branch or tag) of the workflow file to use
	Repository_id int `json:"repository_id"` // The ID of the repository where the workflow is defined
	Sha string `json:"sha,omitempty"` // The commit SHA of the workflow file to use
	Path string `json:"path"` // The path to the workflow file
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Milestone Milestone `json:"milestone,omitempty"` // A collection of related issues and pull requests.
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request Webhookspullrequest5 `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Effective_date string `json:"effective_date"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Previous_marketplace_purchase Webhookspreviousmarketplacepurchase `json:"previous_marketplace_purchase,omitempty"`
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Rule Webhooksrule `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"` // If the action was `edited`, the changes to the rule.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert map[string]interface{} `json:"alert"` // The security alert of the vulnerable dependency.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Browser_download_url string `json:"browser_download_url"`
	Created_at string `json:"created_at"`
	Label string `json:"label"`
	Uploader GeneratedType `json:"uploader"` // A GitHub user.
	Content_type string `json:"content_type"`
	Digest string `json:"digest"`
	Name string `json:"name"` // The file name of the asset.
	State string `json:"state"` // State of the release asset.
	Size int `json:"size"`
	Updated_at string `json:"updated_at"`
	Download_count int `json:"download_count"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Body string `json:"body"`
	Html_url string `json:"html_url"`
	Path string `json:"path"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Line int `json:"line"`
	Node_id string `json:"node_id"`
	Position int `json:"position"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comments int `json:"comments"`
	Deletions int `json:"deletions"`
	Mergeable_state string `json:"mergeable_state"`
	Id int64 `json:"id"`
	Additions int `json:"additions"`
	Base map[string]interface{} `json:"base"`
	Closed_at string `json:"closed_at"`
	Diff_url string `json:"diff_url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	Requested_reviewers []GeneratedType `json:"requested_reviewers,omitempty"`
	Requested_teams []GeneratedType `json:"requested_teams,omitempty"`
	Patch_url string `json:"patch_url"`
	Auto_merge GeneratedType `json:"auto_merge"` // The status of auto merging a pull request.
	Head map[string]interface{} `json:"head"`
	Html_url string `json:"html_url"`
	Labels []map[string]interface{} `json:"labels"`
	Number int `json:"number"` // Number uniquely identifying the pull request within its repository.
	Review_comment_url string `json:"review_comment_url"`
	Title string `json:"title"` // The title of the pull request.
	Issue_url string `json:"issue_url"`
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Mergeable bool `json:"mergeable"`
	Url string `json:"url"`
	Merged_by GeneratedType `json:"merged_by"` // A GitHub user.
	Maintainer_can_modify bool `json:"maintainer_can_modify"` // Indicates whether maintainers can modify the pull request.
	State string `json:"state"` // State of this Pull Request. Either `open` or `closed`.
	Comments_url string `json:"comments_url"`
	Commits_url string `json:"commits_url"`
	Links map[string]interface{} `json:"_links"`
	Review_comments_url string `json:"review_comments_url"`
	Updated_at string `json:"updated_at"`
	Changed_files int `json:"changed_files"`
	Node_id string `json:"node_id"`
	Body string `json:"body"`
	Locked bool `json:"locked"`
	Merged_at string `json:"merged_at"`
	Review_comments int `json:"review_comments"`
	Statuses_url string `json:"statuses_url"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Commits int `json:"commits"`
	Created_at string `json:"created_at"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Rebaseable bool `json:"rebaseable,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Merged bool `json:"merged"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged.
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., "Merge pull request #123 from branch-name").
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.**
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow auto-merge for pull requests.
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether to allow updating the pull request's branch.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Properties []GeneratedType `json:"properties"` // List of custom property names and associated values
	Repository_full_name string `json:"repository_full_name"`
	Repository_id int `json:"repository_id"`
	Repository_name string `json:"repository_name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Apps_url string `json:"apps_url"`
	Teams []map[string]interface{} `json:"teams"`
	Teams_url string `json:"teams_url"`
	Url string `json:"url"`
	Users []map[string]interface{} `json:"users"`
	Users_url string `json:"users_url"`
	Apps []map[string]interface{} `json:"apps"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alert Webhooksalert `json:"alert"` // The security alert of the vulnerable dependency.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Token_last_used_at string `json:"token_last_used_at"` // Date and time when the associated fine-grained personal access token was last used for authentication.
	Token_name string `json:"token_name"` // The name given to the user's token. This field can also be found in an organization's settings page for Active Tokens.
	Repository_selection string `json:"repository_selection"` // Type of repository selection requested.
	Token_expires_at string `json:"token_expires_at"` // Date and time when the associated fine-grained personal access token expires.
	Repositories_url string `json:"repositories_url"` // URL to the list of repositories the fine-grained personal access token can access. Only follow when `repository_selection` is `subset`.
	Token_expired bool `json:"token_expired"` // Whether the associated fine-grained personal access token has expired.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Access_granted_at string `json:"access_granted_at"` // Date and time when the fine-grained personal access token was approved to access the organization.
	Id int `json:"id"` // Unique identifier of the fine-grained personal access token grant. The `pat_id` used to get details about an approved fine-grained personal access token.
	Permissions map[string]interface{} `json:"permissions"` // Permissions requested, categorized by type of permission.
	Token_id int `json:"token_id"` // Unique identifier of the user's token. This field can also be found in audit log events and the organization's settings for their PAT grants.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Database_commit_sha string `json:"database_commit_sha,omitempty"` // The SHA of the commit the CodeQL database was built against. This is only available for successful analyses.
	Failure_message string `json:"failure_message,omitempty"` // The reason of the failure of this repo task. This is only available if the repository task has failed.
	Repository GeneratedType `json:"repository"` // A GitHub repository.
	Result_count int `json:"result_count,omitempty"` // The number of results in the case of a successful analysis. This is only available for successful analyses.
	Source_location_prefix string `json:"source_location_prefix,omitempty"` // The source location prefix to use. This is only available for successful analyses.
	Analysis_status string `json:"analysis_status"` // The new status of the CodeQL variant analysis repository task.
	Artifact_size_in_bytes int `json:"artifact_size_in_bytes,omitempty"` // The size of the artifact. This is only available for successful analyses.
	Artifact_url string `json:"artifact_url,omitempty"` // The URL of the artifact. This is only available for successful analyses.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the ping JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Github_owned_allowed bool `json:"github_owned_allowed,omitempty"` // Whether GitHub-owned actions are allowed. For example, this includes the actions in the `actions` organization.
	Patterns_allowed []string `json:"patterns_allowed,omitempty"` // Specifies a list of string-matching patterns to allow specific action(s) and reusable workflow(s). Wildcards, tags, and SHAs are allowed. For example, `monalisa/octocat@*`, `monalisa/octocat@v2`, `monalisa/*`. > [!NOTE] > The `patterns_allowed` setting only applies to public repositories.
	Verified_allowed bool `json:"verified_allowed,omitempty"` // Whether actions from GitHub Marketplace verified creators are allowed. Set to `true` to allow all actions by GitHub Marketplace verified creators.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Links map[string]interface{} `json:"_links"`
	Path string `json:"path"`
	Sha string `json:"sha"`
	TypeField string `json:"type"`
	Html_url string `json:"html_url"`
	Url string `json:"url"`
	Download_url string `json:"download_url"`
	Git_url string `json:"git_url"`
	Size int `json:"size"`
	Target string `json:"target"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Sub_issue_repo Repository `json:"sub_issue_repo"` // A repository on GitHub.
	Action string `json:"action"`
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Email string `json:"email,omitempty"`
	Id int `json:"id"`
	Login string `json:"login"`
	Marketplace_pending_change map[string]interface{} `json:"marketplace_pending_change,omitempty"`
	Marketplace_purchase map[string]interface{} `json:"marketplace_purchase"`
	Organization_billing_email string `json:"organization_billing_email,omitempty"`
	TypeField string `json:"type"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Commit_url string `json:"commit_url"`
	Event string `json:"event"`
	Lock_reason string `json:"lock_reason"`
	Node_id string `json:"node_id"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Created_at string `json:"created_at"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Released string `json:"released"`
	Created_at string `json:"created_at"`
	Name string `json:"name"`
	Description string `json:"description"`
	Logo_url string `json:"logo_url,omitempty"`
	Score float64 `json:"score"`
	Featured bool `json:"featured"`
	Curated bool `json:"curated"`
	Updated_at string `json:"updated_at"`
	Related []map[string]interface{} `json:"related,omitempty"`
	Aliases []map[string]interface{} `json:"aliases,omitempty"`
	Created_by string `json:"created_by"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Short_description string `json:"short_description"`
	Repository_count int `json:"repository_count,omitempty"`
	Display_name string `json:"display_name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Effective_date string `json:"effective_date,omitempty"` // The `pending_cancellation` and `pending_tier_change` event types will include the date the cancellation or tier change will take effect.
	Changes Webhookschanges8 `json:"changes"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// Webhooksprojectcolumn represents the Webhooksprojectcolumn schema from the OpenAPI specification
type Webhooksprojectcolumn struct {
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	After_id int `json:"after_id,omitempty"`
	Cards_url string `json:"cards_url"`
	Id int `json:"id"` // The unique identifier of the project column
	Node_id string `json:"node_id"`
	Name string `json:"name"` // Name of the project column
	Created_at string `json:"created_at"`
	Project_url string `json:"project_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Can_approve_pull_request_reviews bool `json:"can_approve_pull_request_reviews"` // Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
	Default_workflow_permissions string `json:"default_workflow_permissions"` // The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"` // A unique identifier of the repository.
	Name string `json:"name"` // The name of the repository.
	Private bool `json:"private"` // Whether the repository is private.
	Stargazers_count int `json:"stargazers_count"`
	Updated_at string `json:"updated_at"`
	Full_name string `json:"full_name"` // The full, globally unique, name of the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Permissions string `json:"permissions"` // The permission associated with the invitation.
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Url string `json:"url"` // URL for the repository invitation
	Created_at string `json:"created_at"`
	Expired bool `json:"expired,omitempty"` // Whether or not the invitation has expired
	Id int64 `json:"id"` // Unique identifier of the repository invitation.
	Inviter GeneratedType `json:"inviter"` // A GitHub user.
	Html_url string `json:"html_url"`
	Invitee GeneratedType `json:"invitee"` // A GitHub user.
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card Webhooksprojectcard `json:"project_card"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Definition GeneratedType `json:"definition"` // Custom property defined on an organization
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Merge_type string `json:"merge_type,omitempty"`
	Message string `json:"message,omitempty"`
	Base_branch string `json:"base_branch,omitempty"`
}

// Webhooksmilestone3 represents the Webhooksmilestone3 schema from the OpenAPI specification
type Webhooksmilestone3 struct {
	Title string `json:"title"` // The title of the milestone.
	Id int `json:"id"`
	Created_at string `json:"created_at"`
	Closed_issues int `json:"closed_issues"`
	Labels_url string `json:"labels_url"`
	Open_issues int `json:"open_issues"`
	Number int `json:"number"` // The number of the milestone.
	Updated_at string `json:"updated_at"`
	Node_id string `json:"node_id"`
	Html_url string `json:"html_url"`
	Description string `json:"description"`
	State string `json:"state"` // The state of the milestone.
	Due_on string `json:"due_on"`
	Closed_at string `json:"closed_at"`
	Creator map[string]interface{} `json:"creator"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) the comment belongs to.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhooksissuecomment `json:"comment"` // The [comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment) itself.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Rate GeneratedType `json:"rate"`
	Resources map[string]interface{} `json:"resources"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Event string `json:"event,omitempty"` // The event that triggered the deployment protection rule.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_requests []GeneratedType `json:"pull_requests,omitempty"`
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action,omitempty"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Deployment_callback_url string `json:"deployment_callback_url,omitempty"` // The URL to review the deployment protection rule.
	Environment string `json:"environment,omitempty"` // The name of the environment that has the deployment protection rule.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// Language represents the Language schema from the OpenAPI specification
type Language struct {
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Workflow_job interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Subscribers_url string `json:"subscribers_url"`
	Languages_url string `json:"languages_url"`
	Releases_url string `json:"releases_url"`
	Has_wiki bool `json:"has_wiki"` // Whether the wiki is enabled.
	Forks_url string `json:"forks_url"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Issue_events_url string `json:"issue_events_url"`
	Use_squash_pr_title_as_default bool `json:"use_squash_pr_title_as_default,omitempty"` // Whether a squash merge commit can use the pull request title as default. **This property is closing down. Please use `squash_merge_commit_title` instead.
	Notifications_url string `json:"notifications_url"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"` // Whether to delete head branches when pull requests are merged
	Mirror_url string `json:"mirror_url"`
	Has_projects bool `json:"has_projects"` // Whether projects are enabled.
	Size int `json:"size"` // The size of the repository, in kilobytes. Size is calculated hourly. When a repository is initially created, the size is 0.
	Merge_commit_message string `json:"merge_commit_message,omitempty"` // The default value for a merge commit message. - `PR_TITLE` - default to the pull request's title. - `PR_BODY` - default to the pull request's body. - `BLANK` - default to a blank commit message.
	Forks int `json:"forks"`
	License GeneratedType `json:"license"` // License Simple
	Is_template bool `json:"is_template,omitempty"` // Whether this repository acts as a template that can be used to generate new repositories.
	Labels_url string `json:"labels_url"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"` // Whether to allow rebase merges for pull requests.
	Pulls_url string `json:"pulls_url"`
	Updated_at string `json:"updated_at"`
	Assignees_url string `json:"assignees_url"`
	Comments_url string `json:"comments_url"`
	Stargazers_url string `json:"stargazers_url"`
	Clone_url string `json:"clone_url"`
	Archive_url string `json:"archive_url"`
	Has_pages bool `json:"has_pages"`
	Has_downloads bool `json:"has_downloads"` // Whether downloads are enabled.
	Keys_url string `json:"keys_url"`
	Name string `json:"name"` // The name of the repository.
	Allow_update_branch bool `json:"allow_update_branch,omitempty"` // Whether or not a pull request head branch that is behind its base branch can always be updated even if it is not required to be up to date before merging.
	Archived bool `json:"archived"` // Whether the repository is archived.
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"` // Whether to allow merge commits for pull requests.
	Has_discussions bool `json:"has_discussions,omitempty"` // Whether discussions are enabled.
	Contributors_url string `json:"contributors_url"`
	Events_url string `json:"events_url"`
	Private bool `json:"private"` // Whether the repository is private or public.
	Svn_url string `json:"svn_url"`
	Git_url string `json:"git_url"`
	Html_url string `json:"html_url"`
	Squash_merge_commit_title string `json:"squash_merge_commit_title,omitempty"` // The default value for a squash merge commit title: - `PR_TITLE` - default to the pull request's title. - `COMMIT_OR_PR_TITLE` - default to the commit's title (if only one commit) or the pull request's title (when more than one commit).
	Tags_url string `json:"tags_url"`
	Created_at string `json:"created_at"`
	Blobs_url string `json:"blobs_url"`
	Git_commits_url string `json:"git_commits_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"` // Whether to allow Auto-merge to be used on pull requests.
	Squash_merge_commit_message string `json:"squash_merge_commit_message,omitempty"` // The default value for a squash merge commit message: - `PR_BODY` - default to the pull request's body. - `COMMIT_MESSAGES` - default to the branch's commit messages. - `BLANK` - default to a blank commit message.
	Branches_url string `json:"branches_url"`
	Forks_count int `json:"forks_count"`
	Milestones_url string `json:"milestones_url"`
	Ssh_url string `json:"ssh_url"`
	Downloads_url string `json:"downloads_url"`
	Subscription_url string `json:"subscription_url"`
	Description string `json:"description"`
	Watchers_count int `json:"watchers_count"`
	Node_id string `json:"node_id"`
	Allow_forking bool `json:"allow_forking,omitempty"` // Whether to allow forking this repo
	Contents_url string `json:"contents_url"`
	Trees_url string `json:"trees_url"`
	Stargazers_count int `json:"stargazers_count"`
	Homepage string `json:"homepage"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Issues_url string `json:"issues_url"`
	Merges_url string `json:"merges_url"`
	Deployments_url string `json:"deployments_url"`
	Git_refs_url string `json:"git_refs_url"`
	Hooks_url string `json:"hooks_url"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub user.
	Watchers int `json:"watchers"`
	Anonymous_access_enabled bool `json:"anonymous_access_enabled,omitempty"` // Whether anonymous git access is enabled for this repository
	Custom_properties map[string]interface{} `json:"custom_properties,omitempty"` // The custom properties that were defined for the repository. The keys are the custom property names, and the values are the corresponding custom property values.
	Open_issues_count int `json:"open_issues_count"`
	Topics []string `json:"topics,omitempty"`
	Default_branch string `json:"default_branch"` // The default branch of the repository.
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"` // Whether to require contributors to sign off on web-based commits
	Open_issues int `json:"open_issues"`
	Git_tags_url string `json:"git_tags_url"`
	Url string `json:"url"`
	Fork bool `json:"fork"`
	Teams_url string `json:"teams_url"`
	Commits_url string `json:"commits_url"`
	Language string `json:"language"`
	Statuses_url string `json:"statuses_url"`
	Subscribers_count int `json:"subscribers_count,omitempty"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Full_name string `json:"full_name"`
	Template_repository map[string]interface{} `json:"template_repository,omitempty"`
	Collaborators_url string `json:"collaborators_url"`
	Master_branch string `json:"master_branch,omitempty"`
	Id int64 `json:"id"` // Unique identifier of the repository
	Merge_commit_title string `json:"merge_commit_title,omitempty"` // The default value for a merge commit title. - `PR_TITLE` - default to the pull request's title. - `MERGE_MESSAGE` - default to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Starred_at string `json:"starred_at,omitempty"`
	Compare_url string `json:"compare_url"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Has_issues bool `json:"has_issues"` // Whether issues are enabled.
	Issue_comment_url string `json:"issue_comment_url"`
	Network_count int `json:"network_count,omitempty"`
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"` // Whether to allow squash merges for pull requests.
	Pushed_at string `json:"pushed_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"`
	Review_requester GeneratedType `json:"review_requester"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Requested_team Team `json:"requested_team,omitempty"` // Groups of organization members that gives permissions on specified repositories.
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Requested_reviewer GeneratedType `json:"requested_reviewer,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"` // The REST API URL of the alert resource.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Number int `json:"number"` // The security alert number.
	State string `json:"state"` // State of a code scanning alert.
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Rule GeneratedType `json:"rule"`
	Dismissal_approved_by GeneratedType `json:"dismissal_approved_by,omitempty"` // A GitHub user.
	Most_recent_instance GeneratedType `json:"most_recent_instance"`
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Tool GeneratedType `json:"tool"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State string `json:"state,omitempty"` // Code scanning default setup has been configured or not.
	Threat_model string `json:"threat_model,omitempty"` // Threat model to be used for code scanning analysis. Use `remote` to analyze only network sources and `remote_and_local` to include local sources like filesystem access, command-line arguments, database reads, environment variable and standard input.
	Updated_at string `json:"updated_at,omitempty"` // Timestamp of latest configuration update.
	Languages []string `json:"languages,omitempty"` // Languages to be analyzed.
	Query_suite string `json:"query_suite,omitempty"` // CodeQL query suite to be used.
	Runner_label string `json:"runner_label,omitempty"` // Runner label to be used if the runner type is labeled.
	Runner_type string `json:"runner_type,omitempty"` // Runner type to be used.
	Schedule string `json:"schedule,omitempty"` // The frequency of the periodic analysis.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body_html string `json:"body_html,omitempty"`
	Event string `json:"event"`
	Submitted_at string `json:"submitted_at,omitempty"`
	Body string `json:"body"` // The text of the review.
	State string `json:"state"`
	User GeneratedType `json:"user"` // A GitHub user.
	Links map[string]interface{} `json:"_links"`
	Body_text string `json:"body_text,omitempty"`
	Id int `json:"id"` // Unique identifier of the review
	Pull_request_url string `json:"pull_request_url"`
	Commit_id string `json:"commit_id"` // A commit SHA for the review.
}

// Collaborator represents the Collaborator schema from the OpenAPI specification
type Collaborator struct {
	Site_admin bool `json:"site_admin"`
	Login string `json:"login"`
	Subscriptions_url string `json:"subscriptions_url"`
	TypeField string `json:"type"`
	User_view_type string `json:"user_view_type,omitempty"`
	Id int64 `json:"id"`
	Events_url string `json:"events_url"`
	Followers_url string `json:"followers_url"`
	Repos_url string `json:"repos_url"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Gists_url string `json:"gists_url"`
	Email string `json:"email,omitempty"`
	Name string `json:"name,omitempty"`
	Organizations_url string `json:"organizations_url"`
	Role_name string `json:"role_name"`
	Starred_url string `json:"starred_url"`
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Received_events_url string `json:"received_events_url"`
	Following_url string `json:"following_url"`
	Gravatar_id string `json:"gravatar_id"`
}

// Event represents the Event schema from the OpenAPI specification
type Event struct {
	Payload map[string]interface{} `json:"payload"`
	Public bool `json:"public"`
	Repo map[string]interface{} `json:"repo"`
	TypeField string `json:"type"`
	Actor Actor `json:"actor"` // Actor
	Created_at string `json:"created_at"`
	Id string `json:"id"`
	Org Actor `json:"org,omitempty"` // Actor
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Branch string `json:"branch"`
	Path string `json:"path"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request GeneratedType `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name of the role.
	Organization GeneratedType `json:"organization"` // A GitHub user.
	Permissions []string `json:"permissions"` // A list of permissions included in this role.
	Id int64 `json:"id"` // The unique identifier of the role.
	Source string `json:"source,omitempty"` // Source answers the question, "where did this role come from?"
	Updated_at string `json:"updated_at"` // The date and time the role was last updated.
	Description string `json:"description,omitempty"` // A short description about who this role is for or what permissions it grants.
	Base_role string `json:"base_role,omitempty"` // The system role from which this role inherits permissions.
	Created_at string `json:"created_at"` // The date and time the role was created.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	PackageField map[string]interface{} `json:"package"` // Information about the package.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Requester interface{} `json:"requester,omitempty"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repositories []map[string]interface{} `json:"repositories,omitempty"` // An array of repository objects that the installation can access.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Spdx_id string `json:"spdx_id"`
	Url string `json:"url"`
	Html_url string `json:"html_url,omitempty"`
	Key string `json:"key"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sbom map[string]interface{} `json:"sbom"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Message string `json:"message,omitempty"`
	Schemas []string `json:"schemas,omitempty"`
	Scimtype string `json:"scimType,omitempty"`
	Status int `json:"status,omitempty"`
	Detail string `json:"detail,omitempty"`
	Documentation_url string `json:"documentation_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description"` // The description of an autofix.
	Started_at string `json:"started_at"` // The start time of an autofix in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Status string `json:"status"` // The status of an autofix.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Full_name string `json:"full_name"` // The full, globally unique name of the repository.
	Html_url string `json:"html_url"` // The URL to view the repository on GitHub.com.
	Id int `json:"id"` // A unique identifier of the repository.
	Node_id string `json:"node_id"` // The GraphQL identifier of the repository.
	Private bool `json:"private"` // Whether the repository is private.
	Default_branch string `json:"default_branch"` // The default branch for the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes Webhooksprojectchanges `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Runner_label string `json:"runner_label,omitempty"` // Runner label to be used if the runner type is labeled.
	Runner_type string `json:"runner_type,omitempty"` // Runner type to be used.
	State string `json:"state,omitempty"` // The desired state of code scanning default setup.
	Threat_model string `json:"threat_model,omitempty"` // Threat model to be used for code scanning analysis. Use `remote` to analyze only network sources and `remote_and_local` to include local sources like filesystem access, command-line arguments, database reads, environment variable and standard input.
	Languages []string `json:"languages,omitempty"` // CodeQL languages to be analyzed.
	Query_suite string `json:"query_suite,omitempty"` // CodeQL query suite to be used.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Sub_issue_repo Repository `json:"sub_issue_repo"` // A repository on GitHub.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// Project represents the Project schema from the OpenAPI specification
type Project struct {
	State string `json:"state"` // State of the project; either 'open' or 'closed'
	Body string `json:"body"` // Body of the project
	Private bool `json:"private,omitempty"` // Whether or not this project can be seen by everyone. Only present if owner is an organization.
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Organization_permission string `json:"organization_permission,omitempty"` // The baseline permission that all organization members have on this project. Only present if owner is an organization.
	Html_url string `json:"html_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Name string `json:"name"` // Name of the project
	Number int `json:"number"`
	Columns_url string `json:"columns_url"`
	Node_id string `json:"node_id"`
	Owner_url string `json:"owner_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Context string `json:"context"`
	Id int `json:"id"`
	State string `json:"state"`
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Description string `json:"description"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Required bool `json:"required,omitempty"`
	Target_url string `json:"target_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Display_name string `json:"display_name"` // The display name of the machine includes cores, memory, and storage.
	Memory_in_bytes int `json:"memory_in_bytes"` // How much memory is available to the codespace.
	Name string `json:"name"` // The name of the machine.
	Operating_system string `json:"operating_system"` // The operating system of the machine.
	Prebuild_availability string `json:"prebuild_availability"` // Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be "null" if prebuilds are not supported or prebuild availability could not be determined. Value will be "none" if no prebuild is available. Latest values "ready" and "in_progress" indicate the prebuild availability status.
	Storage_in_bytes int `json:"storage_in_bytes"` // How much storage is available to the codespace.
	Cpus int `json:"cpus"` // How many cores are available to the codespace.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Milestone Webhooksmilestone `json:"milestone,omitempty"` // A collection of related issues and pull requests.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Deleted_at string `json:"deleted_at,omitempty"`
	License string `json:"license,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Url string `json:"url"`
	Html_url string `json:"html_url,omitempty"`
	Id int `json:"id"` // Unique identifier of the package version.
	Description string `json:"description,omitempty"`
	Name string `json:"name"` // The name of the package version.
	Package_html_url string `json:"package_html_url"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Invitation_teams_url string `json:"invitation_teams_url"`
	Failed_at string `json:"failed_at,omitempty"`
	Node_id string `json:"node_id"`
	Team_count int `json:"team_count"`
	Failed_reason string `json:"failed_reason,omitempty"`
	Login string `json:"login"`
	Role string `json:"role"`
	Created_at string `json:"created_at"`
	Email string `json:"email"`
	Invitation_source string `json:"invitation_source,omitempty"`
	Inviter GeneratedType `json:"inviter"` // A GitHub user.
	Id int64 `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Personal_access_token_request GeneratedType `json:"personal_access_token_request"` // Details of a Personal Access Token Request.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// Root represents the Root schema from the OpenAPI specification
type Root struct {
	Issue_search_url string `json:"issue_search_url"`
	Current_user_repositories_url string `json:"current_user_repositories_url"`
	User_repositories_url string `json:"user_repositories_url"`
	Hub_url string `json:"hub_url,omitempty"`
	Public_gists_url string `json:"public_gists_url"`
	Events_url string `json:"events_url"`
	Code_search_url string `json:"code_search_url"`
	Organization_repositories_url string `json:"organization_repositories_url"`
	User_url string `json:"user_url"`
	Emails_url string `json:"emails_url"`
	Feeds_url string `json:"feeds_url"`
	User_search_url string `json:"user_search_url"`
	Starred_url string `json:"starred_url"`
	Organization_url string `json:"organization_url"`
	User_organizations_url string `json:"user_organizations_url"`
	Followers_url string `json:"followers_url"`
	Topic_search_url string `json:"topic_search_url,omitempty"`
	Keys_url string `json:"keys_url"`
	Current_user_url string `json:"current_user_url"`
	Repository_url string `json:"repository_url"`
	Gists_url string `json:"gists_url"`
	Issues_url string `json:"issues_url"`
	Commit_search_url string `json:"commit_search_url"`
	Starred_gists_url string `json:"starred_gists_url"`
	Current_user_authorizations_html_url string `json:"current_user_authorizations_html_url"`
	Label_search_url string `json:"label_search_url"`
	Notifications_url string `json:"notifications_url"`
	Repository_search_url string `json:"repository_search_url"`
	Emojis_url string `json:"emojis_url"`
	Following_url string `json:"following_url"`
	Organization_teams_url string `json:"organization_teams_url"`
	Rate_limit_url string `json:"rate_limit_url"`
	Authorizations_url string `json:"authorizations_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Source string `json:"source"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Issue_comment_url string `json:"issue_comment_url"` // The API URL to get the issue comment where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Comment Webhookscomment `json:"comment"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Membership Webhooksmembership `json:"membership,omitempty"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Checks []map[string]interface{} `json:"checks"`
	Contexts []string `json:"contexts"`
	Contexts_url string `json:"contexts_url,omitempty"`
	Enforcement_level string `json:"enforcement_level,omitempty"`
	Strict bool `json:"strict,omitempty"`
	Url string `json:"url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repositories_removed []map[string]interface{} `json:"repositories_removed"` // An array of repository objects, which were removed from the installation.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Repository_selection string `json:"repository_selection"` // Describe whether all repositories have been selected or there's a selection involved
	Installation Installation `json:"installation"` // Installation
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Requester Webhooksuser `json:"requester"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Repositories_added []map[string]interface{} `json:"repositories_added"` // An array of repository objects, which were added to the installation.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Website_url string `json:"website_url,omitempty"` // The enterprise's website URL.
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the enterprise
	Avatar_url string `json:"avatar_url"`
	Name string `json:"name"` // The name of the enterprise.
	Created_at string `json:"created_at"`
	Description string `json:"description,omitempty"` // A short description of the enterprise.
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Slug string `json:"slug"` // The slug url identifier for the enterprise.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion_comment_url string `json:"discussion_comment_url"` // The API URL to get the discussion comment where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Status string `json:"status"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	PackageField map[string]interface{} `json:"package"` // Information about the package.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion_body_url string `json:"discussion_body_url"` // The URL to the discussion where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Visibility string `json:"visibility"` // Visibility of a secret
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the secret.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Review_id int `json:"review_id"`
	State string `json:"state"`
	Dismissal_commit_id string `json:"dismissal_commit_id,omitempty"`
	Dismissal_message string `json:"dismissal_message"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Statuses_url string `json:"statuses_url"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Merged_at string `json:"merged_at"`
	Labels []map[string]interface{} `json:"labels"`
	User GeneratedType `json:"user"` // A GitHub user.
	Body string `json:"body"`
	Html_url string `json:"html_url"`
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	Review_comments_url string `json:"review_comments_url"`
	Auto_merge GeneratedType `json:"auto_merge"` // The status of auto merging a pull request.
	Closed_at string `json:"closed_at"`
	Created_at string `json:"created_at"`
	State string `json:"state"`
	Url string `json:"url"`
	Diff_url string `json:"diff_url"`
	Requested_teams []Team `json:"requested_teams,omitempty"`
	Comments_url string `json:"comments_url"`
	Review_comment_url string `json:"review_comment_url"`
	Number int `json:"number"`
	Locked bool `json:"locked"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Issue_url string `json:"issue_url"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Id int64 `json:"id"`
	Title string `json:"title"`
	Node_id string `json:"node_id"`
	Patch_url string `json:"patch_url"`
	Links map[string]interface{} `json:"_links"`
	Base map[string]interface{} `json:"base"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Head map[string]interface{} `json:"head"`
	Requested_reviewers []GeneratedType `json:"requested_reviewers,omitempty"`
	Commits_url string `json:"commits_url"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Merge_base_commit Commit `json:"merge_base_commit"` // Commit
	Status string `json:"status"`
	Total_commits int `json:"total_commits"`
	Url string `json:"url"`
	Behind_by int `json:"behind_by"`
	Diff_url string `json:"diff_url"`
	Files []GeneratedType `json:"files,omitempty"`
	Html_url string `json:"html_url"`
	Patch_url string `json:"patch_url"`
	Base_commit Commit `json:"base_commit"` // Commit
	Commits []Commit `json:"commits"`
	Permalink_url string `json:"permalink_url"`
	Ahead_by int `json:"ahead_by"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Rate_limited_request_count int64 `json:"rate_limited_request_count,omitempty"` // The total number of requests that were rate limited within the queried time period
	Total_request_count int64 `json:"total_request_count,omitempty"` // The total number of requests within the queried time period
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes Webhookschanges8 `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Sponsorship Webhookssponsorship `json:"sponsorship"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job map[string]interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Hook map[string]interface{} `json:"hook"` // The deleted webhook. This will contain different keys based on the type of webhook it is: repository, organization, business, app, or GitHub Marketplace.
	Hook_id int `json:"hook_id"` // The id of the modified webhook.
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Protected bool `json:"protected"`
	Protection GeneratedType `json:"protection,omitempty"` // Branch Protection
	Protection_url string `json:"protection_url,omitempty"`
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha256_checksum string `json:"sha256_checksum,omitempty"`
	Temp_download_token string `json:"temp_download_token,omitempty"` // A short lived bearer token used to download the runner, if needed.
	Architecture string `json:"architecture"`
	Download_url string `json:"download_url"`
	Filename string `json:"filename"`
	Os string `json:"os"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Configuration GeneratedType `json:"configuration,omitempty"` // A code security configuration
	Status string `json:"status,omitempty"` // The attachment status of the code security configuration on the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Documentation_url string `json:"documentation_url"`
	Errors []string `json:"errors,omitempty"`
	Message string `json:"message"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Supported secret scanning patterns](https://docs.github.com/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
	Validity string `json:"validity,omitempty"` // The token status as of the latest validity check.
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
	Resolution_comment string `json:"resolution_comment,omitempty"` // An optional comment to resolve an alert.
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	Push_protection_bypassed_by GeneratedType `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypass_request_html_url string `json:"push_protection_bypass_request_html_url,omitempty"` // The URL to a push protection bypass request.
	Number int `json:"number,omitempty"` // The security alert number.
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Multi_repo bool `json:"multi_repo,omitempty"` // Whether the detected secret was found in multiple repositories in the same organization or business.
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
	Publicly_leaked bool `json:"publicly_leaked,omitempty"` // Whether the detected secret was publicly leaked.
	Resolved_by GeneratedType `json:"resolved_by,omitempty"` // A GitHub user.
	Push_protection_bypass_request_comment string `json:"push_protection_bypass_request_comment,omitempty"` // An optional comment when requesting a push protection bypass.
	Push_protection_bypass_request_reviewer GeneratedType `json:"push_protection_bypass_request_reviewer,omitempty"` // A GitHub user.
	Resolution string `json:"resolution,omitempty"` // The reason for resolving the alert.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Push_protection_bypass_request_reviewer_comment string `json:"push_protection_bypass_request_reviewer_comment,omitempty"` // An optional comment when reviewing a push protection bypass.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Cpus int `json:"cpus"` // How many cores are available to the codespace.
	Display_name string `json:"display_name"` // The display name of the machine includes cores, memory, and storage.
	Memory_in_bytes int `json:"memory_in_bytes"` // How much memory is available to the codespace.
	Name string `json:"name"` // The name of the machine.
	Operating_system string `json:"operating_system"` // The operating system of the machine.
	Prebuild_availability string `json:"prebuild_availability"` // Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be "null" if prebuilds are not supported or prebuild availability could not be determined. Value will be "none" if no prebuild is available. Latest values "ready" and "in_progress" indicate the prebuild availability status.
	Storage_in_bytes int `json:"storage_in_bytes"` // How much storage is available to the codespace.
}

// Environment represents the Environment schema from the OpenAPI specification
type Environment struct {
	Updated_at string `json:"updated_at"` // The time that the environment was last updated, in ISO 8601 format.
	Url string `json:"url"`
	Id int64 `json:"id"` // The id of the environment.
	Node_id string `json:"node_id"`
	Created_at string `json:"created_at"` // The time that the environment was created, in ISO 8601 format.
	Html_url string `json:"html_url"`
	Deployment_branch_policy GeneratedType `json:"deployment_branch_policy,omitempty"` // The type of deployment branch policy for this environment. To allow all branches to deploy, set to `null`.
	Protection_rules []interface{} `json:"protection_rules,omitempty"` // Built-in deployment protection rules for the environment.
	Name string `json:"name"` // The name of the environment.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Size int `json:"size"`
	TypeField string `json:"type"`
	Content string `json:"content"`
	Download_url string `json:"download_url"`
	Url string `json:"url"`
	Encoding string `json:"encoding"`
	Html_url string `json:"html_url"`
	Links map[string]interface{} `json:"_links"`
	License GeneratedType `json:"license"` // License Simple
	Sha string `json:"sha"`
	Git_url string `json:"git_url"`
	Name string `json:"name"`
	Path string `json:"path"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField GeneratedType `json:"type,omitempty"` // The type of issue.
	State_reason string `json:"state_reason,omitempty"` // The reason for the current state
	Repository_url string `json:"repository_url"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Body string `json:"body,omitempty"` // Contents of the issue
	Sub_issues_summary GeneratedType `json:"sub_issues_summary,omitempty"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Comments int `json:"comments"`
	Title string `json:"title"` // Title of the issue
	Body_html string `json:"body_html,omitempty"`
	Url string `json:"url"` // URL for the issue
	Html_url string `json:"html_url"`
	Draft bool `json:"draft,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Events_url string `json:"events_url"`
	Created_at string `json:"created_at"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Closed_by GeneratedType `json:"closed_by,omitempty"` // A GitHub user.
	Number int `json:"number"` // Number uniquely identifying the issue within its repository
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Updated_at string `json:"updated_at"`
	Body_text string `json:"body_text,omitempty"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Id int64 `json:"id"`
	Node_id string `json:"node_id"`
	Comments_url string `json:"comments_url"`
	Locked bool `json:"locked"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Labels []interface{} `json:"labels"` // Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository
	Assignees []GeneratedType `json:"assignees,omitempty"`
	State string `json:"state"` // State of the issue; either 'open' or 'closed'
	Timeline_url string `json:"timeline_url,omitempty"`
	Labels_url string `json:"labels_url"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Closed_at string `json:"closed_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"` // The date and time at which the variable was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Value string `json:"value"` // The value of the variable.
	Visibility string `json:"visibility"` // Visibility of a variable
	Created_at string `json:"created_at"` // The date and time at which the variable was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the variable.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Billable map[string]interface{} `json:"billable"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
}

// Webhooksteam1 represents the Webhooksteam1 schema from the OpenAPI specification
type Webhooksteam1 struct {
	Html_url string `json:"html_url,omitempty"`
	Name string `json:"name"` // Name of the team
	Parent map[string]interface{} `json:"parent,omitempty"`
	Notification_setting string `json:"notification_setting,omitempty"` // Whether team members will receive notifications when their team is @mentioned
	Url string `json:"url,omitempty"` // URL for the team
	Id int `json:"id"` // Unique identifier of the team
	Members_url string `json:"members_url,omitempty"`
	Privacy string `json:"privacy,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Permission string `json:"permission,omitempty"` // Permission that the team will have for its repositories
	Repositories_url string `json:"repositories_url,omitempty"`
	Slug string `json:"slug,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Description string `json:"description,omitempty"` // Description of the team
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Id int `json:"id"`
	Milestone map[string]interface{} `json:"milestone"`
}

// Stargazer represents the Stargazer schema from the OpenAPI specification
type Stargazer struct {
	Starred_at string `json:"starred_at"`
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Public_gists int `json:"public_gists"`
	Updated_at string `json:"updated_at"`
	Starred_url string `json:"starred_url"`
	Bio string `json:"bio"`
	Login string `json:"login"`
	TypeField string `json:"type"`
	Plan map[string]interface{} `json:"plan,omitempty"`
	Id int64 `json:"id"`
	Collaborators int `json:"collaborators,omitempty"`
	Site_admin bool `json:"site_admin"`
	Gravatar_id string `json:"gravatar_id"`
	Url string `json:"url"`
	Disk_usage int `json:"disk_usage,omitempty"`
	Following int `json:"following"`
	User_view_type string `json:"user_view_type,omitempty"`
	Notification_email string `json:"notification_email,omitempty"`
	Events_url string `json:"events_url"`
	Following_url string `json:"following_url"`
	Owned_private_repos int `json:"owned_private_repos,omitempty"`
	Twitter_username string `json:"twitter_username,omitempty"`
	Hireable bool `json:"hireable"`
	Blog string `json:"blog"`
	Private_gists int `json:"private_gists,omitempty"`
	Public_repos int `json:"public_repos"`
	Followers int `json:"followers"`
	Location string `json:"location"`
	Total_private_repos int `json:"total_private_repos,omitempty"`
	Name string `json:"name"`
	Received_events_url string `json:"received_events_url"`
	Organizations_url string `json:"organizations_url"`
	Html_url string `json:"html_url"`
	Subscriptions_url string `json:"subscriptions_url"`
	Company string `json:"company"`
	Repos_url string `json:"repos_url"`
	Avatar_url string `json:"avatar_url"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Gists_url string `json:"gists_url"`
	Email string `json:"email"`
	Followers_url string `json:"followers_url"`
}

// Artifact represents the Artifact schema from the OpenAPI specification
type Artifact struct {
	Expired bool `json:"expired"` // Whether or not the artifact has expired.
	Expires_at string `json:"expires_at"`
	Node_id string `json:"node_id"`
	Workflow_run map[string]interface{} `json:"workflow_run,omitempty"`
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the artifact.
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Id int `json:"id"`
	Size_in_bytes int `json:"size_in_bytes"` // The size in bytes of the artifact.
	Archive_download_url string `json:"archive_download_url"`
	Digest string `json:"digest,omitempty"` // The SHA256 digest of the artifact. This field will only be populated on artifacts uploaded with upload-artifact v4 or newer. For older versions, this field will be null.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Expires_at string `json:"expires_at"`
	Has_multiple_single_files bool `json:"has_multiple_single_files,omitempty"`
	Permissions GeneratedType `json:"permissions,omitempty"` // The permissions granted to the user access token.
	Repositories []Repository `json:"repositories,omitempty"`
	Repository_selection string `json:"repository_selection,omitempty"`
	Single_file string `json:"single_file,omitempty"`
	Single_file_paths []string `json:"single_file_paths,omitempty"`
	Token string `json:"token"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Dependencies []map[string]interface{} `json:"dependencies,omitempty"`
	Version_info map[string]interface{} `json:"version_info,omitempty"`
	Commit_oid string `json:"commit_oid,omitempty"`
	Homepage string `json:"homepage,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Name string `json:"name,omitempty"`
	Platform string `json:"platform,omitempty"`
	Description string `json:"description,omitempty"`
	Readme string `json:"readme,omitempty"`
	Repo string `json:"repo,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	State string `json:"state"`
	Statuses []GeneratedType `json:"statuses"`
	Total_count int `json:"total_count"`
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Sha string `json:"sha"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Ignored bool `json:"ignored"` // Determines if all notifications should be blocked from this repository.
	Reason string `json:"reason"`
	Repository_url string `json:"repository_url"`
	Subscribed bool `json:"subscribed"` // Determines if notifications should be received from this repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Status string `json:"status"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Commit string `json:"commit"`
	Created_at string `json:"created_at"`
	Duration int `json:"duration"`
	ErrorField map[string]interface{} `json:"error"`
	Pusher GeneratedType `json:"pusher"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Comment Webhookscomment `json:"comment"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Changes map[string]interface{} `json:"changes"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Alert map[string]interface{} `json:"alert"` // The code scanning alert involved in the event.
	Commit_oid string `json:"commit_oid"` // The commit SHA of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Ref string `json:"ref"` // The Git reference of the code scanning alert. When the action is `reopened_by_user` or `closed_by_user`, the event was triggered by the `sender` and this value will be empty.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	App GeneratedType `json:"app"` // A GitHub App that is providing a custom deployment protection rule.
	Enabled bool `json:"enabled"` // Whether the deployment protection rule is enabled for the environment.
	Id int `json:"id"` // The unique identifier for the deployment protection rule.
	Node_id string `json:"node_id"` // The node ID for the deployment protection rule.
}

// Discussion represents the Discussion schema from the OpenAPI specification
type Discussion struct {
	Html_url string `json:"html_url"`
	Title string `json:"title"`
	Comments int `json:"comments"`
	Created_at string `json:"created_at"`
	User map[string]interface{} `json:"user"`
	Answer_chosen_by map[string]interface{} `json:"answer_chosen_by"`
	Answer_html_url string `json:"answer_html_url"`
	Category map[string]interface{} `json:"category"`
	Locked bool `json:"locked"`
	Labels []Label `json:"labels,omitempty"`
	Id int `json:"id"`
	Number int `json:"number"`
	Repository_url string `json:"repository_url"`
	Node_id string `json:"node_id"`
	Reactions map[string]interface{} `json:"reactions,omitempty"`
	State string `json:"state"` // The current state of the discussion. `converting` means that the discussion is being converted from an issue. `transferring` means that the discussion is being transferred from another repository.
	Answer_chosen_at string `json:"answer_chosen_at"`
	Active_lock_reason string `json:"active_lock_reason"`
	State_reason string `json:"state_reason"` // The reason for the current state
	Timeline_url string `json:"timeline_url,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Updated_at string `json:"updated_at"`
	Body string `json:"body"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Base map[string]interface{} `json:"base"`
	Closed_at string `json:"closed_at"`
	Diff_url string `json:"diff_url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Draft bool `json:"draft,omitempty"` // Indicates whether or not the pull request is a draft.
	Requested_reviewers []GeneratedType `json:"requested_reviewers,omitempty"`
	Requested_teams []GeneratedType `json:"requested_teams,omitempty"`
	Patch_url string `json:"patch_url"`
	Auto_merge GeneratedType `json:"auto_merge"` // The status of auto merging a pull request.
	Head map[string]interface{} `json:"head"`
	Html_url string `json:"html_url"`
	Labels []map[string]interface{} `json:"labels"`
	Number int `json:"number"` // Number uniquely identifying the pull request within its repository.
	Review_comment_url string `json:"review_comment_url"`
	Title string `json:"title"` // The title of the pull request.
	Issue_url string `json:"issue_url"`
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Mergeable bool `json:"mergeable"`
	Url string `json:"url"`
	Merged_by GeneratedType `json:"merged_by"` // A GitHub user.
	Maintainer_can_modify bool `json:"maintainer_can_modify"` // Indicates whether maintainers can modify the pull request.
	State string `json:"state"` // State of this Pull Request. Either `open` or `closed`.
	Comments_url string `json:"comments_url"`
	Commits_url string `json:"commits_url"`
	Links map[string]interface{} `json:"_links"`
	Review_comments_url string `json:"review_comments_url"`
	Updated_at string `json:"updated_at"`
	Changed_files int `json:"changed_files"`
	Node_id string `json:"node_id"`
	Body string `json:"body"`
	Locked bool `json:"locked"`
	Merged_at string `json:"merged_at"`
	Review_comments int `json:"review_comments"`
	Statuses_url string `json:"statuses_url"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Commits int `json:"commits"`
	Created_at string `json:"created_at"`
	Merge_commit_sha string `json:"merge_commit_sha"`
	Rebaseable bool `json:"rebaseable,omitempty"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Merged bool `json:"merged"`
	Comments int `json:"comments"`
	Deletions int `json:"deletions"`
	Mergeable_state string `json:"mergeable_state"`
	Id int64 `json:"id"`
	Additions int `json:"additions"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project_card interface{} `json:"project_card"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) the comment belongs to.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Changes Webhookschanges `json:"changes"` // The changes to the comment.
	Comment Webhooksissuecomment `json:"comment"` // The [comment](https://docs.github.com/rest/issues/comments#get-an-issue-comment) itself.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Pull_request map[string]interface{} `json:"pull_request"`
	Review Webhooksreview `json:"review"` // The review that was affected.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Changes map[string]interface{} `json:"changes"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Invite_link string `json:"invite_link"` // The link that a student can use to accept the assignment.
	Slug string `json:"slug"` // Sluggified name of the assignment.
	Students_are_repo_admins bool `json:"students_are_repo_admins"` // Whether students are admins on created repository when a student accepts the assignment.
	Invitations_enabled bool `json:"invitations_enabled"` // Whether the invitation link is enabled. Visiting an enabled invitation link will accept the assignment.
	Max_members int `json:"max_members"` // The maximum allowable members per team.
	Classroom Classroom `json:"classroom"` // A GitHub Classroom classroom
	Passing int `json:"passing"` // The number of students that have passed the assignment.
	Feedback_pull_requests_enabled bool `json:"feedback_pull_requests_enabled"` // Whether feedback pull request will be created when a student accepts the assignment.
	Language string `json:"language"` // The programming language used in the assignment.
	Editor string `json:"editor"` // The selected editor for the assignment.
	Submitted int `json:"submitted"` // The number of students that have submitted the assignment.
	Id int `json:"id"` // Unique identifier of the repository.
	Deadline string `json:"deadline"` // The time at which the assignment is due.
	Public_repo bool `json:"public_repo"` // Whether an accepted assignment creates a public repository.
	Title string `json:"title"` // Assignment title.
	TypeField string `json:"type"` // Whether it's a group assignment or individual assignment.
	Max_teams int `json:"max_teams"` // The maximum allowable teams for the assignment.
	Starter_code_repository GeneratedType `json:"starter_code_repository"` // A GitHub repository view for Classroom
	Accepted int `json:"accepted"` // The number of students that have accepted the assignment.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Errors []map[string]interface{} `json:"errors"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_active_users int `json:"total_active_users,omitempty"` // The total number of Copilot users with activity belonging to any Copilot feature, globally, for the given day. Includes passive activity such as receiving a code suggestion, as well as engagement activity such as accepting a code suggestion or prompting chat. Does not include authentication events. Is not limited to the individual features detailed on the endpoint.
	Total_engaged_users int `json:"total_engaged_users,omitempty"` // The total number of Copilot users who engaged with any Copilot feature, for the given day. Examples include but are not limited to accepting a code suggestion, prompting Copilot chat, or triggering a PR Summary. Does not include authentication events. Is not limited to the individual features detailed on the endpoint.
	Copilot_dotcom_chat GeneratedType `json:"copilot_dotcom_chat,omitempty"` // Usage metrics for Copilot Chat in GitHub.com
	Copilot_dotcom_pull_requests GeneratedType `json:"copilot_dotcom_pull_requests,omitempty"` // Usage metrics for Copilot for pull requests.
	Copilot_ide_chat GeneratedType `json:"copilot_ide_chat,omitempty"` // Usage metrics for Copilot Chat in the IDE.
	Copilot_ide_code_completions GeneratedType `json:"copilot_ide_code_completions,omitempty"` // Usage metrics for Copilot editor code completions in the IDE.
	Date string `json:"date"` // The date for which the usage metrics are aggregated, in `YYYY-MM-DD` format.
}

// Webhooksalert represents the Webhooksalert schema from the OpenAPI specification
type Webhooksalert struct {
	State string `json:"state"`
	Fix_reason string `json:"fix_reason,omitempty"`
	Ghsa_id string `json:"ghsa_id"`
	Affected_range string `json:"affected_range"`
	Created_at string `json:"created_at"`
	Fixed_in string `json:"fixed_in,omitempty"`
	Affected_package_name string `json:"affected_package_name"`
	Dismiss_reason string `json:"dismiss_reason,omitempty"`
	External_reference string `json:"external_reference"`
	Node_id string `json:"node_id"`
	Dismissed_at string `json:"dismissed_at,omitempty"`
	Severity string `json:"severity"`
	Dismisser map[string]interface{} `json:"dismisser,omitempty"`
	Number int `json:"number"`
	Fixed_at string `json:"fixed_at,omitempty"`
	Id int `json:"id"`
	External_identifier string `json:"external_identifier"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Hook map[string]interface{} `json:"hook,omitempty"` // The webhook that is being pinged
	Hook_id int `json:"hook_id,omitempty"` // The ID of the webhook that triggered the ping.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Zen string `json:"zen,omitempty"` // Random string of GitHub zen.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job map[string]interface{} `json:"workflow_job"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Wait_timer_started_at string `json:"wait_timer_started_at"` // The time that the wait timer began.
	Current_user_can_approve bool `json:"current_user_can_approve"` // Whether the currently authenticated user can approve the deployment
	Environment map[string]interface{} `json:"environment"`
	Reviewers []map[string]interface{} `json:"reviewers"` // The people or teams that may approve jobs that reference the environment. You can list up to six users or teams as reviewers. The reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	Wait_timer int `json:"wait_timer"` // The set duration of the wait timer
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Description string `json:"description,omitempty"`
	Group_name string `json:"group_name,omitempty"`
	Name string `json:"name"`
	Organization_selection_type string `json:"organization_selection_type,omitempty"`
	Html_url string `json:"html_url"`
	Id int64 `json:"id"`
	Members_url string `json:"members_url"`
	Slug string `json:"slug"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Group_id string `json:"group_id,omitempty"`
	Sync_to_organizations string `json:"sync_to_organizations,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enabled bool `json:"enabled,omitempty"` // Whether public IP is enabled.
	Length int `json:"length,omitempty"` // The length of the IP prefix.
	Prefix string `json:"prefix,omitempty"` // The prefix for the public IP.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Context string `json:"context"` // The status check context name that must be present on the commit.
	Integration_id int `json:"integration_id,omitempty"` // The optional integration ID that this status check must originate from.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Source string `json:"source,omitempty"` // The source of the repository property. Defaults to 'custom' if not specified.
	Name string `json:"name"` // The name of the repository property to target
	Property_values []string `json:"property_values"` // The values to match for the repository property
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"` // The GraphQL identifier of the repository.
	Html_url string `json:"html_url"` // The URL to view the repository on GitHub.com.
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Notifications_url string `json:"notifications_url"` // A template for the API URL to get information about notifications on the repository.
	Issue_events_url string `json:"issue_events_url"` // A template for the API URL to get information about issue events on the repository.
	Downloads_url string `json:"downloads_url"` // The API URL to list the downloads on the repository.
	Commits_url string `json:"commits_url"` // A template for the API URL to get information about commits on the repository.
	Full_name string `json:"full_name"` // The full, globally unique, name of the repository.
	Id int64 `json:"id"` // A unique identifier of the repository.
	Subscribers_url string `json:"subscribers_url"` // The API URL to list the subscribers on the repository.
	Teams_url string `json:"teams_url"` // The API URL to list the teams on the repository.
	Branches_url string `json:"branches_url"` // A template for the API URL to get information about branches in the repository.
	Contributors_url string `json:"contributors_url"` // A template for the API URL to list the contributors to the repository.
	Merges_url string `json:"merges_url"` // The API URL to merge branches in the repository.
	Git_tags_url string `json:"git_tags_url"` // A template for the API URL to get information about Git tags of the repository.
	Statuses_url string `json:"statuses_url"` // A template for the API URL to get information about statuses of a commit.
	Labels_url string `json:"labels_url"` // A template for the API URL to get information about labels of the repository.
	Private bool `json:"private"` // Whether the repository is private.
	Forks_url string `json:"forks_url"` // The API URL to list the forks of the repository.
	Issues_url string `json:"issues_url"` // A template for the API URL to get information about issues on the repository.
	Git_commits_url string `json:"git_commits_url"` // A template for the API URL to get information about Git commits of the repository.
	Assignees_url string `json:"assignees_url"` // A template for the API URL to list the available assignees for issues in the repository.
	Tags_url string `json:"tags_url"` // The API URL to get information about tags on the repository.
	Git_refs_url string `json:"git_refs_url"` // A template for the API URL to get information about Git refs of the repository.
	Name string `json:"name"` // The name of the repository.
	Collaborators_url string `json:"collaborators_url"` // A template for the API URL to get information about collaborators of the repository.
	Description string `json:"description"` // The repository description.
	Keys_url string `json:"keys_url"` // A template for the API URL to get information about deploy keys on the repository.
	Blobs_url string `json:"blobs_url"` // A template for the API URL to create or retrieve a raw Git blob in the repository.
	Archive_url string `json:"archive_url"` // A template for the API URL to download the repository as an archive.
	Deployments_url string `json:"deployments_url"` // The API URL to list the deployments of the repository.
	Pulls_url string `json:"pulls_url"` // A template for the API URL to get information about pull requests on the repository.
	Stargazers_url string `json:"stargazers_url"` // The API URL to list the stargazers on the repository.
	Fork bool `json:"fork"` // Whether the repository is a fork.
	Compare_url string `json:"compare_url"` // A template for the API URL to compare two commits or refs.
	Contents_url string `json:"contents_url"` // A template for the API URL to get the contents of the repository.
	Issue_comment_url string `json:"issue_comment_url"` // A template for the API URL to get information about issue comments on the repository.
	Releases_url string `json:"releases_url"` // A template for the API URL to get information about releases on the repository.
	Subscription_url string `json:"subscription_url"` // The API URL to subscribe to notifications for this repository.
	Milestones_url string `json:"milestones_url"` // A template for the API URL to get information about milestones of the repository.
	Comments_url string `json:"comments_url"` // A template for the API URL to get information about comments on the repository.
	Url string `json:"url"` // The URL to get more information about the repository from the GitHub API.
	Hooks_url string `json:"hooks_url"` // The API URL to list the hooks on the repository.
	Languages_url string `json:"languages_url"` // The API URL to get information about the languages of the repository.
	Trees_url string `json:"trees_url"` // A template for the API URL to create or retrieve a raw Git tree of the repository.
	Events_url string `json:"events_url"` // The API URL to list the events of the repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name of the secret.
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Payload string `json:"payload"` // A URL-encoded string of the check_run.requested_action JSON payload. The decoded payload is a JSON object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Is_enabled bool `json:"is_enabled"` // Whether or not the issue type is enabled at the organization level.
	Name string `json:"name"` // Name of the issue type.
	Color string `json:"color,omitempty"` // Color for the issue type.
	Description string `json:"description,omitempty"` // Description of the issue type.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Duration float64 `json:"duration,omitempty"`
	Id string `json:"id"`
	Start_date string `json:"start_date,omitempty"`
	Title string `json:"title"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Status string `json:"status,omitempty"` // The current status of the deployment.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone3 `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_advisory GeneratedType `json:"repository_advisory"` // A repository security advisory.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Change_status map[string]interface{} `json:"change_status"`
	Committed_at string `json:"committed_at"`
	Url string `json:"url"`
	User GeneratedType `json:"user"` // A GitHub user.
	Version string `json:"version"`
}

// Webhookssponsorship represents the Webhookssponsorship schema from the OpenAPI specification
type Webhookssponsorship struct {
	Created_at string `json:"created_at"`
	Maintainer map[string]interface{} `json:"maintainer,omitempty"`
	Node_id string `json:"node_id"`
	Privacy_level string `json:"privacy_level"`
	Sponsor map[string]interface{} `json:"sponsor"`
	Sponsorable map[string]interface{} `json:"sponsorable"`
	Tier map[string]interface{} `json:"tier"` // The `tier_changed` and `pending_tier_change` will include the original tier before the change or pending change. For more information, see the pending tier change payload.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Spdx_id string `json:"spdx_id"`
	Url string `json:"url"`
	Html_url string `json:"html_url,omitempty"`
	Key string `json:"key"`
	Name string `json:"name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Client_payload map[string]interface{} `json:"client_payload"` // The `client_payload` that was specified in the `POST /repos/{owner}/{repo}/dispatches` request body.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The `event_type` that was specified in the `POST /repos/{owner}/{repo}/dispatches` request body.
	Branch string `json:"branch"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Author GeneratedType `json:"author"` // A GitHub user.
	Body_version string `json:"body_version"` // The current version of the body content. If provided, this update operation will be rejected if the given version does not match the latest version on the server.
	Reactions GeneratedType `json:"reactions,omitempty"`
	Discussion_url string `json:"discussion_url"`
	Html_url string `json:"html_url"`
	Number int `json:"number"` // The unique sequence number of a team discussion comment.
	Body_html string `json:"body_html"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Body string `json:"body"` // The main text of the comment.
	Last_edited_at string `json:"last_edited_at"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id,omitempty"` // Unique identifier of the label.
	Name string `json:"name"` // Name of the label.
	TypeField string `json:"type,omitempty"` // The type of label. Read-only labels are applied automatically when the runner is configured.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Version string `json:"version,omitempty"`
	Change_status map[string]interface{} `json:"change_status,omitempty"`
	Committed_at string `json:"committed_at,omitempty"`
	Url string `json:"url,omitempty"`
	User GeneratedType `json:"user,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Uploader GeneratedType `json:"uploader"` // A GitHub user.
	Id int `json:"id"` // The ID of the CodeQL database.
	Name string `json:"name"` // The name of the CodeQL database.
	Size int `json:"size"` // The size of the CodeQL database file in bytes.
	Updated_at string `json:"updated_at"` // The date and time at which the CodeQL database was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Commit_oid string `json:"commit_oid,omitempty"` // The commit SHA of the repository at the time the CodeQL database was created.
	Url string `json:"url"` // The URL at which to download the CodeQL database. The `Accept` header must be set to the value of the `content_type` property.
	Content_type string `json:"content_type"` // The MIME type of the CodeQL database file.
	Created_at string `json:"created_at"` // The date and time at which the CodeQL database was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Language string `json:"language"` // The language of the CodeQL database.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action,omitempty"`
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Git_url string `json:"git_url"`
	Name string `json:"name"`
	Html_url string `json:"html_url"`
	Last_modified_at string `json:"last_modified_at,omitempty"`
	Url string `json:"url"`
	File_size int `json:"file_size,omitempty"`
	Language string `json:"language,omitempty"`
	Line_numbers []string `json:"line_numbers,omitempty"`
	Path string `json:"path"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Score float64 `json:"score"`
	Sha string `json:"sha"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ide_chat string `json:"ide_chat,omitempty"` // The organization policy for allowing or disallowing Copilot Chat in the IDE.
	Plan_type string `json:"plan_type,omitempty"` // The Copilot plan of the organization, or the parent enterprise, when applicable.
	Platform_chat string `json:"platform_chat,omitempty"` // The organization policy for allowing or disallowing Copilot features on GitHub.com.
	Public_code_suggestions string `json:"public_code_suggestions"` // The organization policy for allowing or blocking suggestions matching public code (duplication detection filter).
	Seat_breakdown GeneratedType `json:"seat_breakdown"` // The breakdown of Copilot Business seats for the organization.
	Seat_management_setting string `json:"seat_management_setting"` // The mode of assigning new seats.
	Cli string `json:"cli,omitempty"` // The organization policy for allowing or disallowing Copilot in the CLI.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Blocked_user Webhooksuser `json:"blocked_user"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name pattern that branches must match in order to deploy to the environment. Wildcard characters will not match `/`. For example, to match branches that begin with `release/` and contain an additional single slash, use `release/*/*`. For more information about pattern matching syntax, see the [Ruby File.fnmatch documentation](https://ruby-doc.org/core-2.5.1/File.html#method-c-fnmatch).
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Reactions GeneratedType `json:"reactions,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	Issue_url string `json:"issue_url"`
	Url string `json:"url"` // URL for the issue comment
	User GeneratedType `json:"user"` // A GitHub user.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Body string `json:"body,omitempty"` // Contents of the issue comment
	Body_html string `json:"body_html,omitempty"`
	Html_url string `json:"html_url"`
	Id int64 `json:"id"` // Unique identifier of the issue comment
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pull_request_review_url string `json:"pull_request_review_url"` // The API URL to get the pull request review where the secret was detected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Score float64 `json:"score"`
	Url string `json:"url"`
	Author GeneratedType `json:"author"` // A GitHub user.
	Commit map[string]interface{} `json:"commit"`
	Html_url string `json:"html_url"`
	Node_id string `json:"node_id"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Sha string `json:"sha"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Comments_url string `json:"comments_url"`
	Parents []map[string]interface{} `json:"parents"`
	Committer GeneratedType `json:"committer"` // Metaproperties for Git author/committer information.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Node_id string `json:"node_id"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Url string `json:"url"`
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_url string `json:"commit_url"`
	Id int `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository_id map[string]interface{} `json:"repository_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Definition GeneratedType `json:"definition"` // Custom property defined on an organization
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	Tag string `json:"tag"` // Name of the tag
	Tagger map[string]interface{} `json:"tagger"`
	Url string `json:"url"` // URL for the tag
	Verification Verification `json:"verification,omitempty"`
	Message string `json:"message"` // Message describing the purpose of the tag
	Node_id string `json:"node_id"`
	Object map[string]interface{} `json:"object"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Login string `json:"login"`
	Node_id string `json:"node_id,omitempty"`
	Organization_billing_email string `json:"organization_billing_email,omitempty"`
	TypeField string `json:"type"`
	Url string `json:"url"`
	Email string `json:"email,omitempty"`
	Id int `json:"id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Blocked_user Webhooksuser `json:"blocked_user"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Target_ref string `json:"target_ref,omitempty"` // The Git reference of target branch for the commit. For more information, see "[Git References](https://git-scm.com/book/en/v2/Git-Internals-Git-References)" in the Git documentation.
	Sha string `json:"sha,omitempty"` // SHA of commit with autofix.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Check_suite map[string]interface{} `json:"check_suite"` // The [check_suite](https://docs.github.com/rest/checks/suites#get-a-check-suite).
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Webhooksusermannequin represents the Webhooksusermannequin schema from the OpenAPI specification
type Webhooksusermannequin struct {
	Received_events_url string `json:"received_events_url,omitempty"`
	Site_admin bool `json:"site_admin,omitempty"`
	Url string `json:"url,omitempty"`
	Avatar_url string `json:"avatar_url,omitempty"`
	Gists_url string `json:"gists_url,omitempty"`
	Organizations_url string `json:"organizations_url,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Login string `json:"login"`
	Subscriptions_url string `json:"subscriptions_url,omitempty"`
	Starred_url string `json:"starred_url,omitempty"`
	Gravatar_id string `json:"gravatar_id,omitempty"`
	Id int `json:"id"`
	Deleted bool `json:"deleted,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Email string `json:"email,omitempty"`
	Events_url string `json:"events_url,omitempty"`
	Followers_url string `json:"followers_url,omitempty"`
	Repos_url string `json:"repos_url,omitempty"`
	Following_url string `json:"following_url,omitempty"`
	Name string `json:"name,omitempty"`
	TypeField string `json:"type,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Billable map[string]interface{} `json:"billable"`
	Run_duration_ms int `json:"run_duration_ms,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"`
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Target_type string `json:"target_type"`
	Action string `json:"action"`
	Account map[string]interface{} `json:"account"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Output map[string]interface{} `json:"output"`
	Html_url string `json:"html_url"`
	Started_at string `json:"started_at"`
	External_id string `json:"external_id"`
	Name string `json:"name"` // The name of the check.
	Completed_at string `json:"completed_at"`
	Url string `json:"url"`
	Details_url string `json:"details_url"`
	App Integration `json:"app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Check_suite GeneratedType `json:"check_suite"` // A suite of checks performed on the code of a given code change
	Deployment GeneratedType `json:"deployment,omitempty"` // A deployment created as the result of an Actions check run from a workflow that references an environment
	Head_sha string `json:"head_sha"` // The SHA of the commit that is being checked.
	Pull_requests []GeneratedType `json:"pull_requests"`
	Conclusion string `json:"conclusion"`
	Node_id string `json:"node_id"`
	Id int `json:"id"` // The id of the check.
	Status string `json:"status"` // The phase of the lifecycle that the check is currently in.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Percentile float64 `json:"percentile,omitempty"`
	Percentage float64 `json:"percentage,omitempty"`
}

// Traffic represents the Traffic schema from the OpenAPI specification
type Traffic struct {
	Count int `json:"count"`
	Timestamp string `json:"timestamp"`
	Uniques int `json:"uniques"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Forks_url string `json:"forks_url,omitempty"`
	Git_push_url string `json:"git_push_url,omitempty"`
	Id string `json:"id,omitempty"`
	Comments_enabled bool `json:"comments_enabled,omitempty"`
	Owner GeneratedType `json:"owner,omitempty"` // A GitHub user.
	Comments int `json:"comments,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Forks []map[string]interface{} `json:"forks,omitempty"`
	Node_id string `json:"node_id,omitempty"`
	Commits_url string `json:"commits_url,omitempty"`
	History []GeneratedType `json:"history,omitempty"`
	Files map[string]interface{} `json:"files,omitempty"`
	Url string `json:"url,omitempty"`
	Description string `json:"description,omitempty"`
	Html_url string `json:"html_url,omitempty"`
	Comments_url string `json:"comments_url,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Fork_of map[string]interface{} `json:"fork_of,omitempty"` // Gist
	Git_pull_url string `json:"git_pull_url,omitempty"`
	Public bool `json:"public,omitempty"`
	User string `json:"user,omitempty"`
	Truncated bool `json:"truncated,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Membership Webhooksmembership `json:"membership"` // The membership between the user and the organization. Not present when the action is `member_invited`.
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Start_line float64 `json:"start_line"` // Line number at which the secret starts in the file
	Start_column float64 `json:"start_column"` // The column at which the secret starts within the start line when the file is interpreted as 8-bit ASCII.
	Commit_sha string `json:"commit_sha"` // SHA-1 hash ID of the associated commit
	Page_url string `json:"page_url"` // The GitHub URL to get the associated wiki page
	Path string `json:"path"` // The file path of the wiki page
	Blob_sha string `json:"blob_sha"` // SHA-1 hash ID of the associated blob
	Commit_url string `json:"commit_url"` // The GitHub URL to get the associated wiki commit
	End_column float64 `json:"end_column"` // The column at which the secret ends within the end line when the file is interpreted as 8-bit ASCII.
	End_line float64 `json:"end_line"` // Line number at which the secret ends in the file
}

// Integration represents the Integration schema from the OpenAPI specification
type Integration struct {
	Created_at string `json:"created_at"`
	Owner interface{} `json:"owner"`
	Permissions map[string]interface{} `json:"permissions"` // The set of permissions for the GitHub app
	Slug string `json:"slug,omitempty"` // The slug name of the GitHub app
	Description string `json:"description"`
	External_url string `json:"external_url"`
	Name string `json:"name"` // The name of the GitHub app
	Events []string `json:"events"` // The list of events for the GitHub app. Note that the `installation_target`, `security_advisory`, and `meta` events are not included because they are global events and not specific to an installation.
	Html_url string `json:"html_url"`
	Id int `json:"id"` // Unique identifier of the GitHub app
	Installations_count int `json:"installations_count,omitempty"` // The number of installations associated with the GitHub app. Only returned when the integration is requesting details about itself.
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Client_id string `json:"client_id,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository map[string]interface{} `json:"repository,omitempty"` // A git repository
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Team Webhooksteam1 `json:"team"` // Groups of organization members that gives permissions on specified repositories.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actor_type string `json:"actor_type"` // The type of actor that can bypass a ruleset.
	Bypass_mode string `json:"bypass_mode,omitempty"` // When the specified actor can bypass the ruleset. `pull_request` means that an actor can only bypass rules on pull requests. `pull_request` is not applicable for the `DeployKey` actor type. Also, `pull_request` is only applicable to branch rulesets.
	Actor_id int `json:"actor_id,omitempty"` // The ID of the actor that can bypass a ruleset. If `actor_type` is `OrganizationAdmin`, this should be `1`. If `actor_type` is `DeployKey`, this should be null. `OrganizationAdmin` is not applicable for personal repositories.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id string `json:"id,omitempty"` // An identifier for the upload.
	Url string `json:"url,omitempty"` // The REST API URL for checking the status of the upload.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes Webhooksprojectchanges `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sha string `json:"sha"`
	Path string `json:"path"`
	Ref string `json:"ref,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Count int `json:"count"`
	Referrer string `json:"referrer"`
	Uniques int `json:"uniques"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Description string `json:"description"` // The repository's current description.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Master_branch string `json:"master_branch"` // The name of the repository's default branch (usually `main`).
	Ref_type string `json:"ref_type"` // The type of Git ref object created in the repository.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Ref string `json:"ref"` // The [`git ref`](https://docs.github.com/rest/git/refs#get-a-reference) resource.
	Pusher_type string `json:"pusher_type"` // The pusher type for the event. Can be either `user` or a deploy key.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// Authorization represents the Authorization schema from the OpenAPI specification
type Authorization struct {
	Expires_at string `json:"expires_at"`
	Note_url string `json:"note_url"`
	App map[string]interface{} `json:"app"`
	Id int64 `json:"id"`
	Token string `json:"token"`
	Url string `json:"url"`
	Fingerprint string `json:"fingerprint"`
	Hashed_token string `json:"hashed_token"`
	Installation GeneratedType `json:"installation,omitempty"`
	Note string `json:"note"`
	User GeneratedType `json:"user,omitempty"` // A GitHub user.
	Token_last_eight string `json:"token_last_eight"`
	Updated_at string `json:"updated_at"`
	Created_at string `json:"created_at"`
	Scopes []string `json:"scopes"` // A list of scopes that this authorization is in.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key_id string `json:"key_id"`
	Revoked bool `json:"revoked"`
	Created_at string `json:"created_at"`
	Primary_key_id int `json:"primary_key_id"`
	Can_sign bool `json:"can_sign"`
	Can_certify bool `json:"can_certify"`
	Id int64 `json:"id"`
	Subkeys []map[string]interface{} `json:"subkeys"`
	Can_encrypt_storage bool `json:"can_encrypt_storage"`
	Public_key string `json:"public_key"`
	Emails []map[string]interface{} `json:"emails"`
	Can_encrypt_comms bool `json:"can_encrypt_comms"`
	Expires_at string `json:"expires_at"`
	Name string `json:"name,omitempty"`
	Raw_key string `json:"raw_key"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_active_caches_count int `json:"total_active_caches_count"` // The count of active caches across all repositories of an enterprise or an organization.
	Total_active_caches_size_in_bytes int `json:"total_active_caches_size_in_bytes"` // The total size in bytes of all active cache items across all repositories of an enterprise or an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// Page represents the Page schema from the OpenAPI specification
type Page struct {
	Build_type string `json:"build_type,omitempty"` // The process in which the Page will be built.
	Https_enforced bool `json:"https_enforced,omitempty"` // Whether https is enabled on the domain
	Pending_domain_unverified_at string `json:"pending_domain_unverified_at,omitempty"` // The timestamp when a pending domain becomes unverified.
	Protected_domain_state string `json:"protected_domain_state,omitempty"` // The state if the domain is verified
	Custom_404 bool `json:"custom_404"` // Whether the Page has a custom 404 page.
	Source GeneratedType `json:"source,omitempty"`
	Html_url string `json:"html_url,omitempty"` // The web address the Page can be accessed from.
	Cname string `json:"cname"` // The Pages site's custom domain
	Https_certificate GeneratedType `json:"https_certificate,omitempty"`
	Public bool `json:"public"` // Whether the GitHub Pages site is publicly visible. If set to `true`, the site is accessible to anyone on the internet. If set to `false`, the site will only be accessible to users who have at least `read` access to the repository that published the site.
	Status string `json:"status"` // The status of the most recent build of the Page.
	Url string `json:"url"` // The API address for accessing this Page resource.
}

// Vulnerability represents the Vulnerability schema from the OpenAPI specification
type Vulnerability struct {
	First_patched_version string `json:"first_patched_version"` // The package version that resolves the vulnerability.
	PackageField map[string]interface{} `json:"package"` // The name of the package affected by the vulnerability.
	Vulnerable_functions []string `json:"vulnerable_functions"` // The functions in the package that are affected by the vulnerability.
	Vulnerable_version_range string `json:"vulnerable_version_range"` // The range of the package versions affected by the vulnerability.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Cards_url string `json:"cards_url"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // The unique identifier of the project column
	Name string `json:"name"` // Name of the project column
	Node_id string `json:"node_id"`
	Project_url string `json:"project_url"`
}

// Blob represents the Blob schema from the OpenAPI specification
type Blob struct {
	Content string `json:"content"`
	Encoding string `json:"encoding"`
	Highlighted_content string `json:"highlighted_content,omitempty"`
	Node_id string `json:"node_id"`
	Sha string `json:"sha"`
	Size int `json:"size"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Confirm_delete_url string `json:"confirm_delete_url"` // Next deletable analysis in chain, with last analysis deletion confirmation
	Next_analysis_url string `json:"next_analysis_url"` // Next deletable analysis in chain, without last analysis deletion confirmation
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	User Webhooksuser `json:"user,omitempty"`
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Invitation map[string]interface{} `json:"invitation"` // The invitation for the user or email if the action is `member_invited`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comments []GeneratedType `json:"comments,omitempty"`
	Event string `json:"event,omitempty"`
	Node_id string `json:"node_id,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Most_recent_instance GeneratedType `json:"most_recent_instance"`
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	State string `json:"state"` // State of a code scanning alert.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Number int `json:"number"` // The security alert number.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Rule GeneratedType `json:"rule"`
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Tool GeneratedType `json:"tool"`
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Dismissal_approved_by GeneratedType `json:"dismissal_approved_by,omitempty"` // A GitHub user.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // Full Repository
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Export_url string `json:"export_url,omitempty"` // Url for fetching export details
	Html_url string `json:"html_url,omitempty"` // Web url for the exported branch
	Id string `json:"id,omitempty"` // Id for the export details
	Sha string `json:"sha,omitempty"` // Git commit SHA of the exported branch
	State string `json:"state,omitempty"` // State of the latest export
	Branch string `json:"branch,omitempty"` // Name of the exported branch
	Completed_at string `json:"completed_at,omitempty"` // Completion time of the last export operation
}

// Issue represents the Issue schema from the OpenAPI specification
type Issue struct {
	Repository_url string `json:"repository_url"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	Labels []interface{} `json:"labels"` // Labels to associate with this issue; pass one or more label names to replace the set of labels on this issue; send an empty array to clear all labels from the issue; note that the labels are silently dropped for users without push access to the repository
	Locked bool `json:"locked"`
	State_reason string `json:"state_reason,omitempty"` // The reason for the current state
	Body string `json:"body,omitempty"` // Contents of the issue
	Title string `json:"title"` // Title of the issue
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Id int64 `json:"id"`
	TypeField GeneratedType `json:"type,omitempty"` // The type of issue.
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Url string `json:"url"` // URL for the issue
	Body_html string `json:"body_html,omitempty"`
	Reactions GeneratedType `json:"reactions,omitempty"`
	Node_id string `json:"node_id"`
	Comments int `json:"comments"`
	Labels_url string `json:"labels_url"`
	Html_url string `json:"html_url"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Created_at string `json:"created_at"`
	User GeneratedType `json:"user"` // A GitHub user.
	Events_url string `json:"events_url"`
	Updated_at string `json:"updated_at"`
	Closed_at string `json:"closed_at"`
	Closed_by GeneratedType `json:"closed_by,omitempty"` // A GitHub user.
	Timeline_url string `json:"timeline_url,omitempty"`
	Body_text string `json:"body_text,omitempty"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Comments_url string `json:"comments_url"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Sub_issues_summary GeneratedType `json:"sub_issues_summary,omitempty"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	State string `json:"state"` // State of the issue; either 'open' or 'closed'
	Draft bool `json:"draft,omitempty"`
	Number int `json:"number"` // Number uniquely identifying the issue within its repository
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Release Webhooksrelease1 `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url,omitempty"`
	Sha string `json:"sha"`
	Tree []map[string]interface{} `json:"tree"` // Objects specifying a tree structure
	Truncated bool `json:"truncated"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2 GeneratedType `json:"projects_v2"` // A projects v2 project
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Release map[string]interface{} `json:"release"` // The [release](https://docs.github.com/rest/releases/releases/#get-a-release) object.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// Milestone represents the Milestone schema from the OpenAPI specification
type Milestone struct {
	Created_at string `json:"created_at"`
	Title string `json:"title"` // The title of the milestone.
	Labels_url string `json:"labels_url"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Html_url string `json:"html_url"`
	Number int `json:"number"` // The number of the milestone.
	State string `json:"state"` // The state of the milestone.
	Url string `json:"url"`
	Closed_at string `json:"closed_at"`
	Description string `json:"description"`
	Due_on string `json:"due_on"`
	Updated_at string `json:"updated_at"`
	Open_issues int `json:"open_issues"`
	Closed_issues int `json:"closed_issues"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Tarball_url string `json:"tarball_url"`
	Zipball_url string `json:"zipball_url"`
	Commit map[string]interface{} `json:"commit"`
	Name string `json:"name"`
	Node_id string `json:"node_id"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Controller_repo GeneratedType `json:"controller_repo"` // A GitHub repository.
	Created_at string `json:"created_at,omitempty"` // The date and time at which the variant analysis was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Failure_reason string `json:"failure_reason,omitempty"` // The reason for a failure of the variant analysis. This is only available if the variant analysis has failed.
	Query_language string `json:"query_language"` // The language targeted by the CodeQL query
	Query_pack_url string `json:"query_pack_url"` // The download url for the query pack.
	Skipped_repositories map[string]interface{} `json:"skipped_repositories,omitempty"` // Information about repositories that were skipped from processing. This information is only available to the user that initiated the variant analysis.
	Status string `json:"status"`
	Updated_at string `json:"updated_at,omitempty"` // The date and time at which the variant analysis was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Completed_at string `json:"completed_at,omitempty"` // The date and time at which the variant analysis was completed, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ. Will be null if the variant analysis has not yet completed or this information is not available.
	Scanned_repositories []map[string]interface{} `json:"scanned_repositories,omitempty"`
	Id int `json:"id"` // The ID of the variant analysis.
	Actions_workflow_run_id int `json:"actions_workflow_run_id,omitempty"` // The GitHub Actions workflow run used to execute this variant analysis. This is only available if the workflow run has started.
	Actor GeneratedType `json:"actor"` // A GitHub user.
}

// License represents the License schema from the OpenAPI specification
type License struct {
	Name string `json:"name"`
	Spdx_id string `json:"spdx_id"`
	Description string `json:"description"`
	Featured bool `json:"featured"`
	Url string `json:"url"`
	Conditions []string `json:"conditions"`
	Implementation string `json:"implementation"`
	Key string `json:"key"`
	Permissions []string `json:"permissions"`
	Body string `json:"body"`
	Node_id string `json:"node_id"`
	Html_url string `json:"html_url"`
	Limitations []string `json:"limitations"`
}

// Deployment represents the Deployment schema from the OpenAPI specification
type Deployment struct {
	Node_id string `json:"node_id"`
	Payload interface{} `json:"payload"`
	Repository_url string `json:"repository_url"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Description string `json:"description"`
	Id int64 `json:"id"` // Unique identifier of the deployment
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Original_environment string `json:"original_environment,omitempty"`
	Production_environment bool `json:"production_environment,omitempty"` // Specifies if the given environment is one that end-users directly interact with. Default: false.
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Ref string `json:"ref"` // The ref to deploy. This can be a branch, tag, or sha.
	Sha string `json:"sha"`
	Statuses_url string `json:"statuses_url"`
	Created_at string `json:"created_at"`
	Environment string `json:"environment"` // Name for the target deployment environment.
	Task string `json:"task"` // Parameter to specify a task to execute
	Transient_environment bool `json:"transient_environment,omitempty"` // Specifies if the given environment is will no longer exist at some point in the future. Default: false.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Marketplace_purchase Webhooksmarketplacepurchase `json:"marketplace_purchase"`
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Previous_marketplace_purchase map[string]interface{} `json:"previous_marketplace_purchase,omitempty"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Effective_date string `json:"effective_date"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_item GeneratedType `json:"projects_v2_item"` // An item belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Content map[string]interface{} `json:"content"`
	Commit map[string]interface{} `json:"commit"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Dismissal_approved_by GeneratedType `json:"dismissal_approved_by,omitempty"` // A GitHub user.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_reason string `json:"dismissed_reason"` // **Required when the state is dismissed.** The reason for dismissing or closing the alert.
	Instances_url string `json:"instances_url"` // The REST API URL for fetching the list of instances for an alert.
	Rule GeneratedType `json:"rule"`
	Tool GeneratedType `json:"tool"`
	Repository GeneratedType `json:"repository"` // A GitHub repository.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Fixed_at string `json:"fixed_at,omitempty"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	State string `json:"state"` // State of a code scanning alert.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Most_recent_instance GeneratedType `json:"most_recent_instance"`
	Dismissed_comment string `json:"dismissed_comment,omitempty"` // The dismissal comment associated with the dismissal of the alert.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Number int `json:"number"` // The security alert number.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Creator GeneratedType `json:"creator,omitempty"` // A GitHub user.
	Project_node_id string `json:"project_node_id,omitempty"`
	Id float64 `json:"id"`
	Node_id string `json:"node_id"`
	Status string `json:"status,omitempty"`
	Target_date string `json:"target_date,omitempty"`
	Updated_at string `json:"updated_at"`
	Body string `json:"body,omitempty"` // Body of the status update
	Created_at string `json:"created_at"`
	Start_date string `json:"start_date,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue Webhooksissue `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"` // The action that was performed.
	Assignee Webhooksusermannequin `json:"assignee,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Commit_id string `json:"commit_id"`
	Created_at string `json:"created_at"`
	Node_id string `json:"node_id"`
	Url string `json:"url"`
	Commit_url string `json:"commit_url"`
	Id int `json:"id"`
	Event string `json:"event"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Project_card map[string]interface{} `json:"project_card,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Is_base64_encoded bool `json:"is_base64_encoded,omitempty"` // A boolean value representing whether or not alert is base64 encoded
	Number int `json:"number,omitempty"` // The security alert number.
	Push_protection_bypassed_by GeneratedType `json:"push_protection_bypassed_by,omitempty"` // A GitHub user.
	Publicly_leaked bool `json:"publicly_leaked,omitempty"` // Whether the detected secret was publicly leaked.
	Push_protection_bypassed_at string `json:"push_protection_bypassed_at,omitempty"` // The time that push protection was bypassed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Resolved_at string `json:"resolved_at,omitempty"` // The time that the alert was resolved in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Secret_type_display_name string `json:"secret_type_display_name,omitempty"` // User-friendly name for the detected secret, matching the `secret_type`. For a list of built-in patterns, see "[Supported secret scanning patterns](https://docs.github.com/code-security/secret-scanning/introduction/supported-secret-scanning-patterns#supported-secrets)."
	Validity string `json:"validity,omitempty"` // The token status as of the latest validity check.
	Created_at string `json:"created_at,omitempty"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypassed bool `json:"push_protection_bypassed,omitempty"` // Whether push protection was bypassed for the detected secret.
	State string `json:"state,omitempty"` // Sets the state of the secret scanning alert. You must provide `resolution` when you set the state to `resolved`.
	Html_url string `json:"html_url,omitempty"` // The GitHub URL of the alert resource.
	Resolved_by GeneratedType `json:"resolved_by,omitempty"` // A GitHub user.
	Url string `json:"url,omitempty"` // The REST API URL of the alert resource.
	Locations_url string `json:"locations_url,omitempty"` // The REST API URL of the code locations for this alert.
	Push_protection_bypass_request_reviewer GeneratedType `json:"push_protection_bypass_request_reviewer,omitempty"` // A GitHub user.
	Secret string `json:"secret,omitempty"` // The secret that was detected.
	First_location_detected GeneratedType `json:"first_location_detected,omitempty"` // Details on the location where the token was initially detected. This can be a commit, wiki commit, issue, discussion, pull request.
	Updated_at string `json:"updated_at,omitempty"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Push_protection_bypass_request_html_url string `json:"push_protection_bypass_request_html_url,omitempty"` // The URL to a push protection bypass request.
	Push_protection_bypass_request_reviewer_comment string `json:"push_protection_bypass_request_reviewer_comment,omitempty"` // An optional comment when reviewing a push protection bypass.
	Has_more_locations bool `json:"has_more_locations,omitempty"` // A boolean value representing whether or not the token in the alert was detected in more than one location.
	Multi_repo bool `json:"multi_repo,omitempty"` // Whether the detected secret was found in multiple repositories under the same organization or enterprise.
	Push_protection_bypass_request_comment string `json:"push_protection_bypass_request_comment,omitempty"` // An optional comment when requesting a push protection bypass.
	Secret_type string `json:"secret_type,omitempty"` // The type of secret that secret scanning detected.
	Resolution string `json:"resolution,omitempty"` // **Required when the `state` is `resolved`.** The reason for resolving the alert.
	Resolution_comment string `json:"resolution_comment,omitempty"` // An optional comment to resolve an alert.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total_paid_gigabytes_bandwidth_used int `json:"total_paid_gigabytes_bandwidth_used"` // Total paid storage space (GB) for GitHuub Packages.
	Included_gigabytes_bandwidth int `json:"included_gigabytes_bandwidth"` // Free storage space (GB) for GitHub Packages.
	Total_gigabytes_bandwidth_used int `json:"total_gigabytes_bandwidth_used"` // Sum of the free and paid storage space (GB) for GitHuub Packages.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Subscriptions_url string `json:"subscriptions_url"`
	Login string `json:"login"`
	Node_id string `json:"node_id"`
	Followers_url string `json:"followers_url"`
	Gists_url string `json:"gists_url"`
	Organizations_url string `json:"organizations_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Avatar_url string `json:"avatar_url"`
	Email string `json:"email,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Html_url string `json:"html_url"`
	TypeField string `json:"type"`
	Id int64 `json:"id"`
	Repos_url string `json:"repos_url"`
	Starred_url string `json:"starred_url"`
	Following_url string `json:"following_url"`
	Gravatar_id string `json:"gravatar_id"`
	Name string `json:"name,omitempty"`
	Url string `json:"url"`
	Events_url string `json:"events_url"`
	Received_events_url string `json:"received_events_url"`
	Site_admin bool `json:"site_admin"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	User GeneratedType `json:"user"` // Collaborator
	Permission string `json:"permission"`
	Role_name string `json:"role_name"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Number int `json:"number"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Pull_request map[string]interface{} `json:"pull_request"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Vulnerable_version_range string `json:"vulnerable_version_range"` // The range of the package versions affected by the vulnerability.
	PackageField map[string]interface{} `json:"package"` // The name of the package affected by the vulnerability.
	Patched_versions string `json:"patched_versions"` // The package version(s) that resolve the vulnerability.
	Vulnerable_functions []string `json:"vulnerable_functions"` // The functions in the package that are affected.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Personal_access_token_request GeneratedType `json:"personal_access_token_request"` // Details of a Personal Access Token Request.
}

// Reaction represents the Reaction schema from the OpenAPI specification
type Reaction struct {
	Content string `json:"content"` // The reaction to use
	Created_at string `json:"created_at"`
	Id int `json:"id"`
	Node_id string `json:"node_id"`
	User GeneratedType `json:"user"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Updated_at string `json:"updated_at,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Id int `json:"id,omitempty"`
	Pattern string `json:"pattern"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Auto_dismissed_at string `json:"auto_dismissed_at,omitempty"` // The time that the alert was auto-dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Dismissed_at string `json:"dismissed_at"` // The time that the alert was dismissed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Security_vulnerability GeneratedType `json:"security_vulnerability"` // Details pertaining to one vulnerable version range for the advisory.
	Updated_at string `json:"updated_at"` // The time that the alert was last updated in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Html_url string `json:"html_url"` // The GitHub URL of the alert resource.
	Url string `json:"url"` // The REST API URL of the alert resource.
	Dismissed_reason string `json:"dismissed_reason"` // The reason that the alert was dismissed.
	Created_at string `json:"created_at"` // The time that the alert was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Fixed_at string `json:"fixed_at"` // The time that the alert was no longer detected and was considered fixed in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	Number int `json:"number"` // The security alert number.
	State string `json:"state"` // The state of the Dependabot alert.
	Dismissed_comment string `json:"dismissed_comment"` // An optional comment associated with the alert's dismissal.
	Security_advisory GeneratedType `json:"security_advisory"` // Details for the GitHub Security Advisory.
	Dismissed_by GeneratedType `json:"dismissed_by"` // A GitHub user.
	Dependency map[string]interface{} `json:"dependency"` // Details for the vulnerable dependency.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id,omitempty"`
	Key string `json:"key"` // The Base64 encoded public key.
	Key_id string `json:"key_id"` // The identifier for the key.
	Title string `json:"title,omitempty"`
	Url string `json:"url,omitempty"`
	Created_at string `json:"created_at,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Id int `json:"id"` // ID of the reviewer which must review changes to matching files.
	TypeField string `json:"type"` // The type of the reviewer
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"` // The name pattern that branches or tags must match in order to deploy to the environment. Wildcard characters will not match `/`. For example, to match branches that begin with `release/` and contain an additional single slash, use `release/*/*`. For more information about pattern matching syntax, see the [Ruby File.fnmatch documentation](https://ruby-doc.org/core-2.5.1/File.html#method-c-fnmatch).
	TypeField string `json:"type,omitempty"` // Whether this rule targets a branch or tag
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Limit int `json:"limit"`
	Remaining int `json:"remaining"`
	Reset int `json:"reset"`
	Used int `json:"used"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Summary string `json:"summary"` // A short summary of the advisory.
	Vulnerabilities []map[string]interface{} `json:"vulnerabilities,omitempty"` // An array of products affected by the vulnerability detailed in a repository security advisory.
	Cvss_vector_string string `json:"cvss_vector_string,omitempty"` // The CVSS vector that calculates the severity of the advisory. You must choose between setting this field or `severity`.
	Cwe_ids []string `json:"cwe_ids,omitempty"` // A list of Common Weakness Enumeration (CWE) IDs.
	Description string `json:"description"` // A detailed description of what the advisory impacts.
	Severity string `json:"severity,omitempty"` // The severity of the advisory. You must choose between setting this field or `cvss_vector_string`.
	Start_private_fork bool `json:"start_private_fork,omitempty"` // Whether to create a temporary private fork of the repository to collaborate on a fix.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Comment string `json:"comment"` // Comment associated with the pending deployment protection rule. **Required when state is not provided.**
	Environment_name string `json:"environment_name"` // The name of the environment to approve or reject.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Changes int `json:"changes"`
	Contents_url string `json:"contents_url"`
	Deletions int `json:"deletions"`
	Previous_filename string `json:"previous_filename,omitempty"`
	Raw_url string `json:"raw_url"`
	Sha string `json:"sha"`
	Status string `json:"status"`
	Blob_url string `json:"blob_url"`
	Filename string `json:"filename"`
	Patch string `json:"patch,omitempty"`
	Additions int `json:"additions"`
}

// Email represents the Email schema from the OpenAPI specification
type Email struct {
	Email string `json:"email"`
	Primary bool `json:"primary"`
	Verified bool `json:"verified"`
	Visibility string `json:"visibility"`
}

// Webhooksreviewcomment represents the Webhooksreviewcomment schema from the OpenAPI specification
type Webhooksreviewcomment struct {
	Side string `json:"side"` // The side of the first line of the range for a multi-line comment.
	Html_url string `json:"html_url"` // HTML URL for the pull request review comment.
	Updated_at string `json:"updated_at"`
	Path string `json:"path"` // The relative path of the file to which the comment applies.
	Start_side string `json:"start_side"` // The side of the first line of the range for a multi-line comment.
	User map[string]interface{} `json:"user"`
	Line int `json:"line"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Original_start_line int `json:"original_start_line"` // The first line of the range for a multi-line comment.
	Position int `json:"position"` // The line index in the diff to which the comment applies.
	Start_line int `json:"start_line"` // The first line of the range for a multi-line comment.
	Url string `json:"url"` // URL for the pull request review comment
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Node_id string `json:"node_id"` // The node ID of the pull request review comment.
	Original_line int `json:"original_line"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Original_position int `json:"original_position"` // The index of the original line in the diff to which the comment applies.
	Original_commit_id string `json:"original_commit_id"` // The SHA of the original commit to which the comment applies.
	Subject_type string `json:"subject_type,omitempty"` // The level at which the comment is targeted, can be a diff line or a file.
	Body string `json:"body"` // The text of the comment.
	Created_at string `json:"created_at"`
	Links map[string]interface{} `json:"_links"`
	Diff_hunk string `json:"diff_hunk"` // The diff of the line that the comment refers to.
	Id int `json:"id"` // The ID of the pull request review comment.
	Commit_id string `json:"commit_id"` // The SHA of the commit to which the comment applies.
	In_reply_to_id int `json:"in_reply_to_id,omitempty"` // The comment ID to reply to.
	Pull_request_review_id int `json:"pull_request_review_id"` // The ID of the pull request review to which the comment belongs.
	Pull_request_url string `json:"pull_request_url"` // URL for the pull request that the review comment belongs to.
	Reactions map[string]interface{} `json:"reactions"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Key Webhooksdeploykey `json:"key"` // The [`deploy key`](https://docs.github.com/rest/deploy-keys/deploy-keys#get-a-deploy-key) resource.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Provider string `json:"provider"`
	Url string `json:"url"`
}

// Dependency represents the Dependency schema from the OpenAPI specification
type Dependency struct {
	Package_url string `json:"package_url,omitempty"` // Package-url (PURL) of dependency. See https://github.com/package-url/purl-spec for more details.
	Relationship string `json:"relationship,omitempty"` // A notation of whether a dependency is requested directly by this manifest or is a dependency of another dependency.
	Scope string `json:"scope,omitempty"` // A notation of whether the dependency is required for the primary build artifact (runtime) or is only used for development. Future versions of this specification may allow for more granular scopes.
	Dependencies []string `json:"dependencies,omitempty"` // Array of package-url (PURLs) of direct child dependencies.
	Metadata Metadata `json:"metadata,omitempty"` // User-defined metadata to store domain-specific information limited to 8 keys with scalar values.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"`
	Id int `json:"id"`
	Monthly_price_in_cents int `json:"monthly_price_in_cents"`
	Number int `json:"number"`
	State string `json:"state"`
	Accounts_url string `json:"accounts_url"`
	Has_free_trial bool `json:"has_free_trial"`
	Name string `json:"name"`
	Price_model string `json:"price_model"`
	Unit_name string `json:"unit_name"`
	Yearly_price_in_cents int `json:"yearly_price_in_cents"`
	Bullets []string `json:"bullets"`
	Description string `json:"description"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Project Webhooksproject `json:"project"`
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ecosystem string `json:"ecosystem"` // The package's language or package management ecosystem.
	Name string `json:"name"` // The unique package name within its ecosystem.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Total int `json:"total"`
	Completed int `json:"completed"`
	Percent_completed int `json:"percent_completed"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Timestamp string `json:"timestamp"` // Timestamp of the commit
	Tree_id string `json:"tree_id"` // SHA for the commit's tree
	Author map[string]interface{} `json:"author"` // Information about the Git author
	Committer map[string]interface{} `json:"committer"` // Information about the Git committer
	Id string `json:"id"` // SHA for the commit
	Message string `json:"message"` // Message describing the purpose of the commit
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Changes map[string]interface{} `json:"changes"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Discussion Discussion `json:"discussion"` // A Discussion in a repository.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	TypeField string `json:"type"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Rules []GeneratedType `json:"rules,omitempty"`
	Source_type string `json:"source_type,omitempty"` // The type of the source of the ruleset
	Conditions interface{} `json:"conditions,omitempty"`
	Current_user_can_bypass string `json:"current_user_can_bypass,omitempty"` // The bypass type of the user making the API request for this ruleset. This field is only returned when querying the repository-level endpoint.
	Enforcement string `json:"enforcement"` // The enforcement level of the ruleset. `evaluate` allows admins to test rules before enforcing them. Admins can view insights on the Rule Insights page (`evaluate` is only available with GitHub Enterprise).
	Updated_at string `json:"updated_at,omitempty"`
	Bypass_actors []GeneratedType `json:"bypass_actors,omitempty"` // The actors that can bypass the rules in this ruleset
	Id int `json:"id"` // The ID of the ruleset
	Name string `json:"name"` // The name of the ruleset
	Source string `json:"source"` // The name of the source
	Target string `json:"target,omitempty"` // The target of the ruleset
	Links map[string]interface{} `json:"_links,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Node_id string `json:"node_id,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Projects_v2_status_update GeneratedType `json:"projects_v2_status_update"` // An status update belonging to a project
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name"`
	Git_tags_url string `json:"git_tags_url"`
	Fork bool `json:"fork"`
	Url string `json:"url"`
	Has_projects bool `json:"has_projects"`
	Updated_at string `json:"updated_at"`
	Clone_url string `json:"clone_url"`
	Archived bool `json:"archived"`
	Forks_url string `json:"forks_url"`
	Trees_url string `json:"trees_url"`
	Has_issues bool `json:"has_issues"`
	Statuses_url string `json:"statuses_url"`
	Web_commit_signoff_required bool `json:"web_commit_signoff_required,omitempty"`
	Contributors_url string `json:"contributors_url"`
	Compare_url string `json:"compare_url"`
	Has_wiki bool `json:"has_wiki"`
	Private bool `json:"private"`
	Git_commits_url string `json:"git_commits_url"`
	Issue_events_url string `json:"issue_events_url"`
	Visibility string `json:"visibility,omitempty"` // The repository visibility: public, private, or internal.
	Collaborators_url string `json:"collaborators_url"`
	Homepage string `json:"homepage"`
	Tags_url string `json:"tags_url"`
	Commits_url string `json:"commits_url"`
	Language string `json:"language"`
	Node_id string `json:"node_id"`
	Delete_branch_on_merge bool `json:"delete_branch_on_merge,omitempty"`
	Issues_url string `json:"issues_url"`
	Mirror_url string `json:"mirror_url"`
	License GeneratedType `json:"license"` // License Simple
	Allow_merge_commit bool `json:"allow_merge_commit,omitempty"`
	Has_pages bool `json:"has_pages"`
	Assignees_url string `json:"assignees_url"`
	Issue_comment_url string `json:"issue_comment_url"`
	Id int `json:"id"`
	Allow_forking bool `json:"allow_forking,omitempty"`
	Comments_url string `json:"comments_url"`
	Subscribers_url string `json:"subscribers_url"`
	Forks int `json:"forks"`
	Default_branch string `json:"default_branch"`
	Open_issues int `json:"open_issues"`
	Score float64 `json:"score"`
	Deployments_url string `json:"deployments_url"`
	Full_name string `json:"full_name"`
	Labels_url string `json:"labels_url"`
	Svn_url string `json:"svn_url"`
	Subscription_url string `json:"subscription_url"`
	Ssh_url string `json:"ssh_url"`
	Keys_url string `json:"keys_url"`
	Stargazers_url string `json:"stargazers_url"`
	Stargazers_count int `json:"stargazers_count"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Languages_url string `json:"languages_url"`
	Notifications_url string `json:"notifications_url"`
	Archive_url string `json:"archive_url"`
	Blobs_url string `json:"blobs_url"`
	Pushed_at string `json:"pushed_at"`
	Created_at string `json:"created_at"`
	Temp_clone_token string `json:"temp_clone_token,omitempty"`
	Watchers int `json:"watchers"`
	Owner GeneratedType `json:"owner"` // A GitHub user.
	Downloads_url string `json:"downloads_url"`
	Size int `json:"size"`
	Contents_url string `json:"contents_url"`
	Merges_url string `json:"merges_url"`
	Topics []string `json:"topics,omitempty"`
	Disabled bool `json:"disabled"` // Returns whether or not this repository disabled.
	Allow_squash_merge bool `json:"allow_squash_merge,omitempty"`
	Branches_url string `json:"branches_url"`
	Has_downloads bool `json:"has_downloads"`
	Git_refs_url string `json:"git_refs_url"`
	Hooks_url string `json:"hooks_url"`
	Is_template bool `json:"is_template,omitempty"`
	Pulls_url string `json:"pulls_url"`
	Description string `json:"description"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Master_branch string `json:"master_branch,omitempty"`
	Has_discussions bool `json:"has_discussions,omitempty"`
	Allow_rebase_merge bool `json:"allow_rebase_merge,omitempty"`
	Open_issues_count int `json:"open_issues_count"`
	Events_url string `json:"events_url"`
	Forks_count int `json:"forks_count"`
	Html_url string `json:"html_url"`
	Watchers_count int `json:"watchers_count"`
	Git_url string `json:"git_url"`
	Releases_url string `json:"releases_url"`
	Teams_url string `json:"teams_url"`
	Allow_auto_merge bool `json:"allow_auto_merge,omitempty"`
	Milestones_url string `json:"milestones_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name,omitempty"` // The name of the tool used to generate the code scanning analysis.
	Version string `json:"version,omitempty"` // The version of the tool used to generate the code scanning analysis.
	Guid string `json:"guid,omitempty"` // The GUID of the tool used to generate the code scanning analysis, if provided in the uploaded SARIF data.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	In_reply_to_id int `json:"in_reply_to_id,omitempty"`
	Subject_type string `json:"subject_type,omitempty"` // The level at which the comment is targeted, can be a diff line or a file.
	Original_start_line int `json:"original_start_line,omitempty"` // The original first line of the range for a multi-line comment.
	Body_text string `json:"body_text,omitempty"`
	Html_url string `json:"html_url"`
	Position int `json:"position"`
	Id int64 `json:"id"`
	Diff_hunk string `json:"diff_hunk"`
	Original_position int `json:"original_position"`
	Side string `json:"side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Pull_request_review_id int64 `json:"pull_request_review_id"`
	Updated_at string `json:"updated_at"`
	Start_side string `json:"start_side,omitempty"` // The side of the first line of the range for a multi-line comment.
	Links map[string]interface{} `json:"_links"`
	Body string `json:"body"`
	Pull_request_url string `json:"pull_request_url"`
	Original_commit_id string `json:"original_commit_id"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Path string `json:"path"`
	Commit_id string `json:"commit_id"`
	Node_id string `json:"node_id"`
	Start_line int `json:"start_line,omitempty"` // The first line of the range for a multi-line comment.
	Url string `json:"url"`
	Created_at string `json:"created_at"`
	Original_line int `json:"original_line,omitempty"` // The original line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Reactions GeneratedType `json:"reactions,omitempty"`
	User GeneratedType `json:"user"` // A GitHub user.
	Line int `json:"line,omitempty"` // The line of the blob to which the comment applies. The last line of the range for a multi-line comment
	Body_html string `json:"body_html,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Marketplace_purchase map[string]interface{} `json:"marketplace_purchase"`
	Previous_marketplace_purchase Webhookspreviousmarketplacepurchase `json:"previous_marketplace_purchase,omitempty"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Effective_date string `json:"effective_date"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Content_type string `json:"content_type,omitempty"` // The media type used to serialize the payloads. Supported values include `json` and `form`. The default is `form`.
	Insecure_ssl GeneratedType `json:"insecure_ssl,omitempty"`
	Secret string `json:"secret,omitempty"` // If provided, the `secret` will be used as the `key` to generate the HMAC hex digest value for [delivery signature headers](https://docs.github.com/webhooks/event-payloads/#delivery-headers).
	Url string `json:"url,omitempty"` // The URL to which the payloads will be delivered.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Milestone Webhooksmilestone `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Check_run GeneratedType `json:"check_run"` // A check performed on the code of a given code change
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Requested_action map[string]interface{} `json:"requested_action,omitempty"` // The action requested by the user.
}

// Webhooksrule represents the Webhooksrule schema from the OpenAPI specification
type Webhooksrule struct {
	Required_conversation_resolution_level string `json:"required_conversation_resolution_level"`
	Admin_enforced bool `json:"admin_enforced"`
	Repository_id int `json:"repository_id"`
	Require_last_push_approval bool `json:"require_last_push_approval,omitempty"` // Whether the most recent push must be approved by someone other than the person who pushed it
	Strict_required_status_checks_policy bool `json:"strict_required_status_checks_policy"`
	Required_status_checks []string `json:"required_status_checks"`
	Lock_allows_fork_sync bool `json:"lock_allows_fork_sync,omitempty"` // Whether users can pull changes from upstream when the branch is locked. Set to `true` to allow users to pull changes from upstream when the branch is locked. This setting is only applicable for forks.
	Create_protected bool `json:"create_protected,omitempty"`
	Ignore_approvals_from_contributors bool `json:"ignore_approvals_from_contributors"`
	Allow_deletions_enforcement_level string `json:"allow_deletions_enforcement_level"`
	Id int `json:"id"`
	Name string `json:"name"`
	Pull_request_reviews_enforcement_level string `json:"pull_request_reviews_enforcement_level"`
	Lock_branch_enforcement_level string `json:"lock_branch_enforcement_level"` // The enforcement level of the branch lock setting. `off` means the branch is not locked, `non_admins` means the branch is read-only for non_admins, and `everyone` means the branch is read-only for everyone.
	Required_approving_review_count int `json:"required_approving_review_count"`
	Authorized_actor_names []string `json:"authorized_actor_names"`
	Required_status_checks_enforcement_level string `json:"required_status_checks_enforcement_level"`
	Dismiss_stale_reviews_on_push bool `json:"dismiss_stale_reviews_on_push"`
	Signature_requirement_enforcement_level string `json:"signature_requirement_enforcement_level"`
	Allow_force_pushes_enforcement_level string `json:"allow_force_pushes_enforcement_level"`
	Merge_queue_enforcement_level string `json:"merge_queue_enforcement_level"`
	Authorized_actors_only bool `json:"authorized_actors_only"`
	Authorized_dismissal_actors_only bool `json:"authorized_dismissal_actors_only"`
	Linear_history_requirement_enforcement_level string `json:"linear_history_requirement_enforcement_level"`
	Created_at string `json:"created_at"`
	Require_code_owner_review bool `json:"require_code_owner_review"`
	Required_deployments_enforcement_level string `json:"required_deployments_enforcement_level"`
	Updated_at string `json:"updated_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repositories []GeneratedType `json:"repositories"` // A list of repositories that were skipped. This list may not include all repositories that were skipped. This is only available when the repository was found and the user has access to it.
	Repository_count int `json:"repository_count"` // The total number of repositories that were skipped for this reason.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repos_url string `json:"repos_url"`
	TypeField string `json:"type"`
	Id int `json:"id"`
	Url string `json:"url"`
	Inherited_from []GeneratedType `json:"inherited_from,omitempty"` // Team the user has gotten the role through
	Name string `json:"name,omitempty"`
	Organizations_url string `json:"organizations_url"`
	Starred_at string `json:"starred_at,omitempty"`
	Site_admin bool `json:"site_admin"`
	Html_url string `json:"html_url"`
	Following_url string `json:"following_url"`
	Gravatar_id string `json:"gravatar_id"`
	Received_events_url string `json:"received_events_url"`
	Node_id string `json:"node_id"`
	Email string `json:"email,omitempty"`
	User_view_type string `json:"user_view_type,omitempty"`
	Assignment string `json:"assignment,omitempty"` // Determines if the user has a direct, indirect, or mixed relationship to a role
	Events_url string `json:"events_url"`
	Followers_url string `json:"followers_url"`
	Gists_url string `json:"gists_url"`
	Login string `json:"login"`
	Avatar_url string `json:"avatar_url"`
	Starred_url string `json:"starred_url"`
	Subscriptions_url string `json:"subscriptions_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Environment_url string `json:"environment_url,omitempty"` // The URL for accessing your environment.
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Id int64 `json:"id"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Url string `json:"url"`
	Repository_url string `json:"repository_url"`
	State string `json:"state"` // The state of the status.
	Updated_at string `json:"updated_at"`
	Deployment_url string `json:"deployment_url"`
	Environment string `json:"environment,omitempty"` // The environment of the deployment that the status is for.
	Node_id string `json:"node_id"`
	Log_url string `json:"log_url,omitempty"` // The URL to associate with this status.
	Target_url string `json:"target_url"` // Closing down notice: the URL to associate with this status.
	Created_at string `json:"created_at"`
	Description string `json:"description"` // A short description of the status.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Parent_issue Issue `json:"parent_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Parent_issue_id float64 `json:"parent_issue_id"` // The ID of the parent issue.
	Sub_issue Issue `json:"sub_issue"` // Issues are a great way to keep track of tasks, enhancements, and bugs for your projects.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Sub_issue_id float64 `json:"sub_issue_id"` // The ID of the sub-issue.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Parent_issue_repo Repository `json:"parent_issue_repo"` // A repository on GitHub.
	Repository GeneratedType `json:"repository,omitempty"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Action string `json:"action"`
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Workflow_job interface{} `json:"workflow_job"`
	Action string `json:"action"`
	Deployment Deployment `json:"deployment,omitempty"` // A request for a specific ref(branch,sha,tag) to be deployed
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pusher_type string `json:"pusher_type"` // The pusher type for the event. Can be either `user` or a deploy key.
	Ref string `json:"ref"` // The [`git ref`](https://docs.github.com/rest/git/refs#get-a-reference) resource.
	Ref_type string `json:"ref_type"` // The type of Git ref object deleted in the repository.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Merged bool `json:"merged"`
	Message string `json:"message"`
	Sha string `json:"sha"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Selected_actions_url string `json:"selected_actions_url,omitempty"` // The API URL to use to get or set the actions and reusable workflows that are allowed to run, when `allowed_actions` is set to `selected`.
	Selected_repositories_url string `json:"selected_repositories_url,omitempty"` // The API URL to use to get or set the selected repositories that are allowed to run GitHub Actions, when `enabled_repositories` is set to `selected`.
	Allowed_actions string `json:"allowed_actions,omitempty"` // The permissions policy that controls the actions and reusable workflows that are allowed to run.
	Enabled_repositories string `json:"enabled_repositories"` // The policy that controls the repositories in the organization that are allowed to run GitHub Actions.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Errors []map[string]interface{} `json:"errors,omitempty"`
	Message string `json:"message"`
	Documentation_url string `json:"documentation_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Cvss_v3 map[string]interface{} `json:"cvss_v3,omitempty"`
	Cvss_v4 map[string]interface{} `json:"cvss_v4,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Label Webhookslabel `json:"label"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender,omitempty"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Due_on string `json:"due_on"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Url string `json:"url"`
	State string `json:"state"` // The state of the milestone.
	Title string `json:"title"` // The title of the milestone.
	Closed_issues int `json:"closed_issues"`
	Labels_url string `json:"labels_url"`
	Number int `json:"number"` // The number of the milestone.
	Updated_at string `json:"updated_at"`
	Id int `json:"id"`
	Open_issues int `json:"open_issues"`
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Html_url string `json:"html_url"`
	Closed_at string `json:"closed_at"`
}

// Verification represents the Verification schema from the OpenAPI specification
type Verification struct {
	Verified_at string `json:"verified_at"`
	Payload string `json:"payload"`
	Reason string `json:"reason"`
	Signature string `json:"signature"`
	Verified bool `json:"verified"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pages string `json:"pages,omitempty"` // The level of permission to grant the access token to retrieve Pages statuses, configuration, and builds, as well as create new builds.
	Organization_custom_roles string `json:"organization_custom_roles,omitempty"` // The level of permission to grant the access token for custom repository roles management.
	Metadata string `json:"metadata,omitempty"` // The level of permission to grant the access token to search repositories, list collaborators, and access repository metadata.
	Organization_personal_access_token_requests string `json:"organization_personal_access_token_requests,omitempty"` // The level of permission to grant the access token for viewing and managing fine-grained personal access tokens that have been approved by an organization.
	Organization_projects string `json:"organization_projects,omitempty"` // The level of permission to grant the access token to manage organization projects and projects public preview (where available).
	Repository_hooks string `json:"repository_hooks,omitempty"` // The level of permission to grant the access token to manage the post-receive hooks for a repository.
	Security_events string `json:"security_events,omitempty"` // The level of permission to grant the access token to view and manage security events like code scanning alerts.
	Checks string `json:"checks,omitempty"` // The level of permission to grant the access token for checks on code.
	Packages string `json:"packages,omitempty"` // The level of permission to grant the access token for packages published to GitHub Packages.
	Team_discussions string `json:"team_discussions,omitempty"` // The level of permission to grant the access token to manage team discussions and related comments.
	Issues string `json:"issues,omitempty"` // The level of permission to grant the access token for issues and related comments, assignees, labels, and milestones.
	Contents string `json:"contents,omitempty"` // The level of permission to grant the access token for repository contents, commits, branches, downloads, releases, and merges.
	Codespaces string `json:"codespaces,omitempty"` // The level of permission to grant the access token to create, edit, delete, and list Codespaces.
	Organization_administration string `json:"organization_administration,omitempty"` // The level of permission to grant the access token to manage access to an organization.
	Workflows string `json:"workflows,omitempty"` // The level of permission to grant the access token to update GitHub Actions workflow files.
	Starring string `json:"starring,omitempty"` // The level of permission to grant the access token to list and manage repositories a user is starring.
	Repository_custom_properties string `json:"repository_custom_properties,omitempty"` // The level of permission to grant the access token to view and edit custom properties for a repository, when allowed by the property.
	Gpg_keys string `json:"gpg_keys,omitempty"` // The level of permission to grant the access token to view and manage GPG keys belonging to a user.
	Members string `json:"members,omitempty"` // The level of permission to grant the access token for organization teams and members.
	Secret_scanning_alerts string `json:"secret_scanning_alerts,omitempty"` // The level of permission to grant the access token to view and manage secret scanning alerts.
	Organization_custom_org_roles string `json:"organization_custom_org_roles,omitempty"` // The level of permission to grant the access token for custom organization roles management.
	Administration string `json:"administration,omitempty"` // The level of permission to grant the access token for repository creation, deletion, settings, teams, and collaborators creation.
	Organization_hooks string `json:"organization_hooks,omitempty"` // The level of permission to grant the access token to manage the post-receive hooks for an organization.
	Pull_requests string `json:"pull_requests,omitempty"` // The level of permission to grant the access token for pull requests and related comments, assignees, labels, milestones, and merges.
	Single_file string `json:"single_file,omitempty"` // The level of permission to grant the access token to manage just a single file.
	Organization_plan string `json:"organization_plan,omitempty"` // The level of permission to grant the access token for viewing an organization's plan.
	Followers string `json:"followers,omitempty"` // The level of permission to grant the access token to manage the followers belonging to a user.
	Organization_announcement_banners string `json:"organization_announcement_banners,omitempty"` // The level of permission to grant the access token to view and manage announcement banners for an organization.
	Organization_user_blocking string `json:"organization_user_blocking,omitempty"` // The level of permission to grant the access token to view and manage users blocked by the organization.
	Repository_projects string `json:"repository_projects,omitempty"` // The level of permission to grant the access token to manage repository projects, columns, and cards.
	Actions string `json:"actions,omitempty"` // The level of permission to grant the access token for GitHub Actions workflows, workflow runs, and artifacts.
	Secrets string `json:"secrets,omitempty"` // The level of permission to grant the access token to manage repository secrets.
	Organization_personal_access_tokens string `json:"organization_personal_access_tokens,omitempty"` // The level of permission to grant the access token for viewing and managing fine-grained personal access token requests to an organization.
	Environments string `json:"environments,omitempty"` // The level of permission to grant the access token for managing repository environments.
	Organization_self_hosted_runners string `json:"organization_self_hosted_runners,omitempty"` // The level of permission to grant the access token to view and manage GitHub Actions self-hosted runners available to an organization.
	Deployments string `json:"deployments,omitempty"` // The level of permission to grant the access token for deployments and deployment statuses.
	Profile string `json:"profile,omitempty"` // The level of permission to grant the access token to manage the profile settings belonging to a user.
	Organization_packages string `json:"organization_packages,omitempty"` // The level of permission to grant the access token for organization packages published to GitHub Packages.
	Git_ssh_keys string `json:"git_ssh_keys,omitempty"` // The level of permission to grant the access token to manage git SSH keys.
	Interaction_limits string `json:"interaction_limits,omitempty"` // The level of permission to grant the access token to view and manage interaction limits on a repository.
	Statuses string `json:"statuses,omitempty"` // The level of permission to grant the access token for commit statuses.
	Email_addresses string `json:"email_addresses,omitempty"` // The level of permission to grant the access token to manage the email addresses belonging to a user.
	Organization_events string `json:"organization_events,omitempty"` // The level of permission to grant the access token to view events triggered by an activity in an organization.
	Dependabot_secrets string `json:"dependabot_secrets,omitempty"` // The level of permission to grant the access token to manage Dependabot secrets.
	Organization_copilot_seat_management string `json:"organization_copilot_seat_management,omitempty"` // The level of permission to grant the access token for managing access to GitHub Copilot for members of an organization with a Copilot Business subscription. This property is in public preview and is subject to change.
	Organization_custom_properties string `json:"organization_custom_properties,omitempty"` // The level of permission to grant the access token for custom property management.
	Vulnerability_alerts string `json:"vulnerability_alerts,omitempty"` // The level of permission to grant the access token to manage Dependabot alerts.
	Organization_secrets string `json:"organization_secrets,omitempty"` // The level of permission to grant the access token to manage organization secrets.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Url string `json:"url"` // The URL to the workflow run.
	Status string `json:"status"`
	Head_sha string `json:"head_sha"` // The SHA of the head commit that points to the version of the workflow being run.
	Display_title string `json:"display_title"` // The event-specific title associated with the run or the run-name if set, or the value of `run-name` if it is set in the workflow.
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Event string `json:"event"`
	Html_url string `json:"html_url"`
	Head_repository_id int `json:"head_repository_id,omitempty"`
	Run_started_at string `json:"run_started_at,omitempty"` // The start time of the latest run. Resets on re-run.
	Previous_attempt_url string `json:"previous_attempt_url,omitempty"` // The URL to the previous attempted run of this workflow, if one exists.
	Head_commit GeneratedType `json:"head_commit"` // A commit.
	Triggering_actor GeneratedType `json:"triggering_actor,omitempty"` // A GitHub user.
	Logs_url string `json:"logs_url"` // The URL to download the logs for the workflow run.
	Head_repository GeneratedType `json:"head_repository"` // Minimal Repository
	Check_suite_node_id string `json:"check_suite_node_id,omitempty"` // The node ID of the associated check suite.
	Workflow_id int `json:"workflow_id"` // The ID of the parent workflow.
	Referenced_workflows []GeneratedType `json:"referenced_workflows,omitempty"`
	Workflow_url string `json:"workflow_url"` // The URL to the workflow.
	Check_suite_id int `json:"check_suite_id,omitempty"` // The ID of the associated check suite.
	Head_branch string `json:"head_branch"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
	Run_attempt int `json:"run_attempt,omitempty"` // Attempt number of the run, 1 for first attempt and higher if the workflow was re-run.
	Actor GeneratedType `json:"actor,omitempty"` // A GitHub user.
	Path string `json:"path"` // The full path of the workflow
	Id int `json:"id"` // The ID of the workflow run.
	Jobs_url string `json:"jobs_url"` // The URL to the jobs for the workflow run.
	Run_number int `json:"run_number"` // The auto incrementing run number for the workflow run.
	Pull_requests []GeneratedType `json:"pull_requests"` // Pull requests that are open with a `head_sha` or `head_branch` that matches the workflow run. The returned pull requests do not necessarily indicate pull requests that triggered the run.
	Cancel_url string `json:"cancel_url"` // The URL to cancel the workflow run.
	Check_suite_url string `json:"check_suite_url"` // The URL to the associated check suite.
	Conclusion string `json:"conclusion"`
	Created_at string `json:"created_at"`
	Name string `json:"name,omitempty"` // The name of the workflow run.
	Artifacts_url string `json:"artifacts_url"` // The URL to the artifacts for the workflow run.
	Rerun_url string `json:"rerun_url"` // The URL to rerun the workflow run.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	User GeneratedType `json:"user"` // A GitHub user.
	Sub_issues_summary map[string]interface{} `json:"sub_issues_summary,omitempty"`
	Locked bool `json:"locked"`
	State string `json:"state"`
	Title string `json:"title"`
	Id int64 `json:"id"`
	Number int `json:"number"`
	Active_lock_reason string `json:"active_lock_reason,omitempty"`
	State_reason string `json:"state_reason,omitempty"`
	Text_matches []map[string]interface{} `json:"text_matches,omitempty"`
	Assignees []GeneratedType `json:"assignees,omitempty"`
	Html_url string `json:"html_url"`
	Closed_at string `json:"closed_at"`
	Draft bool `json:"draft,omitempty"`
	Events_url string `json:"events_url"`
	Author_association string `json:"author_association"` // How the author is associated with the repository.
	Score float64 `json:"score"`
	TypeField GeneratedType `json:"type,omitempty"` // The type of issue.
	Comments int `json:"comments"`
	Body_text string `json:"body_text,omitempty"`
	Labels_url string `json:"labels_url"`
	Timeline_url string `json:"timeline_url,omitempty"`
	Comments_url string `json:"comments_url"`
	Updated_at string `json:"updated_at"`
	Performed_via_github_app GeneratedType `json:"performed_via_github_app,omitempty"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Reactions GeneratedType `json:"reactions,omitempty"`
	Pull_request map[string]interface{} `json:"pull_request,omitempty"`
	Labels []map[string]interface{} `json:"labels"`
	Repository_url string `json:"repository_url"`
	Created_at string `json:"created_at"`
	Url string `json:"url"`
	Milestone GeneratedType `json:"milestone"` // A collection of related issues and pull requests.
	Node_id string `json:"node_id"`
	Repository Repository `json:"repository,omitempty"` // A repository on GitHub.
	Assignee GeneratedType `json:"assignee"` // A GitHub user.
	Body_html string `json:"body_html,omitempty"`
	Body string `json:"body,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id"`
	Updated_at string `json:"updated_at"`
	Column_name string `json:"column_name,omitempty"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
	Created_at string `json:"created_at"`
	Project_url string `json:"project_url"`
	Note string `json:"note"`
	Project_id string `json:"project_id,omitempty"`
	Url string `json:"url"`
	Archived bool `json:"archived,omitempty"` // Whether or not the card is archived
	Column_url string `json:"column_url"`
	Content_url string `json:"content_url,omitempty"`
	Id int64 `json:"id"` // The project card's ID
}

// Import represents the Import schema from the OpenAPI specification
type Import struct {
	Authors_url string `json:"authors_url"`
	Error_message string `json:"error_message,omitempty"`
	Project_choices []map[string]interface{} `json:"project_choices,omitempty"`
	Svc_root string `json:"svc_root,omitempty"`
	Authors_count int `json:"authors_count,omitempty"`
	Message string `json:"message,omitempty"`
	Url string `json:"url"`
	Html_url string `json:"html_url"`
	Large_files_size int `json:"large_files_size,omitempty"`
	Commit_count int `json:"commit_count,omitempty"`
	Failed_step string `json:"failed_step,omitempty"`
	Use_lfs bool `json:"use_lfs,omitempty"`
	Vcs_url string `json:"vcs_url"` // The URL of the originating repository.
	Tfvc_project string `json:"tfvc_project,omitempty"`
	Repository_url string `json:"repository_url"`
	Has_large_files bool `json:"has_large_files,omitempty"`
	Import_percent int `json:"import_percent,omitempty"`
	Large_files_count int `json:"large_files_count,omitempty"`
	Svn_root string `json:"svn_root,omitempty"`
	Vcs string `json:"vcs"`
	Status_text string `json:"status_text,omitempty"`
	Status string `json:"status"`
	Push_percent int `json:"push_percent,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Key Webhooksdeploykey `json:"key"` // The [`deploy key`](https://docs.github.com/rest/deploy-keys/deploy-keys#get-a-deploy-key) resource.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Issue map[string]interface{} `json:"issue"` // The [issue](https://docs.github.com/rest/issues/issues#get-an-issue) itself.
	Milestone Webhooksmilestone `json:"milestone"` // A collection of related issues and pull requests.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Visibility string `json:"visibility"` // The type of repositories in the organization that the secret is visible to
	Created_at string `json:"created_at"` // The date and time at which the secret was created, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
	Name string `json:"name"` // The name of the secret
	Selected_repositories_url string `json:"selected_repositories_url"` // The API URL at which the list of repositories this secret is visible to can be retrieved
	Updated_at string `json:"updated_at"` // The date and time at which the secret was last updated, in ISO 8601 format':' YYYY-MM-DDTHH:MM:SSZ.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Warning string `json:"warning"` // Warning generated when processing the analysis
	Sarif_id string `json:"sarif_id"` // An identifier for the upload.
	Results_count int `json:"results_count"` // The total number of results in the analysis.
	Url string `json:"url"` // The REST API URL of the analysis resource.
	Deletable bool `json:"deletable"`
	Id int `json:"id"` // Unique identifier for this analysis.
	Category string `json:"category,omitempty"` // Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code.
	Commit_sha string `json:"commit_sha"` // The SHA of the commit to which the analysis you are uploading relates.
	Tool GeneratedType `json:"tool"`
	Rules_count int `json:"rules_count"` // The total number of rules used in the analysis.
	Environment string `json:"environment"` // Identifies the variable values associated with the environment in which this analysis was performed.
	Created_at string `json:"created_at"` // The time that the analysis was created in ISO 8601 format: `YYYY-MM-DDTHH:MM:SSZ`.
	ErrorField string `json:"error"`
	Ref string `json:"ref"` // The Git reference, formatted as `refs/pull/<number>/merge`, `refs/pull/<number>/head`, `refs/heads/<branch name>` or simply `<branch name>`.
	Analysis_key string `json:"analysis_key"` // Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Milestone Webhooksmilestone `json:"milestone"` // A collection of related issues and pull requests.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ruleset_source string `json:"ruleset_source,omitempty"` // The name of the source of the ruleset that includes this rule.
	Ruleset_source_type string `json:"ruleset_source_type,omitempty"` // The type of source for the ruleset that includes this rule.
	Ruleset_id int `json:"ruleset_id,omitempty"` // The ID of the ruleset that includes this rule.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Grade string `json:"grade"` // Most recent grade.
	Id int `json:"id"` // Unique identifier of the repository.
	Passing bool `json:"passing"` // Whether a submission passed.
	Repository GeneratedType `json:"repository"` // A GitHub repository view for Classroom
	Students []GeneratedType `json:"students"`
	Submitted bool `json:"submitted"` // Whether an accepted assignment has been submitted.
	Assignment GeneratedType `json:"assignment"` // A GitHub Classroom assignment
	Commit_count int `json:"commit_count"` // Count of student commits.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Actions_caches []map[string]interface{} `json:"actions_caches"` // Array of caches
	Total_count int `json:"total_count"` // Total number of caches
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Registry_type string `json:"registry_type"` // The registry type.
	Selected_repository_ids []int `json:"selected_repository_ids,omitempty"` // An array of repository IDs that can access the organization private registry when `visibility` is set to `selected`.
	Updated_at string `json:"updated_at"`
	Username string `json:"username,omitempty"` // The username to use when authenticating with the private registry.
	Visibility string `json:"visibility"` // Which type of organization repositories have access to the private registry. `selected` means only the repositories specified by `selected_repository_ids` can access the private registry.
	Created_at string `json:"created_at"`
	Name string `json:"name"` // The name of the private registry configuration.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Member Webhooksuser `json:"member"`
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Preferences map[string]interface{} `json:"preferences"`
	Repository GeneratedType `json:"repository"` // Minimal Repository
}

// Webhooksworkflowjobrun represents the Webhooksworkflowjobrun schema from the OpenAPI specification
type Webhooksworkflowjobrun struct {
	Environment string `json:"environment"`
	Html_url string `json:"html_url"`
	Id int `json:"id"`
	Name interface{} `json:"name"`
	Status string `json:"status"`
	Updated_at string `json:"updated_at"`
	Conclusion interface{} `json:"conclusion"`
	Created_at string `json:"created_at"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Owner GeneratedType `json:"owner,omitempty"` // A GitHub user.
	Created_at string `json:"created_at"`
	History []interface{} `json:"history,omitempty"`
	Id string `json:"id"`
	Updated_at string `json:"updated_at"`
	Git_push_url string `json:"git_push_url"`
	Forks_url string `json:"forks_url"`
	Public bool `json:"public"`
	Description string `json:"description"`
	Forks []interface{} `json:"forks,omitempty"`
	Node_id string `json:"node_id"`
	Comments_enabled bool `json:"comments_enabled,omitempty"`
	Commits_url string `json:"commits_url"`
	Files map[string]interface{} `json:"files"`
	Truncated bool `json:"truncated,omitempty"`
	Git_pull_url string `json:"git_pull_url"`
	Comments int `json:"comments"`
	User GeneratedType `json:"user"` // A GitHub user.
	Html_url string `json:"html_url"`
	Url string `json:"url"`
	Comments_url string `json:"comments_url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Node_id string `json:"node_id,omitempty"`
	Requester GeneratedType `json:"requester"` // A GitHub user.
	Account interface{} `json:"account"`
	Created_at string `json:"created_at"`
	Id int `json:"id"` // Unique identifier of the request installation.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Number int `json:"number"` // The pull request number.
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Label Webhookslabel `json:"label,omitempty"`
	Pull_request map[string]interface{} `json:"pull_request"`
	Action string `json:"action"`
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Title string `json:"title"` // Assignment title.
	Invite_link string `json:"invite_link"` // The link that a student can use to accept the assignment.
	Passing int `json:"passing"` // The number of students that have passed the assignment.
	Max_members int `json:"max_members,omitempty"` // The maximum allowable members per team.
	Public_repo bool `json:"public_repo"` // Whether an accepted assignment creates a public repository.
	Id int `json:"id"` // Unique identifier of the repository.
	Submitted int `json:"submitted"` // The number of students that have submitted the assignment.
	Students_are_repo_admins bool `json:"students_are_repo_admins"` // Whether students are admins on created repository on accepted assignment.
	Accepted int `json:"accepted"` // The number of students that have accepted the assignment.
	TypeField string `json:"type"` // Whether it's a Group Assignment or Individual Assignment.
	Feedback_pull_requests_enabled bool `json:"feedback_pull_requests_enabled"` // Whether feedback pull request will be created on assignment acceptance.
	Deadline string `json:"deadline"` // The time at which the assignment is due.
	Classroom GeneratedType `json:"classroom"` // A GitHub Classroom classroom
	Editor string `json:"editor"` // The selected editor for the assignment.
	Invitations_enabled bool `json:"invitations_enabled"` // Whether the invitation link is enabled. Visiting an enabled invitation link will accept the assignment.
	Language string `json:"language"` // The programming language used in the assignment.
	Slug string `json:"slug"` // Sluggified name of the assignment.
	Max_teams int `json:"max_teams,omitempty"` // The maximum allowable teams for the assignment.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Alert GeneratedType `json:"alert"` // A Dependabot alert.
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Commit_id string `json:"commit_id"`
	Commit_url string `json:"commit_url"`
	Created_at string `json:"created_at"`
	Event string `json:"event"`
	Id int `json:"id"`
	Performed_via_github_app Integration `json:"performed_via_github_app"` // GitHub apps are a new way to extend GitHub. They can be installed directly on organizations and user accounts and granted access to specific repositories. They come with granular permissions and built-in webhooks. GitHub apps are first class actors within GitHub.
	Actor GeneratedType `json:"actor"` // A GitHub user.
	Node_id string `json:"node_id"`
	Project_card map[string]interface{} `json:"project_card,omitempty"`
	Url string `json:"url"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Repository GeneratedType `json:"repository"` // The repository on GitHub where the event occurred. Webhook payloads contain the `repository` property when the event occurs from activity in a repository.
	Rule Webhooksrule `json:"rule"` // The branch protection rule. Includes a `name` and all the [branch protection settings](https://docs.github.com/github/administering-a-repository/defining-the-mergeability-of-pull-requests/about-protected-branches#about-branch-protection-settings) applied to branches that match the name. Binary settings are boolean. Multi-level configurations are one of `off`, `non_admins`, or `everyone`. Actor and build lists are arrays of strings.
	Sender GeneratedType `json:"sender"` // A GitHub user.
	Action string `json:"action"`
	Enterprise GeneratedType `json:"enterprise,omitempty"` // An enterprise on GitHub. Webhook payloads contain the `enterprise` property when the webhook is configured on an enterprise account or an organization that's part of an enterprise account. For more information, see "[About enterprise accounts](https://docs.github.com/admin/overview/about-enterprise-accounts)."
	Installation GeneratedType `json:"installation,omitempty"` // The GitHub App installation. Webhook payloads contain the `installation` property when the event is configured for and sent to a GitHub App. For more information, see "[Using webhooks with GitHub Apps](https://docs.github.com/apps/creating-github-apps/registering-a-github-app/using-webhooks-with-github-apps)."
	Organization GeneratedType `json:"organization,omitempty"` // A GitHub organization. Webhook payloads contain the `organization` property when the webhook is configured for an organization, or when the event occurs from activity in a repository owned by an organization.
}

// Status represents the Status schema from the OpenAPI specification
type Status struct {
	State string `json:"state"`
	Id int `json:"id"`
	Created_at string `json:"created_at"`
	Description string `json:"description"`
	Node_id string `json:"node_id"`
	Target_url string `json:"target_url"`
	Updated_at string `json:"updated_at"`
	Url string `json:"url"`
	Avatar_url string `json:"avatar_url"`
	Context string `json:"context"`
	Creator GeneratedType `json:"creator"` // A GitHub user.
}
