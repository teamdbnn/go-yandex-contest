package models

type ExtendedRole struct {
	Metas     []*RoleMeta `json:"metas"`
	Principal *Principal  `json:"principal"`
	Resource  *Resource   `json:"resource,omitempty"`
	Role      *Role       `json:"role"`
}
