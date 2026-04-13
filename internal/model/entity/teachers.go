// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Teachers is the golang structure for table teachers.
type Teachers struct {
	Id        int64       `json:"id"        orm:"id"         description:"Primary key"`                   // Primary key
	GoogleSub string      `json:"googleSub" orm:"google_sub" description:"Google unique user identifier"` // Google unique user identifier
	Email     string      `json:"email"     orm:"email"      description:"Google email"`                  // Google email
	Name      string      `json:"name"      orm:"name"       description:"Display name"`                  // Display name
	AvatarUrl string      `json:"avatarUrl" orm:"avatar_url" description:"Avatar URL"`                    // Avatar URL
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"Created at"`                    // Created at
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"Updated at"`                    // Updated at
}
