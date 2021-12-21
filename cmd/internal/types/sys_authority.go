package types

type SysAuthorityResponse struct {
	Authority SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      SysAuthority `json:"authority"`
	OldAuthorityId string       `json:"oldAuthorityId"` // 旧角色ID
}
