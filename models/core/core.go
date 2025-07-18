// Package core provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package core

import (
	"time"

	"github.com/gofrs/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// AcceptedTermsAt defines model for accepted_terms_at.
type AcceptedTermsAt = string

// AvatarUrl Link for profile picture
type AvatarUrl = string

// Bio defines model for bio.
type Bio = string

// CreatedAt Timestamp when the resource was created.
type CreatedAt = time.Time

// CredentialID defines model for credential_uuid.
type CredentialID = uuid.UUID

// DeletedAt Timestamp when the resource was deleted.
type DeletedAt = time.Time

// DesignId defines model for design_id.
type DesignId = uuid.UUID

// Email email
type Email = openapi_types.Email

// EmailPreference defines model for email_preference.
type EmailPreference struct {
	NotifyRoleChange bool `json:"notify_role_change,omitempty"`
	WelcomeEmail     bool `json:"welcome_email,omitempty"`
}

// Emails defines model for emails.
type Emails = []Email

// Empty Body for empty request
type Empty = map[string]interface{}

// Endpoint defines model for endpoint.
type Endpoint = Text

// EnvironmentId defines model for environment_id.
type EnvironmentId = uuid.UUID

// GeneralId defines model for general_id.
type GeneralId = uuid.UUID

// Id defines model for id.
type Id = uuid.UUID

// KubernetesServerID defines model for kubernetes_server_uuid.
type KubernetesServerID = uuid.UUID

// MapObject defines model for map_object.
type MapObject map[string]string

// MesheryInstanceID defines model for meshery_instance_uuid.
type MesheryInstanceID = uuid.UUID

// NOTE:
// this is not in use
// and I assumed that this file is not autogenerated by anu script now
// commented this out to avoid name collision 
//
// // NullTime SQL null Timestamp to handle null values of time.
// type NullTime = sql.NullTime

// Number defines model for number.
type Number = int

// OperationID defines model for operation_id.
type OperationID = uuid.UUID

// OrganizationId defines model for organization_id.
type OrganizationId = uuid.UUID

// Price defines model for price.
type Price = int32

// Provider One of (x-oapi-codegen-extra-tags-cloud, github, google)
type Provider = string

// RecordsPage defines model for recordsPage.
type RecordsPage struct {
	Page         int    `json:"page,omitempty"`
	PageSize     int    `json:"page_size,omitempty"`
	RecordType   string `json:"recordType,omitempty"`
	RecordsTotal int    `json:"records_total,omitempty"`
}

// ResultsPage defines model for resultsPage.
type ResultsPage struct {
	Page       int    `json:"page,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	ResultType string `json:"resultType,omitempty"`
	TotalCount int    `json:"total_count,omitempty"`
}

// RoleNames defines model for roleNames.
type RoleNames = []string

// Status defines model for status.
type Status = string

// SystemID defines model for system_id.
type SystemID = uuid.UUID

// TeamId defines model for team_id.
type TeamId = uuid.UUID

// Text defines model for text.
type Text = string

// Time defines model for time.
type Time = time.Time

// UpdatedAt Timestamp when the resource was updated.
type UpdatedAt = time.Time

// UserId user's email or username
type UserId = string

// UserIds defines model for user_ids.
type UserIds = []Id

// UserID defines model for user_uuid.
type UserID = uuid.UUID

// Username defines model for username.
type Username = string

// ViewId defines model for view_id.
type ViewId = uuid.UUID

// WorkspaceId defines model for workspace_id.
type WorkspaceId = uuid.UUID
