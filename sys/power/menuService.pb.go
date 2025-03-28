// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: menuService.proto

package services

import (
	common "github.com/geiqin/micro-kit/protobuf/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 导航
type Menu struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                                   //ID
	ParentId   int64   `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`       //父ID
	Title      string  `protobuf:"bytes,3,opt,name=title,proto3" json:"title"`                              //导航标题
	Type       string  `protobuf:"bytes,4,opt,name=type,proto3" json:"type"`                                //类型 1目录 2菜单 3按钮
	Path       string  `protobuf:"bytes,5,opt,name=path,proto3" json:"path"`                                //路由地址
	Name       string  `protobuf:"bytes,6,opt,name=name,proto3" json:"name"`                                //导航名称
	Component  string  `protobuf:"bytes,7,opt,name=component,proto3" json:"component"`                      //组件路径
	Redirect   string  `protobuf:"bytes,8,opt,name=redirect,proto3" json:"redirect"`                        //路由重定向
	Icon       string  `protobuf:"bytes,9,opt,name=icon,proto3" json:"icon"`                                //图标
	Permission string  `protobuf:"bytes,10,opt,name=permission,proto3" json:"permission"`                   //权限标识
	Locale     string  `protobuf:"bytes,11,opt,name=locale,proto3" json:"locale"`                           //语言包键名
	IsCache    string  `protobuf:"bytes,12,opt,name=is_cache,json=isCache,proto3" json:"is_cache"`          //是否缓存: 1是 0否
	IsHidden   string  `protobuf:"bytes,13,opt,name=is_hidden,json=isHidden,proto3" json:"is_hidden"`       //是否隐藏: 1是 0否
	IsExternal string  `protobuf:"bytes,14,opt,name=is_external,json=isExternal,proto3" json:"is_external"` //是否外链: 1是 0否
	IsPlugin   string  `protobuf:"bytes,15,opt,name=is_plugin,json=isPlugin,proto3" json:"is_plugin"`       //是否插件: 1是 0否
	PluginId   int64   `protobuf:"varint,16,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id"`      //插件ID
	Sort       int32   `protobuf:"varint,17,opt,name=sort,proto3" json:"sort"`                              //排序值
	Status     string  `protobuf:"bytes,18,opt,name=status,proto3" json:"status"`                           //状态: （1：启用；2：禁用）
	CreatedAt  int64   `protobuf:"varint,20,opt,name=created_at,json=createdAt,proto3" json:"created_at"`
	UpdatedAt  int64   `protobuf:"varint,21,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at"`
	Parent     *Menu   `protobuf:"bytes,22,opt,name=parent,proto3" json:"parent"`
	Children   []*Menu `protobuf:"bytes,23,rep,name=children,proto3" json:"children"`
}

func (x *Menu) Reset() {
	*x = Menu{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menuService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Menu) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Menu) ProtoMessage() {}

func (x *Menu) ProtoReflect() protoreflect.Message {
	mi := &file_menuService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Menu.ProtoReflect.Descriptor instead.
func (*Menu) Descriptor() ([]byte, []int) {
	return file_menuService_proto_rawDescGZIP(), []int{0}
}

func (x *Menu) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Menu) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *Menu) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Menu) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Menu) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Menu) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Menu) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *Menu) GetRedirect() string {
	if x != nil {
		return x.Redirect
	}
	return ""
}

func (x *Menu) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Menu) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *Menu) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Menu) GetIsCache() string {
	if x != nil {
		return x.IsCache
	}
	return ""
}

func (x *Menu) GetIsHidden() string {
	if x != nil {
		return x.IsHidden
	}
	return ""
}

func (x *Menu) GetIsExternal() string {
	if x != nil {
		return x.IsExternal
	}
	return ""
}

func (x *Menu) GetIsPlugin() string {
	if x != nil {
		return x.IsPlugin
	}
	return ""
}

func (x *Menu) GetPluginId() int64 {
	if x != nil {
		return x.PluginId
	}
	return 0
}

