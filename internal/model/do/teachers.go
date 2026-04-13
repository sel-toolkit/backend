// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Teachers is the golang structure of table teachers for DAO operations like Where/Data.
type Teachers struct {
	g.Meta    `orm:"table:teachers, do:true"`
	Id        interface{} // Primary key
	GoogleSub interface{} // Google unique user identifier
	Email     interface{} // Google email
	Name      interface{} // Display name
	AvatarUrl interface{} // Avatar URL
	CreatedAt *gtime.Time // Created at
	UpdatedAt *gtime.Time // Updated at
}
