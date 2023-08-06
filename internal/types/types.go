// Code generated by goctl. DO NOT EDIT.
package types

type DoLoginReq struct {
	SecretKey string `form:"secret_key,optional"`
}

type StartContainerReq struct {
	Name string `json:"name"`
}

type StopContainerReq struct {
	Name string `json:"name"`
}

type RenameContainerReq struct {
	OldName string `json:"oldName"`
	NewName string `json:"newName"`
}

type CreateContainerReq struct {
	OldName         string `json:"old_name"`
	NewName         string `json:"new_name"`
	ImageNameAndTag string `json:"image_name_and_tag"`
}

type RemoveContainerReq struct {
	Name string `json:"name"`
}

type GetNewImageReq struct {
	ImageNameAndTag string `json:"image_name_and_tag"`
}

type MsgResp struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}