func (x *Menu) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *Menu) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Menu) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Menu) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Menu) GetParent() *Menu {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *Menu) GetChildren() []*Menu {
	if x != nil {
		return x.Children
	}
	return nil
}

type RouteItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id"`                  //ID
	ParentId   int64        `protobuf:"varint,2,opt,name=parentId,proto3" json:"parentId"`      //父ID
	Title      string       `protobuf:"bytes,3,opt,name=title,proto3" json:"title"`             //导航标题
	Type       int32        `protobuf:"varint,4,opt,name=type,proto3" json:"type"`              //类型 1目录 2菜单 3按钮
	Path       string       `protobuf:"bytes,5,opt,name=path,proto3" json:"path"`               //路由地址
	Name       string       `protobuf:"bytes,6,opt,name=name,proto3" json:"name"`               //导航名称
	Component  string       `protobuf:"bytes,7,opt,name=component,proto3" json:"component"`     //组件路径
	Redirect   string       `protobuf:"bytes,8,opt,name=redirect,proto3" json:"redirect"`       //路由重定向
	Icon       string       `protobuf:"bytes,9,opt,name=icon,proto3" json:"icon"`               //图标
	Permission string       `protobuf:"bytes,10,opt,name=permission,proto3" json:"permission"`  //权限标识
	Locale     string       `protobuf:"bytes,11,opt,name=locale,proto3" json:"locale"`          //语言包键名
	IsCache    bool         `protobuf:"varint,12,opt,name=isCache,proto3" json:"isCache"`       //是否缓存: 1是 0否
	IsHidden   bool         `protobuf:"varint,13,opt,name=isHidden,proto3" json:"isHidden"`     //是否隐藏: 1是 0否
	IsExternal bool         `protobuf:"varint,14,opt,name=isExternal,proto3" json:"isExternal"` //是否外链: 1是 0否
	Sort       int32        `protobuf:"varint,15,opt,name=sort,proto3" json:"sort"`             //排序值
	Status     string       `protobuf:"bytes,16,opt,name=status,proto3" json:"status"`          //状态: （1：启用；2：禁用）
	ActiveMenu string       `protobuf:"bytes,17,opt,name=activeMenu,proto3" json:"activeMenu"`  //
	AlwaysShow bool         `protobuf:"varint,18,opt,name=alwaysShow,proto3" json:"alwaysShow"` //
	Breadcrumb bool         `protobuf:"varint,19,opt,name=breadcrumb,proto3" json:"breadcrumb"` //
	ShowInTabs bool         `protobuf:"varint,20,opt,name=showInTabs,proto3" json:"showInTabs"` //
	Affix      bool         `protobuf:"varint,21,opt,name=affix,proto3" json:"affix"`           //
	Roles      []string     `protobuf:"bytes,22,rep,name=roles,proto3" json:"roles"`
	Children   []*RouteItem `protobuf:"bytes,23,rep,name=children,proto3" json:"children"`
}

func (x *RouteItem) Reset() {
	*x = RouteItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menuService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteItem) ProtoMessage() {}

func (x *RouteItem) ProtoReflect() protoreflect.Message {
	mi := &file_menuService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteItem.ProtoReflect.Descriptor instead.
func (*RouteItem) Descriptor() ([]byte, []int) {
	return file_menuService_proto_rawDescGZIP(), []int{1}
}

func (x *RouteItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RouteItem) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *RouteItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RouteItem) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *RouteItem) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *RouteItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RouteItem) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *RouteItem) GetRedirect() string {
	if x != nil {
		return x.Redirect
	}
	return ""
}

func (x *RouteItem) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *RouteItem) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *RouteItem) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *RouteItem) GetIsCache() bool {
	if x != nil {
		return x.IsCache
	}
	return false
}

func (x *RouteItem) GetIsHidden() bool {
	if x != nil {
		return x.IsHidden
	}
	return false
}

func (x *RouteItem) GetIsExternal() bool {
	if x != nil {
		return x.IsExternal
	}
	return false
}

func (x *RouteItem) GetSort() int32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *RouteItem) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *RouteItem) GetActiveMenu() string {
	if x != nil {
		return x.ActiveMenu
	}
	return ""
}

func (x *RouteItem) GetAlwaysShow() bool {
	if x != nil {
		return x.AlwaysShow
	}
	return false
}

func (x *RouteItem) GetBreadcrumb() bool {
	if x != nil {
		return x.Breadcrumb
	}
	return false
}

func (x *RouteItem) GetShowInTabs() bool {
	if x != nil {
		return x.ShowInTabs
	}
	return false
}

func (x *RouteItem) GetAffix() bool {
	if x != nil {
		return x.Affix
	}
	return false
}

func (x *RouteItem) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *RouteItem) GetChildren() []*RouteItem {
	if x != nil {
		return x.Children
	}
	return nil
}

type MenuRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paged    int64  `protobuf:"varint,1,opt,name=paged,proto3" json:"paged"`
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size"`
	Sorting  string `protobuf:"bytes,3,opt,name=sorting,proto3" json:"sorting"`
	//base params
	Id         int64   `protobuf:"varint,4,opt,name=id,proto3" json:"id"`
	Name       string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	Title      string  `protobuf:"bytes,6,opt,name=title,proto3" json:"title"`
	Type       string  `protobuf:"bytes,7,opt,name=type,proto3" json:"type"`
	ParentId   int64   `protobuf:"varint,8,opt,name=parent_id,json=parentId,proto3" json:"parent_id"`
	IsHidden   int64   `protobuf:"varint,9,opt,name=is_hidden,json=isHidden,proto3" json:"is_hidden"`
	System     string  `protobuf:"bytes,10,opt,name=system,proto3" json:"system"`
	IsPlugin   string  `protobuf:"bytes,11,opt,name=is_plugin,json=isPlugin,proto3" json:"is_plugin"`
	PluginId   int32   `protobuf:"varint,12,opt,name=plugin_id,json=pluginId,proto3" json:"plugin_id"`
	StoreId    int64   `protobuf:"varint,13,opt,name=store_id,json=storeId,proto3" json:"store_id"`
	PluginIds  []int32 `protobuf:"varint,14,rep,packed,name=plugin_ids,json=pluginIds,proto3" json:"plugin_ids"`
	NavTypeIds []int64 `protobuf:"varint,15,rep,packed,name=nav_type_ids,json=navTypeIds,proto3" json:"nav_type_ids"`
	Ids        []int64 `protobuf:"varint,16,rep,packed,name=ids,proto3" json:"ids"`
}

func (x *MenuRequest) Reset() {
	*x = MenuRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menuService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuRequest) ProtoMessage() {}

func (x *MenuRequest) ProtoReflect() protoreflect.Message {
	mi := &file_menuService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuRequest.ProtoReflect.Descriptor instead.
func (*MenuRequest) Descriptor() ([]byte, []int) {
	return file_menuService_proto_rawDescGZIP(), []int{2}
}

func (x *MenuRequest) GetPaged() int64 {
	if x != nil {
		return x.Paged
	}
	return 0
}

func (x *MenuRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *MenuRequest) GetSorting() string {
	if x != nil {
		return x.Sorting
	}
	return ""
}

func (x *MenuRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MenuRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MenuRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MenuRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MenuRequest) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *MenuRequest) GetIsHidden() int64 {
	if x != nil {
		return x.IsHidden
	}
	return 0
}

func (x *MenuRequest) GetSystem() string {
	if x != nil {
		return x.System
	}
	return ""
}

func (x *MenuRequest) GetIsPlugin() string {
	if x != nil {
		return x.IsPlugin
	}
	return ""
}

func (x *MenuRequest) GetPluginId() int32 {
	if x != nil {
		return x.PluginId
	}
	return 0
}

func (x *MenuRequest) GetStoreId() int64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *MenuRequest) GetPluginIds() []int32 {
	if x != nil {
		return x.PluginIds
	}
	return nil
}

func (x *MenuRequest) GetNavTypeIds() []int64 {
	if x != nil {
		return x.NavTypeIds
	}
	return nil
}

func (x *MenuRequest) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type MenuResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity *Menu         `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity"`
	Pager  *common.Pager `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager"`
	Items  []*Menu       `protobuf:"bytes,3,rep,name=items,proto3" json:"items"`
	Routes []*RouteItem  `protobuf:"bytes,4,rep,name=routes,proto3" json:"routes"`
	Tree   []*TreeItem   `protobuf:"bytes,5,rep,name=tree,proto3" json:"tree"`
	Msg    string        `protobuf:"bytes,6,opt,name=msg,proto3" json:"msg"`
}

func (x *MenuResponse) Reset() {
	*x = MenuResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_menuService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuResponse) ProtoMessage() {}

func (x *MenuResponse) ProtoReflect() protoreflect.Message {
	mi := &file_menuService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuResponse.ProtoReflect.Descriptor instead.
func (*MenuResponse) Descriptor() ([]byte, []int) {
	return file_menuService_proto_rawDescGZIP(), []int{3}
}

func (x *MenuResponse) GetEntity() *Menu {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *MenuResponse) GetPager() *common.Pager {
	if x != nil {
		return x.Pager
	}
	return nil
}

func (x *MenuResponse) GetItems() []*Menu {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *MenuResponse) GetRoutes() []*RouteItem {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *MenuResponse) GetTree() []*TreeItem {
	if x != nil {
		return x.Tree
	}
	return nil
}

func (x *MenuResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_menuService_proto protoreflect.FileDescriptor

var file_menuService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x65, 0x6e, 0x75, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x11, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x04, 0x0a, 0x04, 0x4d, 0x65, 0x6e, 0x75,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x73, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x73, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x69, 0x73, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x73, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65,
	0x6e, 0x75, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0xee, 0x04, 0x0a, 0x09, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x63, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x43, 0x61,
	0x63, 0x68, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x69, 0x73, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x6c,
	0x77, 0x61, 0x79, 0x73, 0x53, 0x68, 0x6f, 0x77, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x61, 0x6c, 0x77, 0x61, 0x79, 0x73, 0x53, 0x68, 0x6f, 0x77, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x72,
	0x65, 0x61, 0x64, 0x63, 0x72, 0x75, 0x6d, 0x62, 0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x62, 0x72, 0x65, 0x61, 0x64, 0x63, 0x72, 0x75, 0x6d, 0x62, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x68,
	0x6f, 0x77, 0x49, 0x6e, 0x54, 0x61, 0x62, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x73, 0x68, 0x6f, 0x77, 0x49, 0x6e, 0x54, 0x61, 0x62, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x66,
	0x66, 0x69, 0x78, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x66, 0x66, 0x69, 0x78,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x16, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72,
	0x65, 0x6e, 0x18, 0x17, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x08, 0x63,
	0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22, 0xa2, 0x03, 0x0a, 0x0b, 0x4d, 0x65, 0x6e, 0x75,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x68, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x69, 0x73, 0x48, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x73, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x6e, 0x61, 0x76, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a,
	0x6e, 0x61, 0x76, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xe8, 0x01, 0x0a,
	0x0c, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x72, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x2b, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x26, 0x0a,
	0x04, 0x74, 0x72, 0x65, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x54, 0x72, 0x65, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x04, 0x74, 0x72, 0x65, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x84, 0x05, 0x0a, 0x0b, 0x4d, 0x65, 0x6e, 0x75,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x1a, 0x16,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65,
	0x6e, 0x75, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65,
	0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x36, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x69, 0x64, 0x65, 0x12, 0x0e,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x1a, 0x16,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65,
	0x6e, 0x75, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65,
	0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x04, 0x54, 0x72, 0x65, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d,
	0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x10, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x4d, 0x65,
	0x6e, 0x75, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x10, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x4d, 0x65, 0x6e, 0x75, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x15, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x4d, 0x65, 0x6e, 0x75, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d,
	0x5a, 0x0b, 0x2f, 0x2e, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_menuService_proto_rawDescOnce sync.Once
	file_menuService_proto_rawDescData = file_menuService_proto_rawDesc
)

func file_menuService_proto_rawDescGZIP() []byte {
	file_menuService_proto_rawDescOnce.Do(func() {
		file_menuService_proto_rawDescData = protoimpl.X.CompressGZIP(file_menuService_proto_rawDescData)
	})
	return file_menuService_proto_rawDescData
}

var file_menuService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_menuService_proto_goTypes = []interface{}{
	(*Menu)(nil),         // 0: services.Menu
	(*RouteItem)(nil),    // 1: services.RouteItem
	(*MenuRequest)(nil),  // 2: services.MenuRequest
	(*MenuResponse)(nil), // 3: services.MenuResponse
	(*common.Pager)(nil), // 4: common.Pager
	(*TreeItem)(nil),     // 5: services.TreeItem
}
var file_menuService_proto_depIdxs = []int32{
	0,  // 0: services.Menu.parent:type_name -> services.Menu
	0,  // 1: services.Menu.children:type_name -> services.Menu
	1,  // 2: services.RouteItem.children:type_name -> services.RouteItem
	0,  // 3: services.MenuResponse.entity:type_name -> services.Menu
	4,  // 4: services.MenuResponse.pager:type_name -> common.Pager
	0,  // 5: services.MenuResponse.items:type_name -> services.Menu
	1,  // 6: services.MenuResponse.routes:type_name -> services.RouteItem
	5,  // 7: services.MenuResponse.tree:type_name -> services.TreeItem
	0,  // 8: services.MenuService.Get:input_type -> services.Menu
	0,  // 9: services.MenuService.Create:input_type -> services.Menu
	0,  // 10: services.MenuService.Update:input_type -> services.Menu
	0,  // 11: services.MenuService.UpdateHide:input_type -> services.Menu
	0,  // 12: services.MenuService.Delete:input_type -> services.Menu
	2,  // 13: services.MenuService.Search:input_type -> services.MenuRequest
	2,  // 14: services.MenuService.List:input_type -> services.MenuRequest
	2,  // 15: services.MenuService.Tree:input_type -> services.MenuRequest
	2,  // 16: services.MenuService.Routes:input_type -> services.MenuRequest
	2,  // 17: services.MenuService.PluginMenuCreate:input_type -> services.MenuRequest
	2,  // 18: services.MenuService.PluginMenuDelete:input_type -> services.MenuRequest
	3,  // 19: services.MenuService.Get:output_type -> services.MenuResponse
	3,  // 20: services.MenuService.Create:output_type -> services.MenuResponse
	3,  // 21: services.MenuService.Update:output_type -> services.MenuResponse
	3,  // 22: services.MenuService.UpdateHide:output_type -> services.MenuResponse
	3,  // 23: services.MenuService.Delete:output_type -> services.MenuResponse
	3,  // 24: services.MenuService.Search:output_type -> services.MenuResponse
	3,  // 25: services.MenuService.List:output_type -> services.MenuResponse
	3,  // 26: services.MenuService.Tree:output_type -> services.MenuResponse
	3,  // 27: services.MenuService.Routes:output_type -> services.MenuResponse
	3,  // 28: services.MenuService.PluginMenuCreate:output_type -> services.MenuResponse
	3,  // 29: services.MenuService.PluginMenuDelete:output_type -> services.MenuResponse
	19, // [19:30] is the sub-list for method output_type
	8,  // [8:19] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_menuService_proto_init() }
func file_menuService_proto_init() {
	if File_menuService_proto != nil {
		return
	}
	file_baseInfoService_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_menuService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Menu); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_menuService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_menuService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_menuService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_menuService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_menuService_proto_goTypes,
		DependencyIndexes: file_menuService_proto_depIdxs,
		MessageInfos:      file_menuService_proto_msgTypes,
	}.Build()
	File_menuService_proto = out.File
	file_menuService_proto_rawDesc = nil
	file_menuService_proto_goTypes = nil
	file_menuService_proto_depIdxs = nil
}
